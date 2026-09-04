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

/** Palettes. `light` and `dark` say which of Quartz's two modes each one is. */
export const PALETTES: { id: string; label: string }[] = [
  { id: "paper", label: "Paper" },
  { id: "sepia", label: "Sepia" },
  { id: "solarized", label: "Solarized" },
  { id: "slate", label: "Slate" },
  { id: "ink", label: "Ink" },
  { id: "gruvbox", label: "Gruvbox" },
  { id: "nord", label: "Nord" },
]

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
    <div class={classNames(displayClass, "reading-dock")}>
      <div class="reading-panel" id="reading-panel" hidden>
        <div class="reading-row">
          <label class="reading-row-label" for="palette-select">
            Theme
          </label>
          <select class="reading-select" id="palette-select">
            {PALETTES.map((p) => (
              <option value={p.id}>{p.label}</option>
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
          <span class="reading-row-label">Reader</span>
          {/* `readermode` is the hook Quartz's own script binds to. */}
          <button
            type="button"
            class="reading-switch readermode"
            id="reader-toggle"
            role="switch"
            aria-checked="false"
            aria-label="Reader mode"
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
        id="sidebar-toggle"
        aria-label="Toggle sidebar"
        aria-expanded="true"
        title="Toggle sidebar"
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
          <rect x="3" y="4" width="18" height="16" rx="1" />
          <path d="M9 4v16" />
        </svg>
      </button>

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
  )
}

ReadingControls.beforeDOMLoaded = `
  try {
    var root = document.documentElement
    if (localStorage.getItem("bionic") === "off") root.classList.add("bionic-off")
    var face = localStorage.getItem("typeface")
    if (face && face !== "source-serif") root.setAttribute("data-typeface", face)
    var palette = localStorage.getItem("palette")
    if (palette) root.setAttribute("data-palette", palette)
    if (localStorage.getItem("focus") === "on") root.classList.add("focus-mode")
    var step = parseInt(localStorage.getItem("type-step") || "0", 10)
    if (step) root.style.setProperty("--type-step", String(step))
  } catch (e) {}
`

ReadingControls.afterDOMLoaded = script
ReadingControls.css = style

export default (() => ReadingControls) satisfies QuartzComponentConstructor
