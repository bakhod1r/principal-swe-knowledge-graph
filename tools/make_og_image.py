#!/usr/bin/env python3
"""Draw site/static/og-image.png — the link preview every page falls back to.

Quartz's CustomOgImages emitter renders one image per page, which 9k+ notes make
far too slow, so the whole site shares this single card. Colours are the light
theme from site/quartz.config.ts.

    tools/make_og_image.py            # rewrites site/static/og-image.png

Serif faces come from macOS; the script is a one-off asset generator, not part
of the build, so it is fine for it to be macOS-only.
"""

import os

from PIL import Image, ImageDraw, ImageFont

W, H = 1200, 630
BG = "#f7f4ee"
INK = "#191510"
MUTED = "#8c8579"
ACCENT = "#8a3324"

BOLD = "/System/Library/Fonts/Supplemental/Georgia Bold.ttf"
REGULAR = "/System/Library/Fonts/Supplemental/Georgia.ttf"

TITLE = "Principal SWE"
SUBTITLE = "Knowledge Graph"
TAGLINE = "Systems, architecture, infrastructure, AI and leadership —"
TAGLINE2 = "an engineering vault for Staff+ engineers."

OUT = os.path.join(os.path.dirname(os.path.abspath(__file__)), "..", "site", "static", "og-image.png")


def main():
    img = Image.new("RGB", (W, H), BG)
    d = ImageDraw.Draw(img)

    # Accent rule down the left edge, the same red the site uses for links.
    d.rectangle([0, 0, 16, H], fill=ACCENT)

    title = ImageFont.truetype(BOLD, 96)
    subtitle = ImageFont.truetype(REGULAR, 72)
    body = ImageFont.truetype(REGULAR, 34)
    label = ImageFont.truetype(BOLD, 26)

    x = 96
    d.text((x, 118), TITLE, font=title, fill=INK)
    d.text((x, 232), SUBTITLE, font=subtitle, fill=ACCENT)

    d.line([x, 372, W - 96, 372], fill="#ded7c9", width=2)

    d.text((x, 410), TAGLINE, font=body, fill=INK)
    d.text((x, 456), TAGLINE2, font=body, fill=INK)

    d.text((x, 540), "bakhod1r.github.io/principal-swe-knowledge-graph", font=label, fill=MUTED)

    os.makedirs(os.path.dirname(OUT), exist_ok=True)
    img.save(os.path.normpath(OUT), "PNG", optimize=True)
    print(f"wrote {os.path.normpath(OUT)} ({W}x{H})")


if __name__ == "__main__":
    main()
