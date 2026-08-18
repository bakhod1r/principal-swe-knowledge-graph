---
title: CGO & External Toolchain Environment
tags:
  - golang
  - basics
  - environment
  - cgo
parent: "[[Settings Environment]]"
---

# 🔌 CGO & External Toolchain Environment

Configuration for C/C++ language interoperability, external compilers, linkers, and flags.

---

## 🗂️ CGO & Host Toolchain Variables

- [[CGO_ENABLED]] — Enables (`1`) or disables (`0`) Cgo compilation.
- [[CGO_CFLAGS]] — Additional flags passed to the host C compiler.
- [[CGO_CPPFLAGS]] — Additional flags passed to the C preprocessor.
- [[CGO_CXXFLAGS]] — Additional flags passed to the host C++ compiler.
- [[CGO_LDFLAGS]] — Additional flags passed to the host linker.
- [[CC]] — Command to invoke the C compiler (`gcc`, `clang`).
- [[CXX]] — Command to invoke the C++ compiler (`g++`, `clang++`).
- [[AR]] — Command to invoke the archive tool (`ar`).
- [[PKG_CONFIG]] — Path to the `pkg-config` utility.

---

## 🔗 Navigation
- ⬆️ Parent: [[Settings Environment]]
