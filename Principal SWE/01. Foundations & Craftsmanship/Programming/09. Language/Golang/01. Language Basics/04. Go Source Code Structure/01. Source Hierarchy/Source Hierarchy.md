---
title: Source Hierarchy
tags:
  - golang
  - source-structure
  - principal-swe
parent: "[[Go Source Code Structure]]"
---

# Source Hierarchy

Standard library source root ($GOROOT/src), runtime engine, compiler, and internal packages.

```text
Source Hierarchy
│
├── [[src Directory Layout]]
├── [[src-runtime Architecture]]
├── [[src-cmd Toolchain Source]]
├── [[src-internal Private Packages]]
└── [[GOROOT bin and pkg Layout]]
```

---

## 🗂️ Topics

- [[src Directory Layout]] — Root source tree for all standard library packages and runtime engine.
- [[src-runtime Architecture]] — Core runtime engine sources: proc.go, mgc.go, malloc.go, chan.go, panic.go.
- [[src-cmd Toolchain Source]] — Compiler and linker source code: cmd/compile, cmd/link, cmd/asm, cmd/go.
- [[src-internal Private Packages]] — Standard library private unimportable helper packages (internal/cpu, internal/bytealg).
- [[GOROOT bin and pkg Layout]] — Executable compiler binaries and precompiled package metadata archives.

---

## 🔗 References
- ⬆️ Parent: [[Go Source Code Structure]]

