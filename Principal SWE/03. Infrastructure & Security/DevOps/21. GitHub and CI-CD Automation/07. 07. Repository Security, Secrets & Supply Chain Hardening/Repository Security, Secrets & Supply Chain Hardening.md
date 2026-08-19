---
title: Repository Security, Secrets & Supply Chain Hardening
tags:
  - devops
  - git-and-github
  - version-control
  - repository-security,-secrets-and-supply-chain-hardening
  - principal-swe
parent: "[[GitHub and CI-CD Automation]]"
---

# 🐙 Repository Security, Secrets & Supply Chain Hardening

Software supply chain security: GPG/SSH commit signing, Secret Scanning & Push Protection, Dependabot, CodeQL (SAST), SBOM generation, and SLSA provenance.

```text
Repository Security, Secrets & Supply Chain Hardening
│
├── [[Cryptographic Commit Signing with GPG and SSH Keys|01. 01. Cryptographic Commit Signing GPG and SSH Keys]]
├── [[GitHub Secret Scanning, Push Protection, and Pre-Commit Hooks|02. 02. Secret Scanning and Push Protection Gates]]
├── [[Dependabot Vulnerability Alerts and Automated Security Updates|03. 03. Automated Dependency Vulnerability Management Dependabot]]
├── [[CodeQL Static Application Security Testing (SAST) in GitHub Actions|04. 04. CodeQL Static Analysis and Security Scanning SAST]]
└── [[Artifact Provenance, GitHub Artifact Attestations, and SLSA|05. 05. Supply Chain Attestation and Artifact Provenance]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Cryptographic Commit Signing with GPG and SSH Keys|01. Cryptographic Commit Signing GPG and SSH Keys]] — Generating GPG/SSH signing keys, configuring Git to sign commits (`git commit -S`), verified badge in GitHub, and blocking unsigned commits.
- 📂 [[GitHub Secret Scanning, Push Protection, and Pre-Commit Hooks|02. Secret Scanning and Push Protection Gates]] — Detecting leaked API keys/tokens in real-time, blocking git pushes with Push Protection, and local pre-commit hooks (Gitleaks, TruffleHog).
- 📂 [[Dependabot Vulnerability Alerts and Automated Security Updates|03. Automated Dependency Vulnerability Management Dependabot]] — Configuring `.github/dependabot.yml`, automated security update PRs, grouping dependency PRs, and vulnerability triage workflows.
- 📂 [[CodeQL Static Application Security Testing (SAST) in GitHub Actions|04. CodeQL Static Analysis and Security Scanning SAST]] — Writing custom CodeQL queries, scanning pull requests for CWEs (SQLi, XSS, RCE), and managing security alerts in GitHub Security tab.
- 📂 [[Artifact Provenance, GitHub Artifact Attestations, and SLSA|05. Supply Chain Attestation and Artifact Provenance]] — Signing build artifacts and container images in GitHub Actions using Sigstore, generating in-toto build attestations, and verifying build integrity.

---

## 🔗 References
- ⬆️ Parent: `Git & GitHub Version Control & CI-CD Automation`
- 📚 Module: `DevOps`

