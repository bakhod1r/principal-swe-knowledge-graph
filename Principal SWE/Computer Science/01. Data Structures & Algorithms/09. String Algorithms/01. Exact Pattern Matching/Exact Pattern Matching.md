---
title: Exact Pattern Matching
tags:
  - computer-science
  - algorithms
  - string-algorithms
  - principal-swe
parent: "[[String Algorithms]]"
---

# Exact Pattern Matching

KMP failure function, Boyer-Moore bad character/good suffix, Rabin-Karp rolling hash, Aho-Corasick, Z-Algorithm.

```text
Exact Pattern Matching
│
├── [[Knuth-Morris-Pratt (KMP) and Prefix Function]]
├── [[Boyer-Moore-Horspool Algorithm and Bad Character Rule]]
├── [[Rabin-Karp Rolling Hash and Polynomial Fingerprinting]]
├── [[Aho-Corasick Multi-Pattern Automaton]]
└── [[Z-Algorithm and Z-Array Construction]]
```

---

## 🗂️ Topics

- [[Knuth-Morris-Pratt (KMP) and Prefix Function]] — O(N + M) pattern matching using Longest Prefix Suffix (LPS) lookup table to prevent backtracking.
- [[Boyer-Moore-Horspool Algorithm and Bad Character Rule]] — Sublinear average time pattern searching scanning pattern right-to-left and jumping characters.
- [[Rabin-Karp Rolling Hash and Polynomial Fingerprinting]] — O(N) search using rolling polynomial modular hashing with prime base arithmetic.
- [[Aho-Corasick Multi-Pattern Automaton]] — Trie augmented with failure and dictionary links matching thousands of keywords simultaneously in O(N + M).
- [[Z-Algorithm and Z-Array Construction]] — O(N) prefix-matching array computing longest common prefix with string head for every position.

---

## 🔗 References
- ⬆️ Parent: [[String Algorithms]]
- 🎓 Root: [[Principal SWE]]
