---
title: Core Environment Variables
tags:
  - golang
  - environment
  - principal-swe
parent: "[[Settings Environment]]"
---

# Core Environment Variables

Fundamental Go directories, paths, binaries, workspace controls, experimental flags, and persistent environment configurations.

```text
Core Environment Variables
│
├── [[PATH and Go Binaries]]
├── [[GOROOT]]
├── [[GOPATH Workspace & Module Cache]]
├── [[GOBIN Binary Destination]]
├── [[GOENV Persistent Configuration]]
├── [[GOMODCACHE Module Cache]]
├── [[GOTOOLCHAIN Version Selection]]
├── [[GOWORK Environment Variable & Workspace Control]]
├── [[GOEXPERIMENT Flags (BoringCrypto, SwissMap, RangeFunc, AllocHeaders)]]
├── [[GOTELEMETRY Privacy & Crash Reporting (Go 1.23+)]]
└── [[Shell Startup & Profile Persistence]]
```

---

## 🗂️ Topics

- [[PATH and Go Binaries]] — OS search path resolution for Go binary executable and installed CLI tools.
- [[GOROOT]] — Root directory of the Go SDK distribution containing standard library source.
- [[GOPATH Workspace & Module Cache]] — Workspace structure, `pkg/mod` module cache, and `bin` output directories.
- [[GOBIN Binary Destination]] — Target directory where `go install` outputs compiled executable binaries.
- [[GOENV Persistent Configuration]] — Location of persistent environment variable settings file (`$HOME/.config/go/env`).
- [[GOMODCACHE Module Cache]] — Download cache for modules and toolchains; read-only layout and `go clean -modcache`.
- [[GOTOOLCHAIN Version Selection]] — Which Go toolchain actually runs the build; `auto`/`local`/`path` and the `toolchain` directive.
- [[GOWORK Environment Variable & Workspace Control]] — Explicitly specifying `go.work` file path or disabling workspace mode (`GOWORK=off`).
- [[GOEXPERIMENT Flags (BoringCrypto, SwissMap, RangeFunc, AllocHeaders)]] — Build-time compiler/runtime experimental feature toggles.
- [[GOTELEMETRY Privacy & Crash Reporting (Go 1.23+)]] — Go toolchain anonymous usage telemetry configuration (`on`/`off`/`local`).
- [[Shell Startup & Profile Persistence]] — Configuring environment variables in `.zshrc`, `.bashrc`, or `/etc/profile`.

---

## 🔗 References
- ⬆️ Parent: [[Settings Environment]]
- 📚 Module: `Go Environment & Commands`
