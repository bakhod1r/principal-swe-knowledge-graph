---
title: Strings
tags:
  - golang
  - types
  - principal-swe
parent: "[[Data Types]]"
---

# Strings

Immutable byte sequences, string literals, indexing, slicing, and string operations.

```text
Strings
│
├── [[String Immutability]]
├── [[Raw String Literals]]
├── [[Interpreted String Literals]]
├── [[String Indexing and Slicing]]
├── [[String Concatenation Performance]]
└── [[strings.Builder Deep Dive]]
```

---

## 🗂️ Topics

- [[String Immutability]] — Why Go strings are immutable byte sequences and cannot be modified in-place.
- [[Raw String Literals]] — Multi-line unescaped strings enclosed in backticks (`).
- [[Interpreted String Literals]] — Double-quoted strings supporting escape sequences (\n, \t, \x).
- [[String Indexing and Slicing]] — Accessing raw byte offsets (s[i]) vs sub-slice expressions (s[i:j]).
- [[String Concatenation Performance]] — Comparing +, fmt.Sprintf, strings.Join, bytes.Buffer, and strings.Builder.
- [[strings.Builder Deep Dive]] — Zero-allocation string builder pattern and String() memory copy elimination.

---

## 🔗 References
- ⬆️ Parent: [[Data Types]]
- 🎓 Root: [[Principal SWE]]
