---
title: go.env
tags:
  - golang
  - goroot
  - configuration
parent: "[[Installation Directories]]"
---

# `$GOROOT/go.env`

Toolchain-shipped defaults for `go env` values. The lowest layer of the
configuration stack.

## 1. Contents

```bash
cat "$(go env GOROOT)/go.env"
```

```text
# This file contains the initial defaults for go command configuration.
# Values set by 'go env -w' and written to the user's go/env file override these.
GOPROXY=https://proxy.golang.org,direct
GOSUMDB=sum.golang.org
GOTOOLCHAIN=auto
```

## 2. Precedence

```text
OS environment variable
        ↓ overrides
user go/env file        ($(go env GOENV), written by `go env -w`)
        ↓ overrides
$GOROOT/go.env          ← this file
        ↓ overrides
compiled-in defaults
```

## 3. Why It Matters

It is where `GOPROXY`, `GOSUMDB`, and `GOTOOLCHAIN` get their public
defaults. An organization distributing a custom toolchain edits this file to
point every developer at an internal proxy with no per-machine setup.

## 4. Gotchas

- Edits are lost on every Go upgrade — for a persistent per-user change use
  `go env -w`, which writes to `GOENV` instead.
- `go env -u <VAR>` reverts to the value **here**, not to empty.
- Being read from `GOROOT` means a per-user Go install and a system-wide one can
  legitimately disagree.

---

## 🔗 References
- ⬆️ Parent: [[Installation Directories]]
