import { QuartzComponent, QuartzComponentConstructor, QuartzComponentProps } from "./types"
import { classNames } from "../util/lang"
// @ts-ignore
import script from "./scripts/readingControls.inline"
import style from "./styles/readingControls.scss"

const GF = "https://fonts.googleapis.com/css2?family="

/**
 * Body typefaces offered in the panel. Order groups serif, sans, then mono.
 *
 * `stack` is the CSS font stack and `href` the Google Fonts stylesheet to
 * fetch, where the family is not already on the page. Both ride on the
 * `<option>` as data attributes, which is where every consumer reads them
 * from — the panel script and the legacy fallback below alike. A second
 * hand-kept copy would drift the first time a face is added.
 *
 * Fonts load on demand: pulling twenty families up front would cost more than
 * the whole rest of the page, and a reader uses exactly one.
 */
export const TYPEFACES: {
  id: string
  label: string
  group: string
  stack: string
  href?: string
}[] = [
  { id: "source-serif", label: "Source Serif", group: "Serif", stack: '"Source Serif 4", Georgia, serif' },
  { id: "newsreader", label: "Newsreader", group: "Serif", stack: '"Newsreader", Georgia, serif' },
  {
    id: "literata",
    label: "Literata",
    group: "Serif",
    stack: '"Literata", Georgia, serif',
    href: GF + "Literata:ital,wght@0,400;0,600;1,400&display=swap",
  },
  {
    id: "lora",
    label: "Lora",
    group: "Serif",
    stack: '"Lora", Georgia, serif',
    href: GF + "Lora:ital,wght@0,400;0,600;1,400&display=swap",
  },
  {
    id: "merriweather",
    label: "Merriweather",
    group: "Serif",
    stack: '"Merriweather", Georgia, serif',
    href: GF + "Merriweather:ital,wght@0,400;0,700;1,400&display=swap",
  },
  {
    id: "libre-baskerville",
    label: "Libre Baskerville",
    group: "Serif",
    stack: '"Libre Baskerville", Georgia, serif',
    href: GF + "Libre+Baskerville:ital,wght@0,400;0,700;1,400&display=swap",
  },
  {
    id: "eb-garamond",
    label: "EB Garamond",
    group: "Serif",
    stack: '"EB Garamond", Georgia, serif',
    href: GF + "EB+Garamond:ital,wght@0,400;0,600;1,400&display=swap",
  },
  {
    id: "crimson-pro",
    label: "Crimson Pro",
    group: "Serif",
    stack: '"Crimson Pro", Georgia, serif',
    href: GF + "Crimson+Pro:ital,wght@0,400;0,600;1,400&display=swap",
  },
  {
    id: "spectral",
    label: "Spectral",
    group: "Serif",
    stack: '"Spectral", Georgia, serif',
    href: GF + "Spectral:ital,wght@0,400;0,600;1,400&display=swap",
  },
  { id: "system-sans", label: "System Sans", group: "Sans", stack: "var(--uiFont)" },
  {
    id: "inter",
    label: "Inter",
    group: "Sans",
    stack: '"Inter", system-ui, sans-serif',
    href: GF + "Inter:wght@400;600&display=swap",
  },
  {
    id: "work-sans",
    label: "Work Sans",
    group: "Sans",
    stack: '"Work Sans", system-ui, sans-serif',
    href: GF + "Work+Sans:ital,wght@0,400;0,600;1,400&display=swap",
  },
  {
    id: "public-sans",
    label: "Public Sans",
    group: "Sans",
    stack: '"Public Sans", system-ui, sans-serif',
    href: GF + "Public+Sans:ital,wght@0,400;0,600;1,400&display=swap",
  },
  {
    id: "ibm-plex-sans",
    label: "IBM Plex Sans",
    group: "Sans",
    stack: '"IBM Plex Sans", system-ui, sans-serif',
    href: GF + "IBM+Plex+Sans:ital,wght@0,400;0,600;1,400&display=swap",
  },
  {
    id: "source-sans",
    label: "Source Sans 3",
    group: "Sans",
    stack: '"Source Sans 3", system-ui, sans-serif',
    href: GF + "Source+Sans+3:ital,wght@0,400;0,600;1,400&display=swap",
  },
  {
    id: "nunito-sans",
    label: "Nunito Sans",
    group: "Sans",
    stack: '"Nunito Sans", system-ui, sans-serif',
    href: GF + "Nunito+Sans:ital,wght@0,400;0,600;1,400&display=swap",
  },
  {
    id: "manrope",
    label: "Manrope",
    group: "Sans",
    stack: '"Manrope", system-ui, sans-serif',
    href: GF + "Manrope:wght@400;600&display=swap",
  },
  { id: "intel-mono", label: "Intel One Mono", group: "Mono", stack: "var(--codeFont)" },
  {
    id: "ibm-plex-mono",
    label: "IBM Plex Mono",
    group: "Mono",
    stack: '"IBM Plex Mono", ui-monospace, monospace',
  },
  {
    id: "jetbrains-mono",
    label: "JetBrains Mono",
    group: "Mono",
    stack: '"JetBrains Mono", ui-monospace, monospace',
    href: GF + "JetBrains+Mono:ital,wght@0,400;0,600;1,400&display=swap",
  },
]

/**
 * Palettes. `mode` says which of Quartz's two modes each one is.
 *
 * This list is the only place a palette is declared. The pre-paint script
 * below serialises it, and the panel's `<option>`s carry the mode as a
 * `data-mode` attribute, which is where readingControls.inline.ts reads it
 * from — a second hand-kept copy would drift the first time one is added.
 */
export const PALETTES: { id: string; label: string; mode: "light" | "dark" }[] = [
  { id: "paper", label: "Paper", mode: "light" },
  { id: "sepia", label: "Sepia", mode: "light" },
  { id: "solarized", label: "Solarized", mode: "light" },
  { id: "slate", label: "Slate", mode: "light" },
  { id: "cream", label: "Cream", mode: "light" },
  { id: "linen", label: "Linen", mode: "light" },
  { id: "sage", label: "Sage", mode: "light" },
  { id: "dawn", label: "Dawn", mode: "light" },
  { id: "newsprint", label: "Newsprint", mode: "light" },
  { id: "ink", label: "Ink", mode: "dark" },
  { id: "gruvbox", label: "Gruvbox", mode: "dark" },
  { id: "nord", label: "Nord", mode: "dark" },
  { id: "dracula", label: "Dracula", mode: "dark" },
  { id: "solarized-dark", label: "Solarized Dark", mode: "dark" },
  { id: "one-dark", label: "One Dark", mode: "dark" },
  { id: "midnight", label: "Midnight", mode: "dark" },
  { id: "mocha", label: "Mocha", mode: "dark" },
]

const PALETTE_MODES = Object.fromEntries(PALETTES.map((p) => [p.id, p.mode]))

const GROUPS = ["Serif", "Sans", "Mono"]

/**
 * Reading affordances Quartz does not ship, plus the two it does (theme and
 * reader mode) gathered into one place.
 *
 * They live in a floating panel rather than the sidebar: stacked in a 320px
 * rail they wrap into a ragged column and compete with navigation, and these
 * are settings you reach for while reading, not while browsing.
 *
 * Every preference persists in localStorage and is applied before first paint,
 * so the page never flashes the wrong setting.
 */
const ReadingControls: QuartzComponent = ({ displayClass }: QuartzComponentProps) => {
  return (
    <>
      {/* Each rail gets its own tab, pinned to the rail's inner edge. */}
      <button
        type="button"
        class="rail-toggle rail-toggle--left"
        id="rail-left-toggle"
        aria-label="Toggle navigation sidebar"
        aria-expanded="true"
        title="Navigation"
      >
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" aria-hidden="true">
          <path d="M15 18l-6-6 6-6" stroke-linecap="round" stroke-linejoin="round" />
        </svg>
      </button>

      <button
        type="button"
        class="rail-toggle rail-toggle--right"
        id="rail-right-toggle"
        aria-label="Toggle contents sidebar"
        aria-expanded="true"
        title="Contents"
      >
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" aria-hidden="true">
          <path d="M9 18l6-6-6-6" stroke-linecap="round" stroke-linejoin="round" />
        </svg>
      </button>

      <div class={classNames(displayClass, "reading-dock")}>
      <div class="reading-panel" id="reading-panel" hidden>
        <div class="reading-row">
          <label class="reading-row-label" for="palette-select">
            Theme
          </label>
          <select class="reading-select" id="palette-select">
            {(["light", "dark"] as const).map((mode) => (
              <optgroup label={mode === "light" ? "Light" : "Dark"}>
                {PALETTES.filter((p) => p.mode === mode).map((p) => (
                  <option value={p.id} data-mode={p.mode}>
                    {p.label}
                  </option>
                ))}
              </optgroup>
            ))}
          </select>
        </div>

        <div class="reading-row">
          <label class="reading-row-label" for="typeface-select">
            Font
          </label>
          <select class="reading-select" id="typeface-select">
            {GROUPS.map((group) => (
              <optgroup label={group}>
                {TYPEFACES.filter((f) => f.group === group).map((f) => (
                  <option value={f.id} data-stack={f.stack} data-href={f.href}>
                    {f.label}
                  </option>
                ))}
              </optgroup>
            ))}
          </select>
        </div>

        <div class="reading-row">
          <span class="reading-row-label">Size</span>
          <div class="reading-segment">
            <button type="button" class="reading-key" id="type-smaller" aria-label="Smaller text">
              <span class="type-smaller">A</span>
            </button>
            <button type="button" class="reading-key" id="type-larger" aria-label="Larger text">
              <span class="type-larger">A</span>
            </button>
          </div>
        </div>

        <div class="reading-row">
          <span class="reading-row-label">Bold</span>
          <button
            type="button"
            class="reading-switch"
            id="bold-toggle"
            role="switch"
            aria-checked="false"
            aria-label="Bold text"
          >
            <span class="reading-switch-track" aria-hidden="true" />
          </button>
        </div>

        <div class="reading-row">
          <span class="reading-row-label">Margin</span>
          <div class="reading-segment">
            {/* Wider margin is a shorter line, so this button shrinks the measure.
                The widest stop is immersive reading: both rails hide and the
                page goes full screen. */}
            <button type="button" class="reading-key" id="margin-wider" aria-label="Wider margin">
              <span class="margin-glyph margin-glyph--wide" aria-hidden="true" />
            </button>
            <button type="button" class="reading-key" id="margin-narrower" aria-label="Narrower margin">
              <span class="margin-glyph margin-glyph--narrow" aria-hidden="true" />
            </button>
          </div>
        </div>

        <div class="reading-row">
          <span class="reading-row-label">Bionic</span>
          <button
            type="button"
            class="reading-switch"
            id="bionic-toggle"
            role="switch"
            aria-checked="true"
            aria-label="Bionic reading"
          >
            <span class="reading-switch-track" aria-hidden="true" />
          </button>
        </div>

        <div class="reading-row">
          <span class="reading-row-label">Focus</span>
          <button
            type="button"
            class="reading-switch"
            id="focus-toggle"
            role="switch"
            aria-checked="false"
            aria-label="Focus mode"
          >
            <span class="reading-switch-track" aria-hidden="true" />
          </button>
        </div>

        <div class="reading-row">
          <span class="reading-row-label">Full screen</span>
          <button
            type="button"
            class="reading-switch"
            id="fullscreen-toggle"
            role="switch"
            aria-checked="false"
            aria-label="Full screen"
          >
            <span class="reading-switch-track" aria-hidden="true" />
          </button>
        </div>
      </div>

      <button
        type="button"
        class="reading-fab"
        id="reading-fab"
        aria-label="Reading settings"
        aria-expanded="false"
        aria-controls="reading-panel"
        title="Reading settings"
      >
        <svg
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="1.6"
          stroke-linecap="round"
          stroke-linejoin="round"
          aria-hidden="true"
        >
          <circle cx="12" cy="12" r="3" />
          <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 1 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 1 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 1 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 1 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z" />
        </svg>
        </button>
      </div>
    </>
  )
}

ReadingControls.beforeDOMLoaded = `
  try {
    var root = document.documentElement
    if (localStorage.getItem("bionic") === "off") root.classList.add("bionic-off")
    var face = localStorage.getItem("typeface")
    if (face && face !== "source-serif") root.setAttribute("data-typeface", face)
    var modes = ${JSON.stringify(PALETTE_MODES)}
    var palette = localStorage.getItem("palette")
    // An unknown id (a palette that was renamed or dropped) matches no CSS at
    // all, so fall back rather than paint an unstyled page.
    if (!palette || !modes[palette]) palette = "paper"
    root.setAttribute("data-palette", palette)
    root.setAttribute("saved-theme", modes[palette])
    if (localStorage.getItem("bold") === "on") root.classList.add("bold-text")
    if (localStorage.getItem("focus") === "on") root.classList.add("focus-mode")
    if (localStorage.getItem("rail-left") === "closed") root.classList.add("rail-left-hidden")
    if (localStorage.getItem("rail-right") === "closed") root.classList.add("rail-right-hidden")
    var step = parseInt(localStorage.getItem("type-step") || "0", 10)
    if (step) root.style.setProperty("--type-step", String(step))
    var measure = parseInt(localStorage.getItem("measure-step") || "0", 10)
    if (measure) root.style.setProperty("--measure-step", String(measure))
    // -5 is MIN_MEASURE_STEP in readingControls.inline.ts: the widest margin,
    // which reads as immersive and hides both rails.
    if (measure <= -5) {
      root.classList.add("immersive", "rail-left-hidden", "rail-right-hidden")
    }
  } catch (e) {}

  /**
   * Legacy fallback: the same panel, wired without ES modules.
   *
   * Quartz serves its whole interactive bundle as a single
   * <script type="module">, which a browser without module support silently
   * never runs — no error, no console, just a page where every control is
   * inert. The Kindle's WebKit is exactly that browser, and a Kindle is the
   * device this reading panel was built for.
   *
   * This runs from the classic prescript instead, so it parses and executes
   * there. It stands down the moment the module bundle sets its flag, so a
   * modern browser never reaches any of it.
   *
   * Everything below is ES5 on purpose: no arrow functions, no let/const, no
   * template literals, no optional chaining, no for-of, no Element.closest,
   * and no second argument to classList.toggle — that engine has none of them,
   * and a single unsupported token would fail to parse the whole prescript,
   * taking the pre-paint above down with it.
   */
  ;(function () {
    var doc = document
    var root = doc.documentElement
    var MIN_TYPE = -2
    var MAX_TYPE = 4
    var MIN_MEASURE = -5
    var MAX_MEASURE = 5

    function id(name) {
      return doc.getElementById(name)
    }

    function read(key, fallback) {
      try {
        var v = localStorage.getItem(key)
        return v === null || v === undefined ? fallback : v
      } catch (e) {
        return fallback
      }
    }

    function write(key, value) {
      try {
        localStorage.setItem(key, value)
      } catch (e) {}
    }

    function hasClass(name) {
      return (" " + root.className + " ").indexOf(" " + name + " ") > -1
    }

    /* classList.toggle's force argument is unsupported here, so do it by hand. */
    function setClass(name, on) {
      if (on === hasClass(name)) return
      if (on) {
        root.className = root.className === "" ? name : root.className + " " + name
      } else {
        var out = (" " + root.className + " ").split(" " + name + " ").join(" ")
        root.className = out.replace(/^\\s+/, "").replace(/\\s+$/, "")
      }
    }

    function setChecked(name, on) {
      var el = id(name)
      if (el) el.setAttribute("aria-checked", on ? "true" : "false")
    }

    function flag(cls, on, name) {
      setClass(cls, on)
      setChecked(name, on)
    }

    function num(key) {
      var n = parseInt(read(key, "0"), 10)
      return isNaN(n) ? 0 : n
    }

    function clamp(v, lo, hi) {
      return v < lo ? lo : v > hi ? hi : v
    }

    /* ------------------------------------------------------------- panel */

    function panelOpen() {
      var p = id("reading-panel")
      return p !== null && !p.hasAttribute("hidden")
    }

    /* The [hidden] attribute alone is unreliable on this engine, so the
       inline display carries the state too. */
    function setPanel(open) {
      var p = id("reading-panel")
      if (p) {
        if (open) {
          p.removeAttribute("hidden")
          p.style.display = "block"
        } else {
          p.setAttribute("hidden", "hidden")
          p.style.display = "none"
        }
      }
      var fab = id("reading-fab")
      if (fab) fab.setAttribute("aria-expanded", open ? "true" : "false")
    }

    /* ------------------------------------------------------------- rails */

    function setRail(side, hidden) {
      setClass("rail-" + side + "-hidden", hidden)
      var el = id("rail-" + side + "-toggle")
      if (el) el.setAttribute("aria-expanded", hidden ? "false" : "true")
    }

    function syncImmersive(step) {
      var on = step <= MIN_MEASURE
      setClass("immersive", on)
      setRail("left", on || read("rail-left", "open") === "closed")
      setRail("right", on || read("rail-right", "open") === "closed")
    }

    /* ------------------------------------------------------------- steps */

    function applyStep(cssVar, value, lo, hi, decId, incId) {
      var v = clamp(value, lo, hi)
      if (root.style.setProperty) root.style.setProperty(cssVar, String(v))
      var dec = id(decId)
      var inc = id(incId)
      if (dec) dec.disabled = v <= lo
      if (inc) inc.disabled = v >= hi
      return v
    }

    function applyTypeStep(v) {
      return applyStep("--type-step", v, MIN_TYPE, MAX_TYPE, "type-smaller", "type-larger")
    }

    /* A wider margin is a shorter line, so "wider" walks the step down. */
    function applyMeasureStep(v) {
      return applyStep("--measure-step", v, MIN_MEASURE, MAX_MEASURE, "margin-wider", "margin-narrower")
    }

    /* --------------------------------------------------- selects by option */

    function option(selectId, value) {
      var sel = id(selectId)
      if (!sel) return null
      var opts = sel.getElementsByTagName("option")
      for (var i = 0; i < opts.length; i++) {
        if (opts[i].value === value) return opts[i]
      }
      return null
    }

    function applyPalette(value) {
      var opt = option("palette-select", value)
      var chosen = opt ? value : "paper"
      var mode = opt ? opt.getAttribute("data-mode") : null
      if (!mode) {
        var fb = option("palette-select", "paper")
        mode = fb ? fb.getAttribute("data-mode") : "light"
      }
      root.setAttribute("data-palette", chosen)
      if (mode) root.setAttribute("saved-theme", mode)
      var sel = id("palette-select")
      if (sel) sel.value = chosen
      return chosen
    }

    function applyTypeface(value) {
      var opt = option("typeface-select", value)
      var chosen = opt ? value : "source-serif"
      if (!opt) opt = option("typeface-select", chosen)
      if (opt) {
        var href = opt.getAttribute("data-href")
        if (href && !doc.querySelector('link[data-font="' + chosen + '"]')) {
          var link = doc.createElement("link")
          link.rel = "stylesheet"
          link.href = href
          link.setAttribute("data-font", chosen)
          doc.getElementsByTagName("head")[0].appendChild(link)
        }
        if (chosen === "source-serif") {
          root.removeAttribute("data-typeface")
          if (root.style.removeProperty) root.style.removeProperty("--bodyFont")
        } else {
          root.setAttribute("data-typeface", chosen)
          var stack = opt.getAttribute("data-stack")
          if (stack && root.style.setProperty) root.style.setProperty("--bodyFont", stack)
        }
      }
      var sel = id("typeface-select")
      if (sel) sel.value = chosen
      return chosen
    }

    /* -------------------------------------------------------------- wire */

    function onClick(name, handler) {
      var el = id(name)
      if (el) el.addEventListener("click", handler, false)
    }

    function onChange(name, handler) {
      var el = id(name)
      if (el) el.addEventListener("change", handler, false)
    }

    function start() {
      /* The module bundle got here first; it owns the panel. */
      if (window.__readingControlsModern) return
      if (!id("reading-fab")) return

      var bionicOff = read("bionic", "on") === "off"
      setClass("bionic-off", bionicOff)
      setChecked("bionic-toggle", !bionicOff)
      applyPalette(read("palette", "paper"))
      applyTypeface(read("typeface", "source-serif"))
      applyTypeStep(num("type-step"))
      applyMeasureStep(num("measure-step"))
      flag("bold-text", read("bold", "off") === "on", "bold-toggle")
      flag("focus-mode", read("focus", "off") === "on", "focus-toggle")
      syncImmersive(num("measure-step"))
      setPanel(false)

      onClick("reading-fab", function () {
        setPanel(!panelOpen())
      })

      onClick("rail-left-toggle", function () {
        var hidden = !hasClass("rail-left-hidden")
        setRail("left", hidden)
        write("rail-left", hidden ? "closed" : "open")
      })

      onClick("rail-right-toggle", function () {
        var hidden = !hasClass("rail-right-hidden")
        setRail("right", hidden)
        write("rail-right", hidden ? "closed" : "open")
      })

      onClick("bionic-toggle", function () {
        var off = !hasClass("bionic-off")
        setClass("bionic-off", off)
        setChecked("bionic-toggle", !off)
        write("bionic", off ? "off" : "on")
      })

      onClick("bold-toggle", function () {
        var on = !hasClass("bold-text")
        flag("bold-text", on, "bold-toggle")
        write("bold", on ? "on" : "off")
      })

      onClick("focus-toggle", function () {
        var on = !hasClass("focus-mode")
        flag("focus-mode", on, "focus-toggle")
        write("focus", on ? "on" : "off")
      })

      onChange("palette-select", function () {
        write("palette", applyPalette(id("palette-select").value))
      })

      onChange("typeface-select", function () {
        write("typeface", applyTypeface(id("typeface-select").value))
      })

      function stepType(delta) {
        write("type-step", String(applyTypeStep(num("type-step") + delta)))
      }
      onClick("type-smaller", function () {
        stepType(-1)
      })
      onClick("type-larger", function () {
        stepType(1)
      })

      function stepMargin(delta) {
        var step = applyMeasureStep(num("measure-step") + delta)
        write("measure-step", String(step))
        syncImmersive(step)
      }
      onClick("margin-wider", function () {
        stepMargin(-1)
      })
      onClick("margin-narrower", function () {
        stepMargin(1)
      })

      /* No Fullscreen API on this engine; hide the row rather than offer a
         switch that cannot do anything. */
      var fs = id("fullscreen-toggle")
      if (fs && !root.requestFullscreen && !root.webkitRequestFullscreen) {
        var row = fs.parentNode
        if (row && row.style) row.style.display = "none"
      } else if (fs) {
        fs.addEventListener(
          "click",
          function () {
            try {
              if (doc.fullscreenElement || doc.webkitFullscreenElement) {
                if (doc.exitFullscreen) doc.exitFullscreen()
                else if (doc.webkitExitFullscreen) doc.webkitExitFullscreen()
              } else if (root.requestFullscreen) {
                root.requestFullscreen()
              } else if (root.webkitRequestFullscreen) {
                root.webkitRequestFullscreen()
              }
            } catch (e) {}
          },
          false,
        )
      }

      /* Dismiss on an outside tap. Element.closest does not exist here, so
         walk the parents by hand. */
      doc.addEventListener(
        "click",
        function (ev) {
          if (!panelOpen()) return
          var node = ev.target
          while (node && node !== doc) {
            if (
              node.className &&
              typeof node.className === "string" &&
              (" " + node.className + " ").indexOf(" reading-dock ") > -1
            ) {
              return
            }
            node = node.parentNode
          }
          setPanel(false)
        },
        false,
      )
    }

    /* Module scripts are deferred, so the flag is already set by the time
       DOMContentLoaded fires — but give a late bundle one more beat before
       taking over, and never bind twice. */
    function boot() {
      if (window.__readingControlsFallback) return
      window.__readingControlsFallback = true
      start()
    }

    if (doc.readyState === "complete" || doc.readyState === "interactive") {
      setTimeout(boot, 0)
    } else {
      doc.addEventListener("DOMContentLoaded", function () {
        setTimeout(boot, 0)
      }, false)
    }
  })()
`

ReadingControls.afterDOMLoaded = script
ReadingControls.css = style

export default (() => ReadingControls) satisfies QuartzComponentConstructor
