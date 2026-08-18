---
title: go mod download
tags:
  - golang
  - basics
  - cli
  - toolchain
  - modules
  - ci
parent: "[[Module Commands]]"
---

# `go mod download`

Fetches modules into `GOMODCACHE` without building anything.

```bash
go mod download           # everything in the build list
go mod download all
go mod download -x        # show the underlying commands
go mod download -json     # machine-readable, includes cache paths
```

## 1. The Docker Layer Pattern

```dockerfile
COPY go.mod go.sum ./
RUN go mod download          # cached layer — only busted when go.mod changes
COPY . .
RUN go build -o /app ./cmd/app
```

Dependencies are re-downloaded only when the manifests change, not on every
source edit. This is the single highest-value use of the command.

## 2. `-json` Output

```json
{
  "Path": "github.com/go-chi/chi/v5",
  "Version": "v5.1.0",
  "Dir": "/root/go/pkg/mod/github.com/go-chi/chi/v5@v5.1.0",
  "Zip": "/root/go/pkg/mod/cache/download/.../v5.1.0.zip",
  "Sum": "h1:..."
}
```

Useful for vendoring pipelines and license scanners.

## 3. Gotchas

- Since Go 1.17 the bare form downloads only modules **needed to build the main
  module**, not the entire graph. Use `all` for the old behaviour.
- It verifies against `go.sum` and `GOSUMDB` — a download failure here is
  frequently an integrity failure, not a network one. Read the message.
- Does not update `go.mod`; it is safe to run with `-mod=readonly`.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
