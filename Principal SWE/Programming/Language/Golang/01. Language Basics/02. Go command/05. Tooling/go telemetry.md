---
title: Telemetry Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - telemetry
  - privacy
parent: "[[Tooling]]"
---

# `go telemetry`

Controls the Go toolchain's opt-in telemetry. See `GOTELEMETRY` and
`GOTELEMETRYDIR`.

## 1. Usage

```bash
go telemetry            # show current mode
go telemetry local      # collect locally, upload nothing  (default)
go telemetry on         # collect and upload to telemetry.go.dev
go telemetry off        # collect nothing
```

## 2. Modes

| Mode | Collects | Uploads |
|---|---|---|
| `off` | no | no |
| `local` | yes, to disk | no |
| `on` | yes | yes, weekly |

**`local` is the default** — nothing leaves the machine unless you opt in with `on`.

## 3. Where Data Lives

```bash
go env GOTELEMETRYDIR    # counter files and, if enabled, upload records
```

Counters are aggregate (command names, error categories, version) — no file
paths, module names, or source content.

## 4. Gotchas

- Setting `GOTELEMETRY=off` as an environment variable also works and is easier
  to enforce fleet-wide.
- `gopls` reports through the same mechanism and honours the same mode.

---

## 🔗 References
- ⬆️ Parent: [[Tooling]]
