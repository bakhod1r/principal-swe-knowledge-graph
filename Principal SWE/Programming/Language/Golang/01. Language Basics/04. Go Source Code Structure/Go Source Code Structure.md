---
title: Go Source Code Structure
tags:
  - golang
  - basics
parent: "[[Language Basics]]"
---

# Go Source Code Structure

Internal organization of $GOROOT: standard library (src/), runtime, compiler, and internal packages.

```text
Go Source Code Structure
│
├── [[src Directory]]
├── [[src-runtime]]
├── [[src-cmd]]
├── [[src-internal]]
└── [[GOROOT bin and pkg]]
```

---

## 🗂️ Topics

- [[src Directory]] — Root source tree for all standard library packages and runtime.
- [[src-runtime]] — Core runtime engine (proc.go, mgc.go, malloc.go, chan.go).
- [[src-cmd]] — Toolchain source (cmd/go, cmd/compile, cmd/link, cmd/asm).
- [[src-internal]] — Private unimportable standard library helper packages.
- [[GOROOT bin and pkg]] — Executable toolchain binaries and precompiled metadata archives.

---

## 🔗 References
- ⬆️ Parent: [[Language Basics]]
- 🎓 Root: [[Principal SWE]]
