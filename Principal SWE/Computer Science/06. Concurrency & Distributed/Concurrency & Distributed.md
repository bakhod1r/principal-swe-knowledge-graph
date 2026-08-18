---
title: Concurrency & Distributed
tags:
  - computer-science
  - concurrency-&-distributed
  - principal-swe
parent: "[[Computer Science]]"
---

# 🏛️ Concurrency & Distributed (Foundations & Systems Architecture)

Hardware and software memory models, happens-before consistency, lock-free/wait-free data structures, Amdahl/USL scaling, CAP/PACELC trade-offs, FLP impossibility, Paxos/Raft consensus, and Byzantine fault tolerance.

```text
Concurrency & Distributed
│
├── [[Hardware and Software Memory Models|01. Hardware Memory Models]]
├── [[Happens Before Relationship and Memory Ordering|02. Happens Before Invariants]]
├── [[Atomic Instructions and Compare and Swap (cas)|03. Atomics and CAS]]
├── [[Lock Free and Wait Free Concurrent Algorithms|04. Lock Free and Wait Free]]
├── [[Amdahl's Law and Universal Scalability Law (usl)|05. Amdahl and USL Scaling]]
├── [[Race Conditions, Data Races, and the ABA Problem|06. Race Conditions and ABA]]
├── [[Actor Model and Communicating Sequential Processes (csp)|07. Actor Model and CSP]]
├── [[Software Transactional Memory (stm)|08. Transactional Memory]]
├── [[Work Stealing Task Schedulers|09. Work Stealing Schedulers]]
├── [[Read Copy Update (rcu) and Hazard Pointers|10. RCU and Hazard Pointers]]
├── [[CAP Theorem and PACELC Framework|11. CAP and PACELC Theorems]]
├── [[FLP Impossibility Theorem|12. FLP Impossibility]]
├── [[Distributed Consistency Models Hierarchy|13. Distributed Consistency Models]]
├── [[Lamport Logical Clocks and Vector Clocks|14. Logical and Vector Clocks]]
├── [[Quorum Systems and Gossip Protocols|15. Quorums and Gossip Protocols]]
├── [[Foundations of Distributed Consensus|16. Distributed Consensus Foundations]]
├── [[Byzantine Fault Tolerance and PBFT|17. Byzantine Fault Tolerance PBFT]]
├── [[State Machine Replication (smr) Principle|18. State Machine Replication]]
└── [[Cluster Membership and Failure Detectors|19. Membership and Failure Detectors]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Hardware and Software Memory Models|01. Hardware Memory Models]] — Sequential Consistency (SC), TSO, Release Consistency, Java/Go memory models, and data-race-free (DRF).
- 📂 [[Happens Before Relationship and Memory Ordering|02. Happens Before Invariants]] — Partial ordering of program execution, synchronization edges, volatile/atomic variables, and transitive orderings.
- 📂 [[Atomic Instructions and Compare and Swap (cas)|03. Atomics and CAS]] — Hardware atomic RMW instructions (LOCK CMPXCHG), load-linked/store-conditional (LL/SC), and atomic fetch-and-add.
- 📂 [[Lock Free and Wait Free Concurrent Algorithms|04. Lock Free and Wait Free]] — Non-blocking progress guarantees: lock-free (at least one thread makes progress) vs wait-free (every thread bounded).
- 📂 [[Amdahl's Law and Universal Scalability Law (usl)|05. Amdahl and USL Scaling]] — Quantifying multi-core speedup limits: serial contention bottlenecks (Amdahl) and coherency crosstalk penalties (USL).
- 📂 [[Race Conditions, Data Races, and the ABA Problem|06. Race Conditions and ABA]] — Data races vs race conditions, pointer recycling ABA hazard, tagged pointers, and double-word CAS (CAS2).
- 📂 [[Actor Model and Communicating Sequential Processes (csp)|07. Actor Model and CSP]] — Erlang/Akka asynchronous message passing vs Go channel synchronization and share-memory-by-communicating.
- 📂 [[Software Transactional Memory (stm)|08. Transactional Memory]] — Optimistic concurrency control in RAM: speculative memory reads/writes, conflict detection, and atomic commit/abort.
- 📂 [[Work Stealing Task Schedulers|09. Work Stealing Schedulers]] — Chase-Lev lock-free deques, randomized victim stealing, child-stealing vs continuation-stealing, and Go runtime scheduler.
- 📂 [[Read Copy Update (rcu) and Hazard Pointers|10. RCU and Hazard Pointers]] — Lock-free memory reclamation: quiescent states, grace periods, and protecting shared pointers from premature free.
- 📂 [[CAP Theorem and PACELC Framework|11. CAP and PACELC Theorems]] — Formal distributed trade-offs: Consistency vs Availability under Partitions; Latency vs Consistency under Normal operations.
- 📂 [[FLP Impossibility Theorem|12. FLP Impossibility]] — Mathematical proof that deterministic consensus is impossible in an asynchronous network with a single unannounced crash.
- 📂 [[Distributed Consistency Models Hierarchy|13. Distributed Consistency Models]] — Strict Linearizability, Sequential Consistency, Causal Consistency, Read-Your-Writes, and Eventual Consistency.
- 📂 [[Lamport Logical Clocks and Vector Clocks|14. Logical and Vector Clocks]] — Partial orderings of distributed events, tracking causality, concurrent event detection, and interval tree clocks.
- 📂 [[Quorum Systems and Gossip Protocols|15. Quorums and Gossip Protocols]] — Strict majority quorums (R + W > N), sloppy quorums with hinted handoff, and anti-entropy gossip dissemination.
- 📂 [[Foundations of Distributed Consensus|16. Distributed Consensus Foundations]] — Safety vs liveness invariants, leader election, log replication, commit rules, and epoch numbers in Raft/Paxos.
- 📂 [[Byzantine Fault Tolerance and PBFT|17. Byzantine Fault Tolerance PBFT]] — Tolerating arbitrary malicious and corrupted nodes with 3f+1 node quorums (PBFT 3-phase pre-prepare, prepare, commit).
- 📂 [[State Machine Replication (smr) Principle|18. State Machine Replication]] — Deterministic state machines executing identical ordered command logs to achieve fault-tolerant active replication.
- 📂 [[Cluster Membership and Failure Detectors|19. Membership and Failure Detectors]] — Heartbeating, Phi Accrual failure detectors, SWIM gossip membership, and split-brain fencing tokens.

---

## 🔗 References
- ⬆️ Parent: [[Computer Science]]
- 🎓 Root: [[Principal SWE]]
