import { PageLayout, SharedLayout } from "./quartz/cfg"
import * as Component from "./quartz/components"
import ReadingControls from "./quartz/components/ReadingControls"
import PagefindSearch from "./quartz/components/PagefindSearch"
import SiblingNav from "./quartz/components/SiblingNav"

// components shared across all pages
export const sharedPageComponents: SharedLayout = {
  head: Component.Head(),
  header: [],
  // SiblingNav renders nothing on index and folder pages, so it is safe here;
  // afterBody is shared across page types in Quartz. Reading dock is floating
  // and shows everywhere.
  afterBody: [SiblingNav(), ReadingControls()],
  footer: Component.Footer({
    links: {
      GitHub: "https://github.com/bakhod1r/principal-swe-knowledge-graph",
    },
  }),
}

// components for pages that display a single page (e.g. a single note)
export const defaultContentPageLayout: PageLayout = {
  beforeBody: [
    Component.ConditionalRender({
      component: Component.Breadcrumbs(),
      condition: (page) => page.fileData.slug !== "index",
    }),
    Component.ArticleTitle(),
    Component.ContentMeta(),
    Component.TagList(),
  ],
  left: [
    Component.PageTitle(),
    Component.MobileOnly(Component.Spacer()),
    Component.Flex({
      components: [
        {
          // Pagefind instead of Component.Search(): Quartz's index would ship
          // all 9411 notes to the browser on first paint.
          Component: PagefindSearch(),
          grow: true,
        },
        // Rendered but hidden: the reading panel drives reader mode, and
        // keeping the component in the layout is what bundles its script
        // and styles.
        { Component: Component.ReaderMode() },
      ],
    }),
    // Quartz's default sort is already folders-first with numeric collation,
    // which is what the vault's `01. …`, `02. …` names need.
    Component.Explorer(),
  ],
  // Graph view intentionally omitted.
  right: [Component.DesktopOnly(Component.TableOfContents()), Component.Backlinks()],
}

// components for pages that display lists of pages  (e.g. tags or folders)
export const defaultListPageLayout: PageLayout = {
  beforeBody: [Component.Breadcrumbs(), Component.ArticleTitle(), Component.ContentMeta()],
  left: [
    Component.PageTitle(),
    Component.MobileOnly(Component.Spacer()),
    Component.Flex({
      components: [
        {
          Component: PagefindSearch(),
          grow: true,
        },
      ],
    }),
    // Quartz's default sort is already folders-first with numeric collation,
    // which is what the vault's `01. …`, `02. …` names need.
    Component.Explorer(),
  ],
  right: [],
}
