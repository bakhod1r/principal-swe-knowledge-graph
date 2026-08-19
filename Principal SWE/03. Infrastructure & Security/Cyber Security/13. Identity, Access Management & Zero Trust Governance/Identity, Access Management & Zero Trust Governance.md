---
title: Identity, Access Management & Zero Trust Governance
tags:
  - cyber-security
  - security-engineering
  - identity,-access-management-and-zero-trust-governance
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🛡️ Identity, Access Management & Zero Trust Governance

Modern identity architecture: SSO (SAML 2.0, OIDC), Phishing-resistant MFA (FIDO2, WebAuthn), Privileged Access Management (PAM), RBAC vs ABAC vs ReBAC, Service Mesh mTLS with SPIFFE/SPIRE, and Continuous Adaptive Risk (CARTA).

```text
Identity, Access Management & Zero Trust Governance
│
├── [[Single Sign on (sso), SAML 2.0, and Openid Connect (oidc) Federation|01. Single Sign on SSO and SAML OIDC Federation]]
├── [[Phishing Resistant Multi Factor Authentication (mfa) and Fido2 Webauthn|02. Phishing Resistant MFA and Fido2 Webauthn]]
├── [[Privilege Access Management (pam), Session Recording, and Bastions|03. Privilege Access Management PAM and Bastion Architecture]]
├── [[Access Control Models: Role Based (rbac), Attribute Based (abac), and Relationship Based (rebac)|04. Access Control Models Rbac, Abac, and Rebac]]
├── [[Workload Identity, SPIFFE SPIRE Standards, and Service Mesh mTLS|05. Service Mesh mTLS and Workload Identity SPIFFE SPIRE]]
└── [[Continuous Adaptive Risk and Trust Assessment (carta) Architecture|06. Continuous Adaptive Risk and Trust Assessment Carta]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Single Sign on (sso), SAML 2.0, and Openid Connect (oidc) Federation|01. Single Sign on SSO and SAML OIDC Federation]] — Identity Provider (IdP) vs Service Provider (SP), SAML XML assertion validation, OIDC ID Tokens, user provisioning via SCIM 2.0, and cross-domain federation.
- 📂 [[Phishing Resistant Multi Factor Authentication (mfa) and Fido2 Webauthn|02. Phishing Resistant MFA and Fido2 Webauthn]] — Asymmetric public-key authentication, hardware security keys (YubiKey), WebAuthn browser API, passkeys, and eliminating SMS/TOTP phishing vulnerabilities.
- 📂 [[Privilege Access Management (pam), Session Recording, and Bastions|03. Privilege Access Management PAM and Bastion Architecture]] — Dynamic credential injection, ephemeral SSH certificate authorities (HashiCorp Vault / Teleport), session recording, and dual-authorization approvals.
- 📂 [[Access Control Models: Role Based (rbac), Attribute Based (abac), and Relationship Based (rebac)|04. Access Control Models Rbac, Abac, and Rebac]] — Comparing RBAC role explosion with ABAC dynamic contextual policies (XACML, OPA Rego), and Google Zanzibar relationship-based access control (ReBAC).
- 📂 [[Workload Identity, SPIFFE SPIRE Standards, and Service Mesh mTLS|05. Service Mesh mTLS and Workload Identity SPIFFE SPIRE]] — Zero-trust service-to-service authentication without shared secrets: cryptographic SPIFFE IDs, short-lived X.509 SVID certificates, and automated Envoy mTLS rotation.
- 📂 [[Continuous Adaptive Risk and Trust Assessment (carta) Architecture|06. Continuous Adaptive Risk and Trust Assessment Carta]] — Real-time risk scoring, behavioral anomaly detection, step-up authentication triggers, device health posture checks, and dynamic session termination.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]

