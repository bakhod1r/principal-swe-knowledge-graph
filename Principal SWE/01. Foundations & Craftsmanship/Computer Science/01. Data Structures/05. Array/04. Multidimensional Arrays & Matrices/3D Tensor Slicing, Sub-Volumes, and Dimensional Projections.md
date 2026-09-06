---
title: "3D Tensor Slicing, Sub-Volumes, and Dimensional Projections"
tags:
  - review
  - computer-science
  - data-structures
  - matrix
  - tensors
  - principal-swe
parent: "[[Matrix & Multidimensional Arrays]]"
---

# 3D Tensor Slicing, Sub-Volumes, and Dimensional Projections

## 1. Definition
A **Tensor View** allows inspecting, projecting, and modifying a sub-region of a multidimensional buffer without allocating new memory or copying bytes.

A strided tensor view is completely defined by:
1. `data_ptr`: Pointer to the first element of the sub-volume.
2. `shape`: Tuple of dimension extents $(D', H', W')$.
3. `strides`: Tuple of stride multipliers $(S_d, S_h, S_w)$.

## 2. Orthogonal Hyperplane Projections

```text
3D Volume Slicing across 3 Anatomical/Spatial Planes:
1. Axial Slice (Z-plane): Fix d = k   ==> Strides: (S_h, S_w) [Contiguous in row-major]
2. Coronal Slice (Y-plane): Fix r = k ==> Strides: (S_d, S_w) [Strided jumps]
3. Sagittal Slice (X-plane): Fix c = k==> Strides: (S_d, S_h) [Non-contiguous jumps]
```

## 3. Stride Manipulation Algebra
Slicing tensor $A[d_0:d_1, r_0:r_1, c_0:c_1]$ with step sizes $(s_d, s_r, s_c)$:
- $\text{New Pointer} = \text{Base} + d_0 \cdot S_d + r_0 \cdot S_r + c_0 \cdot S_w$
- $\text{New Strides} = (S_d \cdot s_d, S_r \cdot s_r, S_w \cdot s_c)$
- $\text{New Shape} = ((d_1 - d_0)/s_d, (r_1 - r_0)/s_r, (c_1 - c_0)/s_c)$

## 4. Gotchas
- **Non-Contiguous Views**: Transposing or slicing across non-stride-1 dimensions makes the tensor non-contiguous (`is_contiguous() == false`). Attempting flat memory operations (`memcpy`, SIMD vectorization) on non-contiguous tensor views corrupts data. Always invoke `.contiguous()` before vector operations.

---

## 🔗 References
- ⬆️ Parent: [[Matrix & Multidimensional Arrays]]
- 📚 Module: `Data Structures`
