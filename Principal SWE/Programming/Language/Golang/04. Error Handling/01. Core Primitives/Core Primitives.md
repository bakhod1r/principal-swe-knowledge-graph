---
title: Core Primitives
tags:
  - golang
  - error-handling
parent: "[[Error Handling]]"
---

# Core Primitives

Error values, the error interface, errors.New, formatting, and nil traps.

```text
Core Primitives
│
├── [[Errors as First-Class Values]]
├── [[error Interface Contract]]
├── [[errors.New]]
├── [[fmt.Errorf Formatting]]
├── [[Error String Conventions]]
└── [[nil Error Pitfall]]
```

---

## 🗂️ Topics

- [[Errors as First-Class Values]] — Explicit error values, if err != nil idiom, and error control flow.
- [[error Interface Contract]] — The single-method built-in interface contract: Error() string.
- [[errors.New]] — Creating simple static error values using errors.New().
- [[fmt.Errorf Formatting]] — Formatting dynamic error messages with %s, %d, and %v.
- [[Error String Conventions]] — Lowercase, no trailing punctuation, descriptive error phrasing standards.
- [[nil Error Pitfall]] — Typed nil pointer assigned to error interface creating non-nil interface trap.

---

## 🔗 References
- ⬆️ Parent: [[Error Handling]]
- 🎓 Root: [[Principal SWE]]
