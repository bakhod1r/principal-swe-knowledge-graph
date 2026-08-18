---
title: Go Source Code Folder Structure
tags:
  - golang
  - basics
  - internals
  - architecture
parent: "[[Language Basic]]"
---

# 📂 Go Source Code Folder Structure

The **Go Source Code Structure** represents the complete layout inside `GOROOT` (`/usr/local/go` or your active Go installation), containing the Go standard library, runtime engine, compiler, and developer tools.

```text
Go Source Code Structure (GOROOT)
│
├── [[Standard Library]] (src/)
│   ├── fmt, net/http, os, io
│   └── sync, context, crypto, encoding, reflect
│
├── [[Runtime Internals]] (src/runtime/)
│   ├── proc.go (GMP Scheduler) ⭐⭐⭐
│   ├── mgc.go (Garbage Collector) ⭐⭐⭐
│   ├── malloc.go (Memory Allocator) ⭐⭐⭐
│   └── chan.go, map.go, iface.go, panic.go
│
├── [[Toolchain & Compiler]] (src/cmd/)
│   ├── cmd/go (CLI Coordinator) ⭐⭐⭐
│   ├── cmd/compile (Compiler Frontend & Backend) ⭐⭐⭐
│   └── cmd/link, cmd/asm, cmd/vet
│
├── [[Internal Packages]] (src/internal/)
│   └── Internal packages and compiler visibility rules
│
└── [[Installation Directories]]
    └── bin/, pkg/, doc/, test/, misc/
```

---

## 🗂️ Categories

1. 📦 **[[Standard Library]]** (`src/`) — Built-in packages (`fmt`, `net/http`, `os`, `io`, `sync`, `context`, etc.).
2. ⚡ **[[Runtime Internals]]** (`src/runtime/`) — Scheduler (`proc.go`), Garbage Collector (`mgc.go`), Memory Allocator (`malloc.go`), Channels, and Maps.
3. 🛠️ **[[Toolchain & Compiler]]** (`src/cmd/`) — Go CLI orchestration (`cmd/go`), compiler (`cmd/compile`), linker (`cmd/link`), and assembler (`cmd/asm`).
4. 🔒 **[[Internal Packages]]** (`src/internal/`) — Private standard library helpers isolated by compiler enforcement.
5. 📁 **[[Installation Directories]]** (`GOROOT/`) — Binaries (`bin/`), toolchain artifacts (`pkg/`), documentation (`doc/`), tests (`test/`), and utilities (`misc/`).

---

## 🧠 Mental Model: User Project vs Go Source

```text
YOUR APPLICATION PROJECT
───────────────────────────────────
~/projects/order-service/
├── go.mod
├── go.sum
├── cmd/main.go
└── internal/order/

GO TOOLCHAIN & STANDARD LIBRARY (GOROOT)
───────────────────────────────────
/usr/local/go/
├── bin/ (go, gofmt)
└── src/ (fmt, net/http, runtime, cmd/compile)
```

---

## 🔗 Navigation
- ⬆️ Parent: [[Language Basic]]
- 📂 Setup Guide: `Settings Environment`
- 📦 Dependencies: `Dependencies`
