---
title: Networking & Data Source
tags:
  - golang
  - source-reading
  - principal-swe
parent: "[[Go Standard Library Source Reading]]"
---

# Networking & Data Source

Analyzing the source code of net/http, database/sql, and encoding/json.

```text
Networking & Data Source
│
├── [[net-http Server.Serve & Transport Source]]
├── [[database-sql Connection Pool Source]]
└── [[encoding-json Scanner & Encoder Source]]
```

---

## 🗂️ Topics

- [[net-http Server.Serve & Transport Source]] — Server.Serve loop, conn.serve, Transport connection pooling, RoundTripper interface.
- [[database-sql Connection Pool Source]] — DB connection pool management, driver interface, tx isolation level handling.
- [[encoding-json Scanner & Encoder Source]] — Reflect-based encoder compilation cache, state machine scanner, stream buffer reuse.

---

## 🔗 References
- ⬆️ Parent: [[Go Standard Library Source Reading]]
- 🎓 Root: [[Principal SWE]]
