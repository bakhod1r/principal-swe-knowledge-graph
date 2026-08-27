---
title: "Backtracking Concept"
tags:
  - review
  - computer-science
  - algorithms
  - dsa
  - backtracking
  - principal-swe
parent: "[[Fundamentals]]"
---

# Backtracking Concept

## 1. What Is Backtracking?

**Backtracking** is an algorithmic technique for solving problems by:

> **Building a solution incrementally, and whenever the current partial solution cannot lead to a valid complete solution, undoing the last choice and trying another one.**

The core idea is:

```text
Choose
  ↓
Explore
  ↓
Valid?
 ├── Yes → Continue
 └── No  → Undo choice → Try another choice
```

A useful mental model is:

> **Backtracking = DFS over a decision tree + pruning + undo**

It is especially useful when the problem asks you to **enumerate or find configurations** satisfying a set of constraints.

---

# 2. The Core Mental Model

Imagine you need to construct a solution one decision at a time.

For example, suppose we want all permutations of:

```text
[1, 2, 3]
```

At the first position we have three choices:

```text
              []
          /    |    \
        [1]   [2]   [3]
       /  \   / \   / \
   [1,2][1,3]...
```

This is a **decision tree**.

Backtracking performs a DFS traversal of this tree.

When we reach:

```text
[1, 2, 3]
```

we have a complete solution.

Then we **undo** the last decision:

```text
[1, 2, 3]
       ↑
     undo 3

[1, 2]
```

and try another choice.

---

# 3. Why Does Backtracking Exist?

Many combinatorial problems have an enormous number of possible solutions.

For example:

### Permutations

For `n` elements:

  n!  

possibilities.

For `n = 10`:

 10! = 3,628,800 

For `n = 20`:

20! > 2.43 times 10^18 

We cannot blindly enumerate every possibility efficiently.

Backtracking helps because we can **stop exploring branches that are already impossible**.

This is called **pruning**.

---

# 4. Backtracking vs Brute Force

The distinction is important.

### Pure brute force

Generate every possible candidate:

```text
generate candidate
check candidate
generate candidate
check candidate
...
```

### Backtracking

Build candidates incrementally:

```text
make decision
check whether partial solution is still possible

if impossible:
    stop this branch

otherwise:
    continue
```

So:

> **Backtracking is essentially structured brute force with early pruning.**

It does **not magically eliminate exponential complexity**.

It tries to avoid exploring unnecessary parts of the search space.

---

# 5. The Three Fundamental Components

Almost every backtracking algorithm has three parts.

## 1. Choice

What decision can I make next?

```text
for each candidate:
```

## 2. Constraint

Can I make this choice?

```text
if violatesConstraint(candidate):
    continue
```

## 3. Undo

After exploring the choice, restore the previous state.

```text
choose(candidate)

backtrack()

undo(candidate)
```

This gives the canonical structure:

```text
func backtrack(state):
    if solution(state):
        record(state)
        return

    for choice in choices(state):
        if invalid(choice, state):
            continue

        make(choice, state)

        backtrack(state)

        undo(choice, state)
```

This pattern is worth memorizing conceptually, not syntactically.

---

# 6. Simple Example — Generate Binary Strings

Suppose we want all binary strings of length `3`.

Expected:

```text
000
001
010
011
100
101
110
111
```

At every position we have two choices:

```text
0
1
```

Decision tree:

```text
                 ""
              /      \
             0        1
           /  \      /  \
         00   01    10   11
        / \   / \   / \   / \
      000 001 ...
```

The algorithm:

```go
func generate(path []byte, n int) {
    if len(path) == n {
        fmt.Println(string(path))
        return
    }

    for _, choice := range []byte{'0', '1'} {
        path = append(path, choice)

        generate(path, n)

        path = path[:len(path)-1]
    }
}
```

The critical operation is:

```go
path = path[:len(path)-1]
```

That is the **backtrack / undo** operation.

---

# 7. A More Interesting Example — N-Queens

The classic backtracking problem is **N-Queens**.

Problem:

> Place `N` queens on an `N × N` chessboard so that no two queens attack each other.

For example, for `N = 4`:

```text
.Q..
...Q
Q...
..Q.
```

A queen attacks:

```text
← → horizontal
↑ ↓ vertical
↗ ↘ diagonals
```

Instead of placing queens randomly, we can make one decision per row:

```text
Row 0 → choose column
Row 1 → choose column
Row 2 → choose column
...
```

The search tree becomes:

```text
row 0
 ├── col 0
 │    ├── row 1 col 0 ❌
 │    ├── row 1 col 1 ❌
 │    └── row 1 col 2 ✓
 │          ...
 ├── col 1
 │    ...
 └── col 2
```

If a queen placement causes a conflict, we immediately stop that branch.

That is **pruning**.

---

# 8. The Most Important Insight: State

A good backtracking solution starts by defining:

> **What information completely describes my current partial solution?**

For N-Queens:

```text
row
occupied columns
occupied diagonals
```

For permutations:

```text
current path
used elements
```

For subset generation:

```text
current index
current subset
```

For Sudoku:

```text
board state
```

This is one of the most important algorithm-design skills:

> **Choose a state representation that makes constraint checking cheap.**

---

# 9. Backtracking Template

A generic implementation looks like:

```go
func backtrack(state State) {
    if isComplete(state) {
        save(state)
        return
    }

    for _, choice := range getChoices(state) {
        if !isValid(state, choice) {
            continue
        }

        apply(state, choice)

        backtrack(state)

        undo(state, choice)
    }
}
```

Think of it as:

```text
             State
               │
          generate choices
               │
       ┌───────┼───────┐
       ↓       ↓       ↓
    Choice A Choice B Choice C
       │       │       │
     valid?  valid?  valid?
       │
     apply
       │
     recurse
       │
     undo
```

---

# 10. The `undo` Operation Is Fundamental

Consider:

```go
path = append(path, choice)

backtrack(path)

path = path[:len(path)-1]
```

Why do we need the undo?

Because the same state object is reused across different branches.

Suppose:

```text
path = [1]
```

We choose:

```text
2
```

Now:

```text
path = [1,2]
```

After exploring that branch, we need:

```text
path = [1]
```

before trying:

```text
3
```

Otherwise the next branch would incorrectly see:

```text
[1,2,3]
```

instead of:

```text
[1,3]
```

So:

> **Backtracking = mutation + recursion + restoration**

This is one of the most important implementation invariants.

---

# 11. Backtracking and DFS

Backtracking is closely related to DFS.

### DFS

DFS says:

> Explore one branch deeply before exploring another.

### Backtracking

Backtracking adds:

> While exploring, maintain constraints and undo decisions when returning from recursion.

Therefore:

```text
Backtracking
    =
DFS
+
Decision making
+
Constraint checking
+
State restoration
+
Pruning
```

Not every DFS is backtracking.

But many backtracking algorithms use DFS.

---

# 12. Pruning

Pruning is where backtracking becomes significantly more powerful.

Suppose we are solving:

```text
Subset Sum
```

Given:

```text
[2, 4, 7, 10]
```

and target:

```text
8
```

If the current sum is already:

```text
11
```

then:

```text
11 > 8
```

and assuming all numbers are non-negative:

```text
this branch can never produce 8
```

So we immediately return.

```go
if sum > target {
    return
}
```

Instead of:

```text
explore everything
```

we get:

```text
explore only potentially valid branches
```

---

# 13. Constraint Checking

A backtracking algorithm typically has a predicate:

```text
Is this partial solution still potentially valid?
```

Examples:

### N-Queens

```text
Is this column occupied?
Is this diagonal occupied?
```

### Sudoku

```text
Is this number already present in row?
Is it present in column?
Is it present in box?
```

### Graph coloring

```text
Does this color conflict with neighboring vertices?
```

### Combination Sum

```text
Has the current sum exceeded target?
```

This constraint function is often where most of the algorithm's performance comes from.

---

# 14. Backtracking vs Dynamic Programming

These are frequently confused.

### Backtracking

Explores:

```text
different decisions
```

and usually searches for:

```text
all solutions
or one valid configuration
```

### Dynamic Programming

Exploits:

```text
overlapping subproblems
+
optimal substructure
```

and tries to avoid solving the same subproblem repeatedly.

For example:

```text
Backtracking:

        A
       / \
      B   C
     / \
    D   E
```

If multiple branches reach the same logical state:

```text
B
```

DP may cache its result.

Backtracking generally does not.

Therefore:

> If the search tree contains many repeated equivalent states, consider **memoization / DP**.

---

# 15. Backtracking vs Greedy

Greedy algorithms make a locally optimal choice:

```text
choose best-looking option now
never reconsider
```

Backtracking does:

```text
make a choice
explore
if it fails:
    undo
    try another
```

So the key difference is:

> **Greedy commits. Backtracking is willing to reconsider.**

Example:

```text
Greedy:
A → B → C → done

Backtracking:
A → B → C → fail
        ↓
      undo
        ↓
A → B → D
```

---

# 16. Complexity

Backtracking is often exponential.

If every state has `b` choices and maximum depth is `d`:

 O(b^d) 

For example, generating all subsets:

```text
Each element:
    include
    exclude
```

Therefore:

 O(2^n)  

Permutations:

O(n!)  

N-Queens has a large combinatorial search space and is commonly described with exponential/factorial-style upper bounds depending on the formulation and pruning.

The important point is:

> **Backtracking does not guarantee polynomial time.**

Its practical performance depends heavily on:

- branching factor
    
- depth
    
- constraint strength
    
- ordering of choices
    
- pruning effectiveness
    
- state representation
    

---

# 17. Choice Ordering Matters

This is a more advanced insight.

Suppose two branches exist:

```text
A
B
```

If branch `A` is likely to fail quickly:

```text
A → fail
```

then exploring it first can be beneficial.

Why?

Because we discover contradictions earlier.

This is especially useful when searching for **one solution**.

For example:

```text
choose most constrained variable first
```

is a powerful heuristic.

This idea appears in:

- Sudoku solvers
    
- SAT solvers
    
- constraint programming
    
- graph coloring
    
- scheduling
    

---

# 18. A Useful Principle

When solving a backtracking problem, ask:

> **Can I detect failure before reaching a complete solution?**

If yes:

```text
prune
```

The earlier the failure is detected, the smaller the effective search tree.

Conceptually:

```text
Without pruning:

              root
        /      |      \
       ...    ...     ...
      / | \  / | \   / | \
     huge search tree
```

With strong pruning:

```text
              root
        /      |      \
       ✓       ❌       ❌
      / \
     ✓   ❌
     |
   solution
```

---

# 19. Common Backtracking Problems

You should recognize these patterns immediately.

### Permutations

```text
[1,2,3]
```

Generate:

```text
123
132
213
231
312
321
```

Pattern:

```text
choose unused element
→ recurse
→ unchoose
```

---

### Combinations

Choose `k` elements from `n`.

Pattern:

```text
choose next element
→ recurse from i+1
→ undo
```

---

### Subsets

For every element:

```text
include
exclude
```

---

### N-Queens

```text
choose column for each row
→ validate
→ recurse
→ remove queen
```

---

### Sudoku

```text
choose empty cell
→ try 1..9
→ validate
→ recurse
→ undo
```

---

### Maze / Grid Search

```text
move
→ mark visited
→ recurse
→ unmark
```

---

### Graph Coloring

```text
assign color
→ check neighbors
→ recurse
→ remove color
```

---

# 20. A Powerful Generic Pattern

You can reduce many problems to:

```text
function backtrack(state):

    if solution(state):
        record solution
        return

    choices = generateChoices(state)

    for choice in choices:

        if violatesConstraints(state, choice):
            continue

        apply(choice)

        backtrack(state)

        undo(choice)
```

When facing a new problem, fill in these five questions:

```text
1. What is my state?
2. What are my choices?
3. What makes a choice invalid?
4. When is the solution complete?
5. How do I undo the choice?
```

If you can answer those clearly, you are usually close to the solution.

---

# 21. Common Bugs

## Bug 1 — Forgetting `undo`

```go
apply(choice)
backtrack()
```

without:

```go
undo(choice)
```

causes state contamination between branches.

---

## Bug 2 — Checking constraints too late

Bad:

```text
build entire candidate
then check validity
```

Better:

```text
check validity after every decision
```

---

## Bug 3 — Incorrect state sharing

In Go, slices are descriptors over an underlying array.

So code involving:

```go
append(path, x)
```

requires careful reasoning about whether branches can observe mutations to the same backing array.

A clean pattern is often:

```go
path = append(path, choice)

backtrack(path)

path = path[:len(path)-1]
```

provided the recursive function does not retain the slice incorrectly.

---

## Bug 4 — Copying too much

You could create a new state for every recursive call:

```go
newState := copy(state)
```

This simplifies ownership reasoning but can create substantial allocation and copying overhead.

Alternatively:

```text
mutate
recurse
undo
```

is usually faster.

But it requires a strict invariant:

> **Every mutation must be completely reversed before returning.**

---

# 22. Production-Level Mental Model

For algorithmic problems, think of backtracking as a **search engine**.

It has:

```text
State
   ↓
Decision
   ↓
Constraint
   ↓
Transition
   ↓
Recursive search
   ↓
Rollback
```

The key engineering problem is not recursion itself.

It is:

> **How cheaply can I represent state, generate choices, detect invalid states, and restore state?**

That determines practical performance.

---

# 23. Backtracking Optimization Hierarchy

When performance is poor, optimize in this order:

### 1. Better pruning

Largest potential win.

```text
Can I prove this branch cannot work earlier?
```

### 2. Better state representation

Replace expensive scans with:

```text
set
bitset
boolean array
counter
```

### 3. Better choice ordering

Try likely-to-fail or likely-to-succeed choices first depending on the goal.

### 4. Avoid unnecessary copying

Prefer controlled mutation + undo where appropriate.

### 5. Memoization

If the same logical state appears repeatedly:

```text
state → cached result
```

may transform the problem substantially.

### 6. More advanced algorithms

Depending on the problem:

```text
DP
Branch and Bound
Constraint Programming
SAT
ILP
Graph algorithms
```

may be superior.

---

# 24. Backtracking vs Branch and Bound

They are related but not identical.

### Backtracking

Usually prunes because:

```text
current state violates constraints
```

### Branch and Bound

Prunes because:

```text
even the best possible continuation cannot beat the current best solution
```

Example optimization problem:

```text
Current best = 100

This branch can achieve at most 80

→ discard branch
```

So:

```text
Backtracking:
    "This branch cannot produce a valid solution."

Branch and Bound:
    "This branch cannot produce a better solution."
```

---

# 25. The Principal-Level Insight

The most important thing to learn is not the recursive template.

It is **search-space engineering**.

Suppose the theoretical search space is:

10^{15}  

You cannot make that fast simply by writing cleaner recursion.

You need to reduce the **effective search space**.

That comes from:

```text
better constraints
      ↓
better pruning
      ↓
better state representation
      ↓
better ordering
      ↓
memoization / symmetry breaking
      ↓
possibly a fundamentally different algorithm
```

This is how experienced algorithm engineers think.

---

# 26. A Practical Decision Framework

When you encounter a problem, ask:

```text
Does the problem involve choosing among alternatives?
        │
        ├── No → probably not backtracking
        │
        └── Yes
             │
             ↓
Can choices be made incrementally?
             │
             ├── No → consider another technique
             │
             └── Yes
                  │
                  ↓
Can invalid partial solutions be detected early?
                  │
                  ├── No → brute force may be necessary
                  │
                  └── Yes
                       │
                       ↓
                 Backtracking
```

Then ask:

```text
Can equivalent states repeat?
        ↓
     Memoization / DP

Is this an optimization problem?
        ↓
     Branch and Bound

Is there strong mathematical structure?
        ↓
     Maybe a specialized algorithm is better
```

---

# 27. One Sentence to Remember

> **Backtracking systematically explores a decision tree, immediately abandons impossible partial solutions, and restores the previous state before trying the next choice.**

The canonical mental equation is:


`Backtracking  = DFS + Constraints + Pruning + Undo`


And the most important implementation invariant is:

`apply(choice) -> recurse -> undo(choice)`


That invariant is the foundation for solving **permutations, combinations, subsets, N-Queens, Sudoku, graph coloring, maze problems, constraint satisfaction**, and many other combinatorial search problems.

---

## 🔗 References
- ⬆️ Parent: [[Fundamentals]]
- 📚 Module: `Backtracking`
