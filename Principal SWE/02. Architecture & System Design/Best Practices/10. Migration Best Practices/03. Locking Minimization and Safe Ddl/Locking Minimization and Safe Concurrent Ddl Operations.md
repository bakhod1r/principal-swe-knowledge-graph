---
title: Locking Minimization and Safe Concurrent Ddl Operations
tags:
  - review
  - best-practices
  - software-engineering
  - database-schema-design-and-migration-best-practices
  - principal-swe
parent: "[[Migration Best Practices]]"
---

# 📦 Locking Minimization and Safe Concurrent Ddl Operations

PostgreSQL `CREATE INDEX CONCURRENTLY`, setting short `lock_timeout`, avoiding table rewrite locks during column additions with default values.

```text
Locking Minimization and Safe Concurrent Ddl Operations
│
├── [[Locking Minimization and Safe Concurrent Ddl Operations Engineering Standards and Principles]]
├── [[Locking Minimization and Safe Concurrent Ddl Operations Production Implementation Patterns]]
└── [[Locking Minimization and Safe Concurrent Ddl Operations Failure Modes and Anti Pattern Mitigations]]
```

---

## 🗂️ Engineering Standards & Patterns

- [[Locking Minimization and Safe Concurrent Ddl Operations Engineering Standards and Principles]]
- [[Locking Minimization and Safe Concurrent Ddl Operations Production Implementation Patterns]]
- [[Locking Minimization and Safe Concurrent Ddl Operations Failure Modes and Anti Pattern Mitigations]]

---

## 🔗 References
- ⬆️ Parent: `Database Schema Design & Migration Best Practices`
- 📚 Module: `Best Practices`

