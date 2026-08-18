---
title: Language Basics
tags:
  - golang
  - basics
  - principal-swe
parent: "[[Golang]]"
---

# 🔤 Language Basics

Core language syntax, environment setup, toolchain commands, modules, GOROOT layout, types, control flow, functions, and pointers.

```text
Language Basics
│
├── [[Settings Environment|01. Settings Environment]]
│   ├── [[PATH]]
│   ├── [[GOROOT]]
│   ├── [[GOPATH]]
│   ├── [[GOBIN]]
│   ├── [[GOENV]]
│   ├── [[GOOS and GOARCH]]
│   ├── [[GOCACHE and GOTMPDIR]]
│   └── [[Shell Startup]]
├── [[Go Commands|02. Go command]]
│   ├── [[go build]]
│   ├── [[go run]]
│   ├── [[go install]]
│   ├── [[go test]]
│   ├── [[go fmt]]
│   ├── [[go vet]]
│   ├── [[go generate]]
│   ├── [[go doc]]
│   ├── [[go clean]]
│   └── [[go work]]
├── [[Dependencies & Go Modules|03. Dependencies]]
│   ├── [[go.mod]]
│   ├── [[go.sum]]
│   ├── [[require Directive]]
│   ├── [[replace Directive]]
│   ├── [[retract Directive]]
│   ├── [[Semantic Versioning]]
│   ├── [[MVS (Minimal Version Selection)]]
│   ├── [[GOPROXY]]
│   ├── [[GOSUMDB]]
│   ├── [[GOPRIVATE]]
│   └── [[Vendoring]]
├── [[Go Source Code Structure|04. Go Source Code Structure]]
│   ├── [[src Directory]]
│   ├── [[src-runtime]]
│   ├── [[src-cmd]]
│   ├── [[src-internal]]
│   └── [[GOROOT bin and pkg]]
├── [[Variables & Constants|05. Variables & Constants]]
│   ├── [[var Declaration]]
│   ├── [[Short Variable Declaration (:=)]]
│   ├── [[Zero Values]]
│   ├── [[const Declaration]]
│   ├── [[iota Enumerator]]
│   ├── [[Variable Shadowing]]
│   └── [[Blank Identifier (_)]]
├── [[Data Types|06. Data Types]]
│   ├── [[Integers (Signed and Unsigned)]]
│   ├── [[Floating Points]]
│   ├── [[Complex Numbers]]
│   ├── [[Boolean]]
│   ├── [[Runes & UTF-8]]
│   ├── [[Strings]]
│   ├── [[Raw String Literals]]
│   ├── [[Interpreted String Literals]]
│   ├── [[String Internals (stringStruct)]]
│   ├── [[Numeric Conversions]]
│   └── [[Overflow and Precision]]
├── [[Composite Types|07. Composite Types]]
│   ├── [[Arrays]]
│   ├── [[Slices]]
│   ├── [[Slice Header Internals]]
│   ├── [[Slice Capacity and Growth]]
│   ├── [[make() for Slices and Maps]]
│   ├── [[Slice to Array Conversion]]
│   ├── [[Slice Tricks]]
│   ├── [[Maps]]
│   ├── [[Map Internals (hmap and bmap)]]
│   ├── [[Comma-Ok Idiom for Maps]]
│   ├── [[Structs]]
│   ├── [[Struct Memory Layout & Padding]]
│   ├── [[Struct Tags & JSON]]
│   ├── [[Embedding Structs]]
│   ├── [[Empty Struct (struct{})]]
│   └── [[Anonymous Structs]]
├── [[Conditionals|08. Conditionals]]
│   ├── [[if Statement]]
│   ├── [[Short Statement in If]]
│   ├── [[if-else Chains]]
│   ├── [[switch Statement]]
│   ├── [[Tagless Switch]]
│   └── [[Type Switch]]
├── [[Loops & Iteration|09. Loops]]
│   ├── [[for Loop]]
│   ├── [[for range]]
│   ├── [[Range over Integer (Go 1.22+)]]
│   ├── [[Range over Func (Go 1.23+)]]
│   ├── [[Iterating Maps]]
│   ├── [[Iterating Strings]]
│   ├── [[break Statement]]
│   ├── [[continue Statement]]
│   ├── [[goto Statement]]
│   └── [[Labeled break and continue]]
├── [[Functions & Closures|10. Functions]]
│   ├── [[Function Declarations]]
│   ├── [[Multiple Return Values]]
│   ├── [[Named Return Values]]
│   ├── [[Variadic Functions]]
│   ├── [[Anonymous Functions]]
│   ├── [[Closures]]
│   ├── [[Closure Internals (Heap Escape)]]
│   ├── [[Call by Value]]
│   ├── [[defer Statement]]
│   ├── [[defer Ordering and Evaluation]]
│   └── [[init() Function]]
└── [[Pointers & Memory|11. Pointers]]
│   ├── [[Pointers Basics]]
│   ├── [[Pointer vs Value Semantics]]
│   ├── [[Pointers with Structs]]
│   ├── [[Pointers with Slices & Maps]]
│   ├── [[unsafe.Pointer and uintptr]]
│   ├── [[nil Pointer Dereference]]
│   ├── [[Memory Management & Escape Analysis]]
│   └── [[Garbage Collection Overview]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Settings Environment|01. Settings Environment]]
- [[PATH]] — OS search path for Go binaries and installed tools.
- [[GOROOT]] — Installation directory of the Go SDK and standard library.
- [[GOPATH]] — User workspace, module cache, and binary download directory.
- [[GOBIN]] — Target directory where go install writes executable binaries.
- [[GOENV]] — Location of persistent Go environment variable settings file.
- [[GOOS and GOARCH]] — Target operating system and CPU architecture for compilation.
- [[GOCACHE and GOTMPDIR]] — Build artifact cache and temporary compilation directory.
- [[Shell Startup]] — Configuring environment persistence in .zshrc, .bashrc, or profile.
### 2. 📂 [[Go Commands|02. Go command]]
- [[go build]] — Compiling packages and dependencies into executables or packages.
- [[go run]] — Compiling and executing Go source files on the fly.
- [[go install]] — Compiling and installing executables into GOBIN.
- [[go test]] — Automated testing and benchmark execution.
- [[go fmt]] — Standard code formatting with gofmt.
- [[go vet]] — Compiler static analysis for suspicious constructs.
- [[go generate]] — Executing code generators via //go:generate directives.
- [[go doc]] — Displaying documentation for packages and symbols from the terminal.
- [[go clean]] — Removing object files and cached package artifacts.
- [[go work]] — Multi-module local development workflow commands.
### 3. 📂 [[Dependencies & Go Modules|03. Dependencies]]
- [[go.mod]] — Module declaration file, module path, and dependency requirements.
- [[go.sum]] — Cryptographic checksums of direct and indirect module dependencies.
- [[require Directive]] — Declaring minimum required versions of external modules.
- [[replace Directive]] — Substituting module dependencies with local forks or paths.
- [[retract Directive]] — Retracting broken or retracted module version releases.
- [[Semantic Versioning]] — SemVer rules and v2+ major version import path suffixes.
- [[MVS (Minimal Version Selection)]] — Deterministic dependency resolution algorithm in Go.
- [[GOPROXY]] — HTTP download mirror and proxy for public Go modules.
- [[GOSUMDB]] — Cryptographic notary database for module checksum verification.
- [[GOPRIVATE]] — Configuring private corporate repositories bypassing GOPROXY/GOSUMDB.
- [[Vendoring]] — Embedding dependencies in local vendor/ directory with go mod vendor.
### 4. 📂 [[Go Source Code Structure|04. Go Source Code Structure]]
- [[src Directory]] — Root source tree for all standard library packages and runtime.
- [[src-runtime]] — Core runtime engine (proc.go, mgc.go, malloc.go, chan.go).
- [[src-cmd]] — Toolchain source (cmd/go, cmd/compile, cmd/link, cmd/asm).
- [[src-internal]] — Private unimportable standard library helper packages.
- [[GOROOT bin and pkg]] — Executable toolchain binaries and precompiled metadata archives.
### 5. 📂 [[Variables & Constants|05. Variables & Constants]]
- [[var Declaration]] — Explicit variable declarations with type annotations.
- [[Short Variable Declaration (:=)]] — Type-inferred local variable declaration syntax.
- [[Zero Values]] — Default memory initialization for all Go types (0, false, "", nil).
- [[const Declaration]] — Compile-time immutable constants and untyped numeric constants.
- [[iota Enumerator]] — Sequential compile-time constant generator and bitmask idioms.
- [[Variable Shadowing]] — Inner block variable redeclaration masking outer scope variables.
- [[Blank Identifier (_)]] — Ignoring unused return values, imports, and interface checks.
### 6. 📂 [[Data Types|06. Data Types]]
- [[Integers (Signed and Unsigned)]] — int8, int16, int32, int64, uint8, uint16, uint32, uint64, int, uint, uintptr.
- [[Floating Points]] — IEEE-754 float32, float64 precision and NaN/Infinity behaviors.
- [[Complex Numbers]] — complex64, complex128 arithmetic and real/imag builtins.
- [[Boolean]] — bool type (true, false) and logical operators (&&, ||, !).
- [[Runes & UTF-8]] — rune (int32) representing Unicode code points and UTF-8 encoding.
- [[Strings]] — Immutable byte slices, UTF-8 indexing, len() in bytes vs utf8.RuneCountInString.
- [[Raw String Literals]] — Multi-line unescaped strings enclosed in backticks.
- [[Interpreted String Literals]] — Double-quoted strings with escape sequences (\n, \t, \x).
- [[String Internals (stringStruct)]] — Two-word struct: pointer to byte array and length in bytes.
- [[Numeric Conversions]] — Explicit type casting between numeric types without implicit widening.
- [[Overflow and Precision]] — Integer arithmetic overflow behavior and floating-point precision loss.
### 7. 📂 [[Composite Types|07. Composite Types]]
- [[Arrays]] — Fixed-length contiguous memory sequences with value semantics.
- [[Slices]] — Dynamic views over arrays with length, capacity, and pointer.
- [[Slice Header Internals]] — sliceHeader struct: unsafe.Pointer, int len, int cap.
- [[Slice Capacity and Growth]] — Slice append allocation formula and doubling thresholds.
- [[make() for Slices and Maps]] — Allocating and initializing slices, maps, and channels.
- [[Slice to Array Conversion]] — Pointer-to-array conversions and Go 1.20 slice-to-array casting.
- [[Slice Tricks]] — Idiomatic cut, delete, insert, push, pop, and copy patterns.
- [[Maps]] — Hash table implementation with O(1) average lookup, insert, and delete.
- [[Map Internals (hmap and bmap)]] — Buckets, overflow buckets, hash seeds, and incremental evacuation.
- [[Comma-Ok Idiom for Maps]] — Distinguishing between zero values and missing keys (v, ok := m[k]).
- [[Structs]] — User-defined aggregate types grouping named fields.
- [[Struct Memory Layout & Padding]] — CPU memory word alignment, padding bytes, and field ordering.
- [[Struct Tags & JSON]] — Field reflection metadata for serialization (json, xml, db).
- [[Embedding Structs]] — Composition and field/method promotion without inheritance.
- [[Empty Struct (struct{})]] — Zero-byte type for signaling channels and set data structures.
- [[Anonymous Structs]] — Ad-hoc inline struct definitions for local grouping or test cases.
### 8. 📂 [[Conditionals|08. Conditionals]]
- [[if Statement]] — Boolean condition branching.
- [[Short Statement in If]] — Scoped initialization statement before condition (if err := ...; err != nil).
- [[if-else Chains]] — Multi-way conditional branching.
- [[switch Statement]] — Multi-case matching with automatic break and fallthrough keyword.
- [[Tagless Switch]] — Evaluating arbitrary boolean conditions in switch cases.
- [[Type Switch]] — Dynamic type dispatching on interface values (switch v := x.(type)).
### 9. 📂 [[Loops & Iteration|09. Loops]]
- [[for Loop]] — Three-component loop (init; condition; post), while-style condition, infinite loop.
- [[for range]] — Iterating over slices, arrays, maps, strings, and channels.
- [[Range over Integer (Go 1.22+)]] — Syntactic sugar for counting loops (for i := range 10).
- [[Range over Func (Go 1.23+)]] — Custom user iterator functions yielding elements to for-range.
- [[Iterating Maps]] — Randomized iteration order in Go maps and sorted key patterns.
- [[Iterating Strings]] — Byte iteration vs Unicode rune decoding during string range.
- [[break Statement]] — Exiting the innermost loop or switch.
- [[continue Statement]] — Skipping to the next loop iteration.
- [[goto Statement]] — Unconditional jumps within function scope.
- [[Labeled break and continue]] — Breaking out of outer loops from nested inner loops or switches.
### 10. 📂 [[Functions & Closures|10. Functions]]
- [[Function Declarations]] — Signatures, parameter lists, return types, first-class citizen functions.
- [[Multiple Return Values]] — Idiomatic Go function signatures returning (result, error).
- [[Named Return Values]] — Naked returns and return value variable scoping.
- [[Variadic Functions]] — Passing variable numbers of arguments with ...T syntax.
- [[Anonymous Functions]] — Inline lambda functions and immediately invoked function expressions.
- [[Closures]] — Functions capturing variables from enclosing lexical scopes.
- [[Closure Internals (Heap Escape)]] — How the compiler moves captured variables to the heap.
- [[Call by Value]] — Go is strictly pass-by-value: passing pointers copies the pointer address.
- [[defer Statement]] — LIFO deferred function execution on function return.
- [[defer Ordering and Evaluation]] — Argument evaluation at defer time vs execution at return time.
- [[init() Function]] — Package initialization lifecycle and execution order across packages.
### 11. 📂 [[Pointers & Memory|11. Pointers]]
- [[Pointers Basics]] — Memory addresses, pointer types (*T), address-of operator (&), dereference (*).
- [[Pointer vs Value Semantics]] — Copying data vs sharing memory, mutation vs immutability.
- [[Pointers with Structs]] — Dot notation auto-dereferencing (p.Field equivalent to (*p).Field).
- [[Pointers with Slices & Maps]] — Reference types containing internal pointers vs pointer-to-slice.
- [[unsafe.Pointer and uintptr]] — Bypassing Go type safety, pointer arithmetic, and alignment.
- [[nil Pointer Dereference]] — Runtime panics on dereferencing nil and defensive nil checking patterns.
- [[Memory Management & Escape Analysis]] — Stack vs heap allocation rules determined at compile time.
- [[Garbage Collection Overview]] — Automatic memory reclamation with concurrent tricolor GC.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
