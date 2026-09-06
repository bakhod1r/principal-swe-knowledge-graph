---
title: Hash Tables
tags:
  - review
  - computer-science
  - data-structures
  - hash-tables
  - principal-swe
parent: "[[Data Structures]]"
---

# 📦 Hash Tables (Associative Topologies & Collision Resolution)

Comprehensive engineering catalog of associative key-value mapping, uniform hash algorithms, bucket memory architecture, collision resolution strategies (separate chaining, open addressing, Robin Hood, Cuckoo, Swiss Tables SIMD), and distributed hash rings across 6 specialized domains.

```text
Hash Tables
│
├── 01. Separate Chaining Hash Table
│   ├── [[Separate Chaining Hash Table]]
│   ├── [[Hash Table Memory Layout and Bucket Array]]
│   ├── [[Hash Function Uniformity (MurmurHash3, xxHash, CityHash)]]
│   ├── [[Hash Table Load Factor and Dynamic Incremental Rehashing]]
│   └── [[Hash Table Separate Chaining (Linked Lists vs RB-Trees)]]
│
├── 02. Open Addressing Hash Table
│   ├── [[Open Addressing Hash Table]]
│   ├── [[Hash Table Open Addressing (Linear Probing Mechanics)]]
│   ├── [[Hash Table Quadratic Probing and Double Hashing]]
│   ├── [[Hash Table Insert and Put Operations]]
│   ├── [[Hash Table Lookup and Get Operations]]
│   └── [[Hash Table Delete and Tombstone Markers]]
│
├── 03. Robin Hood Hash Table
│   ├── [[Robin Hood Hash Table]]
│   └── [[Hash Table Robin Hood Hashing (PSL Equalization)]]
│
├── 04. Cuckoo Hash Table
│   ├── [[Cuckoo Hash Table]]
│   └── [[Hash Table Cuckoo Hashing (Worst-Case O(1) Lookups)]]
│
├── 05. Swiss Table (SIMD Probing)
│   ├── [[Swiss Table]]
│   └── [[Swiss Tables SIMD Control Byte Group Probing (Abseil)]]
│
└── 06. Distributed Hash Ring
    ├── [[Distributed Hash Ring]]
    ├── [[Consistent Hashing and Virtual Nodes on Hash Ring]]
    └── [[Cryptographic Hash Indexing and Merkle Trees]]
```

---

## 🗂️ Concrete Types & Knowledge Domains

### 1. 📂 [[Separate Chaining Hash Table|01. Separate Chaining Hash Table]]
- [[Separate Chaining Hash Table]] — Master topology of bucket array chaining.
- [[Hash Table Memory Layout and Bucket Array]] — Contiguous bucket array layout, hash slot calculation (hash & mask), and cache locality.
- [[Hash Function Uniformity (MurmurHash3, xxHash, CityHash)]] — Mapping arbitrary byte keys into uniformly distributed 64-bit integers with avalanche properties.
- [[Hash Table Load Factor and Dynamic Incremental Rehashing]] — Doubling bucket capacity and incrementally evacuating keys (Redis/Go map).
- [[Hash Table Separate Chaining (Linked Lists vs RB-Trees)]] — Bucket collision lists converting to Red-Black trees (Java 8 HashMap) when chain length > 8.

### 2. 📂 [[Open Addressing Hash Table|02. Open Addressing Hash Table]]
- [[Open Addressing Hash Table]] — Master topology of contiguous array probing.
- [[Hash Table Open Addressing (Linear Probing Mechanics)]] — Sequential contiguous slot probing with cache line spatial locality.
- [[Hash Table Quadratic Probing and Double Hashing]] — Eliminating primary clustering using quadratic offsets (i^2) or secondary hash functions.
- [[Hash Table Insert and Put Operations]] — O(1) expected key-value insertion, collision resolution, and load factor evaluation.
- [[Hash Table Lookup and Get Operations]] — O(1) expected search via bucket probing or chaining list traversal.
- [[Hash Table Delete and Tombstone Markers]] — Marking deleted slots with tombstones in open addressing to prevent broken probe chains.

### 3. 📂 [[Robin Hood Hash Table|03. Robin Hood Hash Table]]
- [[Robin Hood Hash Table]] — Master topology of PSL-equalized hashing.
- [[Hash Table Robin Hood Hashing (PSL Equalization)]] — Equalizing probe sequence lengths (PSL) by displacing richer keys during insert.

### 4. 📂 [[Cuckoo Hash Table|04. Cuckoo Hash Table]]
- [[Cuckoo Hash Table]] — Master topology of dual-hash displacement.
- [[Hash Table Cuckoo Hashing (Worst-Case O(1) Lookups)]] — Two independent hash functions guaranteeing max 2 memory reads for any lookup.

### 5. 📂 [[Swiss Table|05. Swiss Table (SIMD Probing)]]
- [[Swiss Table]] — Master architecture of vectorized control-byte group probing.
- [[Swiss Tables SIMD Control Byte Group Probing (Abseil)]] — 1-byte control metadata matching 16 slots per SSE2 instruction (Rust hashbrown).

### 6. 📂 [[Distributed Hash Ring|06. Distributed Hash Ring]]
- [[Distributed Hash Ring]] — Master architecture of distributed circular hash rings.
- [[Consistent Hashing and Virtual Nodes on Hash Ring]] — 2^32-1 circular hash ring minimizing key migrations during server scaling.
- [[Cryptographic Hash Indexing and Merkle Trees]] — Hierarchical cryptographic hash verification in distributed stores and git.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`



