---
title: Minimal Version Selection (MVS)
tags:
  - golang
  - dependencies
  - mvs
  - semver
  - principal-swe
parent: "[[Dependencies]]"
---

# Minimal Version Selection (MVS)

MVS deterministic algorithm, pseudo-versions, diamond dependency resolution, SemVer compatibility rules, and major version upgrades.

```text
Minimal Version Selection (MVS)
│
├── [[MVS Graph Algorithm Mechanics]]
├── [[Diamond Dependency Problem in Go Modules]]
├── [[Pseudo-Versions & Untagged Commit Resolution (v0.0.0-timestamp-hash)]]
├── [[Semantic Import Versioning (v2+ Rules)]]
├── [[Major Version Upgrades & Dual-Import Compatibility]]
└── [[retract Directive & Broken Releases]]
```

---

## 🗂️ Topics

- [[MVS Graph Algorithm Mechanics]] — Why Go rejected SAT-solvers in favor of deterministic Minimal Version Selection.
- [[Diamond Dependency Problem in Go Modules]] — Deterministic convergence on the highest requested minimum version in dependency diamond graphs.
- [[Pseudo-Versions & Untagged Commit Resolution (v0.0.0-timestamp-hash)]] — Commit timestamp and revision hash format for untagged git dependencies.
- [[Semantic Import Versioning (v2+ Rules)]] — Major version suffixes in import paths (`pkg/v2`) and backwards compatibility contracts.
- [[Major Version Upgrades & Dual-Import Compatibility]] — Coexisting `v1` and `v2` imports within the same binary during phased migrations.
- [[retract Directive & Broken Releases]] — Retracting broken or compromised module releases in `go.mod`.

---

## 🔗 References
- ⬆️ Parent: [[Dependencies]]
- 📚 Module: `Go Environment & Commands`
