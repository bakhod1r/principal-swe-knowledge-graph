---
title: go mod verify
tags:
  - golang
  - basics
  - cli
  - toolchain
  - modules
  - security
parent: "[[Module Commands]]"
---

# `go mod verify`

Checks that the modules in `GOMODCACHE` still hash to the values recorded when
they were downloaded.

```bash
go mod verify
```

```text
all modules verified
```

## 1. What It Proves and Does Not Prove

| | |
|---|---|
| Proves | Nothing on **this machine** modified the cached module contents |
| Does not prove | That the content matches upstream — that is `GOSUMDB`'s job at download time |

The check is against the `.ziphash` files written at download, so it detects
local tampering and disk corruption, not a compromised upstream.

## 2. Where It Belongs

```bash
go mod download
go mod verify        # cache integrity
go build ./...
```

In a build pipeline that shares a warm module cache between jobs, this is the
step that catches a poisoned cache volume.

## 3. Gotchas

- A failure names the module and its directory; the fix is
  `go clean -modcache` followed by a fresh `go mod download`.
- Modules extracted read-only will fail verification if a tool "helpfully"
  reformatted them — a classic symptom of running `gofmt` across `GOPATH`.
- Verification of *upstream* authenticity lives in `go.sum` and
  `Distribution & Integrity`.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
