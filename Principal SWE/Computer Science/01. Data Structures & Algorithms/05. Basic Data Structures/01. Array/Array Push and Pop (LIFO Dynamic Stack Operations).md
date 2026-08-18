---
title: "Array Push and Pop (LIFO Dynamic Stack Operations)"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Push and Pop (LIFO Dynamic Stack Operations)

## 1. Definition
**Array Push and Pop** implements Last-In, First-Out (LIFO) stack operations on a dynamic array:
- **Push:** Appends an element to index $\text{len}$ in amortized $\mathcal{O}(1)$ time.
- **Pop:** Reads and removes the element at $\text{len}-1$ in strict $\mathcal{O}(1)$ time by decrementing $\text{len}$.

---

## 2. Mental Model
```text
Stack Index Top (T):
Push(50): Arr[T] = 50; T = T + 1;
Pop():    T = T - 1; Val = Arr[T]; Return Val;
```

---

## 3. Usage
```go
type Stack[T any] struct {
    items []T
}
func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 { var zero T; return zero, false }
    idx := len(s.items) - 1
    val := s.items[idx]
    var zero T
    s.items[idx] = zero // GC leak prevention
    s.items = s.items[:idx]
    return val, true
}
```

---

## 4. Gotchas
- **Popping an Empty Array:** Calling Pop on an empty array without length checking causes immediate runtime panic (`index out of range [-1]`).

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]

