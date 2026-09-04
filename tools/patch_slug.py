#!/usr/bin/env python3
"""Strip backticks inside Quartz's slug function.

Quartz's `sluggify` leaves a backtick untouched, so it reaches the URL as %60
and GitHub Pages answers 404 for that path — every other punctuation mark the
vault uses (: [ ] * < > & parens) serves fine. tools/check_links.py rejects
backticked filenames before the build; this patch makes an escaped one harmless
instead of invisible.

Run against the checked-out Quartz copy: patch_slug.py quartz/quartz/util/path.ts
Idempotent, and fails loudly if Quartz's source moves out from under it.
"""

import sys

ANCHOR = '.replace(/\\s/g, "-")'
INSERT = '\n        .replace(/`/g, "")'


def main(path):
    with open(path, encoding="utf-8") as f:
        src = f.read()

    if '.replace(/`/g, "")' in src:
        print(f"patch_slug: {path} already patched")
        return 0

    if src.count(ANCHOR) != 1:
        print(
            f"patch_slug FAILED: expected exactly one `{ANCHOR}` in {path}, "
            f"found {src.count(ANCHOR)} — Quartz's slugify changed shape",
            file=sys.stderr,
        )
        return 1

    with open(path, "w", encoding="utf-8") as f:
        f.write(src.replace(ANCHOR, ANCHOR + INSERT))
    print(f"patch_slug: backtick stripping added to {path}")
    return 0


if __name__ == "__main__":
    if len(sys.argv) != 2:
        sys.exit("usage: patch_slug.py <path to quartz/util/path.ts>")
    sys.exit(main(sys.argv[1]))
