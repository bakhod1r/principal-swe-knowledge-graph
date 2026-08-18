---
title: Internal Visibility Rule
tags:
  - golang
  - goroot
  - packages
  - encapsulation
parent: "[[Internal Packages]]"
---

# The `internal` Visibility Rule

A compiler-enforced access boundary based purely on directory position. Go's only
mechanism for package-level encapsulation beyond exported/unexported identifiers.

## 1. The Rule

> A package under a directory named `internal` may be imported **only** by code
> rooted in the parent of that `internal` directory.

```text
github.com/me/api/
├── internal/
│   └── store/          ← importable by anything under github.com/me/api/
├── cmd/api/            ✅ can import internal/store
└── pkg/client/         ✅ can import internal/store

github.com/you/app/     ❌ cannot import github.com/me/api/internal/store
```

## 2. The Error

```text
use of internal package github.com/me/api/internal/store not allowed
```

Enforced by the `go` command at load time, before compilation.

## 3. Why It Is the Most Useful Layout Rule

Exported identifiers in a non-internal package are a **permanent API commitment**
— anyone can import them, so changing them breaks strangers. Moving the same code
under `internal/` makes it freely refactorable while staying exported within your
own module.

```text
pkg/       → public contract, semver applies
internal/  → yours, change it freely
```

## 4. Nesting

The rule applies at **every** level:

```text
api/service/internal/cache   ← importable only under api/service/
```

This scopes a helper to one subsystem rather than the whole module.

## 5. Gotchas

- The boundary is the **directory tree**, not the module. Two modules in one
  repository do not share each other's `internal/`.
- The standard library uses it heavily — see `internal` and
  [[Internal Packages]]. That is why you cannot import `internal/abi`.
- `vendor/` copies of internal packages keep their restriction.

---

## 🔗 References
- ⬆️ Parent: [[Internal Packages]]
