---
title: Conflict Resolution & Interactive Debugging
tags:
  - review
  - devops
  - git-and-github
  - version-control
  - conflict-resolution-and-interactive-debugging
  - principal-swe
parent: "[[Git Version Control]]"
---

# 🐙 Conflict Resolution & Interactive Debugging

Resolving complex merge conflicts: 3-way diff tools, conflict markers, `git merge-base`, binary file conflicts, submodules vs subtrees, and binary search debugging with `git bisect`.

```text
Conflict Resolution & Interactive Debugging
│
├── [[Anatomy of Merge Conflicts, Diff3 Formatting, and Markers|01. 01. Anatomy of Merge Conflicts and Conflict Markers]]
├── [[Automated Regression Hunting with Git Bisect and Test Scripts|02. 02. Automated Regression Hunting with Git Bisect]]
├── [[Git Log Archaeology, Pickaxe Search (-S), and Line History|03. 03. Git Blame, Log Archaeology, and Git Pickaxe]]
├── [[Git Submodules vs Git Subtrees for Monorepos and Dependency Sharing|04. 04. Git Submodules and Monorepo Subtrees]]
├── [[Binary Asset Versioning with Git Large File Storage (Git LFS)|05. 05. Binary Asset Versioning and Git LFS]]
└── [[Git Stash Internals, Stash Stack, and Patch Formatting|06. 06. Git Stash Internals and Patch Queues]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Anatomy of Merge Conflicts, Diff3 Formatting, and Markers|01. Anatomy of Merge Conflicts and Conflict Markers]] — Understanding `<<<<<<<`, `=======`, `>>>>>>>` markers, enabling diff3 / zdiff3 style showing the common ancestor base, and manual conflict resolution.
- 📂 [[Automated Regression Hunting with Git Bisect and Test Scripts|02. Automated Regression Hunting with Git Bisect]] — Binary search through commit history (`git bisect start`, `bad`, `good`), automating bug isolation with test scripts (`git bisect run pytest`), and bisect logs.
- 📂 [[Git Log Archaeology, Pickaxe Search (-S), and Line History|03. Git Blame, Log Archaeology, and Git Pickaxe]] — Investigating code origin with `git blame -w -M -C`, searching commit diffs for specific string additions/deletions (`git log -S'func_name'`), and patch diffs.
- 📂 [[Git Submodules vs Git Subtrees for Monorepos and Dependency Sharing|04. Git Submodules and Monorepo Subtrees]] — Managing submodules (`.gitmodules`), updating recursive submodules, detangling submodule pointer detached states, and merging external repos via `git subtree`.
- 📂 [[Binary Asset Versioning with Git Large File Storage (Git LFS)|05. Binary Asset Versioning and Git LFS]] — Tracking large binaries (.psd, .onnx, .tar.gz), pointer files in git history, smudge/clean filters, and configuring LFS storage backends (S3/Artifactory).
- 📂 [[Git Stash Internals, Stash Stack, and Patch Formatting|06. Git Stash Internals and Patch Queues]] — How `git stash` creates two dangling commit objects (working tree + index), stashing untracked files (`--include-untracked`), and creating email patches with `git format-patch`.

---

## 🔗 References
- ⬆️ Parent: `Git & GitHub Version Control & CI-CD Automation`
- 📚 Module: `DevOps`

