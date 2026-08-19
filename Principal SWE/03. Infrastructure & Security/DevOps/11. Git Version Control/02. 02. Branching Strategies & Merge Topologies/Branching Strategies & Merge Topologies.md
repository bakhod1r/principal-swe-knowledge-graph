---
title: Branching Strategies & Merge Topologies
tags:
  - review
  - devops
  - git-and-github
  - version-control
  - branching-strategies-and-merge-topologies
  - principal-swe
parent: "[[Git Version Control]]"
---

# 🐙 Branching Strategies & Merge Topologies

Branching topologies: Trunk-Based Development, GitFlow, GitHub Flow, 3-way recursive merges, fast-forward merges, merge commits vs squash merges, and octopus merges.

```text
Branching Strategies & Merge Topologies
│
├── [[Enterprise Branching Strategies: Trunk-Based Development vs GitFlow|01. 01. Branching Workflows Trunk Based vs GitFlow]]
├── [[Merge Topologies: Fast-Forward Merges vs True Three-Way Merges|02. 02. Fast-Forward vs True Three-Way Merges]]
├── [[Squash Merging, PR Cleanliness, and Commit Atomicity|03. 03. Squash Merging and Clean Commit History]]
├── [[Managing Long-Lived Release Branches and Hotfix Backports|04. 04. Long-Lived Release Branches and Hotfix Topologies]]
├── [[Git Worktrees: Checking Out Multiple Branches Simultaneously|05. 05. Git Worktrees for Multi-Branch Parallelism]]
└── [[Remote Tracking Branches, Upstream Remotes, and Fetch vs Pull|06. 06. Remote Tracking Branches and Upstream Synchronization]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Enterprise Branching Strategies: Trunk-Based Development vs GitFlow|01. Branching Workflows Trunk Based vs GitFlow]] — Short-lived feature branches, continuous integration in Trunk-Based Development, release branches in GitFlow, and feature flag decoupling.
- 📂 [[Merge Topologies: Fast-Forward Merges vs True Three-Way Merges|02. Fast-Forward vs True Three-Way Merges]] — Three-way merge base calculation (`merge-base`), creating explicit merge commits (`--no-ff`), linear git history, and merge conflict physics.
- 📂 [[Squash Merging, PR Cleanliness, and Commit Atomicity|03. Squash Merging and Clean Commit History]] — Consolidating noisy work-in-progress commits into a single atomic commit, writing structured conventional commit messages, and bisectability.
- 📂 [[Managing Long-Lived Release Branches and Hotfix Backports|04. Long-Lived Release Branches and Hotfix Topologies]] — Cherry-picking emergency hotfixes across release branches (`git cherry-pick`), semantic release tagging, and backporting bug fixes to trunk.
- 📂 [[Git Worktrees: Checking Out Multiple Branches Simultaneously|05. Git Worktrees for Multi-Branch Parallelism]] — Creating linked working trees (`git worktree add`), testing PRs without stashing/switching branches, and disk-efficient multi-branch development.
- 📂 [[Remote Tracking Branches, Upstream Remotes, and Fetch vs Pull|06. Remote Tracking Branches and Upstream Synchronization]] — Tracking branches (`origin/main`), configuring multiple remotes (upstream vs fork), `git fetch` vs `git pull --rebase`, and remote ref pruning.

---

## 🔗 References
- ⬆️ Parent: `Git & GitHub Version Control & CI-CD Automation`
- 📚 Module: `DevOps`

