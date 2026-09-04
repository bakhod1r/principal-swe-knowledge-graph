import { FullSlug, isFolderPath, resolveRelative } from "../util/path"
import { QuartzPluginData } from "../plugins/vfile"
import { QuartzComponent, QuartzComponentProps } from "./types"
import { GlobalConfiguration } from "../cfg"

export type SortFn = (f1: QuartzPluginData, f2: QuartzPluginData) => number

/**
 * Folder and tag listings, ordered the way the vault is written.
 *
 * Quartz sorts these by modified date, which on a vault where every file was
 * touched the same day produces arbitrary order — `09.`, `06.`, `10.`, `08.`.
 * The numbers in the names *are* the intended order, so sort naturally on the
 * title and drop the date column entirely: it carries no information here.
 */
const collator = new Intl.Collator("en", { numeric: true, sensitivity: "base" })

function titleOf(f: QuartzPluginData): string {
  return f.frontmatter?.title ?? decodeURIComponent((f.slug ?? "").split("/").pop() ?? "")
}

export function byNatural(_cfg?: GlobalConfiguration): SortFn {
  return (f1, f2) => collator.compare(titleOf(f1), titleOf(f2))
}

export function byDateAndAlphabetical(cfg: GlobalConfiguration): SortFn {
  // Kept for RecentNotes, which imports it by name.
  return byNatural(cfg)
}

export function byDateAndAlphabeticalFolderFirst(cfg: GlobalConfiguration): SortFn {
  const natural = byNatural(cfg)
  return (f1, f2) => {
    const f1IsFolder = isFolderPath(f1.slug ?? "")
    const f2IsFolder = isFolderPath(f2.slug ?? "")
    if (f1IsFolder && !f2IsFolder) return -1
    if (!f1IsFolder && f2IsFolder) return 1
    return natural(f1, f2)
  }
}

type Props = {
  limit?: number
  sort?: SortFn
} & QuartzComponentProps

export const PageList: QuartzComponent = ({ cfg, fileData, allFiles, limit, sort }: Props) => {
  const sorter = sort ?? byDateAndAlphabeticalFolderFirst(cfg)
  let list = allFiles.sort(sorter)
  if (limit) {
    list = list.slice(0, limit)
  }

  return (
    <ul class="section-ul">
      {list.map((page) => {
        const tags = page.frontmatter?.tags ?? []

        return (
          <li class="section-li">
            <div class="section">
              <div class="desc">
                <h3>
                  <a href={resolveRelative(fileData.slug!, page.slug!)} class="internal">
                    {titleOf(page)}
                  </a>
                </h3>
              </div>
              <ul class="tags">
                {tags.map((tag) => (
                  <li>
                    <a
                      class="internal tag-link"
                      href={resolveRelative(fileData.slug!, `tags/${tag}` as FullSlug)}
                    >
                      {tag}
                    </a>
                  </li>
                ))}
              </ul>
            </div>
          </li>
        )
      })}
    </ul>
  )
}

PageList.css = `
.section h3 {
  margin: 0;
}

.section > .tags {
  margin: 0;
}

/* Two columns now that the date column is gone. */
li.section-li > .section {
  grid-template-columns: 3fr 1fr;
}

.popover .section {
  grid-template-columns: 1fr !important;
}
`
