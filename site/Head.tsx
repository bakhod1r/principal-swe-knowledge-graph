import { i18n } from "../i18n"
import {
  FullSlug,
  SimpleSlug,
  getFileExtension,
  joinSegments,
  pathToRoot,
  simplifySlug,
} from "../util/path"
import { CSSResourceToStyleElement, JSResourceToScriptElement } from "../util/resources"
import { googleFontHref, googleFontSubsetHref } from "../util/theme"
import { QuartzComponent, QuartzComponentConstructor, QuartzComponentProps } from "./types"
import { unescapeHTML } from "../util/escape"
import { CustomOgImagesEmitterName } from "../plugins/emitters/ogImage"

/**
 * Quartz's Head with the SEO the upstream one leaves out:
 *
 *   - a canonical URL, built from the simplified slug so a folder page and its
 *     `/index` spelling do not read as two pages;
 *   - `noindex, follow` on skeleton notes (ComingSoon marks them) — thousands of
 *     outlines indexed as thin content would drag the whole site down, and
 *     `follow` still lets the crawler walk through them to the real ones;
 *   - JSON-LD: WebSite on the landing page, Article on every note;
 *   - `og:locale` and `article:modified_time`.
 *
 * Everything else is upstream's, kept in upstream's order so a Quartz bump is a
 * readable diff.
 */
export default (() => {
  const Head: QuartzComponent = ({
    cfg,
    fileData,
    externalResources,
    ctx,
  }: QuartzComponentProps) => {
    const titleSuffix = cfg.pageTitleSuffix ?? ""
    const title =
      (fileData.frontmatter?.title ?? i18n(cfg.locale).propertyDefaults.title) + titleSuffix
    const description =
      fileData.frontmatter?.socialDescription ??
      fileData.frontmatter?.description ??
      unescapeHTML(fileData.description?.trim() ?? i18n(cfg.locale).propertyDefaults.description)

    const { css, js, additionalHead } = externalResources

    const url = new URL(`https://${cfg.baseUrl ?? "example.com"}`)
    const path = url.pathname as FullSlug
    const baseDir = fileData.slug === "404" ? path : pathToRoot(fileData.slug!)
    const iconPath = joinSegments(baseDir, "static/icon.png")

    // simplifySlug drops the trailing `index`, so a folder page has exactly one
    // address. Upstream uses the raw slug here, which yields two.
    //
    // `canonicalSlug` in frontmatter overrides it, which is what the vault's
    // root note needs: the build copies it to `index.md`, so the same text is
    // served at `/` and at `/Principal-SWE`. Both carry the frontmatter, so both
    // point search engines at `/`.
    const frontmatterCanonical = fileData.frontmatter?.canonicalSlug as string | undefined
    const canonicalSlug =
      fileData.slug === "404"
        ? ("" as SimpleSlug)
        : ((frontmatterCanonical ?? simplifySlug(fileData.slug!)) as SimpleSlug)
    const socialUrl =
      fileData.slug === "404" ? url.toString() : joinSegments(url.toString(), canonicalSlug)

    const usesCustomOgImage = ctx.cfg.plugins.emitters.some(
      (e) => e.name === CustomOgImagesEmitterName,
    )
    const ogImageDefaultPath = `https://${cfg.baseUrl}/static/og-image.png`

    const isHome = fileData.slug === "index"
    // A 404 must never be indexed; nor must an outline that has no prose yet.
    const noIndex = fileData.slug === "404" || fileData.comingSoon === true
    const tags = fileData.frontmatter?.tags ?? []
    const modified = fileData.dates?.modified
    const published = fileData.dates?.published ?? fileData.dates?.created

    const jsonLd = isHome
      ? {
          "@context": "https://schema.org",
          "@type": "WebSite",
          name: cfg.pageTitle,
          url: url.toString(),
          description,
          inLanguage: cfg.locale,
        }
      : {
          "@context": "https://schema.org",
          "@type": "Article",
          headline: title,
          description,
          url: socialUrl,
          inLanguage: cfg.locale,
          keywords: tags.length > 0 ? tags.join(", ") : undefined,
          datePublished: published?.toISOString(),
          dateModified: modified?.toISOString(),
          isPartOf: {
            "@type": "WebSite",
            name: cfg.pageTitle,
            url: url.toString(),
          },
        }

    return (
      <head>
        <title>{title}</title>
        <meta charSet="utf-8" />
        {cfg.theme.cdnCaching && cfg.theme.fontOrigin === "googleFonts" && (
          <>
            <link rel="preconnect" href="https://fonts.googleapis.com" />
            <link rel="preconnect" href="https://fonts.gstatic.com" />
            <link rel="stylesheet" href={googleFontHref(cfg.theme)} />
            {cfg.theme.typography.title && (
              <link rel="stylesheet" href={googleFontSubsetHref(cfg.theme, cfg.pageTitle)} />
            )}
          </>
        )}
        <link rel="preconnect" href="https://cdnjs.cloudflare.com" crossOrigin="anonymous" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />

        <meta name="og:site_name" content={cfg.pageTitle}></meta>
        <meta property="og:title" content={title} />
        <meta property="og:type" content={isHome ? "website" : "article"} />
        <meta property="og:locale" content={cfg.locale.replace("-", "_")} />
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:title" content={title} />
        <meta name="twitter:description" content={description} />
        <meta property="og:description" content={description} />
        <meta property="og:image:alt" content={description} />
        {modified && <meta property="article:modified_time" content={modified.toISOString()} />}

        {!usesCustomOgImage && (
          <>
            <meta property="og:image" content={ogImageDefaultPath} />
            <meta property="og:image:url" content={ogImageDefaultPath} />
            <meta name="twitter:image" content={ogImageDefaultPath} />
            {/* getFileExtension keeps the dot, so upstream emits `image/.png`. */}
            <meta
              property="og:image:type"
              content={`image/${(getFileExtension(ogImageDefaultPath) ?? ".png").replace(/^\./, "")}`}
            />
          </>
        )}

        {cfg.baseUrl && (
          <>
            <meta property="twitter:domain" content={cfg.baseUrl}></meta>
            <meta property="og:url" content={socialUrl}></meta>
            <meta property="twitter:url" content={socialUrl}></meta>
            <link rel="canonical" href={socialUrl} />
          </>
        )}

        <link rel="icon" href={iconPath} />
        <meta name="description" content={description} />
        {tags.length > 0 && <meta name="keywords" content={tags.join(", ")} />}
        <meta name="robots" content={noIndex ? "noindex, follow" : "index, follow"} />
        <meta name="generator" content="Quartz" />
        <script
          type="application/ld+json"
          dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
        />

        {css.map((resource) => CSSResourceToStyleElement(resource, true))}
        {js
          .filter((resource) => resource.loadTime === "beforeDOMReady")
          .map((res) => JSResourceToScriptElement(res, true))}
        {additionalHead.map((resource) => {
          if (typeof resource === "function") {
            return resource(fileData)
          } else {
            return resource
          }
        })}
      </head>
    )
  }

  return Head
}) satisfies QuartzComponentConstructor
