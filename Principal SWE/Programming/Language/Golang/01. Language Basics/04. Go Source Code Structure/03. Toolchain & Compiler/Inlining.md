---
title: Inlining
tags:
  - golang
  - goroot
  - compiler
  - optimization
  - performance
parent: "[[Toolchain & Compiler]]"
---

# Inlining

Replacing a call with the callee's body. The enabling optimization — it exposes
constants and escape information the other passes then act on.

## 1. Observing Decisions

```bash
go build -gcflags='-m' ./...
```

```text
./user.go:12:6: can inline (*Store).get with cost 42 as: ...
./user.go:31:14: inlining call to (*Store).get
./user.go:44:6: cannot inline Process: function too complex: cost 128 exceeds budget 80
```

## 2. The Budget

Each function gets a **cost budget of 80 nodes**. Statements add cost; some
constructs historically set cost to infinity.

| Blocks inlining | Status |
|---|---|
| `recover` | Never inlinable |
| Type switches | Inlinable since Go 1.18 |
| `for ... range` over funcs | Depends on release |
| Closures | Inlinable in many cases since Go 1.20+ |
| Cost > 80 | Not inlined (mid-stack inlining still applies at depth) |

## 3. Why It Compounds

```go
func (p Point) X() int { return p.x }   // inlined
d := p.X() - q.X()                       // now pure field arithmetic
```

After inlining, the SSA passes see field loads instead of opaque calls, so
constant folding and bounds-check elimination can proceed. See `SSA Backend`.

## 4. Forcing and Preventing

```go
//go:noinline
func benchmarkTarget() { ... }
```

There is **no** `//go:inline` — you cannot force it. Restructure the function
instead, or set `-gcflags='-l=4'` for aggressive inlining while experimenting.

## 5. Gotchas

- Chasing "can inline" messages rarely pays; profile first with
  `go tool pprof`.
- Interface method calls are only inlined when devirtualization proves the
  concrete type — see `iface (Interfaces)`.
- `-gcflags='-l'` disables inlining and typically costs 10–30% throughput.

---

## 🔗 References
- ⬆️ Parent: [[Toolchain & Compiler]]
