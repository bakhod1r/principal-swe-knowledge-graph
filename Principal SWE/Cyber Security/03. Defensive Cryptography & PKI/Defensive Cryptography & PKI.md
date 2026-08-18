---
title: Defensive Cryptography & PKI
tags:
  - cyber-security
  - appsec
  - defensive-cryptography-and-pki
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🏛️ Defensive Cryptography & PKI

Applied modern cryptographic engineering: Symmetric ciphers (AES-256-GCM, ChaCha20-Poly1305), asymmetric crypto (RSA-4096, ECC Ed25519), cryptographic hashing, password KDFs (Argon2id, bcrypt), PKI/X.509 certificates, and KMS envelope encryption.

```text
Defensive Cryptography & PKI
│
├── [[Symmetric Ciphers (aes GCM and Chacha20 Poly1305)|01. Symmetric Ciphers AES and Chacha20]]
├── [[Asymmetric Cryptography (rsa and Elliptic Curves)|02. Asymmetric Cryptography RSA and ECC]]
├── [[Cryptographic Hash Functions and Hmac|03. Cryptographic Hashes and Macs]]
├── [[Password Hashing and Key Derivation Functions|04. Password Hashing and Kdfs]]
├── [[Public Key Infrastructure (pki) and X.509 Certificates|05. Public Key Infrastructure PKI and Certificates]]
└── [[Key Management Services (kms) and Hardware Security Modules|06. Key Management Service KMS and Hsms]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Symmetric Ciphers (aes GCM and Chacha20 Poly1305)|01. Symmetric Ciphers AES and Chacha20]] — Authenticated Encryption with Associated Data (AEAD), nonce reuse hazards, block cipher modes (GCM, CBC), and key sizing.
- 📂 [[Asymmetric Cryptography (rsa and Elliptic Curves)|02. Asymmetric Cryptography RSA and ECC]] — Diffie-Hellman ephemeral key exchange (ECDHE), digital signatures (Ed25519), public key encryption, and quantum vulnerability.
- 📂 [[Cryptographic Hash Functions and Hmac|03. Cryptographic Hashes and Macs]] — Collision resistance, SHA-256, SHA-3, BLAKE3, and HMAC integrity message authentication codes.
- 📂 [[Password Hashing and Key Derivation Functions|04. Password Hashing and Kdfs]] — Memory-hard password hashing with Argon2id, bcrypt, PBKDF2, work factor tuning, and salt generation.
- 📂 [[Public Key Infrastructure (pki) and X.509 Certificates|05. Public Key Infrastructure PKI and Certificates]] — Certificate Authorities (CAs), certificate signing requests (CSRs), OCSP stapling, Certificate Transparency (CT), and mTLS.
- 📂 [[Key Management Services (kms) and Hardware Security Modules|06. Key Management Service KMS and Hsms]] — Envelope encryption (DEK/KEK), automatic key rotation, AWS KMS, HashiCorp Vault, and FIPS 140-2/3 HSM hardware.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]
- 🎓 Root: [[Principal SWE]]
