---
title: Incident Response, Forensics & Runtime Security Auditing
tags:
  - golang
  - security
  - principal-swe
parent: "[[Go Security]]"
---

# Incident Response, Forensics & Runtime Security Auditing

Netlink auditd event streaming, memory forensics for IoC, automated STRIDE threat modeling, and cryptographic audit trails.

```text
Incident Response, Forensics & Runtime Security Auditing
│
├── [[Linux Auditd & Netlink Security Event Ingestion]]
├── [[Memory Forensics & Runtime Heap Inspection for IoC]]
├── [[Automated STRIDE Threat Modeling for Go Architectures]]
└── [[Security Telemetry & Cryptographic Audit Trails]]
```

---

## 🗂️ Topics

- [[Linux Auditd & Netlink Security Event Ingestion]] — Streaming kernel audit logs via Netlink sockets in Go for real-time Host-based Intrusion Detection (HIDS).
- [[Memory Forensics & Runtime Heap Inspection for IoC]] — Extracting live heap dumps, parsing runtime symbol tables, and detecting in-memory code injections or Indicators of Compromise.
- [[Automated STRIDE Threat Modeling for Go Architectures]] — Generating threat matrices across Go service boundaries, data stores, external APIs, and RPC contracts.
- [[Security Telemetry & Cryptographic Audit Trails]] — Building tamper-evident audit logs with cryptographic SHA256 hash chaining and structured event schemas.

---

## 🔗 References
- ⬆️ Parent: `Security, Cryptography & Hardening in Go`

