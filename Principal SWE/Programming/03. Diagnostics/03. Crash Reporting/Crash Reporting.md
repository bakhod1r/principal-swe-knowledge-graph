---
title: Crash Reporting
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Crash Reporting

Minidump/coredump generation, DWARF symbolication pipelines, crash clustering, and crash-loop backoff defense.

```text
Crash Reporting
│
├── [[Automated Minidump and Core Dump Generation in Production]]
├── [[Symbolication Pipelines and DWARF-PDB Debugging Symbol Management]]
├── [[Crash Deduplication, Clustering, and Blast Radius Grouping]]
└── [[Circuit Breaking and Crash-Loop Backoff Defense]]
```

---

## 🗂️ Topics

- [[Automated Minidump and Core Dump Generation in Production]] — Capturing CPU register states and memory stack traces upon process segfaults or panics.
- [[Symbolication Pipelines and DWARF-PDB Debugging Symbol Management]] — Mapping raw stack addresses to exact source code files, git commits, and line numbers.
- [[Crash Deduplication, Clustering, and Blast Radius Grouping]] — Grouping millions of crashes by stack signature to isolate top critical production incidents.
- [[Circuit Breaking and Crash-Loop Backoff Defense]] — Preventing cascading crash loops from overwhelming storage backends or message brokers.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]
- 🎓 Root: [[Principal SWE]]
