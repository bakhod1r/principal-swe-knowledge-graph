---
title: "GOOS & GOARCH Matrix"
tags:
  - review
  - golang
  - environment
  - principal-swe
parent: "[[Target OS & Architecture]]"
---
# GOOS & GOARCH Matrix in Go

`GOOS` and `GOARCH` are the two core environment variables that tell the Go toolchain **what operating system** and **what CPU architecture** you are compiling for.

The mental model is:

```text
GOOS   = Target Operating System
GOARCH = Target CPU Architecture

                ┌──────────────┐
                │   Go Source  │
                └──────┬───────┘
                       │
                GOOS + GOARCH
                       │
                       ▼
              ┌────────────────┐
              │   Go Compiler   │
              └───────┬────────┘
                      │
                      ▼
             Target Executable
```

For example:

```bash
GOOS=linux GOARCH=amd64 go build
```

means:

> Build a Linux executable for 64-bit x86 CPUs.

---

## 1. What problem does the matrix solve?

Without cross-compilation, you would generally build software on the same platform where it will run.

For example:

```text
MacBook ARM64
      │
      ▼
    build
      │
      ▼
macOS ARM64 binary
```

But production environments are often different:

```text
Developer laptop
    macOS / arm64
          │
          │ cross compile
          ▼
Production
    Linux / amd64
```

Go makes this particularly straightforward because the standard compiler can target many OS/architecture combinations.

```bash
GOOS=linux GOARCH=amd64 go build
```

No VM is required.

No Linux machine is required.

For pure-Go programs, you usually do not even need a Linux environment to produce the Linux binary.

---

# 2. GOOS

`GOOS` identifies the target operating system.

Common values include:

|GOOS|Operating System|
|---|---|
|`linux`|Linux|
|`darwin`|macOS|
|`windows`|Windows|
|`freebsd`|FreeBSD|
|`openbsd`|OpenBSD|
|`netbsd`|NetBSD|
|`dragonfly`|DragonFly BSD|
|`android`|Android|
|`ios`|iOS|
|`js`|JavaScript/WebAssembly|
|`wasip1`|WASI Preview 1|

You can inspect supported targets with:

```bash
go tool dist list
```

Typical output looks like:

```text
aix/ppc64
android/386
android/amd64
android/arm
android/arm64
darwin/amd64
darwin/arm64
freebsd/amd64
linux/amd64
linux/arm
linux/arm64
windows/amd64
windows/arm64
...
```

The important point:

> `GOOS` does not describe your current OS. It describes the **target** OS.

---

# 3. GOARCH

`GOARCH` identifies the target CPU architecture.

Common architectures:

|GOARCH|CPU family|
|---|---|
|`amd64`|x86-64|
|`386`|x86 32-bit|
|`arm64`|ARM 64-bit|
|`arm`|ARM 32-bit|
|`ppc64le`|PowerPC 64-bit little-endian|
|`s390x`|IBM Z|
|`riscv64`|RISC-V 64-bit|
|`wasm`|WebAssembly|

For modern backend/DevOps work, the most important combinations are usually:

```text
linux/amd64
linux/arm64
darwin/amd64
darwin/arm64
windows/amd64
windows/arm64
```

---

# 4. The GOOS/GOARCH Matrix

Think of the target as a Cartesian product:

```text
             GOARCH
             │
             ├── amd64
             ├── arm64
             ├── arm
             ├── 386
             ├── riscv64
             └── ...
             │
GOOS ─────────┼────────────────────
             │
linux        ├── linux/amd64
             ├── linux/arm64
             ├── linux/arm
             └── ...
             │
darwin       ├── darwin/amd64
             └── darwin/arm64
             │
windows      ├── windows/amd64
             └── windows/arm64
             │
freebsd      └── freebsd/amd64
```

Not every combination exists.

For example, you cannot assume:

```text
darwin/riscv64
windows/s390x
ios/amd64
```

are valid targets.

That's why:

```bash
go tool dist list
```

is the authoritative practical way to inspect the matrix supported by your installed Go release.

---

# 5. Your host vs your target

This distinction is fundamental.

Suppose you have:

```text
Host:
macOS arm64
```

Check it:

```bash
go env GOOS GOARCH
```

You might get:

```text
darwin
arm64
```

Now build for Linux:

```bash
GOOS=linux GOARCH=amd64 go build
```

The resulting binary is:

```text
linux/amd64
```

Your host remains:

```text
darwin/arm64
```

You have simply instructed Go to produce a binary for another platform.

Mental model:

```text
HOST
macOS / arm64
     │
     │ Go compiler
     │
     ├───────────────► linux / amd64
     │
     ├───────────────► linux / arm64
     │
     └───────────────► windows / amd64
```

---

# 6. Practical examples

## Linux AMD64

```bash
GOOS=linux GOARCH=amd64 go build -o app
```

Produces:

```text
app
  └── Linux x86-64 executable
```

Useful for conventional x86-64 servers.

---

## Linux ARM64

```bash
GOOS=linux GOARCH=arm64 go build -o app
```

Common for:

- AWS Graviton
    
- ARM Kubernetes nodes
    
- Raspberry Pi 64-bit
    
- ARM cloud infrastructure
    

---

## macOS Apple Silicon

```bash
GOOS=darwin GOARCH=arm64 go build
```

For modern Apple Silicon Macs:

```text
M1
M2
M3
M4
...
```

---

## macOS Intel

```bash
GOOS=darwin GOARCH=amd64 go build
```

---

## Windows

```bash
GOOS=windows GOARCH=amd64 go build -o app.exe
```

---

# 7. `GOOS`/`GOARCH` affects source code

This is one of the most important parts.

Go supports platform-specific files.

Suppose:

```text
network_linux.go
network_windows.go
network_darwin.go
```

The Go toolchain automatically selects the appropriate implementation.

For Linux:

```text
network_linux.go
       ▲
       │
    GOOS=linux
```

For Windows:

```text
network_windows.go
       ▲
       │
   GOOS=windows
```

You don't manually import one of them.

Go's build system handles the selection.

---

# 8. Build constraints

You can also explicitly specify platform constraints.

For example:

```go
//go:build linux
```

This tells Go:

> Include this file only when building for Linux.

Architecture constraints are also possible:

```go
//go:build linux && amd64
```

Now the file is selected only for:

```text
linux/amd64
```

You can combine constraints:

```go
//go:build linux && (amd64 || arm64)
```

Meaning:

```text
linux/amd64
       OR
linux/arm64
```

---

# 9. Why platform-specific code exists

Operating systems expose different primitives.

For example:

```text
Linux
 ├── epoll
 ├── /proc
 ├── namespaces
 └── cgroups

Windows
 ├── Windows APIs
 ├── IOCP
 └── Windows services

macOS
 ├── kqueue
 └── Mach APIs
```

A single implementation cannot always use the same OS primitives.

Go therefore allows:

```text
common abstraction
       │
       ├── Linux implementation
       ├── Windows implementation
       └── Darwin implementation
```

This is much cleaner than:

```go
if runtime.GOOS == "linux" {
    ...
} else if runtime.GOOS == "windows" {
    ...
}
```

everywhere in your application.

---

# 10. `runtime.GOOS` and `runtime.GOARCH`

Go also exposes the target platform at runtime:

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.GOOS)
    fmt.Println(runtime.GOARCH)
}
```

On Linux ARM64:

```text
linux
arm64
```

But there is an important architectural distinction.

### Compile-time selection

Use build constraints when the **implementation itself differs**.

```text
foo_linux.go
foo_windows.go
```

### Runtime selection

Use `runtime.GOOS` when behavior can genuinely be selected dynamically.

```go
if runtime.GOOS == "windows" {
    ...
}
```

Do not use runtime branching merely because build constraints exist.

---

# 11. `GOARM`, `GOAMD64`, and friends

`GOOS` and `GOARCH` are not always sufficient.

Some architectures have additional tuning levels.

For example:

```bash
GOARCH=amd64
```

can be combined with:

```bash
GOAMD64=v1
GOAMD64=v2
GOAMD64=v3
GOAMD64=v4
```

These represent progressively stronger AMD64 CPU feature requirements.

Conceptually:

```text
GOARCH=amd64
       │
       ├── GOAMD64=v1
       ├── GOAMD64=v2
       ├── GOAMD64=v3
       └── GOAMD64=v4
```

Higher levels can enable newer CPU instructions, but they reduce compatibility with older CPUs.

This is a classic engineering trade-off:

```text
More CPU features
       ↓
Potentially better performance
       ↓
Smaller compatible CPU population
```

Do not blindly use `v4` in a generic production binary.

---

# 12. ARM has additional considerations

For 32-bit ARM:

```bash
GOARCH=arm
```

there is:

```bash
GOARM
```

For example:

```bash
GOOS=linux GOARCH=arm GOARM=7 go build
```

This targets ARMv7-class systems.

Again:

```text
GOARCH
  │
  └── CPU architecture family

GOARM
  │
  └── ARM-specific target level
```

---

# 13. GOOS/GOARCH is not the whole story

This is where engineers often develop the wrong mental model.

This:

```bash
GOOS=linux GOARCH=amd64 go build
```

does **not** guarantee:

> "I can run this binary on every Linux AMD64 machine."

Other factors matter.

For example:

```text
GOOS/GOARCH
     │
     ├── libc requirements
     ├── CGO
     ├── kernel requirements
     ├── CPU feature level
     ├── dynamic libraries
     ├── filesystem assumptions
     └── runtime behavior
```

A binary may technically be:

```text
linux/amd64
```

yet depend on:

```text
glibc
```

or external shared libraries.

This becomes especially important with:

```text
CGO_ENABLED=1
```

versus:

```text
CGO_ENABLED=0
```

---

# 14. Pure Go vs CGO

This is one of the most important cross-compilation boundaries.

### Pure Go

```bash
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64 \
go build
```

Cross-compilation is usually straightforward.

### CGO

```bash
CGO_ENABLED=1 \
GOOS=linux \
GOARCH=amd64 \
go build
```

Now you need a C compiler/toolchain appropriate for the target.

Conceptually:

```text
Pure Go

macOS/arm64
     │
     ▼
Go compiler
     │
     ▼
linux/amd64
```

With CGO:

```text
macOS/arm64
     │
     ▼
Go compiler
     +
target C compiler
     +
target libc
     │
     ▼
linux/amd64
```

This is why CGO cross-compilation becomes substantially more complicated.

---

# 15. Docker multi-platform builds

This matrix becomes especially important in modern container infrastructure.

A common image might support:

```text
linux/amd64
linux/arm64
```

You can build both:

```bash
GOOS=linux GOARCH=amd64 go build
GOOS=linux GOARCH=arm64 go build
```

Then publish a multi-platform container image.

The registry can expose one logical image:

```text
myapp:1.0
```

with manifests for:

```text
linux/amd64
linux/arm64
```

The container runtime selects the appropriate image.

Mental model:

```text
                 myapp:1.0
                     │
              Multi-platform
                 manifest
               /          \
              /            \
     linux/amd64        linux/arm64
```

This is a foundational pattern for modern Kubernetes/cloud deployments.

---

# 16. A production build matrix

For a backend service, you might define:

```text
linux/amd64
linux/arm64
```

Then CI builds both.

For example:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build
```

This gives you:

```text
app-linux-amd64
app-linux-arm64
```

You can verify the resulting binaries with tools such as:

```bash
file app-linux-amd64
file app-linux-arm64
```

---

# 17. Testing the matrix

Building successfully is not equivalent to testing successfully.

Suppose:

```text
Build:
  linux/amd64     ✓
  linux/arm64     ✓
```

You still need confidence that:

```text
Runtime:
  linux/amd64     ✓
  linux/arm64     ✓
```

For platform-sensitive software, the stronger pipeline is:

```text
Source
  │
  ▼
Build matrix
  │
  ├── linux/amd64 ──► test
  │
  └── linux/arm64 ──► test
                     │
                     ▼
                  package
                     │
                     ▼
                  deploy
```

This matters particularly when dealing with:

- CGO
    
- architecture-specific assembly
    
- CPU instructions
    
- syscall behavior
    
- filesystem semantics
    
- endian assumptions
    
- unsafe code
    
- platform-specific dependencies
    

---

# 18. Common mistakes

### Mistake 1 — Confusing host and target

```bash
go env GOOS GOARCH
```

tells you the defaults for the environment/toolchain.

It does not mean you are restricted to those values.

You can override them:

```bash
GOOS=linux GOARCH=arm64 go build
```

---

### Mistake 2 — Assuming every combination works

Don't assume:

```text
GOOS × GOARCH = everything
```

Use:

```bash
go tool dist list
```

---

### Mistake 3 — Ignoring CGO

This command:

```bash
GOOS=linux GOARCH=amd64 go build
```

can become much more complicated when:

```text
CGO_ENABLED=1
```

is involved.

---

### Mistake 4 — Testing only amd64

If production supports:

```text
amd64
arm64
```

but CI tests only:

```text
amd64
```

you have an untested production platform.

---

### Mistake 5 — Using runtime branching for platform implementations

Avoid spreading:

```go
if runtime.GOOS == ...
```

throughout the codebase.

Prefer:

```text
foo_linux.go
foo_windows.go
foo_darwin.go
```

when implementations genuinely differ.

---

# 19. Useful commands

Show current target:

```bash
go env GOOS GOARCH
```

Show all supported targets:

```bash
go tool dist list
```

Build Linux AMD64:

```bash
GOOS=linux GOARCH=amd64 go build
```

Build Linux ARM64:

```bash
GOOS=linux GOARCH=arm64 go build
```

Build Windows:

```bash
GOOS=windows GOARCH=amd64 go build
```

Build without CGO:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

Inspect Go environment:

```bash
go env
```

---

# 20. The Principal Engineer mental model

Don't memorize:

```text
GOOS=linux
GOARCH=amd64
```

Instead think in layers:

```text
                    Target
                      │
             ┌────────┴────────┐
             │                 │
            GOOS             GOARCH
             │                 │
          linux              amd64
             │                 │
             └────────┬────────┘
                      │
               CPU feature level
                      │
                  GOAMD64
                      │
                      ▼
                 Go compiler
                      │
              ┌───────┴────────┐
              │                │
          Pure Go             CGO
              │                │
              │          C compiler + libc
              │                │
              └───────┬────────┘
                      ▼
                Executable
                      │
                      ▼
              Runtime environment
                      │
          ┌───────────┼───────────┐
          ▼           ▼           ▼
        Kernel      libc      CPU features
```

So **GOOS/GOARCH identifies the basic target platform, but it does not fully describe runtime compatibility**.

That distinction is critical in production.

---

## Key takeaways

1. **`GOOS` = target operating system.**
    
2. **`GOARCH` = target CPU architecture.**
    
3. Go supports extensive cross-compilation.
    
4. `go tool dist list` shows the supported target matrix.
    
5. Build constraints allow platform-specific implementations.
    
6. `runtime.GOOS`/`runtime.GOARCH` are runtime information; build tags are compile-time selection.
    
7. `CGO_ENABLED=1` significantly changes cross-compilation complexity.
    
8. `GOAMD64`, `GOARM`, etc. provide architecture-specific tuning.
    
9. Container platforms commonly target `linux/amd64` and `linux/arm64`.
    
10. **A successful cross-build does not automatically prove runtime compatibility.**
    

The deeper engineering lesson is:

> **Platform portability is not just "can the compiler produce a binary?" It is "can that binary execute correctly, efficiently, and observably in the target environment?"**

That is the level at which `GOOS/GOARCH` becomes a system-design concern rather than merely a Go build flag.

---

## 🔗 References
- ⬆️ Parent: [[Target OS & Architecture]]
- 📚 Module: `Go Environment & Commands`
