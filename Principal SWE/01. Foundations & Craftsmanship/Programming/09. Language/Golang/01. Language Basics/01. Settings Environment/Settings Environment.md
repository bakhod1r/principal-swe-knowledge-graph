---
title: Settings Environment
tags:
  - golang
  - environment
  - principal-swe
parent: "[[Language Basics]]"
---


# Setting Up the Go Environment

When setting up a Go development environment, the key is to distinguish between **Go itself**, **Go's environment variables**, and your **shell environment**.

## 1. Install Go

After installing Go, verify:

```bash
go version
```

Example:

```text
go version go1.25.0 linux/amd64
```

Then inspect Go's configuration:

```bash
go env
```

For individual values:

```bash
go env GOROOT
go env GOPATH
go env GOBIN
go env GOPROXY
```

---

## 2. Understand `GOROOT`

`GOROOT` points to the Go installation:

```bash
go env GOROOT
```

Example:

```text
/usr/local/go
```

You generally **should not manually set `GOROOT`** when Go is installed normally.

The Go toolchain knows where its installation is.

So avoid adding this to your shell profile unless you have a specific reason:

```bash
export GOROOT=/usr/local/go
```

Modern Go installations usually manage this automatically.

---

## 3. Understand `GOPATH`

`GOPATH` is Go's workspace/cache location.

Check it:

```bash
go env GOPATH
```

Typical Linux value:

```text
/home/user/go
```

Inside it you may encounter:

```text
~/go/
├── bin/
├── pkg/
└── ...
```

The important directory for developers is:

```text
~/go/bin
```

because tools installed with:

```bash
go install example.com/tool@latest
```

may be placed there.

---

# 4. Add Go binaries to `PATH`

This is one of the most useful shell configuration steps.

First check:

```bash
go env GOPATH
```

Then add its `bin` directory to your `PATH`.

For Bash:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

For Zsh:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Place this in the appropriate shell startup file, such as:

```text
~/.bashrc
```

or:

```text
~/.zshrc
```

Then reload:

```bash
source ~/.bashrc
```

Verify:

```bash
echo "$PATH"
```

and:

```bash
command -v your-tool
```

---

# 5. Why `GOPATH/bin` matters

Suppose you install a CLI tool:

```bash
go install golang.org/x/tools/gopls@latest
```

The binary may be installed into:

```text
~/go/bin/gopls
```

If `~/go/bin` is not in `PATH`, this fails:

```bash
gopls
```

even though the binary exists.

You can verify:

```bash
ls "$(go env GOPATH)/bin"
```

Then:

```bash
command -v gopls
```

should resolve to something like:

```text
/home/user/go/bin/gopls
```

---

# 6. `GOBIN`

`GOBIN` controls where `go install` places executables.

Check:

```bash
go env GOBIN
```

If empty, Go uses the default:

```text
$GOPATH/bin
```

You can explicitly configure it:

```bash
go env -w GOBIN="$HOME/go/bin"
```

Then:

```bash
go env GOBIN
```

returns:

```text
/home/user/go/bin
```

In most cases, **you don't need to configure `GOBIN`**. The default behavior is sufficient.

---

# 7. `GOMODCACHE`

Go also maintains a module download cache.

Check:

```bash
go env GOMODCACHE
```

Usually something like:

```text
~/go/pkg/mod
```

This is where downloaded module source is cached.

Conceptually:

```text
GOPATH
├── bin       ← installed Go tools
├── pkg
│   └── mod   ← module cache
└── ...
```

Do not confuse this with your application's source code.

Modern Go projects normally live anywhere convenient:

```text
~/projects/my-service
```

They **do not need to live under `GOPATH/src`**.

---

# 8. Go Modules

Modern Go development should generally use modules.

Create a project:

```bash
mkdir my-service
cd my-service

go mod init example.com/my-service
```

This creates:

```text
my-service/
├── go.mod
└── ...
```

For example:

```go
module example.com/my-service

go 1.25
```

Then dependencies are managed through:

```bash
go get
go mod tidy
go mod download
```

The important mental model is:

```text
Your project
    │
    ├── go.mod
    ├── go.sum
    │
    └── source code
             │
             ▼
        Go modules
             │
             ▼
       module cache
```

---

# 9. `GOPROXY`

Go uses `GOPROXY` to determine where modules are downloaded from.

Check:

```bash
go env GOPROXY
```

A common default is:

```text
https://proxy.golang.org,direct
```

This means roughly:

```text
Try Go module proxy
       │
       ├── success → use it
       │
       └── unavailable → direct source lookup
```

For most developers, **leave the default unless your organization requires an internal proxy**.

---

# 10. `GOPRIVATE`

This becomes important when working with private repositories.

For example:

```bash
go env -w GOPRIVATE=github.com/mycompany/*
```

This tells Go that matching modules are private.

You might have:

```text
github.com/mycompany/payment-service
github.com/mycompany/internal-lib
```

and configure:

```bash
go env -w GOPRIVATE=github.com/mycompany/*
```

This is especially important for avoiding attempts to retrieve private modules through public module infrastructure.

---

# 11. Recommended developer setup

For a typical Linux Go development machine, keep the setup simple:

```bash
go version
go env GOPATH
go env GOROOT
go env GOPROXY
```

Then make sure:

```bash
$(go env GOPATH)/bin
```

is in your `PATH`.

For example:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Then:

```bash
go install golang.org/x/tools/gopls@latest
command -v gopls
```

---

# 12. What should actually be persistent?

A clean setup generally looks like this:

```text
Shell profile
│
└── PATH
    └── $GOPATH/bin
```

While Go-specific configuration can be managed through:

```bash
go env -w NAME=value
```

For example:

```bash
go env -w GOPRIVATE=github.com/mycompany/*
```

This is preferable to putting every Go setting into `.bashrc`.

### Good separation

```text
.bashrc / .zshrc
    ↓
Shell configuration

go env -w
    ↓
Go configuration

go.mod
    ↓
Project/module configuration

Kubernetes / systemd / CI
    ↓
Production runtime configuration
```

That separation prevents a common engineering mistake: **mixing developer-machine configuration with application configuration**.

---

## 13. Debugging checklist

When a Go command is not found:

```bash
which go
go version
go env GOROOT
go env GOPATH
go env GOBIN
echo "$PATH"
```

When an installed Go tool is not found:

```bash
ls "$(go env GOPATH)/bin"
echo "$PATH"
command -v tool-name
```

When a dependency cannot be downloaded:

```bash
go env GOPROXY
go env GOPRIVATE
go env GOPATH
go env GOMODCACHE
```

And for a complete snapshot:

```bash
go env
```

### The key mental model

```text
Go installation
    ↓
GOROOT
    ↓
Go toolchain

Developer workspace
    ↓
GOPATH
    ├── bin
    └── pkg/mod

Project
    ↓
go.mod
    ↓
Go Modules

Shell
    ↓
PATH
    ↓
Can locate Go executables
```

**Do not configure everything just because the variables exist.** A strong Go setup uses Go's defaults wherever possible and changes only the settings required by the environment or organization.

---

## 🔗 References
- ⬆️ Parent: [[Language Basics]]

