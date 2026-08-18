---
title: Suffix Structures and Text Indexing
tags:
  - computer-science
  - algorithms
  - string-algorithms
  - principal-swe
parent: "[[String Algorithms]]"
---

# Suffix Structures and Text Indexing

Suffix Automaton (SAM), Suffix Array & LCP Array, Suffix Tree (Ukkonen), FM-Index and Burrows-Wheeler Transform.

```text
Suffix Structures and Text Indexing
│
├── [[Suffix Automaton (SAM) Directed Acyclic Word Graph]]
├── [[Suffix Array and Longest Common Prefix (LCP) Array]]
├── [[Suffix Tree Construction (Ukkonen Algorithm)]]
└── [[FM-Index and Burrows-Wheeler Transform (BWT)]]
```

---

## 🗂️ Topics

- [[Suffix Automaton (SAM) Directed Acyclic Word Graph]] — Minimal deterministic automaton recognizing all suffixes of a string with O(N) states and transitions.
- [[Suffix Array and Longest Common Prefix (LCP) Array]] — Sorted lexicographical suffix indices in O(N log N) or O(N) using SA-IS and Kasai s LCP algorithm.
- [[Suffix Tree Construction (Ukkonen Algorithm)]] — Linear-time O(N) on-line suffix tree construction using implicit suffix extensions and skip/count tricks.
- [[FM-Index and Burrows-Wheeler Transform (BWT)]] — Compressed text index enabling exact substring search in O(P) time within entropy-compressed memory space.

---

## 🔗 References
- ⬆️ Parent: [[String Algorithms]]
- 🎓 Root: [[Principal SWE]]
