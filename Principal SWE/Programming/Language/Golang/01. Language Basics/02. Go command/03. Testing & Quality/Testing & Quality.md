---
title: Testing & Quality Commands
tags:
  - golang
  - basics
  - cli
  - toolchain
parent: "[[Go Commands]]"
---

# 🧪 Testing & Quality Commands

```text
_test.go files
  │
  ├── [[go test]]  → run tests, benchmarks, examples, fuzz targets
  ├── [[go vet]]   → static analysis for real mistakes
  ├── [[go fmt]]   → canonical formatting (gofmt -l -w)
  └── [[go fix]]   → rewrite deprecated API usage
```

## 🗂️ Commands

- **[[go test]]** — `-race`, `-cover`, `-bench`, `-fuzz`, `-shuffle`, result caching.
- **[[go vet]]** — printf, copylocks, lostcancel, httpresponse and friends.
- **[[go fmt]]** — formatting is not configurable; that is the feature.
- **[[go fix]]** — migration rewrites, rarely needed on modern code.

## ✅ Minimum CI Gate

```bash
test -z "$(gofmt -l .)"
go vet ./...
go test -race -count=1 ./...
```

---

## 🔗 References
- ⬆️ Parent: `Go Commands`
