---
title: module directive
tags:
  - golang
  - basics
  - dependencies
  - modules
  - go.mod
parent: "[[Core Concepts]]"
---

# `module`

The first directive in `go.mod`. Declares the module's **import path** — the
prefix every package inside it is imported by.

## 1. Syntax

```go
module github.com/me/api
```

## 2. The Path Is Not a Directory

```text
module github.com/me/api
                      │
        directory api/internal/user
                      ▼
        import "github.com/me/api/internal/user"
```

The path must match where the module is fetched from, because `GOPROXY` derives
the download URL from it. A mismatch produces:

```text
module declares its path as: github.com/me/api
        but was required as: github.com/you/api
```

## 3. Major Version Suffix

From `v2` onward the suffix is **part of the path**:

```go
module github.com/me/api/v2
```

See `Major Version Suffix`. This is why a major bump breaks every import line.

## 4. Gotchas

- Renaming a module is a breaking change for every consumer; there is no alias
  mechanism. `retract` plus a new path is the migration route.
- For a private repository the path must still be the real host
  (`git.corp.internal/team/api`) so `GOPRIVATE` patterns match.
- A module path with an uppercase letter is legal but the proxy stores it
  case-encoded (`!m` for `M`) — a frequent source of confusing cache paths.

---

## 🔗 References
- ⬆️ Parent: `01. Core Concepts`
