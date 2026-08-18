---
title: internal/bytealg
tags:
  - golang
  - goroot
  - runtime
  - performance
  - assembly
parent: "[[Internal Packages]]"
---

# `internal/bytealg`

Hand-written assembly implementations of the byte and string primitives, per
architecture. The reason `strings.Index` and `bytes.Equal` are fast.

## 1. What It Implements

| Function | Used by |
|---|---|
| `Compare`, `Equal` | `bytes.Compare`, `bytes.Equal`, string comparison |
| `IndexByte`, `IndexByteString` | `strings.IndexByte`, `bufio` scanning |
| `Index`, `IndexString` | `strings.Index` — Rabin-Karp fallback plus SIMD |
| `Count`, `CountString` | `strings.Count` |

## 2. Per-Architecture Assembly

```text
internal/bytealg/
├── index_amd64.s      ← AVX2 / SSE4.2 paths
├── index_arm64.s      ← NEON
├── compare_amd64.s
└── equal_generic.go   ← portable Go fallback
```

Runtime CPU feature detection picks the widest available path, which is why the
same binary is fast on both old and new hardware.

## 3. Why It Is `internal`

These functions assume invariants the compiler guarantees (alignment, no nil
slices with non-zero length) and skip the checks a public API would need. Exposing
them would make those assumptions part of the compatibility surface.

## 4. Gotchas

- **Unimportable** — see `internal visibility rule`. Use `bytes` and `strings`;
  they compile down to these.
- Hand-rolling a byte-scan loop is almost always slower than `bytes.IndexByte` —
  you are competing with SIMD assembly.
- Small inputs take a different path than large ones; microbenchmarks with tiny
  buffers mislead.

---

## 🔗 References
- ⬆️ Parent: [[Internal Packages]]
