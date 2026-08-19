---
title: Devsecops, Secure SDLC & Supply Chain Hardening
tags:
  - cyber-security
  - security-engineering
  - devsecops,-secure-sdlc-and-supply-chain-hardening
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🛡️ Devsecops, Secure SDLC & Supply Chain Hardening

Automated security in CI/CD: Shift-Left security, SAST, DAST, IAST, SBOM (CycloneDX/SPDX), dependency vulnerability scanning, SLSA provenance, and signed container supply chains.

```text
Devsecops, Secure SDLC & Supply Chain Hardening
│
├── [[Secure Software Development Life Cycle (s Sdlc) and Shift Left Culture|01. Secure SDLC and Shift Left Security Culture]]
├── [[Static Application Security Testing (sast) Integration and Rule Tuning|02. Static Application Security Testing SAST]]
├── [[Dynamic Application Security Testing (dast) and Api Fuzzing|03. Dynamic Application Security Testing DAST]]
├── [[Interactive Application Security Testing (iast) and Runtime Agents|04. Interactive Application Security Testing IAST]]
├── [[Software Bill of Materials (sbom) and Open Source Vulnerability Management|05. Software Bill of Materials SBOM and Dependency Scanning]]
└── [[Supply Chain Security, SLSA Framework, and Cosign Container Signing|06. Supply Chain Levels for Software Artifacts SLSA]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Secure Software Development Life Cycle (s Sdlc) and Shift Left Culture|01. Secure SDLC and Shift Left Security Culture]] — Embedding security gates at requirements, design, coding, CI/CD, and production; security champions network, and developer security enablement.
- 📂 [[Static Application Security Testing (sast) Integration and Rule Tuning|02. Static Application Security Testing SAST]] — AST taint analysis, source-to-sink flow tracking (Semgrep, SonarQube), eliminating false positives, and blocking high-severity vulnerabilities on PRs.
- 📂 [[Dynamic Application Security Testing (dast) and Api Fuzzing|03. Dynamic Application Security Testing DAST]] — Automated black-box web vulnerability scanning (OWASP ZAP), REST/GraphQL API fuzzing, boundary input injection, and automated authenticated testing.
- 📂 [[Interactive Application Security Testing (iast) and Runtime Agents|04. Interactive Application Security Testing IAST]] — Bytecode instrumentation inside runtime (JVM, Node.js), correlating static code paths with real-time test execution data for zero false-positive detection.
- 📂 [[Software Bill of Materials (sbom) and Open Source Vulnerability Management|05. Software Bill of Materials SBOM and Dependency Scanning]] — Generating CycloneDX/SPDX manifests, scanning dependencies for CVEs (Trivy, Dependabot), and blocking compromised upstream packages in CI.
- 📂 [[Supply Chain Security, SLSA Framework, and Cosign Container Signing|06. Supply Chain Levels for Software Artifacts SLSA]] — Hermetic reproducible builds, generating in-toto provenance attestations, signing container images with Cosign/Sigstore, and admission control policies (Kyverno).

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]

