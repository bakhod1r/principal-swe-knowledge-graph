---
title: go work sync
tags:
  - golang
  - basics
  - cli
  - toolchain
  - workspaces
parent: "[[Module Commands]]"
---

# `go work sync`

Pushes the workspace's resolved build list back into each member module's
`go.mod`.

```bash
go work sync
```

## 1. The Problem It Solves

```text
inside the workspace   → api builds against the local shared/ (any version)
outside the workspace  → api builds against shared v1.2.0 from go.mod
```

Those can drift silently: your workspace build passes, CI fails. `go work sync`
raises each member's requirements to the versions `MVS` selected across the
whole workspace, so both agree.

## 2. Typical Loop

```bash
go work use -r .
# ... edit across modules ...
go work sync
go test ./...           # inside workspace
cd api && GOWORK=off go build ./...   # verify standalone
```

That last line is the real check — see `GOWORK`.

## 3. Gotchas

- It only **raises** versions; it never adds a requirement that exists solely
  because of a `use` directive. A module importing a workspace sibling still needs
  a real `require` line before it can be published.
- It rewrites several `go.mod` files at once — commit before running.
- It does not touch `go.sum`; follow with `go mod tidy` per module.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
