---
title: Application Hardening & OWASP Top 10 Defense
tags:
  - golang
  - security
  - principal-swe
parent: "[[Security, Cryptography & Hardening in Go]]"
---

# Application Hardening & OWASP Top 10 Defense

SQL injection defense, SSRF mitigation with IP pinning, os.Root path traversal sandboxing, DoS protection, and XSS auto-escaping.

```text
Application Hardening & OWASP Top 10 Defense
│
├── [[SQL Injection Defense & Parameterized Queries]]
├── [[Server-Side Request Forgery (SSRF) Defense & IP Pinning]]
├── [[Path Traversal & Zip Slip Prevention (os.Root)]]
├── [[Denial-of-Service Defense (Slowloris, Body Limits, Timeouts)]]
└── [[Cross-Site Scripting (XSS) & Template Escaping (html-template)]]
```

---

## 🗂️ Topics

- [[SQL Injection Defense & Parameterized Queries]] — Strict enforcement of parameterized queries in database/sql and pgx.
- [[Server-Side Request Forgery (SSRF) Defense & IP Pinning]] — Custom http.Transport dialer validating against private IP ranges (RFC 1918) and DNS rebinding.
- [[Path Traversal & Zip Slip Prevention (os.Root)]] — Go 1.24+ os.Root directory sandboxing and path cleaning.
- [[Denial-of-Service Defense (Slowloris, Body Limits, Timeouts)]] — Configuring ReadTimeout, WriteTimeout, IdleTimeout, and http.MaxBytesReader.
- [[Cross-Site Scripting (XSS) & Template Escaping (html-template)]] — Context-aware auto-escaping rules in html/template vs text/template.

---

## 🔗 References
- ⬆️ Parent: [[Security, Cryptography & Hardening in Go]]

