package main

import (
	"fmt"
	"html"
	"os"
	"strings"
)

// WritePreviewHTML renders the cards into a single standalone page using the
// exact CSS and templates Anki will use, so the result can be checked in a
// browser before importing.
func WritePreviewHTML(path string, cards []Card, model string) error {
	var b strings.Builder

	b.WriteString(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>md-to-anki preview</title>
<style>
body { margin: 0; background: #0b0d11; }
.wrap { max-width: 860px; margin: 0 auto; padding: 32px 20px 60px; }
.head { font-family: -apple-system, Helvetica, sans-serif; color: #7f8ea6; font-size: 13px;
        letter-spacing: .06em; text-transform: uppercase; margin: 0 0 20px; }
.card { border: 1px solid #232a36; border-radius: 12px; margin: 0 0 22px; }
.meta { font-family: ui-monospace, Menlo, monospace; font-size: 11px; color: #5c6c7d;
        padding: 8px 26px 0; }
</style>
<style>
`)
	b.WriteString(cardCSS)
	b.WriteString("</style>\n</head>\n<body>\n<div class=\"wrap\">\n")
	fmt.Fprintf(&b, "<p class=\"head\">%d cards · note type %s</p>\n", len(cards), html.EscapeString(model))

	for _, c := range cards {
		fmt.Fprintf(&b, "<div class=\"meta\">%s</div>\n", html.EscapeString(c.Deck))
		b.WriteString(`<div class="card">`)
		fmt.Fprintf(&b, `<div class="q question">%s</div>`, html.EscapeString(c.Front))
		b.WriteString(`<hr id="answer">`)
		fmt.Fprintf(&b, `<div class="a answer-body">%s</div>`, c.Back)
		b.WriteString("</div>\n")
	}

	// The card templates address one element by id; on this page every card
	// carries the same class instead, so the scripts run over all of them.
	b.WriteString("</div>\n<script>\n")
	b.WriteString(previewScript())
	b.WriteString("\n</script>\n</body>\n</html>\n")

	return os.WriteFile(path, []byte(b.String()), 0o644)
}

// previewScript reuses the template scripts by feeding them each card in turn.
func previewScript() string {
	return `
document.querySelectorAll(".question").forEach(function (el) {
  el.id = "question";
` + stripScriptTags(frontTemplate) + `
  el.removeAttribute("id");
});
document.querySelectorAll(".answer-body").forEach(function (el) {
  el.id = "answer-body";
` + stripScriptTags(backTemplate) + `
  el.removeAttribute("id");
});
`
}

// stripScriptTags pulls the JavaScript out of a card template.
func stripScriptTags(tpl string) string {
	start := strings.Index(tpl, "<script>")
	end := strings.LastIndex(tpl, "</script>")
	if start == -1 || end == -1 || end < start {
		return ""
	}
	return tpl[start+len("<script>") : end]
}
