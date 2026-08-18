---
title: Incompatible Versions
tags:
  - golang
  - basics
  - dependencies
  - modules
  - versioning
  - legacy
parent: "[[Version Resolution]]"
---

# `+incompatible` Versions

A compatibility escape hatch for repositories tagged `v2.0.0` or higher that
never adopted the `Major Version Suffix`.

## 1. What It Looks Like

```go
require github.com/old/lib v3.2.1+incompatible
```

## 2. When the Go Command Generates It

```text
repo tagged v3.2.1
        │
        ├── has go.mod with module path ending /v3  → normal v3.2.1
        └── no go.mod, or path lacks /v3            → v3.2.1+incompatible
```

The suffix is not in the tag; the go command adds it to record "this predates
modules".

## 3. What It Implies

| | `+incompatible` |
|---|---|
| Module-aware dependencies | None — the repo has no `go.mod` |
| Transitive requirements | Not declared; must be discovered by import scanning |
| Coexistence with v1 | Impossible — same import path |

Because its own dependencies are undeclared, `go mod tidy` has to infer them,
which is why upgrading such a dependency often churns unrelated lines.

## 4. Gotchas

- Not a build error — it works. It is a **maintenance smell**: the upstream has
  been unmaintained since before Go 1.11.
- A repo that later adds a proper `/v3` path publishes an entirely separate
  module; migration is a manual import rewrite.
- `go list -m all | grep incompatible` is a fast audit of pre-modules baggage.
  See `Dependency Auditing`.

---

## 🔗 References
- ⬆️ Parent: `02. Version Resolution`
