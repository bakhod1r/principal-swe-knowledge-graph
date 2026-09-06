---
title: "3D Array Layout, Voxel Tensors, and Coordinate Flattening"
tags:
  - review
  - computer-science
  - data-structures
  - matrix
  - tensors
  - principal-swe
parent: "[[Matrix & Multidimensional Arrays]]"
---

# 3D Array Layout, Voxel Tensors, and Coordinate Flattening

## 1. Definition
A 3-dimensional array (or rank-3 tensor) $A \in \mathbb{R}^{D \times H \times W}$ represents volumetric data (e.g., MRI scans, voxel grids, video frames $T \times H \times W$, or multi-channel CNN feature maps $C \times H \times W$). Coordinate triples $(d, r, c)$ are flattened into a 1D scalar buffer.

$$\text{Offset}(d, r, c) = d \times (H \times W) + r \times W + c$$

General generalized tensor stride formula for rank-$K$ tensor:
$$\text{Offset}(i_0, i_1, \dots, i_{K-1}) = \sum_{k=0}^{K-1} i_k \cdot S_k, \quad S_k = \prod_{m=k+1}^{K-1} D_m$$

## 2. Mental Model & Coordinate Slices

```text
3D Volume (Depth D, Height H, Width W):
Slice d=0:          Slice d=1:
[ 0,0,0 ][ 0,0,1 ]  [ 1,0,0 ][ 1,0,1 ]
[ 0,1,0 ][ 0,1,1 ]  [ 1,1,0 ][ 1,1,1 ]

Contiguous 1D Storage:
[ (0,0,0) (0,0,1) (0,1,0) (0,1,1) | (1,0,0) (1,0,1) (1,1,0) (1,1,1) ]
|<---------- Slice 0 ------------>| |<---------- Slice 1 ------------>|
```

## 3. Usage & Coordinate Inversion
Given a linear flat index $idx$, the 3D coordinates can be recovered without branching:
- $d = idx / (H \times W)$
- $rem = idx \% (H \times W)$
- $r = rem / W$
- $c = rem \% W$

## 4. Gotchas
- Integer overflow: A $1024 \times 1024 \times 1024$ voxel volume of 32-bit floats exceeds $4\text{ GB}$. Flattening offsets with 32-bit signed integers causes silent arithmetic overflow; always use `uint64_t` or `size_t`.
- Memory fragmentation: Allocating 3D arrays as `T***` (pointer to pointer to pointer) requires $D \times H$ separate dynamic heap allocations. Always allocate a single contiguous flat buffer: `malloc(D * H * W * sizeof(T))`.

---

## 🔗 References
- ⬆️ Parent: [[Matrix & Multidimensional Arrays]]
- 📚 Module: `Data Structures`
