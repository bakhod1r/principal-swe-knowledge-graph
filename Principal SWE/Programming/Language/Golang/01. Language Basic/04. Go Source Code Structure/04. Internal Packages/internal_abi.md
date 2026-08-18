---
title: internal/abi
tags:
  - golang
  - goroot
  - runtime
  - internals
parent: "[[Internal Packages]]"
---

# `internal/abi`

Defines the memory layout contracts shared by the compiler, the linker, and the
runtime. Changing a struct here changes the generated machine code.

## 1. What Lives There

| Type | Describes |
|---|---|
| `abi.Type` | The runtime type descriptor every interface value points at |
| `abi.ITab` | Interface method table — see `iface (Interfaces)` |
| `abi.SwissMapType` | Map header layout — see `map (Hash Maps)` |
| `abi.FuncPCABI0` / `ABIInternal` | Calling-convention selection intrinsics |
| Register ABI constants | How arguments are passed in registers |

## 2. Why It Exists

`cmd_compile` emits data structures that `cmd_link` lays out and the runtime
reads at execution time. All three must agree byte-for-byte. `internal/abi` is
that single shared definition, replacing constants that were previously
duplicated across the three.

```text
cmd/compile ─┐
cmd/link    ─┼──► internal/abi (one definition) ──► runtime
runtime     ─┘
```

## 3. The Register ABI

Go 1.17 moved from stack-based to register-based argument passing — a
~5–10% across-the-board speedup. The register assignment rules are encoded here.

## 4. Gotchas

- **Unimportable** by user code — see `internal visibility rule`.
- `unsafe` tricks that hard-code type-descriptor offsets break whenever this
  package changes. It is the reason such tricks are unsupported.
- Reading it is the fastest route to understanding how interfaces and maps are
  actually represented.

---

## 🔗 References
- ⬆️ Parent: [[Internal Packages]]
