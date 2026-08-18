---
title: Env Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - environment
parent: "[[Inspection]]"
---

# `go env`

Prints and persists the Go toolchain configuration. See `Settings Environment`
for the full variable architecture.

## 1. Usage

```bash
go env                          # everything
go env GOROOT GOPATH GOOS GOARCH
go env -json                    # machine-readable
go env -changed                 # only values differing from the default (Go 1.23+)
go env -w GOPRIVATE='github.com/mycompany/*'
go env -u GOPRIVATE             # revert to default
```

## 2. Where `-w` Writes

```text
go env -w  →  $(go env GOENV)  →  ~/.config/go/env   (Linux)
                                  ~/Library/Application Support/go/env  (macOS)
```

See `GOENV`. Values there are lower precedence than real environment variables.

## 3. Precedence

```text
OS environment variable
        ↓ overrides
go env -w value (in the go.env file)
        ↓ overrides
toolchain default
```

## 4. Gotchas

- `go env -w` cannot set every variable — read-only ones like `GOVERSION` and
  `GOTOOLDIR` are rejected. See `GOVERSION`, `GOTOOLDIR`.
- Editing the `go env` file by hand works but skips validation; prefer `-w`.
- In CI, prefer real environment variables — they are visible in the job log.

---

## 🔗 References
- ⬆️ Parent: [[Inspection]]
