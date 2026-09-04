---
title: "Thread Pinning (LockOSThread & UnlockOSThread)"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Goroutines]]"
---
# Thread Pinning in Go — `LockOSThread` & `UnlockOSThread`

Thread pinning means **forcing the current goroutine to stay associated with the same OS thread**.

In Go, this is controlled by:

```go
runtime.LockOSThread()
runtime.UnlockOSThread()
```

It is an **advanced runtime feature** and should be used only when some external system requires OS-thread affinity.

---

## 1. The Problem

Normally, Go deliberately separates:

```text
Goroutine
    ↓
Go scheduler
    ↓
OS thread
```

A goroutine can execute on one OS thread now and another later:

```text
time ─────────────────────────────>

G1:    OS Thread 1 ──────┐
                         │
                         └────── OS Thread 2
```

This is normally exactly what you want.

But some APIs have **thread-affinity requirements**:

```text
Goroutine
    │
    │ must remain here
    ▼
OS Thread #42
    │
    ├── initialize library
    ├── call library
    └── cleanup library
```

Examples include certain:

- GUI frameworks
    
- OpenGL contexts
    
- OS-specific APIs
    
- thread-local storage
    
- C libraries with thread-local state
    
- foreign-function interfaces
    

For these cases, Go provides thread pinning.

---

# 2. `runtime.LockOSThread()`

Calling:

```go
runtime.LockOSThread()
```

binds the **calling goroutine** to its current OS thread.

Example:

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Println("running on a pinned OS thread")

	// Thread-sensitive operations...
}
```

The important distinction is:

> `LockOSThread()` pins a **goroutine**, not the OS thread itself.

The OS thread remains an OS thread managed by the operating system. Go simply prevents the locked goroutine from being migrated to another OS thread.

---

# 3. Mental Model

Without pinning:

```text
              Go Scheduler
                   │
          ┌────────┼────────┐
          ▼        ▼        ▼
        M1        M2       M3
         │         │        │
       G1/G2      G3       G4
```

A goroutine can move:

```text
G1
 │
 ├── M1
 │
 ├── M2
 │
 └── M3
```

With:

```go
runtime.LockOSThread()
```

the relationship becomes:

```text
G1 ───────────────► M1

G1 cannot migrate away from M1
while locked.
```

Other goroutines can still execute on that OS thread when the scheduler permits it; the important guarantee is that the **locked goroutine remains tied to that thread**.

---

# 4. Why Does This Exist?

Go's scheduler normally abstracts away OS threads.

That abstraction is extremely useful:

```text
Application
    │
    ▼
Goroutines
    │
    ▼
Go runtime
    │
    ▼
OS threads
```

Most Go code should never care which OS thread executes a goroutine.

But external APIs sometimes do care.

For example, imagine a C library maintaining state in thread-local storage:

```text
OS Thread A
    └── TLS:
         context = X

OS Thread B
    └── TLS:
         context = Y
```

Suppose:

```go
initialize()

operation()

cleanup()
```

must all happen on the same OS thread.

Without pinning:

```text
initialize() → Thread A
operation()  → Thread B
cleanup()    → Thread A
```

The external library may observe completely different thread-local state.

With pinning:

```text
initialize() → Thread A
operation()  → Thread A
cleanup()    → Thread A
```

---

# 5. Important: Pinning Is About the Goroutine

Consider:

```go
func worker() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// ...
}
```

The goroutine executing `worker()` becomes locked to its current OS thread.

It does **not** mean:

```text
"Reserve this OS thread exclusively for me forever."
```

It means:

```text
"This goroutine must continue running on this OS thread."
```

That distinction is critical.

---

# 6. `UnlockOSThread()`

To release the affinity:

```go
runtime.UnlockOSThread()
```

Example:

```go
func worker() {
	runtime.LockOSThread()

	// thread-affine work

	runtime.UnlockOSThread()
}
```

Usually prefer:

```go
func worker() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// thread-affine work
}
```

The `defer` is particularly important because forgetting to unlock can cause serious scheduler/resource problems.

---

# 7. Lock/Unlock Must Be Balanced

Think of thread locking as a nesting counter.

Conceptually:

```text
LockOSThread()
    lock depth = 1

LockOSThread()
    lock depth = 2

UnlockOSThread()
    lock depth = 1

UnlockOSThread()
    lock depth = 0
```

Therefore:

```go
runtime.LockOSThread()
runtime.LockOSThread()

runtime.UnlockOSThread()
runtime.UnlockOSThread()
```

fully unlocks the goroutine.

This matters when libraries or helper functions themselves manipulate thread locking.

---

# 8. The `go` Statement Is a Major Trap

Consider:

```go
runtime.LockOSThread()

go func() {
	// ...
}()
```

The new goroutine is **not automatically pinned to the parent's OS thread**.

You have:

```text
G1 ── locked ──► M1

        go
        │
        ▼

G2 ────────────► scheduler chooses M
```

If G2 needs thread affinity, it must establish its own lock:

```go
go func() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// thread-affine work
}()
```

---

# 9. `LockOSThread` and Blocking

A very important production concern is blocking.

Suppose:

```go
runtime.LockOSThread()

someBlockingOperation()

runtime.UnlockOSThread()
```

The locked goroutine remains associated with that OS thread while blocked.

Go can create/use additional OS threads when necessary, but you should not interpret this as:

> "Thread pinning is free."

It can increase scheduler pressure and OS-thread usage.

This becomes particularly relevant when many goroutines are pinned and/or block frequently.

---

# 10. Bad Pattern: Pinning Everything

This is usually wrong:

```go
func handler() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// entire HTTP request
	// database call
	// network request
	// business logic
	// logging
}
```

You have transformed normal Go concurrency into unnecessary thread affinity.

Instead:

```go
func handler() {
	// normal Go execution

	threadSensitiveOperation()
}
```

and isolate the pinning:

```go
func threadSensitiveOperation() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Only the truly thread-affine section
}
```

Even better, if the external API requires a persistent thread context, design a dedicated worker architecture.

---

# 11. Dedicated Thread Worker Pattern

For serious thread-affine systems, a useful pattern is:

```text
                    Requests
                       │
                       ▼
                 ┌───────────┐
                 │ Channel   │
                 └─────┬─────┘
                       │
                       ▼
              ┌─────────────────┐
              │ Locked Goroutine│
              │                 │
              │ LockOSThread()  │
              │                 │
              │ External API    │
              │ External API    │
              │ External API    │
              └─────────────────┘
                       │
                       ▼
                    Results
```

Example:

```go
type Request struct {
	Input string
	Reply chan error
}

func worker(requests <-chan Request) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Initialize thread-local / thread-affine state.

	for req := range requests {
		err := process(req.Input)
		req.Reply <- err
	}

	// Cleanup must happen on the same OS thread.
}
```

This is often much safer than sprinkling `LockOSThread()` throughout business logic.

---

# 12. Why a Channel Worker Is Powerful

Suppose an external library requires:

```text
Initialize → Thread A
Use        → Thread A
Cleanup    → Thread A
```

A dedicated worker gives you a clear ownership boundary:

```text
Thread-affine state
        │
        ▼
   worker goroutine
        │
        ├── initialization
        ├── requests
        └── cleanup
```

The rest of the application remains normal Go:

```text
Goroutines
   │
   ▼
Channel
   │
   ▼
Pinned worker
   │
   ▼
External API
```

This is a strong architectural pattern because it isolates the unusual constraint.

---

# 13. A Subtle Point: Goroutine Migration After Unlock

Consider:

```go
runtime.LockOSThread()

// work

runtime.UnlockOSThread()

// more work
```

After unlocking, the goroutine is again free to be scheduled on another OS thread.

Therefore:

```text
Before Unlock:

G1 ─────────► M1

After Unlock:

G1 ───────► M1
       │
       └────► M2
```

You must not assume it remains on the same thread.

---

# 14. `runtime.LockOSThread()` Is Not a Synchronization Primitive

This is another common misconception.

It does **not** provide:

- mutual exclusion
    
- memory synchronization
    
- ordering between goroutines
    
- protection of shared data
    
- race prevention
    

This:

```go
runtime.LockOSThread()
counter++
runtime.UnlockOSThread()
```

does **not** make `counter++` safe against concurrent access.

For synchronization use appropriate primitives:

```go
sync.Mutex
sync.RWMutex
sync/atomic
channels
```

Thread affinity and synchronization are different concepts.

---

# 15. Thread Pinning vs Mutex

### Mutex

Question:

> "Who can access this shared resource?"

```text
G1 ──┐
G2 ──┼──► Mutex ──► resource
G3 ──┘
```

### Thread pinning

Question:

> "Which OS thread must this goroutine execute on?"

```text
G1 ─────────► OS Thread 7
```

They solve fundamentally different problems.

---

# 16. Thread Pinning vs `runtime.GOMAXPROCS`

These are also unrelated.

`GOMAXPROCS` controls approximately how many OS threads can execute Go code simultaneously.

```text
GOMAXPROCS = N
        │
        ▼
parallel Go execution capacity
```

`LockOSThread` controls whether a specific goroutine can migrate between OS threads.

```text
LockOSThread
     │
     ▼
goroutine ↔ OS-thread affinity
```

Do not use one as a substitute for the other.

---

# 17. Common Use Cases

### GUI

Some GUI frameworks require:

```text
UI operations → main OS thread
```

A Go application may need:

```go
runtime.LockOSThread()
```

very early in startup.

---

### OpenGL

An OpenGL context may be associated with a particular OS thread.

Conceptually:

```text
Thread A
   │
   └── OpenGL Context

Goroutine
   │
   └── must remain on Thread A
```

---

### cgo

A C library may depend on:

```text
thread-local state
```

For example:

```text
Thread A:
    C TLS = context A

Thread B:
    C TLS = context B
```

Go thread migration can violate those assumptions unless you explicitly maintain affinity.

---

# 18. cgo Does Not Automatically Mean You Need Pinning

This is an important engineering judgment point.

Do **not** conclude:

> "I'm using cgo, therefore I need `LockOSThread()`."

Usually you don't.

Go's runtime has extensive support for calling C code and managing OS threads.

You need explicit pinning only when the external API has a **real OS-thread-affinity requirement**.

The correct reasoning is:

```text
Does external API require thread affinity?
              │
        ┌─────┴─────┐
       NO           YES
        │             │
     Don't pin     Consider pinning
```

---

# 19. `defer` Is Usually the Right Default

Prefer:

```go
func callAPI() error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	return externalAPI()
}
```

over:

```go
func callAPI() error {
	runtime.LockOSThread()

	err := externalAPI()

	runtime.UnlockOSThread()

	return err
}
```

Because the second version can accidentally leak the lock:

```go
runtime.LockOSThread()

if err != nil {
	return err // BUG: forgot UnlockOSThread()
}

runtime.UnlockOSThread()
```

The first version is structurally safer.

---

# 20. Failure Modes

The biggest problems are usually architectural rather than syntactic.

### 1. Forgetting `UnlockOSThread`

```go
runtime.LockOSThread()

return err
```

The goroutine remains locked.

---

### 2. Pinning around slow operations

```go
runtime.LockOSThread()
defer runtime.UnlockOSThread()

http.Get(...)
```

You may unnecessarily hold thread affinity during network I/O.

---

### 3. Pinning high-cardinality goroutines

Imagine:

```text
100,000 goroutines
        │
        └── LockOSThread()
```

This is fundamentally different from ordinary Go concurrency and can create significant OS-thread pressure.

---

### 4. Assuming goroutine identity = thread identity

Normally:

```text
G1 ≠ permanently associated with M1
```

Pinning is an explicit exception.

---

# 21. Production Checklist

Before using `LockOSThread`, ask:

### Requirement

- Does the external API genuinely require OS-thread affinity?
    
- Is this documented by the API/library?
    

### Scope

- Can I isolate the thread-sensitive section?
    
- Can I use a dedicated worker instead?
    

### Blocking

- Can the pinned goroutine block?
    
- Can it perform network or disk I/O?
    
- What happens under load?
    

### Lifecycle

- Where is initialization performed?
    
- Where is cleanup performed?
    
- Must cleanup happen on the same thread?
    

### Concurrency

- How many pinned goroutines can exist?
    
- Is the number bounded?
    

### Failure

- What happens if the worker crashes/panics?
    
- Can requests queue indefinitely?
    
- Is there backpressure?
    

### Observability

- Can I measure queue length?
    
- Processing latency?
    
- Worker utilization?
    
- Errors/timeouts?
    

These questions are much more important than simply knowing the API.

---

# 22. The Principal Engineer Mental Model

Think of `LockOSThread` as an **escape hatch from Go's scheduler abstraction**.

Normally:

```text
                    Go Runtime
                       │
Application ───────► Goroutine
                       │
                       ▼
                    Scheduler
                       │
                ┌──────┼──────┐
                ▼      ▼      ▼
               M1     M2     M3
```

You should let the runtime optimize this.

With thread affinity:

```text
Application
     │
     ▼
Special constraint
     │
     ▼
LockOSThread
     │
     ▼
G ─────────────► M
```

You're telling the runtime:

> "This goroutine has an external dependency that makes normal thread migration unsafe."

That is a **real constraint**, not a performance optimization.

---

## Key Takeaways

1. **`LockOSThread()` pins the calling goroutine to its current OS thread.**
    
2. **`UnlockOSThread()` releases that affinity.**
    
3. It exists primarily for **OS-thread-affine external APIs**.
    
4. It does **not** provide mutual exclusion or synchronization.
    
5. `go func()` does not inherit the parent's thread affinity.
    
6. Do not pin merely because you're using cgo.
    
7. Avoid holding a pinned goroutine during unnecessary blocking work.
    
8. For persistent thread-affine state, a **dedicated locked worker goroutine** is often the cleanest architecture.
    
9. Always make lock/unlock lifecycle explicit, usually with `defer`.
    
10. Treat thread pinning as an **exception to Go's normal concurrency model**, not as a general-purpose concurrency technique.

---

## 🔗 References
- ⬆️ Parent: [[Goroutines]]
- 📚 Module: `Concurrency & Synchronization`
