---
title: Binary Inspection Tools (go tool)
tags:
  - golang
  - toolchain
  - principal-swe
parent: "[[Go Commands]]"
---

# Binary Inspection Tools (go tool)

Internal toolchain utilities: compile -S, objdump, nm, cgo, pprof, trace, asm, link, buildid, and addr2line.

```text
Binary Inspection Tools (go tool)
│
├── [[go tool compile -S (Plan 9 Assembly)]]
├── [[go tool objdump (Disassembly)]]
├── [[go tool nm (Symbol Table)]]
├── [[go tool pprof (CPU & Heap Profiles)]]
├── [[go tool trace (Execution Tracer)]]
├── [[go tool asm (Plan 9 Assembler)]]
├── [[go tool link (Linker)]]
├── [[go tool buildid (Build Identity)]]
├── [[go tool addr2line (Address Mapping)]]
└── [[go tool cgo (Bridge Stubs)]]
```

---

## 🗂️ Topics

- [[go tool compile -S (Plan 9 Assembly)]] — Inspecting generated Plan 9 assembly instructions directly from source compilation.
- [[go tool objdump (Disassembly)]] — Disassembling compiled binaries into raw CPU machine instructions.
- [[go tool nm (Symbol Table)]] — Extracting and inspecting binary symbol tables, functions, and global variable offsets.
- [[go tool cgo (Bridge Stubs)]] — Inspecting auto-generated C bridge stubs and memory wrapper files.
- [[go tool pprof (CPU & Heap Profiles)]] — Interactive CPU, heap, mutex, and block profile analysis.
- [[go tool trace (Execution Tracer)]] — Goroutine scheduling, GC, and syscall timeline visualisation.
- [[go tool asm (Plan 9 Assembler)]] — Assembling `.s` Plan 9 assembly sources into object files.
- [[go tool link (Linker)]] — Object linking, `-ldflags` variable injection, and symbol stripping.
- [[go tool buildid (Build Identity)]] — Reading and rewriting the build ID that keys the build cache.
- [[go tool addr2line (Address Mapping)]] — Mapping runtime addresses back to file and line numbers.

---

## 🔗 References
- ⬆️ Parent: [[Go Commands]]

