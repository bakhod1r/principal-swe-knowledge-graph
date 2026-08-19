---
title: Primitives & Concurrency Performance
parent: "[[Go Performance Engineering]]"
---

- [[Stack vs Heap Allocation Latency & Throughput]] — Comparing sub-nanosecond bump-allocator stack allocation with TCMalloc mcache heap allocation.

- [[Stack Frame Reallocation & copystack Overhead]] — Measuring CPU cycle cost of resizing stack from 2KB to 4KB/8KB and pointer fixup table traversals.

- [[Stack Inlining & Leaf Function Stack Preservation]] — How inlining keeps variables in the caller stack frame, eliminating heap escape allocations.

- [[Stack Splitting Preambles & morestack Elimination via nosplit]] — The cost of runtime.morestack preamble checks on tight loops and eliminating via //go:nosplit.

- [[Pointer Indirection & CPU Memory Latency (Cache Misses)]] — Measuring the CPU penalty of dereferencing non-contiguous pointers through L1/L2/L3 caches.

- [[Pointer Chasing vs Value Locality in Large Data Structures]] — Benchmarking slices of pointers []*T vs slices of contiguous structs []T for cache line utilization.

- [[Escape Analysis Flow Graphs & Heap Escape Penalties]] — How taking a pointer &v and escaping function scope triggers GC heap allocation and write barrier tracking.

- [[Unsafe Pointer Arithmetic vs Safe Indexing Benchmarks]] — Measuring raw unsafe.Pointer offset traversal speed vs compiler-checked slice indexing.

- [[Pointer Pinning (runtime.Pinner) Overhead in Cgo Hot-Paths]] — Pinning Go memory addresses for foreign C code without GC relocation overhead.

---
title: Primitives & Concurrency Performance
tags:
  - golang
  - performance
  - principal-swe
parent: "`Performance Engineering & Profiling`"
---

# Primitives & Concurrency Performance

Performance characteristics, memory retention hazards, and microbenchmarks for Slices, Maps, Channels, and Goroutines.

```text
Primitives & Concurrency Performance
│
├── [[Slice Capacity Pre-Allocation & Reallocation Cost]]
├── [[In-Place Slice Filtering & Zero-Alloc Compaction]]
├── [[Slice Sub-Slicing Memory Retention & Pinning Hazards]]
├── [[Map Bucket Pre-Allocation & Evacuation Avoidance]]
├── [[Map Bucket Memory Reclamation & Shrinking Workarounds]]
├── [[Swiss Table Map SIMD Optimization (Go 1.24+)]]
├── `sync.Map vs Partitioned Sharded Map Benchmarks`
├── `Buffered vs Unbuffered Channel Performance Profiles`
├── `Channel Batching vs Single Message Transfer Overhead`
├── `Lock-Free SPSC Queues vs Channel Benchmarks`
├── `Goroutine Spawning Overhead vs Bounded Pool Tuning`
├── `Goroutine Stack Copying Penalty & Pre-Sizing`
└── `GOMAXPROCS Over-Subscription & Context Switch Thrashing`
```

---

## 🗂️ Topics

- [[Slice Capacity Pre-Allocation & Reallocation Cost]] — Measuring allocation latency of append() reallocation vs make([]T, 0, cap) pre-allocation in tight loops.
- [[In-Place Slice Filtering & Zero-Alloc Compaction]] — Zero-allocation slice filtering using two-pointer in-place sub-slicing (b := s[:0]) and avoiding leaks.
- [[Slice Sub-Slicing Memory Retention & Pinning Hazards]] — How sub-slicing a tiny piece of a huge backing array keeps the entire memory block pinned in GC heap.
- [[Map Bucket Pre-Allocation & Evacuation Avoidance]] — Pre-sizing maps with make(map[K]V, hint) to avoid expensive 6.5 load factor incremental rehashing.
- [[Map Bucket Memory Reclamation & Shrinking Workarounds]] — Why Go maps never shrink allocated bucket memory after deletions and how to reclaim memory by re-creating.
- [[Swiss Table Map SIMD Optimization (Go 1.24+)]] — Benchmarking legacy bucket-based hmap against Go 1.24+ Swiss Tables with SSE2/NEON SIMD control byte probing.
- `sync.Map vs Partitioned Sharded Map Benchmarks` — Comparing sync.Map read-heavy performance against custom partitioned sharded mutex maps under write-heavy loads.
- `Buffered vs Unbuffered Channel Performance Profiles` — Measuring throughput and latency differences between synchronous stack-copy rendezvous vs ring-buffer lock queuing.
- `Channel Batching vs Single Message Transfer Overhead` — Reducing context switching and hchan mutex contention by transferring slice batches across channels.
- `Lock-Free SPSC Queues vs Channel Benchmarks` — Measuring sub-microsecond latency gains of lock-free ring buffers over channel mutexes in low-latency systems.
- `Goroutine Spawning Overhead vs Bounded Pool Tuning` — Benchmarking 100k unbounded goroutine allocations against pre-warmed worker pools under peak load.
- `Goroutine Stack Copying Penalty & Pre-Sizing` — Measuring copystack CPU pauses during sudden deep recursive calls and mitigating with stack pre-warming.
- `GOMAXPROCS Over-Subscription & Context Switch Thrashing` — Diagnosing thread thrashing, CPU cache invalidation, and CFS quota starvation under high goroutine concurrency.

---

## 🔗 References
- ⬆️ Parent: `Performance Engineering & Profiling`

