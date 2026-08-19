---
title: Error Handling
tags:
  - golang
  - error-handling
  - principal-swe
parent: "[[Golang]]"
---

# 🛑 Error Handling

Enterprise error handling architecture in Go: error interface mechanics, wrapping (%w), multi-errors (errors.Join), domain error codes, stack traces, panic/recover boundaries, and anti-patterns.

```text
Error Handling
│
├── `01. Error Interface Mechanics & Sentinel Errors`
│   ├── `The error Interface Contract (Error() string)`
│   ├── `Sentinel Errors Design (io.EOF, sql.ErrNoRows)`
│   ├── `Constant Errors Pattern via Defined String Types`
│   ├── `Custom Error Structs with Contextual Fields`
│   └── `The Typed Nil Error Return Trap`
├── `02. Modern Error Wrapping & Inspection (errors package)`
│   ├── `Error Wrapping Protocol with fmt.Errorf and %w`
│   ├── `errors.Is Semantic Inspection (Unwrap Single)`
│   ├── `errors.As Type Extraction & Pattern Matching`
│   ├── `Multi-Error Aggregation with errors.Join (Go 1.20+)`
│   └── `Error Tree Traversal Mechanics (Breadth-First vs Depth-First)`
├── `03. Production Error Architecture & Domain Codes`
│   ├── `Clean Architecture Domain Error Hierarchy`
│   ├── `Standard Enterprise Error Codes & Error Taxonomies`
│   ├── `HTTP Status & gRPC Status Code Translation Matrices`
│   ├── `Error Obfuscation & Security Boundaries (Information Leakage)`
│   └── `Retryable vs Non-Retryable Error Classification`
├── `04. Stack Traces & Diagnostic Enrichment`
│   ├── `Stack Trace Capture Mechanics (pkg-errors & custom)`
│   ├── `Structured Error Logging Integration with slog`
│   ├── `Distributed Trace ID & Span Attachment to Errors`
│   └── `Error Monitoring & Crash Reporting Integration (Sentry, Rollbar)`
├── `05. Panic, Recover & Boundary Isolation`
│   ├── `Panic vs Error Decision Tree (Exceptional vs Expected)`
│   ├── `Recover Boundary Middleware in HTTP & gRPC Servers`
│   ├── `Goroutine Panic Isolation Hazard (Orphan Panics)`
│   ├── `Converting Panics to Clean Domain Errors`
│   └── `Structured Stack Frame Dumps during Panic Recovery`
└── `06. Error Handling Anti-Patterns & Code Smells`
│   ├── `The Blind Error Ignoration Anti-Pattern (_ = fn())`
│   ├── `Double Logging & Double Handling Anti-Pattern`
│   ├── `String Matching on Errors Anti-Pattern`
│   ├── `Catch-All Generic Internal Error Anti-Pattern`
│   └── `Staff-Level Error Handling Principles & Guidelines`
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 `01. Error Interface Mechanics & Sentinel Errors`
- `The error Interface Contract (Error() string)` — Single-method interface simplicity, nil error representations, and zero-allocation static errors.
- `Sentinel Errors Design (io.EOF, sql.ErrNoRows)` — Predeclared exported error variables, comparing with errors.Is, and immutability pitfalls.
- `Constant Errors Pattern via Defined String Types` — Defining immutable package-level sentinel errors (type Error string) preventing global variable mutation.
- `Custom Error Structs with Contextual Fields` — Creating structured errors carrying HTTP status codes, request IDs, retry after durations, and timestamps.
- `The Typed Nil Error Return Trap` — Returning a concrete typed nil pointer as an error interface resulting in iface.tab != nil (err != nil).
### 2. 📂 `02. Modern Error Wrapping & Inspection (errors package)`
- `Error Wrapping Protocol with fmt.Errorf and %w` — Constructing causal error chains using %w vs formatting without wrapping (%v).
- `errors.Is Semantic Inspection (Unwrap Single)` — Recursive error tree traversal, custom Is(target error) bool method implementation.
- `errors.As Type Extraction & Pattern Matching` — Extracting concrete error types across nested wrapping layers with As(target any) bool.
- `Multi-Error Aggregation with errors.Join (Go 1.20+)` — Combining concurrent/batch errors into a single error, multi-branch Unwrap() []error.
- `Error Tree Traversal Mechanics (Breadth-First vs Depth-First)` — How errors.Is and errors.As navigate tree structures created by errors.Join.
### 3. 📂 `03. Production Error Architecture & Domain Codes`
- `Clean Architecture Domain Error Hierarchy` — Decoupling transport errors (HTTP/gRPC), infrastructure errors (SQL/Redis), and domain errors.
- `Standard Enterprise Error Codes & Error Taxonomies` — Uniform error catalog: NOT_FOUND, UNAUTHENTICATED, PERMISSION_DENIED, CONFLICT, INTERNAL.
- `HTTP Status & gRPC Status Code Translation Matrices` — Centralized middleware mapping internal domain errors to RFC 7807 Problem Details and gRPC status codes.
- `Error Obfuscation & Security Boundaries (Information Leakage)` — Stripping internal database query details and stack traces before responding to external API clients.
- `Retryable vs Non-Retryable Error Classification` — Defining Retryable() bool and Temporary() bool behavior interfaces for resilient client retries.
### 4. 📂 `04. Stack Traces & Diagnostic Enrichment`
- `Stack Trace Capture Mechanics (pkg-errors & custom)` — Capturing caller program counters (runtime.Callers) at error creation sites with minimal CPU overhead.
- `Structured Error Logging Integration with slog` — Emitting error causes, stack frames, and contextual key-value pairs to structured log streams.
- `Distributed Trace ID & Span Attachment to Errors` — Correlating errors with active OpenTelemetry span contexts for instant observability triaging.
- `Error Monitoring & Crash Reporting Integration (Sentry, Rollbar)` — Automated capture, fingerprinting, and grouping of unhandled production errors.
### 5. 📂 `05. Panic, Recover & Boundary Isolation`
- `Panic vs Error Decision Tree (Exceptional vs Expected)` — Staff-level heuristics: when is panic acceptable (programmer bug, startup failure) vs error.
- `Recover Boundary Middleware in HTTP & gRPC Servers` — Capturing uncaught goroutine panics, logging stack traces, and returning 500 Internal Server Error.
- `Goroutine Panic Isolation Hazard (Orphan Panics)` — Why a panic in a spawned background goroutine terminates the entire process if unrecovered inside itself.
- `Converting Panics to Clean Domain Errors` — Safely capturing third-party panics inside library boundaries and translating them to structured errors.
- `Structured Stack Frame Dumps during Panic Recovery` — Using debug.Stack() to capture and sanitize full multi-goroutine traces upon panic.
### 6. 📂 `06. Error Handling Anti-Patterns & Code Smells`
- `The Blind Error Ignoration Anti-Pattern (_ = fn())` — Security and state corruption risks of silently swallowing return errors.
- `Double Logging & Double Handling Anti-Pattern` — Logging an error and returning it upward, leading to log spam and duplicate alert storms.
- `String Matching on Errors Anti-Pattern` — Fragile strings.Contains(err.Error(), "not found") checks vs robust errors.Is/errors.As.
- `Catch-All Generic Internal Error Anti-Pattern` — Wrapping everything in generic errors.New("something went wrong") destroying observability.
- `Staff-Level Error Handling Principles & Guidelines` — Non-negotiable engineering principles: wrap once at the boundary, handle once, log once.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

