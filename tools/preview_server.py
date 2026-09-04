#!/usr/bin/env python3
"""Static preview server that resolves Quartz's extensionless URLs.

Quartz emits `Page.html`; every internal link points at `Page` with no
extension, because GitHub Pages resolves that itself. `python3 -m http.server`
does not, so a plain static preview 404s on every note while folder pages —
which are real `index.html` files — keep working.

    preview_server.py <root> [--port 8081]
"""

import argparse
import functools
import http.server
import os
import posixpath
import urllib.parse


class PrettyURLHandler(http.server.SimpleHTTPRequestHandler):
    def translate_path(self, path):
        local = super().translate_path(path)
        if os.path.isdir(local) or os.path.exists(local):
            return local

        # Only extensionless paths get the .html fallback; a missing .css stays
        # a 404 rather than silently serving HTML.
        decoded = urllib.parse.unquote(urllib.parse.urlsplit(path).path)
        if not decoded.endswith("/") and not posixpath.splitext(decoded)[1]:
            candidate = local + ".html"
            if os.path.isfile(candidate):
                return candidate
        return local


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("root")
    ap.add_argument("--port", type=int, default=8081)
    ap.add_argument("--bind", default="127.0.0.1")
    args = ap.parse_args()

    handler = functools.partial(PrettyURLHandler, directory=args.root)
    server = http.server.ThreadingHTTPServer((args.bind, args.port), handler)
    print(f"serving {args.root} on http://{args.bind}:{args.port}")
    server.serve_forever()


if __name__ == "__main__":
    main()
