---
title: Audit Logging
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Audit Logging

Tamper-evident audit trails, cryptographic hash chaining, compliance event schemas, PII redaction, and async ingestion pipelines.

```text
Audit Logging
│
├── [[Tamper-Evident Audit Trails with Cryptographic Hash Chaining]]
├── [[Structured Audit Event Schemas (Actor, Action, Resource, Context)]]
├── [[PII Redaction and Secret Masking at Ingestion Boundaries]]
└── [[Asynchronous Non-Blocking Audit Ingestion and Dead-Letter Queues]]
```

---

## 🗂️ Topics

- [[Tamper-Evident Audit Trails with Cryptographic Hash Chaining]] — Securing financial and security audit logs with immutable SHA-256 block hashing.
- [[Structured Audit Event Schemas (Actor, Action, Resource, Context)]] — Standardizing SOC2 and NIST compliance event schemas across distributed microservices.
- [[PII Redaction and Secret Masking at Ingestion Boundaries]] — Zero-leakage automated redaction of credit card numbers, passwords, and sensitive identity data.
- [[Asynchronous Non-Blocking Audit Ingestion and Dead-Letter Queues]] — Preventing audit logging network stalls from degrading critical user transaction paths.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]

