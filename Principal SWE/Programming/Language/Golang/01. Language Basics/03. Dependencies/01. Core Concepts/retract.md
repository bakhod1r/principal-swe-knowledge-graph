---
title: retract directive
tags:
  - golang
  - basics
  - dependencies
  - modules
  - go.mod
  - publishing
parent: "[[Core Concepts]]"
---

# `retract`

Lets a module **author** mark their own published versions as withdrawn. Go 1.16+.

## 1. Syntax

```go
retract (
    v1.0.3                  // contains a data-loss bug
    [v1.1.0, v1.1.4]        // accidentally published from a fork
)
```

The trailing comment is shown to users — write it as an explanation.

## 2. What Users See

```bash
$ go get github.com/me/lib@latest
go: warning: github.com/me/lib@v1.0.3: retracted by module author:
        contains a data-loss bug
```

Retracted versions are skipped by `@latest`, hidden from `go list -m -versions`,
and warned about if already required.

## 3. Why Not Just Delete the Tag

Because module content is **immutable**. Deleting a tag does not remove it from
`GOPROXY` or `GOSUMDB` — the transparency log guarantees it stays fetchable.
`retract` is the only supported withdrawal mechanism. See
`Distribution & Integrity`.

## 4. The Self-Reference Trick

To retract a version you must publish a **newer** version containing the
`retract` line. To retract the latest release you therefore publish `v1.0.4`
whose only change is retracting `v1.0.3` — and it may retract itself:

```go
retract (
    v1.0.3   // bad release
    v1.0.4   // this release exists only to retract v1.0.3
)
```

## 5. Gotchas

- Retraction is advisory: builds that already pin the version keep working.
- Only the retracting module's own versions can be listed.
- Users on older Go versions ignore it entirely.

---

## 🔗 References
- ⬆️ Parent: `01. Core Concepts`
