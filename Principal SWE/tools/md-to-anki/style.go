package main

// Card styling and templates for the note type this tool creates.
//
// Anki renders each card in a WebView, so the answer side ships a small
// regex highlighter that colours the <pre><code class="language-X"> blocks
// produced by renderHTML. No network, no external library.

// cardTemplateName marks the note types this tool owns — restyleModel only
// touches a model whose template carries this name.
const cardTemplateName = "Principal SWE Card"

const cardCSS = `
.card {
  font-family: -apple-system, "SF Pro Text", Helvetica, Arial, sans-serif;
  font-size: 17px;
  line-height: 1.6;
  text-align: left;
  color: #d8dee9;
  background: #14171c;
  padding: 22px 26px;
}

/* ── question ── */
.q {
  font-size: 21px;
  font-weight: 600;
  color: #eceff4;
  letter-spacing: -0.01em;
}
.q .num {
  display: inline-block;
  min-width: 2.1em;
  margin-right: 6px;
  padding: 1px 8px;
  border-radius: 5px;
  background: #2b3a52;
  color: #8fb8ff;
  font-size: 15px;
  font-weight: 700;
  text-align: center;
}
.q .note-title {
  display: block;
  margin-bottom: 6px;
  font-size: 13px;
  font-weight: 500;
  letter-spacing: 0.04em;
  text-transform: uppercase;
  color: #6b7688;
}

hr#answer {
  height: 1px;
  margin: 20px 0;
  border: 0;
  background: linear-gradient(90deg, #3b4252, transparent);
}

/* ── answer body ── */
.a p { margin: 0 0 12px; }
.a b { color: #ffffff; font-weight: 600; }
.a i { color: #cbd3e1; }

.a ul {
  margin: 0 0 12px;
  padding-left: 22px;
}
.a li { margin: 4px 0; }
.a li::marker { color: #7f92b0; }

.a blockquote {
  margin: 12px 0;
  padding: 8px 14px;
  border-left: 3px solid #5e81ac;
  border-radius: 0 6px 6px 0;
  background: #1b202a;
  color: #b8c2d4;
}

.a a { color: #88c0d0; text-decoration: none; }
.a a:hover { text-decoration: underline; }

/* ── code ── */
.a :not(pre) > code {
  padding: 1px 6px;
  border-radius: 4px;
  background: #232936;
  color: #a3d0ff;
  font-family: ui-monospace, "SF Mono", Menlo, monospace;
  font-size: 0.88em;
}

.a pre {
  position: relative;
  margin: 14px 0;
  padding: 14px 16px;
  border: 1px solid #262d3a;
  border-radius: 8px;
  background: #0f1319;
  overflow-x: auto;
}
.a pre code {
  font-family: ui-monospace, "SF Mono", Menlo, monospace;
  font-size: 14px;
  line-height: 1.55;
  color: #d8dee9;
  white-space: pre;
}
.a pre[data-lang]::before {
  content: attr(data-lang);
  position: absolute;
  top: 0;
  right: 0;
  padding: 2px 9px;
  border-radius: 0 7px 0 7px;
  background: #202634;
  color: #7f8ea6;
  font-family: ui-monospace, Menlo, monospace;
  font-size: 11px;
  letter-spacing: 0.05em;
}

/* plain-text diagram blocks stay unhighlighted, just monospace */
.a pre[data-lang="text"] code { color: #9fb0c8; }

/* ── highlight tokens ── */
.tok-kw   { color: #c792ea; }
.tok-str  { color: #c3e88d; }
.tok-com  { color: #5c6c7d; font-style: italic; }
.tok-num  { color: #f78c6c; }
.tok-fn   { color: #82aaff; }
.tok-type { color: #ffcb6b; }

/* ── night mode / light theme fallback ── */
.card.night_mode, .nightMode .card { background: #14171c; }
`

// backTemplate renders the answer side and runs the highlighter.
const backTemplate = `{{FrontSide}}

<hr id=answer>

<div class="a" id="answer-body">{{Back}}</div>

<script>
(function () {
  var KEYWORDS = {
    go: "break case chan const continue default defer else fallthrough for func go goto if import interface map package range return select struct switch type var nil true false iota make new len cap append copy delete panic recover string int int8 int16 int32 int64 uint uintptr byte rune float32 float64 bool error any",
    python: "def class return if elif else for while import from as pass break continue with try except finally raise lambda yield None True False and or not in is global nonlocal async await self",
    javascript: "const let var function return if else for while class new this import export from async await try catch finally throw typeof instanceof null undefined true false switch case break continue default extends super yield",
    typescript: "const let var function return if else for while class new this import export from async await try catch finally throw typeof instanceof null undefined true false switch case break continue default extends super interface type enum implements public private protected readonly",
    rust: "fn let mut const struct enum impl trait pub use mod match if else for while loop return self Self where async await move ref dyn Box Option Result Some None Ok Err true false",
    sql: "select from where group by order having join left right inner outer on as insert into values update set delete create table drop alter index limit offset union distinct and or not null"
  };
  var ALIAS = { js: "javascript", ts: "typescript", py: "python", golang: "go", sh: "bash", shell: "bash", bash: "bash", zsh: "bash" };
  var LINE_COMMENT = { go: "//", javascript: "//", typescript: "//", rust: "//", python: "#", bash: "#", sql: "--" };

  function esc(s) {
    return s.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;");
  }

  function highlight(src, lang) {
    var kw = (KEYWORDS[lang] || "").split(" ").filter(Boolean);
    var comment = LINE_COMMENT[lang];
    var out = "";
    var i = 0;

    while (i < src.length) {
      var ch = src[i];

      // line comment
      if (comment && src.substr(i, comment.length) === comment) {
        var nl = src.indexOf("\n", i);
        if (nl === -1) nl = src.length;
        out += '<span class="tok-com">' + esc(src.slice(i, nl)) + "</span>";
        i = nl;
        continue;
      }

      // block comment
      if (src.substr(i, 2) === "/*") {
        var end = src.indexOf("*/", i + 2);
        end = end === -1 ? src.length : end + 2;
        out += '<span class="tok-com">' + esc(src.slice(i, end)) + "</span>";
        i = end;
        continue;
      }

      // string / char / backtick literal
      if (ch === '"' || ch === "'" || ch === "` + "`" + `") {
        var j = i + 1;
        while (j < src.length && src[j] !== ch) {
          if (src[j] === "\\") j++;
          j++;
        }
        out += '<span class="tok-str">' + esc(src.slice(i, Math.min(j + 1, src.length))) + "</span>";
        i = j + 1;
        continue;
      }

      // number
      if (/[0-9]/.test(ch) && !/[A-Za-z_]/.test(src[i - 1] || "")) {
        var n = i;
        while (n < src.length && /[0-9a-fA-FxX._]/.test(src[n])) n++;
        out += '<span class="tok-num">' + esc(src.slice(i, n)) + "</span>";
        i = n;
        continue;
      }

      // identifier: keyword, call, or plain
      if (/[A-Za-z_]/.test(ch)) {
        var w = i;
        while (w < src.length && /[A-Za-z0-9_]/.test(src[w])) w++;
        var word = src.slice(i, w);
        var next = src[w];
        if (kw.indexOf(word) !== -1) {
          out += '<span class="tok-kw">' + esc(word) + "</span>";
        } else if (next === "(") {
          out += '<span class="tok-fn">' + esc(word) + "</span>";
        } else if (/^[A-Z]/.test(word)) {
          out += '<span class="tok-type">' + esc(word) + "</span>";
        } else {
          out += esc(word);
        }
        i = w;
        continue;
      }

      out += esc(ch);
      i++;
    }
    return out;
  }

  var blocks = document.querySelectorAll("#answer-body pre code");
  for (var b = 0; b < blocks.length; b++) {
    var code = blocks[b];
    var cls = code.className || "";
    var m = cls.match(/language-([A-Za-z0-9_+-]+)/);
    var lang = m ? m[1].toLowerCase() : "";
    lang = ALIAS[lang] || lang;
    if (lang) code.parentNode.setAttribute("data-lang", lang);
    // Plain text blocks are usually ASCII diagrams — leave them alone.
    if (!lang || lang === "text" || lang === "plain") continue;
    code.innerHTML = highlight(code.textContent, lang);
  }
})();
</script>
`

// frontTemplate splits the "Note title — 3. Section" front into two lines.
const frontTemplate = `<div class="q" id="question">{{Front}}</div>

<script>
(function () {
  var el = document.getElementById("question");
  if (!el) return;
  var parts = el.textContent.split(" — ");
  if (parts.length < 2) return;
  var title = parts.shift();
  var rest = parts.join(" — ");
  var m = rest.match(/^([0-9.]+)\.\s*(.*)$/);
  var esc = function (s) {
    return s.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;");
  };
  var head = '<span class="note-title">' + esc(title) + "</span>";
  el.innerHTML = m
    ? head + '<span class="num">' + esc(m[1]) + "</span>" + esc(m[2])
    : head + esc(rest);
})();
</script>
`
