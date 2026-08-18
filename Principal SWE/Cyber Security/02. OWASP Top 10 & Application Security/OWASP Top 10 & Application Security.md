---
title: OWASP Top 10 & Application Security
tags:
  - cyber-security
  - appsec
  - owasp-top-10-and-application-security
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🏛️ OWASP Top 10 & Application Security

Comprehensive application security defense: SQL/Command injection prevention, broken authentication, sensitive data exposure, XML External Entity (XXE), broken access control (BOLA), security misconfigurations, XSS, insecure deserialization, vulnerable dependencies, and logging failures.

```text
OWASP Top 10 & Application Security
│
├── [[Injection Flaws (sqli, Nosqli, and Os Command)|01. Injection Flaws SQLi and Command]]
├── [[Broken Authentication and Session Hijacking|02. Broken Authentication and Session Hijacking]]
├── [[Sensitive Data Exposure and Cryptographic Protection|03. Sensitive Data Exposure and Encryption]]
├── [[Xml External Entity (xxe) Vulnerabilities|04. Xml External Entity XXE Prevention]]
├── [[Broken Access Control and Bola Defenses|05. Broken Access Control and Bola]]
├── [[Security Misconfiguration and Default Hardening|06. Security Misconfiguration]]
├── [[Cross Site Scripting (xss) Defenses|07. Cross Site Scripting XSS Defenses]]
├── [[Insecure Deserialization and Object Injection|08. Insecure Deserialization]]
├── [[Vulnerable Components and Supply Chain Dependencies|09. Vulnerable and Outdated Components]]
└── [[Insufficient Logging and Monitoring Failures|10. Insufficient Logging and Monitoring Failures]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Injection Flaws (sqli, Nosqli, and Os Command)|01. Injection Flaws SQLi and Command]] — Parameterized queries, prepared statements, input escaping, ORM query builders, and preventing shell execution vulnerabilities.
- 📂 [[Broken Authentication and Session Hijacking|02. Broken Authentication and Session Hijacking]] — Credential stuffing defenses, brute-force mitigation, secure session tokens (HttpOnly, Secure, SameSite), and session fixation prevention.
- 📂 [[Sensitive Data Exposure and Cryptographic Protection|03. Sensitive Data Exposure and Encryption]] — Data in transit and at rest encryption, masking PII in logs, TLS 1.3 cipher suites, and preventing memory dumps.
- 📂 [[Xml External Entity (xxe) Vulnerabilities|04. Xml External Entity XXE Prevention]] — Disabling XML external DTD parsing in XML parsers, SOAP security, and safe JSON/Protobuf alternatives.
- 📂 [[Broken Access Control and Bola Defenses|05. Broken Access Control and Bola]] — Direct object reference verification (IDOR/BOLA), multi-tenant isolation, principle of least privilege, and role validation.
- 📂 [[Security Misconfiguration and Default Hardening|06. Security Misconfiguration]] — Eliminating default credentials, disabling debugging endpoints in production, removing unnecessary HTTP headers, and strict CORS.
- 📂 [[Cross Site Scripting (xss) Defenses|07. Cross Site Scripting XSS Defenses]] — Contextual output encoding, Content Security Policy (CSP Level 3), DOMPurify sanitization, and HttpOnly cookies.
- 📂 [[Insecure Deserialization and Object Injection|08. Insecure Deserialization]] — Remote code execution via Java/Python pickle/PHP serialization, type checking, and safe JSON serialization.
- 📂 [[Vulnerable Components and Supply Chain Dependencies|09. Vulnerable and Outdated Components]] — Dependency vulnerability scanning (Snyk, Dependabot), Software Bill of Materials (SBOM), and CVE patching workflows.
- 📂 [[Insufficient Logging and Monitoring Failures|10. Insufficient Logging and Monitoring Failures]] — High-fidelity audit trails, real-time alert thresholds, zero PII logging, and correlation IDs for security incident forensics.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]

