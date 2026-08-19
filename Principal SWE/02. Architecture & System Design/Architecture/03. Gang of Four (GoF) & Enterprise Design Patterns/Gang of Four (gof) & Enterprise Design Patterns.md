---
title: Gang of Four (gof) & Enterprise Design Patterns
tags:
  - review
  - architecture
  - systems-architecture
  - gang-of-four-(gof)-and-enterprise-design-patterns
  - principal-swe
parent: "[[Architecture]]"
---

# 🏛️ Gang of Four (gof) & Enterprise Design Patterns

Classic object-oriented and enterprise design patterns: Creational (Factory, Builder, Singleton), Structural (Adapter, Decorator, Facade, Proxy), Behavioral (Observer, Strategy, Command, State, Visitor, Memento, Iterator), and modern idioms.

```text
Gang of Four (gof) & Enterprise Design Patterns
│
├── [[Creational Patterns: Factory Method, Abstract Factory, and Object Families|01. Creational Patterns Factory Method and Abstract Factory]]
├── [[Creational Patterns: Fluent Builder, Step Builder, and Prototype Cloning|02. Creational Patterns Builder and Prototype]]
├── [[Creational Patterns: Thread Safe Singleton, Double Checked Locking, and Object Pools|03. Creational Patterns Singleton and Object Pool]]
├── [[Structural Patterns: Object Adapter, Class Adapter, and Bridge Decoupling|04. Structural Patterns Adapter and Bridge]]
├── [[Structural Patterns: Composite Tree Hierarchies and Dynamic Decorators|05. Structural Patterns Composite and Decorator]]
├── [[Structural Patterns: Subsystem Facade and Memory Efficient Flyweight|06. Structural Patterns Facade and Flyweight]]
├── [[Structural Patterns: Virtual Proxy, Remote Proxy, and Protection Gateways|07. Structural Patterns Proxy Virtual, Remote, and Protection]]
├── [[Behavioral Patterns: Strategy Algorithm Swapping and Template Method Invariants|08. Behavioral Patterns Strategy and Template Method]]
├── [[Behavioral Patterns: Observer Event Notification and Subject Decoupling|09. Behavioral Patterns Observer and Publish Subscribe]]
├── [[Behavioral Patterns: Command Encapsulation, Undo Redo, and Memento Snapshots|10. Behavioral Patterns Command and Memento]]
├── [[Behavioral Patterns: State Pattern and Object Oriented State Machines|11. Behavioral Patterns State and Finite State Machines]]
└── [[Behavioral Patterns: Double Dispatch Visitor, Iterator, and Centralized Mediator|12. Behavioral Patterns Visitor, Iterator, and Mediator]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Creational Patterns: Factory Method, Abstract Factory, and Object Families|01. Creational Patterns Factory Method and Abstract Factory]] — Decoupling object creation from consumption, abstract creation interfaces, parameterizing factory families, and dependency injection integration.
- 📂 [[Creational Patterns: Fluent Builder, Step Builder, and Prototype Cloning|02. Creational Patterns Builder and Prototype]] — Constructing complex composite objects step-by-step, immutable object construction, and deep cloning with the Prototype pattern.
- 📂 [[Creational Patterns: Thread Safe Singleton, Double Checked Locking, and Object Pools|03. Creational Patterns Singleton and Object Pool]] — Lazy initialization, double-checked locking, thread safety, connection pooling (DB, HTTP), and why singletons are often considered anti-patterns.
- 📂 [[Structural Patterns: Object Adapter, Class Adapter, and Bridge Decoupling|04. Structural Patterns Adapter and Bridge]] — Converting incompatible interfaces, wrapping legacy APIs, and separating an abstraction from its implementation with the Bridge pattern.
- 📂 [[Structural Patterns: Composite Tree Hierarchies and Dynamic Decorators|05. Structural Patterns Composite and Decorator]] — Treating individual objects and compositions uniformly, dynamic runtime behavior wrapping with Decorator, and combining with stream I/O.
- 📂 [[Structural Patterns: Subsystem Facade and Memory Efficient Flyweight|06. Structural Patterns Facade and Flyweight]] — Providing simplified entry points into complex subsystem graphs (Facade), and sharing fine-grained immutable state to minimize memory footprint (Flyweight).
- 📂 [[Structural Patterns: Virtual Proxy, Remote Proxy, and Protection Gateways|07. Structural Patterns Proxy Virtual, Remote, and Protection]] — Lazy loading large resources with Virtual Proxy, remote RPC stubs, access control verification with Protection Proxy, and dynamic proxies in frameworks.
- 📂 [[Behavioral Patterns: Strategy Algorithm Swapping and Template Method Invariants|08. Behavioral Patterns Strategy and Template Method]] — Encapsulating interchangeable algorithms behind interfaces (Strategy), and defining algorithmic skeletons with polymorphic hooks (Template Method).
- 📂 [[Behavioral Patterns: Observer Event Notification and Subject Decoupling|09. Behavioral Patterns Observer and Publish Subscribe]] — One-to-many dependency notifications, thread-safe subject state changes, memory leak prevention (Lapsed Listener Problem), and reactive event streams.
- 📂 [[Behavioral Patterns: Command Encapsulation, Undo Redo, and Memento Snapshots|10. Behavioral Patterns Command and Memento]] — Encapsulating operations as first-class objects (Command), queuing and scheduling commands, and restoring internal object state with Memento snapshots.
- 📂 [[Behavioral Patterns: State Pattern and Object Oriented State Machines|11. Behavioral Patterns State and Finite State Machines]] — Encapsulating state-specific behavior in dedicated classes, eliminating massive switch statements, and transition logic in workflow engines.
- 📂 [[Behavioral Patterns: Double Dispatch Visitor, Iterator, and Centralized Mediator|12. Behavioral Patterns Visitor, Iterator, and Mediator]] — Adding new operations to object structures without modifying them (Visitor), uniform traversal (Iterator), and decoupling complex multi-object communication (Mediator).

---

## 🔗 References
- ⬆️ Parent: [[Architecture]]

