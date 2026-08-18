---
title: Hash Tables (Basic Data Structures)
tags:
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Hash Tables (Basic Data Structures)

Key-value associative mapping, collision resolution, Robin Hood hashing, and tombstone deletions.

```text
Hash Tables (Basic Data Structures)
│
├── [[Hash Function Uniformity (MurmurHash3, xxHash, CityHash)]]
├── [[Hash Table Memory Layout and Bucket Array]]
├── [[Hash Table Insert and Put Operations]]
├── [[Hash Table Lookup and Get Operations]]
├── [[Hash Table Delete and Tombstone Markers]]
├── [[Hash Table Separate Chaining (Linked Lists vs RB-Trees)]]
├── [[Hash Table Open Addressing (Linear Probing Mechanics)]]
├── [[Hash Table Quadratic Probing and Double Hashing]]
├── [[Hash Table Robin Hood Hashing (PSL Equalization)]]
├── [[Hash Table Cuckoo Hashing (Worst-Case O(1) Lookups)]]
├── [[Hash Table Load Factor and Dynamic Incremental Rehashing]]
├── [[Swiss Tables SIMD Control Byte Group Probing (Abseil)]]
├── [[Consistent Hashing and Virtual Nodes on Hash Ring]]
└── [[Cryptographic Hash Indexing and Merkle Trees]]
```

---

## 🗂️ Operations & Topics

- [[Hash Function Uniformity (MurmurHash3, xxHash, CityHash)]] — Mapping arbitrary byte keys into uniformly distributed 64-bit integers with avalanche properties.
- [[Hash Table Memory Layout and Bucket Array]] — Contiguous bucket array layout, hash slot calculation (hash & mask), and cache locality.
- [[Hash Table Insert and Put Operations]] — O(1) expected key-value insertion, collision resolution, and load factor evaluation.
- [[Hash Table Lookup and Get Operations]] — O(1) expected search via bucket probing or chaining list traversal.
- [[Hash Table Delete and Tombstone Markers]] — Marking deleted slots with tombstones in open addressing to prevent broken probe chains.
- [[Hash Table Separate Chaining (Linked Lists vs RB-Trees)]] — Bucket collision lists converting to Red-Black trees (Java 8 HashMap) when chain length > 8.
- [[Hash Table Open Addressing (Linear Probing Mechanics)]] — Sequential contiguous slot probing with cache line spatial locality.
- [[Hash Table Quadratic Probing and Double Hashing]] — Eliminating primary clustering using quadratic offsets (i^2) or secondary hash functions.
- [[Hash Table Robin Hood Hashing (PSL Equalization)]] — Equalizing probe sequence lengths (PSL) by displacing richer keys during insert.
- [[Hash Table Cuckoo Hashing (Worst-Case O(1) Lookups)]] — Two independent hash functions guaranteeing max 2 memory reads for any lookup.
- [[Hash Table Load Factor and Dynamic Incremental Rehashing]] — Doubling bucket capacity and incrementally evacuating keys (Redis/Go map).
- [[Swiss Tables SIMD Control Byte Group Probing (Abseil)]] — 1-byte control metadata matching 16 slots per SSE2 instruction (Rust hashbrown).
- [[Consistent Hashing and Virtual Nodes on Hash Ring]] — 2^32-1 circular hash ring minimizing key migrations during server scaling.
- [[Cryptographic Hash Indexing and Merkle Trees]] — Hierarchical cryptographic hash verification in distributed stores and git.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: [[Data Structures & Algorithms]]

