---
title: Git Core Foundations & Architecture
tags:
  - git
  - github
  - version-control
  - git-core-foundations-and-architecture
  - principal-swe
parent: "[[Git & GitHub]]"
---

# 🏛️ Git Core Foundations & Architecture

Git internal directed acyclic graph (DAG) storage architecture: Content-addressable object store (.git/objects), 4 fundamental object types (Blobs, Trees, Commits, Annotated Tags), SHA-1/SHA-256 hashing, index staging mechanics, and packfile compression.

```text
Git Core Foundations & Architecture
│
├── [[Git Content Addressable Object Store (.git-objects)|01. Content Addressable Object Store]]
├── [[Git Three Tree Architecture (working Tree, Index, Head)|02. Three Trees Architecture]]
├── [[Git Configuration Hierarchy and Cryptographic Signing|03. Git Configuration and Identity]]
├── [[Git Packfiles, Deltas, and Garbage Collection (git Gc)|04. Packfiles and Garbage Collection]]
├── [[Git References, Symbolic Head, and Detached Head|05. Git References and Symbolic Head]]
└── [[Git Content Diffing and Patch Generation|06. Content Diffing and Index Mechanics]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Git Content Addressable Object Store (.git-objects)|01. Content Addressable Object Store]] — How Git stores data as immutable zlib-compressed objects keyed by SHA cryptographic hashes (blobs, trees, commits, tags).
- 📂 [[Git Three Tree Architecture (working Tree, Index, Head)|02. Three Trees Architecture]] — The state machine mechanics governing file transitions across working directory, staging area index, and commit history.
- 📂 [[Git Configuration Hierarchy and Cryptographic Signing|03. Git Configuration and Identity]] — Local, global, and system configs (.gitconfig), SSH key authentication, and commit signing with GPG/SSH.
- 📂 [[Git Packfiles, Deltas, and Garbage Collection (git Gc)|04. Packfiles and Garbage Collection]] — Thin packfile compression, delta offset encoding, unreachable object pruning, and repository maintenance.
- 📂 [[Git References, Symbolic Head, and Detached Head|05. Git References and Symbolic Head]] — How branches and tags are lightweight file pointers in .git/refs, direct commit checkout, and HEAD dereferencing.
- 📂 [[Git Content Diffing and Patch Generation|06. Content Diffing and Index Mechanics]] — Myers diff algorithm, tracking staged vs unstaged file state changes, and binary delta representation.

---

## 🔗 References
- ⬆️ Parent: [[Git & GitHub]]

