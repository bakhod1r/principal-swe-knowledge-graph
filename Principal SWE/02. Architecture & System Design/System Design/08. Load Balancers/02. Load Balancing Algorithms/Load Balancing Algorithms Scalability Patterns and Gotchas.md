---
title: "Load Balancing Algorithms Scalability Patterns and Gotchas"
tags:
  - review
  - system-design
  - architecture
  - distributed-systems
  - load-balancers
  - principal-swe
parent: "[[Load Balancing Algorithms]]"
---

# Load Balancing Algorithms Scalability Patterns and Gotchas

## 1. Definition
**Load Balancing Algorithms Scalability Patterns** defines the horizontal sharding models, multi-tier caching architectures, and throughput scaling strategies for **Load Balancing Algorithms**.
It enables systems to scale linearly from $10^3$ to $10^7$ QPS while maintaining bounded memory footprints, sub-millisecond p99 latencies, and zero data loss across distributed compute nodes.

---

## 2. Mental Model
```text
High-Throughput Sharded Dataflow:
Input Traffic (1M QPS)
          │
          ▼
 [ Consistent Hash Ring / Partition Router ]
   │                 │                 │
   ▼ (Hash Slot 0)   ▼ (Hash Slot 1)   ▼ (Hash Slot 2)
[ Shard Node 0 ]   [ Shard Node 1 ]   [ Shard Node 2 ]
  ├── L1 Local Cache ├── L1 Local Cache ├── L1 Local Cache
  └── L2 Distributed └── L2 Distributed └── L2 Distributed
```
- **Read-Write Decoupling:** Directs high-frequency reads to distributed read replicas and caches (CQRS pattern) while routing mutations through an append-only write path.

---

## 3. Usage
```yaml
# Kubernetes & Envoy Horizontal Autoscaling Topology for Load Balancing Algorithms
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: load-balancing-algorithms-autoscaler
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: load-balancing-algorithms-service
  minReplicas: 10
  maxReplicas: 200
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 65
```

---

## 4. Gotchas
- **Hot Shard / Celebrity Key Problem:** A single high-volume key (e.g., viral user) routing to a single partition overwhelms the node, bypassing horizontal sharding benefits. Mitigate with two-tier salt hashing (`key_salt_0..N`).
- **Thundering Herd / Cache Stampede:** Cache expiration of critical metadata triggers millions of concurrent backend DB hits. Mitigate using single-flight mutexes and probabilistic early expiration (XFetch algorithm).

---

## 🔗 References
- ⬆️ Parent: [[Load Balancing Algorithms]]
- 📚 Module: `Load Balancers`

