---
title: Major Version Suffix
tags:
  - golang
  - basics
  - dependencies
  - modules
  - versioning
parent: "[[Version Resolution]]"
---

# Major Version Suffix

From `v2` onward, the major version is part of the module path and therefore part
of every import line.

## 1. The Rule

```text
v0, v1  →  github.com/me/lib
v2      →  github.com/me/lib/v2
v3      →  github.com/me/lib/v3
```

Declared in the `module` directive and repeated in every consumer's imports.

## 2. Why

Two major versions must be able to coexist in one build:

```go
import (
    old "github.com/me/lib"       // v1, pulled in by a dependency
    new "github.com/me/lib/v2"    // v2, used by your code
)
```

Different paths → different packages → no conflict in `MVS`. Without the
suffix, MVS would have to pick one and break somebody.

## 3. Releasing v2

Two supported layouts:

```text
A) major branch      B) major subdirectory
   repo root            repo root
   └── go.mod              ├── go.mod          (v1)
       module .../v2       └── v2/
   on branch v2                └── go.mod      (module .../v2)
```

Both require editing `module` **and** every internal import inside the module.

## 4. Gotchas

- Forgetting to update internal imports gives the notorious
  `module .../v2 found, but does not contain package .../v2/x` error.
- Tag and path must agree: tag `v2.0.0` with a path lacking `/v2` is only usable
  as `+incompatible` — see `Incompatible Versions`.
- `go get lib@v2.0.0` on the old path fails; the path itself changes.

---

## 🔗 References
- ⬆️ Parent: `02. Version Resolution`
