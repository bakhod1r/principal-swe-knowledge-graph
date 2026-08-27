---
title: "State, Choice, Constraint"
tags:
  - review
  - computer-science
  - algorithms
  - dsa
  - backtracking
  - principal-swe
parent: "[[Fundamentals]]"
---
# State, Choice, Constraint in Backtracking

Backtracking can be understood through three fundamental concepts:

> **State → Choice → Constraint**

These three define what the algorithm **knows**, what it **can try**, and what it **must not violate**.

---

## 1. State

### Definition

**State** is the current partial solution at a particular point in the search.

It answers:

> **“What have I decided so far?”**

For example, in the **N-Queens** problem, the state might be:

```text
[1, 3]
```

This could mean:

- Queen 1 → column 1
    
- Queen 2 → column 3
    

The remaining queens have not been placed yet.

### Mental Model

Think of state as a snapshot:

```text
Start
  ↓
State A
  ↓
State B
  ↓
State C
```

Each recursive call usually represents a new state.

### Examples of State

Depending on the problem, state can be:

```text
current index
current path
selected elements
board configuration
visited nodes
remaining capacity
current sum
current permutation
```

For example:

```go
path := []int{1, 4, 7}
```

Here:

```text
State = [1, 4, 7]
```

---

# 2. Choice

### Definition

A **choice** is one possible decision that can extend the current state.

It answers:

> **“What can I try next?”**

Suppose:

```text
State = [1, 4]
```

Possible choices might be:

```text
2
3
5
6
```

The algorithm explores these alternatives one by one.

Conceptually:

```text
             State
          /    |    \
       Choice Choice Choice
         A       B      C
```

Each choice creates a new state.

---

# 3. Constraint

### Definition

A **constraint** is a rule that determines whether a choice is valid.

It answers:

> **“Is this choice allowed?”**

For example, suppose we are solving a subset problem with:

```text
target = 10
current = 7
```

Trying:

```text
+5
```

produces:

```text
12
```

If the problem requires:

```text
sum <= 10
```

then:

```text
7 + 5 = 12
```

violates the constraint.

Therefore, that branch can be rejected immediately.

This is the key power of backtracking:

> **Don't explore branches that are already known to be invalid.**

---

# 4. The Three Concepts Together

The fundamental loop is:

```text
State
  ↓
Generate Choices
  ↓
Check Constraint
  ↓
Make Choice
  ↓
Recurse
  ↓
Undo Choice
```

In pseudocode:

```text
backtrack(state):
    if solution(state):
        record solution
        return

    for choice in choices(state):
        if violates_constraint(choice, state):
            continue

        apply(choice, state)

        backtrack(state)

        undo(choice, state)
```

This is the core structure behind many backtracking algorithms.

---

# 5. Example: Permutations

Suppose:

```text
nums = [1, 2, 3]
```

We want all permutations.

### Initial State

```text
[]
```

Choices:

```text
1
2
3
```

Choose `1`:

```text
[1]
```

Now choices:

```text
2
3
```

Choose `2`:

```text
[1, 2]
```

Choose `3`:

```text
[1, 2, 3]
```

Solution found.

Then we **backtrack**:

```text
[1, 2, 3]
       ↑
    undo 3

[1, 2]
```

Try another choice:

```text
[1, 3]
```

Then:

```text
[1, 3, 2]
```

The search tree looks like:

```text
                    []
                 /  |  \
                1   2   3
               / \ / \ / \
              2  3 1 3 1 2
              |  |
             3   2
```

Here:

- **State** → current permutation prefix
    
- **Choice** → unused number
    
- **Constraint** → cannot reuse a number
    

---

# 6. Constraint Is What Makes Backtracking Powerful

A naive brute-force algorithm may explore everything.

Backtracking instead asks at every step:

```text
Can this partial solution still become a valid solution?
```

If:

```text
NO
```

stop exploring that branch.

Example:

```text
                    root
                  /      \
                valid    invalid
                         ×
                       prune
```

This is called **pruning**.

So an important mental model is:

> **Backtracking = systematic search + constraint checking + pruning.**

---

# 7. State vs Choice vs Constraint

A useful way to distinguish them:

|Concept|Question|
|---|---|
|**State**|What have I decided so far?|
|**Choice**|What can I decide next?|
|**Constraint**|Which decisions are legal?|

Example: N-Queens

```text
State:
[1, 3]

Choice:
place next queen in column 1..N

Constraint:
cannot share row, column, or diagonal
```

Example: Combination Sum

```text
State:
current combination + current sum

Choice:
choose another candidate

Constraint:
sum must not exceed target
```

Example: Sudoku

```text
State:
current board

Choice:
put a number 1..9 into an empty cell

Constraint:
number cannot violate row/column/box rules
```

---

# 8. The Most Important Design Question

When solving a backtracking problem, don't immediately write recursion.

First identify:

### 1. What is my state?

```text
What information completely describes my current partial solution?
```

### 2. What are my choices?

```text
What decisions can I make from this state?
```

### 3. What are my constraints?

```text
Which choices are invalid?
```

### 4. What is the terminal condition?

```text
When is the state a complete solution?
```

### 5. What must I undo?

```text
Which mutation did I make before recursion?
```

That gives you the recursion almost mechanically.

---

# 9. Production-Level Mental Model

At a deeper level, backtracking is a **DFS over a state-space tree**.

```text
                    Root State
                   /          \
              State A       State B
             /     \           |
          State C  State D   State E
```

Each edge represents a **choice**.

Each node represents a **state**.

Constraints determine whether an edge/node should be explored.

Therefore:

```text
Backtracking
    =
Depth-First Search
    +
Mutable/partial state
    +
Choices
    +
Constraint validation
    +
Undo
    +
Pruning
```

The most important invariant is:

> **After returning from a recursive call, the state must be exactly as it was before the choice was made.**

That invariant is what makes `undo` correct.

---

## Key Takeaways

```text
STATE      → Where am I?
CHOICE     → What can I try?
CONSTRAINT → What is forbidden?
RECURSION  → Explore the choice.
UNDO       → Restore the previous state.
PRUNING    → Stop impossible branches early.
```

If you can correctly identify **State + Choice + Constraint**, most classical backtracking problems—**Subsets, Permutations, Combination Sum, N-Queens, Sudoku, Word Search, Partitioning**—become variations of the same underlying algorithmic pattern.
---

## 🔗 References
- ⬆️ Parent: [[Fundamentals]]
- 📚 Module: `Backtracking`
