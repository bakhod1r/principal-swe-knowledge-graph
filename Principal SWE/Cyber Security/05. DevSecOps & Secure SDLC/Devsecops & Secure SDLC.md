---
title: Devsecops & Secure SDLC
tags:
  - cyber-security
  - appsec
  - devsecops-and-secure-sdlc
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🏛️ Devsecops & Secure SDLC

Automated security integration across the CI/CD pipeline: Shift-Left security, SAST static code analysis, DAST dynamic scanning, IAST, Software Bill of Materials (SBOM), container image scanning, and Policy-as-Code (OPA/Gatekeeper).

```text
Devsecops & Secure SDLC
│
├── [[Secure SDLC and Shift Left Security Culture|01. Secure SDLC and Shift Left Culture]]
├── [[Static Application Security Testing (sast)|02. Static Application Security Testing SAST]]
├── [[Dynamic Application Security Testing (dast)|03. Dynamic Application Security Testing DAST]]
├── [[Software Supply Chain Security and SBOM|04. Software Bill of Materials SBOM]]
├── [[Container Image Scanning and Runtime Security|05. Container and Image Security]]
└── [[Policy As Code with OPA and Gatekeeper|06. Policy As Code with OPA and Gatekeeper]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Secure SDLC and Shift Left Security Culture|01. Secure SDLC and Shift Left Culture]] — Embedding security gates into sprint planning, design reviews, automated PR testing, and developer security champions.
- 📂 [[Static Application Security Testing (sast)|02. Static Application Security Testing SAST]] — Automated code scanning (Semgrep, SonarQube), custom rule authoring, abstract syntax tree security rules, and triage.
- 📂 [[Dynamic Application Security Testing (dast)|03. Dynamic Application Security Testing DAST]] — Black-box web vulnerability fuzzing (OWASP ZAP, Burp Suite Enterprise), API fuzzing, and ephemeral staging integration.
- 📂 [[Software Supply Chain Security and SBOM|04. Software Bill of Materials SBOM]] — CycloneDX, SPDX SBOM generation, signing artifacts with Sigstore/Cosign, and SLSA framework compliance.
- 📂 [[Container Image Scanning and Runtime Security|05. Container and Image Security]] — Vulnerability scanning with Trivy/Clair, distroless images, minimal attack surfaces, and Falco runtime eBPF monitoring.
- 📂 [[Policy As Code with OPA and Gatekeeper|06. Policy As Code with OPA and Gatekeeper]] — Rego policy language, admission controllers in Kubernetes, enforcing compliance rules, and automated PR policy checks.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]

