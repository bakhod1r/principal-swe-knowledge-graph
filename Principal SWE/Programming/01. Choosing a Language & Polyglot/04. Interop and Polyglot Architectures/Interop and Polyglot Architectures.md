---
title: Interop and Polyglot Architectures
tags:
  - programming
  - polyglot
  - principal-swe
parent: "[[Choosing a Language & Polyglot]]"
---

# Interop and Polyglot Architectures

C-ABI foreign function interfaces, gRPC/Protobuf contracts, WebAssembly WASI runtimes, shared memory IPC, and serialization taxes.

```text
Interop and Polyglot Architectures
│
├── [[Foreign Function Interface (FFI) and C-ABI Compatibility]]
├── [[Contract-First Polyglot Microservices (gRPC and Protobuf)]]
├── [[WebAssembly (WASM and WASI) as a Universal Polyglot Engine]]
├── [[Shared Memory, Ring Buffers, and Low-Latency IPC]]
└── [[Managing Polyglot Data Serialization Overhead]]
```

---

## 🗂️ Topics

- [[Foreign Function Interface (FFI) and C-ABI Compatibility]] — Zero-copy memory sharing, cross-language pointer passing, and data structure memory layout across language boundaries.
- [[Contract-First Polyglot Microservices (gRPC and Protobuf)]] — Decoupling language runtimes behind backward-compatible, strictly typed RPC interface definitions.
- [[WebAssembly (WASM and WASI) as a Universal Polyglot Engine]] — Embedding untrusted polyglot plugins (Rust, Go, C++, Python) inside secure host application runtimes.
- [[Shared Memory, Ring Buffers, and Low-Latency IPC]] — High-speed inter-process communication using mmap, Unix domain sockets, and shared memory ring buffers.
- [[Managing Polyglot Data Serialization Overhead]] — Evaluating serialization and deserialization CPU taxes across JSON, Protobuf, FlatBuffers, Cap n Proto, and Avro.

---

## 🔗 References
- ⬆️ Parent: [[Choosing a Language & Polyglot]]

