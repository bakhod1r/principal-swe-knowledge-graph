---
title: "Multidimensional Array Striding (Row-Major vs Column-Major)"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Multidimensional Array Striding (Row-Major vs Column-Major)

## 1. Definition
**Row-Major Order** (C, C++, Go, Python) lays out consecutive elements of a row adjacently in linear memory ($	ext{index} = i 	imes 	ext{Cols} + j$).
**Column-Major Order** (Fortran, Julia, MATLAB) lays out consecutive elements of a column adjacently in linear memory ($	ext{index} = j 	imes 	ext{Rows} + i$).

---

## 2. Mental Model
```text
2D Matrix 2x3:
[ [A, B, C],
  [D, E, F] ]

Row-Major Memory:    [ A ][ B ][ C ][ D ][ E ][ F ]
Sequential Traversal: Loop i { Loop j { Access(i, j) } } -> Sequential Cache Fetch (Fast!)

Column-Major Traversal on Row-Major Memory:
Access A -> Access D (Stride jump!) -> Access B -> Access E (Constant Cache Misses!)
```

---

## 3. Usage
```go
// Optimal Cache-Aware Row-Major Matrix Iteration
func SumMatrix(matrix [][]float64, rows, cols int) float64 {
    var sum float64
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            sum += matrix[r][c] // Sequential memory stride = 8 bytes
        }
    }
    return sum
}
```

---

## 4. Gotchas
- **Stride Thrashing:** Inverting loop nests (`for c ... for r ... matrix[r][c]`) on large matrices ($N > 4096$) causes every single read to miss L1/L2 caches and pull from main RAM, slowing execution down by $20	ext{x}	ext{--}50	ext{x}$.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]

