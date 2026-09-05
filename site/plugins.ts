import { QuartzTransformerPlugin } from "../types"
import { visit } from "unist-util-visit"
import { Root, Element, Text, Parent } from "hast"
import { VFile } from "vfile"

declare module "vfile" {
  interface DataMap {
    /** Set by ComingSoon on skeleton notes; Head turns it into `noindex`. */
    comingSoon: boolean
  }
}

/**
 * Site-specific transformers for the Principal SWE knowledge graph.
 *
 * Both run on the HTML (hast) tree, after markdown parsing, so they see the
 * final rendered shape of a note rather than its markdown source.
 */

/** Containers whose text must never be rewritten. */
const OPAQUE = new Set(["code", "pre", "kbd", "samp", "var", "math", "script", "style"])

/** Headings keep their normal weight — bionic markup only touches body copy. */
const NOT_BODY = new Set(["h1", "h2", "h3", "h4", "h5", "h6", ...OPAQUE])

function asClassList(node: Element): string[] {
  const cls = node.properties?.className
  return Array.isArray(cls) ? cls.map(String) : []
}

/**
 * How much of a word to embolden. Short words get a single character, longer
 * ones roughly the leading 40%, which is the ratio bionic reading is built on.
 */
function fixationLength(word: string): number {
  if (word.length <= 1) return word.length
  if (word.length <= 3) return 1
  return Math.ceil(word.length * 0.4)
}

interface BionicOptions {
  /** Words shorter than this are left alone, so the page keeps some rhythm. */
  minWordLength: number
}

const bionicDefaults: BionicOptions = {
  minWordLength: 3,
}

export const Bionic: QuartzTransformerPlugin<Partial<BionicOptions>> = (userOpts) => {
  const opts = { ...bionicDefaults, ...userOpts }

  return {
    name: "Bionic",
    htmlPlugins() {
      return [
        () => (tree: Root) => {
          visit(tree, "element", (node: Element) => {
            if (NOT_BODY.has(node.tagName)) return "skip"

            // Do not descend into markup this plugin just produced, or the
            // prefix would be re-emboldened on every pass down the tree.
            if (node.tagName === "b" && asClassList(node).includes("bionic")) return "skip"

            const next: (Element | Text)[] = []
            let changed = false

            for (const child of node.children as (Element | Text)[]) {
              if (child.type !== "text") {
                next.push(child)
                continue
              }

              // Split on whitespace but keep it, so spacing survives intact.
              const parts = child.value.split(/(\s+)/)
              const built: (Element | Text)[] = []

              for (const part of parts) {
                if (part.length === 0) continue

                // Leading punctuation must stay outside the bolded prefix.
                const match = /^([^\p{L}\p{N}]*)([\p{L}\p{N}][\p{L}\p{N}'’-]*)(.*)$/u.exec(part)
                if (!match || match[2].length < opts.minWordLength) {
                  built.push({ type: "text", value: part })
                  continue
                }

                const [, lead, word, trail] = match
                const cut = fixationLength(word)
                if (lead) built.push({ type: "text", value: lead })
                built.push({
                  type: "element",
                  tagName: "b",
                  properties: { className: ["bionic"] },
                  children: [{ type: "text", value: word.slice(0, cut) }],
                })
                built.push({ type: "text", value: word.slice(cut) + trail })
                changed = true
              }

              next.push(...built)
            }

            if (changed) {
              node.children = next
            }

            return undefined
          })
        },
      ]
    },
  }
}

interface ComingSoonOptions {
  /**
   * Notes with fewer body words than this are treated as skeletons — the
   * headings exist but nothing has been written under them yet. The vault
   * splits cleanly here: real notes carry hundreds of words, outlines none.
   */
  minWords: number
}

const comingSoonDefaults: ComingSoonOptions = {
  minWords: 15,
}

/** Text that every skeleton note carries, so it must not count as content. */
const BOILERPLATE = /^(⬆️|📚|🔗|References?|Parent|Module|Tags?)\b/u

export const ComingSoon: QuartzTransformerPlugin<Partial<ComingSoonOptions>> = (userOpts) => {
  const opts = { ...comingSoonDefaults, ...userOpts }

  return {
    name: "ComingSoon",
    htmlPlugins() {
      return [
        () => (tree: Root, file: VFile) => {
          let words = 0

          visit(tree, "element", (node: Element) => {
            // Only prose counts: headings and reference lists are scaffolding.
            if (!["p", "td", "blockquote"].includes(node.tagName)) return

            visit(node, "text", (text: Text) => {
              const value = text.value.trim()
              if (value.length === 0 || BOILERPLATE.test(value)) return
              words += value.split(/\s+/).length
            })
          })

          // Code samples are real content even without prose around them.
          let hasCode = false
          visit(tree, "element", (node: Element) => {
            if (node.tagName === "pre") hasCode = true
          })

          if (words >= opts.minWords || hasCode) return

          // An outline is thin content: worth reading in the tree, not worth
          // indexing. Head reads this back and emits `noindex, follow`.
          file.data.comingSoon = true

          const banner: Element = {
            type: "element",
            tagName: "div",
            properties: { className: ["coming-soon"] },
            children: [
              {
                type: "element",
                tagName: "p",
                properties: { className: ["coming-soon-label"] },
                children: [{ type: "text", value: "Coming soon" }],
              },
              {
                type: "element",
                tagName: "p",
                properties: { className: ["coming-soon-body"] },
                children: [
                  {
                    type: "text",
                    value:
                      "This note is an outline. The sections below are placeholders and have not been written yet.",
                  },
                ],
              },
            ],
          }

          ;(tree as Parent).children.unshift(banner)
        },
      ]
    },
  }
}
