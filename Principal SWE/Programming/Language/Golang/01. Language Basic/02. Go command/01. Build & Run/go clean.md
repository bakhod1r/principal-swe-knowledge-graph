---
title: Clean Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - clean
parent: "[[Build & Run]]"
---

# `go clean`

Removes build artifacts and, with flags, the various caches.

## 1. Usage

```bash
go clean                # object files in the package directory
go clean -i ./...       # also remove installed binaries/archives
go clean -r ./...       # recurse into dependencies
go clean -n             # dry run: print what would be removed
```

## 2. Cache Flags

| Flag | Removes | Location |
|---|---|---|
| `-cache` | build cache | `GOCACHE` |
| `-testcache` | cached test results only | `GOCACHE` |
| `-modcache` | downloaded modules | `GOMODCACHE` |
| `-fuzzcache` | fuzzing corpus | `GOCACHE` |

```bash
go clean -cache -testcache -modcache
```

## 3. Gotchas

- `-modcache` deletes read-only files; on older Go versions this could fail with
  permission errors — `go clean -modcache` handles it, manual `rm -rf` may not.
- Clearing the build cache makes the next build slow but never incorrect. If a
  build seems wrong, the bug is almost never the cache.
- `go clean -testcache` is the right way to force test re-runs; `-count=1` on
  `go test` does the same for one invocation.

---

## 🔗 References
- ⬆️ Parent: [[Build & Run]]
