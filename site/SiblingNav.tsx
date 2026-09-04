import { QuartzComponent, QuartzComponentConstructor, QuartzComponentProps } from "./types"
import { classNames } from "../util/lang"
import { FullSlug, resolveRelative } from "../util/path"
import style from "./styles/siblingNav.scss"

/**
 * Previous / next links across the notes in the same folder.
 *
 * The vault is written to be read in order — folders and files are numbered
 * `01. …`, `02. …` — but the site only offered the tree and backlinks, so
 * reading straight through meant a trip back to the sidebar after every note.
 * Ordering matches the Explorer: natural (numeric-aware) sort on the file name.
 */
const collator = new Intl.Collator("en", { numeric: true, sensitivity: "base" })

function parentOf(slug: string): string {
  const idx = slug.lastIndexOf("/")
  return idx === -1 ? "" : slug.slice(0, idx)
}

function titleOf(file: { slug?: string; frontmatter?: { title?: string } }): string {
  return file.frontmatter?.title ?? decodeURIComponent((file.slug ?? "").split("/").pop() ?? "")
}

const SiblingNav: QuartzComponent = ({
  fileData,
  allFiles,
  displayClass,
}: QuartzComponentProps) => {
  const slug = fileData.slug
  if (!slug || slug === "index" || slug.endsWith("/index")) return null

  const folder = parentOf(slug)
  const siblings = allFiles
    .filter((f) => {
      const s = f.slug
      if (!s || s === slug) return false
      // Folder index pages are the parent of the list, not a member of it.
      if (s === "index" || s.endsWith("/index")) return false
      return parentOf(s) === folder
    })
    .concat(allFiles.filter((f) => f.slug === slug))
    .sort((a, b) => collator.compare(a.slug!, b.slug!))

  const here = siblings.findIndex((f) => f.slug === slug)
  if (here === -1 || siblings.length < 2) return null

  const prev = siblings[here - 1]
  const next = siblings[here + 1]
  if (!prev && !next) return null

  const link = (file: (typeof siblings)[number], dir: "prev" | "next") => (
    <a
      class={`sibling-link sibling-${dir}`}
      href={resolveRelative(slug as FullSlug, file.slug as FullSlug)}
    >
      <span class="sibling-dir">{dir === "prev" ? "Previous" : "Next"}</span>
      <span class="sibling-title">{titleOf(file)}</span>
    </a>
  )

  return (
    <nav class={classNames(displayClass, "sibling-nav")} aria-label="Sibling notes">
      {prev ? link(prev, "prev") : <span class="sibling-link sibling-empty" />}
      {next ? link(next, "next") : <span class="sibling-link sibling-empty" />}
    </nav>
  )
}

SiblingNav.css = style

export default (() => SiblingNav) satisfies QuartzComponentConstructor
