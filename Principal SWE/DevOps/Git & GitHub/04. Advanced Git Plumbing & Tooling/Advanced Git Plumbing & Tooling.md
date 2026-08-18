---
title: Advanced Git Plumbing & Tooling
tags:
  - git
  - github
  - version-control
  - advanced-git-plumbing-and-tooling
  - principal-swe
parent: "[[Git & GitHub]]"
---

# 🏛️ Advanced Git Plumbing & Tooling

Plumbing tools and power-user workflows: Binary search regression isolation (git bisect), multi-branch concurrent checkouts (git worktree), large file management (Git LFS), submodules & subtrees, git attributes, and client/server hooks.

```text
Advanced Git Plumbing & Tooling
│
├── [[Git Bisect Automated Binary Search Debugging|01. Git Bisect Automated Regression Isolation]]
├── [[Git Worktree for Concurrent Branch Execution|02. Git Worktree Multi Branch Checkouts]]
├── [[Git Large File Storage (git Lfs) Architecture|03. Git Large File Storage LFS]]
├── [[Git Submodules vs Git Subtrees for Monorepos|04. Git Submodules and Subtrees]]
├── [[Git Attributes (.gitattributes) and Filter Drivers|05. Git Attributes and Path Specific Rules]]
└── [[Client Side and Server Side Git Hooks|06. Client and Server Git Hooks]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Git Bisect Automated Binary Search Debugging|01. Git Bisect Automated Regression Isolation]] — Automating defect localization across thousands of commits in O(log N) steps using git bisect run with automated test scripts.
- 📂 [[Git Worktree for Concurrent Branch Execution|02. Git Worktree Multi Branch Checkouts]] — Checking out multiple branches into distinct filesystem directories simultaneously sharing a single .git database.
- 📂 [[Git Large File Storage (git Lfs) Architecture|03. Git Large File Storage LFS]] — Replacing heavy binary assets (videos, model weights, datasets) with pointer files backed by remote deduplicated storage.
- 📂 [[Git Submodules vs Git Subtrees for Monorepos|04. Git Submodules and Subtrees]] — Embedding nested external Git repositories with commit pinning vs merging external repo trees into subdirectories.
- 📂 [[Git Attributes (.gitattributes) and Filter Drivers|05. Git Attributes and Path Specific Rules]] — Customizing line ending normalization (CRLF/LF), custom diff drivers for binary files, and smudge/clean filters.
- 📂 [[Client Side and Server Side Git Hooks|06. Client and Server Git Hooks]] — Automating pre-commit linting, commit-msg format validation, and server pre-receive push authorization policies.

---

## 🔗 References
- ⬆️ Parent: [[Git & GitHub]]
- 🎓 Root: [[Principal SWE]]
