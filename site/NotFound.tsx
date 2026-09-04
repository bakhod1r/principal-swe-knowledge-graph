import { QuartzComponent, QuartzComponentConstructor, QuartzComponentProps } from "../types"
import PagefindSearchConstructor from "../PagefindSearch"

/**
 * 404 page with search on it.
 *
 * Quartz's stock page is a dead end: a heading and a link home. Most 404s here
 * are a renamed or mistyped note, so the page that has to be useful is the one
 * that lets you look the note up. The search component's script and styles are
 * already in the global bundle (it lives in the sidebar of every real page), so
 * rendering it here costs nothing extra.
 */
const PagefindSearch = PagefindSearchConstructor()

const NotFound: QuartzComponent = (props: QuartzComponentProps) => {
  const { cfg } = props
  const url = new URL(`https://${cfg.baseUrl ?? "example.com"}`)
  const baseDir = url.pathname

  return (
    <article class="popover-hint not-found">
      <h1>404</h1>
      <p>That page does not exist — it may have been renamed or moved.</p>
      <PagefindSearch {...props} />
      <p>
        <a href={baseDir}>Back to the index</a>
      </p>
    </article>
  )
}

export default (() => NotFound) satisfies QuartzComponentConstructor
