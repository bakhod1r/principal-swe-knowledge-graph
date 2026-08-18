---
title: api
tags:
  - golang
  - goroot
  - compatibility
  - stdlib
parent: "[[Installation Directories]]"
---

# `$GOROOT/api`

Machine-checked records of the standard library's exported surface, one file per
Go release. The enforcement mechanism behind the **Go 1 compatibility promise**.

## 1. Contents

```text
$GOROOT/api/
├── go1.txt        ← everything exported in Go 1.0
├── go1.1.txt
├── ...
├── go1.26.txt     ← what 1.26 added
├── next/          ← pending additions on the development branch
└── except.txt     ← the small, deliberate list of approved breaks
```

## 2. Line Format

```text
pkg net/http, func NewResponseController(ResponseWriter) *ResponseController #54136
pkg slices, func Sorted[$0 interface{ ~... }](iter.Seq[$0]) []$0 #61899
```

Package, declaration, and the issue number that approved it.

## 3. How It Is Enforced

`cmd/api` runs during the release test suite. Removing or changing an exported
declaration fails the build unless the change is listed in `except.txt`. That is
why a Go upgrade essentially never breaks compilation.

```bash
grep 'pkg slices' "$(go env GOROOT)/api/go1.21.txt"   # what slices added in 1.21
```

## 4. Practical Use

Answering "which Go version introduced this function?":

```bash
grep -rn 'func Sorted' "$(go env GOROOT)/api/"
```

Faster and more authoritative than searching release notes.

## 5. Gotchas

- The promise covers **compilation**, not behaviour. Behavioural changes are
  gated separately by `godebug` and `GODEBUG`.
- `internal/` packages are excluded — see `Internal Packages`.
- Additions are allowed; only removals and signature changes are blocked.

---

## 🔗 References
- ⬆️ Parent: [[Installation Directories]]
