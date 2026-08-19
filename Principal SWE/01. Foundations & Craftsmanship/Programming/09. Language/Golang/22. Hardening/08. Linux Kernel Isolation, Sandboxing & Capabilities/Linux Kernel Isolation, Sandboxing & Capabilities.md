---
title: Linux Kernel Isolation, Sandboxing & Capabilities
tags:
  - golang
  - security
  - principal-swe
parent: "[[Hardening]]"
---

# Linux Kernel Isolation, Sandboxing & Capabilities

Seccomp-BPF syscall filtering, Linux namespaces, POSIX capability dropping, and Landlock LSM sandboxing.

```text
Linux Kernel Isolation, Sandboxing & Capabilities
│
├── [[Linux Seccomp-BPF Syscall Filtering in Go]]
├── [[Linux Namespaces & Containerization from Scratch]]
├── [[POSIX Capability Dropping (libcap in Go)]]
├── [[Landlock LSM Unprivileged Application Sandboxing]]
└── [[chroot and pivot_root Filesystem Jail Mechanics]]
```

---

## 🗂️ Topics

- [[Linux Seccomp-BPF Syscall Filtering in Go]] — Restricting permitted kernel system calls using libseccomp-golang to prevent privilege escalation.
- [[Linux Namespaces & Containerization from Scratch]] — Spawning isolated processes with CLONE_NEWPID, CLONE_NEWNET, CLONE_NEWNS in pure Go.
- [[POSIX Capability Dropping (libcap in Go)]] — Dropping root privileges (cap_set_proc) and running with least privilege in containerized workloads.
- [[Landlock LSM Unprivileged Application Sandboxing]] — Linux 5.13+ Landlock LSM restricting filesystem access directly from Go unprivileged processes.
- [[chroot and pivot_root Filesystem Jail Mechanics]] — Constructing isolated root filesystems for untrusted code execution environments.

---

## 🔗 References
- ⬆️ Parent: `Security, Cryptography & Hardening in Go`

