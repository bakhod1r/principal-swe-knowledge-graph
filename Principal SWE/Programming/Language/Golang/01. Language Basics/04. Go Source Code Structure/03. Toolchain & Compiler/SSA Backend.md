---
title: SSA Backend
tags:
  - golang
  - goroot
  - compiler
  - optimization
parent: "[[Toolchain & Compiler]]"
---

# SSA Backend

The optimizing middle and back end of `cmd_compile`. Converts the typed AST into
Static Single Assignment form, runs ~50 optimization passes, then emits machine
code.

## 1. Pipeline

```text
typed AST
   │
   ▼
SSA construction        (every value assigned exactly once)
   │
   ▼ ~40 machine-independent passes
   ├── constant folding / propagation
   ├── dead code elimination
   ├── common subexpression elimination
   ├── nil-check and bounds-check elimination
   ├── loop-invariant hoisting
   └── inlining decisions already applied upstream
   │
   ▼ lowering — generic ops become machine ops
   │
   ▼ ~10 machine-dependent passes
   ├── register allocation
   ├── stack frame layout
   └── instruction scheduling
   │
   ▼
machine code → `cmd_link`
```

## 2. Seeing It

```bash
GOSSAFUNC=Fibonacci go build ./...   # writes ssa.html — every pass, clickable
go build -gcflags='-S' ./...          # final assembly
go build -gcflags='-m -m' ./...       # inlining + escape analysis decisions
```

`ssa.html` is the single best tool for understanding why Go generated the code it
did.

## 3. Bounds-Check Elimination

```go
func sum(s []int) int {
    t := 0
    for i := 0; i < len(s); i++ {   // BCE proves 0 <= i < len(s)
        t += s[i]                    // no bounds check emitted
    }
    return t
}
```

```bash
go build -gcflags='-d=ssa/check_bce/debug=1' ./...
```

Reports which checks survived — the checks that remain are the ones costing time.

## 4. Gotchas

- SSA is not a language feature; nothing here is guaranteed across releases.
- `-gcflags='all=-N -l'` disables optimization and inlining for debugging, and
  makes performance measurements meaningless.
- Escape analysis runs **before** SSA — see `Escape Analysis`.

---

## 🔗 References
- ⬆️ Parent: [[Toolchain & Compiler]]
