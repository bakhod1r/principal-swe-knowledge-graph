---
title: GOPATH
tags:
  - golang
  - basics
  - environment
parent: "[[Core]]"
---

# `GOPATH`

**`GOPATH` = the directory Go uses for user-specific Go data, tools, and module cache.**

Check it:

```bash
go env GOPATH
```

Typical result:

```text
/home/user/go
```

## 1. Mental Model

```text
GOPATH
  │
  ├── bin/
  │    └── installed Go tools
  │
  └── pkg/
       └── mod/
            └── downloaded modules
```

Example:

```text
/home/user/go/
├── bin/
│   ├── gopls
│   └── other-tools
└── pkg/
    └── mod/
        └── cached dependencies
```

---

## 2. `GOPATH` vs `GOROOT`

This distinction is critical:

```text
GOROOT
   ↓
Go itself
   ↓
/usr/local/go
```

```text
GOPATH
   ↓
Your Go-related data
   ↓
/home/user/go
```

|Variable|Meaning|
|---|---|
|`GOROOT`|Go installation|
|`GOPATH`|User Go workspace/cache|
|`GOBIN`|Where installed binaries go|
|`PATH`|Where shell searches for executables|

---

## 3. Default GOPATH

Usually:

```text
Linux/macOS:
$HOME/go
```

Check rather than assuming:

```bash
go env GOPATH
```

---

## 4. `GOPATH/bin`

When you run:

```bash
go install some-tool@version
```

the executable normally goes to:

```text
$GOPATH/bin
```

For example:

```text
~/go/bin/gopls
```

To execute:

```bash
gopls
```

you need:

```text
~/go/bin
```

in your `PATH`.

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Relationship:

```text
go install
     ↓
GOBIN
     ↓
GOPATH/bin
     ↓
PATH
     ↓
gopls
```

---

## 5. `GOPATH/pkg/mod`

Modern Go uses:

```text
$GOPATH/pkg/mod
```

as the default module cache.

For example:

```text
~/go/pkg/mod/
├── github.com/
├── golang.org/
└── ...
```

When your project has:

```go
require github.com/example/foo v1.2.3
```

Go can download/cache that dependency in the module cache.

Check the exact location:

```bash
go env GOMODCACHE
```

---

## 6. Important: Modern Go Does NOT Require `GOPATH/src`

Old Go development used:

```text
$GOPATH/src/github.com/user/project
```

Modern Go uses **Go Modules**:

```text
~/projects/my-api/
├── go.mod
├── go.sum
└── main.go
```

So your project can live anywhere.

```text
GOPATH ≠ project location
```

This is one of the most important modern Go concepts.

---

## 7. Should You Manually Set `GOPATH`?

Usually **no**.

Check:

```bash
go env GOPATH
```

If it gives:

```text
/home/user/go
```

you can normally leave it alone.

You can configure it with:

```bash
go env -w GOPATH=$HOME/go
```

but there is generally no reason to customize the default unless your environment requires it.

---

## 8. Core Mental Model

Memorize this:

```text
GOROOT
  ↓
Go installation
  ↓
/usr/local/go

GOPATH
  ↓
User Go environment
  ├── bin/
  └── pkg/mod/

GOBIN
  ↓
Executable installation destination

PATH
  ↓
How the shell finds executables
```

### One-line definition

> **`GOPATH` is Go's user-level workspace/cache location; in modern Go, it is mainly important for `bin` and the module cache, not for storing your source code.**

---

## 🔗 References
- ⬆️ Parent: [[Core]]
