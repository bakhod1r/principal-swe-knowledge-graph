---
title: Maps
tags:
  - golang
  - maps
  - composite-types
  - principal-swe
parent: "[[Composite Types]]"
---

# Maps

Go hash table internals, `hmap`/`bmap` bucket architecture, Swiss Table maps (Go 1.24+), lookup/insert/delete, `sync.Map`, and concurrency safety.

```text
Maps
│
├── [[Map Internals (hmap and bmap)]]
├── [[Swiss Table Map Architecture (Go 1.24+)]]
├── [[SIMD Control Byte Probing (SSE2-NEON)]]
├── [[make() for Maps]]
├── [[Comma-Ok Idiom for Maps]]
├── [[Map Insert, Update, and Lookup]]
├── [[Map delete() Builtin]]
├── [[clear() on Maps (Go 1.21+)]]
├── [[Map Iteration Randomization & Hash-DoS Safety]]
├── [[Map Key Types Requirements (Equality Contract)]]
├── [[Map Key Struct with Unexported Fields & Incomparable Fields]]
├── [[Map Evacuation & Load Factor (6.5 Threshold)]]
├── [[Map Evacuation & Incremental Rehashing Mechanics]]
├── [[Map Concurrency Hazard & Race Detector Instrumentation]]
├── [[sync.Map Architecture & Read-Only - Dirty Promotion]]
├── [[Concurrent-Safe Map Alternatives (orcaman-concurrent-map vs sync.Map)]]
├── [[Map Memory Leaks & Bucket Growth]]
└── [[maps Standard Package (Go 1.21+)]]
```

---

## 🗂️ Topics

- [[Map Internals (hmap and bmap)]] — Buckets of 8 key-value pairs, `tophash` array, overflow buckets, hash seeds, and incremental load-factor evacuation.
- [[Swiss Table Map Architecture (Go 1.24+)]] — New Swiss Table hash map implementation using SIMD probing for ultra-fast lookups.
- [[SIMD Control Byte Probing (SSE2-NEON)]] — Vectorized 16-byte control group checks in modern Go runtime map implementations.
- [[make() for Maps]] — Pre-allocating map bucket capacity (`make(map[K]V, hint)`) to eliminate expensive incremental rehashing.
- [[Comma-Ok Idiom for Maps]] — Distinguishing between stored zero values and missing keys (`val, ok := m[key]`).
- [[Map Insert, Update, and Lookup]] — Key hashing, bucket lookup, memory assignment, and zero-value returns on absent keys.
- [[Map delete() Builtin]] — Deleting keys from hash map via `delete(m, key)` and bucket slot clearing.
- [[clear() on Maps (Go 1.21+)]] — Emptying all keys while retaining allocated bucket memory for reuse.
- [[Map Iteration Randomization & Hash-DoS Safety]] — Why Go uses fastrand hash seeds to protect against algorithmic complexity Hash-DoS attacks.
- [[Map Key Types Requirements (Equality Contract)]] — Legal map key types, comparable constraints, and why slices/funcs cannot be map keys.
- [[Map Key Struct with Unexported Fields & Incomparable Fields]] — Using composite structs as keys, field comparability rules, and private fields.
- [[Map Evacuation & Load Factor (6.5 Threshold)]] — Load factor calculations ($L = \text{count} / 2^B$) and doubling trigger points.
- [[Map Evacuation & Incremental Rehashing Mechanics]] — Step-by-step movement of buckets from `oldbuckets` to new bucket arrays on mutation.
- [[Map Concurrency Hazard & Race Detector Instrumentation]] — Race detector (`-race`) TSAN instrumentation detecting concurrent map writes.
- [[sync.Map Architecture & Read-Only - Dirty Promotion]] — Lock-free `readOnly` atomic map cache, dirty promotion, and high-read concurrency optimizations.
- [[Concurrent-Safe Map Alternatives (orcaman-concurrent-map vs sync.Map)]] — Sharded / striped mutex maps vs `sync.Map` for write-heavy workloads.
- [[Map Memory Leaks & Bucket Growth]] — Maps never shrink bucket allocations after deletions; reclaiming memory via re-creation.
- [[maps Standard Package (Go 1.21+)]] — Standard generic map helpers: `maps.Clone`, `maps.Copy`, `maps.Equal`, `maps.DeleteFunc`.

---

## 🔗 References
- ⬆️ Parent: [[Composite Types]]
- 📚 Module: `Language Basics`
