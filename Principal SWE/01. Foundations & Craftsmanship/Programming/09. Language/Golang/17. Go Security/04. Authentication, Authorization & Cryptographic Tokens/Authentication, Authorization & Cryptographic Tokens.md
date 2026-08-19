---
title: Authentication, Authorization & Cryptographic Tokens
tags:
  - golang
  - security
  - principal-swe
parent: "[[Go Security]]"
---

# Authentication, Authorization & Cryptographic Tokens

JWT signature verification pitfalls, PASETO and macaroons, OAuth2/OIDC with PKCE, and RBAC/ABAC authorization engines.

```text
Authentication, Authorization & Cryptographic Tokens
│
├── [[JWT Security Architecture & Signature Verification Pitfalls]]
├── [[PASETO & Macaroons (Modern Token Alternatives)]]
├── [[OAuth2 & OpenID Connect (OIDC) Integration in Go]]
└── [[Role-Based & Attribute-Based Access Control (RBAC, ABAC)]]
```

---

## 🗂️ Topics

- [[JWT Security Architecture & Signature Verification Pitfalls]] — Preventing the none algorithm exploit, key confusion attacks, and token validation.
- [[PASETO & Macaroons (Modern Token Alternatives)]] — Platform-Agnostic Security Tokens and decentralized authorization with contextual caveats.
- [[OAuth2 & OpenID Connect (OIDC) Integration in Go]] — Implementing PKCE flow, state parameter CSRF defense, and identity token validation.
- [[Role-Based & Attribute-Based Access Control (RBAC, ABAC)]] — Building high-speed in-memory permission evaluation engines in Go.

---

## 🔗 References
- ⬆️ Parent: `Security, Cryptography & Hardening in Go`

