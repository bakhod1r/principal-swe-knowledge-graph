---
title: GOPROXY Protocol
tags:
  - golang
  - basics
  - dependencies
  - modules
  - proxy
  - http
parent: "[[Distribution & Integrity]]"
---

# GOPROXY Protocol

The module proxy is a plain, cacheable **HTTP GET API**. Understanding it turns
proxy problems from mysteries into `curl` commands.

## 1. Endpoints

| Path | Returns |
|---|---|
| `/<module>/@v/list` | Known versions, one per line |
| `/<module>/@v/<ver>.info` | JSON: `{"Version":"v1.2.3","Time":"..."}` |
| `/<module>/@v/<ver>.mod` | That version's `go.mod` |
| `/<module>/@v/<ver>.zip` | The module source archive |
| `/<module>/@latest` | Newest version, when `list` is empty |

```bash
curl https://proxy.golang.org/github.com/go-chi/chi/v5/@v/list
curl https://proxy.golang.org/github.com/go-chi/chi/v5/@v/v5.1.0.mod
```

## 2. Case Encoding

Module paths are case-insensitive on some filesystems, so uppercase letters are
escaped with `!`:

```text
github.com/BurntSushi/toml  →  github.com/!burnt!sushi/toml
```

This is why `GOMODCACHE` paths look mangled.

## 3. Fallback Chain

```bash
GOPROXY='https://corp.proxy,https://proxy.golang.org,direct'
```

`,` → fall through on 404/410 only. `|` → fall through on **any** error.
`direct` → clone from the VCS. `off` → no downloads at all.

## 4. Gotchas

- The proxy serves **immutable** content; a moved tag is not re-fetched. This is
  the point — see [[Distribution & Integrity]].
- A 410 Gone means "known to not exist"; a 404 may just be a cold cache.
- Private paths must be excluded with `GONOPROXY`/`GOPRIVATE` or the module
  path leaks to the public proxy in a request URL.

---

## 🔗 References
- ⬆️ Parent: `03. Distribution & Integrity`
