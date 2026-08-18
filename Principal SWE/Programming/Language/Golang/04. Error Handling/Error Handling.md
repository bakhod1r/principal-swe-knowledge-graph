---
title: Error Handling
tags:
  - golang
  - error-handling
  - principal-swe
parent: "[[Golang]]"
---

# 🚨 Error Handling

Go error handling primitives, error wrapping, error inspection, custom error hierarchies, panic, and recover.

```text
Error Handling
│
├── [[Core Primitives|01. Core Primitives]]
│   ├── [[Errors as First-Class Values]]
│   ├── [[error Interface Contract]]
│   ├── [[errors.New]]
│   ├── [[fmt.Errorf Formatting]]
│   ├── [[Error String Conventions]]
│   └── [[nil Error Pitfall]]
├── [[Wrapping & Inspection|02. Wrapping & Inspection]]
│   ├── [[Error Wrapping (%w)]]
│   ├── [[errors.Unwrap]]
│   ├── [[errors.Is and Custom Is()]]
│   ├── [[errors.As and Custom As()]]
│   ├── [[errors.Join (Multi-Error)]]
│   ├── [[Sentinel Errors]]
│   └── [[Custom Error Types]]
└── [[Panic, Recover & Architecture|03. Panic, Recover & Architecture]]
│   ├── [[panic Semantics]]
│   ├── [[recover in Deferred Functions]]
│   ├── [[Goroutine Panic Isolation]]
│   ├── [[Stack Traces & runtime-debug]]
│   ├── [[Domain vs Infrastructure Errors]]
│   ├── [[Error Design Best Practices]]
│   └── [[Handle Errors, Don't Just Check]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Core Primitives|01. Core Primitives]]
- [[Errors as First-Class Values]] — Errors as explicit return values, checking if err != nil.
- [[error Interface Contract]] — The built-in single-method interface contract (Error() string).
- [[errors.New]] — Creating simple static error values.
- [[fmt.Errorf Formatting]] — Formatting dynamic error messages with %v, %s, and %d.
- [[Error String Conventions]] — Lowercase, no trailing punctuation, descriptive message rules.
- [[nil Error Pitfall]] — Typed nil pointer assigned to error interface causing non-nil interface bug.
### 2. 📂 [[Wrapping & Inspection|02. Wrapping & Inspection]]
- [[Error Wrapping (%w)]] — Creating causal error chains with fmt.Errorf and %w verb.
- [[errors.Unwrap]] — Extracting the underlying wrapped error in an error chain.
- [[errors.Is and Custom Is()]] — Value equality matching across wrapped error trees.
- [[errors.As and Custom As()]] — Type-based error matching and extracting target custom errors.
- [[errors.Join (Multi-Error)]] — Combining multiple independent errors into a single aggregated error.
- [[Sentinel Errors]] — Exported package-level error constants (io.EOF, sql.ErrNoRows).
- [[Custom Error Types]] — Implementing custom error structs with structured metadata and context.
### 3. 📂 [[Panic, Recover & Architecture|03. Panic, Recover & Architecture]]
- [[panic Semantics]] — Unwinding the goroutine call stack on unrecoverable conditions.
- [[recover in Deferred Functions]] — Catching runtime panics and restoring program execution safely.
- [[Goroutine Panic Isolation]] — Panics inside goroutines crash the entire process unless caught locally.
- [[Stack Traces & runtime-debug]] — Capturing and formatting stack traces for observability.
- [[Domain vs Infrastructure Errors]] — Separating business logic errors from database/network failures.
- [[Error Design Best Practices]] — Decorating errors without losing context, avoid string matching errors.
- [[Handle Errors, Don't Just Check]] — Meaningful error mitigation vs blindly propagating nil errors.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

