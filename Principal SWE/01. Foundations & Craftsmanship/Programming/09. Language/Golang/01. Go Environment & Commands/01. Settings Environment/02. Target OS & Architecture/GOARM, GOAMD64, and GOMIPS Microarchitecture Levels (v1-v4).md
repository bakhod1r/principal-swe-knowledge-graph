---
title: "GOARM, GOAMD64, and GOMIPS Microarchitecture Levels (v1-v4)"
tags:
  - golang
  - architecture
  - microarchitecture
  - goamd64
  - goarm
  - principal-swe
parent: "[[Target OS & Architecture]]"
---
# GOARM, GOAMD64, and GOMIPS Microarchitecture Levels (v1–v4)

These variables control **which CPU instruction set level the Go compiler is allowed to target** for certain architectures.

The key mental model is:

> `GOARCH` answers **“Which CPU family?”**  
> `GOARM / GOAMD64 / GOMIPS` answer **“Which capabilities within that CPU family?”**

This matters because choosing a higher level can produce binaries that use newer CPU instructions, but those binaries may **fail to run on older CPUs**.

---

## 1. Why do these variables exist?

Consider `GOARCH=amd64`.

Two CPUs can both be x86-64:

```text
Old x86-64 CPU
    ↓
supports baseline AMD64 instructions

Modern x86-64 CPU
    ↓
supports AVX2, BMI2, etc.
```

If Go always targeted the newest instructions, the binary could be faster—but it would no longer run on older x86-64 machines.

So Go provides architecture-specific levels:

```text
GOARCH=amd64
        │
        ├── GOAMD64=v1
        ├── GOAMD64=v2
        ├── GOAMD64=v3
        └── GOAMD64=v4
```

The same idea exists for ARM and MIPS.

---

# 2. GOARM

`GOARM` controls the ARM architecture level for **32-bit ARM** (`GOARCH=arm`).

Typical values:

```bash
GOOS=linux
GOARCH=arm
GOARM=5
```

or:

```bash
GOARM=6
GOARM=7
```

### Mental model

```text
GOARCH=arm
    │
    ├── GOARM=5
    ├── GOARM=6
    └── GOARM=7
```

Higher values generally mean a newer ARM ISA / CPU capability level.

### Important distinction

Do not confuse:

```text
GOARCH=arm
```

with:

```text
GOARCH=arm64
```

They are different architectures.

```text
arm       → 32-bit ARM
arm64     → 64-bit ARM / AArch64
```

`GOARM` applies to:

```text
GOARCH=arm
```

not `arm64`.

---

## GOARM levels

### GOARM=5

Targets older ARMv5-class processors.

Useful when supporting very old ARM hardware.

Trade-off:

```text
maximum compatibility
        ↓
less modern CPU capability
```

---

### GOARM=6

Targets ARMv6-class CPUs.

Typical historical examples include older Raspberry Pi generations.

---

### GOARM=7

Targets ARMv7-class processors.

This is common for older 32-bit ARM systems.

Example:

```bash
GOOS=linux GOARCH=arm GOARM=7 go build
```

---

# 3. GOAMD64

`GOAMD64` is particularly important for modern backend/DevOps work.

It controls the minimum x86-64 CPU feature level.

```text
GOARCH=amd64
    │
    ├── v1
    ├── v2
    ├── v3
    └── v4
```

Example:

```bash
GOOS=linux
GOARCH=amd64
GOAMD64=v3
go build
```

Now the resulting binary is allowed to use instructions available at the v3 level.

---

# 4. GOAMD64=v1

`v1` is the **baseline x86-64 instruction set**.

Think:

```text
v1
│
└── broadest amd64 compatibility
```

This is generally the safest choice when your deployment environment contains unknown or old x86-64 CPUs.

For example:

```bash
GOAMD64=v1 go build
```

Mental model:

```text
Compatibility ████████████████████
CPU features  █████
Optimization  █████
```

---

# 5. GOAMD64=v2

`v2` requires a newer x86-64 feature set.

Conceptually:

```text
v1
 ↓
v2
```

It enables additional CPU instructions such as features from later x86-64 generations.

This can be useful when your fleet is known to support the required instruction set.

---

# 6. GOAMD64=v3

`v3` is where this becomes particularly interesting for performance-sensitive workloads.

It targets CPUs with substantially newer SIMD and bit-manipulation capabilities.

For example:

```bash
GOAMD64=v3 go build
```

This can allow the compiler/runtime to make use of instructions unavailable on baseline x86-64.

But there is a production constraint:

> Every machine that executes the binary must support the selected level.

For example:

```text
Build:
GOAMD64=v3

        ↓

Docker image
        ↓

Kubernetes cluster
        ↓
┌───────────────┐
│ Node A: v3 ✓  │
│ Node B: v3 ✓  │
│ Node C: v1 ✗  │
└───────────────┘
```

The deployment may fail only when the pod lands on Node C.

This is why CPU architecture compatibility is a **deployment concern**, not merely a compiler concern.

---

# 7. GOAMD64=v4

`v4` targets an even newer x86-64 feature level.

Conceptually:

```text
v1
 │
 v2
 │
 v3
 │
 v4
```

The compatibility set becomes progressively smaller.

Therefore:

```text
higher GOAMD64
        ↓
newer CPU requirement
        ↓
potentially better instruction-level optimization
        ↓
smaller compatible hardware fleet
```

Do **not** automatically choose `v4` because it is "faster."

The correct question is:

> "Does my production CPU fleet guarantee this instruction set, and does benchmarking demonstrate a meaningful benefit?"

---

# 8. GOMIPS

`GOMIPS` controls MIPS architecture floating-point behavior.

It is used with MIPS targets such as:

```text
GOARCH=mips
GOARCH=mipsle
GOARCH=mips64
GOARCH=mips64le
```

Typical values include:

```text
GOMIPS=hardfloat
GOMIPS=softfloat
```

Mental model:

```text
MIPS
 │
 ├── hardfloat
 │
 └── softfloat
```

### hardfloat

Use hardware floating-point instructions.

```bash
GOARCH=mips GOMIPS=hardfloat go build
```

Requires appropriate hardware support.

### softfloat

Floating-point operations are implemented without relying on hardware floating-point instructions.

```bash
GOARCH=mips GOMIPS=softfloat go build
```

This improves compatibility with CPUs lacking the required floating-point hardware, potentially at a performance cost.

---

# 9. The important difference: GOARM vs GOAMD64 vs GOMIPS

These variables look similar but solve slightly different problems.

|Variable|Architecture|Purpose|
|---|---|---|
|`GOARM`|`arm`|ARM ISA/CPU level|
|`GOAMD64`|`amd64`|x86-64 ISA feature level|
|`GOMIPS`|MIPS variants|Floating-point mode|

So:

```text
GOARCH
  │
  ├── arm
  │     └── GOARM
  │
  ├── amd64
  │     └── GOAMD64
  │
  └── mips*
        └── GOMIPS
```

---

# 10. How to inspect your current configuration

Run:

```bash
go env GOARCH GOOS GOARM GOAMD64 GOMIPS
```

For example:

```text
amd64
linux
7
v3
hardfloat
```

Not every variable is meaningful for every architecture.

For example:

```bash
GOARCH=arm64
```

does not mean `GOARM` controls the ARM64 target.

---

# 11. Cross-compilation example

Suppose your development machine is:

```text
Apple Silicon
arm64
```

but your production servers are:

```text
Linux
amd64
```

You can build:

```bash
GOOS=linux GOARCH=amd64 GOAMD64=v3 go build
```

Now:

```text
Developer
   │
   │ arm64
   ▼
Go compiler
   │
   │ cross compile
   ▼
linux/amd64
   │
   │ GOAMD64=v3
   ▼
Production binary
```

The build machine's CPU does **not** determine the target CPU level.

The target is explicitly specified by the environment.

---

# 12. Docker/Kubernetes trap

This is a very practical production issue.

Imagine:

```dockerfile
FROM golang:1.XX AS builder

ENV GOOS=linux
ENV GOARCH=amd64
ENV GOAMD64=v3

RUN go build -o /app
```

The resulting binary requires the v3 CPU feature level.

Now your Kubernetes cluster contains heterogeneous nodes:

```text
Node 1 → AMD EPYC        → v3 ✓
Node 2 → modern Intel    → v3 ✓
Node 3 → old Xeon        → v1 ✗
```

Kubernetes generally schedules based on resources/labels, not automatically on arbitrary CPU instruction-level requirements.

So you can get:

```text
Deployment
   ↓
Pod scheduled
   ↓
Binary starts
   ↓
Illegal instruction
   ↓
CrashLoopBackOff
```

This is one of the most important operational implications of microarchitecture-specific builds.

---

# 13. `illegal instruction` is a key failure mode

A classic symptom is:

```text
SIGILL
Illegal instruction
```

The mental model is:

```text
Binary
  │
  ├── contains instruction requiring CPU feature X
  │
  ▼
CPU
  │
  └── does not support X
          ↓
       SIGILL
```

This is different from a normal Go runtime error.

The process can fail before your application gets a chance to handle anything.

---

# 14. Production strategy

For most backend services, I would start with:

```text
GOAMD64=v1
```

unless you have a concrete reason to raise it.

Then establish the fleet constraint:

```text
Production CPU fleet
        ↓
Determine minimum CPU feature level
        ↓
Benchmark
        ↓
Choose GOAMD64
        ↓
Enforce scheduling/deployment constraints
```

For a controlled infrastructure fleet:

```text
GOAMD64=v3
```

can be reasonable if:

1. all production nodes support v3,
    
2. CI/CD builds explicitly target v3,
    
3. node replacement policies preserve that guarantee,
    
4. benchmarks show meaningful benefit,
    
5. deployment constraints prevent incompatible nodes.
    

---

# 15. Multi-binary strategy

Sometimes you don't want one binary for everyone.

You can build:

```text
service-amd64-v1
service-amd64-v3
service-arm64
```

Then deployment infrastructure selects the appropriate artifact.

Conceptually:

```text
                ┌── amd64-v1
Application ────┼── amd64-v3
                └── arm64
```

This can be useful for large fleets, but introduces:

```text
more artifacts
+ more CI complexity
+ more testing
+ more release management
```

So don't do it unless the performance/compatibility economics justify it.

---

# 16. Principal Engineer mental model

Don't think:

> "`GOAMD64=v4` is faster, therefore use v4."

Think:

```text
                 CPU fleet
                    │
                    ▼
          Minimum supported ISA
                    │
                    ▼
             GOAMD64 level
                    │
          ┌─────────┴─────────┐
          ▼                   ▼
   Compatibility          Performance
          │                   │
          └─────────┬─────────┘
                    ▼
               Benchmark
                    │
                    ▼
             Deployment policy
```

The optimization is only valid if the **whole system** supports the assumption.

### The deeper lesson

`GOARM`, `GOAMD64`, and `GOMIPS` are not merely compiler flags.

They create a **compatibility contract between your binary and your infrastructure**.

That contract should be reflected in:

- CI/CD
    
- container images
    
- Kubernetes node pools
    
- VM instance types
    
- autoscaling
    
- disaster recovery environments
    
- developer/testing environments
    
- benchmarking
    
- deployment documentation
    

If you cannot guarantee the CPU capability, use the more compatible target.

**Rule of thumb:**

> **Choose the lowest microarchitecture level that satisfies your performance requirements and is guaranteed across the entire execution fleet.**

## 🔗 References
- ⬆️ Parent: [[Target OS & Architecture]]
- 📚 Module: `Go Environment & Commands`
