---
title: Runtime Stack Management & Unwinding (stack.go, traceback.go)
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# Runtime Stack Management & Unwinding (stack.go, traceback.go)

Contiguous stack growth (copystack), stack shrinking, gentraceback stack unwinding, and //go:nosplit.

```text
Runtime Stack Management & Unwinding (stack.go, traceback.go)
│
├── [[Stack Growth Engine (runtime.morestack & copystack)]]
├── [[Stack Shrinking Mechanics & GC Stack Scavenging]]
├── [[Stack Traceback Unwinding Algorithm (gentraceback)]]
├── [[go:nosplit Pragma & Stack Overflow Prevention]]
└── [[Stack Segment Bounds & Guard Page Protection]]
```

---

## 🗂️ Topics

- [[Stack Growth Engine (runtime.morestack & copystack)]] — Allocating new contiguous stack, copying frames, and updating pointer fixup tables.
- [[Stack Shrinking Mechanics & GC Stack Scavenging]] — Shrinking stack to half size during GC mark termination when utilization drops below 25%.
- [[Stack Traceback Unwinding Algorithm (gentraceback)]] — Traversing call frames using PC/SP tables for panic dumps, stack traces, and pprof.
- [[go:nosplit Pragma & Stack Overflow Prevention]] — Preventing infinite morestack recursion in scheduler and allocator leaf functions via //go:nosplit.
- [[Stack Segment Bounds & Guard Page Protection]] — Detecting hardware stack overflows via OS memory guard pages and stack bottom boundary checks.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]
- 🎓 Root: [[Principal SWE]]
