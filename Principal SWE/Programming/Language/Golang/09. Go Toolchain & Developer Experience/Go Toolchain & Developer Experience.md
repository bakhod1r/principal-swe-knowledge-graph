---
title: Go Toolchain & Developer Experience
tags:
  - golang
  - toolchain
  - principal-swe
parent: "[[Golang]]"
---

# 🛠️ Go Toolchain & Developer Experience

Go CLI toolchain: go build, go test, go generate, build tags, static analysis (golangci-lint), security (govulncheck), and Delve debugging.

```text
Go Toolchain & Developer Experience
│
├── [[Core Commands & Code Gen|01. Core Commands & Code Gen]]
│   ├── [[Core Go Commands]]
│   ├── [[Code Generation (//go:generate)]]
│   ├── [[Build Tags (//go:build)]]
│   ├── [[go.work Workspaces]]
│   └── [[go tool Suite]]
└── [[Quality, Security & Debugging|02. Quality, Security & Debugging]]
│   ├── [[go vet Static Analyzer]]
│   ├── [[golangci-lint Architecture]]
│   ├── [[Security (govulncheck)]]
│   ├── [[Debugging with Delve (dlv)]]
│   ├── [[Runtime Diagnostic Flags (GODEBUG)]]
│   └── [[Live Reloading (Air)]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Core Commands & Code Gen|01. Core Commands & Code Gen]]
- [[Core Go Commands]] — go build, go install, go run, go clean, go doc, go version.
- [[Code Generation (//go:generate)]] — Automating stringer, mock generators, protobuf compilation via go generate.
- [[Build Tags (//go:build)]] — Conditional compilation based on OS, architecture, compiler tags, or custom tags.
- [[go.work Workspaces]] — Managing multi-module development environments with go work use/sync.
- [[go tool Suite]] — Executing compiler/linker internal tools (compile, link, nm, objdump, pprof, trace).
### 2. 📂 [[Quality, Security & Debugging|02. Quality, Security & Debugging]]
- [[go vet Static Analyzer]] — Standard compiler static analysis checks for unreachable code, print format bugs.
- [[golangci-lint Architecture]] — Configuring multi-linter pipelines, fast caching, linters settings (.golangci.yml).
- [[Security (govulncheck)]] — Scanning dependencies for known CVEs using Go vulnerability database and call-graph analysis.
- [[Debugging with Delve (dlv)]] — Setting breakpoints, inspecting goroutine stacks, evaluating variables with dlv CLI/IDE.
- [[Runtime Diagnostic Flags (GODEBUG)]] — gctrace=1, schedtrace=1000, madvdontneed=1, asyncpreemptoff=1.
- [[Live Reloading (Air)]] — Hot-reloading Go services on file changes during local development.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
