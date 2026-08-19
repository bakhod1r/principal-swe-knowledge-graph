---
title: Maps
parent: "[[Composite Types]]"
---

- [[Map Iteration Randomization & Hash-DoS Safety]] — Why Go uses fastrand hash seeds to protect against algorithmic complexity Hash-DoS attacks.

- [[Map Concurrency Hazard & Race Detector Instrumentation]] — Why concurrent map read/write causes fatal runtime crash concurrent map writes.

- [[Map Key Types Requirements (Equality Contract)]] — Legal map key types, comparable constraints, and why slices/funcs cannot be map keys.

---
title: Maps
tags:
  - golang
  - maps
  - principal-swe
parent: "`Composite Types`"
---

# Maps

Go hash table internals, hmap/bmap bucket architecture, lookup/insert/delete, comma-ok idiom, and concurrency hazards.

```text
Maps
│
├── [[Map Internals (hmap and bmap)]]
├── [[make() for Maps]]
├── [[Comma-Ok Idiom for Maps]]
├── [[Map Insert, Update, and Lookup]]
├── [[Map delete() Builtin]]
├── [[clear() on Maps (Go 1.21+)]]
├── [[Iterating Maps & Randomization]]
├── [[Map Key Constraints & Hashability]]
├── [[Map Concurrency & Fatal Crashes]]
├── [[Map Memory Leaks & Bucket Growth]]
└── [[maps Standard Package (Go 1.21+)]]
```

---

## 🗂️ Topics

- [[Map Internals (hmap and bmap)]] — Buckets of 8 key-value pairs, tophash array, overflow buckets, hash seeds, and incremental load-factor evacuation.
- [[make() for Maps]] — Pre-allocating map bucket capacity (make(map[K]V, hint)) to eliminate expensive incremental rehashing.
- [[Comma-Ok Idiom for Maps]] — Distinguishing between stored zero values and missing keys (val, ok := m[key]).
- [[Map Insert, Update, and Lookup]] — Key hashing, bucket lookup, memory assignment, and zero-value returns on absent keys.
- [[Map delete() Builtin]] — Deleting keys from hash map via delete(m, key) and bucket slot clearing.
- [[clear() on Maps (Go 1.21+)]] — Emptying all keys while retaining allocated bucket memory for reuse.
- [[Iterating Maps & Randomization]] — Randomized hash seed iteration start point and sorting keys pattern.
- [[Map Key Constraints & Hashability]] — Comparable types as map keys, avoiding float keys, unhashable slice/map/func constraints.
- [[Map Concurrency & Fatal Crashes]] — Concurrent map read/write fatal runtime crash, sync.RWMutex guarding, and sync.Map.
- [[Map Memory Leaks & Bucket Growth]] — Maps never shrink bucket allocations after deletions, reclaiming memory via re-creation.
- [[maps Standard Package (Go 1.21+)]] — Standard generic map helpers: maps.Clone, maps.Copy, maps.Equal, maps.DeleteFunc.
- [[Map Concurrency Hazard & Race Detector Instrumentation]]
- [[Map Evacuation & Load Factor (6.5 Threshold)]]
- [[Map Iteration Randomization & Hash-DoS Safety]]
- [[Map Key Types Requirements (Equality Contract)]]
- [[SIMD Control Byte Probing (SSE2-NEON)]]
- [[Swiss Table Map Architecture (Go 1.24+)]]

---

## 🔗 References
- ⬆️ Parent: `Composite Types`

