package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Client talks to the AnkiConnect add-on over its local JSON-RPC endpoint.
type Client struct {
	URL  string
	http *http.Client
}

func NewClient(url string) *Client {
	// AnkiConnect closes the connection after each response, so a reused
	// keep-alive socket fails mid-run ("connection reset by peer").
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.DisableKeepAlives = true
	return &Client{URL: url, http: &http.Client{Timeout: 60 * time.Second, Transport: tr}}
}

// Note is one AnkiConnect note payload.
type Note struct {
	DeckName  string            `json:"deckName"`
	ModelName string            `json:"modelName"`
	Fields    map[string]string `json:"fields"`
	Tags      []string          `json:"tags"`
	Options   map[string]any    `json:"options"`
}

type request struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  any    `json:"params,omitempty"`
}

type response struct {
	Result json.RawMessage `json:"result"`
	Error  *string         `json:"error"`
}

// call runs one AnkiConnect action and decodes its result into out (may be nil).
func (c *Client) call(action string, params, out any) error {
	body, err := json.Marshal(request{Action: action, Version: 6, Params: params})
	if err != nil {
		return err
	}

	resp, err := c.http.Post(c.URL, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("AnkiConnect unreachable at %s on %s (is Anki running with the AnkiConnect add-on?): %w", c.URL, action, err)
	}
	defer resp.Body.Close()

	var r response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}
	if r.Error != nil {
		return fmt.Errorf("AnkiConnect error on %s: %s", action, *r.Error)
	}
	if out == nil {
		return nil
	}
	return json.Unmarshal(r.Result, out)
}

// ─── decks ───

func (c *Client) CreateDeck(name string) error {
	return c.call("createDeck", map[string]any{"deck": name}, nil)
}

// ─── notes ───

// CanAdd reports, per note, whether it would be a new note. addNotes rejects
// the whole batch when it contains a duplicate, so callers filter first.
func (c *Client) CanAdd(notes []Note) ([]bool, error) {
	var ok []bool
	for _, batch := range chunk(notes) {
		var part []bool
		if err := c.call("canAddNotes", map[string]any{"notes": batch}, &part); err != nil {
			return nil, err
		}
		ok = append(ok, part...)
	}
	return ok, nil
}

// batchSize caps how many notes go into one AnkiConnect request — Anki closes
// the connection on very large payloads.
const batchSize = 100

func chunk(notes []Note) [][]Note {
	var out [][]Note
	for i := 0; i < len(notes); i += batchSize {
		end := i + batchSize
		if end > len(notes) {
			end = len(notes)
		}
		out = append(out, notes[i:end])
	}
	return out
}

// AddNotes creates the given notes and returns how many were actually added.
func (c *Client) AddNotes(notes []Note) (int, error) {
	if len(notes) == 0 {
		return 0, nil
	}
	added := 0
	for _, batch := range chunk(notes) {
		var ids []*int64
		if err := c.call("addNotes", map[string]any{"notes": batch}, &ids); err != nil {
			return added, err
		}
		for _, id := range ids {
			if id != nil {
				added++
			}
		}
	}
	return added, nil
}

// UpdateNote rewrites the fields of the existing note sharing this front field.
func (c *Client) UpdateNote(n Note, frontField string) error {
	query := fmt.Sprintf(`deck:"%s" "%s:%s"`, n.DeckName, frontField, escapeQuery(n.Fields[frontField]))

	var found []int64
	if err := c.call("findNotes", map[string]any{"query": query}, &found); err != nil {
		return err
	}
	if len(found) == 0 {
		return fmt.Errorf("no existing note matched")
	}
	return c.call("updateNoteFields", map[string]any{
		"note": map[string]any{"id": found[0], "fields": n.Fields},
	}, nil)
}

// escapeQuery quotes the characters that are operators in Anki's search syntax.
func escapeQuery(s string) string {
	r := strings.NewReplacer(`\`, `\\`, `"`, `\"`, `*`, `\*`, `_`, `\_`)
	return r.Replace(s)
}

// ─── note types ───

// Fields names the two fields a card is written into.
type Fields struct {
	Front string
	Back  string
}

// EnsureModel creates the note type when it is missing, refreshes the styling of
// one this tool owns, and reports which fields to write the card into.
func (c *Client) EnsureModel(model string) (Fields, error) {
	var models []string
	if err := c.call("modelNames", nil, &models); err != nil {
		return Fields{}, err
	}

	if !contains(models, model) {
		if err := c.createModel(model); err != nil {
			return Fields{}, err
		}
		fmt.Fprintf(out, "🆕 Created note type %q\n", model)
		return Fields{Front: "Front", Back: "Back"}, nil
	}

	var fields []string
	if err := c.call("modelFieldNames", map[string]any{"modelName": model}, &fields); err != nil {
		return Fields{}, err
	}
	if len(fields) < 2 {
		return Fields{}, fmt.Errorf("note type %q has fewer than 2 fields", model)
	}

	// A note type of our own shape: refresh its CSS and templates so styling
	// changes reach cards imported earlier.
	if contains(fields, "Front") && contains(fields, "Back") {
		if err := c.restyleModel(model, fields); err != nil {
			return Fields{}, err
		}
		return Fields{Front: "Front", Back: "Back"}, nil
	}

	// Someone else's note type — write into its first two fields, untouched.
	return Fields{Front: fields[0], Back: fields[1]}, nil
}

func (c *Client) createModel(model string) error {
	return c.call("createModel", map[string]any{
		"modelName":     model,
		"inOrderFields": []string{"Front", "Back"},
		"isCloze":       false,
		"css":           cardCSS,
		"cardTemplates": []map[string]string{{
			"Name":  cardTemplateName,
			"Front": frontTemplate,
			"Back":  backTemplate,
		}},
	}, nil)
}

// restyleModel updates the CSS and templates of a note type this tool owns.
// Note types built by hand (or by another tool) are left alone — ours are
// recognised by the template name, or by their plain two-field shape if they
// were created by an earlier version of this tool.
func (c *Client) restyleModel(model string, fields []string) error {
	var templates map[string]map[string]string
	if err := c.call("modelTemplates", map[string]any{"modelName": model}, &templates); err != nil {
		return err
	}

	name, ok := ownedTemplate(templates, fields)
	if !ok {
		return nil
	}

	if err := c.call("updateModelStyling", map[string]any{
		"model": map[string]any{"name": model, "css": cardCSS},
	}, nil); err != nil {
		return err
	}
	return c.call("updateModelTemplates", map[string]any{
		"model": map[string]any{
			"name": model,
			"templates": map[string]any{
				name: map[string]string{"Front": frontTemplate, "Back": backTemplate},
			},
		},
	}, nil)
}

// ownedTemplate returns the template this tool may overwrite, if any.
func ownedTemplate(templates map[string]map[string]string, fields []string) (string, bool) {
	if _, ok := templates[cardTemplateName]; ok {
		return cardTemplateName, true
	}
	if len(fields) != 2 || len(templates) != 1 {
		return "", false
	}

	name := ""
	for only := range templates {
		name = only
	}
	return name, true
}

func contains(list []string, want string) bool {
	for _, s := range list {
		if s == want {
			return true
		}
	}
	return false
}
