---
title: VERSION
tags:
  - golang
  - goroot
  - release
parent: "[[Installation Directories]]"
---

# `$GOROOT/VERSION`

A two-line file naming the exact toolchain release.

## 1. Contents

```bash
cat "$(go env GOROOT)/VERSION"
```

```text
go1.26.2
time 2026-02-05T14:22:11Z
```

## 2. Who Reads It

The `go` command reads this at startup to report `GOVERSION` — which is why
`go version` is instant and needs no subprocess.

```bash
go version                 # go version go1.26.2 darwin/arm64
go env GOVERSION           # go1.26.2
```

## 3. Detecting a Development Toolchain

A toolchain built from source with no release tag reports `devel` plus a commit:

```text
devel go1.27-9f2c1ab3d4 Wed Feb 18 09:14:02 2026 +0000
```

Seeing `devel` in a CI log means someone is building with an unreleased Go.

## 4. Gotchas

- Deleting or corrupting it breaks the toolchain's self-identification and
  therefore `GOTOOLCHAIN` switching.
- The version recorded **in a binary** comes from the toolchain that compiled it,
  readable with `go version` `-m` — not from this file at run time.

---

## 🔗 References
- ⬆️ Parent: [[Installation Directories]]
