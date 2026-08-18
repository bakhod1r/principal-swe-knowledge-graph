---
title: Supply Chain Security, SBOM & Compliance
tags:
  - golang
  - security
  - principal-swe
parent: "[[Security, Cryptography & Hardening in Go]]"
---

# Supply Chain Security, SBOM & Compliance

govulncheck in CI/CD, Cosign/Rekor binary signing, SLSA provenance, FIPS 140-3 BoringCrypto, and CycloneDX SBOM manifests.

```text
Supply Chain Security, SBOM & Compliance
│
├── [[Vulnerability Auditing with govulncheck in CI-CD]]
├── [[Cryptographic Binary Signing with Cosign & Rekor]]
├── [[SLSA Provenance Generation & Attestations]]
├── [[FIPS 140-3 BoringCrypto Enterprise Compliance Mode]]
└── [[Software Bill of Materials (SBOM) Generation (cyclonedx)]]
```

---

## 🗂️ Topics

- [[Vulnerability Auditing with govulncheck in CI-CD]] — Static call-graph scanning detecting actively reachable CVE vulnerabilities in dependencies.
- [[Cryptographic Binary Signing with Cosign & Rekor]] — Keyless container and binary signing with Sigstore transparency logs.
- [[SLSA Provenance Generation & Attestations]] — Generating verifiable build provenance manifests for Go release artifacts.
- [[FIPS 140-3 BoringCrypto Enterprise Compliance Mode]] — Compiling Go binaries with validated cryptographic core modules (CGO_ENABLED=1).
- [[Software Bill of Materials (SBOM) Generation (cyclonedx)]] — Automated dependency tracking and SBOM generation in release pipelines.

---

## 🔗 References
- ⬆️ Parent: [[Security, Cryptography & Hardening in Go]]

