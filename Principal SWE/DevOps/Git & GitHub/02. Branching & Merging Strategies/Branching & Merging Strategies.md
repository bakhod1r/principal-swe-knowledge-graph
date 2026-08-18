---
title: Branching & Merging Strategies
tags:
  - git
  - github
  - version-control
  - branching-and-merging-strategies
  - principal-swe
parent: "[[Git & GitHub]]"
---

# 🏛️ Branching & Merging Strategies

Branching workflows and merge algebra: Fast-forward merges, 3-way recursive merges, rebase vs merge trade-offs, squash merges, cherry-picking, three-way merge conflict resolution, and trunk-based development vs GitFlow.

```text
Branching & Merging Strategies
│
├── [[Fast Forward vs Three Way Recursive Merges|01. Fast Forward vs 3 Way Merge]]
├── [[Git Rebase Mechanics and Linear History|02. Git Rebase and History Linearization]]
├── [[Squash Merging and Atomic Commit Curation|03. Squash Merging and Clean Commit Logs]]
├── [[Git Cherry Pick and Selective Commit Backporting|04. Cherry Picking and Patch Application]]
├── [[Merge Conflict Resolution and Rerere|05. Merge Conflict Resolution Mechanics]]
└── [[Trunk Based Development vs Gitflow Workflows|06. Trunk Based Development vs Gitflow]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Fast Forward vs Three Way Recursive Merges|01. Fast Forward vs 3 Way Merge]] — Pointer advancement vs merge commits, lowest common ancestor (LCA) detection, and merge base calculation.
- 📂 [[Git Rebase Mechanics and Linear History|02. Git Rebase and History Linearization]] — Replaying commit deltas onto a new upstream base, preserving linear audit history, and avoid-rebase-on-shared-branches rule.
- 📂 [[Squash Merging and Atomic Commit Curation|03. Squash Merging and Clean Commit Logs]] — Collapsing multi-commit feature branches into a single clean commit with structured conventional commit messages.
- 📂 [[Git Cherry Pick and Selective Commit Backporting|04. Cherry Picking and Patch Application]] — Applying isolated commit diffs across disparate release branches without merging parent branch histories.
- 📂 [[Merge Conflict Resolution and Rerere|05. Merge Conflict Resolution Mechanics]] — Understanding 3-way conflict markers (ours, theirs, base), diff3 format, and Reuse Recorded Resolution (git rerere).
- 📂 [[Trunk Based Development vs Gitflow Workflows|06. Trunk Based Development vs Gitflow]] — High-velocity short-lived feature branches (<24h) with feature flags vs long-lived release and hotfix branch hierarchies.

---

## 🔗 References
- ⬆️ Parent: [[Git & GitHub]]

