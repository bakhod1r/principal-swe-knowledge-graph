---
title: Actor Model and CSP
tags:
  - programming
  - paradigms
  - principal-swe
parent: "[[Programming Paradigms]]"
---

# Actor Model and CSP

Mailboxes, message immutability, Erlang/Akka supervision trees (Let It Crash), Go CSP channels, and distributed actor ordering.

```text
Actor Model and CSP
│
├── [[Actor Model Architecture (Mailboxes, Message Immutability, Erlang-Akka)]]
├── [[Supervision Trees and Let It Crash Failure Handling]]
├── [[Communicating Sequential Processes (CSP) Channels and Select (Go)]]
└── [[Deadlock and Message Ordering Guarantees in Distributed Actors]]
```

---

## 🗂️ Topics

- [[Actor Model Architecture (Mailboxes, Message Immutability, Erlang-Akka)]] — Autonomous concurrent entities communicating exclusively through immutable asynchronous message passing.
- [[Supervision Trees and Let It Crash Failure Handling]] — Hierarchical fault isolation: restarting crashed child actors rather than littering code with defensive checks.
- [[Communicating Sequential Processes (CSP) Channels and Select (Go)]] — Synchronizing concurrent processes through explicit message channels without shared memory.
- [[Deadlock and Message Ordering Guarantees in Distributed Actors]] — Evaluating causal ordering, FIFO mailboxes, and split-brain scenarios in clustered actor meshes.

---

## 🔗 References
- ⬆️ Parent: [[Programming Paradigms]]
- 🎓 Root: [[Principal SWE]]
