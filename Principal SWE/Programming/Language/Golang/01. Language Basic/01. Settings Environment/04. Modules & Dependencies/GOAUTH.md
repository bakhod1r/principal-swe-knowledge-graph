---
title: GOAUTH
tags:
  - golang
  - basics
  - environment
  - modules
  - security
  - auth
parent: "[[Modules & Dependencies]]"
---

# `GOAUTH`

**Go 1.24+.** Declares how the `go` command authenticates to private module
servers over HTTPS. See `Private Modules`.

## 1. Values

Semicolon-separated list, tried in order:

| Value | Behaviour |
|---|---|
| `off` | Send no credentials |
| `netrc` | **Default.** Read `~/.netrc` (or `%USERPROFILE%\_netrc`) |
| `git <dir>` | Use `git credential fill` from the repo in `<dir>` |
| `<command>` | Run an external command that prints credentials |

```bash
go env -w GOAUTH='git /home/me/work; netrc'
```

## 2. Command Protocol

The command is invoked with no arguments and must print HTTP headers keyed by
URL prefix on stdout:

```text
https://git.corp.internal/

Authorization: Bearer eyJhbGci...
```

Re-invoked with a URL argument when the first attempt gets a 401 — this is how a
short-lived token can be refreshed.

## 3. Why It Replaced `.netrc`-only

`.netrc` stores a long-lived password in plaintext. `GOAUTH` allows delegating to
the OS keychain, a Git credential helper, or a Vault/OIDC token issuer.

## 4. Gotchas

- `GOAUTH` handles **HTTPS** module fetches. Direct VCS fetches (`GONOPROXY`,
  `GOPRIVATE`) still use Git's own credential machinery.
- It does not bypass `GOSUMDB` — for private modules also set `GOPRIVATE`.
- Credentials only go to hosts matched by the printed prefixes; a bare
  `Authorization` line with no URL prefix is rejected.

---

## 🔗 References
- ⬆️ Parent: [[Modules & Dependencies]]
