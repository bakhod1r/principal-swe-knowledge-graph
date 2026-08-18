---
title: Version Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - introspection
parent: "[[Inspection]]"
---

# `go version`

Reports the toolchain version, and — more usefully — reads build metadata back
out of compiled binaries.

## 1. Usage

```bash
go version                     # the toolchain
go version ./bin/api           # the toolchain that built this binary
go version -m ./bin/api        # full module + build settings
go version -m -v ./bin/api
```

## 2. Reading a Binary

```text
$ go version -m ./bin/api
./bin/api: go1.25.0
        path    github.com/me/api/cmd/api
        mod     github.com/me/api  (devel)
        dep     github.com/go-chi/chi/v5  v5.1.0  h1:...
        build   -buildmode=exe
        build   vcs.revision=9f2c1ab...
        build   vcs.time=2026-02-11T09:14:02Z
        build   vcs.modified=false
```

This is the fastest supply-chain audit available — it answers "what is actually
in production". Pairs with `Dependency Auditing` and `Vulnerabilities`.

## 3. Gotchas

- `-buildvcs=false` (or building outside a VCS checkout) removes the `vcs.*` lines.
- `-trimpath` does not remove module info; `-ldflags "-w -s"` does not either.
- The same data is available at runtime via `runtime/debug.ReadBuildInfo()`.
- `govulncheck` can consume a binary directly, using exactly this metadata.

---

## 🔗 References
- ⬆️ Parent: [[Inspection]]
