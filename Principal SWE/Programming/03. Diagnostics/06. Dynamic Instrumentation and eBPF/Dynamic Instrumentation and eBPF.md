---
title: Dynamic Instrumentation and eBPF
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Dynamic Instrumentation and eBPF

eBPF kprobes, uprobes, tracepoints, zero-overhead socket filters, DTrace dynamic probes, and kernel verifier safety.

```text
Dynamic Instrumentation and eBPF
│
├── [[eBPF kprobes, uprobes, and tracepoints Deep Architecture]]
├── [[Zero-Overhead Network Packet Tracing and Socket Filter Probes]]
├── [[Dynamic Tracepoint Injection with DTrace and SystemTap]]
└── [[Security and Safety Guarantees of the eBPF In-Kernel Verifier]]
```

---

## 🗂️ Topics

- [[eBPF kprobes, uprobes, and tracepoints Deep Architecture]] — Attaching non-invasive kernel and user-space probes to running processes without recompilation.
- [[Zero-Overhead Network Packet Tracing and Socket Filter Probes]] — Observing TCP latency, retransmits, and dropped packets directly in the Linux network stack.
- [[Dynamic Tracepoint Injection with DTrace and SystemTap]] — Ad-hoc instrumentation of compiled production binaries without redeploying code.
- [[Security and Safety Guarantees of the eBPF In-Kernel Verifier]] — How the eBPF verifier mathematically guarantees bounded execution time and memory safety.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]
- 🎓 Root: [[Principal SWE]]
