---
title: GOBIN
tags:
  - golang
  - basics
  - environment
parent: "[[Core]]"
---
# `GOBIN`

**`GOBIN` = the directory where `go install` puts executable binaries.**

Check it:

```bash
go env GOBIN
```

If it returns empty:

```text
""
```

Go normally uses:

```text
$GOPATH/bin
```

---

## Mental Model

```text
go install
    │
    ▼
  GOBIN
    │
    ▼
~/go/bin/
    │
    ├── gopls
    ├── air
    └── other tools
```

Example:

```bash
go install golang.org/x/tools/gopls@latest
```

Usually:

```text
~/go/bin/gopls
```

---

## `GOBIN` vs `GOPATH/bin`

### Default

```text
GOBIN = ""
```

means Go uses:

```text
GOPATH/bin
```

For example:

```text
GOPATH=/home/user/go

        ↓

GOPATH/bin=/home/user/go/bin
```

### Custom GOBIN

You can explicitly configure:

```bash
go env -w GOBIN=$HOME/.local/bin
```

Now:

```bash
go install ...
```

puts binaries into:

```text
$HOME/.local/bin
```

instead of:

```text
$HOME/go/bin
```

---

## Important: `GOBIN` ≠ `PATH`

This distinction is critical.

```text
GOBIN
  ↓
Where should Go PUT the binary?

PATH
  ↓
Where should the SHELL LOOK for the binary?
```

Example:

```text
GOBIN
  ↓
~/go/bin
  ↓
gopls
```

But to run:

```bash
gopls
```

your `PATH` must contain:

```text
~/go/bin
```

So:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

## `GOBIN` vs `GOROOT` vs `GOPATH`

|Variable|Purpose|
|---|---|
|`GOROOT`|Where Go itself is installed|
|`GOPATH`|User Go environment/cache|
|`GOBIN`|Where installed Go executables go|
|`PATH`|Where shell searches for executables|

Think:

```text
/usr/local/go
     │
     └── GOROOT
          └── Go itself

~/go
  │
  └── GOPATH
       ├── bin/
       └── pkg/mod/

~/go/bin
     │
     └── GOBIN
          └── gopls

PATH
  │
  └── tells shell where to find gopls
```

### One-line definition

> **`GOBIN` controls where `go install` places executable Go tools; if unset, the default is `$GOPATH/bin`.**

---

## 🔗 References
- ⬆️ Parent: [[Core]]
