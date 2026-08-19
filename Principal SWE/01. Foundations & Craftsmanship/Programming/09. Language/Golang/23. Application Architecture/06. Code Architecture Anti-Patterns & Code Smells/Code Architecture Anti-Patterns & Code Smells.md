---
title: Code Architecture Anti-Patterns & Code Smells
tags:
  - golang
  - architecture
  - principal-swe
parent: "[[Application Architecture]]"
---

# Code Architecture Anti-Patterns & Code Smells

Global state singletons, god packages, package clutter, package stuttering, cyclic dependency hacks, and Staff architecture checklist.

```text
Code Architecture Anti-Patterns & Code Smells
│
├── [[The Global State & Singleton Anti-Pattern]]
├── [[The God Package & Package Clutter Anti-Pattern]]
├── [[The Package Stuttering Anti-Pattern]]
├── [[Cyclic Dependency Workaround Hacks]]
└── [[Staff-Level Code Architecture Checklist]]
```

---

## 🗂️ Topics

- [[The Global State & Singleton Anti-Pattern]] — Concurrency race hazards, testing pollution, and hidden coupling caused by global variables.
- [[The God Package & Package Clutter Anti-Pattern]] — Giant monolithic packages containing everything and dumping code into helpers/.
- [[The Package Stuttering Anti-Pattern]] — Redundant identifiers (user.UserService, client.ClientConfig) harming idiomatic Go readability.
- [[Cyclic Dependency Workaround Hacks]] — Dangerous anti-patterns: using init() hooks or type casting to bypass circular imports.
- [[Staff-Level Code Architecture Checklist]] — Pre-production architectural review checklist for high-scale enterprise Go services.

---

## 🔗 References
- ⬆️ Parent: `Code Organization & Architecture`

