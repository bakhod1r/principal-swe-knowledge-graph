---
title: Compiler Optimizations & PGO
tags:
  - golang
  - performance
  - principal-swe
parent: "[[Performance Engineering & Profiling]]"
---

# Compiler Optimizations & PGO

Profile-Guided Optimization (PGO), CI/CD build pipelines, inlining heuristics, and dead code pruning.

```text
Compiler Optimizations & PGO
│
├── [[Profile-Guided Optimization (PGO) in Go]]
├── [[Automated PGO Pipeline in CI-CD]]
├── [[Function Inlining Heuristics & go:noinline Pragma]]
└── [[Compiler Dead Code Elimination & Branch Pruning]]
```

---

## 🗂️ Topics

- [[Profile-Guided Optimization (PGO) in Go]] — Feeding production .pgo profiles into go build for branch prediction and inlining gains.
- [[Automated PGO Pipeline in CI-CD]] — Automated collection of Kubernetes CPU profiles and auto-injecting into build artifacts.
- [[Function Inlining Heuristics & go:noinline Pragma]] — Understanding inlining budget calculation and selectively preventing inlining.
- [[Compiler Dead Code Elimination & Branch Pruning]] — How the Go compiler removes unreached code branches during SSA compilation.

---

## 🔗 References
- ⬆️ Parent: [[Performance Engineering & Profiling]]
- 🎓 Root: [[Principal SWE]]
