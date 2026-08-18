---
title: Language Basic
tags:
  - golang
  - backend
  - language
  - basics
parent: "[[Golang]]"
---

# 🔤 Language Basics

A comprehensive, deep-dive guide to the fundamentals of the Go programming language, toolchain, environment, runtime source structure, and core syntax.

```text
Language Basics
│
├── [[Settings Environment|01. Settings Environment]]
├── [[Go Commands|02. Go command]]
├── [[Dependencies|03. Dependencies]]
├── [[Go Source Code Structure|04. Go Source Code Structure]]
│
├── [[Variables & Constants|05. Variables & Constants]]
├── [[Data Types|06. Data Types]]
├── [[Composite Types|07. Composite Types]]
├── [[Conditionals|08. Conditionals]]
├── [[Loops|09. Loops]]
├── [[Functions|10. Functions]]
├── [[Pointers|11. Pointers]]
├── [[Methods & Interfaces|12. Methods & Interfaces]]
└── [[Generics|13. Generics]]
```

---

## 🗂️ Core Sections

### 1. ⚙️ [[Settings Environment|01. Settings Environment]]
- Setup, PATH configuration, and full environment variable architecture (`GOROOT`, `GOPATH`, `GOBIN`, `GOOS`, `GOARCH`, `GOCACHE`).

### 2. 💻 [[Go Commands|02. Go command]]
- Go CLI orchestrator: `build`, `run`, `test`, `install`, `vet`, `fmt`, `generate`, `work`, `tool`.

### 3. 📦 [[Dependencies|03. Dependencies]]
- Go Module system, `go.mod`, `go.sum`, MVS (Minimal Version Selection), checksum database (`GOSUMDB`), and private modules (`GOPRIVATE`).

### 4. 📂 [[Go Source Code Structure|04. Go Source Code Structure]]
- Anatomy of `$GOROOT`: standard library (`src/`), runtime engine (`src/runtime/`), compiler & linker (`src/cmd/`), and internal packages.

---

## 🗺️ Language Syntax & Concepts

### 5. 🏷️ [[Variables & Constants|05. Variables & Constants]]
- [[var vs :=]], [[Zero Values]], [[const and iota]], [[Scope and Shadowing]], [[Blank Identifier (_)]].

### 6. 🔢 [[Data Types|06. Data Types]]
- **Numeric Types**: [[Integers (Signed, Unsigned)]], [[Floating Points]], [[Complex Numbers]], [[Overflow and Precision]], [[Boolean]], [[Runes]].
- **Strings**: [[Raw String Literals]], [[Interpreted String Literals]], [[String Internals]].
- **Tooling & Operations**: [[Type Conversion]], [[Commands and Docs]].

### 7. 🧱 [[Composite Types|07. Composite Types]]
- **Arrays**: [[Arrays]].
- **Slices**: [[Capacity and Growth]], [[make()]], [[Slice to Array Conversion]], [[Array to Slice Conversion]], [[Slice Header Internals]], [[Slice Tricks]].
- **Maps**: [[Comma-Ok Idiom]], [[Map Internals]].
- **Structs**: [[Struct Tags & JSON]], [[Embedding Structs]], [[Struct Memory Layout & Padding]], [[Empty Struct]], [[Anonymous Structs]].

### 8. 🔀 [[Conditionals|08. Conditionals]]
- [[if]], [[if-else]], [[switch]], [[Short Statement in If]].

### 9. 🔁 [[Loops|09. Loops]]
- [[for loop]], [[for range]], [[Iterating Maps]], [[Iterating Strings]], [[break]], [[continue]], [[goto]], [[Labeled break and continue]].

### 10. ⚡ [[Functions|10. Functions]]
- [[Function Basics]], [[Multiple Return Values]], [[Named Return Values]], [[Variadic Functions]], [[Anonymous Functions]], [[Closures]], [[Closure Internals]], [[Call by Value]], [[defer Basics]], [[init Function]].

### 11. 🎯 [[Pointers|11. Pointers]]
- [[Pointers Basics]], [[Pointers with Structs]], [[With Maps & Slices]], [[unsafe.Pointer]], [[nil Pointer Dereference]], [[Memory Management]], [[Garbage Collection]].

### 12. 🧩 [[Methods & Interfaces|12. Methods & Interfaces]]
- **Methods**: [[Methods vs Functions]], [[Pointer Receivers]], [[Value Receivers]].
- **Interfaces**: [[Interfaces Basics]], [[Empty Interfaces]], [[Embedding Interfaces]], [[Type Assertions]], [[Type Switch]].

### 13. 🧬 [[Generics|13. Generics]]
- [[Why Generics?]], [[Generic Functions]], [[Generic Types, Interfaces]], [[Type Constraints]], [[Type Inference]].

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 🎓 Root: [[Principal SWE]]
