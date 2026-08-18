---
title: testing Package
tags:
  - golang
  - standard-library
  - testing
parent: "[[Standard Library]]"
---

# `testing`

The test, benchmark, example, and fuzz framework. Driven by `go test`.

## 1. Table-Driven Tests — the Standard Form

```go
func TestSplit(t *testing.T) {
    tests := map[string]struct {
        in   string
        want []string
    }{
        "simple": {"a,b", []string{"a", "b"}},
        "empty":  {"", []string{""}},
    }
    for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
            t.Parallel()
            got := strings.Split(tc.in, ",")
            if !slices.Equal(got, tc.want) {
                t.Errorf("got %v, want %v", got, tc.want)
            }
        })
    }
}
```

## 2. `Error` vs `Fatal`

| | Continues | Use when |
|---|---|---|
| `t.Errorf` | ✅ | One of several independent assertions |
| `t.Fatalf` | ❌ stops the subtest | Setup failed; continuing would panic |

`t.Fatal` may only be called from the test goroutine — from a spawned goroutine it
does not stop the test.

## 3. Cleanup and Helpers

```go
func newDB(t *testing.T) *sql.DB {
    t.Helper()                    // failures report the CALLER's line
    db := open(t)
    t.Cleanup(func() { db.Close() })
    return db
}
```

`t.Cleanup` runs LIFO after the test and its subtests — correct with
`t.Parallel()`, unlike `defer` inside a helper.

## 4. Benchmarks

```go
func BenchmarkMarshal(b *testing.B) {
    v := makeValue()
    b.ReportAllocs()
    for b.Loop() {                 // Go 1.24 — replaces `for i := 0; i < b.N; i++`
        _, _ = json.Marshal(v)
    }
}
```

`b.Loop` keeps the compiler from optimizing the body away and needs no
`ResetTimer`.

## 5. Gotchas

- `t.Parallel()` subtests run **after** the parent function returns; the classic
  loop-variable capture bug is fixed by Go 1.22 semantics — see `go directive`.
- `TestMain(m *testing.M)` must call `os.Exit(m.Run())` or nothing runs.
- Examples with an `// Output:` comment are executed and compared; without it they
  are only compiled.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
