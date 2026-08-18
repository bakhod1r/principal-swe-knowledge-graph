---
title: Module Deprecation
tags:
  - golang
  - basics
  - dependencies
  - modules
  - publishing
parent: "[[Distribution & Integrity]]"
---

# Module Deprecation

Marks an entire module as no longer maintained, so consumers are told on upgrade.
Go 1.17+.

## 1. Syntax

A comment block immediately before the `module` directive:

```go
// Deprecated: use github.com/me/lib/v2 instead.
module github.com/me/lib
```

The paragraph must start with `Deprecated:` — the rest is free text shown to users.

## 2. What Users See

```bash
$ go get -u ./...
go: module github.com/me/lib is deprecated: use github.com/me/lib/v2 instead.
```

Also surfaced by:

```bash
go list -m -u -f '{{.Path}} {{.Deprecated}}' all
```

## 3. Deprecation vs Retraction

| | Deprecation | `retract` |
|---|---|---|
| Scope | The whole module | Specific versions |
| Meaning | "Stop using this" | "This release is broken" |
| Build effect | None — advisory only | Version skipped by `@latest` |

Both require publishing a **new version** carrying the notice.

## 4. Gotchas

- Purely advisory; nothing stops working. Consumers who never run `go get -u`
  never see it.
- The notice is read from the **latest** version of the module, so deprecating
  requires one final release.
- Deprecating `v1` when publishing `v2` is the intended pattern — see
  `Major Version Suffix`.

---

## 🔗 References
- ⬆️ Parent: `03. Distribution & Integrity`
