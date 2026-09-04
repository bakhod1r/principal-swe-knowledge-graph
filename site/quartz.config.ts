import { QuartzConfig } from "./quartz/cfg"
import * as Plugin from "./quartz/plugins"
import { Bionic, ComingSoon } from "./quartz/plugins/transformers/site"

/**
 * Quartz 4 Configuration
 *
 * See https://quartz.jzhao.xyz/configuration for more information.
 */
const config: QuartzConfig = {
  configuration: {
    pageTitle: "Principal SWE Knowledge Graph",
    pageTitleSuffix: "",
    enableSPA: true,
    enablePopovers: true,
    analytics: null,
    locale: "en-US",
    baseUrl: "bakhod1r.github.io/principal-swe-knowledge-graph",
    ignorePatterns: ["private", "templates", ".obsidian", ".trash", "tools", "**/.DS_Store"],
    defaultDateType: "modified",
    theme: {
      fontOrigin: "googleFonts",
      cdnCaching: true,
      typography: {
        header: "Newsreader",
        body: "Source Serif 4",
        code: "IBM Plex Mono",
      },
      colors: {
        lightMode: {
          light: "#f7f4ee",
          lightgray: "#ded7c9",
          gray: "#8c8579",
          darkgray: "#3a352d",
          dark: "#191510",
          secondary: "#8a3324",
          tertiary: "#a8552f",
          highlight: "rgba(138, 51, 36, 0.07)",
          textHighlight: "#e8c99b88",
        },
        darkMode: {
          light: "#15130f",
          lightgray: "#2e2a23",
          gray: "#7d7568",
          darkgray: "#d5cec2",
          dark: "#f3ece0",
          secondary: "#d98e6a",
          tertiary: "#e0a882",
          highlight: "rgba(217, 142, 106, 0.10)",
          textHighlight: "#8a663055",
        },
      },
    },
  },
  plugins: {
    transformers: [
      Plugin.FrontMatter(),
      Plugin.CreatedModifiedDate({
        priority: ["frontmatter", "filesystem"],
      }),
      Plugin.SyntaxHighlighting({
        theme: {
          light: "github-light",
          dark: "github-dark",
        },
        keepBackground: false,
      }),
      Plugin.ObsidianFlavoredMarkdown({ enableInHtmlEmbed: false }),
      Plugin.GitHubFlavoredMarkdown(),
      Plugin.TableOfContents(),
      Plugin.CrawlLinks({ markdownLinkResolution: "shortest" }),
      Plugin.Description(),
      Plugin.Latex({ renderEngine: "katex" }),
      // Order matters: ComingSoon counts words before Bionic splits them up.
      ComingSoon(),
      Bionic(),
    ],
    filters: [Plugin.RemoveDrafts()],
    emitters: [
      Plugin.AliasRedirects(),
      Plugin.ComponentResources(),
      Plugin.ContentPage(),
      Plugin.FolderPage(),
      Plugin.TagPage(),
      Plugin.ContentIndex({
        enableSiteMap: true,
        enableRSS: true,
      }),
      Plugin.Assets(),
      Plugin.Static(),
      Plugin.Favicon(),
      Plugin.NotFoundPage(),
      // CustomOgImages disabled: 9k+ notes would make builds very slow
      // Plugin.CustomOgImages(),
    ],
  },
}

export default config
