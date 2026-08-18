---
title: Binary Inspection Tools (go tool)
tags:
  - golang
  - toolchain
  - principal-swe
parent: "[[Go Commands]]"
---

# Binary Inspection Tools (go tool)

Internal toolchain inspection utilities: compile -S, objdump, nm, and cgo.

```text
Binary Inspection Tools (go tool)
│
├── [[go tool compile -S (Plan 9 Assembly)]]
├── [[go tool objdump (Disassembly)]]
├── [[go tool nm (Symbol Table)]]
└── [[go tool cgo (Bridge Stubs)]]
```

---

## 🗂️ Topics

- [[go tool compile -S (Plan 9 Assembly)]] — Inspecting generated Plan 9 assembly instructions directly from source compilation.
- [[go tool objdump (Disassembly)]] — Disassembling compiled binaries into raw CPU machine instructions.
- [[go tool nm (Symbol Table)]] — Extracting and inspecting binary symbol tables, functions, and global variable offsets.
- [[go tool cgo (Bridge Stubs)]] — Inspecting auto-generated C bridge stubs and memory wrapper files.

---

## 🔗 References
- ⬆️ Parent: [[Go Commands]]
- 🎓 Root: [[Principal SWE]]
