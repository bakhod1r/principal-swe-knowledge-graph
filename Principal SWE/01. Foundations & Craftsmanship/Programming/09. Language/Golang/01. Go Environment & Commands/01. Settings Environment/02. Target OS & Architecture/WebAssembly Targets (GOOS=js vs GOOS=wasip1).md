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
# WebAssembly Targets (`GOOS=js` vs `GOOS=wasip1`) in Go

Go supports WebAssembly (`wasm`) as a compilation target, but there are **two fundamentally different execution environments**:

- `GOOS=js GOARCH=wasm` — WebAssembly running through a **JavaScript host**, typically a browser or Node.js.
    
- `GOOS=wasip1 GOARCH=wasm` — WebAssembly targeting **WASI Preview 1**, designed for non-browser WASM runtimes.
    

The important mental model is:

> **`GOARCH=wasm` describes the CPU/ISA target. `GOOS` describes the operating-system/host environment and therefore the available runtime APIs.**

---

## 1. Why Does Go Have Two WASM Targets?

WebAssembly itself does not provide things such as:

- files
    
- sockets
    
- environment variables
    
- clocks
    
- processes
    
- JavaScript APIs
    

Those capabilities come from the **host environment**.

Go therefore needs to know what kind of host its WebAssembly program will run in.

```text
                    Go source
                       │
                       ▼
                GOARCH=wasm
                       │
             ┌─────────┴─────────┐
             │                   │
        GOOS=js             GOOS=wasip1
             │                   │
             ▼                   ▼
      JavaScript host        WASI host
      Browser / Node         Wasmtime / Wasmer /
                             WasmEdge / etc.
```

So these are **not merely two ways of compiling the same WASM binary**.

They target different host APIs and runtime assumptions.

---

# 2. `GOOS=js GOARCH=wasm`

This target is designed for environments where **JavaScript is the host integration layer**.

Build:

```bash
GOOS=js GOARCH=wasm go build -o main.wasm .
```

The resulting program normally requires Go's JavaScript WebAssembly runtime support.

A common browser setup looks conceptually like:

```text
Browser
 ├── JavaScript
 │    └── Go WASM runtime / wasm_exec.js
 │          └── main.wasm
 │
 └── Web APIs
      ├── DOM
      ├── fetch
      ├── WebSocket
      ├── localStorage
      └── etc.
```

Go exposes JavaScript interoperability through:

```go
import "syscall/js"
```

For example:

```go
package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello from Go WASM")

	window := js.Global()
	fmt.Println(window.Get("location"))
}
```

This is fundamentally different from ordinary Go code running on Linux.

The Go program is interacting with **JavaScript objects**, not directly with browser APIs.

---

# 3. `GOOS=wasip1 GOARCH=wasm`

`wasip1` targets **WASI Preview 1**.

Build:

```bash
GOOS=wasip1 GOARCH=wasm go build -o main.wasm .
```

The execution model becomes:

```text
WASI Runtime
    │
    ├── stdin
    ├── stdout
    ├── filesystem capabilities
    ├── environment
    ├── clocks
    └── other WASI APIs
         │
         ▼
      Go WASM
```

Instead of relying on JavaScript, the program interacts with the **WASI environment**.

This makes `wasip1` much more appropriate for server-side or sandboxed WebAssembly runtimes.

Examples include runtimes such as:

- Wasmtime
    
- Wasmer
    
- WasmEdge
    

The exact capabilities depend on the runtime and how it configures WASI.

---

# 4. The Most Important Difference

A useful comparison:

||`js/wasm`|`wasip1/wasm`|
|---|---|---|
|Host|JavaScript|WASI|
|Browser|Yes|Generally not the primary target|
|Node.js|Yes|No JavaScript dependency required|
|Wasmtime|Not the intended model|Yes|
|JavaScript interop|`syscall/js`|No|
|DOM|Via JS/browser|No|
|WASI APIs|Not the primary interface|Yes|
|Server-side WASM|Possible|Natural fit|
|CLI-style WASM|Awkward|Better fit|
|Browser UI|Natural fit|Not the target|

The key distinction is:

```text
js/wasm

Go
 ↓
JavaScript host
 ↓
Browser / Node APIs
```

versus:

```text
wasip1/wasm

Go
 ↓
WASI
 ↓
WASM runtime
```

---

# 5. Why `GOOS` Matters

This is an important Go cross-compilation mental model.

You might normally think:

```bash
GOOS=linux GOARCH=amd64
```

means:

> Compile for x86-64 Linux.

Similarly:

```bash
GOOS=windows GOARCH=amd64
```

means:

> Compile for x86-64 Windows.

For WebAssembly:

```bash
GOOS=js GOARCH=wasm
```

means:

> Compile WebAssembly assuming a JavaScript-based host environment.

While:

```bash
GOOS=wasip1 GOARCH=wasm
```

means:

> Compile WebAssembly assuming a WASI Preview 1 host environment.

So `GOOS` isn't literally saying that WebAssembly **is an operating system**.

It selects the **host/OS abstraction expected by the Go runtime**.

---

# 6. APIs Available to Your Program

This is where the difference becomes practical.

Consider:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("config.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
```

Whether this works depends heavily on the target.

With:

```bash
GOOS=wasip1 GOARCH=wasm
```

filesystem operations can be provided through WASI, assuming the runtime grants the necessary filesystem capabilities.

For example, a WASI runtime can conceptually run:

```text
main.wasm
   │
   └── open("config.txt")
          │
          ▼
        WASI
          │
          ▼
   sandboxed filesystem
```

This is one of the major advantages of WASI.

---

# 7. JavaScript Interoperability

With:

```bash
GOOS=js GOARCH=wasm
```

you can use:

```go
import "syscall/js"
```

Example:

```go
package main

import "syscall/js"

func main() {
	console := js.Global().Get("console")
	console.Call("log", "Hello from Go")
}
```

This assumes a JavaScript host.

You can access JavaScript globals:

```go
js.Global()
```

and call JavaScript functions:

```go
value.Call("someFunction")
```

This makes `js/wasm` particularly useful when Go is part of a browser application.

---

# 8. Browser Architecture

A typical Go browser application looks like:

```text
                    Browser
                       │
            ┌──────────┴──────────┐
            │                     │
       JavaScript              Web APIs
            │                     │
            │              ┌──────┼──────┐
            │              │      │      │
            │             DOM   fetch  WebSocket
            │
            ▼
       Go WASM
       GOOS=js
       GOARCH=wasm
```

Go provides computation/business logic while JavaScript often acts as the bridge to browser-specific capabilities.

This is useful for:

- computationally heavy logic
    
- existing Go libraries
    
- parsers
    
- cryptographic operations
    
- algorithms
    
- simulation
    
- image/data processing
    

But it doesn't magically turn Go into a native browser language.

---

# 9. WASI Architecture

A server-side WASI application looks more like:

```text
             Host Application
                    │
                    ▼
              WASM Runtime
                    │
             ┌──────┴──────┐
             │             │
            WASI       sandbox
             │
       ┌─────┼─────┐
       │     │     │
      FS    env   clocks
       │
       ▼
   Go WASM module
   GOOS=wasip1
   GOARCH=wasm
```

This architecture is attractive for:

- plugins
    
- sandboxed execution
    
- server-side WASM
    
- untrusted workloads
    
- portable modules
    
- embedded scripting/plugin systems
    
- edge computing
    

---

# 10. A Critical Security Mental Model

WASI does **not** mean:

> "WASM automatically gets access to the filesystem."

Instead:

> **The runtime decides what capabilities the module receives.**

This is capability-oriented sandboxing.

Conceptually:

```text
WASM module
     │
     ├── no filesystem capability
     │       → cannot access files
     │
     ├── read-only directory
     │       → can read only that directory
     │
     └── broader capability
             → more access
```

For production systems, this is extremely important.

Never reason:

> "It's WASM, therefore it is secure."

Instead reason:

> "What capabilities did the host expose to this module?"

---

# 11. `GOOS=js` Has a Different Security Boundary

In browser WASM:

```text
Browser security model
        +
JavaScript security model
        +
WASM sandbox
```

Your Go WASM code doesn't get arbitrary OS access.

For example, it cannot simply do:

```go
os.ReadFile("/etc/passwd")
```

and expect access to the host operating system.

The browser controls what capabilities are exposed.

---

# 12. Runtime Differences

One subtle but important point:

**Compilation target and runtime are different concepts.**

You compile:

```bash
GOOS=wasip1 GOARCH=wasm go build
```

but execution happens in a runtime such as:

```text
Wasmtime
Wasmer
WasmEdge
...
```

The runtime determines many operational details.

Similarly:

```bash
GOOS=js GOARCH=wasm
```

doesn't mean the `.wasm` file independently contains an entire browser.

It expects a JavaScript environment and Go's WASM runtime support.

---

# 13. `wasm_exec.js`

For the JavaScript target, Go historically provides JavaScript-side runtime support through `wasm_exec.js`.

Conceptually:

```text
index.html
    │
    ▼
JavaScript
    │
    ├── loads wasm_exec.js
    │
    └── loads main.wasm
             │
             ▼
          Go runtime
             │
             ▼
          Go code
```

The exact location/version should be taken from the installed Go distribution rather than copied blindly from an old tutorial.

---

# 14. Why You Should Not Choose Based on "WASM"

A common beginner mistake is:

> "I need WebAssembly, therefore I use `GOOS=wasm`."

That's incomplete.

The actual decision should be:

```text
Where will this WASM module run?
             │
       ┌─────┴─────┐
       │           │
   JS host       WASI host
       │           │
     js/wasm    wasip1/wasm
```

### Browser application

Use:

```bash
GOOS=js GOARCH=wasm
```

### JavaScript/Node integration

Usually:

```bash
GOOS=js GOARCH=wasm
```

### WASI runtime

Use:

```bash
GOOS=wasip1 GOARCH=wasm
```

---

# 15. Common Mistake: Treating WASI as "Browser WASM"

WASI and browser APIs solve different problems.

Browser:

```text
DOM
fetch
Web APIs
JavaScript
```

WASI:

```text
filesystem
stdin/stdout
environment
clocks
process-like host capabilities
```

WASI is closer to:

> **A standardized system interface for WebAssembly runtimes**

than:

> "WebAssembly for browsers."

---

# 16. Common Mistake: Assuming All Go Packages Work

WASM is a different target.

A package may depend on:

- OS-specific syscalls
    
- unsupported filesystem behavior
    
- networking assumptions
    
- cgo
    
- assembly
    
- platform-specific runtime functionality
    

and therefore fail to build or behave differently.

Before adopting a dependency, ask:

```text
Does it support GOOS=js?
Does it support GOOS=wasip1?
Does it require cgo?
Does it require OS syscalls?
Does it depend on networking?
Does it assume a filesystem?
```

This is especially important for large Go libraries.

---

# 17. `cgo` Considerations

WebAssembly targets have significant restrictions around cgo.

A package that works perfectly on:

```bash
GOOS=linux GOARCH=amd64
```

doesn't necessarily work for:

```bash
GOOS=wasip1 GOARCH=wasm
```

or:

```bash
GOOS=js GOARCH=wasm
```

This is one reason pure-Go dependencies are often easier to port to WASM.

---

# 18. Networking Is Another Major Difference

Do not assume:

```go
http.Get(...)
```

means exactly the same thing across:

```text
linux/amd64
js/wasm
wasip1/wasm
```

The underlying host capabilities and implementation differ.

For browser WASM, networking is constrained by browser security policies such as:

```text
CORS
same-origin policy
browser networking APIs
```

For WASI, networking support historically differs substantially depending on the WASI version and runtime.

This is an important reason not to design your application around assumptions such as:

> "WASM is basically Linux in a smaller binary."

It isn't.

---

# 19. Performance Mental Model

WASM can provide excellent performance, but:

```text
WASM ≠ automatically faster
```

The performance depends on:

- workload
    
- host/runtime
    
- boundary crossings
    
- memory allocation
    
- GC
    
- JavaScript ↔ WASM calls
    
- WASM ↔ host calls
    
- serialization
    
- data copying
    

For example:

```text
Go WASM
   │
   │ large computation
   ▼
fast

Go WASM
   │
   │ thousands of tiny JS calls
   ▼
potentially expensive
```

Crossing the host boundary repeatedly can become a bottleneck.

A good design often batches operations:

```text
Bad:

Go → JS
Go → JS
Go → JS
Go → JS
Go → JS

Better:

Go ───────────────→ JS
       batch
```

---

# 20. Testing Strategy

Don't test only:

```bash
go test ./...
```

and assume WASM compatibility.

You should have target-specific CI.

For example:

```bash
GOOS=js GOARCH=wasm go test ./...
```

and/or:

```bash
GOOS=wasip1 GOARCH=wasm go test ./...
```

depending on your application.

Then perform **actual runtime tests**, because successful compilation doesn't prove host integration works.

Think:

```text
Compile test
      ↓
Runtime test
      ↓
Host integration test
      ↓
End-to-end test
```

---

# 21. Production Decision Framework

Use this decision tree:

```text
                 Need WebAssembly?
                        │
                        ▼
              Where will it execute?
                        │
             ┌──────────┴──────────┐
             │                     │
        Browser / JS          WASI runtime
             │                     │
             ▼                     ▼
        GOOS=js              GOOS=wasip1
        GOARCH=wasm          GOARCH=wasm
```

Then ask:

### Browser?

Need:

- DOM
    
- JavaScript
    
- browser APIs
    
- Web APIs
    

→ `js/wasm`

### Server-side WASM?

Need:

- sandboxing
    
- filesystem capabilities
    
- stdin/stdout
    
- runtime-hosted execution
    
- portable WASM modules
    

→ `wasip1/wasm`

---

# 22. Practical Commands

Check supported targets:

```bash
go tool dist list
```

You should see WASM-related targets such as:

```text
js/wasm
wasip1/wasm
```

Build JavaScript-targeted WASM:

```bash
GOOS=js GOARCH=wasm go build -o main.wasm .
```

Build WASI-targeted WASM:

```bash
GOOS=wasip1 GOARCH=wasm go build -o main.wasm .
```

The important part is:

```text
          GOOS              GOARCH

          js
           │
           └────────────── wasm

       OR

        wasip1
           │
           └────────────── wasm
```

---

# 23. Architectural Comparison

|Concern|`js/wasm`|`wasip1/wasm`|
|---|---|---|
|Primary host|JavaScript|WASI|
|Browser|⭐⭐⭐|⭐|
|Node.js|⭐⭐⭐|—|
|Server-side WASM|⭐|⭐⭐⭐|
|JS interop|⭐⭐⭐|—|
|Sandbox plugins|⭐|⭐⭐⭐|
|DOM|⭐⭐⭐|—|
|WASI filesystem|—|⭐⭐⭐|
|Portable CLI-like modules|⭐|⭐⭐⭐|
|Browser UI logic|⭐⭐⭐|—|
|Host-controlled capabilities|Browser APIs|WASI capabilities|

---

# 24. The Principal Engineer Mental Model

Don't memorize:

```text
js = browser
wasip1 = server
```

That's too simplistic.

Memorize this instead:

> **`GOARCH=wasm` chooses WebAssembly as the execution architecture; `GOOS` chooses the host contract that Go expects.**

Then:

```text
GOOS=js
    ↓
JavaScript-hosted WASM
    ↓
JS APIs / Browser APIs

GOOS=wasip1
    ↓
WASI-hosted WASM
    ↓
WASI capabilities
```

And the deeper architectural principle is:

> **A WASM module is only useful in relation to its host. The host provides capabilities; the module consumes those capabilities.**

That mental model becomes particularly important when designing **WASM plugins, sandboxed execution platforms, edge runtimes, browser applications, and polyglot systems**.

### Key takeaways

1. `GOARCH=wasm` → WebAssembly target.
    
2. `GOOS=js` → JavaScript host integration.
    
3. `GOOS=wasip1` → WASI Preview 1 host integration.
    
4. `syscall/js` is specific to the JavaScript execution model.
    
5. WASI is capability-oriented; access depends on what the runtime grants.
    
6. WASM compatibility is not guaranteed merely because the code is pure Go.
    
7. Host-boundary crossings can dominate performance.
    
8. Choose the target based on the **execution environment and capability model**, not simply because "I need WASM."
    

For current Go versions, the authoritative references are the Go documentation for **WebAssembly/WASI**, the `GOOS/GOARCH` target list, and the WASI specification.

## 🔗 References
- ⬆️ Parent: [[Target OS & Architecture]]
- 📚 Module: `Go Environment & Commands`
