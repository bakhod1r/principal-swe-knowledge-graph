---
title: Cyber Security
tags:
  - cyber-security
  - security-engineering
  - server-security
  - database-security
  - appsec
  - cloud-security
  - cryptography
  - zero-trust
  - principal-swe
parent: "[[Infrastructure & Security]]"
---

# 🛡️ Cyber Security & Defensive Systems Engineering

Comprehensive, production-grade master architecture covering the complete spectrum of server security, database security, application security, API security, enterprise threat defense, cloud infrastructure hardening, defensive cryptography, and AI red teaming across 13 master pillars and 83 specialized subdomains:

- **Security Foundations & Models:** CIA triad, formal security policies (Bell-LaPadula/Biba), OS hardening, container isolation, and Zero Trust Architecture (ZTA).
- **Server Security & Infrastructure Hardening:** CIS benchmarks, SSH ephemeral certificate authorities (Vault/Teleport), host firewalling (`nftables`), File Integrity Monitoring (FIM), auditd/eBPF telemetry, and immutable OS images (Packer).
- **Database Security & Storage Encryption:** Transparent Data Encryption (TDE), application field encryption, Row-Level Security (RLS), AWS IAM DB auth, SQL injection parameterization, dynamic data masking, and immutable WORM backups.
- **OWASP Top 10 Application Security:** Injection flaws (SQLi/Command), Broken authentication, Crypto failures, XXE, Broken Access Control (BOLA/IDOR), Security misconfiguration, XSS, Insecure deserialization, and SSRF.
- **API & Microservice Security:** OAuth 2.1 & OIDC, JWT hardening & JWKS rotation, API Gateway rate limiting, Service Mesh mTLS with SPIFFE/SPIRE, and GraphQL/gRPC security interceptors.
- **Defensive Cryptography & PKI:** Symmetric AES-GCM/ChaCha20, Asymmetric RSA/ECC, Hashes & HMAC, Password hashing (Argon2id), PKI X.509, and Cloud HSMs/KMS.
- **Threat Modeling & Risk Assessment:** STRIDE methodology, PASTA risk-centric framework, Attack trees & MITRE ATT&CK mapping, CVSS v4.0 scoring, and SOC2/ISO 27001 compliance.
- **DevSecOps & Supply Chain Security:** Secure SDLC, SAST/DAST/IAST automated CI gates, SBOM (CycloneDX/SPDX), and signed container provenance (SLSA/Cosign).
- **Attack Types & Red Teaming:** Social engineering & phishing defenses, Ransomware containment, MitM & network spoofing, DDoS mitigation, Web penetration testing, and Purple Teaming.
- **Cloud & Workload Protection:** Cloud IAM least-privilege, CSPM continuous compliance, CWPP runtime defense, WAF rulesets, and Kubernetes security with Falco eBPF.
- **AI Red Teaming & LLM Safety:** Direct/Indirect prompt injection, Multi-turn jailbreaking, Training data poisoning, Model extraction & weight theft, and LLM Guardrails (NeMo, Llama Guard).
- **Incident Response & Digital Forensics:** CSIRT incident lifecycle, Live memory forensics, SIEM log correlation, SOAR automated playbooks, and proactive threat hunting.
- **IAM & Zero Trust Governance:** SSO (SAML 2.0, OIDC), Phishing-resistant WebAuthn/FIDO2, PAM bastions, RBAC vs ABAC vs ReBAC, and CARTA continuous adaptive risk.

```text
Cyber Security
│
├── [[Core Cyber Security Foundations & Security Models|01. Core Cyber Security Foundations & Security Models]]
│   ├── `01. CIA Triad and Formal Security Models`
│   ├── `02. Operating System Hardening and Kernel Security`
│   ├── `03. Network Security Fundamentals and Dpi`
│   ├── `04. Virtualization and Container Isolation`
│   ├── `05. Privileged Access Management and Identity Lifecycle`
│   └── `06. Zero Trust Architecture Fundamentals`
├── [[Server Security & Infrastructure Hardening|02. Server Security & Infrastructure Hardening]]
│   ├── `01. Linux Host Hardening and Cis Benchmarks`
│   ├── `02. SSH Security, Ephemeral Certificates, and Bastions`
│   ├── `03. Host Firewalling, Micro Segmentation, and Port Lockdown`
│   ├── `04. File Integrity Monitoring Fim and Rootkit Detection`
│   ├── `05. Kernel Auditing, Syscall Telemetry, and eBPF Tracing`
│   └── `06. Immutable Server Images and Ephemeral Infrastructure`
├── `03. Database Security & Storage Encryption`
│   ├── `01. Transparent Data Encryption TDE and Storage Encryption`
│   ├── `02. Application Level Field Encryption and Envelope Encryption`
│   ├── `03. Row Level Security RLS and Fine Grained Access Control`
│   ├── `04. Database Authentication, IAM Roles, and Short Lived Tokens`
│   ├── `05. Sql Injection Prevention and Parameterized Query Engines`
│   ├── `06. Dynamic Data Masking and Database Activity Monitoring Dam`
│   └── `07. Database Backup Encryption and Ransomware Immutability`
├── [[OWASP Top 10 & Web Application Hardening|04. OWASP Top 10 & Web Application Hardening]]
│   ├── `01. Injection Flaws SQLi and Command Injection`
│   ├── `02. Broken Authentication and Session Hijacking`
│   ├── `03. Cryptographic Failures and Data Exposure`
│   ├── `04. Xml External Entity XXE Prevention`
│   ├── `05. Broken Access Control and IDOR`
│   ├── `06. Security Misconfiguration and Default Credentials`
│   ├── `07. Cross Site Scripting XSS Mitigation`
│   ├── `08. Insecure Deserialization and Object Injection`
│   ├── `09. Server Side Request Forgery SSRF Defense`
│   └── `10. Insufficient Logging, Monitoring, and Integrity Failures`
├── [[Api & Microservice Security Architecture|05. API & Microservice Security Architecture]]
│   ├── `01. Oauth 2.1 and Openid Connect OIDC Architectures`
│   ├── `02. Jwt Hardening, Cryptographic Signing, and Revocation`
│   ├── `03. Api Gateway Security, Rate Limiting, and Throttling`
│   ├── `04. Microservice Mutual TLS mTLS and SPIFFE Identities`
│   ├── `05. Graphql Security, Query Depth, and Complexity Analysis`
│   └── `06. Grpc Security, Metadata Interceptors, and TLS`
├── [[Defensive Cryptography, PKI & Key Management|06. Defensive Cryptography, PKI & Key Management]]
│   ├── `01. Symmetric Ciphers and Authenticated Encryption`
│   ├── `02. Asymmetric Cryptography and Digital Signatures`
│   ├── `03. Cryptographic Hashes and Message Authentication Codes`
│   ├── `04. Password Hashing and Key Derivation Functions`
│   ├── `05. Public Key Infrastructure PKI and Certificate Management`
│   └── `06. Hardware Security Modules Hsms and Cloud KMS`
├── [[Threat Modeling, Risk Assessment & Attack Trees|07. Threat Modeling, Risk Assessment & Attack Trees]]
│   ├── `01. STRIDE Threat Modeling Framework`
│   ├── `02. PASTA Risk Centric Threat Analysis`
│   ├── `03. Attack Trees and Mitre Att&ck Mapping`
│   ├── `04. Vulnerability Scoring CVSS V4 and DREAD`
│   ├── `05. Security Architecture Review and Threat Invariants`
│   └── `06. Enterprise Security Governance and Compliance`
├── [[Devsecops, Secure SDLC & Supply Chain Hardening|08. DevSecOps, Secure SDLC & Supply Chain Hardening]]
│   ├── `01. Secure SDLC and Shift Left Security Culture`
│   ├── `02. Static Application Security Testing SAST`
│   ├── `03. Dynamic Application Security Testing DAST`
│   ├── `04. Interactive Application Security Testing IAST`
│   ├── `05. Software Bill of Materials SBOM and Dependency Scanning`
│   └── `06. Supply Chain Levels for Software Artifacts SLSA`
├── [[Attack Types, Red Teaming & Penetration Testing|09. Attack Types, Red Teaming & Penetration Testing]]
│   ├── `01. Social Engineering, Phishing, and MFA Fatigue`
│   ├── `02. Ransomware Defense, Containment, and Recovery`
│   ├── `03. Man in the Middle MitM and Network Spoofing`
│   ├── `04. DDoS Mitigation and Traffic Scrubbing Networks`
│   ├── `05. Web Application Penetration Testing Methodologies`
│   └── `06. Offensive Red Teaming and Adversary Emulation`
├── [[Cloud Security, IAM & Workload Protection|10. Cloud Security, IAM & Workload Protection]]
│   ├── `01. Cloud IAM Security and Ephemeral Escalation`
│   ├── `02. Cloud Security Posture Management CSPM`
│   ├── `03. Cloud Workload Protection Platforms CWPP`
│   ├── `04. Web Application Firewall WAF Rulesets`
│   ├── `05. Kubernetes Security and Container Runtime Hardening`
│   └── `06. Cloud Network Perimeter and Privatelink Isolation`
├── [[AI Red Teaming & LLM Safety Engineering|11. AI Red Teaming & LLM Safety Engineering]]
│   ├── `01. Prompt Injection and Indirect Injection Vectors`
│   ├── `02. Jailbreaking and Multi Turn Adversarial Prompts`
│   ├── `03. Training Data Poisoning and Model Backdoors`
│   ├── `04. Adversarial Robustness and Gradient Attacks`
│   ├── `05. Model Extraction, Inversion, and Weight Theft`
│   └── `06. LLM Guardrails, Input Scrubbing, and Content Moderation`
├── [[Incident Response, Digital Forensics & Soc Operations|12. Incident Response, Digital Forensics & SOC Operations]]
│   ├── `01. Computer Security Incident Response CSIRT Lifecycle`
│   ├── `02. Digital Forensics and Evidence Preservation`
│   ├── `03. Security Information and Event Management SIEM`
│   ├── `04. Security Orchestration, Automation, and Response SOAR`
│   ├── `05. Threat Hunting and Indicator of Compromise IoC Analysis`
│   └── `06. Post Incident Blameless Review and Recovery Verification`
└── [[Identity, Access Management & Zero Trust Governance|13. Identity, Access Management & Zero Trust Governance]]
│   ├── `01. Single Sign on SSO and SAML OIDC Federation`
│   ├── `02. Phishing Resistant MFA and Fido2 Webauthn`
│   ├── `03. Privilege Access Management PAM and Bastion Architecture`
│   ├── `04. Access Control Models Rbac, Abac, and Rebac`
│   ├── `05. Service Mesh mTLS and Workload Identity SPIFFE SPIRE`
│   └── `06. Continuous Adaptive Risk and Trust Assessment Carta`
```

---

## 🛡️ Core Knowledge Pillars

### 1. 📂 [[Core Cyber Security Foundations & Security Models|01. Core Cyber Security Foundations & Security Models]]
- 📂 `01. CIA Triad and Formal Security Models` — Confidentiality, Integrity, Availability triad, Bell-LaPadula confidentiality lattice, Biba integrity model, and Clark-Wilson transaction security.
- 📂 `02. Operating System Hardening and Kernel Security` — Linux capability stripping (`cap_drop`), mandatory access control (SELinux/AppArmor), kernel ASLR, seccomp-bpf syscall filtering, and immutable root filesystems.
- 📂 `03. Network Security Fundamentals and Dpi` — Packet filtering, stateful inspection, Deep Packet Inspection (DPI), DMZ architectures, network segmentation, and intrusion detection (IDS/IPS).
- 📂 `04. Virtualization and Container Isolation` — cgroups resource limits, kernel namespaces (PID, NET, MNT), rootless containers, hypervisor-based microVM isolation (Firecracker/gVisor).
- 📂 `05. Privileged Access Management and Identity Lifecycle` — Principle of least privilege, just-in-time access elevation, bastion hosts, break-glass procedures, and automated access offboarding.
- 📂 `06. Zero Trust Architecture Fundamentals` — NIST SP 800-207 framework, 'never trust, always verify', continuous identity and device posture validation, micro-segmentation, and dynamic trust scoring.
### 2. 📂 [[Server Security & Infrastructure Hardening|02. Server Security & Infrastructure Hardening]]
- 📂 `01. Linux Host Hardening and Cis Benchmarks` — Automating CIS Level 1/2 hardening profiles, disabling legacy filesystems/protocols, restricting `/tmp` mount options (`noexec`, `nosuid`, `nodev`), and sysctl kernel hardening.
- 📂 `02. SSH Security, Ephemeral Certificates, and Bastions` — Disabling passwords and root logins, short-lived SSH certificates signed by HashiCorp Vault, session recording via Teleport, and automated brute-force banning (Fail2ban).
- 📂 `03. Host Firewalling, Micro Segmentation, and Port Lockdown` — Zero-trust local egress filtering, blocking unauthorized outbound C2 traffic, default-deny ingress rules with `nftables`, and service-to-service firewall rules.
- 📂 `04. File Integrity Monitoring Fim and Rootkit Detection` — Monitoring critical system binary integrity (`/bin`, `/sbin`, `/etc`), detecting stealth kernel rootkits (chkrootkit, rkhunter), and real-time auditd telemetry.
- 📂 `05. Kernel Auditing, Syscall Telemetry, and eBPF Tracing` — Tracking executive mutations (`execve`), monitoring file access events, streaming structured auditd logs to SIEM, and real-time behavioral tracing via eBPF Tetragon.
- 📂 `06. Immutable Server Images and Ephemeral Infrastructure` — Baking pre-hardened golden AMIs with HashiCorp Packer, eliminating manual server patching in production, and replacing compromised nodes through automated recycling.
### 3. 📂 `03. Database Security & Storage Encryption`
- 📂 `01. Transparent Data Encryption TDE and Storage Encryption` — Encrypting database datafiles and transaction logs at rest (PostgreSQL, MySQL, Oracle TDE), Linux LUKS volume encryption, and envelope encryption with KMS.
- 📂 `02. Application Level Field Encryption and Envelope Encryption` — Encrypting sensitive fields (SSN, credit cards) before sending to DB, deterministic vs randomized encryption (AES-SIV), format-preserving encryption (FPE), and key rotation.
- 📂 `03. Row Level Security RLS and Fine Grained Access Control` — Enforcing multi-tenant isolation policies in PostgreSQL/MySQL RLS, preventing data leaks across tenant queries, and session variable authorization.
- 📂 `04. Database Authentication, IAM Roles, and Short Lived Tokens` — Eliminating static database passwords, rotating dynamic credentials with HashiCorp Vault DB secrets engine, and authenticating via AWS RDS IAM tokens.
- 📂 `05. Sql Injection Prevention and Parameterized Query Engines` — Prepared statements, parameterized SQL ASTs, escaping untrusted input, securing dynamic query builders, and database connection privilege sandboxing.
- 📂 `06. Dynamic Data Masking and Database Activity Monitoring Dam` — Masking sensitive fields for non-privileged support queries, detecting anomalous high-volume exfiltration queries, and immutable PostgreSQL/Aurora audit logs.
- 📂 `07. Database Backup Encryption and Ransomware Immutability` — Encrypting backup dumps with GPG/KMS, storing database snapshots in AWS S3 Object Lock (WORM compliance), and disaster recovery validation.
### 4. 📂 [[OWASP Top 10 & Web Application Hardening|04. OWASP Top 10 & Web Application Hardening]]
- 📂 `01. Injection Flaws SQLi and Command Injection` — Parameterized queries, prepared statements, stored procedures, escaping untrusted input, and preventing remote code execution (RCE).
- 📂 `02. Broken Authentication and Session Hijacking` — Secure cookie attributes (`HttpOnly`, `Secure`, `SameSite=Strict`), session invalidation on logout, brute-force protections, and preventing credential stuffing.
- 📂 `03. Cryptographic Failures and Data Exposure` — Encrypting data at rest and in transit, avoiding deprecated ciphers (MD5, SHA-1, DES), masking PII, and automated secret scanning.
- 📂 `04. Xml External Entity XXE Prevention` — Disabling XML external entity resolution (`DCL_DISABLE_DTD`), XML schema validation, and migrating from XML to JSON.
- 📂 `05. Broken Access Control and IDOR` — Enforcing record-level ownership authorization checks, indirect reference maps, UUIDs over sequential IDs, and denying by default.
- 📂 `06. Security Misconfiguration and Default Credentials` — Disabling default admin accounts, removing sample apps/debug endpoints, configuring HTTP security headers (CSP, HSTS, X-Frame-Options).
- 📂 `07. Cross Site Scripting XSS Mitigation` — Context-aware output encoding, Content Security Policy (CSP nonce/hash), DOMPurify sanitization, and avoiding `dangerouslySetInnerHTML`/`eval`.
- 📂 `08. Insecure Deserialization and Object Injection` — Safe serialization formats (JSON, Protobuf), validating object integrity with digital signatures (HMAC), and avoiding native language serialization (Python pickle, Java Serializable).
- 📂 `09. Server Side Request Forgery SSRF Defense` — Validating/whitelisting outgoing URLs, blocking private IP ranges (RFC 1918), disabling cloud metadata access (IMDSv2 hop limit = 1), and egress proxying.
- 📂 `10. Insufficient Logging, Monitoring, and Integrity Failures` — Detecting active reconnaissance, alerting on authentication spikes, immutable append-only audit trails, and verifying code signing hashes.
### 5. 📂 [[Api & Microservice Security Architecture|05. API & Microservice Security Architecture]]
- 📂 `01. Oauth 2.1 and Openid Connect OIDC Architectures` — Authorization Code Flow with PKCE, eliminating Implicit Flow, token introspection, refresh token rotation, and scoping API permissions.
- 📂 `02. Jwt Hardening, Cryptographic Signing, and Revocation` — Preventing `alg: none` exploits, verifying RS256/EdDSA signatures via JWKS endpoints, token blacklisting via Redis, and short-lived access tokens.
- 📂 `03. Api Gateway Security, Rate Limiting, and Throttling` — Edge authentication offloading, IP reputation filtering, token bucket & leaky bucket rate limiting, and protecting backend services from denial of wallet.
- 📂 `04. Microservice Mutual TLS mTLS and SPIFFE Identities` — Cryptographic zero-trust service authentication, automatic certificate rotation via Envoy sidecars, and enforcing fine-grained service communication policies.
- 📂 `05. Graphql Security, Query Depth, and Complexity Analysis` — Disabling production introspection, enforcing maximum query depth limits, calculating query execution cost before resolution, and preventing batching attacks.
- 📂 `06. Grpc Security, Metadata Interceptors, and TLS` — Enforcing TLS 1.3 for gRPC transport, auth interceptors for token verification, Protobuf schema validation, and channel credentials.
### 6. 📂 [[Defensive Cryptography, PKI & Key Management|06. Defensive Cryptography, PKI & Key Management]]
- 📂 `01. Symmetric Ciphers and Authenticated Encryption` — Galois/Counter Mode (GCM), ensuring 96-bit nonce uniqueness, hardware AES-NI acceleration, and authenticated encryption guarantees.
- 📂 `02. Asymmetric Cryptography and Digital Signatures` — Public/private key pairs, discrete log hardness, Ed25519 high-speed signatures, and preventing signature malleability attacks.
- 📂 `03. Cryptographic Hashes and Message Authentication Codes` — One-way compression, collision resistance, HMAC authentication, timing attack resistance, and length extension attack immunity (SHA-3/BLAKE3).
- 📂 `04. Password Hashing and Key Derivation Functions` — Memory-hard hashing algorithms (Argon2id), salt generation, iteration cost tuning, and preventing GPU/ASIC rainbow table cracking.
- 📂 `05. Public Key Infrastructure PKI and Certificate Management` — Root CAs, Intermediate CAs, automated ACME protocol (Let's Encrypt), OCSP stapling, Certificate Transparency (CT) logs, and mTLS mesh.
- 📂 `06. Hardware Security Modules Hsms and Cloud KMS` — FIPS 140-2 Level 3 physical security, asymmetric key generation inside HSM, Envelope Encryption (Master Key + Data Encryption Key), and automatic key rotation.
### 7. 📂 [[Threat Modeling, Risk Assessment & Attack Trees|07. Threat Modeling, Risk Assessment & Attack Trees]]
- 📂 `01. STRIDE Threat Modeling Framework` — Analyzing data flow diagrams (DFD), identifying trust boundaries, mapping STRIDE categories to architectural components, and defining mitigation controls.
- 📂 `02. PASTA Risk Centric Threat Analysis` — Business objective alignment, technical scope definition, threat intelligence integration, vulnerability analysis, and residual risk quantification.
- 📂 `03. Attack Trees and Mitre Att&ck Mapping` — Decomposing adversary attack vectors into hierarchical tree branches, calculating attack cost/probability, and mapping enterprise telemetry to MITRE Tactics, Techniques, and Procedures (TTPs).
- 📂 `04. Vulnerability Scoring CVSS V4 and DREAD` — Common Vulnerability Scoring System (Base, Threat, Environmental metrics), DREAD (Damage, Reproducibility, Exploitability, Affected users, Discoverability) prioritization.
- 📂 `05. Security Architecture Review and Threat Invariants` — Reviewing high-level design documents, evaluating authentication/authorization boundaries, zero trust network controls, and third-party data exchange risks.
- 📂 `06. Enterprise Security Governance and Compliance` — Information security management systems (ISMS), continuous audit logging, access review workflows, data residency controls, and audit report generation.
### 8. 📂 [[Devsecops, Secure SDLC & Supply Chain Hardening|08. DevSecOps, Secure SDLC & Supply Chain Hardening]]
- 📂 `01. Secure SDLC and Shift Left Security Culture` — Embedding security gates at requirements, design, coding, CI/CD, and production; security champions network, and developer security enablement.
- 📂 `02. Static Application Security Testing SAST` — AST taint analysis, source-to-sink flow tracking (Semgrep, SonarQube), eliminating false positives, and blocking high-severity vulnerabilities on PRs.
- 📂 `03. Dynamic Application Security Testing DAST` — Automated black-box web vulnerability scanning (OWASP ZAP), REST/GraphQL API fuzzing, boundary input injection, and automated authenticated testing.
- 📂 `04. Interactive Application Security Testing IAST` — Bytecode instrumentation inside runtime (JVM, Node.js), correlating static code paths with real-time test execution data for zero false-positive detection.
- 📂 `05. Software Bill of Materials SBOM and Dependency Scanning` — Generating CycloneDX/SPDX manifests, scanning dependencies for CVEs (Trivy, Dependabot), and blocking compromised upstream packages in CI.
- 📂 `06. Supply Chain Levels for Software Artifacts SLSA` — Hermetic reproducible builds, generating in-toto provenance attestations, signing container images with Cosign/Sigstore, and admission control policies (Kyverno).
### 9. 📂 [[Attack Types, Red Teaming & Penetration Testing|09. Attack Types, Red Teaming & Penetration Testing]]
- 📂 `01. Social Engineering, Phishing, and MFA Fatigue` — Credential harvesting, reverse proxy phishing (Evilginx), MFA prompt bombing, FIDO2/WebAuthn phishing-resistant hardware tokens, and employee security awareness.
- 📂 `02. Ransomware Defense, Containment, and Recovery` — Lateral movement prevention, disabling SMBv1/RDP, air-gapped immutable backups (WORM storage), automated isolate-host playbooks, and disaster recovery drills.
- 📂 `03. Man in the Middle MitM and Network Spoofing` — ARP spoofing detection, DHCP snooping, Dynamic ARP Inspection (DAI), DNSSEC validation, and strict HTTPS enforcement with certificate pinning.
- 📂 `04. DDoS Mitigation and Traffic Scrubbing Networks` — Volumetric attacks (SYN flood, UDP amplification), Layer 7 application floods, Anycast BGP routing, cloud scrubbing centers (Cloudflare, AWS Shield), and rate limiting.
- 📂 `05. Web Application Penetration Testing Methodologies` — Reconnaissance, vulnerability enumeration, manual business logic exploitation, privilege escalation, and writing actionable remediation reports.
- 📂 `06. Offensive Red Teaming and Adversary Emulation` — Assumed breach scenarios, lateral movement techniques (Pass-the-Hash, Kerberoasting), C2 infrastructure, and collaborating with Blue Teams (Purple Teaming).
### 10. 📂 [[Cloud Security, IAM & Workload Protection|10. Cloud Security, IAM & Workload Protection]]
- 📂 `01. Cloud IAM Security and Ephemeral Escalation` — Eliminating long-lived access keys, AWS IAM roles with STS AssumeRole, permission boundaries, and detecting IAM privilege escalation paths.
- 📂 `02. Cloud Security Posture Management CSPM` — Automated cloud configuration auditing (AWS Config, Prowler), detecting misconfigured S3 buckets/security groups, and automated remediation lambdas.
- 📂 `03. Cloud Workload Protection Platforms CWPP` — Agent-based and agentless host monitoring, file integrity monitoring (FIM), anomaly detection, and automated container isolation on malicious process execution.
- 📂 `04. Web Application Firewall WAF Rulesets` — OWASP Core Rule Set (CRS), custom rate limiting rules, IP reputation blocking, geo-blocking, and bot mitigation with CAPTCHA/challenge tokens.
- 📂 `05. Kubernetes Security and Container Runtime Hardening` — Pod Security Standards (Restricted), disabling privileged containers, Network Policies, Kyverno/OPA Gatekeeper admission validation, and real-time Falco eBPF alerts.
- 📂 `06. Cloud Network Perimeter and Privatelink Isolation` — Keeping backend traffic strictly within cloud provider backbones, eliminating internet-facing internal APIs via AWS PrivateLink, and egress firewall filtering.
### 11. 📂 [[AI Red Teaming & LLM Safety Engineering|11. AI Red Teaming & LLM Safety Engineering]]
- 📂 `01. Prompt Injection and Indirect Injection Vectors` — Bypassing system instructions, malicious untrusted RAG document payloads, secondary prompt injection via web search, and delimiter defense strategies.
- 📂 `02. Jailbreaking and Multi Turn Adversarial Prompts` — Universal adversarial suffixes (GCG attacks), DAN exploits, multi-turn psychological priming, base64/rot13 obfuscation, and automated red team probing.
- 📂 `03. Training Data Poisoning and Model Backdoors` — Injecting stealthy triggers into pre-training/fine-tuning datasets, backdoor trigger activation, split-view poisoning, and dataset cryptographic hashing/deduplication.
- 📂 `04. Adversarial Robustness and Gradient Attacks` — Fast Gradient Sign Method (FGSM), Projected Gradient Descent (PGD), adversarial token substitutions, and adversarial training defensive fine-tuning.
- 📂 `05. Model Extraction, Inversion, and Weight Theft` — Stealing proprietary model weights via black-box API querying, reconstructing private training samples from LLM outputs, and differential privacy training.
- 📂 `06. LLM Guardrails, Input Scrubbing, and Content Moderation` — Semantic guardrail classifiers, PII masking before inference, hallucination detection, prompt-leakage prevention, and automated tool-call sandboxing.
### 12. 📂 [[Incident Response, Digital Forensics & Soc Operations|12. Incident Response, Digital Forensics & SOC Operations]]
- 📂 `01. Computer Security Incident Response CSIRT Lifecycle` — Establishing incident severity matrices (P0-P3), forming crisis response teams, running tabletop simulation drills, and executive communication runbooks.
- 📂 `02. Digital Forensics and Evidence Preservation` — Live memory acquisition (LiME, Volatility), disk bit-stream imaging (dd/FTK), maintaining legal chain of custody, and forensic timeline analysis.
- 📂 `03. Security Information and Event Management SIEM` — Centralizing security telemetry (Elastic Security, Splunk, Microsoft Sentinel), writing Sigma correlation detection rules, and alert deduplication.
- 📂 `04. Security Orchestration, Automation, and Response SOAR` — Automated incident enrichment, blocking malicious IPs on firewalls via API, revoking compromised user credentials, and triggering containment actions.
- 📂 `05. Threat Hunting and Indicator of Compromise IoC Analysis` — Proactive hypothesis-driven threat hunting, consuming STIX/TAXII threat feeds, matching YARA memory rules, and focusing on Pyramid of Pain TTPs.
- 📂 `06. Post Incident Blameless Review and Recovery Verification` — Writing formal forensic incident reports, calculating financial/operational impact, conducting blameless retrospectives, and tracking mandatory engineering remediations.
### 13. 📂 [[Identity, Access Management & Zero Trust Governance|13. Identity, Access Management & Zero Trust Governance]]
- 📂 `01. Single Sign on SSO and SAML OIDC Federation` — Identity Provider (IdP) vs Service Provider (SP), SAML XML assertion validation, OIDC ID Tokens, user provisioning via SCIM 2.0, and cross-domain federation.
- 📂 `02. Phishing Resistant MFA and Fido2 Webauthn` — Asymmetric public-key authentication, hardware security keys (YubiKey), WebAuthn browser API, passkeys, and eliminating SMS/TOTP phishing vulnerabilities.
- 📂 `03. Privilege Access Management PAM and Bastion Architecture` — Dynamic credential injection, ephemeral SSH certificate authorities (HashiCorp Vault / Teleport), session recording, and dual-authorization approvals.
- 📂 `04. Access Control Models Rbac, Abac, and Rebac` — Comparing RBAC role explosion with ABAC dynamic contextual policies (XACML, OPA Rego), and Google Zanzibar relationship-based access control (ReBAC).
- 📂 `05. Service Mesh mTLS and Workload Identity SPIFFE SPIRE` — Zero-trust service-to-service authentication without shared secrets: cryptographic SPIFFE IDs, short-lived X.509 SVID certificates, and automated Envoy mTLS rotation.
- 📂 `06. Continuous Adaptive Risk and Trust Assessment Carta` — Real-time risk scoring, behavioral anomaly detection, step-up authentication triggers, device health posture checks, and dynamic session termination.

---

## 🔗 Navigation
- ⬆️ Parent: `Principal SWE`
- 🏛️ Software Architecture: `Architecture`
- 💻 Computer Science Foundations: `Computer Science`
- 🚀 Infrastructure & DevOps: `DevOps`

---

## 🗂️ Topics

- [[AI Red Teaming & LLM Safety Engineering]]
- [[Api & Microservice Security Architecture]]
- [[Attack Types, Red Teaming & Penetration Testing]]
- [[Cloud Security, IAM & Workload Protection]]
- [[Core Cyber Security Foundations & Security Models]]
- [[Database Security]]
- [[Defensive Cryptography, PKI & Key Management]]
- [[Devsecops, Secure SDLC & Supply Chain Hardening]]
- [[Identity, Access Management & Zero Trust Governance]]
- [[Incident Response, Digital Forensics & Soc Operations]]
- [[OWASP Top 10 & Web Application Hardening]]
- [[Server Security & Infrastructure Hardening]]
- [[Storage Encryption]]
- [[Threat Modeling, Risk Assessment & Attack Trees]]
