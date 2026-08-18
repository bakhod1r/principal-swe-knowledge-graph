---
title: Standard Library Mastery
tags:
  - golang
  - stdlib
  - principal-swe
parent: "[[Golang]]"
---

# 📚 Standard Library Mastery

Deep-dive into essential Go standard library packages: io, os, net/http, encoding/json, slog, time, bufio, crypto, and database/sql.

```text
Standard Library Mastery
│
├── [[I-O, OS & System|01. I-O, OS & System]]
│   ├── [[io Package (Reader, Writer, Closer)]]
│   ├── [[bufio Package]]
│   ├── [[os Package]]
│   ├── [[path-filepath Package]]
│   ├── [[flag Package]]
│   └── [[go:embed Directive]]
├── [[Networking & Serialization|02. Networking & Serialization]]
│   ├── [[net Package]]
│   ├── [[net-http Server Lifecycle]]
│   ├── [[net-http Client & Transport]]
│   ├── [[encoding-json]]
│   ├── [[time Package]]
│   └── [[regexp Package]]
└── [[Observability, Security & Persistence|03. Observability, Security & Persistence]]
│   ├── [[log-slog (Structured Logging)]]
│   ├── [[database-sql Connection Pool]]
│   ├── [[database-sql Transactions & Queries]]
│   ├── [[crypto-tls & Certificates]]
│   ├── [[crypto Cryptography Primitives]]
│   ├── [[reflect Laws of Reflection]]
│   └── [[unsafe Zero-Copy Operations]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[I-O, OS & System|01. I-O, OS & System]]
- [[io Package (Reader, Writer, Closer)]] — Streaming abstractions, io.Copy, io.Pipe, io.MultiReader, io.TeeReader.
- [[bufio Package]] — bufio.Reader, bufio.Writer, bufio.Scanner for high-throughput buffered stream processing.
- [[os Package]] — Environment variables, process management, POSIX signal handling, exit codes, os.File.
- [[path-filepath Package]] — Cross-platform file path manipulation, filepath.Walk, filepath.Clean.
- [[flag Package]] — Command-line argument parsing and custom flag value types.
- [[go:embed Directive]] — Embedding static files, templates, and directory trees directly into binaries.
### 2. 📂 [[Networking & Serialization|02. Networking & Serialization]]
- [[net Package]] — net.Dial, net.Listen, raw TCP/UDP socket programming, DNS resolution, connection deadlines.
- [[net-http Server Lifecycle]] — http.Server, http.Handler, http.ServeMux, middleware chaining, timeouts.
- [[net-http Client & Transport]] — http.Client, http.Transport connection pooling, idle connections, RoundTripper.
- [[encoding-json]] — Marshal, Unmarshal, struct tags, custom JSON Marshaler/Unmarshaler, stream decoders.
- [[time Package]] — time.Time monotonic vs wall clock, Timers, Tickers, Duration math, time zones.
- [[regexp Package]] — Regular expression matching, compiling (regexp.MustCompile), and performance considerations.
### 3. 📂 [[Observability, Security & Persistence|03. Observability, Security & Persistence]]
- [[log-slog (Structured Logging)]] — log/slog Logger, JSON/Text Handlers, Level management, Attributes, Groups.
- [[database-sql Connection Pool]] — sql.DB connection lifecycle, SetMaxOpenConns, SetMaxIdleConns, connection recycling.
- [[database-sql Transactions & Queries]] — Prepared statements, QueryRow, Exec, transaction isolation, Context cancellation.
- [[crypto-tls & Certificates]] — crypto/tls, mutual TLS (mTLS), certificate verification, TLS 1.3 cipher suites.
- [[crypto Cryptography Primitives]] — crypto/rand secure randomness, AES-GCM encryption, SHA-256 hashing.
- [[reflect Laws of Reflection]] — reflect.Type, reflect.Value, Interface to Reflection, Settability, Struct inspection.
- [[unsafe Zero-Copy Operations]] — unsafe.Pointer, unsafe.Slice, unsafe.String zero-copy byte/string conversions.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
