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
│   ├── [[Custom Error Types]]
│   └── [[Error Tree Traversal Algorithms]]
└── [[Panic, Recover & Architecture|03. Panic, Recover & Architecture]]
│   ├── [[panic Semantics]]
│   ├── [[recover in Deferred Functions]]
│   ├── [[Goroutine Panic Isolation]]
│   ├── [[Stack Traces & runtime-debug]]
│   ├── [[Domain vs Infrastructure Errors]]
│   ├── [[Error Classification (Transient vs Permanent)]]
│   ├── [[Error Design Best Practices]]
│   └── [[Handle Errors, Don't Just Check]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Core Primitives|01. Core Primitives]]
- [[Errors as First-Class Values]] — Explicit error values, if err != nil idiom, and error control flow.
- [[error Interface Contract]] — The single-method built-in interface contract: Error() string.
- [[errors.New]] — Creating simple static error values using errors.New().
- [[fmt.Errorf Formatting]] — Formatting dynamic error messages with %s, %d, and %v.
- [[Error String Conventions]] — Lowercase, no trailing punctuation, descriptive error phrasing standards.
- [[nil Error Pitfall]] — Typed nil pointer assigned to error interface creating non-nil interface trap.
### 2. 📂 [[Wrapping & Inspection|02. Wrapping & Inspection]]
- [[Error Wrapping (%w)]] — Creating causal error chains using fmt.Errorf with %w verb.
- [[errors.Unwrap]] — Extracting the immediate underlying error from an error chain.
- [[errors.Is and Custom Is()]] — Checking for target error identity across wrapped error trees.
- [[errors.As and Custom As()]] — Extracting custom typed error structs from wrapped error chains.
- [[errors.Join (Multi-Error)]] — Combining multiple concurrent or independent errors into a single error.
- [[Sentinel Errors]] — Exported package constants (io.EOF, sql.ErrNoRows) and comparison rules.
- [[Custom Error Types]] — Implementing structured error structs containing error codes and context.
- [[Error Tree Traversal Algorithms]] — Recursive unwrapping and multi-error branch traversal.
### 3. 📂 [[Panic, Recover & Architecture|03. Panic, Recover & Architecture]]
- [[panic Semantics]] — Unwinding the goroutine call stack on fatal, unrecoverable programmer errors.
- [[recover in Deferred Functions]] — Safely intercepting panics in deferred functions and converting to errors.
- [[Goroutine Panic Isolation]] — Uncaught panics inside spawned goroutines terminate the entire process.
- [[Stack Traces & runtime-debug]] — Capturing, parsing, and logging panic stack traces for observability.
- [[Domain vs Infrastructure Errors]] — Architectural separation of business rule errors vs database/network errors.
- [[Error Classification (Transient vs Permanent)]] — Categorizing errors for intelligent retries, circuit breaking, and alerting.
- [[Error Design Best Practices]] — Enriching errors without losing original context, avoiding string matching.
- [[Handle Errors, Don't Just Check]] — Meaningful error recovery and remediation vs blind error return propagation.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
