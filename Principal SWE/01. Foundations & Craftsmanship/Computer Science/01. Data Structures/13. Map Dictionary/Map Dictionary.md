---
title: Map Dictionary
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# 📦 Map Dictionary

Key-to-value associative dictionaries, iteration randomization, and comma-ok idioms across 3 specialized domains.

```text
Map Dictionary
│
├── 01. Core Invariants & Semantics
│   ├── [[Associative Map Invariants and Key Uniqueness]]
│   ├── [[Ordered Map (RB-Tree) vs Unordered Hash Map]]
│   └── [[Map Key Hashability and Equality Constraints]]
│
├── 02. Key-Value Operations & Lifecycle
│   ├── [[Map Put and Value Updating]]
│   ├── [[Map Get and Zero-Value Fallback]]
│   ├── [[Map Delete and Slot Clearance]]
│   └── [[Map Comma-Ok Pattern (Presence vs Zero Value)]]
│
└── 03. Systems, Memory & Concurrency
    ├── [[Map Memory Footprint and Bucket Re-Allocation]]
    ├── [[Map Iteration Randomization and Hash Seed Security]]
    └── [[Concurrent Map Reads and Writes (sync.Map, Sharded Locks)]]
```

---

## 🗂️ Core Knowledge Domains

### 1. 📂 01. Core Invariants & Semantics
- [[Associative Map Invariants and Key Uniqueness]] — 1-to-1 key-value mapping axioms and memory layout.
- [[Ordered Map (RB-Tree) vs Unordered Hash Map]] — O(log N) sorted key traversal vs O(1) expected hash lookups.
- [[Map Key Hashability and Equality Constraints]] — Requirements for comparable key types and avoiding floating-point map keys.

### 2. 📂 02. Key-Value Operations & Lifecycle
- [[Map Put and Value Updating]] — Inserting new key-value pairs or updating existing values in O(1).
- [[Map Get and Zero-Value Fallback]] — Retrieving value with zero-value fallback on missing keys.
- [[Map Delete and Slot Clearance]] — Removing key-value association and releasing bucket slots.
- [[Map Comma-Ok Pattern (Presence vs Zero Value)]] — Distinguishing between stored zero values and missing keys (val, ok := m[k]).

### 3. 📂 03. Systems, Memory & Concurrency
- [[Map Memory Footprint and Bucket Re-Allocation]] — Understanding why hash maps never shrink allocated bucket memory after key deletions.
- [[Map Iteration Randomization and Hash Seed Security]] — Randomizing map iteration starting point to prevent deterministic order bugs and HashDoS attacks.
- [[Concurrent Map Reads and Writes (sync.Map, Sharded Locks)]] — Mitigating fatal concurrent map write crashes via read-only fast paths and lock sharding.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`


