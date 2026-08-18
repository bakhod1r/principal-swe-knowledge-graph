---
title: GOPRIVATE
tags:
  - golang
  - basics
  - environment
  - modules
parent: "[[Modules & Dependencies]]"
---

# `GOPRIVATE`

> 🔗 **Architecture & Dependencies**: `Private Modules in Dependencies`

**`GOPRIVATE`** controls which modules are considered private and should not be downloaded via the public `GOPROXY` or verified against the public `GOSUMDB`.

---

## ⚙️ Configuration

```bash
# Mark all repositories under an organization as private
go env -w GOPRIVATE=github.com/mycompany/*

# Multiple comma-separated glob patterns
go env -w GOPRIVATE=github.com/mycompany/*,gitlab.internal.net/*
```

---

## 🔗 Related References
- ⬆️ Master Index: `Settings Environment`
- 🔒 Private Modules Architecture: `Private Modules`
- 🌐 Bypass Proxy: `GONOPROXY`
- 🛡️ Bypass Checksum: `GONOSUMDB`

---

## 🔗 References
- ⬆️ Parent: [[Modules & Dependencies]]
