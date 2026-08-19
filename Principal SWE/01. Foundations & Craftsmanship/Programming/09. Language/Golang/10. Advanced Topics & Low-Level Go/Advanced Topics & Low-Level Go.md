---
title: Advanced Topics & Low-Level Go
tags:
  - golang
  - advanced
  - principal-swe
parent: "[[Golang]]"
---

# 🔬 Advanced Topics & Low-Level Go

Low-level systems programming in Go: unsafe pointer arithmetic, runtime reflection, Cgo FFI, Plan 9 assembly, compiler SSA pipeline, custom analyzers, and linker internals.

```text
Advanced Topics & Low-Level Go
│
├── [[Unsafe Pointers & Memory Layout Manipulation|01. Unsafe Pointers & Memory Layout Manipulation]]
│   ├── `unsafe.Pointer Mechanics & Arbitrary Pointer Casting`
│   ├── `uintptr Pointer Arithmetic Rules & GC Safety Constraints`
│   ├── `unsafe.Slice and unsafe.String Zero-Copy Constructors`
│   ├── `Struct Field Offsets (unsafe.Offsetof & unsafe.Alignof)`
│   ├── `Memory Alignment & Word Boundaries (8-byte Alignment)`
│   └── `Pointer Pinning (runtime.Pinner) for Cgo Integration`
├── [[Metaprogramming & Reflection (reflect)|02. Metaprogramming & Reflection (reflect)]]
│   ├── `Laws of Reflection (Interface to reflect.Value, reflect.Type)`
│   ├── `Deep Struct Field Inspection & Dynamic Method Invocation`
│   ├── `Settability & Addressability in Reflection`
│   ├── `Reflect-Based Code Generation vs Static Generation`
│   ├── `Type Creation at Runtime (reflect.StructOf & reflect.ArrayOf)`
│   └── `Go Plugins (.so) Dynamic Loading Architecture`
├── [[FFI & Cgo Architecture|03. FFI & Cgo Architecture]]
│   ├── `Cgo Architecture & Cross-Language Stack Switching`
│   ├── `Cgo Performance Overhead & Call Overhead Benchmarks`
│   ├── `Passing Pointers between Go and C (Pointer Passing Rules)`
│   ├── `Pure Go vs Cgo Compilation (CGO_ENABLED=0 vs 1)`
│   └── `Calling Go Functions from C (export Directives)`
├── [[Plan 9 Go Assembly & Hardware Architecture|04. Plan 9 Go Assembly & Hardware Architecture]]
│   ├── `Plan 9 Assembly Syntax & Architecture Differences`
│   ├── `Stack Frame Layout in Go Assembly Functions`
│   ├── `Go ABIInternal Register Calling Convention Mechanics`
│   ├── `Writing SIMD Vector Instructions in Go Assembly`
│   └── `CPU Feature Detection (internal-cpu & CPUID)`
├── `05. Compiler Pipeline, AST & Static Analysis`
│   ├── `Compiler Pipeline Architecture (Scanning to Code Generation)`
│   ├── `AST Programmatic Manipulation (go-ast & go-parser)`
│   ├── `Building Custom Analyzers with go-analysis Framework`
│   ├── `Compiler SSA (Static Single Assignment) Representation Passes`
│   ├── `Bounds Check Elimination (BCE) Compiler Pass`
│   ├── `Function Inlining Heuristics & Budget Calculation`
│   └── `Compiler Directives & Pragmas (go:noinline, go:nosplit, go:linkname, go:notinheap)`
├── [[Linker Internals & Binary Generation|06. Linker Internals & Binary Generation]]
└── `07. WebAssembly, WASI & Alternative Targets`
│   ├── `Linker Architecture & Symbol Resolution (cmd-link)`
│   ├── `Global Dead Code Elimination in Linker`
│   ├── `Binary Size Optimization Matrix & DWARF Stripping`
│   ├── `Embedding VCS Metadata & Build Info (debug.ReadBuildInfo)`
│   └── `Reproducible Builds & Path Stripping (-trimpath)`
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Unsafe Pointers & Memory Layout Manipulation|01. Unsafe Pointers & Memory Layout Manipulation]]
- `unsafe.Pointer Mechanics & Arbitrary Pointer Casting` — Bypassing Go static type system, casting between arbitrary pointer types, and pointer safety rules.
- `uintptr Pointer Arithmetic Rules & GC Safety Constraints` — Safe conversion patterns: unsafe.Pointer to uintptr for offset arithmetic without GC pointer relocation traps.
- `unsafe.Slice and unsafe.String Zero-Copy Constructors` — Go 1.20+ safe zero-copy slice and string constructors replacing legacy reflect headers.
- `Struct Field Offsets (unsafe.Offsetof & unsafe.Alignof)` — Computing struct memory field byte offsets dynamically for high-speed binary serialization.
- `Memory Alignment & Word Boundaries (8-byte Alignment)` — Natural word boundaries, CPU memory alignment rules, struct padding, and unaligned panic prevention.
- `Pointer Pinning (runtime.Pinner) for Cgo Integration` — Pinning Go heap memory addresses to prevent GC movement during concurrent foreign C function calls.
### 2. 📂 [[Metaprogramming & Reflection (reflect)|02. Metaprogramming & Reflection (reflect)]]
- `Laws of Reflection (Interface to reflect.Value, reflect.Type)` — Dissecting the 3 fundamental Laws of Reflection in Go and type representation.
- `Deep Struct Field Inspection & Dynamic Method Invocation` — Extracting struct tags, reading private fields via unsafe, and invoking methods dynamically.
- `Settability & Addressability in Reflection` — v.CanSet(), v.Elem(), and taking pointer addresses to mutate structs through reflection.
- `Reflect-Based Code Generation vs Static Generation` — Benchmarking runtime reflection performance penalties against build-time code generation.
- `Type Creation at Runtime (reflect.StructOf & reflect.ArrayOf)` — Programmatically constructing new dynamic struct and array types at runtime.
- `Go Plugins (.so) Dynamic Loading Architecture` — Building, compiling, and loading dynamic shared libraries (plugin.Open, plugin.Lookup) at runtime.
### 3. 📂 [[FFI & Cgo Architecture|03. FFI & Cgo Architecture]]
- `Cgo Architecture & Cross-Language Stack Switching` — How Go executes C code: switching from 2KB Go stack to OS C thread stack and back.
- `Cgo Performance Overhead & Call Overhead Benchmarks` — Measuring the 100ns+ CPU overhead per Cgo call and optimizing via batching.
- `Passing Pointers between Go and C (Pointer Passing Rules)` — Rules governing passing Go pointers to C, C.CString, C.free, and memory leak prevention.
- `Pure Go vs Cgo Compilation (CGO_ENABLED=0 vs 1)` — Building pure statically linked binaries, musl vs glibc compatibility, and Docker scratch deployment.
- `Calling Go Functions from C (export Directives)` — Exporting Go functions with //export and building .so/.dylib C-shared dynamic libraries.
### 4. 📂 [[Plan 9 Go Assembly & Hardware Architecture|04. Plan 9 Go Assembly & Hardware Architecture]]
- `Plan 9 Assembly Syntax & Architecture Differences` — Pseudo-registers (FP, SP, SB, PC), instruction operands, and Plan 9 assembly conventions.
- `Stack Frame Layout in Go Assembly Functions` — Writing assembly functions: stack size, argument size declarations (TEXT ·Add(SB), $0-24), and returns.
- `Go ABIInternal Register Calling Convention Mechanics` — Passing integer arguments in RAX, RBX, RCX and floating-point in X0-X14 registers.
- `Writing SIMD Vector Instructions in Go Assembly` — Utilizing AVX2, AVX-512, and ARM NEON vectorized instructions for high-throughput computing.
- `CPU Feature Detection (internal-cpu & CPUID)` — Detecting runtime CPU capabilities (cpu.X86.HasAVX2, cpu.ARM64.HasAES) to dispatch optimized paths.
### 5. 📂 `05. Compiler Pipeline, AST & Static Analysis`
- `Compiler Pipeline Architecture (Scanning to Code Generation)` — Lexer (scanner), Parser (go/ast), Type Checker (go/types), SSA middle-end, Assembler, and Linker.
- `AST Programmatic Manipulation (go-ast & go-parser)` — Parsing, traversing (ast.Walk), inspecting, and rewriting Go Abstract Syntax Trees.
- `Building Custom Analyzers with go-analysis Framework` — Writing custom enterprise static analysis linters checking architectural rules and conventions.
- `Compiler SSA (Static Single Assignment) Representation Passes` — Lowering AST into SSA basic blocks, optimization passes, and Plan 9 code emission.
- `Bounds Check Elimination (BCE) Compiler Pass` — Proving slice bounds to the compiler (-gcflags="-d=ssa/check_bce") to remove runtime bounds checks.
- `Function Inlining Heuristics & Budget Calculation` — Inlining cost scoring algorithm (80 AST node budget), mid-stack inlining, and inlining barriers.
- `Compiler Directives & Pragmas (go:noinline, go:nosplit, go:linkname, go:notinheap)` — Complete deep catalog of compiler pragmas controlling runtime code generation.
### 6. 📂 [[Linker Internals & Binary Generation|06. Linker Internals & Binary Generation]]
- `Linker Architecture & Symbol Resolution (cmd-link)` — Object file parsing, global symbol table resolution, relocations, and internal vs external linking.
- `Global Dead Code Elimination in Linker` — Reachability graph analysis stripping unreachable functions and packages from the final binary.
- `Binary Size Optimization Matrix & DWARF Stripping` — Stripping symbols (-ldflags="-s -w"), removing DWARF debug tables, and binary layout analysis.
- `Embedding VCS Metadata & Build Info (debug.ReadBuildInfo)` — Extracting embedded Git commit hashes, dirty status, and compiler flags from compiled binaries.
- `Reproducible Builds & Path Stripping (-trimpath)` — Generating byte-for-byte identical binary checksums across different developer machines and CI runners.

### 7. 📂 `07. WebAssembly, WASI & Alternative Targets`
- `WebAssembly Compilation Architecture (GOOS=js GOARCH=wasm)` — Compiling Go binaries to .wasm, wasm_exec.js runtime glue, and browser execution lifecycle.
- `syscall-js Bridge & DOM Interop Mechanics` — js.Value, js.Global(), registering DOM event callbacks, and JavaScript promise integration in Go.
- `Zero-Copy Memory Sharing (Go Wasm to JS TypedArrays)` — Direct memory view access via js.CopyBytesToGo and js.CopyBytesToJS without serialization.
- `Server-Side Wasm with WASI (GOOS=wasip1 GOARCH=wasm)` — Go 1.21+ WASI standard support, executing Go Wasm in Wasmtime, Wasmer, and Docker Wasm.
- `Wasm Plugin Architecture in Envoy & Edge Proxies` — Building low-latency proxy filter extensions (Proxy-Wasm) for Envoy, Istio, and API gateways.
- `TinyGo Compiler Architecture for Microcontrollers & Web` — LLVM-based Go compiler, sub-100KB Wasm binaries, and zero-overhead embedded garbage collection.
- `Embedded IoT Programming with TinyGo (ESP32, RP2040, STM32)` — Real-time GPIO, I2C, SPI bus communication, and interrupt handlers on bare-metal silicon.
- `Gomobile Native Bindings for iOS & Android` — Generating native Objective-C/Swift and Java/Kotlin bindings from Go libraries (gomobile bind).

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

---

## 🗂️ Topics

- [[Alternative Targets]]
- [[Compiler Pipeline]]
- [[FFI & Cgo Architecture]]
- [[Linker Internals & Binary Generation]]
- [[Metaprogramming & Reflection (reflect)]]
- [[Plan 9 Go Assembly & Hardware Architecture]]
- [[Static Analysis Tooling]]
- [[Unsafe Pointers & Memory Layout Manipulation]]
- [[WebAssembly and WASI]]
