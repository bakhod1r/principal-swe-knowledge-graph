---
title: Object-Oriented Programming
tags:
  - programming
  - oop
  - principal-swe
parent: "[[Programming]]"
---

# 💻 Object-Oriented Programming

Object memory layouts, vtables, SOLID principles, GoF design patterns, DDD domain modeling, and anti-patterns.

```text
Object-Oriented Programming
│
├── [[Core OOP Mechanics|01. Core OOP Mechanics]]
│   ├── [[Classes, Structs & Object Memory Layouts (vtable, itab)]]
│   ├── [[Encapsulation, Data Hiding & Access Modifiers]]
│   ├── [[Inheritance vs Composition (Composition Over Inheritance)]]
│   └── [[Polymorphism (Dynamic Dispatch, Parametric, Ad-hoc)]]
├── [[SOLID & Design Patterns|02. SOLID & Design Patterns]]
│   ├── [[SOLID Principles Deep Dive & Architectural Intent]]
│   ├── [[Gang of Four (GoF) Patterns Classification]]
│   ├── [[Creational, Structural & Behavioral Patterns in Modern Systems]]
│   └── [[Domain-Driven Design (DDD) Aggregates, Entities & Value Objects]]
└── [[OOP Anti-Patterns & Critiques|03. OOP Anti-Patterns & Critiques]]
│   ├── [[Anemic Domain Model vs Rich Domain Model]]
│   ├── [[God Object & Fragile Base Class Problem]]
│   └── [[Object-Relational Impedance Mismatch]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Core OOP Mechanics|01. Core OOP Mechanics]]
- [[Classes, Structs & Object Memory Layouts (vtable, itab)]] — Virtual method tables (vtable), interface method tables (itab), hidden pointers, and memory alignment.
- [[Encapsulation, Data Hiding & Access Modifiers]] — Protecting internal state invariants, access boundaries, and encapsulation vs data-holding structs.
- [[Inheritance vs Composition (Composition Over Inheritance)]] — Why deep inheritance hierarchies create fragile base classes and how composition enables flexible reuse.
- [[Polymorphism (Dynamic Dispatch, Parametric, Ad-hoc)]] — Comparing subtype polymorphism (vtables), parametric polymorphism (generics), and ad-hoc overloading.
### 2. 📂 [[SOLID & Design Patterns|02. SOLID & Design Patterns]]
- [[SOLID Principles Deep Dive & Architectural Intent]] — Single Responsibility, Open/Closed, Liskov Substitution, Interface Segregation, and Dependency Inversion.
- [[Gang of Four (GoF) Patterns Classification]] — Comprehensive taxonomy: Creational, Structural, and Behavioral patterns in enterprise systems.
- [[Creational, Structural & Behavioral Patterns in Modern Systems]] — Factory, Builder, Adapter, Decorator, Strategy, Observer, Command, and Template Method in modern architectures.
- [[Domain-Driven Design (DDD) Aggregates, Entities & Value Objects]] — Structuring rich domain models: Entity identity, immutable Value Objects, and transactional Aggregate Roots.
### 3. 📂 [[OOP Anti-Patterns & Critiques|03. OOP Anti-Patterns & Critiques]]
- [[Anemic Domain Model vs Rich Domain Model]] — Why separating data from behavior creates procedural code disguised as OOP and how to build rich domains.
- [[God Object & Fragile Base Class Problem]] — Managing sprawling multi-thousand line classes and base class modifications that silently break subclasses.
- [[Object-Relational Impedance Mismatch]] — The fundamental architectural conflict between relational relational tables and graph-like object memory.

---

## 🔗 Navigation
- ⬆️ Parent: [[Programming]]
- 🎓 Root: [[Principal SWE]]
