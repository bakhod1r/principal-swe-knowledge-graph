---
title: iter Package
tags:
  - golang
  - standard-library
  - generics
  - iterators
parent: "[[Standard Library]]"
---

# `iter`

**Go 1.23.** Defines the function signatures that `for ... range` accepts, making
user-defined iteration a language-level feature.

## 1. The Two Types

```go
type Seq[V any]      func(yield func(V) bool)
type Seq2[K, V any]  func(yield func(K, V) bool)
```

## 2. Writing One

```go
func (t *Tree) Walk() iter.Seq[int] {
    return func(yield func(int) bool) {
        var rec func(*node) bool
        rec = func(n *node) bool {
            if n == nil { return true }
            return rec(n.left) && yield(n.val) && rec(n.right)
        }
        rec(t.root)
    }
}

for v := range t.Walk() {
    if v > 100 { break }        // break propagates: yield returns false
}
```

The `bool` return is how `break`, `return`, and `continue` in the caller stop the
producer — no channels, no goroutines, no leaks.

## 3. Pull Iterators

```go
next, stop := iter.Pull(seq)
defer stop()
for {
    v, ok := next()
    if !ok { break }
}
```

For merging two sequences, where push-style iteration cannot interleave.

## 4. Where It Is Used

`slices` `All`/`Values`/`Sorted`/`Collect`, `maps` `Keys`/`Values`,
`strings.SplitSeq` (Go 1.24), `sql.Rows` helpers.

## 5. Gotchas

- Replaces the channel-based iterator pattern, which leaked a goroutine on early
  `break`. Migrate those.
- Always honour `yield`'s `false` return — ignoring it makes `break` not work.
- `iter.Pull` **must** have `stop` called, or the underlying iterator's deferred
  cleanup never runs.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
