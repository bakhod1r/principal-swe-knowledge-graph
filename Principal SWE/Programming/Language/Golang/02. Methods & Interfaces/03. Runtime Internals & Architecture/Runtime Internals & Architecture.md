---
title: Runtime Internals & Architecture
tags:
  - golang
  - methods-and-interfaces
parent: "[[Methods & Interfaces]]"
---

# Runtime Internals & Architecture

Memory representation of interfaces (iface, eface), dynamic dispatch tables (itab), and design patterns.

```text
Runtime Internals & Architecture
│
├── [[iface and eface Structs]]
├── [[itab and Dynamic Dispatch]]
├── [[Interface Allocation Cost]]
├── [[Interface Best Practices]]
└── [[Interface Anti-Patterns]]
```

---

## 🗂️ Topics

- [[iface and eface Structs]] — Two-word structure: _type/itab pointer + data pointer.
- [[itab and Dynamic Dispatch]] — Virtual method table construction, caching, and dispatch cost.
- [[Interface Allocation Cost]] — When assigning a concrete value to an interface causes a heap allocation.
- [[Interface Best Practices]] — Accept interfaces, return structs; small interfaces; consumer-defined interfaces.
- [[Interface Anti-Patterns]] — Premature abstraction, interface pollution, returning interfaces.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]
- 🎓 Root: [[Principal SWE]]
