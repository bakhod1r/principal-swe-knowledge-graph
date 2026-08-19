---
title: "const Declaration"
tags:
  - golang
  - language-basics
  - const-and-iota
  - principal-swe
parent: "[[Const & Iota]]"
---

# const Declaration

## 1. Definition
**const Declaration** is a core operational primitive and fundamental structural paradigm within **Const & Iota**.
Explicit `var` declarations vs short variable declarations `:=`, package-level initialization, and type inference semantics.

It establishes rigorous engineering guarantees and runtime performance bounds:
- **Type-Safety & Semantics:** Guaranteed by Go's strict static type checker with zero implicit runtime conversions.
- **Memory & Allocation Efficiency:** Designed for stack allocation, contiguous memory cachelines, and minimal GC pressure.
- **Production Systems Leverage:** Enables resilient, deterministic systems programming with clear concurrency boundaries and mechanical sympathy with modern multi-core CPU architectures.

---

## 2. Mental Model

```text
Execution Pipeline & Memory Topology for const Declaration:
[ Go Source Expression / AST Representation ]
                     │
                     ▼
[ Compiler Type Check / SSA Optimization Pass ]
                     │
                     ▼
[ Runtime Execution (Goroutine Stack / TCMalloc Heap / GC Engine) ]
                     │
                     ▼
[ Hardware Layer (Registers, 64-byte Cache Lines, Memory Bus) ]
```

### Key Architectural Invariants:
1. **Mechanical Sympathy:** Sequential memory structures maximize CPU L1/L2 data prefetching.
2. **Explicit Error & Memory Boundaries:** Avoid hidden runtime magic; allocations are deterministic and verifiable via `go build -gcflags="-m"`.
3. **Thread Safety Guarantees:** Concurrent state mutations require explicit synchronization via channels, atomics, or mutexes.

---

## 3. Usage

// Production Var vs Short Declaration patterns
package main

import "fmt"

// Package level: explicit var required (:= not permitted at package scope)
var (
	GlobalConfig string = "production"
	DefaultPort  int    = 8080
)

func ProcessRequest(input string) (string, error) {
	// Block scope: := short declaration with automatic type inference
	status := "processing"
	
	// Redeclaration rule: at least one variable on the left must be new
	result, err := executeSubtask(input)
	if err != nil {
		return "", err
	}
	
	// err is reassigned, updatedCount is newly declared
	updatedCount, err := recordMetrics(status)
	if err != nil {
		return "", err
	}
	
	return fmt.Sprintf("%s: %d ops", result, updatedCount), nil
}

func executeSubtask(in string) (string, error) { return in + "_done", nil }
func recordMetrics(st string) (int, error)     { return 1, nil }

---

## 4. Gotchas

- **Variable Shadowing Trap:** Using `:=` inside an inner block (e.g. `if res, err := fn(); ...`) shadows the outer `err` variable, leaving the outer error `nil`.
- **Ignoring Escape Analysis:** Passing local variables to interfaces (`fmt.Println(val)`) or returning pointers causes heap escapes, increasing GC mark-sweep latency.
- **Unchecked Type Assertions:** Performing direct type assertions (`x.(MyType)`) without the comma-ok idiom (`v, ok := x.(MyType)`) triggers panics on mismatched types.

---

## 🔗 References
- ⬆️ Parent: [[Const & Iota]]
- 📚 Module: `Language Basics`
