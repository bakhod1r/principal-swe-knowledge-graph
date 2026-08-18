---
title: Modern Error Wrapping & Inspection (errors package)
tags:
  - golang
  - error-handling
  - principal-swe
parent: "[[Error Handling]]"
---

# Modern Error Wrapping & Inspection (errors package)

fmt.Errorf %w wrapping, errors.Is semantic checking, errors.As type extraction, multi-errors (errors.Join), and tree traversal.

```text
Modern Error Wrapping & Inspection (errors package)
│
├── [[Error Wrapping Protocol with fmt.Errorf and %w]]
├── [[errors.Is Semantic Inspection (Unwrap Single)]]
├── [[errors.As Type Extraction & Pattern Matching]]
├── [[Multi-Error Aggregation with errors.Join (Go 1.20+)]]
└── [[Error Tree Traversal Mechanics (Breadth-First vs Depth-First)]]
```

---

## 🗂️ Topics

- [[Error Wrapping Protocol with fmt.Errorf and %w]] — Constructing causal error chains using %w vs formatting without wrapping (%v).
- [[errors.Is Semantic Inspection (Unwrap Single)]] — Recursive error tree traversal, custom Is(target error) bool method implementation.
- [[errors.As Type Extraction & Pattern Matching]] — Extracting concrete error types across nested wrapping layers with As(target any) bool.
- [[Multi-Error Aggregation with errors.Join (Go 1.20+)]] — Combining concurrent/batch errors into a single error, multi-branch Unwrap() []error.
- [[Error Tree Traversal Mechanics (Breadth-First vs Depth-First)]] — How errors.Is and errors.As navigate tree structures created by errors.Join.

---

## 🔗 References
- ⬆️ Parent: [[Error Handling]]

