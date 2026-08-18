---
title: Private Modules Architecture
tags:
  - golang
  - basics
  - dependencies
  - modules
  - security
parent: "[[Dependencies]]"
---

# 🔒 Private Modules Architecture

Private modules represent internal corporate or proprietary repositories that must not be leaked to public proxy servers (`proxy.golang.org`) or queried against the public checksum database (`sum.golang.org`).

---

## 🗺️ Architectural Flow

```text
go get github.com/mycompany/private-pkg
                   │
                   ▼
       Matches GOPRIVATE pattern?
                   │
        ┌──────────┴──────────┐
        ▼ YES                 ▼ NO
  Direct VCS Fetch      Fetch via GOPROXY
  (Bypass SumDB)       & Verify with GOSUMDB
```

---

## 🛠️ Environment Configuration References

| Environment Variable | Role in Private Modules |
|---|---|
| **`GOPRIVATE`** | Master switch: automatically sets both `GONOPROXY` and `GONOSUMDB` for matching prefixes. |
| **`GONOPROXY`** | Granular control: bypasses proxy and fetches directly via Git/VCS. |
| **`GONOSUMDB`** | Granular control: disables public cryptographic checksum verification. |

### Quick Setup

```bash
# Configure all internal corporate repositories as private
go env -w GOPRIVATE=github.com/mycompany/*,gitlab.internal.corp/*
```

---

## 🔗 Navigation
- ⬆️ Parent: [[Dependencies]]
- 🌐 Environment Variables: `Settings Environment`
