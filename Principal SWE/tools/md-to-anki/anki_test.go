package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// fakeAnki is a stand-in for the AnkiConnect add-on. Each handler returns the
// JSON result for one action; every request is recorded for assertions.
type fakeAnki struct {
	t        *testing.T
	handlers map[string]func(params map[string]any) (any, error)
	calls    []request
	server   *httptest.Server
}

func newFakeAnki(t *testing.T) *fakeAnki {
	t.Helper()
	f := &fakeAnki{t: t, handlers: map[string]func(map[string]any) (any, error){}}
	f.server = httptest.NewServer(http.HandlerFunc(f.serve))
	t.Cleanup(f.server.Close)
	return f
}

func (f *fakeAnki) on(action string, fn func(params map[string]any) (any, error)) *fakeAnki {
	f.handlers[action] = fn
	return f
}

// reply registers a handler that always returns the same result.
func (f *fakeAnki) reply(action string, result any) *fakeAnki {
	return f.on(action, func(map[string]any) (any, error) { return result, nil })
}

func (f *fakeAnki) serve(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	var req request
	if err := json.Unmarshal(body, &req); err != nil {
		f.t.Fatalf("bad request body: %v", err)
	}
	f.calls = append(f.calls, req)

	params, _ := req.Params.(map[string]any)
	handler, ok := f.handlers[req.Action]
	if !ok {
		writeJSON(w, map[string]any{"result": nil, "error": "unsupported action: " + req.Action})
		return
	}

	result, err := handler(params)
	if err != nil {
		writeJSON(w, map[string]any{"result": nil, "error": err.Error()})
		return
	}
	writeJSON(w, map[string]any{"result": result, "error": nil})
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func (f *fakeAnki) client() *Client { return NewClient(f.server.URL) }

// actions lists the actions the fake received, in order.
func (f *fakeAnki) actions() []string {
	var names []string
	for _, c := range f.calls {
		names = append(names, c.Action)
	}
	return names
}

func (f *fakeAnki) paramsOf(action string) map[string]any {
	for _, c := range f.calls {
		if c.Action == action {
			p, _ := c.Params.(map[string]any)
			return p
		}
	}
	f.t.Fatalf("action %q was never called (got %v)", action, f.actions())
	return nil
}

// ─── transport ───

func TestCallReportsAnkiConnectErrors(t *testing.T) {
	fake := newFakeAnki(t)
	err := fake.client().call("createDeck", map[string]any{"deck": "X"}, nil)

	if err == nil || !strings.Contains(err.Error(), "unsupported action") {
		t.Fatalf("err = %v, want the add-on's error message", err)
	}
}

func TestCallReportsAnUnreachableAddon(t *testing.T) {
	c := NewClient("http://127.0.0.1:1") // nothing listens here
	err := c.call("version", nil, nil)

	if err == nil || !strings.Contains(err.Error(), "AnkiConnect unreachable") {
		t.Fatalf("err = %v, want an unreachable-endpoint error", err)
	}
}

func TestCallDecodesTheResult(t *testing.T) {
	fake := newFakeAnki(t).reply("deckNames", []string{"Default", "Go"})

	var decks []string
	if err := fake.client().call("deckNames", nil, &decks); err != nil {
		t.Fatal(err)
	}
	if len(decks) != 2 || decks[1] != "Go" {
		t.Errorf("decks = %v", decks)
	}
}

// ─── decks and notes ───

func TestCreateDeck(t *testing.T) {
	fake := newFakeAnki(t).reply("createDeck", 1)

	if err := fake.client().CreateDeck("Principal SWE::Go"); err != nil {
		t.Fatal(err)
	}
	if got := fake.paramsOf("createDeck")["deck"]; got != "Principal SWE::Go" {
		t.Errorf("deck = %v", got)
	}
}

func TestCanAdd(t *testing.T) {
	fake := newFakeAnki(t).reply("canAddNotes", []bool{true, false})

	got, err := fake.client().CanAdd([]Note{{}, {}})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || !got[0] || got[1] {
		t.Errorf("CanAdd = %v, want [true false]", got)
	}
}

func TestAddNotesCountsOnlyTheAcceptedOnes(t *testing.T) {
	fake := newFakeAnki(t).reply("addNotes", []any{1234, nil, 5678})

	added, err := fake.client().AddNotes([]Note{{}, {}, {}})
	if err != nil {
		t.Fatal(err)
	}
	if added != 2 {
		t.Errorf("added = %d, want 2 (a null id means the note was rejected)", added)
	}
}

func TestAddNotesSkipsTheCallWhenThereIsNothingToAdd(t *testing.T) {
	fake := newFakeAnki(t)

	added, err := fake.client().AddNotes(nil)
	if err != nil || added != 0 {
		t.Fatalf("added = %d, err = %v", added, err)
	}
	if len(fake.calls) != 0 {
		t.Errorf("an empty batch still called %v", fake.actions())
	}
}

func TestAddNotesPropagatesErrors(t *testing.T) {
	fake := newFakeAnki(t).on("addNotes", func(map[string]any) (any, error) {
		return nil, errString("cannot create note because it is a duplicate")
	})

	if _, err := fake.client().AddNotes([]Note{{}}); err == nil {
		t.Fatal("want an error")
	}
}

func TestUpdateNoteFindsThenWrites(t *testing.T) {
	fake := newFakeAnki(t).
		reply("findNotes", []int64{99}).
		reply("updateNoteFields", nil)

	note := Note{
		DeckName: "Go",
		Fields:   map[string]string{"Front": `Chan "x" *y*`, "Back": "body"},
	}
	if err := fake.client().UpdateNote(note, "Front"); err != nil {
		t.Fatal(err)
	}

	query, _ := fake.paramsOf("findNotes")["query"].(string)
	if !strings.Contains(query, `deck:"Go"`) || !strings.Contains(query, `\"x\"`) || !strings.Contains(query, `\*y\*`) {
		t.Errorf("query = %q — deck scope or escaping is wrong", query)
	}

	updated := fake.paramsOf("updateNoteFields")["note"].(map[string]any)
	if updated["id"].(float64) != 99 {
		t.Errorf("updated the wrong note: %v", updated["id"])
	}
}

func TestUpdateNoteWhenNothingMatches(t *testing.T) {
	fake := newFakeAnki(t).reply("findNotes", []int64{})

	err := fake.client().UpdateNote(Note{Fields: map[string]string{"Front": "x"}}, "Front")
	if err == nil || !strings.Contains(err.Error(), "no existing note matched") {
		t.Fatalf("err = %v", err)
	}
}

func TestUpdateNotePropagatesFindErrors(t *testing.T) {
	fake := newFakeAnki(t)

	if err := fake.client().UpdateNote(Note{Fields: map[string]string{"Front": "x"}}, "Front"); err == nil {
		t.Fatal("want an error when findNotes fails")
	}
}

// ─── note types ───

func TestEnsureModelCreatesTheNoteTypeWhenMissing(t *testing.T) {
	fake := newFakeAnki(t).
		reply("modelNames", []string{"Basic"}).
		reply("createModel", map[string]any{})

	fields, err := fake.client().EnsureModel("Principal SWE")
	if err != nil {
		t.Fatal(err)
	}
	if fields != (Fields{Front: "Front", Back: "Back"}) {
		t.Errorf("fields = %+v", fields)
	}

	params := fake.paramsOf("createModel")
	if params["modelName"] != "Principal SWE" {
		t.Errorf("modelName = %v", params["modelName"])
	}
	css, _ := params["css"].(string)
	if !strings.Contains(css, ".tok-kw") {
		t.Error("the created note type does not carry the highlighter styling")
	}
	templates := params["cardTemplates"].([]any)
	if name := templates[0].(map[string]any)["Name"]; name != cardTemplateName {
		t.Errorf("template name = %v, want %q", name, cardTemplateName)
	}
}

func TestEnsureModelRestylesOurOwnNoteType(t *testing.T) {
	fake := newFakeAnki(t).
		reply("modelNames", []string{"Principal SWE"}).
		reply("modelFieldNames", []string{"Front", "Back"}).
		reply("modelTemplates", map[string]any{cardTemplateName: map[string]any{}}).
		reply("updateModelStyling", nil).
		reply("updateModelTemplates", nil)

	if _, err := fake.client().EnsureModel("Principal SWE"); err != nil {
		t.Fatal(err)
	}

	got := strings.Join(fake.actions(), ",")
	if !strings.Contains(got, "updateModelStyling") || !strings.Contains(got, "updateModelTemplates") {
		t.Errorf("our note type was not restyled: %v", fake.actions())
	}
}

func TestEnsureModelLeavesAForeignNoteTypeAlone(t *testing.T) {
	fake := newFakeAnki(t).
		reply("modelNames", []string{"Go Roadmap"}).
		reply("modelFieldNames", []string{"Front", "Back", "Example", "Output"}).
		reply("modelTemplates", map[string]any{"Card 1": map[string]any{}, "Card 2": map[string]any{}})

	fields, err := fake.client().EnsureModel("Go Roadmap")
	if err != nil {
		t.Fatal(err)
	}
	if fields != (Fields{Front: "Front", Back: "Back"}) {
		t.Errorf("fields = %+v", fields)
	}
	for _, a := range fake.actions() {
		if strings.HasPrefix(a, "updateModel") {
			t.Errorf("a foreign note type was modified: %v", fake.actions())
		}
	}
}

func TestEnsureModelUsesTheFirstTwoFieldsOfAnUnrelatedNoteType(t *testing.T) {
	fake := newFakeAnki(t).
		reply("modelNames", []string{"LeetCode"}).
		reply("modelFieldNames", []string{"Title", "Kind", "Puzzle"})

	fields, err := fake.client().EnsureModel("LeetCode")
	if err != nil {
		t.Fatal(err)
	}
	if fields != (Fields{Front: "Title", Back: "Kind"}) {
		t.Errorf("fields = %+v, want the first two fields", fields)
	}
}

func TestEnsureModelRejectsASingleFieldNoteType(t *testing.T) {
	fake := newFakeAnki(t).
		reply("modelNames", []string{"OneField"}).
		reply("modelFieldNames", []string{"Text"})

	if _, err := fake.client().EnsureModel("OneField"); err == nil ||
		!strings.Contains(err.Error(), "fewer than 2 fields") {
		t.Fatalf("err = %v", err)
	}
}

func TestEnsureModelPropagatesErrors(t *testing.T) {
	t.Run("modelNames fails", func(t *testing.T) {
		fake := newFakeAnki(t)
		if _, err := fake.client().EnsureModel("X"); err == nil {
			t.Fatal("want an error")
		}
	})

	t.Run("createModel fails", func(t *testing.T) {
		fake := newFakeAnki(t).reply("modelNames", []string{})
		if _, err := fake.client().EnsureModel("X"); err == nil {
			t.Fatal("want an error")
		}
	})

	t.Run("modelFieldNames fails", func(t *testing.T) {
		fake := newFakeAnki(t).reply("modelNames", []string{"X"})
		if _, err := fake.client().EnsureModel("X"); err == nil {
			t.Fatal("want an error")
		}
	})

	t.Run("modelTemplates fails", func(t *testing.T) {
		fake := newFakeAnki(t).
			reply("modelNames", []string{"X"}).
			reply("modelFieldNames", []string{"Front", "Back"})
		if _, err := fake.client().EnsureModel("X"); err == nil {
			t.Fatal("want an error")
		}
	})

	t.Run("updateModelStyling fails", func(t *testing.T) {
		fake := newFakeAnki(t).
			reply("modelNames", []string{"X"}).
			reply("modelFieldNames", []string{"Front", "Back"}).
			reply("modelTemplates", map[string]any{cardTemplateName: map[string]any{}})
		if _, err := fake.client().EnsureModel("X"); err == nil {
			t.Fatal("want an error")
		}
	})
}

func TestContains(t *testing.T) {
	if !contains([]string{"a", "b"}, "b") {
		t.Error("contains missed a present value")
	}
	if contains([]string{"a"}, "z") {
		t.Error("contains found an absent value")
	}
}

type errString string

func (e errString) Error() string { return string(e) }
