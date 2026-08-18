---
title: History Manipulation & Recovery
tags:
  - git
  - github
  - version-control
  - history-manipulation-and-recovery
  - principal-swe
parent: "[[Git & GitHub]]"
---

# 🏛️ History Manipulation & Recovery

Undoing mistakes, interactive history curation, and disaster recovery: git reset (--soft, --mixed, --hard), git revert, interactive rebase (fixup, squash, edit), git reflog recovery, and large-scale repository cleansing (git filter-repo).

```text
History Manipulation & Recovery
│
├── [[Undoing Changes with Git Reset and Git Revert|01. Undoing Changes with Reset and Revert]]
├── [[Interactive Rebase (git Rebase I) and Commit Editing|02. Interactive Rebase and History Curation]]
├── [[Git Reflog Disaster Recovery and Lost Commit Salvage|03. Git Reflog and Disaster Recovery]]
├── [[Git Stash Mechanics and Working State Shelving|04. Stashing and Context Switching]]
├── [[Large Scale History Rewriting (git Filter Repo)|05. Large Scale History Rewriting]]
└── [[Safe Force Pushing (git Push Force with Lease)|06. Safe Force Pushing with Lease]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Undoing Changes with Git Reset and Git Revert|01. Undoing Changes with Reset and Revert]] — Safe history reversal via forward-compensating git revert vs local state rewinding via git reset (--soft, --mixed, --hard).
- 📂 [[Interactive Rebase (git Rebase I) and Commit Editing|02. Interactive Rebase and History Curation]] — Rewording, squashing, splitting, reordering, and dropping commits to produce publication-grade PR histories.
- 📂 [[Git Reflog Disaster Recovery and Lost Commit Salvage|03. Git Reflog and Disaster Recovery]] — Tracking all HEAD pointer movements (.git/logs/HEAD) to instantly restore dropped branches, hard-reset commits, and broken rebases.
- 📂 [[Git Stash Mechanics and Working State Shelving|04. Stashing and Context Switching]] — Shelving uncommitted work in progress onto stash stack, stash untracked files, and restoring with git stash pop.
- 📂 [[Large Scale History Rewriting (git Filter Repo)|05. Large Scale History Rewriting]] — Purging accidentally committed secrets, private keys, or multi-gigabyte binary files from historical commit trees.
- 📂 [[Safe Force Pushing (git Push Force with Lease)|06. Safe Force Pushing with Lease]] — Overriding remote branch histories safely by validating that no un-fetched remote commits exist before pushing.

---

## 🔗 References
- ⬆️ Parent: [[Git & GitHub]]

