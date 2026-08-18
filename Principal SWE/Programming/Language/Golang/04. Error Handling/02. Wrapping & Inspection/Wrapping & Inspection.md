---
title: Wrapping & Inspection
tags:
  - golang
  - error-handling
parent: "[[Error Handling]]"
---

# Wrapping & Inspection

Error chaining with %w, deep unwrapping, errors.Is, errors.As, and errors.Join.

```text
Wrapping & Inspection
│
├── [[Error Wrapping (%w)]]
├── [[errors.Unwrap]]
├── [[errors.Is and Custom Is()]]
├── [[errors.As and Custom As()]]
├── [[errors.Join (Multi-Error)]]
├── [[Sentinel Errors]]
└── [[Custom Error Types]]
```

---

## 🗂️ Topics

- [[Error Wrapping (%w)]] — Creating causal error chains with fmt.Errorf and %w verb.
- [[errors.Unwrap]] — Extracting the underlying wrapped error in an error chain.
- [[errors.Is and Custom Is()]] — Value equality matching across wrapped error trees.
- [[errors.As and Custom As()]] — Type-based error matching and extracting target custom errors.
- [[errors.Join (Multi-Error)]] — Combining multiple independent errors into a single aggregated error.
- [[Sentinel Errors]] — Exported package-level error constants (io.EOF, sql.ErrNoRows).
- [[Custom Error Types]] — Implementing custom error structs with structured metadata and context.

---

## 🔗 References
- ⬆️ Parent: [[Error Handling]]
- 🎓 Root: [[Principal SWE]]
