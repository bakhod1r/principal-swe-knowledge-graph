---
title: "Array Memory Alignment, Cache Line Striding, and False Sharing"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Memory Alignment, Cache Line Striding, and False Sharing

## 1. Definition
**Memory Alignment** aligns data structures to hardware word boundaries ($4\text{ B}, 8\text{ B}, 64\text{ B}$).
Sequential iterations over aligned arrays allow CPU hardware prefetchers to load entire 64-byte cache lines, executing subsequent element reads with **zero L1 cache miss stalls**.

---

## 2. Mental Model
```text
64-Byte CPU Cache Line:
[ Int32 (4B) ][ Int32 (4B) ][ Int32 (4B) ] ... [ 16 Integers per Line ]
Reading Item 0 automatically loads Items 1..15 into L1 cache for FREE!

False Sharing Across Cores:
Core 1 writes Arr[0] <── Same 64B Line ──> Core 2 writes Arr[1]
MESI Protocol constantly invalidates cache lines across CPU cores (100x slowdown!).
```

---

## 3. Usage
```cpp
// Aligning memory for SIMD AVX2 processing and padding against false sharing
struct alignas(64) PaddedAtomicCounter {
    uint64_t count;
    uint8_t padding[56]; // Fills remainder of 64-byte cache line
};
```

---

## 4. Gotchas
- **False Sharing in Multi-Threaded Arrays:** Unpadded parallel worker counters sharing adjacent array indices trigger continuous cache invalidation cascades.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

