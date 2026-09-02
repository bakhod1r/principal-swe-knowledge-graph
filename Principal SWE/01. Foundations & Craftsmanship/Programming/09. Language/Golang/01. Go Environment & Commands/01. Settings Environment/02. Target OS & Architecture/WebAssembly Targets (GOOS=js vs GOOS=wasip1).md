---
title: "WebAssembly Targets (GOOS=js vs GOOS=wasip1)"
tags:
  - golang
  - webassembly
  - wasi
  - wasip1
  - principal-swe
parent: "[[Target OS & Architecture]]"
---
# WebAssembly Targets: `GOOS=js` vs `GOOS=wasip1`

Go’da WebAssembly uchun **ikkita muhim target** mavjud:

```bash
GOOS=js      GOARCH=wasm
GOOS=wasip1  GOARCH=wasm
```

Ikkalasi ham `.wasm` binary yaratadi, lekin **host environment va system interface butunlay boshqacha**. `js` browser/JavaScript execution modeliga, `wasip1` esa WASI Preview 1 syscall modeliga mo‘ljallangan. `wasip1` Go 1.21’da qo‘shilgan. ([Go.dev](https://go.dev/wiki/WebAssembly?utm_source=chatgpt.com "Go Wiki: WebAssembly - The Go Programming Language"))

---

## 1. Mental Model

Eng muhim farq:

```text
                    Go source code
                          │
                          ▼
                    GOARCH=wasm
                          │
             ┌────────────┴────────────┐
             │                         │
         GOOS=js                  GOOS=wasip1
             │                         │
             ▼                         ▼
      JavaScript host              WASI host
             │                         │
       Browser / Node.js      Wasmtime / Wazero /
                              WasmEdge / Wasmer / Node
```

Ya'ni:

> **`GOARCH=wasm` — instruction architecture.**  
> **`GOOS` — Go program qanday host/system API bilan gaplashishini belgilaydi.**

Go toolchain `wasm`ni `GOARCH`, `js` yoki `wasip1`ni esa `GOOS` sifatida ko‘radi. ([Go.dev](https://go.dev/doc/install/source?utm_source=chatgpt.com "Installing Go from source - The Go Programming Language"))

---

# 2. `GOOS=js`

```bash
GOOS=js GOARCH=wasm go build -o main.wasm
```

Bu Go'ning original WebAssembly targetidir.

Go 1.11'dan mavjud va asosiy maqsadi:

> **WebAssembly module → JavaScript runtime → Browser**

Go program JavaScript bilan `syscall/js` orqali interoperate qiladi. ([Go.dev](https://go.dev/wiki/WebAssembly?utm_source=chatgpt.com "Go Wiki: WebAssembly - The Go Programming Language"))

### Architecture

```text
Browser
│
├── JavaScript
│   │
│   ├── DOM
│   ├── fetch()
│   ├── Web APIs
│   └── events
│
└── Go WASM
    │
    ├── Go runtime
    ├── goroutines
    ├── GC
    └── syscall/js
```

Masalan:

```go
package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello from Go")

	document := js.Global().Get("document")
	body := document.Get("body")

	body.Set("innerText", "Hello from Go WASM!")

	select {}
}
```

Bu yerda:

```go
js.Global()
```

orqali Go JavaScript global environment bilan gaplashmoqda.

Shuning uchun `js/wasm`:

- browser UI
    
- DOM
    
- JavaScript APIs
    
- browser events
    
- Web APIs
    

uchun tabiiy target.

---

# 3. `GOOS=wasip1`

```bash
GOOS=wasip1 GOARCH=wasm go build -o main.wasm
```

Bu esa browserga bog‘liq emas.

WASI — **WebAssembly System Interface**.

Mental model:

```text
Go program
    │
    ▼
Wasm module
    │
    ▼
WASI imports
    │
    ▼
Wasm runtime
    │
    ├── filesystem
    ├── clock
    ├── random
    ├── stdout
    └── other host capabilities
```

Go `wasip1` targeti **WASI Preview 1 (`wasi_snapshot_preview1`)** API'ga target qiladi. Go 1.21 bu targetni qo‘shgan. ([Go.dev](https://go.dev/blog/wasi?utm_source=chatgpt.com "WASI support in Go - The Go Programming Language"))

Masalan:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello WASI")

	data, err := os.ReadFile("/config.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
```

Bu yerda Go filesystem API ishlatmoqda.

Lekin muhim nuance:

```text
WASI filesystem
        ≠
"host filesystem automatically available"
```

Host runtime filesystemni **explicitly expose** qilishi kerak.

Masalan Wasmtime'da:

```bash
wasmtime run \
  --env PWD=/ \
  --dir .::/ \
  main.wasm
```

Go documentation ham WASI runtime'da filesystem preopen/mapping konfiguratsiyasiga e'tibor berishni talab qiladi. ([Go.dev](https://go.dev/wiki/WebAssembly?utm_source=chatgpt.com "Go Wiki: WebAssembly - The Go Programming Language"))

---

# 4. Eng muhim farq

||`js/wasm`|`wasip1/wasm`|
|---|---|---|
|`GOOS`|`js`|`wasip1`|
|`GOARCH`|`wasm`|`wasm`|
|Primary host|Browser / JS|WASI runtime|
|JavaScript|**Native integration**|Optional/host-dependent|
|DOM|✅|❌|
|`syscall/js`|✅|❌|
|Filesystem|Browser/JS model|WASI|
|stdout/stderr|JS environment|WASI|
|Server-side Wasm|Possible, less natural|**Natural**|
|Browser frontend|**Excellent fit**|Poor fit|
|Sandboxed compute|Good|**Excellent fit**|
|Wasmtime|Not the primary model|✅|
|Wazero|Not the primary model|✅|
|Node.js|✅|✅|
|Cloud/server Wasm|Possible|**Typical use case**|

WASI hostlar qatoriga Wasmtime, Wazero, WasmEdge, Wasmer va Node.js kiradi. ([Go.dev](https://go.dev/blog/wasi?utm_source=chatgpt.com "WASI support in Go - The Go Programming Language"))

---

# 5. Why `GOOS` matters

Bu oddiy naming emas.

Masalan:

```go
// file.go

func Load() {
	// ...
}
```

platform-specific implementation:

```text
config_js.go
config_wasip1.go
config_linux.go
```

Build:

```bash
GOOS=js GOARCH=wasm
```

natijasida:

```text
config_js.go       included
config_wasip1.go   excluded
config_linux.go    excluded
```

Aksincha:

```bash
GOOS=wasip1 GOARCH=wasm
```

da:

```text
config_wasip1.go   included
```

Go 1.21 `*_wasip1.go` build constraint semanticsini ham qo‘shgan. ([Go.dev](https://go.dev/doc/go1.21?utm_source=chatgpt.com "Go 1.21 Release Notes - The Go Programming Language"))

Bu architecture design'da juda foydali.

---

# 6. Same `.wasm`, different contract

Bu juda muhim mental model.

Ko‘pchilik:

> "Ikkalasi ham WebAssembly, demak interchangeable."

deb o‘ylaydi.

**Bu noto‘g‘ri.**

`.wasm` faqat binary format.

Real execution contract:

```text
Wasm binary
+
imports
+
exports
+
host capabilities
+
runtime semantics
```

`js/wasm`:

```text
Go
 ↓
Go runtime
 ↓
JavaScript imports
 ↓
Browser / Node
```

`wasip1/wasm`:

```text
Go
 ↓
Go runtime
 ↓
WASI imports
 ↓
WASI runtime
```

Shuning uchun `js` uchun build qilingan module'ni oddiy `wasmtime` executable sifatida ishlatish mumkin deb o‘ylash xato.

---

# 7. Build examples

## Browser

```bash
GOOS=js GOARCH=wasm go build -o main.wasm
```

Keyin Go distribution'dagi `wasm_exec.js` kerak bo‘ladi:

```bash
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" .
```

va JavaScript:

```javascript
const go = new Go();

WebAssembly.instantiateStreaming(
    fetch("main.wasm"),
    go.importObject
).then(result => {
    go.run(result.instance);
});
```

Go documentation `wasm_exec.js` versioni compiler version bilan mos bo‘lishi kerakligini alohida ta'kidlaydi. ([Go.dev](https://go.dev/wiki/WebAssembly?utm_source=chatgpt.com "Go Wiki: WebAssembly - The Go Programming Language"))

---

## WASI

```bash
GOOS=wasip1 GOARCH=wasm go build -o main.wasm
```

Keyin:

```bash
wasmtime main.wasm
```

Oddiy program uchun shu kifoya qilishi mumkin. ([Go.dev](https://go.dev/blog/wasi?utm_source=chatgpt.com "WASI support in Go - The Go Programming Language"))

---

# 8. `GOOS=js` qachon tanlanadi?

### Browser application

Masalan:

```text
Go
 ↓
Wasm
 ↓
Browser
 ↓
DOM
 ↓
User
```

Use cases:

- browser-based Go application
    
- computational frontend
    
- Go-based UI logic
    
- WebAssembly demos
    
- browser-side CPU-heavy operations
    

Agar requirement:

> "Go code browser ichida ishlashi va JavaScript/DOM bilan bevosita ishlashi kerak"

bo‘lsa:

```bash
GOOS=js GOARCH=wasm
```

---

# 9. `GOOS=wasip1` qachon?

Agar:

> "Men portable sandboxed executable xohlayman"

bo‘lsa:

```bash
GOOS=wasip1 GOARCH=wasm
```

Masalan:

```text
API Server
    │
    ▼
Wasm runtime
    │
    ▼
User plugin
    │
    ├── CPU
    ├── limited filesystem
    └── controlled host capabilities
```

Bu plugin architecture uchun juda qiziq.

Masalan:

```text
Application
    │
    ├── Plugin A → plugin.wasm
    ├── Plugin B → plugin.wasm
    └── Plugin C → plugin.wasm
```

Host runtime plugin'ga faqat kerakli capabilities beradi.

Bu:

- sandboxing
    
- plugins
    
- edge computing
    
- serverless workloads
    
- portable execution
    
- untrusted/less-trusted code
    

uchun kuchli model.

---

# 10. Go 1.24+ important evolution

WASI faqat "`main()` run qilib exit qilish" bilan cheklanib qolmayapti.

Go 1.24:

```go
//go:wasmexport add
func add(a, b int32) int32 {
	return a + b
}
```

orqali host application Go functionini Wasm export sifatida chaqirishi mumkin. ([Go.dev](https://go.dev/blog/wasmexport?utm_source=chatgpt.com "Extensible Wasm Applications with Go - The Go Programming Language"))

Bundan tashqari:

```bash
GOOS=wasip1 GOARCH=wasm \
go build -buildmode=c-shared -o reactor.wasm
```

orqali **WASI reactor/library** yaratish mumkin.

Mental model:

### Command

```text
host
  ↓
_start
  ↓
main()
  ↓
exit
```

### Reactor

```text
host
  ↓
_initialize
  ↓
exported function
  ↓
exported function
  ↓
exported function
```

Bu plugin/embedded Wasm architecture uchun ancha muhim. ([Go.dev](https://go.dev/blog/wasmexport?utm_source=chatgpt.com "Extensible Wasm Applications with Go - The Go Programming Language"))

---

# 11. Production architecture perspective

Men buni quyidagicha eslab qolishni tavsiya qilaman:

```text
                 WebAssembly
                     │
          ┌──────────┴──────────┐
          │                     │
        js/wasm              wasip1/wasm
          │                     │
          ▼                     ▼
   "Web platform"          "System interface"
          │                     │
          ▼                     ▼
 JavaScript / Browser          WASI
          │                     │
          ▼                     ▼
      DOM/Web APIs        Filesystem/Clock/etc.
```

### `js`

**Host = JavaScript ecosystem**

### `wasip1`

**Host = WASI ecosystem**

Bu distinction'ni yaxshi tushunish `Wasm`ni faqat "browser technology" deb ko‘rishdan chiqaradi.

---

# 12. Common mistakes

### ❌ Mistake 1

```bash
GOOS=wasip1 GOARCH=wasm
```

ni browser target deb o‘ylash.

**Noto‘g‘ri.**

WASI browser API emas.

---

### ❌ Mistake 2

```go
import "syscall/js"
```

ni `wasip1`da ishlatish.

`syscall/js` JavaScript integration uchun mo‘ljallangan.

---

### ❌ Mistake 3

WASI filesystem:

```text
os.ReadFile("/foo")
```

qilsa, host'dagi `/foo` avtomatik mavjud deb o‘ylash.

**Noto‘g‘ri.**

Host capability/preopen konfiguratsiyasi kerak. ([Go.dev](https://go.dev/wiki/WebAssembly?utm_source=chatgpt.com "Go Wiki: WebAssembly - The Go Programming Language"))

---

### ❌ Mistake 4

"WebAssembly = browser"

Bu eski mental model.

Bugungi architecture:

```text
WebAssembly
├── Browser
├── WASI runtimes
├── Edge
├── Server
├── Plugins
└── Sandboxed execution
```

WASI aynan WebAssembly'ni browserdan tashqarida portable system interface bilan ishlatish uchun paydo bo‘lgan. ([Go.dev](https://go.dev/blog/wasi?utm_source=chatgpt.com "WASI support in Go - The Go Programming Language"))

---

# 13. Principal Engineer mental model

Architecture tanlashda `GOOS`dan boshlamang.

Avval:

```text
Who is the host?
        │
        ▼
What capabilities are needed?
        │
        ▼
What security boundary is required?
        │
        ▼
What API does the host expose?
        │
        ▼
Which Wasm target fits?
```

### Browser-centric

```text
Need DOM / JS / browser APIs
            ↓
        js/wasm
```

### Host-centric / sandboxed

```text
Need portable module + system capabilities
            ↓
        wasip1/wasm
```

### Plugin-centric

```text
Need host ↔ module function calls
            ↓
WASI + wasm exports/imports
```

---

## Final cheat sheet

```text
GOOS=js GOARCH=wasm
        ↓
Browser / JavaScript
        ↓
DOM + Web APIs + syscall/js


GOOS=wasip1 GOARCH=wasm
        ↓
WASI runtime
        ↓
Portable sandboxed application
        ↓
Filesystem / clock / random / stdio / host capabilities
```

**Eng muhim takeaway:**

> `GOARCH=wasm` — **"qaysi instruction format?"**  
> `GOOS=js` — **"JavaScript host bilan qanday ishlayman?"**  
> `GOOS=wasip1` — **"WASI system interface orqali host bilan qanday ishlayman?"**

Go'ning current source/toolchain'ida `js` va `wasip1` alohida `GOOS` sifatida mavjud, ikkalasi ham `GOARCH=wasm` bilan ishlaydi. ([Go.dev](https://go.dev/doc/install/source?utm_source=chatgpt.com "Installing Go from source - The Go Programming Language"))

[Go WebAssembly documentation](https://go.dev/wiki/WebAssembly?utm_source=chatgpt.com) — amaliy `js/wasm` va `wasip1/wasm` ishlatish uchun asosiy reference. ([Go.dev](https://go.dev/wiki/WebAssembly?utm_source=chatgpt.com "Go Wiki: WebAssembly - The Go Programming Language"))

[Go WASI support blog](https://go.dev/blog/wasi?utm_source=chatgpt.com) — WASI architecture va Go integration'ini chuqurroq tushunish uchun. ([Go.dev](https://go.dev/blog/wasi?utm_source=chatgpt.com "WASI support in Go - The Go Programming Language"))

## 🔗 References
- ⬆️ Parent: [[Target OS & Architecture]]
- 📚 Module: `Go Environment & Commands`
