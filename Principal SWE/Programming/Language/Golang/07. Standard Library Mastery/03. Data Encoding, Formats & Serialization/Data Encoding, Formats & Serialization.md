---
title: Data Encoding, Formats & Serialization
tags:
  - golang
  - stdlib
  - principal-swe
parent: "[[Standard Library Mastery]]"
---

# Data Encoding, Formats & Serialization

encoding/json, zero-allocation JSON parsers, high-performance protobuf/flatbuffers, binary varint protocols, and gob streams.

```text
Data Encoding, Formats & Serialization
│
├── [[encoding-json Standard Serialization & Custom Marshaling]]
├── [[Zero-Allocation JSON Parsers (Sonic, Jsoniter)]]
├── [[Zero-Allocation Protobuf & FlatBuffers (vtprotobuf)]]
├── [[encoding-binary & Varint Protocol Encoding]]
└── [[encoding-gob Binary Stream Protocol]]
```

---

## 🗂️ Topics

- [[encoding-json Standard Serialization & Custom Marshaling]] — Marshaler/Unmarshaler interfaces, json.RawMessage, struct tags, and decoder streaming vs Unmarshal.
- [[Zero-Allocation JSON Parsers (Sonic, Jsoniter)]] — SIMD-accelerated JSON parsers (Bytedance Sonic) achieving 3-5x throughput over standard library.
- [[Zero-Allocation Protobuf & FlatBuffers (vtprotobuf)]] — High-performance protobuf code generators bypassing reflect-based marshaling for ultra-low latency.
- [[encoding-binary & Varint Protocol Encoding]] — Byte-level binary packing: binary.BigEndian, binary.LittleEndian, binary.Varint, and Uvarint.
- [[encoding-gob Binary Stream Protocol]] — Go-specific binary stream serialization protocol for inter-process communication and local caching.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]
- 🎓 Root: [[Principal SWE]]
