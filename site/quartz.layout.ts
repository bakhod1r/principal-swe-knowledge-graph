import { PageLayout, SharedLayout } from "./quartz/cfg"
import * as Component from "./quartz/components"
import ReadingControls from "./quartz/components/ReadingControls"
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
    // No search: the tree is the way through the vault, and Quartz's search
    // would ship an index of all 9411 notes to every visitor.
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
    // Quartz's default sort is already folders-first with numeric collation,
    // which is what the vault's `01. …`, `02. …` names need.
    Component.Explorer(),
  ],
  right: [],
}
