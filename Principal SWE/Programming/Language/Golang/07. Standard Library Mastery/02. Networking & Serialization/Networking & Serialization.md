- [[net-http Connection Hijacking]] — Using http.Hijacker to upgrade HTTP connections to WebSockets or raw TCP duplex streams.

- [[net-http Server Graceful Shutdown Lifecycle]] — srv.Shutdown(ctx) vs srv.Close(), draining idle keep-alive connections, context timeouts.

- [[Zero-Allocation Protobuf & FlatBuffers (vtprotobuf)]] — High-performance protobuf code generation with zero memory allocations.

---
title: Networking & Serialization
tags:
  - golang
  - stdlib
  - principal-swe
parent: "[[Standard Library Mastery]]"
---

# Networking & Serialization

Raw TCP/UDP networking, HTTP client/server architecture, JSON encoding, and time manipulation.

```text
Networking & Serialization
│
├── [[net Package]]
├── [[net-http Server Lifecycle]]
├── [[net-http Client & Transport]]
├── [[HTTP-2 Multiplexing & HTTP-3 QUIC]]
├── [[encoding-json]]
├── [[Zero-Allocation JSON Parsers (Sonic, Jsoniter)]]
├── [[time Package]]
└── [[regexp Package]]
```

---

## 🗂️ Topics

- [[net Package]] — net.Dial, net.Listen, raw TCP/UDP socket programming, DNS resolution, and connection deadlines.
- [[net-http Server Lifecycle]] — http.Server, http.Handler, http.ServeMux, middleware chaining, and server timeouts.
- [[net-http Client & Transport]] — http.Client, http.Transport connection pooling, idle connections, and RoundTripper.
- [[HTTP-2 Multiplexing & HTTP-3 QUIC]] — HTTP/2 streaming multiplexing and HTTP/3 QUIC connection architecture in Go.
- [[encoding-json]] — Marshal, Unmarshal, struct tags, custom JSON Marshaler/Unmarshaler, and stream Decoder/Encoder.
- [[Zero-Allocation JSON Parsers (Sonic, Jsoniter)]] — SIMD-accelerated zero-copy JSON parsing libraries.
- [[time Package]] — time.Time monotonic vs wall clock, Timers, Tickers, Duration arithmetic, and time zones.
- [[regexp Package]] — Regular expression matching, compiling (regexp.MustCompile), and linear-time guarantees.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]
- 🎓 Root: [[Principal SWE]]
