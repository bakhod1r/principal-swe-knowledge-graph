---
title: sort Package
tags:
  - golang
  - standard-library
  - collections
  - legacy
parent: "[[Standard Library]]"
---

# `sort`

The pre-generics sorting package. Still present and still correct, but `slices`
is the modern choice for most uses.

## 1. Modern vs Legacy

```go
// Go 1.21+
slices.Sort(nums)
slices.SortFunc(users, func(a, b User) int { return cmp.Compare(a.Age, b.Age) })

// legacy
sort.Ints(nums)
sort.Slice(users, func(i, j int) bool { return users[i].Age < users[j].Age })
```

`slices` is type-specialized; `sort.Slice` uses reflection for the swap and is
measurably slower on large inputs.

## 2. The `sort.Interface` Still Matters

```go
type ByAge []User
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

sort.Sort(ByAge(users))
```

Needed when sorting something that is not a slice, or when the ordering is a
reusable named type.

## 3. Searching

```go
i := sort.SearchInts(nums, 42)
i := sort.Search(len(s), func(i int) bool { return s[i] >= target })
```

`sort.Search` is a general binary search over any monotone predicate — genuinely
useful beyond sorting, and it has no `slices` equivalent.

## 4. Gotchas

- `sort.Slice` is **not stable**; use `sort.SliceStable` or
  `slices.SortStableFunc`.
- `Less` must be a strict weak ordering. A `<=` comparison causes
  `panic: sort: comparison function is not a strict weak ordering` in Go 1.21+.
- `sort.Sort` on NaN-containing float slices produces undefined order.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
