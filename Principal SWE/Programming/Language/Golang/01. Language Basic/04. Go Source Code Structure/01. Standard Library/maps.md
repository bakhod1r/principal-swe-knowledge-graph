---
title: maps Package
tags:
  - golang
  - standard-library
  - generics
  - collections
parent: "[[Standard Library]]"
---

# `maps`

Generic map helpers. Standard library since Go 1.21, extended with iterators in
1.23.

## 1. Functions

```go
keys := slices.Sorted(maps.Keys(m))     // iterator → sorted slice
vals := slices.Collect(maps.Values(m))
maps.Clone(m)                            // shallow copy
maps.Equal(a, b)
maps.Copy(dst, src)                      // merge src into dst
maps.DeleteFunc(m, func(k string, v int) bool { return v == 0 })
```

## 2. Map Iteration Is Randomized

```go
for k, v := range m { }   // DELIBERATELY random order, different every run
```

The runtime randomizes the starting bucket so code cannot depend on order. For
deterministic output:

```go
for _, k := range slices.Sorted(maps.Keys(m)) {
    fmt.Println(k, m[k])
}
```

This is the single most common use of the package. See `slices`.

## 3. `Keys`/`Values` Return Iterators

Since Go 1.23 they return `iter.Seq`, not slices:

```go
maps.Keys(m)                     // iter.Seq[K] — lazy
slices.Collect(maps.Keys(m))     // []K
slices.Sorted(maps.Keys(m))      // []K, sorted
```

Code written against the older `x/exp/maps` (which returned slices) does not
compile against the standard one. See `iter`.

## 4. Gotchas

- Deleting during `range` is safe and defined; adding during `range` is defined
  but the new entry may or may not be visited.
- `maps.Clone` is shallow — nested maps and slices stay shared.
- Map internals: `map (Hash Maps)`.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
