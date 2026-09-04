#!/usr/bin/env bash
# Local preview of the published site.
#
# Mirrors .github/workflows/deploy.yml exactly, so a page that 404s here will
# 404 in production too. The Quartz checkout and the assembled content live
# outside the vault, in $WORK, so they never end up in git.
set -euo pipefail

QUARTZ_REF="v4.5.2"
VAULT_DIR="Principal SWE"
PORT="${PORT:-8081}"

VAULT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
WORK="${WORK:-$HOME/.cache/principal-swe-site}"
QUARTZ="$WORK/quartz"
CONTENT="$WORK/content"

mkdir -p "$WORK"

if [ ! -d "$QUARTZ/.git" ]; then
  echo "==> cloning quartz $QUARTZ_REF"
  git clone --depth 1 --branch "$QUARTZ_REF" https://github.com/jackyzha0/quartz "$QUARTZ"
fi

echo "==> checking links"
VAULT_CONTENT="$VAULT/$VAULT_DIR" python3 "$VAULT/tools/check_links.py"

echo "==> assembling content"
rm -rf "$CONTENT"
cp -R "$VAULT/$VAULT_DIR" "$CONTENT"
# Quartz needs an index.md at the content root for the landing page.
cp "$CONTENT/$VAULT_DIR.md" "$CONTENT/index.md"

echo "==> applying site overlay"
cp "$VAULT/site/quartz.config.ts" "$QUARTZ/quartz.config.ts"
cp "$VAULT/site/quartz.layout.ts" "$QUARTZ/quartz.layout.ts"
cp "$VAULT/site/custom.scss" "$QUARTZ/quartz/styles/custom.scss"
cp "$VAULT/site/plugins.ts" "$QUARTZ/quartz/plugins/transformers/site.ts"
cp "$VAULT/site/ReadingControls.tsx" "$QUARTZ/quartz/components/ReadingControls.tsx"
cp "$VAULT/site/readingControls.inline.ts" "$QUARTZ/quartz/components/scripts/readingControls.inline.ts"
cp "$VAULT/site/readingControls.scss" "$QUARTZ/quartz/components/styles/readingControls.scss"
cp "$VAULT/site/SiblingNav.tsx" "$QUARTZ/quartz/components/SiblingNav.tsx"
cp "$VAULT/site/siblingNav.scss" "$QUARTZ/quartz/components/styles/siblingNav.scss"
cp "$VAULT/site/PageList.tsx" "$QUARTZ/quartz/components/PageList.tsx"
python3 "$VAULT/tools/patch_slug.py" "$QUARTZ/quartz/util/path.ts"
python3 "$VAULT/tools/patch_folder_title.py" "$QUARTZ/quartz/plugins/emitters/folderPage.tsx"

if [ ! -d "$QUARTZ/node_modules" ]; then
  echo "==> npm ci"
  (cd "$QUARTZ" && npm ci)
fi

# Quartz's own dev server 404s on every folder page (it redirects
# `.../Folder/index.html` to `.../Folder`, which it then cannot resolve), so
# build once and serve the output with tools/preview_server.py, which resolves
# both folder `index.html` and the extensionless note URLs the way GitHub Pages
# does.
echo "==> building"
cd "$QUARTZ"
npx quartz build -d "$CONTENT" -o "$WORK/public"

echo "==> serving on http://localhost:$PORT"
exec python3 "$VAULT/tools/preview_server.py" "$WORK/public" --port "$PORT"
