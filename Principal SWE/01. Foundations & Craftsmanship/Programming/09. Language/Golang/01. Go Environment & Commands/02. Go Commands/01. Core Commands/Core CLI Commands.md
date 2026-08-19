---
title: Core CLI Commands
tags:
  - golang
  - commands
  - toolchain
  - principal-swe
parent: "[[Go Commands]]"
---

# Core CLI Commands

All primary top-level `go` subcommands: build flags, test frameworks, module management, environment inspection, and telemetry.

```text
Core CLI Commands
│
├── [[go build Command]]
├── [[go build -gcflags and -ldflags Deep Dive]]
├── [[go run Command]]
├── [[go install Command]]
├── [[go test Command]]
├── [[go test Benchmarking, Profiling & Coverage Flags]]
├── [[go fmt and gofumpt]]
├── [[go vet Static Analysis]]
├── [[go doc CLI Documentation]]
├── [[go clean & Cache Eviction]]
├── [[go mod Module Maintenance]]
├── [[go get Dependency Acquisition]]
├── [[go list Package Queries]]
├── [[go env Environment Inspection]]
├── [[go tool Toolchain Dispatch]]
├── [[go version Build Metadata]]
├── [[go telemetry CLI (Go 1.23+)]]
├── [[go bug Issue Reporting]]
└── [[go work Workspace Workflow]]
```

---

## 🗂️ Topics

- [[go build Command]] — Compiling packages and dependencies into executables or package archives.
- [[go build -gcflags and -ldflags Deep Dive]] — Compiler flags (`-gcflags="-m"` for escape analysis) and linker flags (`-ldflags="-s -w -X"`).
- [[go run Command]] — Compiling and executing Go source files on the fly in memory/temporary directories.
- [[go install Command]] — Compiling and installing executable binaries into `$GOBIN`.
- [[go test Command]] — Automated testing, subtest filtering, benchmark execution, and coverage profiling.
- [[go test Benchmarking, Profiling & Coverage Flags]] — Profiling flags (`-cpuprofile`, `-memprofile`, `-bench`, `-race`, `-coverprofile`).
- [[go fmt and gofumpt]] — Standard formatting with `gofmt` and strict formatting with `gofumpt`.
- [[go vet Static Analysis]] — Standard compiler static analysis checks for unreachable code and format bugs.
- [[go doc CLI Documentation]] — Extracting and displaying documentation for packages and symbols from the terminal.
- [[go clean & Cache Eviction]] — Removing compiled object files and evicting build/test caches (`go clean -cache -modcache`).
- [[go mod Module Maintenance]] — Module file maintenance: `download`, `edit`, `graph`, `init`, `tidy`, `vendor`, `verify`, `why`.
- [[go get Dependency Acquisition]] — Adding, upgrading, and downgrading module requirements in `go.mod`.
- [[go list Package Queries]] — Querying packages and modules with `-f` templates and `-json` output.
- [[go env Environment Inspection]] — Reading, persisting (`-w`), and unsetting (`-u`) toolchain configuration.
- [[go tool Toolchain Dispatch]] — Dispatching to bundled and on-demand toolchain binaries.
- [[go version Build Metadata]] — Toolchain version reporting and `-m` binary build metadata extraction.
- [[go telemetry CLI (Go 1.23+)]] — Viewing and managing local toolchain crash reports and usage telemetry.
- [[go bug Issue Reporting]] — Opening a pre-filled Go issue report with environment details.
- [[go work Workspace Workflow]] — Managing multi-module development environments locally with `go.work`.

---

## 🔗 References
- ⬆️ Parent: [[Go Commands]]
- 📚 Module: `Go Environment & Commands`
