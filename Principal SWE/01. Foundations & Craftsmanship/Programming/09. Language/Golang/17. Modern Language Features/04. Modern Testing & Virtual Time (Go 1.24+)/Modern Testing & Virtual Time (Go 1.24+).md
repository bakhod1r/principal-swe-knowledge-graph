---
title: Modern Testing & Virtual Time (Go 1.24+)
tags:
  - golang
  - modern-go
  - principal-swe
parent: "[[Modern Language Features]]"
---

# Modern Testing & Virtual Time (Go 1.24+)

Deterministic virtual time testing (synctest package), virtual time bubbles, and ThreadSanitizer engine upgrades.

```text
Modern Testing & Virtual Time (Go 1.24+)
│
├── [[Virtual Time Bubbles (synctest package)]]
├── [[Testing Concurrent Goroutines with synctest]]
└── [[ThreadSanitizer & Race Detector Modernization]]
```

---

## 🗂️ Topics

- [[Virtual Time Bubbles (synctest package)]] — Simulating weeks of timers, sleep, and channel delays instantly inside synctest.Run().
- [[Testing Concurrent Goroutines with synctest]] — Eliminating flaky time.Sleep() timeouts in distributed and concurrent unit tests.
- [[ThreadSanitizer & Race Detector Modernization]] — Upgraded v3 ThreadSanitizer engine with lower memory overhead and faster race detection.

---

## 🔗 References
- ⬆️ Parent: [[Modern Language Features]]

