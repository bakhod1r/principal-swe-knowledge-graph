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
├── `01. I-O, OS & Virtual Filesystems`
│   ├── `io Package Architecture (Reader, Writer, Closer)`
│   ├── `io-fs Virtual Filesystem & Directory Sandboxing`
│   ├── `os Package & Process Management`
│   ├── `bufio High-Performance Stream Buffering`
│   ├── `embed Standard Package & Static Assets`
│   ├── `flag & pflag Command-Line Parsing`
│   ├── `path-filepath Cross-Platform Path Handling`
│   ├── `archive-tar & archive-zip Streaming Extraction`
│   ├── `compress-gzip & flate High-Throughput Stream Compression`
│   └── `syscall & golang.org-x-sys-unix Low-Level System Calls`
├── `02. Networking, HTTP & Protocols`
│   ├── `net Package & TCP-UDP Socket Dialing`
│   ├── `net-http Server Lifecycle & Context Propagation`
│   ├── `net-http Client & Transport Connection Pooling`
│   ├── `HTTP-2 Multiplexing & HTTP-3 QUIC Support`
│   ├── `net-http Server Graceful Shutdown Lifecycle`
│   ├── `net-http Connection Hijacking for WebSockets`
│   ├── `net-url Parsing, Encoding & Validation Traps`
│   ├── `mime & multipart Form Streaming Architecture`
│   └── `net-http-httptrace Client Connection Diagnostics`
├── [[Data Encoding, Formats & Serialization|03. Data Encoding, Formats & Serialization]]
│   ├── `encoding-json Standard Serialization & Custom Marshaling`
│   ├── `Zero-Allocation JSON Parsers (Sonic, Jsoniter)`
│   ├── `Zero-Allocation Protobuf & FlatBuffers (vtprotobuf)`
│   ├── `encoding-binary & Varint Protocol Encoding`
│   ├── `encoding-gob Binary Stream Protocol`
│   ├── `encoding-csv High-Speed Stream Reading & Lazy Quotes`
│   ├── `encoding-xml Parser Architecture & Security Constraints`
│   └── `encoding-base64 & encoding-hex Zero-Alloc Appenders`
├── `04. Database & PostgreSQL Engineering (database-sql, pgx)`
│   ├── `database-sql Connection Pool Architecture`
│   ├── `database-sql Transactions, Prepared Statements & Isolation`
│   ├── `pgx Driver Architecture & Protocol vs database-sql`
│   ├── `pgxpool Connection Pool Architecture & Tuning`
│   ├── `High-Performance Batch Queries with pgx.Batch`
│   ├── `Bulk Data Ingestion with pgx.CopyFrom (COPY Protocol)`
│   ├── `Listen-Notify Asynchronous PubSub with pgx`
│   ├── `PostgreSQL Custom Types & JSONB Encoding with pgx`
│   ├── `Transactional Pipelines & Savepoints in pgx`
│   ├── `sqlc with pgx Integration (Type-Safe SQL Compiler)`
│   ├── `Prepared Statement Caching & Query Planning Overhead`
│   ├── `PostgreSQL Logical Replication & CDC in Go (pglogrepl)`
│   ├── `PgBouncer Connection Pooling & pgx Compatibility`
│   ├── `Optimistic vs Pessimistic Locking (SELECT FOR UPDATE)`
│   ├── `Zero-Downtime Database Migrations (golang-migrate)`
│   ├── `PostgreSQL Advisory Locks for Distributed Coordination`
│   ├── `Keyset Pagination vs OFFSET Performance in Go`
│   ├── `Database Circuit Breaking & Query Timeout Budgets`
│   └── `Scanning Dynamic & NULL Values (sql.Null vs pgtype)`
├── `05. Data Structures, Containers & Strings`
│   ├── `container-heap Priority Queue & Custom Sorters`
│   ├── `container-list Doubly Linked List Mechanics`
│   ├── `container-ring Circular Ring Buffer Mechanics`
│   ├── `strings and bytes High-Performance Manipulation`
│   ├── `strings.Builder vs bytes.Buffer Zero-Alloc Comparison`
│   ├── `fmt Package Formatter Internals & Reflection Cost`
│   ├── `strconv Package Fast Numeric Conversion`
│   ├── `sync-atomic Type-Safe Pointers & Value Types`
│   └── `unicode & unicode-utf8 Fast Decoding Engines`
└── `06. Time, Math & Cryptographic Utilities`
    ├── `time Package Monotonic vs Wall Clocks`
    ├── `time.Ticker & time.Timer Resource Management`
    ├── `math-big Arbitrary-Precision Arithmetic`
    ├── `math-rand-v2 Modern Fast Pseudo-Random Generator`
    ├── `context Package Architecture & Cancellation Trees`
    ├── `math-cmplx Complex Mathematical Operations`
    └── `hash Package Hierarchy & CRC32 Hardware Checksums`
```

---

## 🗂️ Core Categories & Topics

### 📂 [[I-O|01. I-O]]
- [[archive-tar & archive-zip Streaming Extraction]]
- [[bufio High-Performance Stream Buffering]]
- [[compress-gzip & flate High-Throughput Stream Compression]]
- [[io Package Architecture (Reader, Writer, Closer)]]

### 📂 [[OS Interfaces|08. OS Interfaces]]
- [[embed Standard Package & Static Assets]]
- [[flag & pflag Command-Line Parsing]]
- [[io-fs Virtual Filesystem & Directory Sandboxing]]
- [[os Package & Process Management]]
- [[path-filepath Cross-Platform Path Handling]]
- [[syscall & golang.org-x-sys-unix Low-Level System Calls]]

### 📂 [[Networking|02. Networking]]

### 📂 [[Data Encoding, Formats & Serialization|03. Data Encoding, Formats & Serialization]]
- [[encoding-json Standard Serialization & Custom Marshaling]] — Marshaler/Unmarshaler interfaces, json.RawMessage, struct tags, and decoder streaming vs Unmarshal.
- [[Zero-Allocation JSON Parsers (Sonic, Jsoniter)]] — SIMD-accelerated JSON parsers (Bytedance Sonic) achieving 3-5x throughput over standard library.
- [[Zero-Allocation Protobuf & FlatBuffers (vtprotobuf)]] — High-performance protobuf code generators bypassing reflect-based marshaling for ultra-low latency.
- [[encoding-binary & Varint Protocol Encoding]] — Byte-level binary packing: binary.BigEndian, binary.LittleEndian, binary.Varint, and Uvarint.
- [[encoding-gob Binary Stream Protocol]] — Go-specific binary stream serialization protocol for inter-process communication and local caching.
- [[encoding-csv High-Speed Stream Reading & Lazy Quotes]] — Handling multiline CSV fields, custom delimiters, lazy quotes, and memory-efficient record streaming.
- [[encoding-xml Parser Architecture & Security Constraints]] — XML decoder streaming, entity expansion limits, custom xml.Marshaler, and namespace parsing.
- [[encoding-base64 & encoding-hex Zero-Alloc Appenders]] — Go 1.22+ base64.Encoding.AppendEncode and hex.AppendEncode for zero-allocation slice appending.

### 📂 [[Database|04. Database]]

### 📂 [[Containers|05. Containers]]
- [[container-heap Priority Queue & Custom Sorters]]
- [[container-list Doubly Linked List Mechanics]]
- [[container-ring Circular Ring Buffer Mechanics]]
- [[sync-atomic Type-Safe Pointers & Value Types]]

### 📂 [[Strings and Text|09. Strings and Text]]
- [[fmt Package Formatter Internals & Reflection Cost]]
- [[strconv Package Fast Numeric Conversion]]
- [[strings and bytes High-Performance Manipulation]]
- [[strings.Builder vs bytes.Buffer Zero-Alloc Comparison]]
- [[unicode & unicode-utf8 Fast Decoding Engines]]

### 📂 [[Time|06. Time]]
- [[context Package Architecture & Cancellation Trees]]
- [[time Package Monotonic vs Wall Clocks]]
- [[time.Ticker & time.Timer Resource Management]]

### 📂 [[Math Utilities|11. Math Utilities]]
- [[hash Package Hierarchy & CRC32 Hardware Checksums]]
- [[math-big Arbitrary-Precision Arithmetic]]
- [[math-cmplx Complex Mathematical Operations]]
- [[math-rand-v2 Modern Fast Pseudo-Random Generator]]

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

---

## 🗂️ Topics

- [[Containers]]
- [[Data Encoding, Formats & Serialization]]
- [[Database]]
- [[HTTP]]
- [[I-O]]
- [[Math Utilities]]
- [[Networking]]
- [[OS Interfaces]]
- [[PostgreSQL Engineering]]
- [[Strings and Text]]
- [[Time]]
