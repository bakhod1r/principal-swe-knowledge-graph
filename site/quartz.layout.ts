import { PageLayout, SharedLayout } from "./quartz/cfg"
import * as Component from "./quartz/components"
import ReadingControls from "./quartz/components/ReadingControls"

// components shared across all pages
export const sharedPageComponents: SharedLayout = {
  head: Component.Head(),
  header: [],
  // Floating reading dock, on every page type.
  afterBody: [ReadingControls()],
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
          Component: Component.Search(),
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
          Component: Component.Search(),
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
