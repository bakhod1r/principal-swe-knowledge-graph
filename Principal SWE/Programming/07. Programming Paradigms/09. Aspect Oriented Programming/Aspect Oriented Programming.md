---
title: Aspect Oriented Programming
tags:
  - programming
  - paradigms
  - principal-swe
parent: "[[Programming Paradigms]]"
---

# Aspect Oriented Programming

Cross-cutting concerns, pointcuts, advices, compile/load/runtime bytecode weaving, and AOP control flow pitfalls.

```text
Aspect Oriented Programming
│
├── [[Cross-Cutting Concerns (Logging, Security, Transactions)]]
├── [[Join Points, Pointcuts, Advices, and Weaving Mechanisms]]
├── [[Compile-Time vs Load-Time vs Runtime Bytecode Weaving]]
└── [[Pitfalls of AOP: Hidden Control Flow and Debugging Friction]]
```

---

## 🗂️ Topics

- [[Cross-Cutting Concerns (Logging, Security, Transactions)]] — Separating orthogonal infrastructural concerns from core business domain logic.
- [[Join Points, Pointcuts, Advices, and Weaving Mechanisms]] — Specifying where and when aspects intercept execution flow (Before, After, Around).
- [[Compile-Time vs Load-Time vs Runtime Bytecode Weaving]] — Evaluating static AST weaving vs JVM agent bytecode manipulation and dynamic proxy wrapping.
- [[Pitfalls of AOP: Hidden Control Flow and Debugging Friction]] — Why implicit aspect side effects create cognitive friction and difficult-to-trace production bugs.

---

## 🔗 References
- ⬆️ Parent: [[Programming Paradigms]]
- 🎓 Root: [[Principal SWE]]
