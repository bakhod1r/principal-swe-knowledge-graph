---
title: "Array Memory Leaks and Garbage Collection Truncation"
tags:
  - review
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Memory Leaks and Garbage Collection Truncation

## 1. Definition
**Array Memory Leaks** occur in garbage-collected environments when unused backing array slots retain pointers to heap objects, preventing the GC from reclaiming unreachable memory.

---

## 2. Mental Model
```text
Memory Leak Scenario:
Slice Header: Ptr -> [ PtrA ][ PtrB ][ PtrC ][ PtrD ] (Len=4, Cap=4)
Truncate:     s = s[:2]
Slice Header: Ptr -> [ PtrA ][ PtrB ] (Len=2, Cap=4)
Backing Memory:      [ PtrA ][ PtrB ][ PtrC (STALE!) ][ PtrD (STALE!) ]
- Object C and D are unreachable by application, but held by slice backing array!
```

---

## 3. Usage
```go
// Safe truncation pattern
func TruncateSafely[T any](s []T, newLen int) []T {
    var zero T
    for i := newLen; i < len(s); i++ {
        s[i] = zero // Break GC reference chain!
    }
    return s[:newLen]
}
```

---

## 4. Gotchas
- **Long-Lived Struct Retention:** Retaining a tiny slice from a heavy parsed JSON buffer keeps the entire 100 MB parent buffer in RAM.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

