- [[WebAssembly Component Model & WASI 0.2]] — WASI Preview 2, Component Model (WIT interfaces), and composable WASM modules.

---
title: WebAssembly (WASM)
tags:
  - golang
  - wasm
  - principal-swe
parent: "[[WebAssembly & Alternative Targets]]"
---

# WebAssembly (WASM)

Browser WASM execution, syscall/js DOM manipulation, WASI server-side execution.

```text
WebAssembly (WASM)
│
├── [[GOOS=js GOARCH=wasm Browser Execution]]
├── [[syscall-js DOM & JavaScript Bridge]]
├── [[WASI (wasip1) Server-Side Execution]]
├── [[WASM Binary Size Optimization]]
└── [[WASM Production Edge Deployments]]
```

---

## 🗂️ Topics

- [[GOOS=js GOARCH=wasm Browser Execution]] — Compiling Go to browser WebAssembly, wasm_exec.js runtime bridge, DOM interop.
- [[syscall-js DOM & JavaScript Bridge]] — Interacting with browser window, document, creating JS callbacks in Go.
- [[WASI (wasip1) Server-Side Execution]] — Server-side WASM execution via wasmtime/wasmer using GOOS=wasip1 GOARCH=wasm.
- [[WASM Binary Size Optimization]] — Stripping symbols (-ldflags="-s -w"), gzip compression, tinygo optimization.
- [[WASM Production Edge Deployments]] — Deploying Go WebAssembly on Cloudflare Workers, Fastly Compute, and plugin engines.

---

## 🔗 References
- ⬆️ Parent: [[WebAssembly & Alternative Targets]]
- 🎓 Root: [[Principal SWE]]
