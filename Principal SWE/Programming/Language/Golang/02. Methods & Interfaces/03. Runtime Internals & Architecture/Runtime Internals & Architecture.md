- [[Devirtualization Compiler Pass]] — How compiler devirtualizes interface method calls into direct function calls when concrete type is known.

- [[Dynamic Interface Dispatch Overhead]] — Branch prediction misses, itab lookups, and register saving costs during interface calls.

---
title: Runtime Internals & Architecture
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Methods & Interfaces]]"
---

# Runtime Internals & Architecture

iface/eface runtime memory representation, itab dispatch tables, allocation profiles, and architectural patterns.

```text
Runtime Internals & Architecture
│
├── [[iface and eface Structs]]
├── [[itab Dynamic Dispatch Table]]
├── [[itab Global Cache & Hash Tables]]
├── [[Interface Allocation Cost]]
├── [[Direct Interface Values (Optimization)]]
├── [[Interface Best Practices]]
└── [[Interface Anti-Patterns]]
```

---

## 🗂️ Topics

- [[iface and eface Structs]] — Two-word runtime interface representation: tab/type pointer and data pointer.
- [[itab Dynamic Dispatch Table]] — Interface table layout, method offset resolution, and runtime dispatch performance.
- [[itab Global Cache & Hash Tables]] — How runtime dynamically computes and caches itab instances for type pairs.
- [[Interface Allocation Cost]] — When converting concrete values to interfaces triggers heap allocations.
- [[Direct Interface Values (Optimization)]] — Compiler optimization storing pointers and small words directly inside the data word.
- [[Interface Best Practices]] — Accept interfaces, return structs; keep interfaces small; define interfaces in consumer package.
- [[Interface Anti-Patterns]] — Interface pollution, premature abstraction, mocking what you do not own.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]
- 🎓 Root: [[Principal SWE]]
