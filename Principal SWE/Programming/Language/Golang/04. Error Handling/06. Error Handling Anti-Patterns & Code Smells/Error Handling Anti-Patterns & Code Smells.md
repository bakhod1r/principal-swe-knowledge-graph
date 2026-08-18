---
title: Error Handling Anti-Patterns & Code Smells
tags:
  - golang
  - error-handling
  - principal-swe
parent: "[[Error Handling]]"
---

# Error Handling Anti-Patterns & Code Smells

Blind error ignoring, double logging, string matching, catch-all errors, and staff-level error handling principles.

```text
Error Handling Anti-Patterns & Code Smells
│
├── [[The Blind Error Ignoration Anti-Pattern (_ = fn())]]
├── [[Double Logging & Double Handling Anti-Pattern]]
├── [[String Matching on Errors Anti-Pattern]]
├── [[Catch-All Generic Internal Error Anti-Pattern]]
└── [[Staff-Level Error Handling Principles & Guidelines]]
```

---

## 🗂️ Topics

- [[The Blind Error Ignoration Anti-Pattern (_ = fn())]] — Security and state corruption risks of silently swallowing return errors.
- [[Double Logging & Double Handling Anti-Pattern]] — Logging an error and returning it upward, leading to log spam and duplicate alert storms.
- [[String Matching on Errors Anti-Pattern]] — Fragile strings.Contains(err.Error(), "not found") checks vs robust errors.Is/errors.As.
- [[Catch-All Generic Internal Error Anti-Pattern]] — Wrapping everything in generic errors.New("something went wrong") destroying observability.
- [[Staff-Level Error Handling Principles & Guidelines]] — Non-negotiable engineering principles: wrap once at the boundary, handle once, log once.

---

## 🔗 References
- ⬆️ Parent: [[Error Handling]]

