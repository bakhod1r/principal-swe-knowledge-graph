---
title: Transport Layer Security & Certificates (crypto-tls, crypto-x509)
tags:
  - golang
  - security
  - principal-swe
parent: "[[Security, Cryptography & Hardening in Go]]"
---

# Transport Layer Security & Certificates (crypto-tls, crypto-x509)

TLS 1.3 architecture, Mutual TLS (mTLS), X.509 certificate chains, zero-downtime certificate rotation, and TLS keylog forensics.

```text
Transport Layer Security & Certificates (crypto-tls, crypto-x509)
│
├── [[TLS 1.3 Architecture & Zero-RTT Sessions (crypto-tls)]]
├── [[Mutual TLS (mTLS) & Client Certificate Verification]]
├── [[X.509 Certificate Generation, Parsing & Revocation (CRL, OCSP)]]
├── [[Certificate Rotation Without Downtime (GetCertificate & Dynamic Reload)]]
└── [[Custom TLS KeyLogWriter for Network Forensics]]
```

---

## 🗂️ Topics

- [[TLS 1.3 Architecture & Zero-RTT Sessions (crypto-tls)]] — TLS 1.3 handshake, cipher suites, session tickets, and ALPN negotiation.
- [[Mutual TLS (mTLS) & Client Certificate Verification]] — Zero-trust service-to-service authentication with custom ClientAuth and RootCAs pools.
- [[X.509 Certificate Generation, Parsing & Revocation (CRL, OCSP)]] — Parsing PEM/DER certificates, building in-memory CA chains, and OCSP stapling.
- [[Certificate Rotation Without Downtime (GetCertificate & Dynamic Reload)]] — Dynamically reloading TLS certificates on file change without restarting HTTP/gRPC servers.
- [[Custom TLS KeyLogWriter for Network Forensics]] — Exporting TLS pre-master secrets for Wireshark inspection during security debugging.

---

## 🔗 References
- ⬆️ Parent: [[Security, Cryptography & Hardening in Go]]

