---
title: Api & Microservice Security Architecture
tags:
  - cyber-security
  - security-engineering
  - api-and-microservice-security-architecture
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🛡️ Api & Microservice Security Architecture

Modern API and microservice security: OAuth 2.1, OpenID Connect (OIDC), JWT security & token signing, API Gateway rate limiting, mTLS zero-trust mesh, GraphQL security (query depth/cost), and gRPC security.

```text
Api & Microservice Security Architecture
│
├── [[Oauth 2.1 Authorization Framework and Openid Connect (oidc)|01. Oauth 2.1 and Openid Connect OIDC Architectures]]
├── [[Json Web Tokens (jwt) Security, Cryptographic Signing, and Revocation|02. Jwt Hardening, Cryptographic Signing, and Revocation]]
├── [[Api Gateway Security, Token Bucket Rate Limiting, and Abuse Prevention|03. Api Gateway Security, Rate Limiting, and Throttling]]
├── [[Microservice Mutual TLS (mtls), Workload Identity, and SPIFFE SPIRE|04. Microservice Mutual TLS mTLS and SPIFFE Identities]]
├── [[Graphql Security: Query Depth Limiting, Complexity Cost, and Introspection|05. Graphql Security, Query Depth, and Complexity Analysis]]
└── [[Grpc Security Architecture, Metadata Interceptors, and TLS Encryption|06. Grpc Security, Metadata Interceptors, and TLS]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Oauth 2.1 Authorization Framework and Openid Connect (oidc)|01. Oauth 2.1 and Openid Connect OIDC Architectures]] — Authorization Code Flow with PKCE, eliminating Implicit Flow, token introspection, refresh token rotation, and scoping API permissions.
- 📂 [[Json Web Tokens (jwt) Security, Cryptographic Signing, and Revocation|02. Jwt Hardening, Cryptographic Signing, and Revocation]] — Preventing `alg: none` exploits, verifying RS256/EdDSA signatures via JWKS endpoints, token blacklisting via Redis, and short-lived access tokens.
- 📂 [[Api Gateway Security, Token Bucket Rate Limiting, and Abuse Prevention|03. Api Gateway Security, Rate Limiting, and Throttling]] — Edge authentication offloading, IP reputation filtering, token bucket & leaky bucket rate limiting, and protecting backend services from denial of wallet.
- 📂 [[Microservice Mutual TLS (mtls), Workload Identity, and SPIFFE SPIRE|04. Microservice Mutual TLS mTLS and SPIFFE Identities]] — Cryptographic zero-trust service authentication, automatic certificate rotation via Envoy sidecars, and enforcing fine-grained service communication policies.
- 📂 [[Graphql Security: Query Depth Limiting, Complexity Cost, and Introspection|05. Graphql Security, Query Depth, and Complexity Analysis]] — Disabling production introspection, enforcing maximum query depth limits, calculating query execution cost before resolution, and preventing batching attacks.
- 📂 [[Grpc Security Architecture, Metadata Interceptors, and TLS Encryption|06. Grpc Security, Metadata Interceptors, and TLS]] — Enforcing TLS 1.3 for gRPC transport, auth interceptors for token verification, Protobuf schema validation, and channel credentials.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]

