---
title: WebAssembly & Alternative Targets
tags:
  - golang
  - wasm
  - principal-swe
parent: "[[Golang]]"
---

# 🌐 WebAssembly & Alternative Targets

Compiling Go to WebAssembly (GOOS=js GOARCH=wasm), WASI (wasip1), TinyGo for embedded microcontrollers and lightweight WASM binaries.

```text
WebAssembly & Alternative Targets
│
├── [[WebAssembly (WASM)|01. WebAssembly (WASM)]]
│   ├── [[GOOS=js GOARCH=wasm Browser Execution]]
│   ├── [[syscall-js DOM & JavaScript Bridge]]
│   ├── [[WASI (wasip1) Server-Side Execution]]
│   ├── [[WASM Binary Size Optimization]]
│   └── [[WASM Production Edge Deployments]]
└── [[Embedded & Microcontrollers|02. Embedded & Microcontrollers]]
│   ├── [[TinyGo LLVM Compiler Architecture]]
│   └── [[Embedded GPIO & Hardware Peripherals]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[WebAssembly (WASM)|01. WebAssembly (WASM)]]
- [[GOOS=js GOARCH=wasm Browser Execution]] — Compiling Go to browser WebAssembly, wasm_exec.js runtime bridge, DOM interop.
- [[syscall-js DOM & JavaScript Bridge]] — Interacting with browser window, document, creating JS callbacks in Go.
- [[WASI (wasip1) Server-Side Execution]] — Server-side WASM execution via wasmtime/wasmer using GOOS=wasip1 GOARCH=wasm.
- [[WASM Binary Size Optimization]] — Stripping symbols (-ldflags="-s -w"), gzip compression, tinygo optimization.
- [[WASM Production Edge Deployments]] — Deploying Go WebAssembly on Cloudflare Workers, Fastly Compute, and plugin engines.
### 2. 📂 [[Embedded & Microcontrollers|02. Embedded & Microcontrollers]]
- [[TinyGo LLVM Compiler Architecture]] — LLVM-based TinyGo compiler, custom tiny runtime, memory-constrained optimization.
- [[Embedded GPIO & Hardware Peripherals]] — Controlling LEDs, sensors, I2C, SPI, UART on Arduino, ESP32, and Raspberry Pi Pico.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
