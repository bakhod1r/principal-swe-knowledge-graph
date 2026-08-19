---
title: Server Security & Infrastructure Hardening
tags:
  - cyber-security
  - security-engineering
  - server-security-and-infrastructure-hardening
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🛡️ Server Security & Infrastructure Hardening

Host-level server security: Linux kernel hardening, SSH bastion security, SSH certificate authorities, firewall lockdown (iptables/nftables), CIS Benchmarks, rootless operations, eBPF intrusion detection, and immutable OS images.

```text
Server Security & Infrastructure Hardening
│
├── [[Linux Host Hardening, Cis Benchmarks, and Kernel Protection|01. Linux Host Hardening and Cis Benchmarks]]
├── [[SSH Server Security, Ephemeral SSH Certificates, and Bastions|02. SSH Security, Ephemeral Certificates, and Bastions]]
├── [[Host Based Packet Filtering, Nftables, and Network Micro Segmentation|03. Host Firewalling, Micro Segmentation, and Port Lockdown]]
├── [[File Integrity Monitoring (fim), Tripwire, and Rootkit Detection|04. File Integrity Monitoring Fim and Rootkit Detection]]
├── [[Kernel Audit Subsystem (auditd), Syscall Telemetry, and eBPF Tracing|05. Kernel Auditing, Syscall Telemetry, and eBPF Tracing]]
└── [[Immutable Server Infrastructure, Golden Images (packer), and Auto Patching|06. Immutable Server Images and Ephemeral Infrastructure]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Linux Host Hardening, Cis Benchmarks, and Kernel Protection|01. Linux Host Hardening and Cis Benchmarks]] — Automating CIS Level 1/2 hardening profiles, disabling legacy filesystems/protocols, restricting `/tmp` mount options (`noexec`, `nosuid`, `nodev`), and sysctl kernel hardening.
- 📂 [[SSH Server Security, Ephemeral SSH Certificates, and Bastions|02. SSH Security, Ephemeral Certificates, and Bastions]] — Disabling passwords and root logins, short-lived SSH certificates signed by HashiCorp Vault, session recording via Teleport, and automated brute-force banning (Fail2ban).
- 📂 [[Host Based Packet Filtering, Nftables, and Network Micro Segmentation|03. Host Firewalling, Micro Segmentation, and Port Lockdown]] — Zero-trust local egress filtering, blocking unauthorized outbound C2 traffic, default-deny ingress rules with `nftables`, and service-to-service firewall rules.
- 📂 [[File Integrity Monitoring (fim), Tripwire, and Rootkit Detection|04. File Integrity Monitoring Fim and Rootkit Detection]] — Monitoring critical system binary integrity (`/bin`, `/sbin`, `/etc`), detecting stealth kernel rootkits (chkrootkit, rkhunter), and real-time auditd telemetry.
- 📂 [[Kernel Audit Subsystem (auditd), Syscall Telemetry, and eBPF Tracing|05. Kernel Auditing, Syscall Telemetry, and eBPF Tracing]] — Tracking executive mutations (`execve`), monitoring file access events, streaming structured auditd logs to SIEM, and real-time behavioral tracing via eBPF Tetragon.
- 📂 [[Immutable Server Infrastructure, Golden Images (packer), and Auto Patching|06. Immutable Server Images and Ephemeral Infrastructure]] — Baking pre-hardened golden AMIs with HashiCorp Packer, eliminating manual server patching in production, and replacing compromised nodes through automated recycling.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]

