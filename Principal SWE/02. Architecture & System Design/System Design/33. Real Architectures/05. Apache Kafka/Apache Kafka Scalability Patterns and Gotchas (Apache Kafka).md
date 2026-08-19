---
title: "Apache Kafka Scalability Patterns and Gotchas (Apache Kafka)"
tags:
  - review
  - system-design
  - architecture
  - distributed-systems
  - real-architectures
  - principal-swe
parent: "[[Apache Kafka (Real Architectures)]]"
---

# Apache Kafka Scalability Patterns and Gotchas (Apache Kafka)

## 1. Definition
**Apache Kafka (Apache Kafka) Scalability Patterns** defines the horizontal sharding models, multi-tier caching architectures, and throughput scaling strategies for **Apache Kafka (Apache Kafka)**.
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
# Kubernetes & Envoy Horizontal Autoscaling Topology for Apache Kafka (Apache Kafka)
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: apache-kafka-(apache-kafka)-autoscaler
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: apache-kafka-(apache-kafka)-service
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
- ⬆️ Parent: [[Apache Kafka (Real Architectures)]]
- 📚 Module: `Real Architectures`

