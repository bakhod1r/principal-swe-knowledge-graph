---
title: exclude directive
tags:
  - golang
  - basics
  - dependencies
  - modules
  - go.mod
parent: "[[Core Concepts]]"
---

# `exclude`

Removes a specific module version from consideration during `MVS`. The graph
behaves as though that version was never published.

## 1. Syntax

```go
exclude (
    github.com/broken/lib v1.4.2
    github.com/broken/lib v1.4.3
)
```

## 2. Effect on Resolution

```text
MVS picks the highest required version
        │
        ├── v1.4.2 excluded → skip
        ├── v1.4.3 excluded → skip
        └── v1.4.4 → selected
```

If nothing higher exists, resolution **fails** rather than falling back to a lower
version.

## 3. `exclude` vs `replace` vs `retract`

| Directive | Who writes it | Effect |
|---|---|---|
| `exclude` | Consumer | Skip a version, pick the next higher one |
| `replace` | Consumer | Substitute a different source or version |
| `retract` | **Author** | Warn every consumer that a release is bad |

## 4. Gotchas

- `exclude` in a **dependency's** `go.mod` is ignored — only the main module's
  exclusions apply. Same rule as `replace`.
- It is version-specific, not a range: excluding `v1.4.2` does nothing about
  `v1.4.5` having the same bug.
- Prefer requiring a fixed version explicitly; `exclude` is for when the broken
  version is pulled in transitively and you cannot bump the intermediary.

---

## 🔗 References
- ⬆️ Parent: `01. Core Concepts`
