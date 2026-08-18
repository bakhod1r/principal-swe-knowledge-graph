---
title: Standard Library Mastery
tags:
  - golang
  - stdlib
  - principal-swe
parent: "[[Golang]]"
---

# 📚 Standard Library Mastery

Deep-dive into essential Go standard library packages: io, os, net/http, encoding/json, slog, time, bufio, crypto, database/sql, and containers.

```text
Standard Library Mastery
│
├── [[I-O, OS & System|01. I-O, OS & System]]
│   ├── [[io Package (Reader, Writer, Closer)]]
│   ├── [[bufio Package]]
│   ├── [[os Package]]
│   ├── [[path-filepath Package]]
│   ├── [[flag Package]]
│   ├── [[embed Standard Package & Static Assets]]
│   └── [[io-fs Virtual Filesystem]]
├── [[Networking & Serialization|02. Networking & Serialization]]
│   ├── [[net Package]]
│   ├── [[net-http Server Lifecycle]]
│   ├── [[net-http Client & Transport]]
│   ├── [[HTTP-2 Multiplexing & HTTP-3 QUIC]]
│   ├── [[encoding-json]]
│   ├── [[Zero-Allocation JSON Parsers (Sonic, Jsoniter)]]
│   ├── [[time Package]]
│   └── [[regexp Package]]
├── [[Observability, Security & Persistence|03. Observability, Security & Persistence]]
│   ├── [[log-slog (Structured Logging)]]
│   ├── [[database-sql Connection Pool]]
│   ├── [[database-sql Transactions & Queries]]
│   ├── [[crypto-tls & Certificates]]
│   ├── [[crypto Cryptography Primitives]]
│   ├── [[crypto-subtle Constant-Time Operations]]
│   ├── [[reflect Laws of Reflection]]
│   └── [[unsafe Zero-Copy Operations]]
└── [[Data Structures, Strings & Formats|04. Data Structures, Strings & Formats]]
│   ├── [[container-list (Doubly Linked List)]]
│   ├── [[container-heap (Priority Queue)]]
│   ├── [[container-ring (Circular List)]]
│   ├── [[strings and bytes Packages]]
│   ├── [[strings.Builder vs Buffer]]
│   ├── [[strconv Package]]
│   └── [[fmt Package Internals & Reflection Cost]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[I-O, OS & System|01. I-O, OS & System]]
- [[io Package (Reader, Writer, Closer)]] — Streaming abstractions, io.Copy, io.Pipe, io.MultiReader, io.TeeReader, io.LimitReader.
- [[bufio Package]] — bufio.Reader, bufio.Writer, bufio.Scanner for high-throughput buffered stream processing.
- [[os Package]] — Environment variables, process management, POSIX signal handling, exit codes, os.File operations.
- [[path-filepath Package]] — Cross-platform file path manipulation, filepath.Walk, filepath.Clean, filepath.Join.
- [[flag Package]] — Command-line argument parsing and custom flag value types.
- [[embed Standard Package & Static Assets]] — Embedding static assets, templates, and directory trees directly into binary executables.
- [[io-fs Virtual Filesystem]] — fs.FS, fs.File, fs.WalkDir virtual filesystem interface abstraction.
### 2. 📂 [[Networking & Serialization|02. Networking & Serialization]]
- [[net Package]] — net.Dial, net.Listen, raw TCP/UDP socket programming, DNS resolution, and connection deadlines.
- [[net-http Server Lifecycle]] — http.Server, http.Handler, http.ServeMux, middleware chaining, and server timeouts.
- [[net-http Client & Transport]] — http.Client, http.Transport connection pooling, idle connections, and RoundTripper.
- [[HTTP-2 Multiplexing & HTTP-3 QUIC]] — HTTP/2 streaming multiplexing and HTTP/3 QUIC connection architecture in Go.
- [[encoding-json]] — Marshal, Unmarshal, struct tags, custom JSON Marshaler/Unmarshaler, and stream Decoder/Encoder.
- [[Zero-Allocation JSON Parsers (Sonic, Jsoniter)]] — SIMD-accelerated zero-copy JSON parsing libraries.
- [[time Package]] — time.Time monotonic vs wall clock, Timers, Tickers, Duration arithmetic, and time zones.
- [[regexp Package]] — Regular expression matching, compiling (regexp.MustCompile), and linear-time guarantees.
### 3. 📂 [[Observability, Security & Persistence|03. Observability, Security & Persistence]]
- [[log-slog (Structured Logging)]] — log/slog Logger, JSON/Text Handlers, Level management, Attributes, and Groups.
- [[database-sql Connection Pool]] — sql.DB connection pool lifecycle, SetMaxOpenConns, SetMaxIdleConns, connection recycling.
- [[database-sql Transactions & Queries]] — Prepared statements, QueryRow, Exec, transaction isolation, Context cancellation.
- [[crypto-tls & Certificates]] — crypto/tls, mutual TLS (mTLS), certificate verification, TLS 1.3 cipher suites.
- [[crypto Cryptography Primitives]] — crypto/rand secure randomness, AES-GCM encryption, SHA-256 hashing.
- [[crypto-subtle Constant-Time Operations]] — Preventing timing side-channel attacks during cryptographic comparisons.
- [[reflect Laws of Reflection]] — reflect.Type, reflect.Value, Interface to Reflection, Settability, Struct inspection.
- [[unsafe Zero-Copy Operations]] — unsafe.Pointer, unsafe.Slice, unsafe.String zero-copy byte/string conversions.
### 4. 📂 [[Data Structures, Strings & Formats|04. Data Structures, Strings & Formats]]
- [[container-list (Doubly Linked List)]] — Standard doubly linked list implementation with element manipulation.
- [[container-heap (Priority Queue)]] — Implementing heap.Interface (Len, Less, Swap, Push, Pop) for priority queues.
- [[container-ring (Circular List)]] — Circular ring buffer data structure in standard library.
- [[strings and bytes Packages]] — High-performance string search, splitting, joining, and bytes.Buffer pool idioms.
- [[strings.Builder vs Buffer]] — Zero-allocation string concatenation and String() memory copy elimination.
- [[strconv Package]] — FormatInt, ParseInt, Atoi, Itoa, and string-to-number conversions.
- [[fmt Package Internals & Reflection Cost]] — Formatting verbs, Sprintf reflection overhead, and avoiding fmt in hot loops.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
