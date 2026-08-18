---
title: Pseudo-versions
tags:
  - golang
  - basics
  - dependencies
  - modules
  - versioning
parent: "[[Version Resolution]]"
---

# Pseudo-versions

Synthetic semantic versions the go command generates for commits that carry no
semver tag.

## 1. Anatomy

```text
v0.0.0-20260218091402-9f2c1ab3d4e5
│      │              │
│      │              └── 12-char commit hash
│      └── commit time, UTC, yyyymmddhhmmss
└── base version
```

## 2. Three Forms

| Form | When |
|---|---|
| `v0.0.0-<time>-<hash>` | No earlier tag exists in the history |
| `v1.2.4-0.<time>-<hash>` | Latest tag before the commit is `v1.2.3` (patch bumped, pre-release) |
| `v1.3.0-pre.0.<time>-<hash>` | Latest tag is a pre-release `v1.3.0-pre` |

The base is always chosen so the pseudo-version sorts **after** the previous tag
and **before** the next real release. See `Semantic Versioning`.

## 3. Getting One

```bash
go get github.com/me/lib@main
go get github.com/me/lib@9f2c1ab
go get github.com/me/lib@2026-02-18   # nearest commit before that date
```

The go command resolves the ref to a commit and constructs the pseudo-version —
you never type one by hand.

## 4. Gotchas

- The **timestamp is the commit time, not the tag or fetch time**. Two branches
  can produce pseudo-versions whose ordering surprises you.
- A pseudo-version pins an exact commit; it is fully reproducible and recorded in
  `go.sum` like any version.
- Many pseudo-versions in `go.mod` usually mean an upstream that never tags —
  a real supply-chain risk flag. See `Dependency Auditing`.
- `+incompatible` and pseudo-versions are unrelated; see `Incompatible Versions`.

---

## 🔗 References
- ⬆️ Parent: `02. Version Resolution`
