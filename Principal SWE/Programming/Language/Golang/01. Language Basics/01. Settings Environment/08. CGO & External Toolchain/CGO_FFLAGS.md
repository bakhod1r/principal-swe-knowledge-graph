---
title: CGO_FFLAGS
tags:
  - golang
  - basics
  - environment
  - cgo
  - fortran
parent: "[[CGO & External Toolchain]]"
---

# `CGO_FFLAGS`

Flags passed to the **Fortran** compiler for `.f`, `.F`, `.f90`, and `.f95` files
in a cgo package. The Fortran member of the `CGO_CFLAGS` family.

## 1. Usage

```bash
CGO_FFLAGS="-O3 -march=native" \
CGO_LDFLAGS="-lgfortran" \
go build ./...
```

The compiler itself comes from `FC` (default `gfortran`), the Fortran analogue of
`CC` and `CXX`.

## 2. The Full cgo Flag Family

| Variable | Applies to |
|---|---|
| `CGO_CPPFLAGS` | Preprocessor, all C-family languages |
| `CGO_CFLAGS` | `.c` |
| `CGO_CXXFLAGS` | `.cc`, `.cpp`, `.cxx` |
| `CGO_FFLAGS` | `.f`, `.F`, `.f90`, `.f95` |
| `CGO_LDFLAGS` | Link step for all of them |

## 3. Where It Shows Up

Wrapping legacy numerical libraries — BLAS/LAPACK kernels, solvers, and physics
codes that were never rewritten.

## 4. Gotchas

- Fortran objects almost always need `-lgfortran` in `CGO_LDFLAGS`; forgetting
  it produces undefined `_gfortran_*` symbols at link time.
- Only honoured when `CGO_ENABLED``=1`.
- `#cgo FFLAGS:` directives in Go source are appended to this variable.

---

## 🔗 References
- ⬆️ Parent: [[CGO & External Toolchain]]
