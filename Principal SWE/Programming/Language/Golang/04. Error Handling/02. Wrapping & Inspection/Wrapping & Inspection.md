---
title: Wrapping & Inspection
tags:
  - golang
  - error-handling
  - principal-swe
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
├── [[Custom Error Types]]
└── [[Error Tree Traversal Algorithms]]
```

---

## 🗂️ Topics

- [[Error Wrapping (%w)]] — Creating causal error chains using fmt.Errorf with %w verb.
- [[errors.Unwrap]] — Extracting the immediate underlying error from an error chain.
- [[errors.Is and Custom Is()]] — Checking for target error identity across wrapped error trees.
- [[errors.As and Custom As()]] — Extracting custom typed error structs from wrapped error chains.
- [[errors.Join (Multi-Error)]] — Combining multiple concurrent or independent errors into a single error.
- [[Sentinel Errors]] — Exported package constants (io.EOF, sql.ErrNoRows) and comparison rules.
- [[Custom Error Types]] — Implementing structured error structs containing error codes and context.
- [[Error Tree Traversal Algorithms]] — Recursive unwrapping and multi-error branch traversal.

---

## 🔗 References
- ⬆️ Parent: [[Error Handling]]
- 🎓 Root: [[Principal SWE]]
