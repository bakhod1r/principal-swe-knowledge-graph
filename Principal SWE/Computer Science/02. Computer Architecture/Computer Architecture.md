---
title: Computer Architecture
tags:
  - computer-science
  - computer-architecture
  - principal-swe
parent: "[[Computer Science]]"
---

# 🏛️ Computer Architecture (Foundations & Systems Architecture)

Instruction pipelining, out-of-order execution, branch prediction, MESI cache coherence, TLB virtual addressing, NUMA nodes, memory ordering barriers, SIMD vectorization, and mechanical sympathy.

```text
Computer Architecture
│
├── [[Instruction Pipelining and Out of Order Execution|01. Instruction Pipelining and OoO]]
├── [[Branch Prediction Mechanisms|02. Branch Prediction]]
├── [[Cache Coherence and MESI Protocol|03. Cache Coherence and MESI]]
├── [[Memory Hierarchy and Translation Lookaside Buffer|04. Memory Hierarchy and TLB]]
├── [[Non Uniform Memory Access (numa)|05. NUMA Architectures]]
├── [[Memory Ordering and Hardware Memory Barriers|06. Memory Ordering and Barriers]]
├── [[SIMD and Vector Processing|07. SIMD and Vectorization]]
├── [[Mechanical Sympathy and Hardware Consciousness|08. Mechanical Sympathy]]
└── [[Assembly Language and Hardware Instructions|09. Assembly and Low Level Programming]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Instruction Pipelining and Out of Order Execution|01. Instruction Pipelining and OoO]] — Superscalar execution, hazard detection, register renaming (ROB), and speculative branch execution.
- 📂 [[Branch Prediction Mechanisms|02. Branch Prediction]] — Two-level adaptive branch predictors, branch target buffers (BTB), and TAGE conditional prediction.
- 📂 [[Cache Coherence and MESI Protocol|03. Cache Coherence and MESI]] — MESI and MOESI cache line states (Modified, Exclusive, Shared, Invalid), bus snooping, and directory protocols.
- 📂 [[Memory Hierarchy and Translation Lookaside Buffer|04. Memory Hierarchy and TLB]] — L1/L2/L3 caches, TLB address translation caching, multi-level page tables, and HugePages.
- 📂 [[Non Uniform Memory Access (numa)|05. NUMA Architectures]] — Socket interconnects (QPI/UPI), local vs remote RAM access latency, thread pinning, and numactl tuning.
- 📂 [[Memory Ordering and Hardware Memory Barriers|06. Memory Ordering and Barriers]] — Total Store Order (TSO) on x86 vs Weak Memory Ordering on ARM, MFENCE, load-acquire, and store-release.
- 📂 [[SIMD and Vector Processing|07. SIMD and Vectorization]] — Data-parallel AVX-512, AVX2, and ARM NEON registers, autovectorization, and fused multiply-add (FMA).
- 📂 [[Mechanical Sympathy and Hardware Consciousness|08. Mechanical Sympathy]] — Designing software to cooperate with hardware cache lines (64B), false sharing, and branchless coding.
- 📂 [[Assembly Language and Hardware Instructions|09. Assembly and Low Level Programming]] — x86-64 and ARM64 instruction decoding, calling conventions (System V AMD64 ABI), and inline assembly.

---

## 🔗 References
- ⬆️ Parent: [[Computer Science]]

