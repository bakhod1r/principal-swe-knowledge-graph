---
title: "Hash Tables"
tags:

  - computer-science
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# Hash Tables

## 1. Definition

## 2. Mental Model

## 3. Usage

## 4. Gotchas

---

## 🗺️ Module Architecture & Sub-Domains

```
Hash Tables/
├── Separate Chaining Hash Table/
├── Open Addressing Hash Table/
├── Robin Hood Hash Table/
├── Cuckoo Hash Table/
├── Swiss Table (SIMD Probing)/
├── Distributed Hash Ring/
├── Hash Functions & Resizing Architectures/
└── Common Hash Table Algorithms & Operations/
```

---

## 📑 Comprehensive Sub-Domain Index

### [[Separate Chaining Hash Table]]
  - [[Cache-Conscious Chaining (Unrolled Bucket Arrays)]]
  - [[Hash Function Uniformity (MurmurHash3, xxHash, CityHash)]]
  - [[Hash Table Memory Layout and Bucket Array]]
  - [[Load Factor Thresholds and Rehashing Economics in Chaining]]
  - [[Separate Chaining Architecture and Bucket Array Topology]]
  - [[Separate Chaining Deletion and Key Unlinking Operations]]
  - [[Separate Chaining Insertion and Collision Resolution]]
  - [[Separate Chaining Search and Get Operations]]
  - [[Separate Chaining Treeification Threshold (Java 8 Red-Black Tree)]]

### [[Open Addressing Hash Table]]
  - [[Double Hashing Collision Resolution]]
  - [[Linear Probing Collision Resolution and Cluster Formation]]
  - [[Open Addressing Insert and Put Operations]]
  - [[Open Addressing Lookup and Get Operations]]
  - [[Open Addressing and Flat Array Storage Invariants]]
  - [[Primary and Secondary Clustering in Open Addressing]]
  - [[Quadratic Probing Mechanics and Secondary Clustering]]
  - [[Tombstone Garbage Collection and Probing Chain Compaction]]
  - [[Tombstone Markers and Lazy Deletion in Open Addressing]]

### [[Robin Hood Hash Table]]
  - [[Backward-Shift Deletion in Robin Hood Hashing (Eliminating Tombstones)]]
  - [[Probe Sequence Length (PSL) Variance Reduction Proof]]
  - [[Robin Hood Hashing and Probe Sequence Length (PSL)]]
  - [[Robin Hood Search and Probe Count Operations]]

### [[Cuckoo Hash Table]]
  - [[Cuckoo Graph Cycles, Stash Mechanics, and Rehash Triggers]]
  - [[Cuckoo Hash Table Lookup, Insert, and Eviction Kicking Operations]]
  - [[Cuckoo Hashing Two-Way Choice and Power of Two Choices]]
  - [[Cuckoo Hashing and Dual-Hash O(1) Worst-Case Lookups]]

### [[Swiss Table (SIMD Probing)]]
  - [[Control Byte Array and SSE-AVX-512 Mask Matching (_mm_cmpeq_epi8)]]
  - [[H1-H2 Hash Splitting and 16-Way Parallel Probe Groups]]
  - [[Swiss Table 16-Way Parallel Group Search and Insert Operations]]
  - [[Swiss Table Architecture (SIMD Metadata Probing)]]

### [[Distributed Hash Ring]]
  - [[Consistent Hashing Topology and Key Distribution Ring]]
  - [[Consistent Hashing Virtual Nodes and Replication Strategies]]
  - [[Cryptographic Hash Indexing and Merkle Trees]]
  - [[Rendezvous (Highest Random Weight - HRW) Hashing vs Consistent Rings]]

### [[Hash Functions & Resizing Architectures]]
  - [[Cryptographic Hashes vs Hash Flooding Denial-of-Service Attacks (SipHash)]]
  - [[Incremental and Progressive Rehashing Mechanics (Redis Dict Resizing)]]
  - [[Non-Cryptographic Fast Hashes (MurmurHash3, xxHash, CityHash)]]

### [[Common Hash Table Algorithms & Operations]]
  - [[Group Anagrams via Frequency Key Encoding]]
  - [[Longest Consecutive Sequence in O(N) Time using Hash Table]]
  - [[Subarray Sum Equals K (Prefix Sum with Hash Map)]]
  - [[Two-Sum and Complement Lookup in O(N) Time via Hash Table]]

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Curriculum: `Computer Science`
