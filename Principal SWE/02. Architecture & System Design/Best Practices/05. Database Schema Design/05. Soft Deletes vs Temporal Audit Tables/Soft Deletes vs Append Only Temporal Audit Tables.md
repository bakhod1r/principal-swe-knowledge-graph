---
title: Soft Deletes vs Append Only Temporal Audit Tables
tags:
  - best-practices
  - software-engineering
  - database-schema-design-and-migration-best-practices
  - principal-swe
parent: "[[Database Schema Design]]"
---

# 📦 Soft Deletes vs Append Only Temporal Audit Tables

Trade-offs of `is_deleted` flags (index poisoning, unique constraint complexity) versus dedicated audit tables, CDC outboxes, and PostgreSQL partition dropping.

```text
Soft Deletes vs Append Only Temporal Audit Tables
│
├── [[Soft Deletes vs Append Only Temporal Audit Tables Engineering Standards and Principles]]
├── [[Soft Deletes vs Append Only Temporal Audit Tables Production Implementation Patterns]]
└── [[Soft Deletes vs Append Only Temporal Audit Tables Failure Modes and Anti Pattern Mitigations]]
```

---

## 🗂️ Engineering Standards & Patterns

- [[Soft Deletes vs Append Only Temporal Audit Tables Engineering Standards and Principles]]
- [[Soft Deletes vs Append Only Temporal Audit Tables Production Implementation Patterns]]
- [[Soft Deletes vs Append Only Temporal Audit Tables Failure Modes and Anti Pattern Mitigations]]

---

## 🔗 References
- ⬆️ Parent: `Database Schema Design & Migration Best Practices`
- 📚 Module: `Best Practices`

