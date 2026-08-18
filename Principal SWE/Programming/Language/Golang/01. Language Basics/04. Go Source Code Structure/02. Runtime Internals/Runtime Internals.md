---
title: Go Runtime Internals (src/runtime/)
tags:
  - golang
  - runtime
  - internals
parent: "[[Go Source Code Structure]]"
---

# ⚡ Go Runtime Internals (`src/runtime/`)

Deep internals powering Goroutine scheduling, memory allocation, garbage collection, and channels.

---

## 🗂️ Core Subsystems

- [[proc (Scheduler)|proc.go (GMP Scheduler)]] ⭐⭐⭐
- [[mgc (Garbage Collector)|mgc.go (Garbage Collector)]] ⭐⭐⭐
- [[malloc (Memory Allocator)|malloc.go (Memory Allocator)]] ⭐⭐⭐
- [[chan (Channels)|chan.go (Channels)]]
- [[map (Hash Maps)|map.go (Hash Maps)]]
- [[iface (Interfaces)|iface.go (Interfaces)]]
- [[panic (Panic & Recover)|panic.go (Panic & Recover)]]

---

## 🔗 Navigation
- ⬆️ Parent: [[Go Source Code Structure]]
