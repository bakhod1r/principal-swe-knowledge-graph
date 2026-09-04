import { QuartzComponent, QuartzComponentConstructor, QuartzComponentProps } from "./types"
import { classNames } from "../util/lang"
// @ts-ignore
import script from "./scripts/pagefindSearch.inline"
import style from "./styles/pagefindSearch.scss"

/**
 * Search backed by Pagefind instead of Quartz's built-in index.
 *
 * Quartz ships the full-text index of every note to the browser up front. At
 * 9411 notes that is by far the heaviest thing on the page. Pagefind builds a
 * sharded index at deploy time (see the workflow's "Build search index" step)
 * and fetches only the shards a query touches, so first paint no longer pays
 * for the whole corpus.
 *
 * The base path is passed through as a data attribute: the site is served from
 * a project subpath, and Pagefind's own bundle and result URLs are root-relative.
 */
const PagefindSearch: QuartzComponent = ({ displayClass, cfg }: QuartzComponentProps) => {
  const basePath = "/" + (cfg.baseUrl?.split("/").slice(1).join("/") ?? "")
  const base = basePath === "/" ? "" : basePath.replace(/\/$/, "")

  return (
    <div class={classNames(displayClass, "pagefind-search")} data-base={base}>
      <button type="button" class="search-button" id="pagefind-open" aria-label="Search">
        <svg
          role="img"
          aria-hidden="true"
          viewBox="0 0 24 24"
          width="18"
          height="18"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <circle cx="11" cy="11" r="7" />
          <line x1="16.5" y1="16.5" x2="21" y2="21" />
        </svg>
        <span class="search-button-label">Search</span>
      </button>

      <div class="search-container" id="pagefind-container" hidden>
        <div class="search-space">
          <input
            type="text"
            class="search-bar"
            id="pagefind-input"
            autocomplete="off"
            autocorrect="off"
            spellcheck={false}
            placeholder="Search notes…"
            aria-label="Search notes"
          />
          <output class="search-layout" id="pagefind-results" />
        </div>
      </div>
    </div>
  )
}

PagefindSearch.afterDOMLoaded = script
PagefindSearch.css = style

export default (() => PagefindSearch) satisfies QuartzComponentConstructor
