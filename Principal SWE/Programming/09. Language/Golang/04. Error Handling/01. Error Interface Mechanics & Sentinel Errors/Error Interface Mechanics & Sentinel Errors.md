---
title: Error Interface Mechanics & Sentinel Errors
tags:
  - golang
  - error-handling
  - principal-swe
parent: "[[Error Handling (Clean Code)]]"
---

# Error Interface Mechanics & Sentinel Errors

error interface contract, sentinel error variables, constant errors, contextual error structs, and the typed nil trap.

```text
Error Interface Mechanics & Sentinel Errors
│
├── [[The error Interface Contract (Error() string)]]
├── [[Sentinel Errors Design (io.EOF, sql.ErrNoRows)]]
├── [[Constant Errors Pattern via Defined String Types]]
├── [[Custom Error Structs with Contextual Fields]]
└── [[The Typed Nil Error Return Trap]]
```

---

## 🗂️ Topics

- [[The error Interface Contract (Error() string)]] — Single-method interface simplicity, nil error representations, and zero-allocation static errors.
- [[Sentinel Errors Design (io.EOF, sql.ErrNoRows)]] — Predeclared exported error variables, comparing with errors.Is, and immutability pitfalls.
- [[Constant Errors Pattern via Defined String Types]] — Defining immutable package-level sentinel errors (type Error string) preventing global variable mutation.
- [[Custom Error Structs with Contextual Fields]] — Creating structured errors carrying HTTP status codes, request IDs, retry after durations, and timestamps.
- [[The Typed Nil Error Return Trap]] — Returning a concrete typed nil pointer as an error interface resulting in iface.tab != nil (err != nil).

---

## 🔗 References
- ⬆️ Parent: [[Error Handling (Clean Code)]]

