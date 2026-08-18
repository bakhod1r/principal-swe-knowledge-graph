---
title: Core Go Environment
tags:
  - golang
  - basics
  - environment
parent: "[[Settings Environment]]"
---

# 🐹 Core Go Environment

Foundational variables governing the Go toolchain SDK installation, workspace, and module cache.

---

## 🗂️ Core Variables

- [[GOENV]] — Path to the Go environment configuration file.
- [[GOROOT]] — Root path where the Go SDK and standard library are installed. (Deep dive: `Go Source Code Structure`)
- [[GOPATH]] — User workspace, downloaded module cache, and installed tool binaries.
- [[GOBIN]] — Directory where `go install` outputs executable binaries.
- [[GOMOD]] — Path to the current active `go.mod` file.
- [[GOWORK]] — Path to the current active `go.work` file.
- [[GOMODCACHE]] — Directory where downloaded modules are stored (`$GOPATH/pkg/mod`).

---

## 🔗 Navigation
- ⬆️ Parent: [[Settings Environment]]
