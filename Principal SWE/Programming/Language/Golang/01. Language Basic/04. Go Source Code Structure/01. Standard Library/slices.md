---
title: slices Package
tags:
  - golang
  - standard-library
  - generics
  - collections
parent: "[[Standard Library]]"
---

# `slices`

Generic slice operations. Standard library since Go 1.21; replaces most hand-rolled
loops and the `sort` boilerplate.

## 1. Sorting

```go
slices.Sort(nums)                                  // ordered types
slices.SortFunc(users, func(a, b User) int {       // cmp semantics: -1, 0, +1
    return cmp.Compare(a.Age, b.Age)
})
slices.SortStableFunc(users, byName)
slices.IsSorted(nums)
```

Replaces `sort.Slice` — and is faster, because it is type-specialized rather than
reflection-based. See `sort`.

## 2. Searching

```go
i := slices.Index(s, target)
ok := slices.Contains(s, target)
i, found := slices.BinarySearch(sortedNums, 42)
```

## 3. Manipulation

```go
s = slices.Insert(s, 2, x, y)
s = slices.Delete(s, 1, 3)          // removes [1,3)
s = slices.Compact(s)               // dedupe ADJACENT equal elements
s2 := slices.Clone(s)
slices.Reverse(s)
slices.Equal(a, b)
```

## 4. Iterators (Go 1.23)

```go
for i, v := range slices.All(s) { ... }
for v := range slices.Values(s) { ... }
sorted := slices.Sorted(maps.Keys(m))     // the canonical stable map iteration
```

See `iter` and `maps`.

## 5. Gotchas

- `Delete` and `Insert` **modify the backing array** and return a new header.
  Always reassign: `s = slices.Delete(s, ...)`.
- `Compact` only removes *adjacent* duplicates — sort first for full dedupe.
- `Clone` is shallow; element pointers are shared.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
