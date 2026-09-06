---
title: "Map Dictionary"
tags:
  - review
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# Map Dictionary

## 1. Definition

## 2. Mental Model

## 3. Usage

## 4. Gotchas

---

## 🗺️ Module Architecture & Sub-Domains

```
Map Dictionary/
├── Hash Map/
├── Ordered Tree Map/
├── Concurrent Map/
├── Specialized Memory Maps/
└── Common Map Algorithms & Operations/
```

---

## 📑 Comprehensive Sub-Domain Index

### [[Hash Map]]
  - [[Associative Map Invariants and Key Uniqueness]]
  - [[Fast Insertion and Lookup Patterns (Compute-If-Absent, Upsert)]]
  - [[Key Immutability Invariant and Hash Code Mutation Gotchas]]
  - [[Map Comma-Ok Pattern (Presence vs Zero Value)]]
  - [[Map Contains Key (Membership Testing)]]
  - [[Map Get and Lookup by Key]]
  - [[Map Iteration Randomization and Hash Seed Security]]
  - [[Map Key Hashability and Equality Constraints]]
  - [[Map Key Set and Entry Set Iteration]]
  - [[Map Load Factor and Dynamic Rehash Resizing]]
  - [[Map Memory Footprint and Bucket Re-Allocation]]
  - [[Map Put and Insert Key-Value Pairs]]
  - [[Map Remove and Delete by Key]]

### [[Ordered Tree Map]]
  - [[Bi-Directional Range Queries on Balanced Tree Maps]]
  - [[NavigableMap Operations (SubMap, HeadMap, TailMap Splices)]]
  - [[Ordered Tree Map Architecture (std::map, Java TreeMap)]]
  - [[Ordered Tree Map Put, Get, and Remove Operations]]
  - [[Prefix-Search and Lexicographical Sorting via Tree Maps]]
  - [[Tree Map Range Queries (Range Scans, Min Key, Max Key)]]

### [[Concurrent Map]]
  - [[Concurrent Map Architecture and Lock Striping]]
  - [[Concurrent Map Lock-Free Reads (Go sync.Map, Java ConcurrentHashMap)]]
  - [[Concurrent Map Thread-Safe Put, Get, and CAS Update Operations]]
  - [[Lock Striping and Fine-Grained Segment Locks (Java 7 ConcurrentHashMap)]]
  - [[Read-Copy-Update (RCU) and Left-Right Concurrency in Maps]]

### [[Specialized Memory Maps]]
  - [[EnumMap and Flat Array Direct-Index Maps for Compact Enums]]
  - [[Memory-Mapped Key-Value Store (LMDB B-Tree Memory Architecture)]]
  - [[Trie and Radix Tree Prefix Maps (Compact Key Compression)]]

### [[Common Map Algorithms & Operations]]
  - [[LFU Cache Implementation (Frequency Map + Doubly Linked Lists)]]
  - [[LRU Cache Implementation (Hash Map + Doubly Linked List)]]
  - [[Trie-Based Prefix Autocomplete and Dictionary Lookups]]
  - [[Word Pattern and Isomorphic String Mapping]]

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Curriculum: `Computer Science`
