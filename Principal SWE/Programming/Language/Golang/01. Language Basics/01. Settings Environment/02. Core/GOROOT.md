---
title: GOROOT
tags:
  - golang
  - basics
  - environment
parent: "[[Core]]"
---
# `GOROOT`

**`GOROOT` = the root directory where the Go SDK / toolchain is installed.**

Check it:

```bash
go env GOROOT
```

Typical paths:
- macOS (Homebrew): `/opt/homebrew/opt/go/libexec` or `/usr/local/go`
- Linux: `/usr/local/go`
- Windows: `C:\Program Files\Go`

---

## 1. Directory Structure of `GOROOT`

```text
GOROOT/
├── bin/
│   ├── go        # The main go CLI
│   └── gofmt     # Code formatter
├── src/          # Go standard library source code (fmt, net/http, etc.)
├── pkg/          # Compiled standard library package objects
└── doc/          # Go documentation
```

---

## 2. When Should You Set `GOROOT`?

Normally, **DO NOT manually set `GOROOT`**.
Modern Go automatically determines `GOROOT` relative to the path of the `go` executable.

Only set `GOROOT` if:
- You maintain multiple custom Go compiler installations.
- Your IDE specifically requires setting the SDK path manually.

---

## 3. Mental Model

```text
             GOROOT
               │
      ┌────────┴────────┐
      ▼                 ▼
 $GOROOT/bin       $GOROOT/src
  (go, gofmt)    (Standard Library)
```

---

## 🔗 Navigation & Deep Dive
- 📂 Deep Dive: `Go Source Code Folder Structure (src, runtime, cmd)`
- ⬆️ Parent: `Settings Environment`
- 📂 Compare: `GOPATH` | `GOBIN`
