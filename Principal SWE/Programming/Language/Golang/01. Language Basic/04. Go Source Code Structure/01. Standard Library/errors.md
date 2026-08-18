---
title: errors Package
tags:
  - golang
  - standard-library
  - errors
parent: "[[Standard Library]]"
---

# `errors`

Error construction, wrapping, and inspection. Small package, load-bearing
semantics.

## 1. Creating

```go
var ErrNotFound = errors.New("not found")          // sentinel — package level
fmt.Errorf("load user %d: %w", id, err)            // wrapping
errors.Join(err1, err2)                            // Go 1.20 — multiple errors
```

## 2. Inspecting

```go
if errors.Is(err, ErrNotFound) { ... }             // identity, through wrapping

var pathErr *fs.PathError
if errors.As(err, &pathErr) {                      // type, through wrapping
    log.Println(pathErr.Path)
}
```

**Never compare with `==`** once wrapping is in play — `err == ErrNotFound` fails
the moment someone adds context with `%w`.

## 3. The Wrapping Contract

```go
return fmt.Errorf("query users: %w", err)   // wrapped — Is/As see through
return fmt.Errorf("query users: %v", err)   // flattened — chain broken
```

`%w` preserves the chain; `%v` deliberately hides it. Hiding is the right choice
when the underlying error is an implementation detail you do not want callers
matching on.

## 4. Custom Errors

```go
type ValidationError struct{ Field string; Err error }

func (e *ValidationError) Error() string { return e.Field + ": " + e.Err.Error() }
func (e *ValidationError) Unwrap() error { return e.Err }
```

Implementing `Unwrap` is what makes `errors.Is`/`As` traverse your type.

## 5. Gotchas

- A `nil` error inside a non-nil interface value is non-nil — the classic
  typed-nil trap. Return the bare `nil`, not a nil concrete pointer.
- `errors.Join` returns an error whose `Unwrap() []error` form is understood by
  `Is`/`As` but not by `%w` in older code.
- Sentinel errors are part of your **public API**; changing one is a breaking
  change.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
