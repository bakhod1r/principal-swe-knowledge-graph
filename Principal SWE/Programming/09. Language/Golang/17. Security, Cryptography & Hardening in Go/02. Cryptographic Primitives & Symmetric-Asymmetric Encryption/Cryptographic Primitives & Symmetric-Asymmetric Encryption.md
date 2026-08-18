---
title: Cryptographic Primitives & Symmetric-Asymmetric Encryption
tags:
  - golang
  - security
  - principal-swe
parent: "[[Security, Cryptography & Hardening in Go]]"
---

# Cryptographic Primitives & Symmetric-Asymmetric Encryption

Authenticated symmetric encryption (AES-GCM, ChaCha20), asymmetric signatures (Ed25519), KDFs, and Post-Quantum cryptography.

```text
Cryptographic Primitives & Symmetric-Asymmetric Encryption
│
├── [[Authenticated Symmetric Encryption (AES-GCM & ChaCha20-Poly1305)]]
├── [[Asymmetric Signatures & Key Pairs (Ed25519, ECDSA, RSA)]]
├── [[Key Derivation Functions (Argon2, Scrypt, PBKDF2, HKDF)]]
├── [[Cryptographic Randomness (crypto-rand vs math-rand)]]
└── [[Post-Quantum Cryptography in Modern Go (ML-KEM & Kyber)]]
```

---

## 🗂️ Topics

- [[Authenticated Symmetric Encryption (AES-GCM & ChaCha20-Poly1305)]] — AEAD encryption, nonce collision hazards, and Galois/Counter Mode guarantees.
- [[Asymmetric Signatures & Key Pairs (Ed25519, ECDSA, RSA)]] — High-speed digital signatures, curve selection (P-256 vs Ed25519), and key serialization.
- [[Key Derivation Functions (Argon2, Scrypt, PBKDF2, HKDF)]] — Password hashing architectures, work factors, memory cost parameters, and salt generation.
- [[Cryptographic Randomness (crypto-rand vs math-rand)]] — Operating system entropy sources (/dev/urandom, getrandom), non-blocking random byte generation.
- [[Post-Quantum Cryptography in Modern Go (ML-KEM & Kyber)]] — Standard Post-Quantum Key Encapsulation Mechanism in Go cryptography engines.

---

## 🔗 References
- ⬆️ Parent: [[Security, Cryptography & Hardening in Go]]

