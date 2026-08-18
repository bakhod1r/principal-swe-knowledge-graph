---
title: Data Structures & Algorithms
tags:
  - computer-science
  - algorithms
  - data-structures
  - principal-swe
parent: "[[Computer Science]]"
---

# ⚡ Data Structures & Algorithms

Comprehensive, hardware-conscious foundation of computational complexity, linear and non-linear memory structures, comparison and non-comparison sorting, search paradigms, graph flows and decompositions, dynamic programming optimizations, matroid greediness, exact cover backtracking, and string automata.

```text
Data Structures & Algorithms
│
├── [[Complexity Analysis|01. Complexity Analysis]]
│   ├── [[Asymptotic Bounds|01. Asymptotic Bounds]]
│   │   ├── [[Big-O Notation and Upper Bounds]]
│   │   ├── [[Big-Omega and Lower Bounds]]
│   │   ├── [[Big-Theta and Tight Bounds]]
│   │   ├── [[Little-o and Little-Omega Strict Bounds]]
│   │   └── [[Dominance Relations and Complexity Classes]]
│   ├── [[Recurrences and Amortization|02. Recurrences and Amortization]]
│   │   ├── [[Master Theorem for Divide and Conquer]]
│   │   ├── [[Akra-Bazzi Method for Non-Uniform Recurrences]]
│   │   ├── [[Accounting and Aggregate Amortized Analysis]]
│   │   ├── [[Potential Method and Energy Functions]]
│   │   └── [[Dynamic Array Resizing Amortization]]
│   └── [[Hardware-Aware Complexity|03. Hardware-Aware Complexity]]
│   │   ├── [[External Memory Model and IO Complexity]]
│   │   ├── [[Cache-Oblivious Algorithm Design]]
│   │   ├── [[Branch Prediction Penalty in Branchless Code]]
│   │   └── [[Memory Allocation Overhead and Fragmentation Cost]]
├── [[Data Structures|02. Data Structures]]
│   ├── [[Linear Structures|01. Linear Structures]]
│   │   ├── [[Dynamic Array Contiguity and Growth Policy]]
│   │   ├── [[Singly and Doubly Linked List Pointer Layouts]]
│   │   ├── [[Unrolled Linked List and Memory Density]]
│   │   ├── [[Circular Ring Buffer and Lock-Free Queues]]
│   │   └── [[Monotonic Stack and Monotonic Deque]]
│   ├── [[Hash-Based Structures|02. Hash-Based Structures]]
│   │   ├── [[Hash Table Chaining vs Open Addressing]]
│   │   ├── [[Robin Hood Hashing and Low Variance Lookup]]
│   │   ├── [[Cuckoo Hashing and Constant Worst-Case Lookup]]
│   │   ├── [[Consistent Hashing and Virtual Nodes]]
│   │   ├── [[Bloom Filter and Counting Bloom Filter]]
│   │   └── [[Cuckoo Filter and Quotient Filter]]
│   ├── [[Trees and Hierarchies|03. Trees and Hierarchies]]
│   │   ├── [[AVL Tree Strict Balance and Rotations]]
│   │   ├── [[Red-Black Tree Invariants and Colored Balancing]]
│   │   ├── [[Splay Tree and Amortized Self-Balancing]]
│   │   ├── [[Treap and Randomized Binary Search Tree]]
│   │   ├── [[B-Tree and B-Plus Tree on Disk Storage]]
│   │   ├── [[Segment Tree and Lazy Propagation]]
│   │   ├── [[Fenwick Tree (Binary Indexed Tree) Indexing Trick]]
│   │   ├── [[Trie, Radix Tree, and Compressed Prefix Tries]]
│   │   └── [[Hash Array Mapped Trie (HAMT)]]
│   ├── [[Heaps and Priority Queues|04. Heaps and Priority Queues]]
│   │   ├── [[Binary Heap Array Representation and Sift Operations]]
│   │   ├── [[D-Ary Heap Cache Tuning in Production]]
│   │   ├── [[Fibonacci Heap and O(1) Decrease Key]]
│   │   ├── [[Pairing Heap Practical Performance]]
│   │   └── [[Hierarchical Timer Wheel Scheduling]]
│   └── [[Disjoint Set and Graph Representations|05. Disjoint Set and Graph Representations]]
│   │   ├── [[Disjoint Set Union (DSU) with Path Compression and Rank]]
│   │   ├── [[Compressed Sparse Row (CSR) Graph Storage]]
│   │   └── [[Adjacency Matrix vs Adjacency List Tradeoffs]]
├── [[Sorting|03. Sorting]]
│   ├── [[Comparison-Based Sorts|01. Comparison-Based Sorts]]
│   │   ├── [[Quick Sort Dual-Pivot and Partitioning Schemes]]
│   │   ├── [[Merge Sort Stability and In-Place Variants]]
│   │   ├── [[Heap Sort In-Place Sorting and Cache Penalties]]
│   │   └── [[Information-Theoretic Sorting Lower Bound (N log N)]]
│   ├── [[Hybrid Production Sorts|02. Hybrid Production Sorts]]
│   │   ├── [[TimSort Run Finding and Galloping Mode]]
│   │   ├── [[IntroSort Depth Limiting and Fallback]]
│   │   └── [[Pattern-Defeating QuickSort (PDQSort)]]
│   └── [[Non-Comparison and External Sorts|03. Non-Comparison and External Sorts]]
│   │   ├── [[Counting Sort and Radix Sort (LSD vs MSD)]]
│   │   ├── [[Bucket Sort and Uniform Distribution Assumptions]]
│   │   └── [[External K-Way Merge Sort for Big Data]]
├── [[Searching|04. Searching]]
│   ├── [[Binary Search Paradigms|01. Binary Search Paradigms]]
│   │   ├── [[Binary Search Invariant and Boundary Conventions]]
│   │   ├── [[Binary Search on Monotonic Answer Space]]
│   │   └── [[Branchless and Cache-Aware Binary Search]]
│   ├── [[Two Pointers and Sliding Window|02. Two Pointers and Sliding Window]]
│   │   ├── [[Two Pointers Opposite Ends and Partitioning]]
│   │   ├── [[Fast and Slow Pointer Cycle Detection (Floyd)]]
│   │   ├── [[Sliding Window State and Shrink Invariants]]
│   │   └── [[Monotonic Queue for Sliding Window Extrema]]
│   └── [[Selection and Order Statistics|03. Selection and Order Statistics]]
│   │   ├── [[Quickselect and Expected O(N) Selection]]
│   │   ├── [[Median of Medians Worst-Case O(N) Selection]]
│   │   └── [[Order Statistics Tree and Rank Queries]]
├── [[Graph Algorithms|05. Graph Algorithms]]
│   ├── [[Traversals and Connectivity|01. Traversals and Connectivity]]
│   │   ├── [[Breadth-First Search (BFS) and Shortest Unweighted Paths]]
│   │   ├── [[Depth-First Search (DFS) and Tree Edge Classification]]
│   │   ├── [[Topological Sort (Kahn vs DFS Post-Order)]]
│   │   ├── [[Strongly Connected Components (Tarjan and Kosaraju)]]
│   │   └── [[Articulation Points and Bridges (Tarjan Low-Link)]]
│   ├── [[Shortest Path Algorithms|02. Shortest Path Algorithms]]
│   │   ├── [[Dijkstra Algorithm with Indexed Priority Queue]]
│   │   ├── [[Bellman-Ford and Negative Cycle Detection]]
│   │   ├── [[Floyd-Warshall All-Pairs Shortest Path]]
│   │   ├── [[Johnson Algorithm for Sparse All-Pairs Shortest Path]]
│   │   └── [[A-Star Heuristic Search and Admissible Heuristics]]
│   ├── [[Spanning Trees and Tree Decompositions|03. Spanning Trees and Tree Decompositions]]
│   │   ├── [[Kruskal Algorithm with Disjoint Set Union]]
│   │   ├── [[Prim Algorithm for Dense Graphs]]
│   │   ├── [[Lowest Common Ancestor (LCA) Binary Lifting]]
│   │   ├── [[Heavy-Light Decomposition (HLD)]]
│   │   └── [[Centroid Decomposition for Tree Divide and Conquer]]
│   └── [[Network Flow and Matching|04. Network Flow and Matching]]
│   │   ├── [[Max-Flow Min-Cut Theorem and Residual Graphs]]
│   │   ├── [[Edmonds-Karp and Dinic Max-Flow Algorithms]]
│   │   ├── [[Push-Relabel Algorithm and Highest Label Selection]]
│   │   └── [[Hopcroft-Karp Maximum Bipartite Matching]]
├── [[Dynamic Programming|06. Dynamic Programming]]
│   ├── [[Classical DP Patterns|01. Classical DP Patterns]]
│   │   ├── [[0-1 Knapsack and Unbounded Knapsack]]
│   │   ├── [[Longest Common Subsequence (LCS) and Diff Algorithms]]
│   │   ├── [[Longest Increasing Subsequence (LIS) in O(N log N)]]
│   │   ├── [[Edit Distance (Levenshtein Distance)]]
│   │   └── [[Matrix Chain Multiplication and Interval DP]]
│   ├── [[Advanced DP Paradigms|02. Advanced DP Paradigms]]
│   │   ├── [[Bitmask Dynamic Programming (TSP)]]
│   │   ├── [[Tree Dynamic Programming (Rerooting Technique)]]
│   │   ├── [[Digit Dynamic Programming for Numerical Ranges]]
│   │   └── [[Broken Profile (Plug DP) for Grid Tiling]]
│   └── [[DP Optimizations|03. DP Optimizations]]
│   │   ├── [[Convex Hull Trick and Li Chao Tree]]
│   │   ├── [[Divide and Conquer DP Optimization]]
│   │   ├── [[Knuth Optimization for Quadrangle Inequality]]
│   │   └── [[Monotonic Queue DP Optimization]]
├── [[Greedy|07. Greedy]]
│   ├── [[Greedy Foundations and Matroids|01. Greedy Foundations and Matroids]]
│   │   ├── [[Greedy Choice Property and Optimal Substructure]]
│   │   ├── [[Exchange Arguments Proof Technique]]
│   │   └── [[Matroid Theory and Greedy Correctness (Rado-Edmonds)]]
│   └── [[Classical Greedy Solutions|02. Classical Greedy Solutions]]
│   │   ├── [[Interval Scheduling and Minimum Meeting Rooms]]
│   │   ├── [[Huffman Optimal Prefix Code Construction]]
│   │   ├── [[Fractional Knapsack and Value-Density Greedy]]
│   │   └── [[Task Scheduler and CPU Cooldown Optimization]]
├── [[Backtracking|08. Backtracking]]
│   ├── [[Systematic Search and Pruning|01. Systematic Search and Pruning]]
│   │   ├── [[State Space Tree Exploration and Rollback Invariants]]
│   │   ├── [[Forward Checking and Constraint Pruning]]
│   │   └── [[Alpha-Beta Pruning for Minimax Game Trees]]
│   └── [[Exact Cover and Branch and Bound|02. Exact Cover and Branch and Bound]]
│   │   ├── [[Exact Cover and Knuth Dancing Links (Algorithm X)]]
│   │   ├── [[N-Queens and Bitmask Constraint Optimization]]
│   │   └── [[Branch and Bound for Travelling Salesperson (TSP)]]
└── [[String Algorithms|09. String Algorithms]]
│   ├── [[Exact Pattern Matching|01. Exact Pattern Matching]]
│   │   ├── [[Knuth-Morris-Pratt (KMP) and Prefix Function]]
│   │   ├── [[Boyer-Moore-Horspool Algorithm and Bad Character Rule]]
│   │   ├── [[Rabin-Karp Rolling Hash and Polynomial Fingerprinting]]
│   │   ├── [[Aho-Corasick Multi-Pattern Automaton]]
│   │   └── [[Z-Algorithm and Z-Array Construction]]
│   ├── [[Suffix Structures and Text Indexing|02. Suffix Structures and Text Indexing]]
│   │   ├── [[Suffix Automaton (SAM) Directed Acyclic Word Graph]]
│   │   ├── [[Suffix Array and Longest Common Prefix (LCP) Array]]
│   │   ├── [[Suffix Tree Construction (Ukkonen Algorithm)]]
│   │   └── [[FM-Index and Burrows-Wheeler Transform (BWT)]]
│   └── [[Palindromes and Sequence Algorithms|03. Palindromes and Sequence Algorithms]]
│   │   ├── [[Manachers Algorithm for Longest Palindromic Substring]]
│   │   └── [[Palindromic Tree (Eertree) Automaton]]
```

---

## 🏛️ Core Knowledge Domains

### 1. 📂 [[Complexity Analysis|01. Complexity Analysis]]
#### 1. 📂 [[Asymptotic Bounds|01. Asymptotic Bounds]]
- [[Big-O Notation and Upper Bounds]] — Formal definition of asymptotic upper bounds and tight worst-case bounding.
- [[Big-Omega and Lower Bounds]] — Asymptotic lower bounds and information-theoretic sorting/searching limits.
- [[Big-Theta and Tight Bounds]] — Tight asymptotic bounding where upper and lower bounds asymptotically converge.
- [[Little-o and Little-Omega Strict Bounds]] — Strict non-tight asymptotic dominance relations and mathematical limits.
- [[Dominance Relations and Complexity Classes]] — Hierarchy of complexity: O(1), O(log n), O(n), O(n log n), O(n^2), O(2^n), O(n!).
#### 2. 📂 [[Recurrences and Amortization|02. Recurrences and Amortization]]
- [[Master Theorem for Divide and Conquer]] — Solving recurrences of the form T(n) = aT(n/b) + f(n) across all three cases.
- [[Akra-Bazzi Method for Non-Uniform Recurrences]] — Generalized recurrence solving for non-integer branch factors and unequal splits.
- [[Accounting and Aggregate Amortized Analysis]] — Charging artificial costs to operations to prove bounded average sequence latency.
- [[Potential Method and Energy Functions]] — Physically inspired amortized analysis using potential energy functions Phi(D_i).
- [[Dynamic Array Resizing Amortization]] — Proving O(1) amortized append operations under geometric expansion factors (2x vs 1.5x).
#### 3. 📂 [[Hardware-Aware Complexity|03. Hardware-Aware Complexity]]
- [[External Memory Model and IO Complexity]] — Aggarwal-Vitter model measuring disk page and cache line block transfers (B and M).
- [[Cache-Oblivious Algorithm Design]] — Optimal cache usage across all memory hierarchy levels without hardware tuning parameters.
- [[Branch Prediction Penalty in Branchless Code]] — Designing branch-free arithmetic to eliminate CPU pipeline stalls on random data.
- [[Memory Allocation Overhead and Fragmentation Cost]] — Analyzing hidden allocator overhead and cache line pollution in pointer-heavy structures.
### 2. 📂 [[Data Structures|02. Data Structures]]
#### 1. 📂 [[Linear Structures|01. Linear Structures]]
- [[Dynamic Array Contiguity and Growth Policy]] — Contiguous memory allocation, geometric reallocation factors, and cache prefetching.
- [[Singly and Doubly Linked List Pointer Layouts]] — Node pointer chasing, intrusive list designs (Linux kernel list_head), and cache misses.
- [[Unrolled Linked List and Memory Density]] — Combining arrays within list nodes to maximize cache line utilization.
- [[Circular Ring Buffer and Lock-Free Queues]] — Fixed-size power-of-two ring buffers with atomic head/tail masking for high-throughput IPC.
- [[Monotonic Stack and Monotonic Deque]] — Maintaining sorted invariant stacks to solve Next Greater Element and Sliding Window Maximum in O(N).
#### 2. 📂 [[Hash-Based Structures|02. Hash-Based Structures]]
- [[Hash Table Chaining vs Open Addressing]] — Separate chaining bucket arrays vs linear/quadratic probing and cache locality.
- [[Robin Hood Hashing and Low Variance Lookup]] — Stealing from the rich to equalize probe sequence lengths and bound P99 search latency.
- [[Cuckoo Hashing and Constant Worst-Case Lookup]] — Multiple hash functions and displacement chains ensuring true O(1) worst-case lookups.
- [[Consistent Hashing and Virtual Nodes]] — Ring-based key distribution for distributed caching, dynamodb-style partitioning, and zero-churn additions.
- [[Bloom Filter and Counting Bloom Filter]] — Probabilistic set membership testing with zero false negatives and tunable false positive bitsets.
- [[Cuckoo Filter and Quotient Filter]] — Dynamic deletion-supporting probabilistic filters with superior space efficiency over Bloom filters.
#### 3. 📂 [[Trees and Hierarchies|03. Trees and Hierarchies]]
- [[AVL Tree Strict Balance and Rotations]] — Rigid balance factor (|h_L - h_R| <= 1) providing optimal read-heavy search performance.
- [[Red-Black Tree Invariants and Colored Balancing]] — Relaxed height bound (2 log n) optimizing insertion and deletion mutation throughput.
- [[Splay Tree and Amortized Self-Balancing]] — Access-driven tree rotation (splaying) providing O(log n) amortized cost for non-uniform access.
- [[Treap and Randomized Binary Search Tree]] — Combining BST keys with randomized max-heap priorities to maintain balance without rotations.
- [[B-Tree and B-Plus Tree on Disk Storage]] — High fan-out multi-way search trees optimized for block storage and database index pages.
- [[Segment Tree and Lazy Propagation]] — O(log N) range queries and range updates using deferred tree node modifications.
- [[Fenwick Tree (Binary Indexed Tree) Indexing Trick]] — Prefix sum queries and point updates using two s complement bitwise operations (i & -i).
- [[Trie, Radix Tree, and Compressed Prefix Tries]] — Prefix-based retrieval structures, edge-compacted Patricia tries, and longest prefix matching.
- [[Hash Array Mapped Trie (HAMT)]] — Persistent, immutable, high-fan-out (32-way) trie structures used in modern functional languages.
#### 4. 📂 [[Heaps and Priority Queues|04. Heaps and Priority Queues]]
- [[Binary Heap Array Representation and Sift Operations]] — Implicit 1-indexed array representation (parent = i/2, children = 2i, 2i+1) and O(N) heapify.
- [[D-Ary Heap Cache Tuning in Production]] — Optimizing heap fan-out (D=4 or D=8) to match hardware CPU cache line transfers.
- [[Fibonacci Heap and O(1) Decrease Key]] — Lazy tree consolidation enabling theoretical O(1) amortized Decrease-Key for Dijkstra.
- [[Pairing Heap Practical Performance]] — Simpler self-adjusting heap variant outperforming Fibonacci heaps in real-world implementations.
- [[Hierarchical Timer Wheel Scheduling]] — O(1) insert and tick timer scheduler used in Linux kernel and Netty network engines.
#### 5. 📂 [[Disjoint Set and Graph Representations|05. Disjoint Set and Graph Representations]]
- [[Disjoint Set Union (DSU) with Path Compression and Rank]] — Near-constant O(alpha(N)) inverse Ackermann set union and connectivity queries.
- [[Compressed Sparse Row (CSR) Graph Storage]] — Memory-dense contiguous array graph representation for high-performance graph processing and GPU algorithms.
- [[Adjacency Matrix vs Adjacency List Tradeoffs]] — Memory footprint vs neighbor traversal speed and edge existence verification trade-offs.
### 3. 📂 [[Sorting|03. Sorting]]
#### 1. 📂 [[Comparison-Based Sorts|01. Comparison-Based Sorts]]
- [[Quick Sort Dual-Pivot and Partitioning Schemes]] — Lomuto, Hoare, and Yaroslavskiy dual-pivot partitioning with randomized pivot selection.
- [[Merge Sort Stability and In-Place Variants]] — Stable divide-and-conquer sorting, linked list sorting, and memory trade-offs.
- [[Heap Sort In-Place Sorting and Cache Penalties]] — Guaranteed O(N log N) worst-case time with zero auxiliary memory but poor cache locality.
- [[Information-Theoretic Sorting Lower Bound (N log N)]] — Decision tree proof establishing Omega(N log N) comparisons for comparison-based sorting.
#### 2. 📂 [[Hybrid Production Sorts|02. Hybrid Production Sorts]]
- [[TimSort Run Finding and Galloping Mode]] — Adaptive stable hybrid of Merge Sort and Insertion Sort optimizing real-world ordered data (Python/Java).
- [[IntroSort Depth Limiting and Fallback]] — Hybrid Quick Sort falling back to Heap Sort upon recursion depth limits to guarantee O(N log N) (C++ std::sort).
- [[Pattern-Defeating QuickSort (PDQSort)]] — Branchless partitioning, fallback to Heap Sort, and detecting pattern repetitions (Go 1.19+ sort).
#### 3. 📂 [[Non-Comparison and External Sorts|03. Non-Comparison and External Sorts]]
- [[Counting Sort and Radix Sort (LSD vs MSD)]] — Linear O(N + K) sorting by digit keys, stability preservation, and cache considerations.
- [[Bucket Sort and Uniform Distribution Assumptions]] — Scatter-gather sorting under uniform probability distributions and variance bounding.
- [[External K-Way Merge Sort for Big Data]] — Sorting terabytes of data exceeding RAM using disk runs, tournament winner trees, and polyphase merging.
### 4. 📂 [[Searching|04. Searching]]
#### 1. 📂 [[Binary Search Paradigms|01. Binary Search Paradigms]]
- [[Binary Search Invariant and Boundary Conventions]] — Half-open [low, high) vs closed [low, high] interval invariants and avoiding infinite loops.
- [[Binary Search on Monotonic Answer Space]] — Transforming optimization problems (min/max) into decision problems with monotonic boolean predicates.
- [[Branchless and Cache-Aware Binary Search]] — Eliminating branch prediction failures using conditional moves (CMOV) and Eytzinger memory layout.
#### 2. 📂 [[Two Pointers and Sliding Window|02. Two Pointers and Sliding Window]]
- [[Two Pointers Opposite Ends and Partitioning]] — Two-pointer convergence for sorted arrays, Dutch National Flag, and container water problems.
- [[Fast and Slow Pointer Cycle Detection (Floyd)]] — Tortoise and Hare algorithm for linked list cycle detection, entry point discovery, and midpoints.
- [[Sliding Window State and Shrink Invariants]] — Dynamic window expansion and contraction maintaining valid frequency counters in O(N) time.
- [[Monotonic Queue for Sliding Window Extrema]] — O(1) amortized window min/max tracking using double-ended monotonic queues.
#### 3. 📂 [[Selection and Order Statistics|03. Selection and Order Statistics]]
- [[Quickselect and Expected O(N) Selection]] — Selecting the K-th smallest element via partition recursion without full array sorting.
- [[Median of Medians Worst-Case O(N) Selection]] — Blum-Floyd-Pratt-Rivest-Tarjan algorithm guaranteeing balanced 70-30 partitions.
- [[Order Statistics Tree and Rank Queries]] — Augmenting balanced BSTs with subtree size counters to perform O(log N) Find-by-Rank.
### 5. 📂 [[Graph Algorithms|05. Graph Algorithms]]
#### 1. 📂 [[Traversals and Connectivity|01. Traversals and Connectivity]]
- [[Breadth-First Search (BFS) and Shortest Unweighted Paths]] — Queue-based level-order traversal, 0-1 BFS with deques, and multi-source BFS.
- [[Depth-First Search (DFS) and Tree Edge Classification]] — Recursive and iterative DFS classifying tree, back, forward, and cross edges.
- [[Topological Sort (Kahn vs DFS Post-Order)]] — Linear ordering of DAG vertices using in-degree queues and reverse post-order DFS.
- [[Strongly Connected Components (Tarjan and Kosaraju)]] — Decomposing directed graphs into maximal strongly connected subgraphs using low-link values.
- [[Articulation Points and Bridges (Tarjan Low-Link)]] — Identifying single points of failure in undirected networks using DFS entry/low times.
#### 2. 📂 [[Shortest Path Algorithms|02. Shortest Path Algorithms]]
- [[Dijkstra Algorithm with Indexed Priority Queue]] — Greedy single-source shortest paths on non-negative weighted graphs in O((V + E) log V).
- [[Bellman-Ford and Negative Cycle Detection]] — Dynamic programming edge relaxation over V-1 iterations and detecting negative arbitrage cycles.
- [[Floyd-Warshall All-Pairs Shortest Path]] — O(V^3) matrix dynamic programming across intermediate vertices k for dense networks.
- [[Johnson Algorithm for Sparse All-Pairs Shortest Path]] — Reweighting graph edges via Bellman-Ford potentials to run V Dijkstra searches in O(V E log V).
- [[A-Star Heuristic Search and Admissible Heuristics]] — Goal-directed shortest path search using Euclidean/Manhattan distance heuristic estimates.
#### 3. 📂 [[Spanning Trees and Tree Decompositions|03. Spanning Trees and Tree Decompositions]]
- [[Kruskal Algorithm with Disjoint Set Union]] — Greedy edge sorting and cycle prevention for Minimum Spanning Trees in O(E log E).
- [[Prim Algorithm for Dense Graphs]] — Vertex-growing MST construction in O(V^2) or O(E + V log V) with Fibonacci heaps.
- [[Lowest Common Ancestor (LCA) Binary Lifting]] — O(N log N) precomputation and O(log N) ancestor queries using 2^k ancestor tables.
- [[Heavy-Light Decomposition (HLD)]] — Decomposing tree paths into heavy chains to support O(log^2 N) range queries with Segment Trees.
- [[Centroid Decomposition for Tree Divide and Conquer]] — Recursively splitting trees at centroid balance nodes to answer all-path distance queries in O(N log N).
#### 4. 📂 [[Network Flow and Matching|04. Network Flow and Matching]]
- [[Max-Flow Min-Cut Theorem and Residual Graphs]] — Duality between maximum network flow and minimum capacity edge cuts in directed graphs.
- [[Edmonds-Karp and Dinic Max-Flow Algorithms]] — Shortest augmenting paths with BFS and Dinic s layered networks with blocking flows in O(V^2 E).
- [[Push-Relabel Algorithm and Highest Label Selection]] — Preflow-push mechanics operating on local vertex heights achieving O(V^2 sqrt(E)).
- [[Hopcroft-Karp Maximum Bipartite Matching]] — Finding maximum matching in bipartite graphs in O(E sqrt(V)) using layered alternating BFS/DFS.
### 6. 📂 [[Dynamic Programming|06. Dynamic Programming]]
#### 1. 📂 [[Classical DP Patterns|01. Classical DP Patterns]]
- [[0-1 Knapsack and Unbounded Knapsack]] — State transitions over capacity constraints and space optimization to 1D arrays.
- [[Longest Common Subsequence (LCS) and Diff Algorithms]] — 2D grid transitions computing minimal edit scripts and file diff comparisons.
- [[Longest Increasing Subsequence (LIS) in O(N log N)]] — Patience sorting and binary search state optimization over greedy tails.
- [[Edit Distance (Levenshtein Distance)]] — Dynamic programming table for string insertions, deletions, and substitutions.
- [[Matrix Chain Multiplication and Interval DP]] — Optimal parenthesization of associative matrix products via subinterval partitioning.
#### 2. 📂 [[Advanced DP Paradigms|02. Advanced DP Paradigms]]
- [[Bitmask Dynamic Programming (TSP)]] — Representing subsets of visited vertices as integer bitmasks in O(2^N * N^2) time.
- [[Tree Dynamic Programming (Rerooting Technique)]] — Computing tree properties for all possible roots in O(N) time via two-pass DFS.
- [[Digit Dynamic Programming for Numerical Ranges]] — Counting integers in [A, B] satisfying digit properties using tight/leading-zero states.
- [[Broken Profile (Plug DP) for Grid Tiling]] — Tracking frontier boundary contours to solve exact domino and polyomino grid tilings.
#### 3. 📂 [[DP Optimizations|03. DP Optimizations]]
- [[Convex Hull Trick and Li Chao Tree]] — Optimizing DP transitions of the form DP[i] = min(DP[j] + m_j * x_i + c_j) from O(N^2) to O(N log N).
- [[Divide and Conquer DP Optimization]] — Accelerating 2D DP when the optimal transition index opt[i][j] is monotonic in O(K N log N).
- [[Knuth Optimization for Quadrangle Inequality]] — Reducing interval DP from O(N^3) to O(N^2) when cost functions satisfy the quadrangle inequality.
- [[Monotonic Queue DP Optimization]] — Sliding window maximum optimization reducing transitions over sliding index bounds to O(N).
### 7. 📂 [[Greedy|07. Greedy]]
#### 1. 📂 [[Greedy Foundations and Matroids|01. Greedy Foundations and Matroids]]
- [[Greedy Choice Property and Optimal Substructure]] — Proving local optimal choices lead to global optimum via mathematical induction and contradiction.
- [[Exchange Arguments Proof Technique]] — Proving greedy optimality by showing any non-greedy solution can be transformed into the greedy solution without loss.
- [[Matroid Theory and Greedy Correctness (Rado-Edmonds)]] — Algebraic structures (hereditary system + exchange property) where greedy algorithms are guaranteed to find maximum weight independent sets.
#### 2. 📂 [[Classical Greedy Solutions|02. Classical Greedy Solutions]]
- [[Interval Scheduling and Minimum Meeting Rooms]] — Sorting by end time for non-overlapping selection vs start-time sorting with min-heaps for room allocation.
- [[Huffman Optimal Prefix Code Construction]] — Building optimal prefix-free variable length binary trees using priority queues in O(N log N).
- [[Fractional Knapsack and Value-Density Greedy]] — Sorting continuous items by value-to-weight ratio for optimal greedy packing.
- [[Task Scheduler and CPU Cooldown Optimization]] — Frequency-based greedy scheduling with idle slot minimization.
### 8. 📂 [[Backtracking|08. Backtracking]]
#### 1. 📂 [[Systematic Search and Pruning|01. Systematic Search and Pruning]]
- [[State Space Tree Exploration and Rollback Invariants]] — Recursive DFS depth exploration, mutating state before recursive calls, and restoring invariants upon backtrack.
- [[Forward Checking and Constraint Pruning]] — Eliminating invalid future search branches early using constraint lookahead.
- [[Alpha-Beta Pruning for Minimax Game Trees]] — Pruning branches in adversarial two-player zero-sum games without altering evaluation scores.
#### 2. 📂 [[Exact Cover and Branch and Bound|02. Exact Cover and Branch and Bound]]
- [[Exact Cover and Knuth Dancing Links (Algorithm X)]] — Circular doubly linked 2D toroidal matrices achieving ultra-fast backtrack row cover/uncover operations.
- [[N-Queens and Bitmask Constraint Optimization]] — Solving N-Queens using column, left-diagonal, and right-diagonal bitmasks in O(1) bitwise operations.
- [[Branch and Bound for Travelling Salesperson (TSP)]] — Bounding search space using Minimum Spanning Tree relaxation and lower bound cost matrices.
### 9. 📂 [[String Algorithms|09. String Algorithms]]
#### 1. 📂 [[Exact Pattern Matching|01. Exact Pattern Matching]]
- [[Knuth-Morris-Pratt (KMP) and Prefix Function]] — O(N + M) pattern matching using Longest Prefix Suffix (LPS) lookup table to prevent backtracking.
- [[Boyer-Moore-Horspool Algorithm and Bad Character Rule]] — Sublinear average time pattern searching scanning pattern right-to-left and jumping characters.
- [[Rabin-Karp Rolling Hash and Polynomial Fingerprinting]] — O(N) search using rolling polynomial modular hashing with prime base arithmetic.
- [[Aho-Corasick Multi-Pattern Automaton]] — Trie augmented with failure and dictionary links matching thousands of keywords simultaneously in O(N + M).
- [[Z-Algorithm and Z-Array Construction]] — O(N) prefix-matching array computing longest common prefix with string head for every position.
#### 2. 📂 [[Suffix Structures and Text Indexing|02. Suffix Structures and Text Indexing]]
- [[Suffix Automaton (SAM) Directed Acyclic Word Graph]] — Minimal deterministic automaton recognizing all suffixes of a string with O(N) states and transitions.
- [[Suffix Array and Longest Common Prefix (LCP) Array]] — Sorted lexicographical suffix indices in O(N log N) or O(N) using SA-IS and Kasai s LCP algorithm.
- [[Suffix Tree Construction (Ukkonen Algorithm)]] — Linear-time O(N) on-line suffix tree construction using implicit suffix extensions and skip/count tricks.
- [[FM-Index and Burrows-Wheeler Transform (BWT)]] — Compressed text index enabling exact substring search in O(P) time within entropy-compressed memory space.
#### 3. 📂 [[Palindromes and Sequence Algorithms|03. Palindromes and Sequence Algorithms]]
- [[Manachers Algorithm for Longest Palindromic Substring]] — O(N) linear time palindrome discovery exploiting symmetry around expanded delimiter centers.
- [[Palindromic Tree (Eertree) Automaton]] — Directed tree structure storing all distinct palindromic substrings with odd/even imaginary root nodes.

---

## 🔗 Navigation
- ⬆️ Parent: [[Computer Science]]
- 🎓 Root: [[Principal SWE]]
