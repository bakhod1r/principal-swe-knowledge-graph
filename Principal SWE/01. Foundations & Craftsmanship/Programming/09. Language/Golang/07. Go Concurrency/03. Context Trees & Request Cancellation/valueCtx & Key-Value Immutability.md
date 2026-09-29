---
title: "valueCtx & Key-Value Immutability"
tags:

  - golang
  - concurrency
  - principal-swe
parent: "[[Context Trees & Request Cancellation]]"
---
# `valueCtx` & Key-Value Immutability in Go

`context.WithValue` is often described as a simple key-value store, but that mental model is incomplete.

The important concept is:

> **`context.Context` is an immutable linked chain of scopes, and `valueCtx` adds one immutable key-value binding to that chain.**

This design is what makes context values safe to share across goroutines without locks.

---

## 1. The Problem

Suppose a request enters your service:

```text
HTTP Request
    │
    ├── request ID
    ├── trace ID
    ├── authenticated user
    └── logger fields
```

Many functions need access to some of this request-scoped information:

```go
func Handle(ctx context.Context) {
    service(ctx)
}

func service(ctx context.Context) {
    repository(ctx)
}

func repository(ctx context.Context) {
    // needs request ID
}
```

Passing every value explicitly becomes noisy:

```go
service(requestID, traceID, userID, ...)
```

`context.Context` provides a propagation mechanism:

```go
ctx = context.WithValue(ctx, requestIDKey, requestID)
```

The crucial question is:

> **How can this shared context be safely accessed concurrently without a mutex?**

The answer is **immutability + persistent chaining**.

---

# 2. Mental Model

Think of a context as a linked list:

```text
valueCtx
   │
   ├── key   = "requestID"
   ├── val   = "abc-123"
   │
   └── parent
         │
         ▼
      valueCtx
         │
         ├── key = "userID"
         ├── val = 42
         │
         └── parent
               │
               ▼
          backgroundCtx
```

Each `WithValue` creates a **new node**.

It does **not** modify the existing context.

Conceptually:

```go
ctx1 := context.Background()

ctx2 := context.WithValue(ctx1, key1, value1)

ctx3 := context.WithValue(ctx2, key2, value2)
```

The structure becomes:

```text
ctx3
 │
 ▼
[key2=value2]
 │
 ▼
[key1=value1]
 │
 ▼
background
```

`ctx1` remains unchanged.

---

# 3. Internally: `valueCtx`

Conceptually, Go's implementation looks like:

```go
type valueCtx struct {
    Context
    key, val any
}
```

The important part is:

```go
Context
```

That is the parent.

So:

```go
context.WithValue(parent, key, value)
```

effectively constructs:

```text
new valueCtx
    key = key
    val = value
    parent = parent
```

It does not mutate `parent`.

---

# 4. Why Immutability Matters

Consider:

```go
ctx1 := context.Background()

ctx2 := context.WithValue(ctx1, key, "A")
ctx3 := context.WithValue(ctx2, key, "B")
```

You now have:

```text
ctx3
 │
 └── key = B
       │
       ▼
     ctx2
       │
       └── key = A
             │
             ▼
           ctx1
```

`ctx2` still means:

```text
key → A
```

while `ctx3` means:

```text
key → B
```

There is no mutation from:

```go
ctx2
```

to:

```go
ctx3
```

This is extremely important.

---

# 5. Persistent Data Structure

This resembles a **persistent data structure**.

Instead of modifying:

```text
A
```

you create:

```text
B
│
└── A
```

The old version remains valid.

For example:

```go
base := context.Background()

ctxA := context.WithValue(base, key, "A")
ctxB := context.WithValue(ctxA, key, "B")
```

Both contexts coexist:

```text
ctxB ──► [key=B] ──► ctxA ──► [key=A] ──► base
                         ▲
                         │
                         └── immutable
```

This allows multiple goroutines to safely share `ctxA`.

---

# 6. `Value()` Lookup

When you call:

```go
ctxB.Value(key)
```

the lookup walks toward the root.

Conceptually:

```text
ctxB
 │
 ├── key == requested key?
 │      YES → return value
 │
 └── NO
      │
      ▼
     ctxA
      │
      ├── key == requested key?
      │
      └── ...
```

So lookup is approximately:

```text
O(depth)
```

where `depth` is the number of relevant context layers traversed.

This is why context should **not** become a general-purpose key-value database.

---

# 7. Shadowing

A child context can override a parent's value:

```go
ctx1 := context.WithValue(parent, key, "production")

ctx2 := context.WithValue(ctx1, key, "staging")
```

Then:

```go
ctx1.Value(key) // "production"
ctx2.Value(key) // "staging"
```

The parent was not changed.

This is **shadowing**, not mutation.

Think lexical scope:

```text
outer scope:
    environment = production

inner scope:
    environment = staging
```

The inner value hides the outer value.

---

# 8. Why This Is Goroutine-Safe

Suppose:

```go
ctx := context.WithValue(
    context.Background(),
    requestIDKey,
    "abc",
)
```

Then:

```go
go worker1(ctx)
go worker2(ctx)
go worker3(ctx)
```

All goroutines can read:

```go
ctx.Value(requestIDKey)
```

without a mutex.

Why?

Because they are all reading the same immutable context structure.

There is no operation like:

```go
ctx.values[key] = value
```

There is no shared mutable map inside `valueCtx`.

Instead:

```text
                 ┌── worker 1
                 │
immutable ctx ───┼── worker 2
                 │
                 └── worker 3
```

This is one of the fundamental reasons contexts can be safely propagated across goroutines.

---

# 9. Important Distinction: Context Immutable ≠ Value Immutable

This is a subtle but **very important** point.

Consider:

```go
type User struct {
    Name string
}
```

Then:

```go
user := &User{Name: "Alice"}

ctx := context.WithValue(
    context.Background(),
    userKey,
    user,
)
```

The **context binding** is immutable.

But the object stored inside it is not necessarily immutable.

You can still do:

```go
user.Name = "Bob"
```

Now:

```go
ctx.Value(userKey).(*User).Name
```

returns:

```text
Bob
```

So:

```text
valueCtx immutability
        ≠
stored value immutability
```

This distinction matters enormously in concurrent code.

---

# 10. The Real Concurrency Boundary

This is safe:

```go
ctx := context.WithValue(
    context.Background(),
    key,
    "request-123",
)
```

because strings are immutable values.

This can be dangerous:

```go
m := map[string]string{
    "role": "admin",
}

ctx := context.WithValue(
    context.Background(),
    key,
    m,
)
```

If multiple goroutines mutate:

```go
m["role"] = "user"
```

you have a normal concurrent-map problem.

`context` does **not** make the map safe.

The ownership model is:

```text
Context node
    │
    └── immutable binding
              │
              └── referenced object
                       │
                       └── may still be mutable
```

---

# 11. Key Immutability and Comparability

`context.WithValue` requires the key to be comparable.

Good:

```go
type requestIDKey struct{}

var requestIDCtxKey requestIDKey
```

Then:

```go
ctx = context.WithValue(ctx, requestIDCtxKey, requestID)
```

Even better, commonly:

```go
type requestIDKey struct{}
```

and use:

```go
var requestIDKey requestIDKey
```

Why not:

```go
context.WithValue(ctx, "requestID", id)
```

Because string keys can collide across packages.

Imagine:

```go
package auth

context.WithValue(ctx, "user", user)
```

and:

```go
package billing

context.WithValue(ctx, "user", billingUser)
```

Both use the same comparable key:

```text
"user"
```

One package can accidentally shadow another package's value.

A package-private key type avoids this.

---

# 12. Type-Based Namespacing

A good pattern:

```go
type requestIDKey struct{}
```

Because the type itself creates a namespace.

For example:

```go
package tracing

type traceIDKey struct{}

var traceIDCtxKey traceIDKey
```

Another package:

```go
package auth

type userIDKey struct{}

var userIDCtxKey userIDKey
```

Even if both variables have similar names, their types are different:

```text
tracing.traceIDKey
auth.userIDKey
```

No accidental collision.

---

# 13. Why `WithValue` Returns a New Context

This API:

```go
ctx = context.WithValue(ctx, key, value)
```

may initially look strange.

Why not:

```go
ctx.SetValue(key, value)
```

?

Because that would require mutation.

Imagine:

```go
ctx := ...
```

and:

```go
go workerA(ctx)
go workerB(ctx)
```

If one goroutine could mutate:

```go
ctx.SetValue(...)
```

then another goroutine could observe the change unexpectedly.

Instead:

```go
ctx2 := context.WithValue(ctx, key, value)
```

means:

```text
old ctx
   │
   ├── remains unchanged
   │
new ctx
   │
   └── contains additional binding
```

This gives contexts **value-like semantics**.

---

# 14. Context as a Scope Chain

A useful mental model is:

> **`Context` behaves more like lexical scope than like a mutable map.**

For example:

```go
root
 │
 ├── requestID = 123
 │
 └── child
      │
      ├── userID = 42
      │
      └── child
           │
           └── tenantID = 7
```

The deepest context sees:

```text
requestID
userID
tenantID
```

while the intermediate context sees:

```text
requestID
userID
```

This is analogous to nested scopes in a programming language.

---

# 15. A Common Misconception

This is **not**:

```text
Context
  └── map[key]value
```

It is closer to:

```text
Context
   │
   ▼
[value]
   │
   ▼
Context
   │
   ▼
[value]
   │
   ▼
Context
```

That implementation choice explains several properties simultaneously:

- immutability
    
- inheritance
    
- shadowing
    
- cheap creation
    
- safe sharing
    
- linear lookup
    

This is the mental model you should retain.

---

# 16. Performance

Creating a value context is cheap.

Conceptually:

```go
ctx2 := &valueCtx{
    Context: ctx,
    key:     key,
    val:     value,
}
```

So creation is approximately:

```text
O(1)
```

Memory:

```text
O(1)
```

per added value layer.

Lookup is different:

```text
Value()
   ↓
valueCtx
   ↓
parent
   ↓
parent
   ↓
...
```

Worst-case:

```text
O(depth)
```

Therefore:

### Good

```go
ctx = context.WithValue(ctx, requestIDKey, requestID)
ctx = context.WithValue(ctx, traceIDKey, traceID)
```

### Bad

Using context as a giant application-wide state container:

```go
ctx = context.WithValue(ctx, key1, ...)
ctx = context.WithValue(ctx, key2, ...)
ctx = context.WithValue(ctx, key3, ...)
...
ctx = context.WithValue(ctx, key1000, ...)
```

That is a design smell.

---

# 17. What Context Values Are Actually For

The official intended use is:

> **request-scoped data that transits API boundaries and processes, not arbitrary optional parameters.**

Good examples:

```text
request ID
trace/span information
authentication information
deadline/cancellation metadata
request-scoped logging metadata
```

Bad examples:

```text
database connection
repository
configuration
service objects
business state
large caches
optional function parameters
```

Instead of:

```go
func GetUser(ctx context.Context) (*User, error) {
    db := ctx.Value(dbKey).(*sql.DB)
    ...
}
```

prefer:

```go
func GetUser(ctx context.Context, db *sql.DB) (*User, error)
```

The second version makes the dependency explicit.

---

# 18. A Subtle Architectural Insight

`context.Context` gives you **implicit dependency propagation**.

That is useful for cross-cutting request metadata.

But excessive use creates **hidden dependencies**.

Compare:

```go
func Authorize(ctx context.Context) error
```

with:

```go
func Authorize(ctx context.Context, user *User) error
```

The first hides where `user` comes from.

The second makes the dependency visible.

Therefore:

> **Use context for request-scoped infrastructure metadata, not to avoid designing explicit APIs.**

This is a Staff/Principal-level distinction.

---

# 19. Context Tree vs Context Chain

There are actually two useful ways to visualize it.

### Structural model

```text
root
 │
 └── valueCtx
      │
      └── valueCtx
           │
           └── cancelCtx
                │
                └── timerCtx
```

This is a **chain** from each context toward its parent.

But operationally, cancellation can propagate through a **tree**:

```text
                  parent
                 /      \
             child A   child B
              /   \       \
            A1    A2      B1
```

So remember:

```text
Value lookup → parent chain

Cancellation → propagation tree
```

These are related but different concepts.

---

# 20. Immutability Enables Structural Sharing

This is the deeper computer-science concept.

Suppose:

```go
base := context.Background()

a := context.WithValue(base, keyA, valueA)
b := context.WithValue(a, keyB, valueB)
c := context.WithValue(a, keyC, valueC)
```

The structure is:

```text
             a
            / \
           b   c
```

More precisely:

```text
b ──► [keyB] ──► a ──► [keyA] ──► base

c ──► [keyC] ──► a ──► [keyA] ──► base
```

`b` and `c` share the same immutable suffix:

```text
[keyA] → base
```

This is **structural sharing**.

No deep copy is required.

That's one of the elegant properties of persistent data structures.

---

# 21. Production Pattern

A clean pattern is to hide context key details inside a package:

```go
package requestctx

import "context"

type requestIDKey struct{}

func WithRequestID(ctx context.Context, id string) context.Context {
    return context.WithValue(ctx, requestIDKey{}, id)
}

func RequestID(ctx context.Context) (string, bool) {
    id, ok := ctx.Value(requestIDKey{}).(string)
    return id, ok
}
```

Usage:

```go
ctx = requestctx.WithRequestID(ctx, "req-123")

id, ok := requestctx.RequestID(ctx)
if !ok {
    // missing
}
```

This is better than exposing:

```go
context.WithValue(ctx, "requestID", ...)
```

throughout the application.

---

# 22. But Don't Over-Abstract

You don't need a huge framework:

```text
ContextManager
ContextValueProvider
ContextRegistry
ContextFactory
ContextAccessor
ContextMiddleware
...
```

That is usually overengineering.

A small package with:

```go
WithX(...)
X(...)
```

is generally enough.

---

# 23. Key Takeaways

The core model:

```text
context.WithValue(parent, key, value)

             creates

              valueCtx
             /        \
          key          val
           |
         parent
```

Properties:

|Property|Meaning|
|---|---|
|Immutable|Existing context never changes|
|Persistent|Old versions remain valid|
|O(1) creation|Add one node|
|O(depth) lookup|Walk parent chain|
|Shadowing|Child value overrides parent|
|Structural sharing|Children share immutable parents|
|Goroutine-safe|Context structure can be shared|
|Value not necessarily immutable|Stored objects may still mutate|
|Comparable key|Required for lookup|
|Request-scoped|Intended use|

### The most important mental model

Don't think:

```text
Context = concurrent map
```

Think:

```text
Context = immutable scope chain
```

And don't confuse:

```text
immutable context binding
```

with:

```text
immutable object stored inside the binding
```

That distinction explains both **why `context.Context` is concurrency-friendly** and **why putting mutable shared state inside it can still create races**.

---

## 🔗 References
- ⬆️ Parent: [[Context Trees & Request Cancellation]]
- 📚 Module: `Concurrency & Synchronization`
