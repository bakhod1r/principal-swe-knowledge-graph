#!/usr/bin/env python3
"""Keep skeleton notes out of the sitemap and the RSS feed.

The ComingSoon transformer (site/plugins.ts) marks every note that is still an
outline, and site/Head.tsx serves those with `noindex`. A sitemap that lists
noindexed URLs is a contradiction a crawler has to resolve on its own, so the
same flag has to reach Quartz's ContentIndex — and that emitter has no option
for it.

Run against the checked-out Quartz copy:
    patch_content_index.py quartz/quartz/plugins/emitters/contentIndex.tsx
Idempotent, and fails loudly if Quartz's source moves out from under it.
"""

import sys

ANCHOR = 'if (opts?.includeEmptyFiles || (file.data.text && file.data.text !== "")) {'
REPLACEMENT = (
    "const canonical = file.data.frontmatter?.canonicalSlug as string | undefined\n"
    "        const isCanonical = canonical === undefined || canonical === simplifySlug(slug)\n"
    "        if (\n"
    "          !file.data.comingSoon &&\n"
    "          isCanonical &&\n"
    '          (opts?.includeEmptyFiles || (file.data.text && file.data.text !== ""))\n'
    "        ) {"
)
MARKER = "!file.data.comingSoon"


def main(path):
    with open(path, encoding="utf-8") as f:
        source = f.read()

    if MARKER in source:
        print(f"patch_content_index: already applied to {path}")
        return

    if source.count(ANCHOR) != 1:
        sys.exit(
            f"patch_content_index: expected exactly one occurrence of\n  {ANCHOR}\n"
            f"in {path}; Quartz's source has changed and the patch needs revisiting."
        )

    with open(path, "w", encoding="utf-8") as f:
        f.write(source.replace(ANCHOR, REPLACEMENT))
    print(f"patch_content_index: patched {path}")


if __name__ == "__main__":
    if len(sys.argv) != 2:
        sys.exit("usage: patch_content_index.py <path to contentIndex.tsx>")
    main(sys.argv[1])
