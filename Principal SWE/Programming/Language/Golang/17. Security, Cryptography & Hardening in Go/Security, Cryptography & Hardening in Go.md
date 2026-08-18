---
title: Security, Cryptography & Hardening in Go
tags:
  - golang
  - security
  - principal-swe
parent: "[[Golang]]"
---

# 🛡️ Security, Cryptography & Hardening in Go

Enterprise cloud security, cryptography, and systems hardening in Go: TLS 1.3/mTLS, symmetric/asymmetric encryption, side-channel attack defenses, memory zeroing, token architectures (JWT, PASETO), OWASP Top 10 mitigation, and SLSA supply chain security.

```text
Security, Cryptography & Hardening in Go
│
├── [[Transport Layer Security & Certificates (crypto-tls, crypto-x509)|01. Transport Layer Security & Certificates (crypto-tls, crypto-x509)]]
│   ├── [[TLS 1.3 Architecture & Zero-RTT Sessions (crypto-tls)]]
│   ├── [[Mutual TLS (mTLS) & Client Certificate Verification]]
│   ├── [[X.509 Certificate Generation, Parsing & Revocation (CRL, OCSP)]]
│   ├── [[Certificate Rotation Without Downtime (GetCertificate & Dynamic Reload)]]
│   └── [[Custom TLS KeyLogWriter for Network Forensics]]
├── [[Cryptographic Primitives & Symmetric-Asymmetric Encryption|02. Cryptographic Primitives & Symmetric-Asymmetric Encryption]]
│   ├── [[Authenticated Symmetric Encryption (AES-GCM & ChaCha20-Poly1305)]]
│   ├── [[Asymmetric Signatures & Key Pairs (Ed25519, ECDSA, RSA)]]
│   ├── [[Key Derivation Functions (Argon2, Scrypt, PBKDF2, HKDF)]]
│   ├── [[Cryptographic Randomness (crypto-rand vs math-rand)]]
│   └── [[Post-Quantum Cryptography in Modern Go (ML-KEM & Kyber)]]
├── [[Side-Channel Attacks & Memory Hardening|03. Side-Channel Attacks & Memory Hardening]]
│   ├── [[Constant-Time Cryptographic Operations (crypto-subtle)]]
│   ├── [[Sensitive Secret Erasure & Memory Zeroing (memclr)]]
│   ├── [[Preventing Secret Leakage in Heap Dumps & Logs]]
│   └── [[Memory Locking (unix.Mlock) to Prevent Swap Spills]]
├── [[Authentication, Authorization & Cryptographic Tokens|04. Authentication, Authorization & Cryptographic Tokens]]
│   ├── [[JWT Security Architecture & Signature Verification Pitfalls]]
│   ├── [[PASETO & Macaroons (Modern Token Alternatives)]]
│   ├── [[OAuth2 & OpenID Connect (OIDC) Integration in Go]]
│   └── [[Role-Based & Attribute-Based Access Control (RBAC, ABAC)]]
├── [[Application Hardening & OWASP Top 10 Defense|05. Application Hardening & OWASP Top 10 Defense]]
│   ├── [[SQL Injection Defense & Parameterized Queries]]
│   ├── [[Server-Side Request Forgery (SSRF) Defense & IP Pinning]]
│   ├── [[Path Traversal & Zip Slip Prevention (os.Root)]]
│   ├── [[Denial-of-Service Defense (Slowloris, Body Limits, Timeouts)]]
│   └── [[Cross-Site Scripting (XSS) & Template Escaping (html-template)]]
├── [[Supply Chain Security, SBOM & Compliance|06. Supply Chain Security, SBOM & Compliance]]
└── [[Code-Level Security & Secure Coding Standards|07. Code-Level Security & Secure Coding Standards]]
│   ├── [[Vulnerability Auditing with govulncheck in CI-CD]]
│   ├── [[Cryptographic Binary Signing with Cosign & Rekor]]
│   ├── [[SLSA Provenance Generation & Attestations]]
│   ├── [[FIPS 140-3 BoringCrypto Enterprise Compliance Mode]]
│   └── [[Software Bill of Materials (SBOM) Generation (cyclonedx)]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Transport Layer Security & Certificates (crypto-tls, crypto-x509)|01. Transport Layer Security & Certificates (crypto-tls, crypto-x509)]]
- [[TLS 1.3 Architecture & Zero-RTT Sessions (crypto-tls)]] — TLS 1.3 handshake, cipher suites, session tickets, and ALPN negotiation.
- [[Mutual TLS (mTLS) & Client Certificate Verification]] — Zero-trust service-to-service authentication with custom ClientAuth and RootCAs pools.
- [[X.509 Certificate Generation, Parsing & Revocation (CRL, OCSP)]] — Parsing PEM/DER certificates, building in-memory CA chains, and OCSP stapling.
- [[Certificate Rotation Without Downtime (GetCertificate & Dynamic Reload)]] — Dynamically reloading TLS certificates on file change without restarting HTTP/gRPC servers.
- [[Custom TLS KeyLogWriter for Network Forensics]] — Exporting TLS pre-master secrets for Wireshark inspection during security debugging.
### 2. 📂 [[Cryptographic Primitives & Symmetric-Asymmetric Encryption|02. Cryptographic Primitives & Symmetric-Asymmetric Encryption]]
- [[Authenticated Symmetric Encryption (AES-GCM & ChaCha20-Poly1305)]] — AEAD encryption, nonce collision hazards, and Galois/Counter Mode guarantees.
- [[Asymmetric Signatures & Key Pairs (Ed25519, ECDSA, RSA)]] — High-speed digital signatures, curve selection (P-256 vs Ed25519), and key serialization.
- [[Key Derivation Functions (Argon2, Scrypt, PBKDF2, HKDF)]] — Password hashing architectures, work factors, memory cost parameters, and salt generation.
- [[Cryptographic Randomness (crypto-rand vs math-rand)]] — Operating system entropy sources (/dev/urandom, getrandom), non-blocking random byte generation.
- [[Post-Quantum Cryptography in Modern Go (ML-KEM & Kyber)]] — Standard Post-Quantum Key Encapsulation Mechanism in Go cryptography engines.
### 3. 📂 [[Side-Channel Attacks & Memory Hardening|03. Side-Channel Attacks & Memory Hardening]]
- [[Constant-Time Cryptographic Operations (crypto-subtle)]] — Eliminating CPU timing attack vulnerabilities using subtle.ConstantTimeCompare.
- [[Sensitive Secret Erasure & Memory Zeroing (memclr)]] — Zeroing byte slices containing private keys/passwords before GC deallocation.
- [[Preventing Secret Leakage in Heap Dumps & Logs]] — Implementing fmt.Stringer and slog.LogValuer masking to redact sensitive fields.
- [[Memory Locking (unix.Mlock) to Prevent Swap Spills]] — Pinning cryptographic keys in RAM using unix.Mlock to prevent swapping to disk.
### 4. 📂 [[Authentication, Authorization & Cryptographic Tokens|04. Authentication, Authorization & Cryptographic Tokens]]
- [[JWT Security Architecture & Signature Verification Pitfalls]] — Preventing the none algorithm exploit, key confusion attacks, and token validation.
- [[PASETO & Macaroons (Modern Token Alternatives)]] — Platform-Agnostic Security Tokens and decentralized authorization with contextual caveats.
- [[OAuth2 & OpenID Connect (OIDC) Integration in Go]] — Implementing PKCE flow, state parameter CSRF defense, and identity token validation.
- [[Role-Based & Attribute-Based Access Control (RBAC, ABAC)]] — Building high-speed in-memory permission evaluation engines in Go.
### 5. 📂 [[Application Hardening & OWASP Top 10 Defense|05. Application Hardening & OWASP Top 10 Defense]]
- [[SQL Injection Defense & Parameterized Queries]] — Strict enforcement of parameterized queries in database/sql and pgx.
- [[Server-Side Request Forgery (SSRF) Defense & IP Pinning]] — Custom http.Transport dialer validating against private IP ranges (RFC 1918) and DNS rebinding.
- [[Path Traversal & Zip Slip Prevention (os.Root)]] — Go 1.24+ os.Root directory sandboxing and path cleaning.
- [[Denial-of-Service Defense (Slowloris, Body Limits, Timeouts)]] — Configuring ReadTimeout, WriteTimeout, IdleTimeout, and http.MaxBytesReader.
- [[Cross-Site Scripting (XSS) & Template Escaping (html-template)]] — Context-aware auto-escaping rules in html/template vs text/template.
### 6. 📂 [[Supply Chain Security, SBOM & Compliance|06. Supply Chain Security, SBOM & Compliance]]
- [[Vulnerability Auditing with govulncheck in CI-CD]] — Static call-graph scanning detecting actively reachable CVE vulnerabilities in dependencies.
- [[Cryptographic Binary Signing with Cosign & Rekor]] — Keyless container and binary signing with Sigstore transparency logs.
- [[SLSA Provenance Generation & Attestations]] — Generating verifiable build provenance manifests for Go release artifacts.
- [[FIPS 140-3 BoringCrypto Enterprise Compliance Mode]] — Compiling Go binaries with validated cryptographic core modules (CGO_ENABLED=1).
- [[Software Bill of Materials (SBOM) Generation (cyclonedx)]] — Automated dependency tracking and SBOM generation in release pipelines.

### 7. 📂 [[Code-Level Security & Secure Coding Standards|07. Code-Level Security & Secure Coding Standards]]
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

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
