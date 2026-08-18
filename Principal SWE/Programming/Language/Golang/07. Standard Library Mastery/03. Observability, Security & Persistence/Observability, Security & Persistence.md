---
title: Observability, Security & Persistence
tags:
  - golang
  - stdlib
parent: "[[Standard Library Mastery]]"
---

# Observability, Security & Persistence

Structured logging, database pools, cryptography, reflection, and zero-copy unsafe operations.

```text
Observability, Security & Persistence
│
├── [[log-slog (Structured Logging)]]
├── [[database-sql Connection Pool]]
├── [[database-sql Transactions & Queries]]
├── [[crypto-tls & Certificates]]
├── [[crypto Cryptography Primitives]]
├── [[reflect Laws of Reflection]]
└── [[unsafe Zero-Copy Operations]]
```

---

## 🗂️ Topics

- [[log-slog (Structured Logging)]] — log/slog Logger, JSON/Text Handlers, Level management, Attributes, and Groups.
- [[database-sql Connection Pool]] — sql.DB connection pool lifecycle, SetMaxOpenConns, SetMaxIdleConns, connection recycling.
- [[database-sql Transactions & Queries]] — Prepared statements, QueryRow, Exec, transaction isolation, Context cancellation.
- [[crypto-tls & Certificates]] — crypto/tls, mutual TLS (mTLS), certificate verification, TLS 1.3 cipher suites.
- [[crypto Cryptography Primitives]] — crypto/rand secure randomness, AES-GCM encryption, SHA-256 hashing.
- [[reflect Laws of Reflection]] — reflect.Type, reflect.Value, Interface to Reflection, Settability, Struct inspection.
- [[unsafe Zero-Copy Operations]] — unsafe.Pointer, unsafe.Slice, unsafe.String zero-copy byte/string conversions.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]
- 🎓 Root: [[Principal SWE]]
