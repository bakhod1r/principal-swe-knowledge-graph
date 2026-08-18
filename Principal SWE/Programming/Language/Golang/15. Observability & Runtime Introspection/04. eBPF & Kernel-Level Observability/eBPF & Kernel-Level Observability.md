---
title: eBPF & Kernel-Level Observability
tags:
  - golang
  - observability
  - principal-swe
parent: "[[Observability & Runtime Introspection]]"
---

# eBPF & Kernel-Level Observability

Zero-code instrumentation, eBPF distributed tracing (Cilium, Pixie), off-CPU profiling, and kernel network inspection.

```text
eBPF & Kernel-Level Observability
│
├── [[eBPF Zero-Code Instrumentation Architecture]]
├── [[eBPF Distributed Tracing with Cilium & Pixie]]
├── [[Continuous Off-CPU Profiling with eBPF]]
└── [[Kernel Network Packet Latency & Drop Analysis]]
```

---

## 🗂️ Topics

- [[eBPF Zero-Code Instrumentation Architecture]] — Attaching Linux eBPF kernel kprobes and user-space uprobes to Go binaries without code changes.
- [[eBPF Distributed Tracing with Cilium & Pixie]] — Automatic protocol parsing (HTTP, gRPC, Kafka) at the Linux socket layer via eBPF.
- [[Continuous Off-CPU Profiling with eBPF]] — Measuring nanoseconds spent sleeping on lock mutexes, epoll waits, and disk I/O.
- [[Kernel Network Packet Latency & Drop Analysis]] — Diagnosing TCP retransmits, SYN queue drops, and connection resets using eBPF TC/XDP filters.

---

## 🔗 References
- ⬆️ Parent: [[Observability & Runtime Introspection]]

