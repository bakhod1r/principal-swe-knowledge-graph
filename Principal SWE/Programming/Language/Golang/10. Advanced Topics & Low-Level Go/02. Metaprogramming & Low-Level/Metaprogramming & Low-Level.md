---
title: Metaprogramming & Low-Level
tags:
  - golang
  - advanced
  - principal-swe
parent: "[[Advanced Topics & Low-Level Go]]"
---

# Metaprogramming & Low-Level

Runtime reflection, pointer arithmetic with unsafe.Pointer and uintptr, go:linkname directives.

```text
Metaprogramming & Low-Level
│
├── [[reflect.Type and reflect.Value]]
├── [[unsafe.Pointer & uintptr Arithmetic]]
├── [[unsafe.Slice and unsafe.String]]
├── [[go:linkname Compiler Directive]]
└── [[Go Plugins (.so) Dynamic Loading]]
```

---

## 🗂️ Topics

- [[reflect.Type and reflect.Value]] — Deep type introspection, struct field inspection, method invocation, performance overhead.
- [[unsafe.Pointer & uintptr Arithmetic]] — Direct memory manipulation, computing struct field offsets, casting pointers.
- [[unsafe.Slice and unsafe.String]] — Zero-copy conversion between byte slices and strings without heap allocation.
- [[go:linkname Compiler Directive]] — Linking to unexported runtime and standard library functions across package boundaries.
- [[Go Plugins (.so) Dynamic Loading]] — Compiling and dynamically loading shared object plugins at runtime via plugin package.

---

## 🔗 References
- ⬆️ Parent: [[Advanced Topics & Low-Level Go]]
- 🎓 Root: [[Principal SWE]]
