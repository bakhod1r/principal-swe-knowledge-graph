---
title: bytes Package
tags:
  - golang
  - standard-library
  - io
  - text
parent: "[[Standard Library]]"
---

# `bytes`

The `[]byte` mirror of `strings` — same function names, mutable data, no
allocation on slicing.

## 1. Why It Exists Separately

```go
strings.Contains(s, sub)   // string  — immutable
bytes.Contains(b, sub)     // []byte  — mutable, no conversion cost
```

Converting `[]byte(s)` or `string(b)` **copies**. In an I/O path that copy is the
cost you are trying to avoid.

## 2. `bytes.Buffer`

```go
var buf bytes.Buffer
buf.WriteString("hello ")
fmt.Fprintf(&buf, "%d", 42)
io.Copy(dst, &buf)
```

Implements both `io.Reader` and `io.Writer` — the standard in-memory pipe. See
`io`.

## 3. `bytes.Buffer` vs `strings.Builder`

| | `bytes.Buffer` | `strings.Builder` |
|---|---|---|
| Readable | ✅ implements `io.Reader` | ❌ write-only |
| `String()` cost | Copies | No copy |
| Reusable via `Reset` | ✅ | ✅ |

Building a string only → `strings.Builder`. Need to read it back or pass as a
reader → `bytes.Buffer`.

## 4. Gotchas

- `bytes.NewReader(b)` is cheaper than `bytes.NewBuffer(b)` when you only read.
- The compiler optimizes some `string([]byte)` conversions away in map lookups
  and comparisons — but not in general.
- `bytes.Equal` is SIMD assembly (`internal_bytealg`); a manual loop is slower.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
