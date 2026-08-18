---
title: API Security Best Practices
tags:
  - best-practices
  - engineering-excellence
  - api-security-best-practices
  - principal-swe
parent: "[[Best Practices]]"
---

# 🏛️ API Security Best Practices

Defensive API engineering standards: Identity verification (OAuth2/OIDC, JWT rotation), access control (RBAC, ABAC, least privilege), input sanitization, rate limiting, cryptographic transport encryption, and audit observability.

```text
API Security Best Practices
│
├── [[API Authentication and Identity Verification|01. Authentication and Identity Verification]]
├── [[API Authorization and Access Control|02. Authorization and Access Control]]
├── [[API Input Validation and Sanitization|03. Input Validation and Sanitization]]
├── [[API Rate Limiting and Traffic Management|04. Rate Limiting and Traffic Management]]
├── [[API Transport Security and Data Encryption|05. Transport Security and Encryption]]
└── [[API Audit Logging and Security Monitoring|06. Audit Logging and Security Monitoring]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[API Authentication and Identity Verification|01. Authentication and Identity Verification]] — Standardized identity protocols (OAuth2, OIDC, mTLS), high-entropy JWT secrets, short-lived tokens, and refresh token rotation.
- 📂 [[API Authorization and Access Control|02. Authorization and Access Control]] — Role-Based Access Control (RBAC), Attribute-Based Access Control (ABAC), scoping, and broken object-level authorization (BOLA) prevention.
- 📂 [[API Input Validation and Sanitization|03. Input Validation and Sanitization]] — Strict schema enforcement, SQL injection, NoSQL injection, SSRF, XML external entity (XXE) defenses, and parameter pollution.
- 📂 [[API Rate Limiting and Traffic Management|04. Rate Limiting and Traffic Management]] — Token bucket, leaky bucket algorithms, per-user/per-IP quotas, distributed Redis rate limiters, and WAF DDoS mitigation.
- 📂 [[API Transport Security and Data Encryption|05. Transport Security and Encryption]] — TLS 1.3 enforcement, HSTS headers, certificate pinning, payload encryption, and envelope encryption with KMS.
- 📂 [[API Audit Logging and Security Monitoring|06. Audit Logging and Security Monitoring]] — Structured SIEM logging, masking sensitive PII/secrets, automated security scanning (SAST/DAST), and rapid incident response.

---

## 🔗 References
- ⬆️ Parent: [[Best Practices]]

