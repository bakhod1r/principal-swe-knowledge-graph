---
title: "CGO_ENABLED (Static vs Dynamic Linking)"
tags:
  - review
  - golang
  - environment
  - principal-swe
parent: "[[Target OS & Architecture]]"
---

# `CGO_ENABLED` — Static vs Dynamic Linking

`CGO_ENABLED` controls whether the Go toolchain is allowed to use **cgo**, which lets Go code call C code and link against C libraries.

```bash
CGO_ENABLED=1
CGO_ENABLED=0
```

The important point is:

> **`CGO_ENABLED` does not directly mean "static linking" vs "dynamic linking".**  
> It controls whether cgo is available. Static/dynamic linking is a consequence of the code, linker, and libraries involved.

---

## 1. What problem does `CGO_ENABLED` solve?

Go was designed to produce self-contained binaries, but sometimes applications need to interact with native C libraries:

- SQLite implementations
    
- OpenSSL/libcrypto
    
- system libraries
    
- OS-specific APIs
    
- existing C/C++ libraries
    
- drivers that depend on native code
    

cgo provides that bridge:

```go
/*
#include <stdio.h>
*/
import "C"

func main() {
    C.puts(C.CString("hello"))
}
```

With:

```bash
CGO_ENABLED=1 go build
```

the Go compiler can compile and link the C portion.

With:

```bash
CGO_ENABLED=0 go build
```

cgo is disabled.

---

# 2. `CGO_ENABLED=0`

This is the common choice for **pure-Go applications**.

```bash
CGO_ENABLED=0 go build -o app
```

Conceptually:

```text
Go source
   │
   ▼
Go compiler
   │
   ▼
Go linker
   │
   ▼
app
```

There is no C compiler/linker dependency for the build.

### Advantages

**Portable builds**

You can often build a Linux binary without having a C toolchain installed.

**Simpler containers**

For example:

```dockerfile
FROM golang:1.25 AS build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 go build -o app ./cmd/app

FROM scratch
COPY --from=build /src/app /app

ENTRYPOINT ["/app"]
```

The resulting image can be extremely small.

**Predictable deployment**

There is less dependency on:

- glibc
    
- musl
    
- system `.so` libraries
    
- runtime C library versions
    

This is one reason Go is popular for containerized services.

---

# 3. `CGO_ENABLED=1`

When enabled, the build can involve the native toolchain:

```text
Go
 │
 ├── Go compiler
 │
 ├── cgo
 │
 └── C compiler
       │
       └── system libraries
              │
              ▼
           linker
              │
              ▼
             app
```

For example:

```bash
CGO_ENABLED=1 go build -o app
```

The binary may now depend on native shared libraries.

Check this with:

```bash
ldd ./app
```

For example:

```text
linux-vdso.so.1
libc.so.6
/lib64/ld-linux-x86-64.so.2
```

That means the executable has dynamic runtime dependencies.

---

# 4. Static vs dynamic linking

This is where an important distinction matters.

### Pure Go + `CGO_ENABLED=0`

Usually produces a binary that is **statically self-contained with respect to the Go/runtime portion**.

You can inspect it:

```bash
file ./app
```

You may see:

```text
ELF 64-bit LSB executable, x86-64, statically linked
```

This is excellent for minimal containers.

---

### cgo + `CGO_ENABLED=1`

The linker may produce:

```text
Go application
      │
      ├── Go runtime
      │
      └── libc.so
             │
             └── system dependencies
```

So the binary requires compatible runtime libraries in the target environment.

For example, this can fail:

```text
./app: error while loading shared libraries: libc.so.6
```

if the target image does not contain the required library.

---

# 5. Why `scratch` matters

Consider:

```dockerfile
FROM scratch
COPY app /app

ENTRYPOINT ["/app"]
```

There is essentially **no userspace environment**.

Therefore a dynamically linked binary can fail because there is no:

```text
libc.so.6
ld-linux-x86-64.so.2
```

A pure-Go binary built with:

```bash
CGO_ENABLED=0 go build
```

is much more suitable for:

```text
scratch
```

or similarly minimal images.

---

# 6. But `CGO_ENABLED=0` has consequences

This is not simply a "better" setting.

Some packages require cgo.

A classic example is SQLite-based software using a C implementation.

You might encounter:

```text
build constraints exclude all Go files
```

or package-specific errors because the implementation requires cgo.

Another important example is Go's `net` and `os/user` packages.

Go may have both:

```text
pure-Go implementation
```

and:

```text
cgo/system implementation
```

depending on the configuration and platform.

So blindly setting:

```bash
CGO_ENABLED=0
```

can change behavior or remove functionality you depend on.

---

# 7. Cross-compilation

This is one of the biggest practical benefits of disabling cgo.

For a pure-Go application:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app
```

You can build a Linux binary from another operating system without requiring a Linux C toolchain.

For example:

```text
macOS developer machine
        │
        │ CGO_ENABLED=0
        ▼
Linux amd64 binary
```

With:

```bash
CGO_ENABLED=1
```

cross-compilation becomes significantly more complicated because you generally need a suitable **target C compiler/toolchain and target libraries**.

For example:

```text
macOS
  │
  ├── Go compiler
  └── cross C compiler
          │
          ▼
       Linux libc
```

This is one reason teams often deliberately avoid cgo unless they actually need it.

---

# 8. Static linking with cgo

Here's the subtle part.

You **can** create statically linked binaries involving cgo, but it is considerably more complicated.

For example, you might use:

```bash
CGO_ENABLED=1 \
CC=musl-gcc \
go build \
    -ldflags '-linkmode external -extldflags "-static"' \
    -o app
```

This depends on your platform, compiler, libraries, and dependencies.

The resulting architecture can be:

```text
Go
 │
 ├── cgo
 │
 ├── C compiler
 │
 └── static libc
       │
       ▼
   static executable
```

But "static cgo binary" is **not automatically equivalent to a pure-Go static binary**.

Native libraries can introduce additional complications.

---

# 9. glibc vs musl

This becomes particularly important in containers.

### Debian/Ubuntu

Typically uses:

```text
glibc
```

### Alpine

Typically uses:

```text
musl
```

A cgo binary dynamically linked against glibc generally cannot simply be dropped into Alpine and expected to work.

For example:

```text
Build:
    Debian
      ↓
    glibc-linked binary

Runtime:
    Alpine
      ↓
    musl
```

Potential result:

```text
❌ incompatible runtime
```

This is a common production mistake.

---

# 10. A better container strategy

For a pure-Go service:

```dockerfile
FROM golang:1.25 AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /app ./cmd/server

FROM scratch

COPY --from=build /app /app

ENTRYPOINT ["/app"]
```

This gives you:

```text
small image
+
minimal runtime dependencies
+
simple deployment
+
good portability
```

However, `scratch` has **no CA certificates, timezone database, shell, users database, etc.**

So production readiness may require explicitly copying things such as:

```dockerfile
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
```

depending on what the application needs.

---

# 11. When should I use each?

|Situation|Recommendation|
|---|---|
|Pure-Go backend|`CGO_ENABLED=0`|
|Minimal container|`CGO_ENABLED=0`|
|`scratch` image|Prefer `CGO_ENABLED=0`|
|Easy cross-compilation|`CGO_ENABLED=0`|
|C library dependency|`CGO_ENABLED=1`|
|Native OS integration|Often `CGO_ENABLED=1`|
|SQLite C implementation|Usually `CGO_ENABLED=1`|
|Existing C/C++ library|`CGO_ENABLED=1`|
|Need OpenSSL through a cgo dependency|`CGO_ENABLED=1`|
|Maximum portability|Prefer avoiding cgo|
|Need native performance/library|Evaluate cgo|

---

# 12. Production mental model

Do **not** think:

```text
CGO_ENABLED=0 → static
CGO_ENABLED=1 → dynamic
```

Think:

```text
CGO_ENABLED
     │
     ▼
Can Go use C?
     │
     ├── NO
     │    └── pure-Go dependency graph
     │          └── usually highly portable/self-contained
     │
     └── YES
          └── native code possible
                │
                ├── static linking
                └── dynamic linking
```

The actual linkage depends on the **entire dependency graph + linker configuration + native toolchain + target OS/libc**.

That distinction is important at Staff/Principal level because otherwise you can make incorrect deployment assumptions.

---

## 13. How to verify instead of guessing

Don't assume your binary is static.

Inspect it:

```bash
file ./app
```

```bash
ldd ./app
```

And inspect ELF dependencies:

```bash
readelf -d ./app
```

For example, if:

```bash
ldd ./app
```

returns:

```text
not a dynamic executable
```

you have a strong indication that it is statically linked.

For a production build pipeline, make this an **artifact property you verify**, rather than an assumption based on `CGO_ENABLED`.

---

## 14. The Principal Engineer takeaway

The real design question isn't:

> "Should I set `CGO_ENABLED=0`?"

It is:

> **"Do I want native-code dependencies in this application's build and runtime dependency graph?"**

If the answer is **no**, disabling cgo usually gives you a much simpler operational model.

If the answer is **yes**, deliberately design around:

- C compiler/toolchain
    
- libc compatibility
    
- cross-compilation
    
- dynamic vs static linking
    
- container base image
    
- native library versions
    
- security updates
    
- runtime dependencies
    
- reproducible builds
    
- deployment portability
    

**Rule of thumb:** for a Go backend, start with `CGO_ENABLED=0` **if your dependency graph permits it**. Enable cgo because you have a concrete requirement—not because it is the default or because "static linking" sounds desirable.

---

## 🔗 References
- ⬆️ Parent: [[Target OS & Architecture]]
- 📚 Module: `Go Environment & Commands`
