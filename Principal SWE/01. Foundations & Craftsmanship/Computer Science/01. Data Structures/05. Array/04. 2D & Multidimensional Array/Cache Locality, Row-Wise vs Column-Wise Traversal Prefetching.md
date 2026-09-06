---
title: "Cache Locality, Row-Wise vs Column-Wise Traversal Prefetching"
tags:
  - review
  - computer-science
  - data-structures
  - matrix
  - principal-swe
parent: "[[2D & Multidimensional Array]]"
---

# Cache Locality, Row-Wise vs Column-Wise Traversal Prefetching

## 1. Definition
Modern CPUs transfer data between main memory (DRAM) and CPU caches in 64-byte chunks known as **Cache Lines**. When accessing element `A[i][j]`, the hardware prefetcher loads neighboring memory addresses into L1 cache ahead of time.

- **Row-Wise Traversal (in Row-Major)**: Accesses contiguous addresses $k, k+4, k+8, \dots$. Yields 1 cache miss every 16 accesses for 32-bit floats.
- **Column-Wise Traversal (in Row-Major)**: Accesses non-contiguous addresses $k, k+N, k+2N, \dots$. Stride exceeds 64 bytes when $N \ge 16$, causing **up to a 10x-50x runtime slowdown**.

## 2. Benchmark Comparison Pattern

```c
// Row-wise Traversal: FAST (L1 Cache Streaming Hit)
for (int i = 0; i < ROWS; i++) {
    for (int j = 0; j < COLS; j++) {
        sum += matrix[i * COLS + j]; // sequential memory addresses
    }
}

// Column-wise Traversal: SLOW (Catastrophic Cache Stride Penalty)
for (int j = 0; j < COLS; j++) {
    for (int i = 0; i < ROWS; i++) {
        sum += matrix[i * COLS + j]; // stride-COLS jumps across DRAM
    }
}
```

## 3. Hardware Mechanical Sympathy
- **Spatial Locality**: Row-wise iteration maximizes spatial locality by consuming all 64 bytes of each fetched cache line.
- **Hardware Stride Prefetchers**: Modern x86 and ARM Neoverse cores recognize sequential memory access within $\approx 2$ cache lines and issue speculative DRAM reads, hiding 200-cycle memory latency.
- **TLB Thrashing**: In massive matrices ($N > 10^4$), stride jumps cross $4\text{ KB}$ virtual memory page boundaries, evicting Translation Lookaside Buffer (TLB) entries and forcing CPU page-table walks.

---

## 🔗 References
- ⬆️ Parent: [[2D & Multidimensional Array]]
- 📚 Module: `Data Structures`
