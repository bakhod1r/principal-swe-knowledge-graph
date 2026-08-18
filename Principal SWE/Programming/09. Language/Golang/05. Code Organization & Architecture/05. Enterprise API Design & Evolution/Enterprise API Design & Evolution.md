---
title: Enterprise API Design & Evolution
tags:
  - golang
  - architecture
  - principal-swe
parent: "[[Code Organization & Architecture]]"
---

# Enterprise API Design & Evolution

RESTful API architecture with net/http and chi, gRPC service contracts, API versioning strategies, OpenAPI generation, and compatibility.

```text
Enterprise API Design & Evolution
│
├── [[RESTful API Architecture with net-http and chi]]
├── [[gRPC and Protocol Buffers Service Architecture]]
├── [[API Versioning Strategies (URL, Header, Subdomain)]]
├── [[OpenAPI & Swagger Documentation Generation]]
└── [[Backward Compatibility & Breaking Change Prevention]]
```

---

## 🗂️ Topics

- [[RESTful API Architecture with net-http and chi]] — Building fast, idiomatic REST services with sub-routers and context middlewares.
- [[gRPC and Protocol Buffers Service Architecture]] — Defining .proto service contracts, generating Go stubs, interceptors, and streaming RPCs.
- [[API Versioning Strategies (URL, Header, Subdomain)]] — Backward-compatible schema evolution, deprecation timelines, and field migrations.
- [[OpenAPI & Swagger Documentation Generation]] — Auto-generating OpenAPI specs from Go doc comments and declarative annotations.
- [[Backward Compatibility & Breaking Change Prevention]] — Protobuf field number rules, additive JSON extensions, and contract testing.

---

## 🔗 References
- ⬆️ Parent: [[Code Organization & Architecture]]

