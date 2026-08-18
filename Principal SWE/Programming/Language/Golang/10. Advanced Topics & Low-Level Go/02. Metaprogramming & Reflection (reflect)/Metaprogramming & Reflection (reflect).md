---
title: Metaprogramming & Reflection (reflect)
tags:
  - golang
  - advanced
  - principal-swe
parent: "[[Advanced Topics & Low-Level Go]]"
---

# Metaprogramming & Reflection (reflect)

The 3 laws of reflection, struct field inspection, settability, dynamic method invocation, and dynamic shared plugins (.so).

```text
Metaprogramming & Reflection (reflect)
│
├── [[Laws of Reflection (Interface to reflect.Value, reflect.Type)]]
├── [[Deep Struct Field Inspection & Dynamic Method Invocation]]
├── [[Settability & Addressability in Reflection]]
├── [[Reflect-Based Code Generation vs Static Generation]]
├── [[Type Creation at Runtime (reflect.StructOf & reflect.ArrayOf)]]
└── [[Go Plugins (.so) Dynamic Loading Architecture]]
```

---

## 🗂️ Topics

- [[Laws of Reflection (Interface to reflect.Value, reflect.Type)]] — Dissecting the 3 fundamental Laws of Reflection in Go and type representation.
- [[Deep Struct Field Inspection & Dynamic Method Invocation]] — Extracting struct tags, reading private fields via unsafe, and invoking methods dynamically.
- [[Settability & Addressability in Reflection]] — v.CanSet(), v.Elem(), and taking pointer addresses to mutate structs through reflection.
- [[Reflect-Based Code Generation vs Static Generation]] — Benchmarking runtime reflection performance penalties against build-time code generation.
- [[Type Creation at Runtime (reflect.StructOf & reflect.ArrayOf)]] — Programmatically constructing new dynamic struct and array types at runtime.
- [[Go Plugins (.so) Dynamic Loading Architecture]] — Building, compiling, and loading dynamic shared libraries (plugin.Open, plugin.Lookup) at runtime.

---

## 🔗 References
- ⬆️ Parent: [[Advanced Topics & Low-Level Go]]

