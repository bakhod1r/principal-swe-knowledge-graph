---
title: Cyber Security
tags:
  - cyber-security
  - appsec
  - infosec
  - principal-swe
parent: "[[Principal SWE]]"
---

# 🛡️ Cyber Security, Application Defense & AI Red Teaming

Comprehensive, production-grade master architecture covering the full offensive and defensive cybersecurity landscape: operating system and network hardening, OWASP Top 10 web application defenses, applied defensive cryptography (AES-GCM, Ed25519, KMS), STRIDE threat modeling, automated DevSecOps (SAST, DAST, SBOM), offensive exploit analysis, enterprise cloud security posture management (CSPM, WAF), and modern AI Red Teaming for Large Language Models across 8 master pillars and 52 specialized subdomains.

```text
Cyber Security
│
├── [[Core Cyber Security Foundations|01. Core Cyber Security Foundations]]
│   ├── [[CIA Triad and Core Security Models|01. CIA Triad and Security Models]]
│   ├── [[Operating System Security Hardening|02. Operating System Hardening]]
│   ├── [[Network Security Protocols and Architecture|03. Network Security Fundamentals]]
│   ├── [[Virtualization and Container Isolation Boundaries|04. Virtualization and Container Isolation]]
│   ├── [[Security Diagnostics and Network Analysis Tools|05. Security Diagnostics and Tools]]
│   └── [[Capture the Flag (ctf) Methodologies and Skills|06. Capture the Flag CTF Methodologies]]
├── [[OWASP Top 10 & Application Security|02. OWASP Top 10 & Application Security]]
│   ├── [[Injection Flaws (sqli, Nosqli, and Os Command)|01. Injection Flaws SQLi and Command]]
│   ├── [[Broken Authentication and Session Hijacking|02. Broken Authentication and Session Hijacking]]
│   ├── [[Sensitive Data Exposure and Cryptographic Protection|03. Sensitive Data Exposure and Encryption]]
│   ├── [[Xml External Entity (xxe) Vulnerabilities|04. Xml External Entity XXE Prevention]]
│   ├── [[Broken Access Control and Bola Defenses|05. Broken Access Control and Bola]]
│   ├── [[Security Misconfiguration and Default Hardening|06. Security Misconfiguration]]
│   ├── [[Cross Site Scripting (xss) Defenses|07. Cross Site Scripting XSS Defenses]]
│   ├── [[Insecure Deserialization and Object Injection|08. Insecure Deserialization]]
│   ├── [[Vulnerable Components and Supply Chain Dependencies|09. Vulnerable and Outdated Components]]
│   └── [[Insufficient Logging and Monitoring Failures|10. Insufficient Logging and Monitoring Failures]]
├── [[Defensive Cryptography & PKI|03. Defensive Cryptography & PKI]]
│   ├── [[Symmetric Ciphers (aes GCM and Chacha20 Poly1305)|01. Symmetric Ciphers AES and Chacha20]]
│   ├── [[Asymmetric Cryptography (rsa and Elliptic Curves)|02. Asymmetric Cryptography RSA and ECC]]
│   ├── [[Cryptographic Hash Functions and Hmac|03. Cryptographic Hashes and Macs]]
│   ├── [[Password Hashing and Key Derivation Functions|04. Password Hashing and Kdfs]]
│   ├── [[Public Key Infrastructure (pki) and X.509 Certificates|05. Public Key Infrastructure PKI and Certificates]]
│   └── [[Key Management Services (kms) and Hardware Security Modules|06. Key Management Service KMS and Hsms]]
├── [[Threat Modeling & Risk Management|04. Threat Modeling & Risk Management]]
│   ├── [[STRIDE Threat Modeling Framework|01. STRIDE Threat Modeling]]
│   ├── [[PASTA Methodology and Attack Tree Modeling|02. PASTA and Attack Trees]]
│   ├── [[Vulnerability Risk Scoring (cvss V4 and Dread)|03. Vulnerability Scoring CVSS and DREAD]]
│   ├── [[Zero Trust Architecture (zta) Principles|04. Zero Trust Architecture Zta]]
│   ├── [[Enterprise Security Risk Assessment and Compliance|05. Enterprise Security Risk Assessment]]
│   └── [[Security Architecture Reviews and Secure Design|06. Security Architecture Reviews]]
├── [[Devsecops & Secure SDLC|05. DevSecOps & Secure SDLC]]
│   ├── [[Secure SDLC and Shift Left Security Culture|01. Secure SDLC and Shift Left Culture]]
│   ├── [[Static Application Security Testing (sast)|02. Static Application Security Testing SAST]]
│   ├── [[Dynamic Application Security Testing (dast)|03. Dynamic Application Security Testing DAST]]
│   ├── [[Software Supply Chain Security and SBOM|04. Software Bill of Materials SBOM]]
│   ├── [[Container Image Scanning and Runtime Security|05. Container and Image Security]]
│   └── [[Policy As Code with OPA and Gatekeeper|06. Policy As Code with OPA and Gatekeeper]]
├── [[Attack Types & Penetration Testing|06. Attack Types & Penetration Testing]]
│   ├── [[Social Engineering, Phishing, and Bec Defenses|01. Social Engineering and Phishing]]
│   ├── [[Ransomware Defense Architecture and Recovery|02. Ransomware Defenses and Recovery]]
│   ├── [[Man in the Middle (mitm) Attacks and Interception|03. Man in the Middle MitM Attacks]]
│   ├── [[DDoS Mitigation and Volumetric Traffic Scrubbing|04. DDoS Mitigation and Traffic Scrubbing]]
│   ├── [[Memory Corruption and Binary Exploitation|05. Memory Corruption and Exploitation]]
│   └── [[Red Team Penetration Testing Methodologies|06. Red Team Penetration Testing Methodologies]]
├── [[Cloud & Infrastructure Security|07. Cloud & Infrastructure Security]]
│   ├── [[Cloud IAM Security and Least Privilege Enforcement|01. Cloud IAM Security and Least Privilege]]
│   ├── [[Cloud Security Posture Management (cspm)|02. Cloud Security Posture Management CSPM]]
│   ├── [[Web Application Firewall (waf) Rulesets and Tuning|03. Web Application Firewall WAF Rulesets]]
│   ├── [[Cloud Workload Protection (cwpp)|04. Cloud Workload Protection CWPP]]
│   ├── [[Cloud SIEM and Security Log Analytics|05. SIEM and Security Log Analytics]]
│   └── [[Security Incident Response and SOAR Playbooks|06. Security Incident Response and SOAR]]
└── [[AI Red Teaming & LLM Security|08. AI Red Teaming & LLM Security]]
│   ├── [[Prompt Injection and LLM Jailbreak Techniques|01. Prompt Injection and Jailbreaks]]
│   ├── [[Training Data Poisoning and Model Backdoors|02. Training Data Poisoning and Backdoors]]
│   ├── [[Adversarial Attacks and Machine Learning Robustness|03. Adversarial Attacks and Robustness]]
│   ├── [[Model Extraction, Inversion, and Weight Theft|04. Model Extraction and Weight Theft]]
│   ├── [[LLM Guardrails and Input-output Sanitization|05. LLM Guardrails and Input Sanitization]]
│   └── [[AI Red Teaming Frameworks and Automated Evals|06. AI Red Teaming Frameworks and Evals]]
```

---

## 🏛️ Core Knowledge Pillars

### 1. 📂 [[Core Cyber Security Foundations|01. Core Cyber Security Foundations]]
- 📂 [[CIA Triad and Core Security Models|01. CIA Triad and Security Models]] — Formal principles of Confidentiality, Integrity, and Availability; Bell-LaPadula, Biba, and Clark-Wilson integrity models.
- 📂 [[Operating System Security Hardening|02. Operating System Hardening]] — Kernel parameters sysctl hardening, SELinux/AppArmor mandatory access control (MAC), ASLR, DEP/NX bit, and secure boot.
- 📂 [[Network Security Protocols and Architecture|03. Network Security Fundamentals]] — Packet filtering firewalls, stateful inspection, IPsec VPN tunnels, TLS/mTLS, 802.1X network access control, and IDS/IPS.
- 📂 [[Virtualization and Container Isolation Boundaries|04. Virtualization and Container Isolation]] — Hypervisor escape prevention (VT-x), Linux namespaces, cgroups, rootless containers, seccomp profiles, and gVisor sandboxing.
- 📂 [[Security Diagnostics and Network Analysis Tools|05. Security Diagnostics and Tools]] — Packet dissection with Wireshark, port scanning with Nmap, vulnerability assessment with OpenVAS, and PCAP inspection.
- 📂 [[Capture the Flag (ctf) Methodologies and Skills|06. Capture the Flag CTF Methodologies]] — Binary exploitation, reverse engineering, web application exploitation, cryptography cracking, and forensics CTF patterns.
### 2. 📂 [[OWASP Top 10 & Application Security|02. OWASP Top 10 & Application Security]]
- 📂 [[Injection Flaws (sqli, Nosqli, and Os Command)|01. Injection Flaws SQLi and Command]] — Parameterized queries, prepared statements, input escaping, ORM query builders, and preventing shell execution vulnerabilities.
- 📂 [[Broken Authentication and Session Hijacking|02. Broken Authentication and Session Hijacking]] — Credential stuffing defenses, brute-force mitigation, secure session tokens (HttpOnly, Secure, SameSite), and session fixation prevention.
- 📂 [[Sensitive Data Exposure and Cryptographic Protection|03. Sensitive Data Exposure and Encryption]] — Data in transit and at rest encryption, masking PII in logs, TLS 1.3 cipher suites, and preventing memory dumps.
- 📂 [[Xml External Entity (xxe) Vulnerabilities|04. Xml External Entity XXE Prevention]] — Disabling XML external DTD parsing in XML parsers, SOAP security, and safe JSON/Protobuf alternatives.
- 📂 [[Broken Access Control and Bola Defenses|05. Broken Access Control and Bola]] — Direct object reference verification (IDOR/BOLA), multi-tenant isolation, principle of least privilege, and role validation.
- 📂 [[Security Misconfiguration and Default Hardening|06. Security Misconfiguration]] — Eliminating default credentials, disabling debugging endpoints in production, removing unnecessary HTTP headers, and strict CORS.
- 📂 [[Cross Site Scripting (xss) Defenses|07. Cross Site Scripting XSS Defenses]] — Contextual output encoding, Content Security Policy (CSP Level 3), DOMPurify sanitization, and HttpOnly cookies.
- 📂 [[Insecure Deserialization and Object Injection|08. Insecure Deserialization]] — Remote code execution via Java/Python pickle/PHP serialization, type checking, and safe JSON serialization.
- 📂 [[Vulnerable Components and Supply Chain Dependencies|09. Vulnerable and Outdated Components]] — Dependency vulnerability scanning (Snyk, Dependabot), Software Bill of Materials (SBOM), and CVE patching workflows.
- 📂 [[Insufficient Logging and Monitoring Failures|10. Insufficient Logging and Monitoring Failures]] — High-fidelity audit trails, real-time alert thresholds, zero PII logging, and correlation IDs for security incident forensics.
### 3. 📂 [[Defensive Cryptography & PKI|03. Defensive Cryptography & PKI]]
- 📂 [[Symmetric Ciphers (aes GCM and Chacha20 Poly1305)|01. Symmetric Ciphers AES and Chacha20]] — Authenticated Encryption with Associated Data (AEAD), nonce reuse hazards, block cipher modes (GCM, CBC), and key sizing.
- 📂 [[Asymmetric Cryptography (rsa and Elliptic Curves)|02. Asymmetric Cryptography RSA and ECC]] — Diffie-Hellman ephemeral key exchange (ECDHE), digital signatures (Ed25519), public key encryption, and quantum vulnerability.
- 📂 [[Cryptographic Hash Functions and Hmac|03. Cryptographic Hashes and Macs]] — Collision resistance, SHA-256, SHA-3, BLAKE3, and HMAC integrity message authentication codes.
- 📂 [[Password Hashing and Key Derivation Functions|04. Password Hashing and Kdfs]] — Memory-hard password hashing with Argon2id, bcrypt, PBKDF2, work factor tuning, and salt generation.
- 📂 [[Public Key Infrastructure (pki) and X.509 Certificates|05. Public Key Infrastructure PKI and Certificates]] — Certificate Authorities (CAs), certificate signing requests (CSRs), OCSP stapling, Certificate Transparency (CT), and mTLS.
- 📂 [[Key Management Services (kms) and Hardware Security Modules|06. Key Management Service KMS and Hsms]] — Envelope encryption (DEK/KEK), automatic key rotation, AWS KMS, HashiCorp Vault, and FIPS 140-2/3 HSM hardware.
### 4. 📂 [[Threat Modeling & Risk Management|04. Threat Modeling & Risk Management]]
- 📂 [[STRIDE Threat Modeling Framework|01. STRIDE Threat Modeling]] — Spoofing, Tampering, Repudiation, Information Disclosure, Denial of Service, and Elevation of Privilege analysis.
- 📂 [[PASTA Methodology and Attack Tree Modeling|02. PASTA and Attack Trees]] — Process for Attack Simulation and Threat Analysis (PASTA), visual attack tree decomposition, and threat actor profiling.
- 📂 [[Vulnerability Risk Scoring (cvss V4 and Dread)|03. Vulnerability Scoring CVSS and DREAD]] — Calculating Base, Temporal, and Environmental CVSS scores, DREAD qualitative assessment, and risk prioritization.
- 📂 [[Zero Trust Architecture (zta) Principles|04. Zero Trust Architecture Zta]] — Never trust, always verify; continuous authentication, micro-segmentation, identity-aware proxies, and device health verification.
- 📂 [[Enterprise Security Risk Assessment and Compliance|05. Enterprise Security Risk Assessment]] — NIST Cybersecurity Framework (CSF), ISO 27001 ISMS, SOC 2 Type II audit readiness, and threat surface mapping.
- 📂 [[Security Architecture Reviews and Secure Design|06. Security Architecture Reviews]] — Defense-in-depth, fail-safe defaults, complete mediation, least common mechanism, and psychological acceptability.
### 5. 📂 [[Devsecops & Secure SDLC|05. DevSecOps & Secure SDLC]]
- 📂 [[Secure SDLC and Shift Left Security Culture|01. Secure SDLC and Shift Left Culture]] — Embedding security gates into sprint planning, design reviews, automated PR testing, and developer security champions.
- 📂 [[Static Application Security Testing (sast)|02. Static Application Security Testing SAST]] — Automated code scanning (Semgrep, SonarQube), custom rule authoring, abstract syntax tree security rules, and triage.
- 📂 [[Dynamic Application Security Testing (dast)|03. Dynamic Application Security Testing DAST]] — Black-box web vulnerability fuzzing (OWASP ZAP, Burp Suite Enterprise), API fuzzing, and ephemeral staging integration.
- 📂 [[Software Supply Chain Security and SBOM|04. Software Bill of Materials SBOM]] — CycloneDX, SPDX SBOM generation, signing artifacts with Sigstore/Cosign, and SLSA framework compliance.
- 📂 [[Container Image Scanning and Runtime Security|05. Container and Image Security]] — Vulnerability scanning with Trivy/Clair, distroless images, minimal attack surfaces, and Falco runtime eBPF monitoring.
- 📂 [[Policy As Code with OPA and Gatekeeper|06. Policy As Code with OPA and Gatekeeper]] — Rego policy language, admission controllers in Kubernetes, enforcing compliance rules, and automated PR policy checks.
### 6. 📂 [[Attack Types & Penetration Testing|06. Attack Types & Penetration Testing]]
- 📂 [[Social Engineering, Phishing, and Bec Defenses|01. Social Engineering and Phishing]] — Spear phishing, Business Email Compromise (BEC), credential harvesting, DMARC/DKIM/SPF email security, and FIDO2 MFA.
- 📂 [[Ransomware Defense Architecture and Recovery|02. Ransomware Defenses and Recovery]] — Lateral movement prevention, immutable air-gapped backups, endpoint detection and response (EDR), and recovery runbooks.
- 📂 [[Man in the Middle (mitm) Attacks and Interception|03. Man in the Middle MitM Attacks]] — ARP poisoning, DNS spoofing, rogue Wi-Fi access points, SSL stripping, HSTS preloading, and public key pinning.
- 📂 [[DDoS Mitigation and Volumetric Traffic Scrubbing|04. DDoS Mitigation and Traffic Scrubbing]] — SYN floods, UDP amplification, HTTP/2 Rapid Reset attacks, Anycast BGP scrubbing centers, and Cloudflare Magic Transit.
- 📂 [[Memory Corruption and Binary Exploitation|05. Memory Corruption and Exploitation]] — Stack buffer overflows, heap exploitation, Return-Oriented Programming (ROP), stack canaries, and Rust memory safety.
- 📂 [[Red Team Penetration Testing Methodologies|06. Red Team Penetration Testing Methodologies]] — OSINT reconnaissance, initial access, privilege escalation, credential dumping (Mimikatz), C2 frameworks, and debriefing.
### 7. 📂 [[Cloud & Infrastructure Security|07. Cloud & Infrastructure Security]]
- 📂 [[Cloud IAM Security and Least Privilege Enforcement|01. Cloud IAM Security and Least Privilege]] — Automated permission boundary enforcement, temporary STS credentials, eliminating static access keys, and IAM Access Analyzer.
- 📂 [[Cloud Security Posture Management (cspm)|02. Cloud Security Posture Management CSPM]] — Automated misconfiguration scanning (Prowler, ScoutSuite), drift detection, CIS Benchmark compliance, and auto-remediation.
- 📂 [[Web Application Firewall (waf) Rulesets and Tuning|03. Web Application Firewall WAF Rulesets]] — Managed OWASP Core Rule Sets (CRS), custom rate-based rules, bot management, geoblocking, and false positive tuning.
- 📂 [[Cloud Workload Protection (cwpp)|04. Cloud Workload Protection CWPP]] — Agent-based and agentless host security, kernel eBPF integrity monitoring, rootkit detection, and file integrity monitoring (FIM).
- 📂 [[Cloud SIEM and Security Log Analytics|05. SIEM and Security Log Analytics]] — Centralized security data lakes, AWS CloudTrail, VPC Flow Logs, Athena queries, and Elastic/Splunk security rules.
- 📂 [[Security Incident Response and SOAR Playbooks|06. Security Incident Response and SOAR]] — Automated quarantine workflows, forensic disk snapshotting, credential revocation webhooks, and blameless post-mortems.
### 8. 📂 [[AI Red Teaming & LLM Security|08. AI Red Teaming & LLM Security]]
- 📂 [[Prompt Injection and LLM Jailbreak Techniques|01. Prompt Injection and Jailbreaks]] — Direct prompt injection, indirect context injection via RAG/web scraping, DAN jailbreaks, base64 payload obfuscation, and defenses.
- 📂 [[Training Data Poisoning and Model Backdoors|02. Training Data Poisoning and Backdoors]] — Poisoning pre-training datasets, trigger phrases inducing hidden backdoors, and data curation integrity validation.
- 📂 [[Adversarial Attacks and Machine Learning Robustness|03. Adversarial Attacks and Robustness]] — Fast Gradient Sign Method (FGSM), adversarial perturbation attacks, evasion attacks on classifiers, and robust model training.
- 📂 [[Model Extraction, Inversion, and Weight Theft|04. Model Extraction and Weight Theft]] — API query distillation attacks, model inversion extracting private training data, and securing model artifact storage (S3/KMS).
- 📂 [[LLM Guardrails and Input-output Sanitization|05. LLM Guardrails and Input Sanitization]] — NeMo Guardrails, Llama Guard, regex/embedding-based semantic filtering, output validation, and preventing toxic/PII leakage.
- 📂 [[AI Red Teaming Frameworks and Automated Evals|06. AI Red Teaming Frameworks and Evals]] — Automated adversarial red teaming pipelines (Garak, PyRIT), benchmark evaluation suites, and responsible disclosure.

---

## 🔗 Navigation
- ⬆️ Parent: [[Principal SWE]]

