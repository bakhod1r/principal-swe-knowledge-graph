---
title: "Direct Stack-to-Stack Copy Optimization"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Channel Architecture]]"
---
# Direct Stack-to-Stack Copy Optimization in Go

This is a **Go compiler/runtime optimization** related to how values move between goroutine stacks and other storage locations. The key idea is:

> When the compiler/runtime can prove that a value only needs to move from one stack location to another, it may copy the value **directly between stack slots**, avoiding an intermediate heap allocation or temporary object.

This matters because Go heavily relies on goroutine stacks, and stack operations are much cheaper than heap allocation + garbage collection.

---

## 1. The Problem

Consider a function returning a struct:

```go
type Result struct {
    A int
    B int
    C int
}

func makeResult() Result {
    return Result{1, 2, 3}
}

func caller() {
    r := makeResult()
    use(r)
}
```

Conceptually, you might imagine:

```text
makeResult stack
      │
      │ create Result
      ▼
temporary Result
      │
      ▼
caller stack
      │
      ▼
use()
```

A naive implementation could involve unnecessary intermediate storage.

The compiler wants something closer to:

```text
caller stack
┌──────────────┐
│ Result       │
│ A            │
│ B            │
│ C            │
└──────────────┘
       ▲
       │ direct copy
       │
makeResult
```

The exact implementation depends on the compiler's generated code and ABI, but the optimization principle is **eliminating unnecessary intermediate copies/storage**.

---

# 2. Mental Model

Think about a goroutine's stack as a collection of memory slots:

```text
goroutine stack

┌──────────────────────┐
│ caller's variables   │
├──────────────────────┤
│ return-value slots   │
├──────────────────────┤
│ callee's variables   │
└──────────────────────┘
```

Suppose:

```go
func f() Large {
    return x
}

func g() {
    y := f()
}
```

The compiler may arrange the ABI/frame layout so that the result of `f()` is written **directly into the destination slot used by `g()`**.

Instead of:

```text
x
 ↓
temporary
 ↓
y
```

it can become effectively:

```text
x ───────────────► y
```

This is the essence of **direct stack-to-stack copying**.

---

# 3. Why It Matters

The cost of moving a value can be roughly thought of as:

```text
unnecessary temporary:

source
  ↓
temporary allocation/storage
  ↓
destination
```

versus:

```text
source ─────► destination
```

The second approach can reduce:

- memory traffic
    
- stack usage
    
- temporary objects
    
- register spills
    
- heap allocations in some cases
    
- GC pressure
    
- generated instructions
    

This is especially relevant for:

```go
struct
array
multiple return values
large values
defer arguments
interface conversions
```

But don't interpret this as:

> "Go always copies values directly."

The compiler only performs transformations it can prove are safe.

---

# 4. Stack vs Heap

This optimization is easier to understand if you distinguish **copying** from **allocation**.

A copy itself isn't necessarily expensive:

```go
dst = src
```

For a small struct:

```go
type Point struct {
    X int
    Y int
}
```

the compiler may simply use registers:

```text
RAX ← X
RBX ← Y
```

For larger values, it may emit memory moves.

The bigger concern is often:

```text
heap allocation
      +
memory initialization
      +
GC tracking
      +
additional copy
```

Therefore:

```text
stack → stack
```

is generally preferable to:

```text
stack → heap temporary → stack
```

when the intermediate object isn't required.

---

# 5. Connection to Escape Analysis

This optimization is closely related to Go's **escape analysis**, but they are not the same thing.

Escape analysis asks:

> Does this value need to outlive the current stack frame or otherwise become accessible from a context that requires heap allocation?

Example:

```go
func f() *int {
    x := 42
    return &x
}
```

`x` cannot remain only in `f`'s stack frame because the pointer escapes:

```text
f stack
   │
   │ x
   ▼
heap
   ▲
   │
returned pointer
```

But:

```go
func f() int {
    x := 42
    return x
}
```

can remain entirely stack/register based:

```text
x
│
└──────► caller's result
```

You can inspect the compiler's escape analysis with:

```bash
go build -gcflags="-m=2" .
```

---

# 6. A More Interesting Example

Consider:

```go
type User struct {
    ID      int64
    Balance int64
    Flags   uint64
}

func loadUser() User {
    return User{
        ID:      42,
        Balance: 1000,
        Flags:   1,
    }
}

func process() {
    u := loadUser()
    consume(u)
}
```

At the source-code level:

```text
loadUser()
    ↓
returned User
    ↓
u
    ↓
consume()
```

But the compiler can potentially arrange the calling convention such that the result is materialized directly in the caller's expected location.

Conceptually:

```text
             caller frame
        ┌───────────────────┐
        │ u.ID              │ ◄────┐
        │ u.Balance         │ ◄────┤ loadUser
        │ u.Flags           │ ◄────┘
        └───────────────────┘
```

There doesn't have to be a conceptual:

```text
temporary User
```

between the producer and consumer.

---

# 7. Multiple Return Values

This becomes particularly interesting with Go's multiple return values:

```go
func lookup() (int, error) {
    return 42, nil
}

func handler() {
    value, err := lookup()

    if err != nil {
        return
    }

    use(value)
}
```

The ABI can pass results through:

- registers
    
- stack slots
    
- a combination of both
    

depending on architecture, value size, compiler decisions, and ABI details.

For example, conceptually:

```text
lookup()

value ─────► register
err   ─────► register
              │
              ▼
          handler()
```

No temporary heap object is necessary.

---

# 8. Important Distinction: Copy Elision vs Direct Stack Copy

Don't confuse this with **C++-style copy elision**.

Go has its own compiler implementation and calling conventions.

The compiler performs many transformations such as:

- SSA optimization
    
- dead-code elimination
    
- register allocation
    
- scalar replacement
    
- store/load elimination
    
- inlining
    
- escape analysis
    
- stack slot reuse
    
- memory movement optimization
    

Therefore "direct stack-to-stack copy" is best understood as an **implementation-level optimization**, not a Go language guarantee.

Your program must remain correct even if the compiler decides to perform an actual copy.

---

# 9. Stack Growth Makes This Interesting

Go goroutine stacks are dynamically managed.

Historically Go used segmented stacks, but modern Go uses **contiguous growable stacks**.

A goroutine may start with a small stack:

```text
┌──────────────┐
│ small stack  │
└──────────────┘
```

and grow:

```text
┌────────────────────────┐
│ larger stack           │
└────────────────────────┘
```

When stack growth occurs, stack-resident pointers and references have to remain valid.

This is one reason the runtime/compiler has strict knowledge about stack layout and stack maps.

The runtime needs to know things like:

```text
which slots contain pointers?
which slots contain scalar data?
where are return values?
where are locals?
```

That metadata is critical for GC and stack copying.

---

# 10. Stack-to-Stack Copy During Stack Growth

There is another meaning worth distinguishing.

When a goroutine's stack grows, the runtime itself may need to copy stack contents:

```text
old stack

┌───────────────┐
│ locals        │
│ pointers      │
│ frames        │
└───────────────┘
        │
        │ copy
        ▼
new stack

┌───────────────────────────┐
│ locals                    │
│ pointers                  │
│ frames                    │
└───────────────────────────┘
```

This is **not the same optimization** as compiler-level direct stack-to-stack value passing.

So distinguish:

### Compiler optimization

```text
source stack slot ─────► destination stack slot
```

### Runtime stack growth

```text
old goroutine stack ───► new goroutine stack
```

They operate at different levels.

---

# 11. How the Compiler Thinks About It

Modern Go compilation goes roughly through:

```text
Go source
   │
   ▼
AST
   │
   ▼
type checking
   │
   ▼
IR
   │
   ▼
SSA
   │
   ├── escape analysis
   ├── dead-code elimination
   ├── value propagation
   ├── bounds-check elimination
   ├── nil-check elimination
   ├── inlining
   └── other optimizations
   │
   ▼
register allocation
   │
   ▼
machine code
```

At the SSA level, the compiler tries to represent values without unnecessary memory operations.

For example, conceptually:

```text
v1 = load source
store temporary, v1
v2 = load temporary
store destination, v2
```

can potentially become:

```text
v1 = load source
store destination, v1
```

or even:

```text
v1 → register → consumer
```

depending on the situation.

---

# 12. SSA Is the Important Mental Model

A strong way to reason about this is:

> Don't think primarily in terms of source-level variables. Think in terms of **values, lifetimes, storage locations, and data flow**.

Source:

```go
a := f()
b := a
use(b)
```

doesn't necessarily mean:

```text
stack[a]
   ↓
stack[b]
```

The compiler may discover:

```text
f result
   ↓
register
   ↓
use
```

or:

```text
f result
   ↓
stack slot reused as b
```

or:

```text
f result
   ↓
destination stack slot
```

The source-level assignment does **not** dictate the physical movement.

---

# 13. How to Verify Instead of Guessing

This is where production-level engineering judgment matters.

Don't assume an optimization exists because the source code looks efficient.

Inspect the generated code.

### Compiler diagnostics

```bash
go build -gcflags="-m=2" .
```

Useful for:

- escape analysis
    
- inlining decisions
    
- compiler reasoning
    

### Assembly

```bash
go tool compile -S file.go
```

or:

```bash
go build -gcflags="-S" .
```

Look for:

```text
MOV
MOVQ
MOVUPS
DUFFCOPY
CALL runtime.memmove
```

depending on the generated code and architecture.

### Benchmark

```go
func BenchmarkLoadUser(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = loadUser()
    }
}
```

Then:

```bash
go test -bench=. -benchmem
```

You want to measure:

```text
ns/op
B/op
allocs/op
```

rather than infer performance from source syntax.

---

# 14. Large Structs

Suppose:

```go
type Large struct {
    Data [1024]byte
}
```

Returning it:

```go
func create() Large {
    var x Large
    return x
}
```

may require substantial memory movement if the compiler cannot optimize it away.

Conceptually:

```text
1024 bytes
     │
     ▼
memory copy
```

This doesn't automatically mean:

> "Returning large structs is bad."

The compiler may optimize aggressively, and sometimes returning a value is preferable to introducing pointers.

For example:

```go
func create() Large
```

can be better than:

```go
func create() *Large
```

because the pointer version can introduce:

```text
heap allocation
GC pressure
pointer indirection
```

So **don't optimize based solely on object size**.

Measure.

---

# 15. A Common Wrong Mental Model

### Wrong

> "Go function return always copies the return value."

Not necessarily.

### Also wrong

> "Go never copies return values."

Also false.

The actual behavior depends on:

- ABI
    
- compiler version
    
- architecture
    
- optimization opportunities
    
- escape analysis
    
- inlining
    
- value size
    
- whether the value is address-taken
    
- whether it crosses a function boundary
    
- register availability
    

Therefore:

```text
Go semantics
     ≠
specific machine-level copy strategy
```

This distinction is fundamental.

---

# 16. Production Engineering Rule

When writing Go, don't write code specifically to force a hypothetical stack-copy optimization unless profiling proves you need it.

Prefer:

```go
func NewUser() User {
    return User{...}
}
```

over prematurely changing the API to:

```go
func NewUser() *User {
    return &User{...}
}
```

simply because you fear copying.

Then measure:

```bash
go test -bench=. -benchmem
```

and profile if necessary.

---

# 17. Performance Hierarchy

A useful mental hierarchy is:

```text
register
   ↓
stack
   ↓
heap
   ↓
remote memory / network
```

Generally, as you move downward:

```text
latency ↑
bandwidth cost ↑
coordination ↑
GC implications ↑
failure possibilities ↑
```

But this is not an absolute performance law.

For example, an unnecessarily huge stack copy can be more expensive than a well-managed heap object.

The correct principle is:

> **Minimize unnecessary data movement and allocation, then verify with measurement.**

---

# 18. Connection to Go's ABI

Modern Go uses an internal ABI designed to efficiently pass arguments and results, commonly referred to in compiler/runtime discussions as **ABIInternal**.

Small values can often be passed through registers.

Conceptually:

```text
Caller
  │
  ├── RAX = result part 1
  ├── RBX = result part 2
  └── ...
  │
  ▼
Callee
```

Larger values may involve stack memory.

Therefore:

```text
small value
    → registers often

larger aggregate
    → registers + stack / memory

escaping object
    → potentially heap
```

This is one reason modern Go performance can differ significantly from simplistic models based on "everything lives on the stack."

---

# 19. Key Takeaway

The deepest mental model is:

```text
             Go source
                 │
                 ▼
          logical values
                 │
                 ▼
              SSA
                 │
        ┌────────┴────────┐
        ▼                 ▼
     registers           stack
        │                 │
        └────────┬────────┘
                 ▼
              machine
                 │
                 ▼
        actual memory traffic
```

**Direct stack-to-stack copy optimization** means the compiler/runtime can arrange value movement so that an unnecessary intermediate representation or temporary storage is avoided.

But the important Staff+/Principal-level lesson is broader:

> **Source-level assignments describe semantics, not physical data movement.**

When performance matters, reason about:

```text
value lifetime
    ↓
escape analysis
    ↓
SSA data flow
    ↓
register allocation
    ↓
stack layout
    ↓
generated assembly
    ↓
benchmark/profile
```

That is the right path from **"I think this copies"** to **"I know what the machine actually does."**
---

## 🔗 References
- ⬆️ Parent: [[Channel Architecture]]
- 📚 Module: `Concurrency & Synchronization`
