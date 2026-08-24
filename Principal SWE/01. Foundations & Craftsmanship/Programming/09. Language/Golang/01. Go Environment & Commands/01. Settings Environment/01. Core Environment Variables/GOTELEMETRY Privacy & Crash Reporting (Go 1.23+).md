---
title: "GOTELEMETRY Privacy & Crash Reporting (Go 1.23+)"
tags:
  - golang
  - environment
  - telemetry
  - tooling
  - principal-swe
parent: "[[Core Environment Variables]]"
---

# `GOTELEMETRY` — Privacy & Crash Reporting (Go 1.23+)

Starting with **Go 1.23**, the Go toolchain introduced **Go Telemetry**, a system for collecting usage, performance, and breakage statistics from Go's own developer tools. It is **opt-in for uploading** and is controlled by the `go telemetry` command. ([Go.dev](https://go.dev/doc/go1.23?utm_source=chatgpt.com "Go 1.23 Release Notes - The Go Programming Language"))

The key mental model is:

> **`GOTELEMETRY` does not mean "send telemetry." It reports the current telemetry mode.**

---

## 1. What is `GOTELEMETRY`?

`GOTELEMETRY` is a **read-only Go environment variable** that reports the current Go telemetry mode. ([Go.dev](https://go.dev/doc/telemetry?utm_source=chatgpt.com "Go Telemetry - The Go Programming Language"))

Check it with:

```bash
go env GOTELEMETRY
```

Possible values are:

```text
local
on
off
```

They mean:

| Mode    | Local collection |                     Upload |
| ------- | ---------------: | -------------------------: |
| `local` |              Yes |                         No |
| `on`    |              Yes | Yes, approved/sampled data |
| `off`   |               No |                         No |

The default mode is:

```text
local
```

So Go can collect telemetry locally by default, but **does not upload it** unless you explicitly opt in. ([Go.dev](https://go.dev/doc/telemetry?utm_source=chatgpt.com "Go Telemetry - The Go Programming Language"))

---

# 2. `GOTELEMETRY` is not a normal environment variable

This distinction is important.

You generally should **not** do:

```bash
export GOTELEMETRY=off
```

Instead, use:

```bash
go telemetry off
```

or:

```bash
go telemetry local
```

or:

```bash
go telemetry on
```

Then inspect the resulting state:

```bash
go env GOTELEMETRY
```

The Go documentation describes `GOTELEMETRY` as a **read-only environment variable** exposing the configured mode. ([Go.dev](https://go.dev/doc/telemetry?utm_source=chatgpt.com "Go Telemetry - The Go Programming Language"))

Think of it like:

```text
go telemetry off
        │
        ▼
 persistent configuration
        │
        ▼
go env GOTELEMETRY
        │
        ▼
      "off"
```

---

# 3. What does Go Telemetry actually collect?

Telemetry is primarily for **Go toolchain programs**, not your applications.

Examples include:

```text
go
compiler
linker
go vet
gopls
govulncheck
```

The official documentation describes telemetry as collecting data about the **performance and usage of Go toolchain programs**. ([Go.dev](https://go.dev/doc/telemetry?utm_source=chatgpt.com "Go Telemetry - The Go Programming Language"))

Examples of information can include:

- Go version
    
- operating system
    
- CPU architecture
    
- tool version
    
- invocation counters
    
- selected feature usage
    
- performance counters
    
- approved stack counters
    

For example, current public telemetry charts include things such as:

```text
cmd/compile
    GOOS
    GOARCH
    GoVersion
    compile/invocations

cmd/go
    GOOS
    GOARCH
    GoVersion
    go/invocations
```

([telemetry.go.dev](https://telemetry.go.dev/?utm_source=chatgpt.com "Go Telemetry"))

---

# 4. What it does NOT collect

A very important privacy distinction:

Go telemetry is **not a mechanism for uploading your application's source code**.

The Go documentation specifically states that stack counters contain function names and line numbers from **Go toolchain programs**, and do not include user inputs such as the names or contents of your source code. ([Go.dev](https://go.dev/doc/telemetry?utm_source=chatgpt.com "Go Telemetry - The Go Programming Language"))

The privacy policy also states that telemetry files do not contain personal or other identifying information about the user or their system. ([telemetry.go.dev](https://telemetry.go.dev/privacy?utm_source=chatgpt.com "Go Telemetry Privacy Policy"))

So this:

```go
package main

func main() {
    // Your application
}
```

is not being uploaded simply because Go telemetry is enabled.

Also important:

> **Go Telemetry is not built into your compiled application binaries.**

The telemetry service explicitly states that it is for Go toolchain programs and is not built into users' binaries. ([telemetry.go.dev](https://telemetry.go.dev/?utm_source=chatgpt.com "Go Telemetry"))

---

# 5. Why does Go need telemetry?

The core problem is that traditional bug reports are incomplete.

Imagine:

```text
100,000 developers
       │
       ├── 99,900 never report problems
       │
       └── 100 report bugs
```

The Go team sees only a tiny fraction of real-world behavior.

Telemetry can expose things such as:

```text
rare execution path
       ↓
unexpected condition
       ↓
stack counter
       ↓
aggregate telemetry
       ↓
Go team discovers bug
```

This is particularly valuable for bugs that:

- happen rarely
    
- are difficult to reproduce
    
- don't cause an obvious failure
    
- users don't bother reporting
    
- only occur on particular platforms
    

The Go team describes `gopls` telemetry finding unexpected execution paths that led to real bugs that were difficult to discover otherwise. ([Go.dev](https://go.dev/blog/gotelemetry?utm_source=chatgpt.com "Telemetry in Go 1.23 and beyond - The Go Programming Language"))

---

# 6. Crash reporting

This is where the title **"Privacy & Crash Reporting"** needs an important clarification.

Go 1.23 introduced:

```go
runtime.SetCrashOutput(...)
```

which allows programs to arrange automated crash reporting through a watchdog process. ([Go.dev](https://go.dev/blog/gotelemetry?utm_source=chatgpt.com "Telemetry in Go 1.23 and beyond - The Go Programming Language"))

For example, the architecture can conceptually look like:

```text
                 gopls
                   │
              unexpected crash
                   │
                   ▼
          SetCrashOutput(...)
                   │
                   ▼
              watchdog
                   │
                   ▼
            crash information
                   │
                   ▼
             telemetry system
```

This is primarily used by tools such as `gopls`.

It does **not** mean:

```text
Go compiler automatically uploads
every crash from every Go application
```

That is not how Go Telemetry works.

---

# 7. `gopls` and crash telemetry

`gopls` provides a useful real-world example.

With Go 1.23, `gopls` can use the crash-output mechanism to report crash information through telemetry. The Go team's documentation describes a `crash/crash` stack counter that can be generated when `gopls` crashes, assuming it was built with Go 1.23+. ([Go.dev](https://go.dev/blog/gotelemetry?utm_source=chatgpt.com "Telemetry in Go 1.23 and beyond - The Go Programming Language"))

This gives the Go team something like:

```text
Crash
  ↓
stack information
  ↓
anonymous/aggregated telemetry
  ↓
identify recurring failure
  ↓
reproduce
  ↓
fix
```

This is much closer to **automated diagnostic reporting** than conventional application analytics.

---

# 8. Three telemetry states

### `local`

```bash
go telemetry local
```

Means:

```text
Collect locally
       │
       ▼
Do NOT upload
```

This is the default mode. ([Go.dev](https://go.dev/doc/telemetry?utm_source=chatgpt.com "Go Telemetry - The Go Programming Language"))

---

### `on`

```bash
go telemetry on
```

Means:

```text
Collect locally
       │
       ▼
Upload approved telemetry
       │
       ▼
approximately weekly
```

Only approved counters are eligible for upload, and uploads can also be sampled. ([Go.dev](https://go.dev/doc/telemetry?utm_source=chatgpt.com "Go Telemetry - The Go Programming Language"))

---

### `off`

```bash
go telemetry off
```

Means:

```text
No collection
       │
       X
No upload
```

This is the strongest privacy setting provided by Go Telemetry. ([Go.dev](https://go.dev/doc/telemetry?utm_source=chatgpt.com "Go Telemetry - The Go Programming Language"))

---

# 9. Where is telemetry stored?

Go stores local telemetry data under:

```text
os.UserConfigDir()/go/telemetry
```

The documentation refers to this as:

```text
<gotelemetry>
```

with local data under:

```text
<gotelemetry>/local
```

and uploaded-report copies under the corresponding upload/remote area depending on the tooling/version. ([Go.dev](https://go.dev/doc/telemetry?utm_source=chatgpt.com "Go Telemetry - The Go Programming Language"))

You can inspect the location with:

```bash
go env GOTELEMETRYDIR
```

So:

```text
GOTELEMETRY
    └── telemetry mode

GOTELEMETRYDIR
    └── telemetry data/config location
```

---

# 10. Inspect your current configuration

Start with:

```bash
go telemetry
```

Then:

```bash
go env GOTELEMETRY
```

And:

```bash
go env GOTELEMETRYDIR
```

A useful diagnostic:

```bash
go env GOTELEMETRY GOTELEMETRYDIR
```

You might see:

```text
local
/home/mrb/.config/go/telemetry
```

The exact directory depends on your OS and user configuration directory.

---

# 11. `GOTELEMETRY` vs `GOTELEMETRYDIR`

Do not confuse these.

```text
GOTELEMETRY
     │
     └── "What mode?"
         local / on / off


GOTELEMETRYDIR
     │
     └── "Where is telemetry data/config?"
```

So:

```bash
go env GOTELEMETRY
```

answers:

> Is telemetry local, enabled for upload, or disabled?

while:

```bash
go env GOTELEMETRYDIR
```

answers:

> Where does Go keep its telemetry state?

---

# 12. CI/CD recommendation

For CI, explicitly decide what your organization's privacy policy requires.

For example, if you want deterministic behavior and don't want developer-machine telemetry settings influencing CI:

```bash
go telemetry off
```

However, be careful about **global configuration changes inside shared build environments**.

A better operational approach is to define the policy at the image/runner level and verify it:

```bash
go env GOTELEMETRY
```

For highly controlled environments, you may want:

```text
CI runner
   │
   ├── no unexpected outbound telemetry
   ├── deterministic tool configuration
   └── explicit network policy
```

From a security perspective, this is less about Go being "unsafe" and more about having an explicit **egress/data-governance policy**.

---

# 13. Important distinction: telemetry vs your application's observability

Do not confuse:

```text
Go Telemetry
```

with:

```text
Application Observability
```

They solve completely different problems.

### Go Telemetry

```text
Go toolchain
    ↓
usage / performance / breakage
    ↓
Go maintainers
```

### Your application observability

```text
Your service
    ↓
logs / metrics / traces
    ↓
your observability platform
    ↓
your SRE / engineering team
```

For example, your Go API should still use proper:

- Prometheus/OpenTelemetry metrics
    
- distributed tracing
    
- structured logging
    
- RED metrics
    
- P95/P99 latency
    
- error rates
    
- resource saturation
    
- application-specific crash reporting
    

Turning Go telemetry on does **not** give you application observability.

---

# 14. Security & privacy mental model

A good way to reason about it:

```text
                    Go Telemetry
                         │
              ┌──────────┴──────────┐
              │                     │
           Local                 Upload
              │                     │
          default              explicit opt-in
              │                     │
              ▼                     ▼
       your machine          approved counters
                                    │
                                    ▼
                             telemetry.go.dev
```

The important privacy guarantees are:

1. **Uploading requires opt-in.**
    
2. Default mode is `local`.
    
3. `off` disables collection entirely.
    
4. Only approved counters are eligible for upload.
    
5. Uploading occurs approximately weekly.
    
6. Uploaded telemetry is publicly available in aggregate/raw datasets.
    
7. Go telemetry is for Go toolchain programs, not automatically for your binaries. ([Go.dev](https://go.dev/doc/telemetry?utm_source=chatgpt.com "Go Telemetry - The Go Programming Language"))
    

---

# Principal Engineer takeaway

The most important mental model is:

```text
GOTELEMETRY
    │
    ├── local → collect, don't upload
    ├── on    → collect + approved upload
    └── off   → don't collect
```

And:

```text
GOTELEMETRY ≠ crash reporting switch
GOTELEMETRY ≠ application telemetry
GOTELEMETRY ≠ environment variable you normally export
```

It is the **read-only representation of Go's telemetry mode**; the mode itself is controlled with `go telemetry`. ([Go.dev](https://go.dev/doc/telemetry?utm_source=chatgpt.com "Go Telemetry - The Go Programming Language"))

For a production engineer, the deeper lesson is broader:

> **Telemetry should be evaluated as a data-flow and trust-boundary problem: What is collected? Where is it stored? Who can receive it? What requires consent? What is the retention model? Can egress be controlled?**

That mental model applies equally to Go tooling, application telemetry, crash reporting, OpenTelemetry, CI systems, and SaaS observability platforms.

## 🔗 References
- ⬆️ Parent: [[Core Environment Variables]]
- 📚 Module: `Go Environment & Commands`
