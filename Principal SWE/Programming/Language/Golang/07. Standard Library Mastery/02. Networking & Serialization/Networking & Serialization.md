---
title: Networking & Serialization
tags:
  - golang
  - stdlib
parent: "[[Standard Library Mastery]]"
---

# Networking & Serialization

Socket programming, HTTP servers and clients, JSON serialization, and time operations.

```text
Networking & Serialization
│
├── [[net Package]]
├── [[net-http Server Lifecycle]]
├── [[net-http Client & Transport]]
├── [[encoding-json]]
├── [[time Package]]
└── [[regexp Package]]
```

---

## 🗂️ Topics

- [[net Package]] — net.Dial, net.Listen, raw TCP/UDP socket programming, DNS resolution, and connection deadlines.
- [[net-http Server Lifecycle]] — http.Server, http.Handler, http.ServeMux, middleware chaining, and server timeouts.
- [[net-http Client & Transport]] — http.Client, http.Transport connection pooling, idle connections, and RoundTripper.
- [[encoding-json]] — Marshal, Unmarshal, struct tags, custom JSON Marshaler/Unmarshaler, and stream Decoder/Encoder.
- [[time Package]] — time.Time monotonic vs wall clock, Timers, Tickers, Duration arithmetic, and time zones.
- [[regexp Package]] — Regular expression matching, compiling (regexp.MustCompile), and performance characteristics.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]
- 🎓 Root: [[Principal SWE]]
