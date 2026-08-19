---
title: Core CLI Commands
tags:
  - review
  - golang
  - toolchain
  - principal-swe
parent: "[[Go Commands]]"
---

# Core CLI Commands

All 19 top-level `go` subcommands: build, run, install, test, fmt, vet, doc, clean, work, mod, get, list, env, tool, version, and bug (generate, fix, and telemetry are covered under Go Toolchain & Developer Experience).

```text
Core CLI Commands
│
├── [[go build Command]]
├── [[go run Command]]
├── [[go install Command]]
├── [[go test Command]]
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
├── [[go bug Issue Reporting]]
└── [[go work Workspace Workflow]]
```

---

## 🗂️ Topics

- [[go build Command]] — Compiling packages and dependencies into executables or package archives.
- [[go run Command]] — Compiling and executing Go source files on the fly in memory/temporary directories.
- [[go install Command]] — Compiling and installing executable binaries into $GOBIN.
- [[go test Command]] — Automated testing, subtest filtering, benchmark execution, and coverage profiling.
- [[go fmt and gofumpt]] — Standard formatting with gofmt and strict formatting with gofumpt.
- [[go vet Static Analysis]] — Standard compiler static analysis checks for unreachable code and format bugs.
- [[go doc CLI Documentation]] — Extracting and displaying documentation for packages and symbols from the terminal.
- [[go clean & Cache Eviction]] — Removing compiled object files and evicting build/test caches.
- [[go work Workspace Workflow]] — Managing multi-module development environments locally.
- [[go mod Module Maintenance]] — Module file maintenance: `download`, `edit`, `graph`, `init`, `tidy`, `vendor`, `verify`, `why`.
- [[go get Dependency Acquisition]] — Adding, upgrading, and downgrading module requirements in `go.mod`.
- [[go list Package Queries]] — Querying packages and modules with `-f` templates and `-json` output.
- [[go env Environment Inspection]] — Reading, persisting (`-w`), and unsetting (`-u`) toolchain configuration.
- [[go tool Toolchain Dispatch]] — Dispatching to bundled and on-demand toolchain binaries.
- [[go version Build Metadata]] — Toolchain version reporting and `-m` binary build metadata extraction.
- [[go bug Issue Reporting]] — Opening a pre-filled Go issue report with environment details.

---

## 🔗 References
- ⬆️ Parent: [[Go Commands]]

