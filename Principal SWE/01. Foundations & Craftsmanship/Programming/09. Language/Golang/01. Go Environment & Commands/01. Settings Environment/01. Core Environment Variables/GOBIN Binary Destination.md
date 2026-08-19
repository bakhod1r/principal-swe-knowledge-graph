---
title: "GOBIN Binary Destination"
tags:
  - review
  - golang
  - environment
  - principal-swe
parent: "[[Core Environment Variables]]"
---

# `GOBIN` — Binary Destination

`GOBIN` answers one specific question:

> **Where should Go install executable binaries produced by `go install`?**

It does **not** determine where Go itself is installed, where modules are cached, or where your source code lives.

---

## 1. Mental model

When you run:

```bash
go install golang.org/x/tools/gopls@latest
```

the flow is:

```text
go install
    │
    ▼
Build Go package
    │
    ▼
Executable binary
    │
    ▼
GOBIN
    │
    ▼
gopls
```

So:

```text
GOBIN = binary destination
```

---

## 2. Check `GOBIN`

Run:

```bash
go env GOBIN
```

There are two important cases.

### Case 1 — `GOBIN` is explicitly configured

```text
/home/user/.local/bin
```

Then:

```bash
go install example.com/tool@latest
```

installs the executable into:

```text
/home/user/.local/bin/tool
```

### Case 2 — `GOBIN` is empty

This is normal.

```bash
go env GOBIN
```

returns nothing.

Go then uses:

```text
$GOPATH/bin
```

So if:

```bash
go env GOPATH
```

returns:

```text
/home/user/go
```

the default binary destination is:

```text
/home/user/go/bin
```

---

# 3. `GOBIN` vs `GOPATH`

These are often confused.

### `GOPATH`

Defines the broader Go workspace/cache location:

```text
GOPATH
├── bin
├── pkg
│   └── mod
└── ...
```

### `GOBIN`

Defines specifically where installed executables go:

```text
GOBIN
└── tool binaries
```

Conceptually:

```text
GOPATH=/home/user/go

          ┌─────────────────┐
          │ /home/user/go   │
          │                 │
          │ bin/ ← default  │
          │ pkg/mod/        │
          └─────────────────┘
```

If `GOBIN` is configured:

```text
GOPATH=/home/user/go
GOBIN=/home/user/.local/bin

/home/user/go
└── pkg/mod/

          +

/home/user/.local/bin
├── gopls
├── goimports
└── mytool
```

---

# 4. Setting `GOBIN`

Prefer Go's own configuration mechanism:

```bash
go env -w GOBIN="$HOME/.local/bin"
```

Verify:

```bash
go env GOBIN
```

Expected:

```text
/home/user/.local/bin
```

Now:

```bash
go install golang.org/x/tools/gopls@latest
```

will place the binary in:

```text
~/.local/bin/gopls
```

---

# 5. Don't confuse `GOBIN` with `PATH`

This is the most important practical distinction.

Suppose:

```bash
go env GOBIN
```

returns:

```text
/home/user/.local/bin
```

and Go installs:

```text
/home/user/.local/bin/gopls
```

But:

```bash
echo "$PATH"
```

does **not** contain:

```text
/home/user/.local/bin
```

Then:

```bash
gopls
```

may fail with:

```text
command not found
```

So:

```text
GOBIN
  ↓
Where Go puts the binary

PATH
  ↓
Where the shell looks for the binary
```

You need both.

For example:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

Then:

```bash
command -v gopls
```

should find it.

---

# 6. Recommended configuration

There are two good approaches.

### Option A — Use the Go default

Don't configure `GOBIN`.

```bash
go env GOBIN
```

returns empty.

Go uses:

```text
$(go env GOPATH)/bin
```

Then add that directory to `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

This is the simplest approach.

### Option B — Explicit user binary directory

Configure:

```bash
go env -w GOBIN="$HOME/.local/bin"
```

Then:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

This is also clean, especially if you want all user-installed executables in one conventional location.

**I would not configure `GOBIN` unless you have a reason to prefer a different binary directory.**

---

# 7. `go install` vs `go build`

`GOBIN` applies to `go install`, not ordinary `go build` output.

For example:

```bash
go install ./cmd/mytool
```

uses `GOBIN`:

```text
GOBIN/
└── mytool
```

But:

```bash
go build ./cmd/mytool
```

doesn't automatically install the binary into `GOBIN`.

You can explicitly choose an output:

```bash
go build -o ./bin/mytool ./cmd/mytool
```

Result:

```text
project/
└── bin/
    └── mytool
```

This distinction is important:

```text
go install
    → install executable
    → GOBIN

go build
    → produce build artifact
    → current/default/explicit output
```

---

# 8. Why `GOBIN` exists

Without `GOBIN`, Go needs a predictable place for user-installed executables.

Historically, Go's workspace model centered around:

```text
$GOPATH/bin
```

Modern Go still uses that as the default when `GOBIN` is not explicitly set.

`GOBIN` provides an override:

```text
Default:

$GOPATH/bin

        ↓ override

GOBIN
```

This lets you separate:

```text
Go module cache
    ~/go/pkg/mod

Go binaries
    ~/.local/bin
```

if that organization is useful to you.

---

# 9. Debugging `GOBIN`

If a tool was installed but you cannot find it:

### Check destination

```bash
go env GOBIN
go env GOPATH
```

### Determine expected directory

If `GOBIN` is empty:

```bash
echo "$(go env GOPATH)/bin"
```

Otherwise:

```bash
go env GOBIN
```

### Search for the binary

For example:

```bash
ls "$(go env GOPATH)/bin"
```

or, if `GOBIN` is configured:

```bash
ls "$(go env GOBIN)"
```

### Check shell resolution

```bash
command -v gopls
```

### Check `PATH`

```bash
echo "$PATH" | tr ':' '\n'
```

This gives you a deterministic debugging sequence:

```text
Where should Go put it?
        ↓
Did Go put it there?
        ↓
Is that directory in PATH?
        ↓
Can the shell resolve it?
```

---

# 10. Important security consideration

`GOBIN` can contain executable programs.

Therefore, adding a directory to `PATH` means you are trusting executables in that directory.

For example:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

means the shell may execute programs from:

```text
~/.local/bin/
```

Be careful about:

- who can write to the directory
    
- ownership and permissions
    
- installing untrusted Go tools
    
- dependency/supply-chain security
    

For developer machines, a user-owned directory is preferable to an arbitrary world-writable directory.

---

# 11. Production perspective

`GOBIN` is primarily a **developer/build-environment concern**.

You generally don't want a production service to depend on:

```text
$HOME/go/bin
```

or:

```text
$HOME/.local/bin
```

Instead, a production pipeline should produce an explicit artifact:

```text
Source
  ↓
go build
  ↓
Binary
  ↓
Container/package
  ↓
Deployment
```

For example:

```bash
go build -o bin/api ./cmd/api
```

Then package `bin/api` into your deployment artifact.

---

## Final mental model

```text
                 Go

        go install tool@version
                  │
                  ▼
            Build executable
                  │
                  ▼
             ┌─────────┐
             │  GOBIN  │
             └────┬────┘
                  │
                  ▼
            ~/.local/bin/tool
                  │
                  │ directory must be
                  │ included in
                  ▼
                PATH
                  │
                  ▼
          command: tool
```

### Remember these three

```text
GOPATH → Go workspace/cache location
GOBIN  → executable installation destination
PATH   → directories the shell searches for executables
```

**Best default:** leave `GOBIN` unset and let Go use `$GOPATH/bin`, unless you deliberately want a different binary layout.

---

## 🔗 References
- ⬆️ Parent: [[Core Environment Variables]]
- 📚 Module: `Go Environment & Commands`
