---
title: Enterprise Secrets Management & Key Vaults (HSM, KMS)
tags:
  - golang
  - security
  - principal-swe
parent: "[[Security, Cryptography & Hardening in Go]]"
---

# Enterprise Secrets Management & Key Vaults (HSM, KMS)

HashiCorp Vault integration, PKCS#11 HSM hardware keys, Cloud KMS envelope encryption, and key rotation.

```text
Enterprise Secrets Management & Key Vaults (HSM, KMS)
│
├── [[HashiCorp Vault Dynamic Secret Leasing & Token Renewal]]
├── [[Hardware Security Module (HSM) Integration via PKCS#11]]
├── [[Cloud KMS Envelope Encryption (AWS KMS, GCP KMS)]]
├── [[Kubernetes Secrets Decryption & External Secrets Operator]]
└── [[Automated Key Rotation Pipelines & Multi-Key Decryption]]
```

---

## 🗂️ Topics

- [[HashiCorp Vault Dynamic Secret Leasing & Token Renewal]] — Integrating vault/api, rotating database credentials, and managing dynamic TLS certificates in Go.
- [[Hardware Security Module (HSM) Integration via PKCS#11]] — Offloading cryptographic operations and private keys to hardware security modules (miekg/pkcs11).
- [[Cloud KMS Envelope Encryption (AWS KMS, GCP KMS)]] — Encrypting data with local Data Encryption Keys (DEKs) wrapped by KMS Key Encryption Keys (KEKs).
- [[Kubernetes Secrets Decryption & External Secrets Operator]] — Ingesting encrypted secrets directly into Go application memory without disk writes.
- [[Automated Key Rotation Pipelines & Multi-Key Decryption]] — Handling versioned encryption keys, decrypting legacy data, and re-wrapping payloads.

---

## 🔗 References
- ⬆️ Parent: [[Security, Cryptography & Hardening in Go]]
- 🎓 Root: [[Principal SWE]]
