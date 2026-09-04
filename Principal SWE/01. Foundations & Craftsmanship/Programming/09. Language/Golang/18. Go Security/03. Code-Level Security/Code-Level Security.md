---
title: Code-Level Security & Secure Coding Standards
tags:
  - golang
  - security
  - principal-swe
parent: "[[Go Security]]"
---

# Code-Level Security & Secure Coding Standards

SAST security scanning (gosec), command injection defense, secret detection (gitleaks), TOCTOU race condition defense, and secure coding standards.

```text
Code-Level Security & Secure Coding Standards
│
├── `Static Application Security Testing (gosec & semgrep-go)`
├── [[Command Injection Defense & exec.Command Sanitization]]
├── [[Insecure Deserialization & XML-JSON Bomb Defense]]
├── [[Regular Expression Denial of Service (ReDoS) Defense (RE2)]]
├── `Hardcoded Secret Detection & Git Pre-Commit Scanning (Gitleaks)`
├── [[Unsafe Pointer Memory Safety Violations & Exploits]]
├── [[Race Condition Exploits & Time-of-Check to Time-of-Use (TOCTOU)]]
├── [[HTTP Header Injection & CRLF Defense]]
├── [[Mass Assignment & Unbounded Request Binding Defense]]
└── `Safe Temporary File Creation & Permission Hardening`
```

---

## 🗂️ Topics

- [[Static Application Security Testing (gosec & semgrep-go)]] — Automated SAST rules scanning Go AST for hardcoded credentials, unsafe integer conversions, and unhandled errors.
- [[Command Injection Defense & exec.Command Sanitization]] — Preventing shell command injection: passing discrete argument slices without sh -c and environment variable scrubbing.
- [[Insecure Deserialization & XML-JSON Bomb Defense]] — Preventing exponential entity expansion and unbounded memory allocations in XML/JSON parsers.
- [[Regular Expression Denial of Service (ReDoS) Defense (RE2)]] — Why Go standard regexp package (RE2 linear O(n) guarantee) is immune to ReDoS, and third-party backtracking engine risks.
- [[Hardcoded Secret Detection & Git Pre-Commit Scanning (Gitleaks)]] — Automated secret detection in pre-commit hooks and CI pipelines to prevent secret leakage into Git history.
- [[Unsafe Pointer Memory Safety Violations & Exploits]] — How improper unsafe.Pointer casting causes arbitrary memory reads/writes and memory corruption vulnerabilities.
- [[Race Condition Exploits & Time-of-Check to Time-of-Use (TOCTOU)]] — Concurrency race hazards leading to security bypasses and file system race conditions.
- [[HTTP Header Injection & CRLF Defense]] — Sanitizing response headers against newline injection (CRLF) and HTTP response splitting attacks.
- [[Mass Assignment & Unbounded Request Binding Defense]] — Preventing attackers from modifying unauthorized fields in incoming JSON request payloads via strict DTOs.
- [[Safe Temporary File Creation & Permission Hardening]] — Preventing insecure /tmp symlink attacks using os.CreateTemp and restrictive 0600 file permissions.

---

## 🔗 References
- ⬆️ Parent: `Security, Cryptography & Hardening in Go`

