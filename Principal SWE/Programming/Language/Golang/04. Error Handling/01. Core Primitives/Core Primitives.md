---
title: Core Primitives
tags:
  - golang
  - error-handling
parent: "[[Error Handling]]"
---

# Core Primitives

The error interface, sentinel errors, errors.New, and fmt.Errorf formatting.

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

- [[Errors as First-Class Values]] — Errors as explicit return values, checking if err != nil.
- [[error Interface Contract]] — The built-in single-method interface contract (Error() string).
- [[errors.New]] — Creating simple static error values.
- [[fmt.Errorf Formatting]] — Formatting dynamic error messages with %v, %s, and %d.
- [[Error String Conventions]] — Lowercase, no trailing punctuation, descriptive message rules.
- [[nil Error Pitfall]] — Typed nil pointer assigned to error interface causing non-nil interface bug.

---

## 🔗 References
- ⬆️ Parent: [[Error Handling]]
- 🎓 Root: [[Principal SWE]]
