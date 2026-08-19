---
title: "Structure of Arrays (SoA) vs Array of Structures (AoS)"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Structure of Arrays (SoA) vs Array of Structures (AoS)

## 1. Definition
- **Array of Structures (AoS):** Storing complete objects contiguously (`[ {x, y, z, id}, {x, y, z, id} ]`). Ideal for OOP, but wastes cache memory when querying single fields.
- **Structure of Arrays (SoA):** Storing individual fields in separate contiguous arrays (`x[], y[], z[], id[]`). Maximizes SIMD vectorization and cache density.

---

## 2. Mental Model
```text
AoS Memory (Querying X only):
[ X0 ][ Y0 ][ Z0 ][ ID0 ][ X1 ][ Y1 ][ Z1 ][ ID1 ]
- 75% of fetched cache line data is discarded!

SoA Memory (Querying X only):
[ X0 ][ X1 ][ X2 ][ X3 ][ X4 ][ X5 ][ X6 ][ X7 ]
- 100% of fetched cache line data is processed! (Full SIMD vectorization!)
```

---

## 3. Usage
```cpp
// Structure of Arrays (SoA) for Game Physics & Graphics
struct ParticleSystemSoA {
    alignas(32) std::vector<float> posX;
    alignas(32) std::vector<float> posY;
    alignas(32) std::vector<float> posZ;
    alignas(32) std::vector<float> velX;
    alignas(32) std::vector<float> velY;
    alignas(32) std::vector<float> velZ;
};
```

---

## 4. Gotchas
- **Object Construction Complexity:** SoA requires managing multiple separate buffers, increasing code complexity compared to simple AoS structs.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

