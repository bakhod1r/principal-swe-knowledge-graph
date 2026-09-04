const MIN_STEP = -2
const MAX_STEP = 4

// 68ch ± 4ch per step: 48ch is still readable, 88ch is as wide as the layout
// gives before the right rail starts fighting for the space.
const MIN_MEASURE_STEP = -5
const MAX_MEASURE_STEP = 5

/**
 * Body typefaces. Each entry carries the CSS stack and, where the family is
 * not already on the page, the Google Fonts stylesheet to fetch for it.
 *
 * Fonts load on demand: pulling twenty families up front would cost more than
 * the whole rest of the page, and a reader uses exactly one.
 */
const FONTS: Record<string, { stack: string; href?: string }> = {
  "source-serif": { stack: '"Source Serif 4", Georgia, serif' },
  newsreader: { stack: '"Newsreader", Georgia, serif' },
  literata: {
    stack: '"Literata", Georgia, serif',
    href: "https://fonts.googleapis.com/css2?family=Literata:ital,wght@0,400;0,600;1,400&display=swap",
  },
  lora: {
    stack: '"Lora", Georgia, serif',
    href: "https://fonts.googleapis.com/css2?family=Lora:ital,wght@0,400;0,600;1,400&display=swap",
  },
  merriweather: {
    stack: '"Merriweather", Georgia, serif',
    href: "https://fonts.googleapis.com/css2?family=Merriweather:ital,wght@0,400;0,700;1,400&display=swap",
  },
  "libre-baskerville": {
    stack: '"Libre Baskerville", Georgia, serif',
    href: "https://fonts.googleapis.com/css2?family=Libre+Baskerville:ital,wght@0,400;0,700;1,400&display=swap",
  },
  "eb-garamond": {
    stack: '"EB Garamond", Georgia, serif',
    href: "https://fonts.googleapis.com/css2?family=EB+Garamond:ital,wght@0,400;0,600;1,400&display=swap",
  },
  "crimson-pro": {
    stack: '"Crimson Pro", Georgia, serif',
    href: "https://fonts.googleapis.com/css2?family=Crimson+Pro:ital,wght@0,400;0,600;1,400&display=swap",
  },
  spectral: {
    stack: '"Spectral", Georgia, serif',
    href: "https://fonts.googleapis.com/css2?family=Spectral:ital,wght@0,400;0,600;1,400&display=swap",
  },
  "system-sans": { stack: "var(--uiFont)" },
  inter: {
    stack: '"Inter", system-ui, sans-serif',
    href: "https://fonts.googleapis.com/css2?family=Inter:wght@400;600&display=swap",
  },
  "work-sans": {
    stack: '"Work Sans", system-ui, sans-serif',
    href: "https://fonts.googleapis.com/css2?family=Work+Sans:ital,wght@0,400;0,600;1,400&display=swap",
  },
  "public-sans": {
    stack: '"Public Sans", system-ui, sans-serif',
    href: "https://fonts.googleapis.com/css2?family=Public+Sans:ital,wght@0,400;0,600;1,400&display=swap",
  },
  "ibm-plex-sans": {
    stack: '"IBM Plex Sans", system-ui, sans-serif',
    href: "https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:ital,wght@0,400;0,600;1,400&display=swap",
  },
  "source-sans": {
    stack: '"Source Sans 3", system-ui, sans-serif',
    href: "https://fonts.googleapis.com/css2?family=Source+Sans+3:ital,wght@0,400;0,600;1,400&display=swap",
  },
  "nunito-sans": {
    stack: '"Nunito Sans", system-ui, sans-serif',
    href: "https://fonts.googleapis.com/css2?family=Nunito+Sans:ital,wght@0,400;0,600;1,400&display=swap",
  },
  manrope: {
    stack: '"Manrope", system-ui, sans-serif',
    href: "https://fonts.googleapis.com/css2?family=Manrope:wght@400;600&display=swap",
  },
  "intel-mono": { stack: 'var(--codeFont)' },
  "ibm-plex-mono": { stack: '"IBM Plex Mono", ui-monospace, monospace' },
  "jetbrains-mono": {
    stack: '"JetBrains Mono", ui-monospace, monospace',
    href: "https://fonts.googleapis.com/css2?family=JetBrains+Mono:ital,wght@0,400;0,600;1,400&display=swap",
  },
}

/**
 * Which of Quartz's two modes each palette is, read off the `data-mode` the
 * panel's `<option>`s carry. ReadingControls.tsx declares the palettes; a
 * copy of the list here would drift the first time one is added.
 */
function paletteModes(): Record<string, string> {
  const modes: Record<string, string> = {}
  for (const el of document.querySelectorAll<HTMLOptionElement>("#palette-select option")) {
    modes[el.value] = el.dataset.mode ?? "light"
  }
  return modes
}

function read(key: string, fallback: string): string {
  try {
    return localStorage.getItem(key) ?? fallback
  } catch (e) {
    return fallback
  }
}

function write(key: string, value: string) {
  try {
    localStorage.setItem(key, value)
  } catch (e) {
    /* private mode: the preference simply does not persist */
  }
}

function setChecked(selector: string, on: boolean) {
  for (const el of document.querySelectorAll(selector)) {
    el.setAttribute("aria-checked", String(on))
  }
}

/** A class on <html> plus a switch that reflects it — the shape of every toggle here. */
function applyFlag(cls: string, on: boolean, selector: string) {
  document.documentElement.classList.toggle(cls, on)
  setChecked(selector, on)
}

/* ----------------------------------------------------------- typeface */

function ensureFontLoaded(id: string) {
  const href = FONTS[id]?.href
  if (href === undefined) return
  if (document.querySelector(`link[data-font="${id}"]`) !== null) return

  const link = document.createElement("link")
  link.rel = "stylesheet"
  link.href = href
  link.dataset.font = id
  document.head.appendChild(link)
}

function applyTypeface(id: string) {
  const chosen = id in FONTS ? id : "source-serif"
  ensureFontLoaded(chosen)
  const root = document.documentElement
  if (chosen === "source-serif") {
    root.removeAttribute("data-typeface")
    root.style.removeProperty("--bodyFont")
  } else {
    root.setAttribute("data-typeface", chosen)
    root.style.setProperty("--bodyFont", FONTS[chosen].stack)
  }
  for (const el of document.querySelectorAll<HTMLSelectElement>("#typeface-select")) {
    el.value = chosen
  }
  return chosen
}

/* ------------------------------------------------------------ palette */

function applyPalette(id: string) {
  const modes = paletteModes()
  const chosen = modes[id] ? id : "paper"
  document.documentElement.setAttribute("data-palette", chosen)
  // Quartz keys syntax highlighting and its icons off `saved-theme`. On a page
  // without the panel there is nothing to read the mode from, so leave the
  // attribute the pre-paint script already set.
  const mode = modes[chosen]
  if (mode) document.documentElement.setAttribute("saved-theme", mode)
  for (const el of document.querySelectorAll<HTMLSelectElement>("#palette-select")) {
    el.value = chosen
  }
  return chosen
}

/* ---------------------------------------------------------- type size */

function applyTypeStep(step: number) {
  const clamped = Math.max(MIN_STEP, Math.min(MAX_STEP, step))
  document.documentElement.style.setProperty("--type-step", String(clamped))
  for (const el of document.querySelectorAll("#type-smaller")) {
    ;(el as HTMLButtonElement).disabled = clamped <= MIN_STEP
  }
  for (const el of document.querySelectorAll("#type-larger")) {
    ;(el as HTMLButtonElement).disabled = clamped >= MAX_STEP
  }
  return clamped
}

function currentStep(): number {
  const raw = document.documentElement.style.getPropertyValue("--type-step").trim()
  const parsed = parseInt(raw === "" ? read("type-step", "0") : raw, 10)
  return Number.isNaN(parsed) ? 0 : parsed
}

/* ------------------------------------------------------------- margin */

/**
 * Line length, in 4ch steps around the 68ch default. A shorter line is a wider
 * margin, which is the way a reader thinks about it, so the "wider margin"
 * button walks the step down.
 */
function applyMeasureStep(step: number) {
  const clamped = Math.max(MIN_MEASURE_STEP, Math.min(MAX_MEASURE_STEP, step))
  document.documentElement.style.setProperty("--measure-step", String(clamped))
  for (const el of document.querySelectorAll("#margin-wider")) {
    ;(el as HTMLButtonElement).disabled = clamped <= MIN_MEASURE_STEP
  }
  for (const el of document.querySelectorAll("#margin-narrower")) {
    ;(el as HTMLButtonElement).disabled = clamped >= MAX_MEASURE_STEP
  }
  return clamped
}

function currentMeasureStep(): number {
  const raw = document.documentElement.style.getPropertyValue("--measure-step").trim()
  const parsed = parseInt(raw === "" ? read("measure-step", "0") : raw, 10)
  return Number.isNaN(parsed) ? 0 : parsed
}

/**
 * The widest margin is immersive reading: at the last stop both rails hide and
 * the page goes full screen, and stepping back off it restores whatever rail
 * state the reader had chosen. Immersive never writes the rail preferences, so
 * that state survives the round trip.
 */
function immersiveAt(step: number): boolean {
  return step <= MIN_MEASURE_STEP
}

function syncImmersive(step: number) {
  const on = immersiveAt(step)
  document.documentElement.classList.toggle("immersive", on)
  applyRail("left", on || read("rail-left", "open") === "closed")
  applyRail("right", on || read("rail-right", "open") === "closed")
}

/* --------------------------------------------------------- fullscreen */

function syncFullscreen() {
  applyFlag("is-fullscreen", document.fullscreenElement !== null, "#fullscreen-toggle")
}

async function toggleFullscreen() {
  try {
    if (document.fullscreenElement) {
      await document.exitFullscreen()
    } else {
      await document.documentElement.requestFullscreen()
    }
  } catch (e) {
    /* the browser refused (iOS Safari, or a permissions policy); ignore */
  }
}

/* --------------------------------------------------------------- dock */

function setPanel(open: boolean) {
  for (const panel of document.querySelectorAll<HTMLElement>("#reading-panel")) {
    panel.hidden = !open
  }
  for (const fab of document.querySelectorAll("#reading-fab")) {
    fab.setAttribute("aria-expanded", String(open))
  }
}

function panelOpen(): boolean {
  const panel = document.querySelector<HTMLElement>("#reading-panel")
  return panel !== null && !panel.hidden
}

function applyRail(side: "left" | "right", hidden: boolean) {
  document.documentElement.classList.toggle(`rail-${side}-hidden`, hidden)
  for (const el of document.querySelectorAll(`#rail-${side}-toggle`)) {
    el.setAttribute("aria-expanded", String(!hidden))
  }
}

/* --------------------------------------------------------------- wire */

function bind(selector: string, handler: (ev: Event) => void) {
  for (const el of document.querySelectorAll(selector)) {
    el.addEventListener("click", handler)
    window.addCleanup(() => el.removeEventListener("click", handler))
  }
}

document.addEventListener("nav", () => {
  const bionicOff = read("bionic", "on") === "off"
  document.documentElement.classList.toggle("bionic-off", bionicOff)
  setChecked("#bionic-toggle", !bionicOff)

  applyPalette(read("palette", "paper"))
  applyTypeface(read("typeface", "source-serif"))
  applyTypeStep(currentStep())
  applyMeasureStep(currentMeasureStep())
  applyFlag("bold-text", read("bold", "off") === "on", "#bold-toggle")
  applyFlag("focus-mode", read("focus", "off") === "on", "#focus-toggle")
  syncImmersive(currentMeasureStep())
  syncFullscreen()
  setPanel(false)

  bind("#reading-fab", () => setPanel(!panelOpen()))

  for (const side of ["left", "right"] as const) {
    bind(`#rail-${side}-toggle`, () => {
      const hidden = !document.documentElement.classList.contains(`rail-${side}-hidden`)
      applyRail(side, hidden)
      write(`rail-${side}`, hidden ? "closed" : "open")
    })
  }

  bind("#bionic-toggle", () => {
    const off = !document.documentElement.classList.contains("bionic-off")
    document.documentElement.classList.toggle("bionic-off", off)
    setChecked("#bionic-toggle", !off)
    write("bionic", off ? "off" : "on")
  })

  bind("#bold-toggle", () => {
    const on = !document.documentElement.classList.contains("bold-text")
    applyFlag("bold-text", on, "#bold-toggle")
    write("bold", on ? "on" : "off")
  })

  bind("#focus-toggle", () => {
    const on = !document.documentElement.classList.contains("focus-mode")
    applyFlag("focus-mode", on, "#focus-toggle")
    write("focus", on ? "on" : "off")
  })

  for (const el of document.querySelectorAll<HTMLSelectElement>("#typeface-select")) {
    const onChange = () => write("typeface", applyTypeface(el.value))
    el.addEventListener("change", onChange)
    window.addCleanup(() => el.removeEventListener("change", onChange))
  }

  for (const el of document.querySelectorAll<HTMLSelectElement>("#palette-select")) {
    const onChange = () => write("palette", applyPalette(el.value))
    el.addEventListener("change", onChange)
    window.addCleanup(() => el.removeEventListener("change", onChange))
  }

  bind("#type-smaller", () => write("type-step", String(applyTypeStep(currentStep() - 1))))
  bind("#type-larger", () => write("type-step", String(applyTypeStep(currentStep() + 1))))
  // The click is the user gesture full screen needs, so the request belongs
  // here rather than in the restore path that runs on every navigation.
  const stepMargin = (delta: number) => {
    const was = immersiveAt(currentMeasureStep())
    const step = applyMeasureStep(currentMeasureStep() + delta)
    write("measure-step", String(step))
    syncImmersive(step)
    const now = immersiveAt(step)
    if (now !== was && now !== (document.fullscreenElement !== null)) toggleFullscreen()
  }
  bind("#margin-wider", () => stepMargin(-1))
  bind("#margin-narrower", () => stepMargin(1))
  bind("#fullscreen-toggle", toggleFullscreen)

  const onFullscreenChange = () => syncFullscreen()
  document.addEventListener("fullscreenchange", onFullscreenChange)
  window.addCleanup(() => document.removeEventListener("fullscreenchange", onFullscreenChange))

  // Dismiss the panel the way every popover should: outside click, or Escape.
  const onPointerDown = (ev: Event) => {
    const target = ev.target as HTMLElement | null
    if (panelOpen() && target !== null && target.closest(".reading-dock") === null) {
      setPanel(false)
    }
  }
  const onKeyDown = (ev: KeyboardEvent) => {
    if (ev.key === "Escape" && panelOpen()) {
      setPanel(false)
      document.querySelector<HTMLElement>("#reading-fab")?.focus()
    }
  }
  document.addEventListener("pointerdown", onPointerDown)
  document.addEventListener("keydown", onKeyDown)
  window.addCleanup(() => {
    document.removeEventListener("pointerdown", onPointerDown)
    document.removeEventListener("keydown", onKeyDown)
  })
})
