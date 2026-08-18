---
title: go mod vendor
tags:
  - golang
  - basics
  - cli
  - toolchain
  - modules
  - vendor
parent: "[[Module Commands]]"
---

# `go mod vendor`

Copies the packages needed to build and test the main module into `vendor/`.
Concept and trade-offs: `Vendoring`.

```bash
go mod vendor
go mod vendor -v      # print copied modules and packages
go mod vendor -e      # continue despite package load errors
```

## 1. What Gets Copied

Only packages that are **actually imported** — not whole modules. A dependency
with 40 packages of which you import 2 contributes 2 directories.

```text
vendor/
├── modules.txt
└── github.com/
    └── go-chi/chi/v5/
        ├── chi.go
        └── middleware/
```

## 2. `modules.txt` Is a Contract

The go command cross-checks it against `go.mod` on every build:

```text
go: inconsistent vendoring in /src:
        github.com/x/y@v1.2.3: is explicitly required in go.mod,
        but not marked as explicit in vendor/modules.txt
```

That error means "you edited `go.mod` without re-running `go mod vendor`", which
is exactly what the check is for.

## 3. Gotchas

- Must be re-run after **every** `go get` or `go mod tidy`.
- Non-Go assets (testdata, templates, `.proto` files) are **not** vendored unless
  a Go file in the same directory is needed. Code generators break on this.
- `LICENSE` files are copied — which is precisely why legal teams like vendoring.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
