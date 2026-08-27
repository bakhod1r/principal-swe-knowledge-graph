---
title: Goroutine Memory Lifecycle
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Go Concurrency]]"
---

# Goroutine Memory Lifecycle

Goroutine memory structures, 2KB initial stack allocation, copystack growth, GC shrinking, stack split preambles, and gfget/gfput allocation pools.

```text
Goroutine Memory Lifecycle
│
├── [[Initial 2KB Stack Allocation (stack.go)]]
├── [[Contiguous Stack Growth (copystack)]]
├── [[Stack Shrinking during GC]]
├── [[Stack Splitting Preambles (morestack & nosplit)]]
└── [[Goroutine Allocation Pool (gfget & gfput)]]
```

---

## 🗂️ Topics

- [[Initial 2KB Stack Allocation (stack.go)]] — How the Go runtime allocates minimal contiguous stacks for goroutines.
- [[Contiguous Stack Growth (copystack)]] — Stack expansion up to 1GB and pointer fixup translation during memory reallocation.
- [[Stack Shrinking during GC]] — Shrinking underutilized stacks during garbage collection cycles to reclaim memory.
- [[Stack Splitting Preambles (morestack & nosplit)]] — Compiler-inserted stack boundary checks and //go:nosplit leaf function directives.
- [[Goroutine Allocation Pool (gfget & gfput)]] — Reusing dead goroutine memory structures from scheduler free lists.

---

## 🔗 References
- ⬆️ Parent: `Concurrency & Synchronization`

