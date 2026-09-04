#!/usr/bin/env python3
"""Give folder pages a readable title.

Quartz titles a folder page `Folder: <slug>`, and the slug is the escaped path —
`Folder: 01.-Foundations--and--Craftsmanship/Programming/09.-Language`. Only the
last segment matters to a reader, and it reads as a name once the slug escaping
is undone.

Run against the checked-out Quartz copy:
    patch_folder_title.py quartz/quartz/plugins/emitters/folderPage.tsx
Idempotent, and fails loudly if Quartz's source moves out from under it.
"""

import sys

ANCHOR = "title: `${i18n(locale).pages.folderContent.folder}: ${folder}`,"
REPLACEMENT = (
    "title: `${i18n(locale).pages.folderContent.folder}: ${folder\n"
    "            .split('/')\n"
    "            .pop()!\n"
    "            .replace(/-and-/g, ' & ')\n"
    "            .replace(/-percent/g, '%')\n"
    "            .replace(/-/g, ' ')\n"
    "            .replace(/\\s+/g, ' ')\n"
    "            .trim()}`,"
)


def main(path):
    with open(path, encoding="utf-8") as f:
        src = f.read()

    if "replace(/-and-/g, ' & ')" in src:
        print(f"patch_folder_title: {path} already patched")
        return 0

    if src.count(ANCHOR) != 1:
        print(
            f"patch_folder_title FAILED: expected exactly one folder-title line in {path}, "
            f"found {src.count(ANCHOR)} — Quartz's folderPage emitter changed shape",
            file=sys.stderr,
        )
        return 1

    with open(path, "w", encoding="utf-8") as f:
        f.write(src.replace(ANCHOR, REPLACEMENT))
    print(f"patch_folder_title: folder titles un-slugified in {path}")
    return 0


if __name__ == "__main__":
    if len(sys.argv) != 2:
        sys.exit("usage: patch_folder_title.py <path to quartz/plugins/emitters/folderPage.tsx>")
    sys.exit(main(sys.argv[1]))
