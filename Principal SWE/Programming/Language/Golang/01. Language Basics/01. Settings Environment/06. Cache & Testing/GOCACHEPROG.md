---
title: GOCACHEPROG
tags:
  - golang
  - basics
  - environment
  - cache
  - build
parent: "[[Cache & Testing]]"
---

# `GOCACHEPROG`

**Go 1.24+** (an experiment since 1.21). Names an external program that replaces
the on-disk build cache in `GOCACHE`.

## 1. Usage

```bash
go env -w GOCACHEPROG='/usr/local/bin/gocacheprog --remote=cache.corp:8080'
```

The `go` command starts the program once and speaks a **JSON-over-stdin/stdout**
protocol to it for every cache lookup and store.

## 2. Protocol Sketch

```text
go command                       cache program
    │  {"Command":"get","ActionID":"..."}   │
    ├──────────────────────────────────────►│
    │  {"Miss":true}                        │
    │◄──────────────────────────────────────┤
    │  {"Command":"put","ObjectID":"...","BodySize":4096}
    ├──────────────────────────────────────►│
```

Commands: `get`, `put`, `close`. The program answers with a disk path for hits,
so the `go` command still reads objects as files.

## 3. Why

Shared/remote build caches across CI agents and developer machines — the thing
Bazel and `sccache` provide, without leaving the `go` command.

## 4. Gotchas

- The protocol is **not covered by the Go 1 compatibility promise**; it can change
  between releases.
- A slow or flaky cache program makes every build slower than no cache at all.
- `go clean` `-cache` cannot clear a remote cache — that is the program's job.
- If unset, the normal `GOCACHE` directory is used.

---

## 🔗 References
- ⬆️ Parent: [[Cache & Testing]]
