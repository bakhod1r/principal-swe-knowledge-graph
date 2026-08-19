---
title: Cloud Security, IAM & Workload Protection
tags:
  - cyber-security
  - security-engineering
  - cloud-security,-iam-and-workload-protection
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🛡️ Cloud Security, IAM & Workload Protection

Multi-cloud security engineering: Cloud IAM, CSPM, CWPP, WAF rulesets, Kubernetes runtime security (Falco, eBPF), and private VPC networking (AWS PrivateLink).

```text
Cloud Security, IAM & Workload Protection
│
├── [[Cloud IAM Security, Role Based Access, and Ephemeral Privilege Escalation|01. Cloud IAM Security and Ephemeral Escalation]]
├── [[Cloud Security Posture Management (cspm) and Continuous Compliance|02. Cloud Security Posture Management CSPM]]
├── [[Cloud Workload Protection Platforms (cwpp) and Runtime Defense|03. Cloud Workload Protection Platforms CWPP]]
├── [[Web Application Firewall (waf) Rulesets and Traffic Filtering|04. Web Application Firewall WAF Rulesets]]
├── [[Kubernetes Cluster Security, Admission Controllers, and Falco eBPF|05. Kubernetes Security and Container Runtime Hardening]]
└── [[Cloud Network Perimeter, Vpc Endpoints, and Privatelink Isolation|06. Cloud Network Perimeter and Privatelink Isolation]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Cloud IAM Security, Role Based Access, and Ephemeral Privilege Escalation|01. Cloud IAM Security and Ephemeral Escalation]] — Eliminating long-lived access keys, AWS IAM roles with STS AssumeRole, permission boundaries, and detecting IAM privilege escalation paths.
- 📂 [[Cloud Security Posture Management (cspm) and Continuous Compliance|02. Cloud Security Posture Management CSPM]] — Automated cloud configuration auditing (AWS Config, Prowler), detecting misconfigured S3 buckets/security groups, and automated remediation lambdas.
- 📂 [[Cloud Workload Protection Platforms (cwpp) and Runtime Defense|03. Cloud Workload Protection Platforms CWPP]] — Agent-based and agentless host monitoring, file integrity monitoring (FIM), anomaly detection, and automated container isolation on malicious process execution.
- 📂 [[Web Application Firewall (waf) Rulesets and Traffic Filtering|04. Web Application Firewall WAF Rulesets]] — OWASP Core Rule Set (CRS), custom rate limiting rules, IP reputation blocking, geo-blocking, and bot mitigation with CAPTCHA/challenge tokens.
- 📂 [[Kubernetes Cluster Security, Admission Controllers, and Falco eBPF|05. Kubernetes Security and Container Runtime Hardening]] — Pod Security Standards (Restricted), disabling privileged containers, Network Policies, Kyverno/OPA Gatekeeper admission validation, and real-time Falco eBPF alerts.
- 📂 [[Cloud Network Perimeter, Vpc Endpoints, and Privatelink Isolation|06. Cloud Network Perimeter and Privatelink Isolation]] — Keeping backend traffic strictly within cloud provider backbones, eliminating internet-facing internal APIs via AWS PrivateLink, and egress firewall filtering.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]

