import { QuartzComponent, QuartzComponentConstructor, QuartzComponentProps } from "./types"
import { classNames } from "../util/lang"
// @ts-ignore
import script from "./scripts/readingControls.inline"
import style from "./styles/readingControls.scss"

/** Body typefaces offered in the panel. Order groups serif, sans, then mono. */
export const TYPEFACES: { id: string; label: string; group: string }[] = [
  { id: "source-serif", label: "Source Serif", group: "Serif" },
  { id: "newsreader", label: "Newsreader", group: "Serif" },
  { id: "literata", label: "Literata", group: "Serif" },
  { id: "lora", label: "Lora", group: "Serif" },
  { id: "merriweather", label: "Merriweather", group: "Serif" },
  { id: "libre-baskerville", label: "Libre Baskerville", group: "Serif" },
  { id: "eb-garamond", label: "EB Garamond", group: "Serif" },
  { id: "crimson-pro", label: "Crimson Pro", group: "Serif" },
  { id: "spectral", label: "Spectral", group: "Serif" },
  { id: "system-sans", label: "System Sans", group: "Sans" },
  { id: "inter", label: "Inter", group: "Sans" },
  { id: "work-sans", label: "Work Sans", group: "Sans" },
  { id: "public-sans", label: "Public Sans", group: "Sans" },
  { id: "ibm-plex-sans", label: "IBM Plex Sans", group: "Sans" },
  { id: "source-sans", label: "Source Sans 3", group: "Sans" },
  { id: "nunito-sans", label: "Nunito Sans", group: "Sans" },
  { id: "manrope", label: "Manrope", group: "Sans" },
  { id: "intel-mono", label: "Intel One Mono", group: "Mono" },
  { id: "ibm-plex-mono", label: "IBM Plex Mono", group: "Mono" },
  { id: "jetbrains-mono", label: "JetBrains Mono", group: "Mono" },
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
                  <option value={f.id}>{f.label}</option>
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
`

ReadingControls.afterDOMLoaded = script
ReadingControls.css = style

export default (() => ReadingControls) satisfies QuartzComponentConstructor
