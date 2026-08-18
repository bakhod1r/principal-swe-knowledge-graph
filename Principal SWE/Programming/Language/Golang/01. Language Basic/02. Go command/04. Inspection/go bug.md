---
title: Bug Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - support
parent: "[[Inspection]]"
---

# `go bug`

Opens the Go issue tracker in a browser with a pre-filled report containing your
environment.

## 1. Usage

```bash
go bug
```

## 2. What It Collects

The body is pre-populated with the full `go env` dump plus `go version`, and on
Linux the `uname -sr` and libc version. That is exactly the information every Go
bug report is otherwise asked for in the first reply.

```text
### What version of Go are you using?
go version go1.26.2 darwin/arm64

### Does this issue reproduce with the latest release?

### What operating system and processor architecture?
<go env output>
```

## 3. Before Filing

```bash
go version              # is it the latest release?
gotip                   # does it reproduce on tip?
```

Most reports are closed as "fixed in a newer release". Checking first is faster
than filing.

## 4. Gotchas

- It only opens a browser — nothing is submitted automatically, and nothing is
  sent anywhere if you close the tab.
- The dump includes local paths (`GOPATH`, `GOMODCACHE`, `GOROOT`) and
  `GOPRIVATE` patterns. **Review before posting publicly** — internal hostnames
  leak this way.
- Unrelated to `go telemetry`, which is a separate opt-in mechanism.

---

## 🔗 References
- ⬆️ Parent: [[Inspection]]
