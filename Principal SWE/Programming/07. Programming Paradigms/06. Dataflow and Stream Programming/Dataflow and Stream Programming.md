---
title: Dataflow and Stream Programming
tags:
  - programming
  - paradigms
  - principal-swe
parent: "[[Programming Paradigms]]"
---

# Dataflow and Stream Programming

DAG execution engines (Flink, Spark), pipelined streaming vs micro-batching, windowing watermarks, and Chandy-Lamport checkpointing.

```text
Dataflow and Stream Programming
│
├── [[Directed Acyclic Graph (DAG) Execution Engines (Flink, Spark)]]
├── [[Pipelined Stream Processing vs Micro-Batching]]
├── [[Windowing Semantics (Tumbling, Sliding, Session Windows) and Watermarks]]
└── [[Stateful Stream Processing and Distributed Checkpointing (Chandy-Lamport)]]
```

---

## 🗂️ Topics

- [[Directed Acyclic Graph (DAG) Execution Engines (Flink, Spark)]] — Constructing distributed execution graphs for continuous data transformation.
- [[Pipelined Stream Processing vs Micro-Batching]] — Evaluating true low-latency continuous event processing vs micro-batched windowing.
- [[Windowing Semantics (Tumbling, Sliding, Session Windows) and Watermarks]] — Handling out-of-order event time data with watermark heuristics and late-arrival triggers.
- [[Stateful Stream Processing and Distributed Checkpointing (Chandy-Lamport)]] — Achieving exactly-once processing state guarantees using asynchronous barrier snapshotting.

---

## 🔗 References
- ⬆️ Parent: [[Programming Paradigms]]
- 🎓 Root: [[Principal SWE]]
