---
title: Networking & Serialization
tags:
  - golang
  - stdlib
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
├── [[encoding-json]]
├── [[time Package]]
└── [[regexp Package]]
```

---

## 🗂️ Topics

- [[net Package]] — net.Dial, net.Listen, raw TCP/UDP socket programming, DNS resolution, connection deadlines.
- [[net-http Server Lifecycle]] — http.Server, http.Handler, http.ServeMux, middleware chaining, timeouts.
- [[net-http Client & Transport]] — http.Client, http.Transport connection pooling, idle connections, RoundTripper.
- [[encoding-json]] — Marshal, Unmarshal, struct tags, custom JSON Marshaler/Unmarshaler, stream decoders.
- [[time Package]] — time.Time monotonic vs wall clock, Timers, Tickers, Duration math, time zones.
- [[regexp Package]] — Regular expression matching, compiling (regexp.MustCompile), and performance considerations.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]
- 🎓 Root: [[Principal SWE]]
