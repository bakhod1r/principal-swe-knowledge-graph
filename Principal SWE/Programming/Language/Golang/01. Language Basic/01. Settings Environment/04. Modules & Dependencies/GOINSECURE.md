---
title: GOINSECURE
tags:
  - golang
  - basics
  - environment
  - modules
  - security
parent: "[[Modules & Dependencies]]"
---

# `GOINSECURE`

Glob patterns of module prefixes that may be fetched **without HTTPS certificate
validation** — plain HTTP or a self-signed TLS certificate.

## 1. Usage

```bash
go env -w GOINSECURE='git.internal.corp/*,10.0.*'
```

Comma-separated, same glob syntax as `GOPRIVATE` (`path.Match` per element).

## 2. What It Does and Does Not Do

| | Effect |
|---|---|
| TLS certificate check | **Disabled** for matching prefixes |
| Proxy usage | Unchanged — still goes through `GOPROXY` unless `GONOPROXY` matches |
| Checksum verification | **Unchanged** — `GOSUMDB` and `go.sum` still apply |

That last row matters: `GOINSECURE` weakens transport security only. Integrity is
still enforced, so a MITM cannot silently swap module content that is already in
`go.sum`.

## 3. When It Is Legitimate

An internal Git server with an internal CA that is not installed on build agents.
The correct fix is installing the CA; `GOINSECURE` is the pragmatic fallback.

## 4. Gotchas

- **Security:** on a first-ever fetch there is no `go.sum` entry yet, so a MITM
  on an insecure prefix can inject the initial content. Pair with `GOPRIVATE`
  only inside a trusted network.
- Does not affect `git clone` over `ssh://` — that has its own trust model.
- Prefer `GOAUTH` plus a proper CA over `GOINSECURE` in anything long-lived.

---

## 🔗 References
- ⬆️ Parent: [[Modules & Dependencies]]
