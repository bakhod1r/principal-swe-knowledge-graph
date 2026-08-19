---
title: Runtime Bootstrapping & Core Architecture
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# Runtime Bootstrapping & Core Architecture

Hardware entry points, runtime initialization (schedinit), main goroutine lifecycle, g/m/p structs, and symbol table layout (symtab.go).

```text
Runtime Bootstrapping & Core Architecture
│
├── [[Runtime Boot Sequence (rt0_go & asm_amd64.s)]]
├── [[schedinit Implementation & Runtime Initialization]]
├── [[The main Goroutine Lifecycle (runtime.main)]]
├── `Goroutine Layout & g Struct Internals`
├── `OS Thread Layout & m Struct Internals`
├── `Logical Processor Layout & p Struct Internals`
├── [[The g0 System Stack & Stack Switching]]
└── `moduledata & Global Symbol Table Layout (symtab.go)`
```

---

## 🗂️ Topics

- [[Runtime Boot Sequence (rt0_go & asm_amd64.s)]] — Hardware entry point, CPU feature detection, argc/argv extraction, and initial OS thread creation.
- [[schedinit Implementation & Runtime Initialization]] — Stack initialization, memory allocator setup (mallocinit), mcommoninit, gcinit, and procresize.
- [[The main Goroutine Lifecycle (runtime.main)]] — Spawning runtime.main, executing package init() dependency graphs, initializing sysmon, calling main.main.
- `Goroutine Layout & g Struct Internals` — Dissecting the 80+ fields of g struct: stack bounds, sched context, m pointer, atomic status, and panic list.
- `OS Thread Layout & m Struct Internals` — Dissecting m struct: g0 system stack, gsignal, curg running goroutine, p pointer, and fastrand state.
- `Logical Processor Layout & p Struct Internals` — Dissecting p struct: lock-free local runqueues (runq), mcache span caches, sudogcache, and timers.
- [[The g0 System Stack & Stack Switching]] — Dedicated OS-sized system stack (8MB) used for scheduler execution, runtime memory allocation, and GC.
- `moduledata & Global Symbol Table Layout (symtab.go)` — First-class binary metadata struct: pclntab line mapping, function descriptors (funcInfo), and type descriptors.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]

