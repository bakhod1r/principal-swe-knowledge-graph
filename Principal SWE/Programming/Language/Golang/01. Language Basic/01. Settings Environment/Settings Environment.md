---
title: Settings Environment — Setup & Variable Architecture
tags:
  - golang
  - basics
  - environment
parent: "[[Language Basic]]"
---

# 🌐 Settings Environment

Two halves: **Part A** — first-time setup and the four variables you actually need. **Part B** — the full environment-variable architecture, split into 8 category hubs.

---
---

# Part A — Golang First-Time Setup

## 1. Install Go

Download and install Go from the official site:

[Go Installation Guide](https://go.dev/doc/install)

Verify:

```bash
go version
```

Example:

```text
go version go1.25.x linux/amd64
```

---

## 2. Check Go Environment

Run:

```bash
go env
```

The most important variables:

```bash
go env GOROOT
go env GOPATH
go env GOBIN
go env GOOS
go env GOARCH
```

Typical Linux/macOS setup:

```text
GOROOT=/usr/local/go
GOPATH=$HOME/go
GOBIN=
GOOS=linux
GOARCH=amd64
```

---

## 3. Understand the Core Variables

### `GOROOT`

Where Go itself is installed:

```text
/usr/local/go
```

Contains:

```text
/usr/local/go/bin/go
/usr/local/go/bin/gofmt
```

Normally, **do not manually set `GOROOT`**.

### `GOPATH`

Go's user workspace/cache:

```text
$HOME/go
```

Important directories:

```text
~/go/
├── bin/       # installed Go tools
└── pkg/mod/   # downloaded modules
```

### `GOBIN`

Where `go install` puts executables.

If empty, Go normally uses:

```text
$GOPATH/bin
```

---

## 4. Configure `PATH`

Your shell needs to find both:

```text
Go itself
    ↓
/usr/local/go/bin

Installed Go tools
    ↓
~/go/bin
```

Linux/macOS:

```bash
export PATH="$PATH:/usr/local/go/bin"
export PATH="$PATH:$HOME/go/bin"
```

For persistence, put these in your shell configuration, e.g. `~/.bashrc` or `~/.zshrc`. See `Shell Startup`.

Then reload:

```bash
source ~/.bashrc
```

or open a new terminal.

---

## 5. Verify PATH

```bash
command -v go
```

Expected:

```text
/usr/local/go/bin/go
```

Check installed tools:

```bash
echo $PATH
```

You should see:

```text
/usr/local/go/bin
```

and:

```text
/home/<user>/go/bin
```

---

## 6. The Core Mental Model

```text
Operating System
       │
       ▼
     PATH
       │
       ├── /usr/local/go/bin
       │        │
       │        └── go
       │
       └── ~/go/bin
                │
                └── installed tools
```

And:

```text
go install
     │
     ▼
   GOBIN
     │
     ▼
 ~/go/bin
     │
     ▼
   PATH
     │
     ▼
  my-tool
```

### Remember these 5 concepts

|Concept|Meaning|Reference Note|
|---|---|---|
|`PATH`|Where the OS/shell searches for executables|`PATH`|
|`GOROOT`|Where Go itself is installed|`GOROOT`|
|`GOPATH`|Go's user workspace/cache|`GOPATH`|
|`GOBIN`|Where installed Go binaries go|`GOBIN`|
|`go env`|Shows Go's environment/configuration|`GOENV`|

**Golden rule:** Don't blindly configure every environment variable. Usually you only need to make sure the Go executable and your installed Go tools are reachable through `PATH`.

---
---

# Part B — Environment Variables: Full Architecture

In Go's toolchain and runtime system, all environment variables are categorized into 8 distinct functional domains:

```text
Go Environment Variables
│
├── [[OS Environment]]
│   └── Variables, PATH, Shell Startup
│
├── [[Core]]
│   ├── GOENV, GOROOT, GOPATH, GOBIN
│   └── GOMOD, GOWORK, GOMODCACHE
│
├── [[Build & Platform]]
│   ├── GOOS, GOARCH, GOAMD64, GOARM, GOARM64
│   └── GOMIPS, GOMIPS64, GOPPC64, GOS390X, GOWASM
│
├── [[Modules & Dependencies]]
│   ├── GOPROXY, GOSUMDB, GOPRIVATE
│   └── GONOPROXY, GONOSUMDB, GOVCS, GOTOOLCHAIN
│
├── [[Build & Tooling]]
│   └── GOFLAGS, GOTOOLDIR, GOVERSION, GOEXPERIMENT, GOFIPS140
│
├── [[Cache & Testing]]
│   └── GOCACHE, GOTMPDIR, GODEBUG, GOTRACEBACK, GOCOVERDIR
│
├── [[Telemetry]]
│   └── GOTELEMETRY, GOTELEMETRYDIR
│
└── [[CGO & External Toolchain]]
    ├── CGO_ENABLED, CGO_CFLAGS, CGO_CPPFLAGS, CGO_CXXFLAGS, CGO_LDFLAGS
    └── CC, CXX, AR, PKG_CONFIG
```

---

## 🗂️ Categories & Hub Notes

1. 🖥️ **[[OS Environment]]** — OS-level variables (`export`), system executable search `PATH`, and shell startup configurations.
2. 🐹 **[[Core]]** — SDK installation `GOROOT`, workspace `GOPATH`, binary installation `GOBIN`, module manifest `GOMOD`, workspace `GOWORK`, and module cache `GOMODCACHE`. (See also: `Go Source Code Structure`)
3. 🏗️ **[[Build & Platform]]** — Target OS `GOOS`, architecture `GOARCH`, and hardware microarchitecture tuning flags (`GOAMD64`, `GOARM`, `GOARM64`, etc.).
4. 📦 **[[Modules & Dependencies]]** — Proxy `GOPROXY`, checksum database `GOSUMDB`, private module rules (`GOPRIVATE`, `GONOPROXY`, `GONOSUMDB`), `GOVCS`, and `GOTOOLCHAIN`.
5. ⚙️ **[[Build & Tooling]]** — Default command flags `GOFLAGS`, internal tool binaries `GOTOOLDIR`, compiler experiments `GOEXPERIMENT`, and `GOFIPS140`.
6. 🚀 **[[Cache & Testing]]** — Build/test cache `GOCACHE`, runtime diagnostics `GODEBUG`, panic traces `GOTRACEBACK`, and coverage profiles `GOCOVERDIR`.
7. 📊 **[[Telemetry]]** — Toolchain anonymous telemetry mode `GOTELEMETRY` and storage `GOTELEMETRYDIR`.
8. 🔌 **[[CGO & External Toolchain]]** — C/C++ compilation toggle `CGO_ENABLED`, host compilers `CC`/`CXX`, and linker/preprocessor flags.

---

## 🛠️ Quick Commands

```bash
# View all Go environment variables
go env

# Output in JSON format
go env -json

# Inspect specific variables
go env GOROOT GOPATH GOOS GOARCH

# Persist a custom variable setting to go.env
go env -w GOPRIVATE=github.com/mycompany/*

# Reset variable to default
go env -u GOPRIVATE
```

---

## 🔗 References
- ⬆️ Parent: [[Language Basic]]
- 💻 Commands: `Go Commands`
- 📦 Module System: `Dependencies`
- 📂 Source Code Layout: `Go Source Code Structure`
