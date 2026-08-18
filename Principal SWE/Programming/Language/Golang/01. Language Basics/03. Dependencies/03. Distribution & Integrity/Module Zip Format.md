---
title: Module Zip Format
tags:
  - golang
  - basics
  - dependencies
  - modules
  - proxy
  - format
parent: "[[Distribution & Integrity]]"
---

# Module Zip Format

The unit of distribution: one deterministic `.zip` per module version, whose
hash is what `go.sum` and `GOSUMDB` actually attest to.

## 1. Layout

Every path inside is prefixed with `module@version`:

```text
github.com/me/lib@v1.2.3/go.mod
github.com/me/lib@v1.2.3/user/user.go
github.com/me/lib@v1.2.3/LICENSE
```

## 2. Rules That Make It Deterministic

| Rule | Reason |
|---|---|
| No `vendor/` directories from dependencies | Dependencies are fetched separately |
| Nested modules excluded | Directories with their own `go.mod` belong elsewhere |
| No symlinks, no irregular files | Reproducible extraction on every OS |
| Case-conflicting paths rejected | Safe on case-insensitive filesystems |
| 500 MB / 16 MB (`go.mod`) size caps | Denial-of-service protection |

Because content is fully normalized, the same tag always produces byte-identical
zips — which is what makes hash verification meaningful.

## 3. Inspecting One

```bash
go mod download -json github.com/go-chi/chi/v5@v5.1.0
unzip -l "$(go env GOMODCACHE)/cache/download/github.com/go-chi/chi/v5/@v/v5.1.0.zip"
```

The `.info`, `.mod`, `.zip`, and `.ziphash` files all live under
`$GOMODCACHE/cache/download/`.

## 4. Gotchas

- The hash in `go.sum` is of a **normalized directory tree**, not the raw zip
  bytes — `sha256sum` on the file will not match.
- Extracted modules in `GOMODCACHE` are read-only on purpose; edit them and
  `go mod verify` fails. Use `replace` instead.
- `go mod download` populates the cache without building — the right warm-up step
  in a Docker layer.

---

## 🔗 References
- ⬆️ Parent: `03. Distribution & Integrity`
