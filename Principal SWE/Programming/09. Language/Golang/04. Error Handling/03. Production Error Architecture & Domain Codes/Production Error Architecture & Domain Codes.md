---
title: Production Error Architecture & Domain Codes
tags:
  - golang
  - error-handling
  - principal-swe
parent: "[[Error Handling]]"
---

# Production Error Architecture & Domain Codes

Clean Architecture domain error hierarchy, standard enterprise error taxonomies, HTTP/gRPC translation, obfuscation, and retryability.

```text
Production Error Architecture & Domain Codes
│
├── [[Clean Architecture Domain Error Hierarchy]]
├── [[Standard Enterprise Error Codes & Error Taxonomies]]
├── [[HTTP Status & gRPC Status Code Translation Matrices]]
├── [[Error Obfuscation & Security Boundaries (Information Leakage)]]
└── [[Retryable vs Non-Retryable Error Classification]]
```

---

## 🗂️ Topics

- [[Clean Architecture Domain Error Hierarchy]] — Decoupling transport errors (HTTP/gRPC), infrastructure errors (SQL/Redis), and domain errors.
- [[Standard Enterprise Error Codes & Error Taxonomies]] — Uniform error catalog: NOT_FOUND, UNAUTHENTICATED, PERMISSION_DENIED, CONFLICT, INTERNAL.
- [[HTTP Status & gRPC Status Code Translation Matrices]] — Centralized middleware mapping internal domain errors to RFC 7807 Problem Details and gRPC status codes.
- [[Error Obfuscation & Security Boundaries (Information Leakage)]] — Stripping internal database query details and stack traces before responding to external API clients.
- [[Retryable vs Non-Retryable Error Classification]] — Defining Retryable() bool and Temporary() bool behavior interfaces for resilient client retries.

---

## 🔗 References
- ⬆️ Parent: [[Error Handling]]

