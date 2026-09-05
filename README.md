# Principal SWE Knowledge Graph

An Obsidian vault of ~9,400 engineering notes for Staff+ / Principal engineers, published as a static site with [Quartz](https://quartz.jzhao.xyz).

- **Site:** https://bakhod1r.github.io/principal-swe-knowledge-graph
- **Entry note:** [`Principal SWE/Principal SWE.md`](Principal%20SWE/Principal%20SWE.md)

## Layout

```
Principal SWE/            the vault — the only thing published
├── 01. Foundations & Craftsmanship     Computer Science, Programming (7 languages)
├── 02. Architecture & System Design    Architecture, System Design, Best Practices
├── 03. Infrastructure & Security       DevOps, Cyber Security
├── 04. AI & Intelligent Systems        ML, LLMs, RAG, agents, inference
├── 05. Leadership & Soft Skills        tech lead, management, writing, influence
└── tools/                              vault-local Go tools (see below)

site/                     Quartz overlay: config, layout, custom components
tools/                    build + link-check scripts used by CI and local preview
.github/workflows/        deploy.yml (GitHub Pages), release.yml (tagged snapshots)
```

Each folder has a folder note of the same name, so the folder is browsable as a page on the site.

## Local preview

```bash
tools/serve.sh          # http://localhost:8081 (PORT=… to change)
```

It mirrors `deploy.yml` step for step, so a page that 404s locally 404s in production. Quartz is cloned and built under `~/.cache/principal-swe-site`, never inside the repo. Requires Node 22 and Python 3.

## Publishing

Push to `main` builds and deploys to GitHub Pages. The build:

1. copies `Principal SWE/` to `content/`, plus `Principal SWE.md` as `content/index.md` (the landing page);
2. runs `tools/check_links.py` — **any unresolved wikilink or relative link fails the build**, since it would be a 404 on the site;
3. copies the `site/` overlay over the Quartz checkout and applies three source patches:
   - `tools/patch_slug.py` — strips backticks from slugs (they reach the URL as `%60` and Pages 404s);
   - `tools/patch_folder_title.py` — makes folder-page titles readable instead of the escaped full path;
   - `tools/patch_content_index.py` — drops skeleton notes from the sitemap and RSS (see [SEO](#seo));
4. runs `npx quartz build`.

## SEO

[`site/Head.tsx`](site/Head.tsx) replaces Quartz's `Head` and adds what upstream leaves out:

- **`rel=canonical`** on every page, built from the simplified slug so a folder page has one address, not two. `canonicalSlug` in frontmatter overrides it — the root note uses it, because the build serves the same text at `/` and at `/Principal-SWE`.
- **`noindex, follow` on skeleton notes.** The `ComingSoon` transformer already marks any note with no prose under its headings; ~7,100 of ~12,700 pages are outlines, and indexing them as thin content would drag the rest down. `follow` still lets crawlers walk through them.
- **JSON-LD** — `WebSite` on the landing page, `CollectionPage` on folder and tag listings, `Article` on notes (headline, description, keywords, published/modified dates).
- **`og:locale`, `article:modified_time`, `keywords`,** and `og:image:type` without upstream's stray dot (`image/.png`).

[`tools/patch_content_index.py`](tools/patch_content_index.py) keeps the same skeleton notes out of `sitemap.xml` and the RSS feed — a sitemap that advertises noindexed URLs is a contradiction the crawler has to resolve. The sitemap ends up at ~2,300 real pages instead of 9,400.

`site/static/og-image.png` is the shared link-preview card; regenerate it with `tools/make_og_image.py` (macOS fonts, Pillow). Quartz's per-page `CustomOgImages` emitter stays off — 9k+ renders per build is far too slow.

No `robots.txt`: this is a GitHub *project* page, so `bakhod1r.github.io/robots.txt` belongs to the user site, not this repo. Submit `.../principal-swe-knowledge-graph/sitemap.xml` to Search Console directly.

`tools/preview_server.py` exists because Quartz's dev server 404s on folder pages; it resolves extensionless note URLs and folder `index.html` the way GitHub Pages does.

Pushing a `v*` tag runs `release.yml`, which publishes a snapshot of the graph with a note count in the release notes.

## Vault tools

Run from `Principal SWE/tools` (`make help` lists everything):

```bash
make stats  PATH_ARG=..            # empty-note statistics per topic
make random PATH_ARG="../01. Foundations & Craftsmanship"   # a random empty note to fill
make all    PATH_ARG=..            # list every empty note

make anki-preview NOTE="../path/Note.md"   # show the Anki cards, touch nothing
make anki-sync                             # import every note not in Anki yet
```

- `empty-note-finder/` — finds notes that are still stubs.
- `md-to-anki/` — one Anki card per numbered heading, pushed over [AnkiConnect](https://foosoft.net/projects/anki-connect/) (`127.0.0.1:8765`; add-on `2055492159`). Duplicates are detected collection-wide, so re-runs only add what is new. Tests: `make anki-test`.

## Conventions

- Notes are Markdown with YAML frontmatter (`title`, `tags`, `parent`).
- Links are Obsidian wikilinks; resolution is "shortest path", matching Quartz.
- Avoid `|`, `#`, `^` and backticks in filenames — Obsidian cannot link to them and Quartz cannot serve them. `check_links.py` rejects them.
- `.obsidian/`, `.trash/` and `tools/` are excluded from the published site via `ignorePatterns` in [`site/quartz.config.ts`](site/quartz.config.ts).
