---
title: Advanced Rebasing, Cherry-Picking & History Rewriting
tags:
  - review
  - devops
  - git-and-github
  - version-control
  - advanced-rebasing,-cherry-picking-and-history-rewriting
  - principal-swe
parent: "[[Git Version Control]]"
---

# 🐙 Advanced Rebasing, Cherry-Picking & History Rewriting

History manipulation and surgical surgery: Interactive rebase, autosquash, fixup commits, cherry-picking, reflog recovery, `filter-repo`, and `rerere` (reuse recorded resolution).

```text
Advanced Rebasing, Cherry-Picking & History Rewriting
│
├── [[Interactive Rebasing (git rebase -i) and History Crafting|01. 01. Interactive Rebasing and History Crafting]]
├── [[Fixup Commits, Amend Idioms, and Git Autosquash|02. 02. Fixup Commits and Automated Autosquash]]
├── [[Advanced Cherry-Picking (git cherry-pick) and Range Extraction|03. 03. Cherry-Picking Commits and Merge Base Tracking]]
├── [[Git Reflog (Reference Logs) and Disaster Recovery Engineering|04. 04. Git Reflog and Disaster Recovery Engineering]]
├── [[Git Rerere (Reuse Recorded Resolution) and Merge Conflict Memory|05. 05. Reuse Recorded Resolution Git Rerere]]
└── [[Large-Scale History Surgery with Git-Filter-Repo and BFG|06. 06. Large Scale History Surgery and Git Filter-Repo]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Interactive Rebasing (git rebase -i) and History Crafting|01. Interactive Rebasing and History Crafting]] — Reordering, rewording, dropping, and editing past commits; splitting a single commit into multiple atomic commits, and avoiding rebasing shared public branches.
- 📂 [[Fixup Commits, Amend Idioms, and Git Autosquash|02. Fixup Commits and Automated Autosquash]] — Streamlining PR reviews with `git commit --fixup` and `git commit --squash`, and automatically folding them into past commits with `git rebase -i --autosquash`.
- 📂 [[Advanced Cherry-Picking (git cherry-pick) and Range Extraction|03. Cherry-Picking Commits and Merge Base Tracking]] — Applying specific commits across branches, cherry-picking commit ranges (`commitA..commitB`), preserving metadata, and tracking duplicate commits with `git cherry`.
- 📂 [[Git Reflog (Reference Logs) and Disaster Recovery Engineering|04. Git Reflog and Disaster Recovery Engineering]] — Recovering accidentally deleted branches, uncommitted resets, and corrupted rebases using `.git/logs/HEAD` and `git reset --hard HEAD@{n}`.
- 📂 [[Git Rerere (Reuse Recorded Resolution) and Merge Conflict Memory|05. Reuse Recorded Resolution Git Rerere]] — Enabling `git config rerere.enabled true`, recording merge conflict resolutions, and automatically resolving identical conflicts during repeated rebases.
- 📂 [[Large-Scale History Surgery with Git-Filter-Repo and BFG|06. Large Scale History Surgery and Git Filter-Repo]] — Purging sensitive credentials/passwords permanently from git history, stripping massive binary blobs from past commits, and rewritten commit hashes.

---

## 🔗 References
- ⬆️ Parent: `Git & GitHub Version Control & CI-CD Automation`
- 📚 Module: `DevOps`

