---
title: "Race Conditions, Data Races, Thread Sanitizer (TSan), and the ABA Problem Failure Modes and Edge Cases"
tags:
  - review
  - computer-science
  - systems-engineering
  - concurrency-and-multithreading
  - principal-swe
parent: "[[Race Conditions, Data Races, Thread Sanitizer (TSan), and the ABA Problem]]"
---

# Race Conditions, Data Races, Thread Sanitizer (TSan), and the ABA Problem Failure Modes and Edge Cases

## 1. Failure Modes Matrix

```text
┌──────────────────────────────────────┬──────────────────────────────────────────┬────────────────────────────────────────┐
│ Failure Mode                         │ Systems Root Cause                       │ Mitigation Strategy                    │
├──────────────────────────────────────┼──────────────────────────────────────────┼────────────────────────────────────────┤
│ 1. State Inconsistency Under Partit. │ Network partition splits cluster quorum  │ Quorum fencing, generation IDs, leases │
├──────────────────────────────────────┼──────────────────────────────────────────┼────────────────────────────────────────┤
│ 2. Unbounded Queue Overflow          │ Consumer latency spike starves workers   │ Bounded ring buffers & backpressure    │
├──────────────────────────────────────┼──────────────────────────────────────────┼────────────────────────────────────────┤
│ 3. Cascading Retry Storms            │ Synchronized client retries upon timeout │ Exponential backoff with full jitter   │
├──────────────────────────────────────┼──────────────────────────────────────────┼────────────────────────────────────────┤
│ 4. Resource / File Descriptor Leak   │ Unclosed channels / connections on error │ Strict defer cleanup & connection pool │
└──────────────────────────────────────┴──────────────────────────────────────────┴────────────────────────────────────────┘
```

---

## 2. Root Cause Diagnostic Playbook
1. **Quorum Verification:** Ensure total replicas $N \ge 2f + 1$ to survive $f$ crash faults.
2. **Backpressure Validation:** Ensure all processing channels feature bounded capacity to avoid out-of-memory panics.

---

## 🔗 References
- ⬆️ Parent: [[Race Conditions, Data Races, Thread Sanitizer (TSan), and the ABA Problem]]
- 📚 Module: `Concurrency, Multithreading & Memory Models`
