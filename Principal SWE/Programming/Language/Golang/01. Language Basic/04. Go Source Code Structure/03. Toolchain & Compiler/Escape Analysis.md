---
title: Escape Analysis
tags:
  - golang
  - goroot
  - compiler
  - memory
  - performance
parent: "[[Toolchain & Compiler]]"
---

# Escape Analysis

The compiler pass that decides whether each allocation lives on the **stack**
(free, reclaimed by returning) or the **heap** (costs GC work).

## 1. The Question It Answers

> Can this value still be referenced after the function returns?

No → stack. Yes, or unknown → heap.

## 2. Observing It

```bash
go build -gcflags='-m' ./...
```

```text
./user.go:14:13: &User{...} escapes to heap
./user.go:22:9:  make([]byte, 64) does not escape
./user.go:31:21: leaking param: s to result ~r0 level=0
```

`-m -m` adds the reasoning chain.

## 3. Common Escape Causes

| Pattern | Why it escapes |
|---|---|
| Returning `&local` | Outlives the frame |
| Storing into an interface | Concrete type unknown at the store |
| `fmt.Println(x)` | Variadic `...any` boxes the argument |
| Sending on a channel | Consumer lifetime unknown |
| Closure capturing by reference | Closure may outlive the frame |
| Slice with non-constant size | Size unknown at compile time |

The `fmt` row is why adding a debug print inside a hot loop can double
allocations.

## 4. It Is Not About `new` vs `&`

```go
x := new(int)      // may be stack-allocated
y := &Point{}      // may be stack-allocated
```

Both stay on the stack if they do not escape. Go has no stack/heap keyword —
the compiler decides, always.

## 5. Gotchas

- Escape analysis runs **before** `Inlining` finishes influencing it; inlining a
  callee can turn a heap allocation into a stack one, which is why small helpers
  are worth keeping small.
- "Does not escape" plus a large size still means a big stack frame — see
  `malloc (Memory Allocator)` for the heap side.
- Interface method calls block analysis; concrete types keep values on the stack.

---

## 🔗 References
- ⬆️ Parent: [[Toolchain & Compiler]]
