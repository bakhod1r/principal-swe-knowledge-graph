---
title: Zero Downtime Database Migrations (expand and Contract Pattern)
tags:
  - review
  - best-practices
  - software-engineering
  - database-schema-design-and-migration-best-practices
  - principal-swe
parent: "[[Migration Best Practices]]"
---

# 📦 Zero Downtime Database Migrations (expand and Contract Pattern)

Three-phase schema evolution: 1) Expand (add nullable column/view), 2) Migrate data & dual-write, 3) Contract (deprecate & drop old column).

```text
Zero Downtime Database Migrations (expand and Contract Pattern)
│
├── [[Zero Downtime Database Migrations (expand and Contract Pattern) Engineering Standards and Principles]]
├── [[Zero Downtime Database Migrations (expand and Contract Pattern) Production Implementation Patterns]]
└── [[Zero Downtime Database Migrations (expand and Contract Pattern) Failure Modes and Anti Pattern Mitigations]]
```

---

## 🗂️ Engineering Standards & Patterns

- [[Zero Downtime Database Migrations (expand and Contract Pattern) Engineering Standards and Principles]]
- [[Zero Downtime Database Migrations (expand and Contract Pattern) Production Implementation Patterns]]
- [[Zero Downtime Database Migrations (expand and Contract Pattern) Failure Modes and Anti Pattern Mitigations]]

---

## 🔗 References
- ⬆️ Parent: `Database Schema Design & Migration Best Practices`
- 📚 Module: `Best Practices`

