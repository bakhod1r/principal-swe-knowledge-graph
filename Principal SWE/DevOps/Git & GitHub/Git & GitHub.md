---
title: Git & GitHub
tags:
  - git
  - github
  - version-control
  - devops
  - principal-swe
parent: "[[Principal SWE]]"
---

# 🐙 Git & GitHub Mastery (Architecture, Collaboration & CI/CD)

Comprehensive, production-grade master architecture covering the complete spectrum of modern distributed version control and GitHub enterprise workflows: Git directed acyclic graph (DAG) object storage (.git/objects), branching and merge strategies, history curation and disaster recovery with reflog, advanced plumbing tools (git bisect, worktrees, LFS, submodules), enterprise GitHub team governance (CODEOWNERS, branch rulesets), automated GitHub Actions CI/CD pipelines, programmatic API/GraphQL integrations, and enterprise security (Secret Scanning, CodeQL SAST, Dependabot, Codespaces) across 8 master pillars and 46 specialized subdomains.

```text
Git & GitHub
│
├── [[Git Core Foundations & Architecture|01. Git Core Foundations & Architecture]]
│   ├── [[Git Content Addressable Object Store (.git-objects)|01. Content Addressable Object Store]]
│   ├── [[Git Three Tree Architecture (working Tree, Index, Head)|02. Three Trees Architecture]]
│   ├── [[Git Configuration Hierarchy and Cryptographic Signing|03. Git Configuration and Identity]]
│   ├── [[Git Packfiles, Deltas, and Garbage Collection (git Gc)|04. Packfiles and Garbage Collection]]
│   ├── [[Git References, Symbolic Head, and Detached Head|05. Git References and Symbolic Head]]
│   └── [[Git Content Diffing and Patch Generation|06. Content Diffing and Index Mechanics]]
├── [[Branching & Merging Strategies|02. Branching & Merging Strategies]]
│   ├── [[Fast Forward vs Three Way Recursive Merges|01. Fast Forward vs 3 Way Merge]]
│   ├── [[Git Rebase Mechanics and Linear History|02. Git Rebase and History Linearization]]
│   ├── [[Squash Merging and Atomic Commit Curation|03. Squash Merging and Clean Commit Logs]]
│   ├── [[Git Cherry Pick and Selective Commit Backporting|04. Cherry Picking and Patch Application]]
│   ├── [[Merge Conflict Resolution and Rerere|05. Merge Conflict Resolution Mechanics]]
│   └── [[Trunk Based Development vs Gitflow Workflows|06. Trunk Based Development vs Gitflow]]
├── [[History Manipulation & Recovery|03. History Manipulation & Recovery]]
│   ├── [[Undoing Changes with Git Reset and Git Revert|01. Undoing Changes with Reset and Revert]]
│   ├── [[Interactive Rebase (git Rebase I) and Commit Editing|02. Interactive Rebase and History Curation]]
│   ├── [[Git Reflog Disaster Recovery and Lost Commit Salvage|03. Git Reflog and Disaster Recovery]]
│   ├── [[Git Stash Mechanics and Working State Shelving|04. Stashing and Context Switching]]
│   ├── [[Large Scale History Rewriting (git Filter Repo)|05. Large Scale History Rewriting]]
│   └── [[Safe Force Pushing (git Push Force with Lease)|06. Safe Force Pushing with Lease]]
├── [[Advanced Git Plumbing & Tooling|04. Advanced Git Plumbing & Tooling]]
│   ├── [[Git Bisect Automated Binary Search Debugging|01. Git Bisect Automated Regression Isolation]]
│   ├── [[Git Worktree for Concurrent Branch Execution|02. Git Worktree Multi Branch Checkouts]]
│   ├── [[Git Large File Storage (git Lfs) Architecture|03. Git Large File Storage LFS]]
│   ├── [[Git Submodules vs Git Subtrees for Monorepos|04. Git Submodules and Subtrees]]
│   ├── [[Git Attributes (.gitattributes) and Filter Drivers|05. Git Attributes and Path Specific Rules]]
│   └── [[Client Side and Server Side Git Hooks|06. Client and Server Git Hooks]]
├── [[GitHub Collaboration & Team Governance|05. GitHub Collaboration & Team Governance]]
│   ├── [[GitHub Organizations, Roles, and Team Permissions|01. GitHub Organizations and Team Permissions]]
│   ├── [[GitHub Branch Protection Rules and Repository Rulesets|02. Branch Protection and Rulesets]]
│   ├── [[Codeowners and Automated Code Review Assignment|03. Codeowners and Automated Review Routing]]
│   ├── [[GitHub Forking Workflow and Cross Repository PRs|04. Fork and Pull Request Workflow]]
│   ├── [[GitHub Collaborative Code Review Features|05. Collaborative Pull Request Reviews]]
│   └── [[GitHub Projects (v2) and Issue Triage Automation|06. GitHub Projects and Issue Automation]]
├── [[GitHub Actions & CI CD Automation|06. GitHub Actions & CI-CD Automation]]
│   ├── [[GitHub Actions Workflow Syntax and Event Triggers|01. GitHub Actions Workflow Syntax and Triggers]]
│   ├── [[GitHub Actions Matrix Strategy and Parallel Builds|02. Matrix Strategy and Parallel Builds]]
│   ├── [[GitHub Actions Runners (cloud Hosted vs Self Hosted)|03. Runners Hosted vs Self Hosted]]
│   ├── [[GitHub Actions Secrets and Cloud Openid Connect (oidc)|04. Secrets Management and Cloud Oidc]]
│   ├── [[GitHub Actions Dependency Caching and Artifact Storage|05. Dependency Caching and Artifact Storage]]
│   └── [[Authoring Custom Composite and Docker Container Actions|06. Custom Composite and Docker Actions]]
├── [[GitHub Developer Ecosystem & APIs|07. GitHub Developer Ecosystem & APIs]]
│   ├── [[GitHub REST API V3 Architecture and Pagination|01. GitHub REST API V3 Integration]]
│   ├── [[GitHub GraphQL API V4 Query Optimization|02. GitHub GraphQL API V4 Integration]]
│   ├── [[GitHub Apps vs Oauth Apps and Fine Grained Permissions|03. GitHub Apps vs Oauth Apps]]
│   ├── [[GitHub Webhook Delivery and Hmac SHA 256 Signatures|04. GitHub Webhooks and Hmac Validation]]
│   └── [[GitHub CLI (gh) Scripting and Extension Authoring|05. GitHub CLI Automation Scripting]]
└── [[Enterprise GitHub Features & Security|08. Enterprise GitHub Features & Security]]
│   ├── [[GitHub Advanced Security and Secret Push Protection|01. GitHub Advanced Security Ghas and Secret Scanning]]
│   ├── [[CodeQL Semantic Code Analysis and Query Authoring|02. CodeQL Static Code Analysis SAST]]
│   ├── [[Dependabot Automated Dependency Updates and Alerts|03. Dependabot and Supply Chain Security]]
│   ├── [[GitHub Codespaces and Dev Container Standard|04. GitHub Codespaces Cloud Environments]]
│   └── [[GitHub Packages (ghcr) and Multi Language Registries|05. GitHub Packages and Container Registry]]
```

---

## 🏛️ Core Knowledge Pillars

### 1. 📂 [[Git Core Foundations & Architecture|01. Git Core Foundations & Architecture]]
- 📂 [[Git Content Addressable Object Store (.git-objects)|01. Content Addressable Object Store]] — How Git stores data as immutable zlib-compressed objects keyed by SHA cryptographic hashes (blobs, trees, commits, tags).
- 📂 [[Git Three Tree Architecture (working Tree, Index, Head)|02. Three Trees Architecture]] — The state machine mechanics governing file transitions across working directory, staging area index, and commit history.
- 📂 [[Git Configuration Hierarchy and Cryptographic Signing|03. Git Configuration and Identity]] — Local, global, and system configs (.gitconfig), SSH key authentication, and commit signing with GPG/SSH.
- 📂 [[Git Packfiles, Deltas, and Garbage Collection (git Gc)|04. Packfiles and Garbage Collection]] — Thin packfile compression, delta offset encoding, unreachable object pruning, and repository maintenance.
- 📂 [[Git References, Symbolic Head, and Detached Head|05. Git References and Symbolic Head]] — How branches and tags are lightweight file pointers in .git/refs, direct commit checkout, and HEAD dereferencing.
- 📂 [[Git Content Diffing and Patch Generation|06. Content Diffing and Index Mechanics]] — Myers diff algorithm, tracking staged vs unstaged file state changes, and binary delta representation.
### 2. 📂 [[Branching & Merging Strategies|02. Branching & Merging Strategies]]
- 📂 [[Fast Forward vs Three Way Recursive Merges|01. Fast Forward vs 3 Way Merge]] — Pointer advancement vs merge commits, lowest common ancestor (LCA) detection, and merge base calculation.
- 📂 [[Git Rebase Mechanics and Linear History|02. Git Rebase and History Linearization]] — Replaying commit deltas onto a new upstream base, preserving linear audit history, and avoid-rebase-on-shared-branches rule.
- 📂 [[Squash Merging and Atomic Commit Curation|03. Squash Merging and Clean Commit Logs]] — Collapsing multi-commit feature branches into a single clean commit with structured conventional commit messages.
- 📂 [[Git Cherry Pick and Selective Commit Backporting|04. Cherry Picking and Patch Application]] — Applying isolated commit diffs across disparate release branches without merging parent branch histories.
- 📂 [[Merge Conflict Resolution and Rerere|05. Merge Conflict Resolution Mechanics]] — Understanding 3-way conflict markers (ours, theirs, base), diff3 format, and Reuse Recorded Resolution (git rerere).
- 📂 [[Trunk Based Development vs Gitflow Workflows|06. Trunk Based Development vs Gitflow]] — High-velocity short-lived feature branches (<24h) with feature flags vs long-lived release and hotfix branch hierarchies.
### 3. 📂 [[History Manipulation & Recovery|03. History Manipulation & Recovery]]
- 📂 [[Undoing Changes with Git Reset and Git Revert|01. Undoing Changes with Reset and Revert]] — Safe history reversal via forward-compensating git revert vs local state rewinding via git reset (--soft, --mixed, --hard).
- 📂 [[Interactive Rebase (git Rebase I) and Commit Editing|02. Interactive Rebase and History Curation]] — Rewording, squashing, splitting, reordering, and dropping commits to produce publication-grade PR histories.
- 📂 [[Git Reflog Disaster Recovery and Lost Commit Salvage|03. Git Reflog and Disaster Recovery]] — Tracking all HEAD pointer movements (.git/logs/HEAD) to instantly restore dropped branches, hard-reset commits, and broken rebases.
- 📂 [[Git Stash Mechanics and Working State Shelving|04. Stashing and Context Switching]] — Shelving uncommitted work in progress onto stash stack, stash untracked files, and restoring with git stash pop.
- 📂 [[Large Scale History Rewriting (git Filter Repo)|05. Large Scale History Rewriting]] — Purging accidentally committed secrets, private keys, or multi-gigabyte binary files from historical commit trees.
- 📂 [[Safe Force Pushing (git Push Force with Lease)|06. Safe Force Pushing with Lease]] — Overriding remote branch histories safely by validating that no un-fetched remote commits exist before pushing.
### 4. 📂 [[Advanced Git Plumbing & Tooling|04. Advanced Git Plumbing & Tooling]]
- 📂 [[Git Bisect Automated Binary Search Debugging|01. Git Bisect Automated Regression Isolation]] — Automating defect localization across thousands of commits in O(log N) steps using git bisect run with automated test scripts.
- 📂 [[Git Worktree for Concurrent Branch Execution|02. Git Worktree Multi Branch Checkouts]] — Checking out multiple branches into distinct filesystem directories simultaneously sharing a single .git database.
- 📂 [[Git Large File Storage (git Lfs) Architecture|03. Git Large File Storage LFS]] — Replacing heavy binary assets (videos, model weights, datasets) with pointer files backed by remote deduplicated storage.
- 📂 [[Git Submodules vs Git Subtrees for Monorepos|04. Git Submodules and Subtrees]] — Embedding nested external Git repositories with commit pinning vs merging external repo trees into subdirectories.
- 📂 [[Git Attributes (.gitattributes) and Filter Drivers|05. Git Attributes and Path Specific Rules]] — Customizing line ending normalization (CRLF/LF), custom diff drivers for binary files, and smudge/clean filters.
- 📂 [[Client Side and Server Side Git Hooks|06. Client and Server Git Hooks]] — Automating pre-commit linting, commit-msg format validation, and server pre-receive push authorization policies.
### 5. 📂 [[GitHub Collaboration & Team Governance|05. GitHub Collaboration & Team Governance]]
- 📂 [[GitHub Organizations, Roles, and Team Permissions|01. GitHub Organizations and Team Permissions]] — Managing enterprise organization membership, team hierarchy mapping, base permissions, and SAML SSO integration.
- 📂 [[GitHub Branch Protection Rules and Repository Rulesets|02. Branch Protection and Rulesets]] — Enforcing mandatory pull request reviews, passing status checks, linear history, signed commits, and admin restrictions.
- 📂 [[Codeowners and Automated Code Review Assignment|03. Codeowners and Automated Review Routing]] — Defining directory-level code ownership rules (.github/CODEOWNERS) and auto-assigning domain specialist reviewers.
- 📂 [[GitHub Forking Workflow and Cross Repository PRs|04. Fork and Pull Request Workflow]] — Managing open-source contributions, syncing forks with upstream, and cross-repo CI permission boundaries.
- 📂 [[GitHub Collaborative Code Review Features|05. Collaborative Pull Request Reviews]] — Review suggestions, thread resolution requirements, batch comments, draft PRs, and PR turnaround metrics.
- 📂 [[GitHub Projects (v2) and Issue Triage Automation|06. GitHub Projects and Issue Automation]] — Automated board workflows, custom fields, iteration planning, roadmap views, and automated status transitions.
### 6. 📂 [[GitHub Actions & CI CD Automation|06. GitHub Actions & CI-CD Automation]]
- 📂 [[GitHub Actions Workflow Syntax and Event Triggers|01. GitHub Actions Workflow Syntax and Triggers]] — Push, pull_request, workflow_dispatch, and scheduled cron triggers; job dependencies, outputs, and conditional expressions.
- 📂 [[GitHub Actions Matrix Strategy and Parallel Builds|02. Matrix Strategy and Parallel Builds]] — Running cross-platform test suites across multiple OS versions and runtime language versions simultaneously.
- 📂 [[GitHub Actions Runners (cloud Hosted vs Self Hosted)|03. Runners Hosted vs Self Hosted]] — Configuring enterprise self-hosted runner fleets on Kubernetes (Actions Runner Controller ARC) with ephemeral pods.
- 📂 [[GitHub Actions Secrets and Cloud Openid Connect (oidc)|04. Secrets Management and Cloud Oidc]] — Eliminating long-lived cloud credentials using short-lived OIDC tokens for passwordless AWS/GCP/Azure role assumption.
- 📂 [[GitHub Actions Dependency Caching and Artifact Storage|05. Dependency Caching and Artifact Storage]] — Accelerating CI build times via actions/cache for package managers and publishing build binaries via actions/upload-artifact.
- 📂 [[Authoring Custom Composite and Docker Container Actions|06. Custom Composite and Docker Actions]] — Building reusable organization-wide Actions (action.yml) with composite steps and containerized execution.
### 7. 📂 [[GitHub Developer Ecosystem & APIs|07. GitHub Developer Ecosystem & APIs]]
- 📂 [[GitHub REST API V3 Architecture and Pagination|01. GitHub REST API V3 Integration]] — Programmatic repository, issue, and release management; rate limiting (5,000 req/hr), and Link header pagination.
- 📂 [[GitHub GraphQL API V4 Query Optimization|02. GitHub GraphQL API V4 Integration]] — Fetching deeply nested issue, PR, and commit graph metadata in a single HTTP request with strict type schemas.
- 📂 [[GitHub Apps vs Oauth Apps and Fine Grained Permissions|03. GitHub Apps vs Oauth Apps]] — Building service integrations as first-class GitHub Apps with granular repository permissions and private key authentication.
- 📂 [[GitHub Webhook Delivery and Hmac SHA 256 Signatures|04. GitHub Webhooks and Hmac Validation]] — Handling asynchronous repository event notifications and verifying authenticity via X-Hub-Signature-256.
- 📂 [[GitHub CLI (gh) Scripting and Extension Authoring|05. GitHub CLI Automation Scripting]] — Automating developer workflows via gh cli commands, json output queries with jq, and building custom gh extensions.
### 8. 📂 [[Enterprise GitHub Features & Security|08. Enterprise GitHub Features & Security]]
- 📂 [[GitHub Advanced Security and Secret Push Protection|01. GitHub Advanced Security Ghas and Secret Scanning]] — Blocking accidental commits containing API tokens, high-entropy secrets, and private keys before reaching remote repos.
- 📂 [[CodeQL Semantic Code Analysis and Query Authoring|02. CodeQL Static Code Analysis SAST]] — Querying codebase syntax trees as data to detect zero-day injection flaws, memory corruptions, and security regressions.
- 📂 [[Dependabot Automated Dependency Updates and Alerts|03. Dependabot and Supply Chain Security]] — Automated vulnerability alerts (CVEs), automated security update PRs, and dependency graph review.
- 📂 [[GitHub Codespaces and Dev Container Standard|04. GitHub Codespaces Cloud Environments]] — Zero-configuration cloud development environments using devcontainer.json, Docker containers, and browser VS Code.
- 📂 [[GitHub Packages (ghcr) and Multi Language Registries|05. GitHub Packages and Container Registry]] — Publishing and distributing Docker container images, npm packages, Go modules, and Maven artifacts directly on GitHub.

---

## 🔗 Navigation
- ⬆️ Parent: [[Principal SWE]]
- 🎓 Root: [[Principal SWE]]
