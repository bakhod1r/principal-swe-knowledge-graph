- [[Binary Size Optimization Matrix]] — Stripping symbols (-ldflags="-s -w"), removing DWARF, UPX compression caveats.

- [[Cross-Compilation with Zig Toolchain (CGO_ENABLED=1)]] — Using zig cc as cross-compiler for compiling Cgo across different OS and libc targets.

---
title: Core Commands & Code Gen
tags:
  - golang
  - toolchain
  - principal-swe
parent: "[[Go Toolchain & Developer Experience]]"
---

# Core Commands & Code Gen

Go CLI coordinator, code generation directives, multi-module workspaces, and internal tools.

```text
Core Commands & Code Gen
│
├── [[Core Go Commands]]
├── [[Code Generation (go:generate)]]
├── [[Build Tags (go:build)]]
├── [[go.work Workspaces]]
└── [[go tool Suite]]
```

---

## 🗂️ Topics

- [[Core Go Commands]] — go build, go install, go run, go clean, go doc, go version.
- [[Code Generation (go:generate)]] — Automating stringer, mock generators, protobuf compilation via go generate.
- [[Build Tags (go:build)]] — Conditional compilation based on OS, architecture, compiler tags, or custom tags.
- [[go.work Workspaces]] — Managing multi-module development environments with go work use/sync.
- [[go tool Suite]] — Executing compiler/linker internal tools (compile, link, nm, objdump, pprof, trace).

---

## 🔗 References
- ⬆️ Parent: [[Go Toolchain & Developer Experience]]
- 🎓 Root: [[Principal SWE]]
