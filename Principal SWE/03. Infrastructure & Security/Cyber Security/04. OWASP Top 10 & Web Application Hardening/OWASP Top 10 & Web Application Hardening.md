---
title: OWASP Top 10 & Web Application Hardening
tags:
  - cyber-security
  - security-engineering
  - owasp-top-10-and-web-application-hardening
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🛡️ OWASP Top 10 & Web Application Hardening

Application layer threat defense: SQLi, Command Injection, Broken Auth, Crypto Failures, XXE, Broken Access Control (BOLA/IDOR), Security Misconfiguration, XSS, Insecure Deserialization, and SSRF.

```text
OWASP Top 10 & Web Application Hardening
│
├── [[Injection Flaws, Sql Injection (sqli), and Command Injection Defense|01. Injection Flaws SQLi and Command Injection]]
├── [[Broken Authentication, Session Fixation, and Token Hijacking|02. Broken Authentication and Session Hijacking]]
├── [[Cryptographic Failures and Sensitive Data Exposure Prevention|03. Cryptographic Failures and Data Exposure]]
├── [[Xml External Entity (xxe) Injection and Parser Hardening|04. Xml External Entity XXE Prevention]]
├── [[Broken Access Control, Bola, and Insecure Direct Object References (idor)|05. Broken Access Control and IDOR]]
├── [[Security Misconfiguration, Default Credentials, and Hardening|06. Security Misconfiguration and Default Credentials]]
├── [[Cross Site Scripting (xss) Prevention (stored, Reflected, Dom Based)|07. Cross Site Scripting XSS Mitigation]]
├── [[Insecure Deserialization and Object Injection Vulnerabilities|08. Insecure Deserialization and Object Injection]]
├── [[Server Side Request Forgery (ssrf) and Cloud Metadata Protection|09. Server Side Request Forgery SSRF Defense]]
└── [[Insufficient Logging, Monitoring, and Software Integrity Failures|10. Insufficient Logging, Monitoring, and Integrity Failures]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Injection Flaws, Sql Injection (sqli), and Command Injection Defense|01. Injection Flaws SQLi and Command Injection]] — Parameterized queries, prepared statements, stored procedures, escaping untrusted input, and preventing remote code execution (RCE).
- 📂 [[Broken Authentication, Session Fixation, and Token Hijacking|02. Broken Authentication and Session Hijacking]] — Secure cookie attributes (`HttpOnly`, `Secure`, `SameSite=Strict`), session invalidation on logout, brute-force protections, and preventing credential stuffing.
- 📂 [[Cryptographic Failures and Sensitive Data Exposure Prevention|03. Cryptographic Failures and Data Exposure]] — Encrypting data at rest and in transit, avoiding deprecated ciphers (MD5, SHA-1, DES), masking PII, and automated secret scanning.
- 📂 [[Xml External Entity (xxe) Injection and Parser Hardening|04. Xml External Entity XXE Prevention]] — Disabling XML external entity resolution (`DCL_DISABLE_DTD`), XML schema validation, and migrating from XML to JSON.
- 📂 [[Broken Access Control, Bola, and Insecure Direct Object References (idor)|05. Broken Access Control and IDOR]] — Enforcing record-level ownership authorization checks, indirect reference maps, UUIDs over sequential IDs, and denying by default.
- 📂 [[Security Misconfiguration, Default Credentials, and Hardening|06. Security Misconfiguration and Default Credentials]] — Disabling default admin accounts, removing sample apps/debug endpoints, configuring HTTP security headers (CSP, HSTS, X-Frame-Options).
- 📂 [[Cross Site Scripting (xss) Prevention (stored, Reflected, Dom Based)|07. Cross Site Scripting XSS Mitigation]] — Context-aware output encoding, Content Security Policy (CSP nonce/hash), DOMPurify sanitization, and avoiding `dangerouslySetInnerHTML`/`eval`.
- 📂 [[Insecure Deserialization and Object Injection Vulnerabilities|08. Insecure Deserialization and Object Injection]] — Safe serialization formats (JSON, Protobuf), validating object integrity with digital signatures (HMAC), and avoiding native language serialization (Python pickle, Java Serializable).
- 📂 [[Server Side Request Forgery (ssrf) and Cloud Metadata Protection|09. Server Side Request Forgery SSRF Defense]] — Validating/whitelisting outgoing URLs, blocking private IP ranges (RFC 1918), disabling cloud metadata access (IMDSv2 hop limit = 1), and egress proxying.
- 📂 [[Insufficient Logging, Monitoring, and Software Integrity Failures|10. Insufficient Logging, Monitoring, and Integrity Failures]] — Detecting active reconnaissance, alerting on authentication spikes, immutable append-only audit trails, and verifying code signing hashes.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]

