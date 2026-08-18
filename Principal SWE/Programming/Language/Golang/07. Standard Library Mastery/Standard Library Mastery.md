---
title: Standard Library Mastery
tags:
  - golang
  - stdlib
  - principal-swe
parent: "[[Golang]]"
---

# 📚 Standard Library Mastery

Deep production mastery of the Go standard library and essential database toolkits: I/O virtual filesystems, HTTP/2 & HTTP/3 networking, high-speed serialization, PostgreSQL engineering with native `pgx/v5` and `database/sql`, data structures, and timing primitives.

```text
Standard Library Mastery
│
├── [[I-O, OS & Virtual Filesystems|01. I-O, OS & Virtual Filesystems]]
│   ├── [[io Package Architecture (Reader, Writer, Closer)]]
│   ├── [[io-fs Virtual Filesystem & Directory Sandboxing]]
│   ├── [[os Package & Process Management]]
│   ├── [[bufio High-Performance Stream Buffering]]
│   ├── [[embed Standard Package & Static Assets]]
│   ├── [[flag & pflag Command-Line Parsing]]
│   └── [[path-filepath Cross-Platform Path Handling]]
├── [[Networking, HTTP & Protocols|02. Networking, HTTP & Protocols]]
│   ├── [[net Package & TCP-UDP Socket Dialing]]
│   ├── [[net-http Server Lifecycle & Context Propagation]]
│   ├── [[net-http Client & Transport Connection Pooling]]
│   ├── [[HTTP-2 Multiplexing & HTTP-3 QUIC Support]]
│   ├── [[net-http Server Graceful Shutdown Lifecycle]]
│   └── [[net-http Connection Hijacking for WebSockets]]
├── [[Data Encoding, Formats & Serialization|03. Data Encoding, Formats & Serialization]]
│   ├── [[encoding-json Standard Serialization & Custom Marshaling]]
│   ├── [[Zero-Allocation JSON Parsers (Sonic, Jsoniter)]]
│   ├── [[Zero-Allocation Protobuf & FlatBuffers (vtprotobuf)]]
│   ├── [[encoding-binary & Varint Protocol Encoding]]
│   └── [[encoding-gob Binary Stream Protocol]]
├── [[Database & PostgreSQL Engineering (database-sql, pgx)|04. Database & PostgreSQL Engineering (database-sql, pgx)]]
│   ├── [[database-sql Connection Pool Architecture]]
│   ├── [[database-sql Transactions, Prepared Statements & Isolation]]
│   ├── [[pgx Driver Architecture & Protocol vs database-sql]]
│   ├── [[pgxpool Connection Pool Architecture & Tuning]]
│   ├── [[High-Performance Batch Queries with pgx.Batch]]
│   ├── [[Bulk Data Ingestion with pgx.CopyFrom (COPY Protocol)]]
│   ├── [[Listen-Notify Asynchronous PubSub with pgx]]
│   ├── [[PostgreSQL Custom Types & JSONB Encoding with pgx]]
│   ├── [[Transactional Pipelines & Savepoints in pgx]]
│   └── [[sqlc with pgx Integration (Type-Safe SQL Compiler)]]
├── [[Data Structures, Containers & Strings|05. Data Structures, Containers & Strings]]
│   ├── [[container-heap Priority Queue & Custom Sorters]]
│   ├── [[container-list Doubly Linked List Mechanics]]
│   ├── [[container-ring Circular Ring Buffer Mechanics]]
│   ├── [[strings and bytes High-Performance Manipulation]]
│   ├── [[strings.Builder vs bytes.Buffer Zero-Alloc Comparison]]
│   ├── [[fmt Package Formatter Internals & Reflection Cost]]
│   └── [[strconv Package Fast Numeric Conversion]]
└── [[Time, Math & Cryptographic Utilities|06. Time, Math & Cryptographic Utilities]]
│   ├── [[time Package Monotonic vs Wall Clocks]]
│   ├── [[time.Ticker & time.Timer Resource Management]]
│   ├── [[math-big Arbitrary-Precision Arithmetic]]
│   └── [[math-rand-v2 Modern Fast Pseudo-Random Generator]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[I-O, OS & Virtual Filesystems|01. I-O, OS & Virtual Filesystems]]
- [[io Package Architecture (Reader, Writer, Closer)]] — Core streaming interfaces: io.Reader, io.Writer, io.Closer, io.TeeReader, io.Pipe, and io.MultiWriter.
- [[io-fs Virtual Filesystem & Directory Sandboxing]] — Go 1.16+ read-only virtual filesystem abstraction (fs.FS), fs.WalkDir, and directory security.
- [[os Package & Process Management]] — File operations, environment variables, process spawning (os.StartProcess, exec.Command), and signal forwarding.
- [[bufio High-Performance Stream Buffering]] — Minimizing OS read/write syscall overhead with bufio.Reader, bufio.Writer, and custom Scanner splits.
- [[embed Standard Package & Static Assets]] — Embedding static templates, HTML/CSS assets, and database migrations directly into compiled Go binaries.
- [[flag & pflag Command-Line Parsing]] — Command-line flag parsing architecture, custom flag.Value interfaces, and POSIX-compliant pflag integration.
- [[path-filepath Cross-Platform Path Handling]] — Cross-platform file path manipulation, filepath.Clean, filepath.Rel, and path traversal defense.
### 2. 📂 [[Networking, HTTP & Protocols|02. Networking, HTTP & Protocols]]
- [[net Package & TCP-UDP Socket Dialing]] — Low-level socket programming: net.Listen, net.Dial, net.TCPConn, keep-alive tuning, and socket buffers.
- [[net-http Server Lifecycle & Context Propagation]] — Handler interface, ServeMux routing, middleware chaining, and request context lifecycle cancellation.
- [[net-http Client & Transport Connection Pooling]] — http.Transport tuning: MaxIdleConns, MaxIdleConnsPerHost, IdleConnTimeout, and TLS handshake timeouts.
- [[HTTP-2 Multiplexing & HTTP-3 QUIC Support]] — HTTP/2 stream multiplexing, header compression (HPACK), and experimental HTTP/3 QUIC transport.
- [[net-http Server Graceful Shutdown Lifecycle]] — Handling SIGINT/SIGTERM, closing listeners, draining active in-flight requests with server.Shutdown(ctx).
- [[net-http Connection Hijacking for WebSockets]] — Taking over raw TCP connections via http.Hijacker to upgrade HTTP requests to WebSockets.
### 3. 📂 [[Data Encoding, Formats & Serialization|03. Data Encoding, Formats & Serialization]]
- [[encoding-json Standard Serialization & Custom Marshaling]] — Marshaler/Unmarshaler interfaces, json.RawMessage, struct tags, and decoder streaming vs Unmarshal.
- [[Zero-Allocation JSON Parsers (Sonic, Jsoniter)]] — SIMD-accelerated JSON parsers (Bytedance Sonic) achieving 3-5x throughput over standard library.
- [[Zero-Allocation Protobuf & FlatBuffers (vtprotobuf)]] — High-performance protobuf code generators bypassing reflect-based marshaling for ultra-low latency.
- [[encoding-binary & Varint Protocol Encoding]] — Byte-level binary packing: binary.BigEndian, binary.LittleEndian, binary.Varint, and Uvarint.
- [[encoding-gob Binary Stream Protocol]] — Go-specific binary stream serialization protocol for inter-process communication and local caching.
### 4. 📂 [[Database & PostgreSQL Engineering (database-sql, pgx)|04. Database & PostgreSQL Engineering (database-sql, pgx)]]
- [[database-sql Connection Pool Architecture]] — Connection pooling mechanics: SetMaxOpenConns, SetMaxIdleConns, SetConnMaxLifetime, and driver connections.
- [[database-sql Transactions, Prepared Statements & Isolation]] — Managing ACID transactions with tx.BeginTx, context cancellation, prepared statement caching, and rollback defers.
- [[pgx Driver Architecture & Protocol vs database-sql]] — Native PostgreSQL binary wire protocol, zero-alloc type codecs, and comparing pgx native vs database/sql.
- [[pgxpool Connection Pool Architecture & Tuning]] — pgxpool.Pool tuning: MaxConns, MinConns, MaxConnLifetime, MaxConnIdleTime, and lifecycle health hooks.
- [[High-Performance Batch Queries with pgx.Batch]] — Sending pipelined multi-statement queries in a single network round-trip for massive throughput gains.
- [[Bulk Data Ingestion with pgx.CopyFrom (COPY Protocol)]] — Streaming bulk data ingestion directly into PostgreSQL tables bypassing slow INSERT statements.
- [[Listen-Notify Asynchronous PubSub with pgx]] — Real-time PostgreSQL event streaming and Change Data Capture (CDC) using LISTEN / NOTIFY in Go.
- [[PostgreSQL Custom Types & JSONB Encoding with pgx]] — Native decoding of UUID, Hstore, IPNet, and JSONB directly into Go structs without reflection overhead.
- [[Transactional Pipelines & Savepoints in pgx]] — Nested transactions with SQL SAVEPOINT, isolation levels (Serializable, RepeatableRead), and deferred rollbacks.
- [[sqlc with pgx Integration (Type-Safe SQL Compiler)]] — Generating compile-time type-safe Go repository code from raw SQL queries targeting native pgx/v5.
### 5. 📂 [[Data Structures, Containers & Strings|05. Data Structures, Containers & Strings]]
- [[container-heap Priority Queue & Custom Sorters]] — Implementing heap.Interface (Len, Less, Swap, Push, Pop) for min/max priority queues.
- [[container-list Doubly Linked List Mechanics]] — Doubly linked list operations: PushBack, MoveToFront, Remove, and pointer memory overhead.
- [[container-ring Circular Ring Buffer Mechanics]] — Fixed-capacity circular linked list for round-robin rotation and fixed window buffers.
- [[strings and bytes High-Performance Manipulation]] — Zero-allocation string searching, splitting, trimming, and bytes slice manipulation algorithms.
- [[strings.Builder vs bytes.Buffer Zero-Alloc Comparison]] — Comparing heap allocation profiles of strings.Builder (String() zero-copy) vs bytes.Buffer.
- [[fmt Package Formatter Internals & Reflection Cost]] — How fmt.Sprintf uses reflection and interface boxing, and writing high-speed custom fmt.Formatter.
- [[strconv Package Fast Numeric Conversion]] — High-performance string-to-number parsing: strconv.ParseInt, strconv.FormatFloat, and zero-alloc byte appenders.
### 6. 📂 [[Time, Math & Cryptographic Utilities|06. Time, Math & Cryptographic Utilities]]
- [[time Package Monotonic vs Wall Clocks]] — Understanding wall clock reading (date/time) vs monotonic clock reading for elapsed duration calculations.
- [[time.Ticker & time.Timer Resource Management]] — Preventing resource leaks: time.After in loops, resetting timers, and ticker stop callbacks.
- [[math-big Arbitrary-Precision Arithmetic]] — Calculating arbitrary-precision integers (big.Int), floats (big.Float), and rational numbers.
- [[math-rand-v2 Modern Fast Pseudo-Random Generator]] — Go 1.22+ math/rand/v2: ChaCha8-based high-speed non-cryptographic random generation.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
