---
title: Work-Stealing Deque
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Deque]]"
---

# 📦 Work-Stealing Deque

Multi-threaded task scheduling deque (Chase-Lev).

```text
Work-Stealing Deque
│
├── [[Work-Stealing Deque Architecture (Chase-Lev Algorithm)]]
├── [[Chase-Lev Work-Stealing Deque Lock-Free Algorithm]]
└── [[Fork-Join Frameworks and Work Balancing via Bottom-Push and Top-Steal]]
```

---

## 🗂️ Topics & Implementations

- [[Work-Stealing Deque Architecture (Chase-Lev Algorithm)]] — Worker thread pushes/pops LIFO from bottom; victim threads steal FIFO from top.
- [[Chase-Lev Work-Stealing Deque Lock-Free Algorithm]] — Memory fences and atomic CAS preventing race conditions during concurrent steals.
- [[Fork-Join Frameworks and Work Balancing via Bottom-Push and Top-Steal]] — Load balancing across multicore CPUs in modern runtimes (Go, Java ForkJoinPool).

---

## 🔗 References
- ⬆️ Parent: [[Deque]]
- 📚 Module: `Data Structures`
