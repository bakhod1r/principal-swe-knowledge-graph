---
title: WebAssembly, WASI & Alternative Targets
tags:
  - golang
  - wasm
  - principal-swe
parent: "[[Advanced Topics & Low-Level Go]]"
---

# WebAssembly, WASI & Alternative Targets

WebAssembly browser compilation (GOOS=js), server-side WASI (GOOS=wasip1), Proxy-Wasm, TinyGo for embedded IoT, and Gomobile bindings.

```text
WebAssembly, WASI & Alternative Targets
│
├── [[WebAssembly Compilation Architecture (GOOS=js GOARCH=wasm)]]
├── [[syscall-js Bridge & DOM Interop Mechanics]]
├── [[Zero-Copy Memory Sharing (Go Wasm to JS TypedArrays)]]
├── [[Server-Side Wasm with WASI (GOOS=wasip1 GOARCH=wasm)]]
├── [[Wasm Plugin Architecture in Envoy & Edge Proxies]]
├── `TinyGo Compiler Architecture for Microcontrollers & Web`
├── `Embedded IoT Programming with TinyGo (ESP32, RP2040, STM32)`
└── `Gomobile Native Bindings for iOS & Android`
```

---

## 🗂️ Topics

- [[WebAssembly Compilation Architecture (GOOS=js GOARCH=wasm)]] — Compiling Go binaries to .wasm, wasm_exec.js runtime glue, and browser execution lifecycle.
- [[syscall-js Bridge & DOM Interop Mechanics]] — js.Value, js.Global(), registering DOM event callbacks, and JavaScript promise integration in Go.
- [[Zero-Copy Memory Sharing (Go Wasm to JS TypedArrays)]] — Direct memory view access via js.CopyBytesToGo and js.CopyBytesToJS without serialization.
- [[Server-Side Wasm with WASI (GOOS=wasip1 GOARCH=wasm)]] — Go 1.21+ WASI standard support, executing Go Wasm in Wasmtime, Wasmer, and Docker Wasm.
- [[Wasm Plugin Architecture in Envoy & Edge Proxies]] — Building low-latency proxy filter extensions (Proxy-Wasm) for Envoy, Istio, and API gateways.
- `TinyGo Compiler Architecture for Microcontrollers & Web` — LLVM-based Go compiler, sub-100KB Wasm binaries, and zero-overhead embedded garbage collection.
- `Embedded IoT Programming with TinyGo (ESP32, RP2040, STM32)` — Real-time GPIO, I2C, SPI bus communication, and interrupt handlers on bare-metal silicon.
- `Gomobile Native Bindings for iOS & Android` — Generating native Objective-C/Swift and Java/Kotlin bindings from Go libraries (gomobile bind).

---

## 🔗 References
- ⬆️ Parent: [[Advanced Topics & Low-Level Go]]

