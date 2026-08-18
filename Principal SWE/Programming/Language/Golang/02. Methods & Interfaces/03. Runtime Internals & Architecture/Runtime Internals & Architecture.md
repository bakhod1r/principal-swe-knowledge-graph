---
title: Runtime Internals & Architecture
tags:
  - golang
  - interfaces
parent: "[[Methods & Interfaces]]"
---

# Runtime Internals & Architecture

Runtime representation (iface/eface), itab dispatch table, allocation profile, and architectural principles.

```text
Runtime Internals & Architecture
│
├── [[iface and eface Structs]]
├── [[itab Dynamic Dispatch Table]]
├── [[Interface Allocation Cost]]
├── [[Interface Best Practices]]
└── [[Interface Anti-Patterns]]
```

---

## 🗂️ Topics

- [[iface and eface Structs]] — Two-word runtime interface representation: tab/type pointer and data pointer.
- [[itab Dynamic Dispatch Table]] — Interface table layout, method offset resolution, and runtime dispatch performance.
- [[Interface Allocation Cost]] — When converting concrete values to interfaces triggers heap allocations.
- [[Interface Best Practices]] — Accept interfaces, return structs; keep interfaces small; define interfaces at point of consumption.
- [[Interface Anti-Patterns]] — Interface pollution, premature abstraction, mocking what you do not own.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]
- 🎓 Root: [[Principal SWE]]
