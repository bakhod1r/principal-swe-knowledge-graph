---
title: Database Security & Storage Encryption
tags:
  - cyber-security
  - security-engineering
  - database-security-and-storage-encryption
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🛡️ Database Security & Storage Encryption

Enterprise database and storage security: Transparent Data Encryption (TDE), Column-level encryption, Row-Level Security (RLS), Database Access Control & IAM auth, SQL injection mitigation, dynamic data masking, database auditing & compliance, and secure backup encryption.

```text
Database Security & Storage Encryption
│
├── `01. Transparent Data Encryption TDE and Storage Encryption`
├── `02. Application Level Field Encryption and Envelope Encryption`
├── [[Row Level Security (rls), Tenant Isolation, and Virtual Private Databases|03. Row Level Security RLS and Fine Grained Access Control]]
├── [[Database Authentication, Aws IAM DB Auth, and Ephemeral Credentials|04. Database Authentication, IAM Roles, and Short Lived Tokens]]
├── [[Sql Injection Prevention, Ast Parameterization, and Orm Security|05. Sql Injection Prevention and Parameterized Query Engines]]
├── [[Dynamic Data Masking (ddm), Database Activity Monitoring (dam), and Audit Trails|06. Dynamic Data Masking and Database Activity Monitoring Dam]]
└── `07. Database Backup Encryption and Ransomware Immutability`
```

---

## 🗂️ Core Knowledge Domains

- 📂 `01. Transparent Data Encryption TDE and Storage Encryption` — Encrypting database datafiles and transaction logs at rest (PostgreSQL, MySQL, Oracle TDE), Linux LUKS volume encryption, and envelope encryption with KMS.
- 📂 `02. Application Level Field Encryption and Envelope Encryption` — Encrypting sensitive fields (SSN, credit cards) before sending to DB, deterministic vs randomized encryption (AES-SIV), format-preserving encryption (FPE), and key rotation.
- 📂 [[Row Level Security (rls), Tenant Isolation, and Virtual Private Databases|03. Row Level Security RLS and Fine Grained Access Control]] — Enforcing multi-tenant isolation policies in PostgreSQL/MySQL RLS, preventing data leaks across tenant queries, and session variable authorization.
- 📂 [[Database Authentication, Aws IAM DB Auth, and Ephemeral Credentials|04. Database Authentication, IAM Roles, and Short Lived Tokens]] — Eliminating static database passwords, rotating dynamic credentials with HashiCorp Vault DB secrets engine, and authenticating via AWS RDS IAM tokens.
- 📂 [[Sql Injection Prevention, Ast Parameterization, and Orm Security|05. Sql Injection Prevention and Parameterized Query Engines]] — Prepared statements, parameterized SQL ASTs, escaping untrusted input, securing dynamic query builders, and database connection privilege sandboxing.
- 📂 [[Dynamic Data Masking (ddm), Database Activity Monitoring (dam), and Audit Trails|06. Dynamic Data Masking and Database Activity Monitoring Dam]] — Masking sensitive fields for non-privileged support queries, detecting anomalous high-volume exfiltration queries, and immutable PostgreSQL/Aurora audit logs.
- 📂 `07. Database Backup Encryption and Ransomware Immutability` — Encrypting backup dumps with GPG/KMS, storing database snapshots in AWS S3 Object Lock (WORM compliance), and disaster recovery validation.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]

