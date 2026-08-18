---
title: Map Dictionary
tags:
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Map Dictionary

Key-to-value associative dictionaries, iteration randomization, and comma-ok idioms.

```text
Map Dictionary
│
├── [[Associative Map Invariants and Key Uniqueness]]
├── [[Map Put and Value Updating]]
├── [[Map Get and Zero-Value Fallback]]
├── [[Map Delete and Slot Clearance]]
├── [[Map Comma-Ok Pattern (Presence vs Zero Value)]]
├── [[Map Key Hashability and Equality Constraints]]
├── [[Map Iteration Randomization and Hash Seed Security]]
├── [[Ordered Map (RB-Tree) vs Unordered Hash Map]]
├── [[Concurrent Map Reads and Writes (sync.Map, Sharded Locks)]]
└── [[Map Memory Footprint and Bucket Re-Allocation]]
```

---

## 🗂️ Operations & Topics

- [[Associative Map Invariants and Key Uniqueness]] — 1-to-1 key-value mapping axioms and memory layout.
- [[Map Put and Value Updating]] — Inserting new key-value pairs or updating existing values in O(1).
- [[Map Get and Zero-Value Fallback]] — Retrieving value with zero-value fallback on missing keys.
- [[Map Delete and Slot Clearance]] — Removing key-value association and releasing bucket slots.
- [[Map Comma-Ok Pattern (Presence vs Zero Value)]] — Distinguishing between stored zero values and missing keys (val, ok := m[k]).
- [[Map Key Hashability and Equality Constraints]] — Requirements for comparable key types and avoiding floating-point map keys.
- [[Map Iteration Randomization and Hash Seed Security]] — Randomizing map iteration starting point to prevent deterministic order bugs and HashDoS attacks.
- [[Ordered Map (RB-Tree) vs Unordered Hash Map]] — O(log N) sorted key traversal vs O(1) expected hash lookups.
- [[Concurrent Map Reads and Writes (sync.Map, Sharded Locks)]] — Mitigating fatal concurrent map write crashes via read-only fast paths and lock sharding.
- [[Map Memory Footprint and Bucket Re-Allocation]] — Understanding why hash maps never shrink allocated bucket memory after key deletions.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: [[Data Structures & Algorithms]]
- 🎓 Root: [[Principal SWE]]
