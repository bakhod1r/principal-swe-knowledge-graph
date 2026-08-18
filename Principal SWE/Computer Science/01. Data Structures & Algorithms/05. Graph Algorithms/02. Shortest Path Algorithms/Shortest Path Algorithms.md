---
title: Shortest Path Algorithms
tags:
  - computer-science
  - algorithms
  - graph-algorithms
  - principal-swe
parent: "[[Graph Algorithms]]"
---

# Shortest Path Algorithms

Dijkstra, Bellman-Ford, Floyd-Warshall, Johnson, A* Search.

```text
Shortest Path Algorithms
│
├── [[Dijkstra Algorithm with Indexed Priority Queue]]
├── [[Bellman-Ford and Negative Cycle Detection]]
├── [[Floyd-Warshall All-Pairs Shortest Path]]
├── [[Johnson Algorithm for Sparse All-Pairs Shortest Path]]
└── [[A-Star Heuristic Search and Admissible Heuristics]]
```

---

## 🗂️ Topics

- [[Dijkstra Algorithm with Indexed Priority Queue]] — Greedy single-source shortest paths on non-negative weighted graphs in O((V + E) log V).
- [[Bellman-Ford and Negative Cycle Detection]] — Dynamic programming edge relaxation over V-1 iterations and detecting negative arbitrage cycles.
- [[Floyd-Warshall All-Pairs Shortest Path]] — O(V^3) matrix dynamic programming across intermediate vertices k for dense networks.
- [[Johnson Algorithm for Sparse All-Pairs Shortest Path]] — Reweighting graph edges via Bellman-Ford potentials to run V Dijkstra searches in O(V E log V).
- [[A-Star Heuristic Search and Admissible Heuristics]] — Goal-directed shortest path search using Euclidean/Manhattan distance heuristic estimates.

---

## 🔗 References
- ⬆️ Parent: [[Graph Algorithms]]
- 🎓 Root: [[Principal SWE]]
