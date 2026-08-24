---
title: "GOMODCACHE Module Cache"
tags:
  - review
  - golang
  - environment
  - principal-swe
parent: "[[Core Environment Variables]]"
---
# `GOMODCACHE`

`GOMODCACHE` is a Go environment variable that specifies **where downloaded Go modules are stored locally**.

In simple terms:

> `GOMODCACHE` = the local cache directory where Go stores downloaded module source code, versions, and related module data.

You can inspect it with:

```bash
go env GOMODCACHE
```

A typical result might be:

```text
$GOPATH/pkg/mod
```

For example:

```text
~/go/pkg/mod
```

---

## 1. Why does `GOMODCACHE` exist?

When your project imports an external module:

```go
import "github.com/google/uuid"
```

Go needs to obtain that module before compiling your application.

For example:

```bash
go mod download
```

Go downloads the required module and stores it in the module cache.

Conceptually:

```text
                    go build
                       │
                       ▼
                 go.mod / go.sum
                       │
                       ▼
                Module resolution
                       │
             ┌─────────┴─────────┐
             │                   │
       Already cached?       Not cached?
             │                   │
            YES                  NO
             │                   │
             ▼                   ▼
       Use GOMODCACHE       Download module
                                 │
                                 ▼
                            GOMODCACHE
```

This avoids downloading the same dependency repeatedly.

---

# 2. Default location

`GOMODCACHE` defaults to:

```text
$GOPATH/pkg/mod
```

For example, if:

```bash
go env GOPATH
```

returns:

```text
/home/mrb/go
```

then:

```bash
go env GOMODCACHE
```

will normally return:

```text
/home/mrb/go/pkg/mod
```

On macOS:

```text
/Users/mrb/go/pkg/mod
```

On Windows:

```text
C:\Users\mrb\go\pkg\mod
```

---

# 3. What is stored there?

Suppose your project has:

```go
require github.com/google/uuid v1.6.0
```

Go may store the module under something conceptually like:

```text
$GOMODCACHE/
└── github.com/
    └── google/
        └── uuid@v1.6.0/
```

The module cache can contain:

- downloaded module source
    
- different versions of the same module
    
- module metadata
    
- downloaded `.zip` files
    
- extracted module contents
    
- checksum-related information
    

For example:

```text
GOMODCACHE
├── github.com/
│   ├── google/
│   │   └── uuid@v1.6.0/
│   └── ...
├── golang.org/
│   └── x/
│       └── ...
└── cache/
    └── download/
```

The exact internal layout is an implementation detail and should not be relied upon by applications.

---

# 4. `GOMODCACHE` vs `GOPATH`

These are related but **not the same thing**.

Historically, `GOPATH` was the central workspace for Go development.

Modern Go modules changed that model.

Today:

```text
GOPATH
 ├── bin/
 ├── pkg/
 │   └── mod/       ← GOMODCACHE
 │
 └── ...
```

So:

```text
GOPATH
   │
   └── pkg/mod
          │
          └── GOMODCACHE
```

More precisely:

```bash
GOMODCACHE=$(go env GOPATH)/pkg/mod
```

is the normal default relationship.

### Important distinction

`GOPATH` is a broader Go environment setting.

`GOMODCACHE` specifically controls the **module download cache**.

---

# 5. Changing `GOMODCACHE`

You can configure it with:

```bash
go env -w GOMODCACHE=/opt/go/modcache
```

Then verify:

```bash
go env GOMODCACHE
```

You can also configure it temporarily:

```bash
GOMODCACHE=/tmp/go-mod-cache go build ./...
```

This is useful for isolated builds and CI environments.

---

# 6. Why would you change it?

There are several legitimate reasons.

### CI/CD

You might want a dedicated persistent cache:

```text
/ci/cache/go-mod
```

Then multiple builds can reuse downloaded dependencies.

Without caching:

```text
Build
  ↓
download dependencies
  ↓
compile
```

With caching:

```text
Build
  ↓
check module cache
  ↓
reuse dependencies
  ↓
compile
```

This can significantly reduce build time and network traffic.

---

### Docker builds

A containerized build might use a dedicated module-cache location.

For example:

```dockerfile
ENV GOMODCACHE=/go/pkg/mod
```

The important part isn't the particular path; it's that the cache can potentially be mounted or persisted between builds.

With BuildKit, you can go further and cache the module directory between builds.

---

### Multiple disks

If the default home directory has limited space:

```text
/home/mrb/go/pkg/mod
```

could be moved to a larger filesystem:

```text
/data/go/modcache
```

---

# 7. Module cache is not your project's dependencies directory

A common misconception is:

```text
GOMODCACHE = project's vendor directory
```

No.

For a project:

```text
my-service/
├── go.mod
├── go.sum
├── cmd/
└── internal/
```

the dependencies are normally resolved from the module cache.

```text
my-service/
        │
        ├── go.mod
        └── go.sum
               │
               ▼
          Go module system
               │
               ▼
         GOMODCACHE
```

If you explicitly use vendoring:

```bash
go mod vendor
```

you get:

```text
my-service/
├── go.mod
├── go.sum
├── vendor/
│   └── ...
└── ...
```

That is a different mechanism.

---

# 8. `GOMODCACHE` and `go.sum`

These have different responsibilities.

### `go.mod`

Declares module requirements:

```go
require (
    github.com/google/uuid v1.6.0
)
```

### `go.sum`

Records cryptographic checksums for module content.

### `GOMODCACHE`

Stores the actual downloaded module data locally.

Mental model:

```text
go.mod
   │
   │ "I need X@v1.6.0"
   ▼
Module resolution
   │
   ├── go.sum → integrity verification
   │
   └── GOMODCACHE → local module data
```

---

# 9. `GOMODCACHE` and `GOPROXY`

These are also easy to confuse.

`GOPROXY` controls **where Go obtains modules from**.

`GOMODCACHE` controls **where Go stores them locally**.

```text
                Go build
                   │
                   ▼
              Need module
                   │
            ┌──────┴──────┐
            │             │
       GOMODCACHE      GOPROXY
        "Where?"       "From where?"
            │             │
            ▼             ▼
       local cache    module proxy
```

For example:

```bash
go env GOPROXY
```

might return:

```text
https://proxy.golang.org,direct
```

while:

```bash
go env GOMODCACHE
```

might return:

```text
/home/mrb/go/pkg/mod
```

---

# 10. Useful commands

Inspect:

```bash
go env GOMODCACHE
```

Download dependencies:

```bash
go mod download
```

Clean the module cache:

```bash
go clean -modcache
```

Check module dependencies:

```bash
go list -m all
```

Show Go environment:

```bash
go env
```

---

# 11. Production / CI perspective

For a production-grade CI pipeline, think about the module cache as a **build-performance optimization**, not as application state.

Good:

```text
Source
  ↓
go.mod/go.sum
  ↓
Restored module cache
  ↓
go build
  ↓
artifact
```

Don't treat the cache as something that must be preserved for correctness.

If the cache disappears, Go should be able to download dependencies again.

Therefore:

> **Module cache is disposable state. `go.mod` and `go.sum` are the reproducible dependency inputs.**

This distinction is important when designing CI/CD systems.

---

## Key takeaway

```text
GOMODCACHE
    │
    ├── What?   Local Go module download cache
    ├── Default? $GOPATH/pkg/mod
    ├── Used by? Go module system
    ├── Persistent? Usually yes for performance
    ├── Required for correctness? No
    └── Can be changed? Yes
```

The most important mental model is:

**`go.mod` says what you need → `GOPROXY` says where to obtain it → `GOMODCACHE` says where to cache it locally → `go.sum` verifies module content.**

---

## 🔗 References
- ⬆️ Parent: [[Core Environment Variables]]
- 📚 Module: `Go Environment & Commands`
