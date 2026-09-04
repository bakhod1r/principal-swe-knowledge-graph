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

if [ ! -d "$QUARTZ/node_modules" ]; then
  echo "==> npm ci"
  (cd "$QUARTZ" && npm ci)
fi

echo "==> serving on http://localhost:$PORT"
cd "$QUARTZ"
exec npx quartz build -d "$CONTENT" -o "$WORK/public" --serve --port "$PORT"
