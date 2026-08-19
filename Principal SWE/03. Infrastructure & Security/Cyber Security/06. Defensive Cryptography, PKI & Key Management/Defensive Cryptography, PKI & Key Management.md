---
title: Defensive Cryptography, PKI & Key Management
tags:
  - cyber-security
  - security-engineering
  - defensive-cryptography,-pki-and-key-management
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🛡️ Defensive Cryptography, PKI & Key Management

Cryptographic implementations and lifecycle management: AES-GCM, ChaCha20, RSA, ECC Curve25519, Argon2id, HSMs, KMS, X.509 certificates, and Post-Quantum Cryptography.

```text
Defensive Cryptography, PKI & Key Management
│
├── [[Symmetric Ciphers, AES Gcm, Chacha20 Poly1305, and Aead|01. Symmetric Ciphers and Authenticated Encryption]]
├── [[Asymmetric Cryptography, RSA Pss, Ecdsa, and Ed25519|02. Asymmetric Cryptography and Digital Signatures]]
├── [[Cryptographic Hashes (sha 256, Sha 3, Blake3) and Hmac|03. Cryptographic Hashes and Message Authentication Codes]]
├── [[Password Hashing Standards and Key Derivation Functions (argon2id, Scrypt)|04. Password Hashing and Key Derivation Functions]]
├── [[Public Key Infrastructure (pki), X.509, and Certificate Lifecycles|05. Public Key Infrastructure PKI and Certificate Management]]
└── [[Hardware Security Modules (hsm), Cloud Kms, and Envelope Encryption|06. Hardware Security Modules Hsms and Cloud KMS]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Symmetric Ciphers, AES Gcm, Chacha20 Poly1305, and Aead|01. Symmetric Ciphers and Authenticated Encryption]] — Galois/Counter Mode (GCM), ensuring 96-bit nonce uniqueness, hardware AES-NI acceleration, and authenticated encryption guarantees.
- 📂 [[Asymmetric Cryptography, RSA Pss, Ecdsa, and Ed25519|02. Asymmetric Cryptography and Digital Signatures]] — Public/private key pairs, discrete log hardness, Ed25519 high-speed signatures, and preventing signature malleability attacks.
- 📂 [[Cryptographic Hashes (sha 256, Sha 3, Blake3) and Hmac|03. Cryptographic Hashes and Message Authentication Codes]] — One-way compression, collision resistance, HMAC authentication, timing attack resistance, and length extension attack immunity (SHA-3/BLAKE3).
- 📂 [[Password Hashing Standards and Key Derivation Functions (argon2id, Scrypt)|04. Password Hashing and Key Derivation Functions]] — Memory-hard hashing algorithms (Argon2id), salt generation, iteration cost tuning, and preventing GPU/ASIC rainbow table cracking.
- 📂 [[Public Key Infrastructure (pki), X.509, and Certificate Lifecycles|05. Public Key Infrastructure PKI and Certificate Management]] — Root CAs, Intermediate CAs, automated ACME protocol (Let's Encrypt), OCSP stapling, Certificate Transparency (CT) logs, and mTLS mesh.
- 📂 [[Hardware Security Modules (hsm), Cloud Kms, and Envelope Encryption|06. Hardware Security Modules Hsms and Cloud KMS]] — FIPS 140-2 Level 3 physical security, asymmetric key generation inside HSM, Envelope Encryption (Master Key + Data Encryption Key), and automatic key rotation.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]

