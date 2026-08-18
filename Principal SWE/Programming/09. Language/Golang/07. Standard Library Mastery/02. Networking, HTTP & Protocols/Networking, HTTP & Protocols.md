---
title: Networking, HTTP & Protocols
tags:
  - golang
  - stdlib
  - principal-swe
parent: "[[Standard Library Mastery]]"
---

# Networking, HTTP & Protocols

net socket dialing, net/http server lifecycle, client transport connection pooling, HTTP/2 & HTTP/3, graceful shutdown, and hijacking.

```text
Networking, HTTP & Protocols
│
├── [[net Package & TCP-UDP Socket Dialing]]
├── [[net-http Server Lifecycle & Context Propagation]]
├── [[net-http Client & Transport Connection Pooling]]
├── [[HTTP-2 Multiplexing & HTTP-3 QUIC Support]]
├── [[net-http Server Graceful Shutdown Lifecycle]]
└── [[net-http Connection Hijacking for WebSockets]]
```

---

## 🗂️ Topics

- [[net Package & TCP-UDP Socket Dialing]] — Low-level socket programming: net.Listen, net.Dial, net.TCPConn, keep-alive tuning, and socket buffers.
- [[net-http Server Lifecycle & Context Propagation]] — Handler interface, ServeMux routing, middleware chaining, and request context lifecycle cancellation.
- [[net-http Client & Transport Connection Pooling]] — http.Transport tuning: MaxIdleConns, MaxIdleConnsPerHost, IdleConnTimeout, and TLS handshake timeouts.
- [[HTTP-2 Multiplexing & HTTP-3 QUIC Support]] — HTTP/2 stream multiplexing, header compression (HPACK), and experimental HTTP/3 QUIC transport.
- [[net-http Server Graceful Shutdown Lifecycle]] — Handling SIGINT/SIGTERM, closing listeners, draining active in-flight requests with server.Shutdown(ctx).
- [[net-http Connection Hijacking for WebSockets]] — Taking over raw TCP connections via http.Hijacker to upgrade HTTP requests to WebSockets.

- [[net-url Parsing, Encoding & Validation Traps]] — Handling ambiguous path segments, query parameter encoding, URL canonicalization, and SSRF validation.
- [[mime & multipart Form Streaming Architecture]] — Zero-disk multipart file uploads, streaming boundary parsing, and MIME sniffing mechanics.
- [[net-http-httptrace Client Connection Diagnostics]] — Tracing DNS lookup duration, TLS handshake timing, TCP connection reuse, and request latencies.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]

