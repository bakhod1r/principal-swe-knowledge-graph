#!/usr/bin/env python3
"""Fail the build when a note links somewhere the site cannot serve.

Every unresolved wikilink or relative markdown link becomes a 404 on the
published site, so this runs in CI before Quartz builds. Resolution mirrors
Quartz's "shortest" markdownLinkResolution: a target matches on full path,
on unique basename, or relative to the linking note.
"""

import collections
import os
import re
import sys

ROOT = os.environ.get("VAULT_CONTENT", "Principal SWE")
IGNORE = {".obsidian", ".trash", "templates", "private", "tools"}

# Obsidian stops a wikilink target at |, # or ^, so a filename containing one
# of those can never be linked to — matching that here is what catches it.
WIKILINK = re.compile(r"!?\[\[([^\]\|#\^]+)([^\]]*)\]\]")
MDLINK = re.compile(r"(?<!!)\[[^\]]*\]\(([^)]+)\)")
EXTERNAL = re.compile(r"^(https?:|mailto:|tel:|#)")


def collect(root):
    paths = []
    for dirpath, dirnames, filenames in os.walk(root):
        dirnames[:] = [d for d in dirnames if d not in IGNORE]
        for f in filenames:
            if f.endswith(".md"):
                paths.append(os.path.relpath(os.path.join(dirpath, f), root)[:-3])
    return paths


def main():
    if not os.path.isdir(ROOT):
        sys.exit(f"content directory not found: {ROOT}")

    paths = collect(ROOT)
    pathset = set(paths)
    basenames = {os.path.basename(p) for p in paths}

    def resolves(target, source):
        t = target.strip().rstrip("/")
        if not t or t in pathset or t in basenames:
            return True
        directory = os.path.dirname(source)
        candidate = os.path.normpath(os.path.join(directory, t) if directory else t)
        if candidate in pathset:
            return True
        # Attachments and other non-markdown files.
        return os.path.exists(os.path.join(ROOT, t)) or os.path.exists(
            os.path.join(ROOT, t + ".md")
        )

    broken = collections.defaultdict(set)
    for p in paths:
        text = open(os.path.join(ROOT, p + ".md"), encoding="utf-8", errors="replace").read()
        text = re.sub(r"```.*?```", "", text, flags=re.S)
        text = re.sub(r"`[^`\n]*`", "", text)

        targets = [m.group(1) for m in WIKILINK.finditer(text)]
        for m in MDLINK.finditer(text):
            url = m.group(1)
            if EXTERNAL.match(url):
                continue
            url = url.split("#")[0].split("?")[0]
            targets.append(url[:-3] if url.endswith(".md") else url)

        for t in targets:
            if t.strip() and not resolves(t, p):
                broken[t.strip()].add(p)

    if not broken:
        print(f"link check: {len(paths)} notes, no broken links")
        return 0

    print(f"link check FAILED: {len(broken)} broken target(s)", file=sys.stderr)
    for target, sources in sorted(broken.items()):
        print(f"  [[{target}]]", file=sys.stderr)
        for s in sorted(sources)[:5]:
            print(f"      from {s}.md", file=sys.stderr)
    return 1


if __name__ == "__main__":
    sys.exit(main())
