---
title: "libc vs musl in Scratch Containers"
tags:
  - review
  - golang
  - environment
  - principal-swe
parent: "[[Target OS & Architecture]]"
---
# libc vs musl in Scratch Containers

## 1. Problem

When building a Go application into a `scratch` container, you often want the smallest possible image:

```dockerfile
FROM scratch
COPY app /app
ENTRYPOINT ["/app"]
```

But `scratch` contains **no userspace libraries at all**.

That creates an important question:

> If my Go binary needs `libc`, where does that library come from?

The answer depends primarily on **how the binary was linked**.

---

# 2. What is `libc`?

`libc` means **C standard library**.

It provides common userspace functionality such as:

- memory allocation
    
- string operations
    
- file operations
    
- networking-related APIs
    
- threads
    
- process management
    
- DNS resolver interfaces
    
- system-call wrappers
    
- POSIX functionality
    

On Linux, two important implementations are:

```text
glibc
musl
```

They both implement the C/POSIX userspace interface, but they have different designs and compatibility characteristics.

### glibc

Common on:

```text
Ubuntu
Debian
Fedora
RHEL
CentOS
```

### musl

Common on:

```text
Alpine Linux
```

Conceptually:

```text
Application
    │
    ▼
   libc
    │
    ▼
Linux kernel
```

With glibc:

```text
Application
    │
    ▼
glibc
    │
    ▼
Linux kernel
```

With musl:

```text
Application
    │
    ▼
musl
    │
    ▼
Linux kernel
```

---

# 3. What is `scratch`?

Docker's `scratch` is essentially an **empty filesystem**.

It does not contain:

```text
/bin/sh
/lib
/usr/lib
/etc/passwd
/etc/resolv.conf
CA certificates
glibc
musl
```

For example:

```dockerfile
FROM scratch

COPY app /app

ENTRYPOINT ["/app"]
```

The only things inside are what you explicitly copy.

Therefore:

> A dynamically linked binary cannot normally run in `scratch` unless all required runtime libraries are also provided.

---

# 4. Static vs Dynamic Linking

This is the most important concept.

Suppose:

```text
app
 │
 ├── libc.so
 ├── libpthread.so
 └── other shared libraries
```

That's **dynamic linking**.

The binary expects the operating system to locate those libraries at runtime.

A statically linked binary instead looks more like:

```text
app
 │
 ├── application
 ├── libc implementation
 ├── runtime dependencies
 └── everything required
```

There is no external `libc.so` dependency.

Therefore:

```text
Dynamic binary + scratch
        ↓
usually fails

Static binary + scratch
        ↓
usually works
```

---

# 5. Why Go makes this interesting

Go can often produce binaries that don't depend on libc.

For example:

```bash
CGO_ENABLED=0 go build -o app .
```

The result is typically a statically linked Go binary.

Then:

```dockerfile
FROM scratch

COPY app /app

ENTRYPOINT ["/app"]
```

can work.

Architecture:

```text
┌─────────────────────────┐
│       Go application    │
│                         │
│      Go runtime         │
│                         │
│      Go networking      │
│      Go DNS resolver    │
└────────────┬────────────┘
             │
             ▼
        Linux kernel
```

No libc is required.

---

# 6. Where CGO changes everything

Consider:

```go
package main

/*
#include <stdio.h>
*/
import "C"

func main() {
    println("hello")
}
```

Now CGO is involved.

The compiler may produce:

```text
Go application
      │
      ▼
    CGO
      │
      ▼
    libc
      │
      ▼
 Linux kernel
```

If you build against glibc:

```text
app
 └── libc.so.6
```

the resulting binary may be dynamically linked.

Putting it into `scratch`:

```dockerfile
FROM scratch
COPY app /app
ENTRYPOINT ["/app"]
```

can result in:

```text
/app: not found
```

or an error indicating that the required dynamic loader/library cannot be found.

The confusing part is that:

> The executable file exists, but its interpreter does not.

---

# 7. The Dynamic Loader Is Important

A dynamically linked ELF binary contains an **ELF interpreter**.

For example, a glibc-linked executable might require:

```text
/lib64/ld-linux-x86-64.so.2
```

The loader then loads:

```text
libc.so.6
libpthread.so.0
libm.so.6
...
```

So the runtime dependency chain becomes:

```text
/app
  │
  ▼
ld-linux-x86-64.so.2
  │
  ├── libc.so.6
  ├── libm.so.6
  └── other libraries
```

But `scratch` contains none of these.

Hence:

```text
scratch
├── app
└── nothing else
```

→ failure.

---

# 8. glibc vs musl

A simplified comparison:

|Property|glibc|musl|
|---|---|---|
|Common distro|Debian/Ubuntu/RHEL|Alpine|
|Size|Larger|Smaller|
|Static linking|Possible|Generally friendly|
|Compatibility|Excellent Linux ecosystem compatibility|Strong, but different|
|NSS ecosystem|Extensive|Different implementation|
|Typical minimal containers|Less common|Very common|
|Alpine default|No|Yes|

The key point:

> musl is not merely a "smaller glibc."

They are different libc implementations with different behavior, ABI details, tooling, and ecosystem compatibility.

---

# 9. Why musl is popular for containers

Consider:

```dockerfile
FROM alpine
```

Alpine uses musl.

A dynamically linked application can therefore run against:

```text
/app
/lib/ld-musl-x86_64.so.1
/lib/libc.musl-x86_64.so.1
```

The image remains small while still providing a userspace environment.

Architecture:

```text
┌──────────────────────────┐
│ Alpine                   │
│                          │
│  Application             │
│       │                  │
│       ▼                  │
│     musl                  │
│       │                  │
│       ▼                  │
│ Linux kernel             │
└──────────────────────────┘
```

This is one reason Alpine became popular for small containers.

---

# 10. `scratch` vs Alpine

A useful mental model:

### `scratch`

```text
┌─────────────────────┐
│ Your binary         │
├─────────────────────┤
│ Nothing             │
└─────────────────────┘
```

### Alpine

```text
┌─────────────────────┐
│ Your application    │
├─────────────────────┤
│ musl                │
│ BusyBox              │
│ certificates        │
│ basic filesystem    │
└─────────────────────┘
```

Therefore:

```text
scratch → maximum minimalism
Alpine  → minimal Linux userspace
```

---

# 11. Three Common Go Deployment Models

## Model A — Pure Go + scratch

Build:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app .
```

Docker:

```dockerfile
FROM scratch

COPY app /app

ENTRYPOINT ["/app"]
```

Architecture:

```text
Go
 │
 └── static executable
          │
          ▼
       kernel
```

This is often the simplest solution.

---

# 12. Model B — CGO + Alpine/musl

Build inside Alpine:

```dockerfile
FROM golang:alpine AS build

RUN apk add --no-cache build-base

WORKDIR /src

COPY . .

RUN CGO_ENABLED=1 go build -o app .
```

Runtime:

```dockerfile
FROM alpine

COPY --from=build /src/app /app

ENTRYPOINT ["/app"]
```

Now:

```text
app
 │
 ▼
musl
 │
 ▼
kernel
```

This works because the runtime image provides musl.

---

# 13. Model C — CGO + static musl + scratch

This is the interesting approach.

You can compile CGO code against musl statically:

```text
Go
 │
 ▼
CGO
 │
 ▼
musl
 │
 ▼
static executable
```

Then:

```dockerfile
FROM scratch

COPY app /app

ENTRYPOINT ["/app"]
```

The resulting binary contains the required libc implementation.

Conceptually:

```text
┌──────────────────────────┐
│ /app                     │
│                          │
│ Go runtime               │
│ Go application           │
│ CGO code                 │
│ musl                     │
│ libc functionality       │
└────────────┬─────────────┘
             │
             ▼
        Linux kernel
```

This gives you:

> CGO functionality + scratch-sized runtime filesystem.

But there is a major caveat.

---

# 14. Static Linking Does Not Mean "No Runtime Dependencies"

This is a common mistake.

A binary can be statically linked but still require external runtime data.

For example:

```text
DNS
TLS certificates
timezone database
/etc/passwd
/etc/group
```

Your executable may be self-contained from a **code/library** perspective while still expecting files from the filesystem.

For example:

```text
/etc/ssl/certs/ca-certificates.crt
/etc/resolv.conf
/etc/hosts
/usr/share/zoneinfo
```

This distinction is critical:

```text
Static linking
      ≠
Self-contained filesystem
```

---

# 15. DNS: One of the Most Important Go Gotchas

Go has its own DNS resolver:

```text
Go net package
      │
      ▼
pure Go resolver
```

But when CGO is enabled, Go may use the system resolver depending on configuration and environment.

Conceptually:

```text
CGO_ENABLED=0

Go
 │
 └── net
      │
      └── Go resolver
```

versus:

```text
CGO_ENABLED=1

Go
 │
 └── net
      │
      └── libc/system resolver
```

This is why DNS behavior can change between:

```text
CGO_ENABLED=0
```

and:

```text
CGO_ENABLED=1
```

---

# 16. `netgo`

Go provides the `netgo` build tag to prefer the Go DNS resolver.

Conceptually:

```bash
go build -tags netgo
```

Then:

```text
application
     │
     ▼
Go resolver
     │
     ▼
DNS configuration
     │
     ▼
network
```

This can reduce libc dependency, but you still need to understand the exact behavior of your application and environment.

---

# 17. TLS Certificates in Scratch

Another classic production failure:

```go
http.Get("https://example.com")
```

works in:

```text
Ubuntu
Alpine
```

but fails in:

```text
scratch
```

with a certificate-related error.

Why?

Because:

```text
scratch
└── /etc/ssl/certs
        └── empty/missing
```

Your application may need CA certificates.

A common multi-stage pattern is:

```dockerfile
FROM alpine AS certs

RUN apk add --no-cache ca-certificates

FROM scratch

COPY --from=certs /etc/ssl/certs/ca-certificates.crt \
                  /etc/ssl/certs/ca-certificates.crt

COPY app /app

ENTRYPOINT ["/app"]
```

Now:

```text
scratch
├── /app
└── /etc/ssl/certs/ca-certificates.crt
```

---

# 18. `os/user` Is Another Gotcha

Some Go functionality interacts with the system user database.

For example:

```go
user.Current()
```

behavior can depend on whether CGO/system facilities are involved.

In a `scratch` image:

```text
/etc/passwd
```

doesn't exist by default.

You can explicitly create/copy one:

```dockerfile
COPY passwd /etc/passwd
```

or avoid requiring it if your application doesn't need user lookup.

Again:

> `scratch` makes you responsible for every runtime assumption.

---

# 19. Time Zones

Your application may use:

```go
time.LoadLocation("Asia/Tashkent")
```

A minimal runtime may not contain:

```text
/usr/share/zoneinfo
```

Depending on how your binary is built and configured, Go can embed timezone data using:

```bash
-importcfg
```

or, more commonly, the `time/tzdata` package:

```go
import _ "time/tzdata"
```

Then timezone data becomes part of the executable.

This is another example of:

```text
static binary
    ≠
complete runtime environment
```

---

# 20. How to Inspect Your Binary

Do not guess.

Use:

```bash
file app
```

Example:

```text
app: ELF 64-bit LSB executable, x86-64,
statically linked
```

or:

```text
app: ELF 64-bit LSB executable, x86-64,
dynamically linked
```

Then:

```bash
ldd app
```

For a dynamically linked binary you might see:

```text
linux-vdso.so.1
libc.so.6
/lib64/ld-linux-x86-64.so.2
```

For static binaries:

```text
not a dynamic executable
```

You can also inspect the ELF interpreter:

```bash
readelf -l app | grep interpreter
```

For glibc:

```text
[Requesting program interpreter:
 /lib64/ld-linux-x86-64.so.2]
```

For musl:

```text
/lib/ld-musl-x86_64.so.1
```

This is much better than trial-and-error inside containers.

---

# 21. The Production Decision Tree

When choosing between `scratch`, Alpine, and a glibc-based image, think about **actual dependencies**, not image size alone.

```text
                 Does application need CGO?
                         │
                ┌────────┴────────┐
                │                 │
               No                Yes
                │                 │
                ▼                 ▼
        Can use pure Go?    Need libc compatibility?
                │                 │
                ▼           ┌─────┴─────┐
             scratch       Yes          No
                │           │             │
                │           ▼             ▼
                │        glibc         musl
                │       runtime       possible
                │
                ▼
        Check runtime data:
        CA certs
        DNS
        timezone
        passwd
```

---

# 22. When Should You Use `scratch`?

Use it when:

- your binary can be statically linked
    
- runtime dependencies are understood
    
- you don't need a shell
    
- you don't need package management
    
- you can explicitly provide required files
    
- you value minimal attack surface
    

Typical example:

```text
Go HTTP API
CGO_ENABLED=0
static binary
CA certificates copied
non-root user configured
```

This is an excellent `scratch` candidate.

---

# 23. When Should You NOT Use `scratch`?

Avoid it when your application requires:

- complex CGO dependencies
    
- dynamic shared libraries
    
- glibc-specific behavior
    
- native libraries
    
- shell-based operational tooling
    
- system packages
    
- complex NSS configuration
    
- runtime debugging utilities
    

For example:

```text
Go
 └── CGO
      └── proprietary native library
           └── glibc
```

Trying to force this into `scratch` can create unnecessary operational complexity.

A small Debian/Ubuntu runtime may be the better engineering decision.

---

# 24. Security Perspective

`scratch` has an important security advantage:

```text
fewer files
     ↓
fewer packages
     ↓
smaller attack surface
```

There is no:

```text
shell
package manager
standard Unix utilities
extra shared libraries
```

But don't confuse:

```text
small image
```

with:

```text
secure application
```

You still need:

- non-root execution
    
- dependency scanning
    
- secret management
    
- TLS
    
- input validation
    
- least privilege
    
- read-only filesystem where appropriate
    
- seccomp/AppArmor/container runtime controls
    
- vulnerability management
    

Minimality is **one security layer**, not the whole security model.

---

# 25. Operational Trade-off

The biggest downside of `scratch` is often not technical compatibility.

It's **debuggability**.

Imagine the container crashes and you enter it:

```bash
docker exec -it container sh
```

With:

```text
scratch
```

you get:

```text
exec: "sh": executable file not found
```

There is no shell.

So production debugging relies more heavily on:

```text
logs
metrics
traces
pprof
debug endpoints
ephemeral debug containers
core dumps
external observability
```

This is an important Staff+/Principal-level trade-off:

> Don't optimize image size at the expense of operability without measuring the operational cost.

---

# 26. A Good Go `scratch` Pattern

A common production-oriented structure:

```dockerfile
FROM golang:1.XX AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build \
      -trimpath \
      -ldflags="-s -w" \
      -o /out/app ./cmd/server

FROM alpine AS certs

RUN apk add --no-cache ca-certificates

FROM scratch

COPY --from=certs /etc/ssl/certs/ca-certificates.crt \
                  /etc/ssl/certs/ca-certificates.crt

COPY --from=build /out/app /app

USER 65532:65532

ENTRYPOINT ["/app"]
```

The important part isn't memorizing this Dockerfile.

The mental model is:

```text
Build environment
      │
      │ compile
      ▼
Static executable
      │
      │ copy only required runtime data
      ▼
Minimal runtime
      │
      ▼
Linux kernel
```

---

# 27. The Deeper Mental Model

Think about a container as **filesystem + process**, not as a VM.

The kernel comes from the host:

```text
                 Host Linux Kernel
                       ▲
                       │
       ┌───────────────┴───────────────┐
       │                               │
   scratch container              Alpine container
       │                               │
    /app                          /app + musl
```

A container image doesn't contain the Linux kernel.

But the application may still need userspace components such as:

```text
libc
dynamic loader
CA certificates
timezone database
NSS configuration
passwd/group
```

Therefore:

> The kernel is shared, but userspace dependencies belong to the container filesystem unless statically incorporated into the executable.

That is the key concept behind `libc vs musl vs scratch`.

---

# 28. Principal Engineer Takeaway

Don't ask:

> "Should I use Alpine or scratch?"

Ask:

> **"What is the complete runtime dependency closure of my executable?"**

For example:

```text
Application
   │
   ├── Go runtime
   ├── CGO?
   │    └── libc?
   │          ├── glibc
   │          └── musl
   │
   ├── dynamic loader?
   ├── CA certificates?
   ├── DNS configuration?
   ├── timezone data?
   ├── passwd/group?
   └── native libraries?
```

Then choose the smallest runtime that satisfies those dependencies **without introducing unreasonable operational complexity**.

### The hierarchy to remember

```text
                    Linux Kernel
                         │
              ┌──────────┴──────────┐
              │                     │
          Pure Go                CGO / C
              │                     │
          Go runtime             libc
              │              ┌──────┴──────┐
              │            glibc          musl
              │              │              │
              └──────┬───────┴──────┬───────┘
                     │              │
                 static          dynamic
                     │              │
                     ▼              ▼
                  scratch      libc runtime
```

**The core rule:** `scratch` is not "a tiny Linux." It is **no userspace at all**. If your executable needs userspace libraries, you must either **statically include them** or **provide them in the image**.

---

## 🔗 References
- ⬆️ Parent: [[Target OS & Architecture]]
- 📚 Module: `Go Environment & Commands`
