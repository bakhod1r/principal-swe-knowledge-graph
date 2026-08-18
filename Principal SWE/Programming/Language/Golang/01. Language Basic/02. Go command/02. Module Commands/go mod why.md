---
title: go mod why
tags:
  - golang
  - basics
  - cli
  - toolchain
  - modules
  - debugging
parent: "[[Module Commands]]"
---

# `go mod why`

Explains the shortest import chain from the main module to a given package or
module.

```bash
go mod why golang.org/x/text/language
go mod why -m golang.org/x/text        # module instead of package
go mod why -m all                      # every module in the build list
```

## 1. Output

```text
# golang.org/x/text/language
github.com/me/api/internal/i18n
golang.org/x/text/language
```

Each line is a package importing the next. The chain answers "who dragged this
in" in one command.

## 2. When Nothing Needs It

```text
# github.com/unused/lib
(main module does not need module github.com/unused/lib)
```

That is the signal to run `go mod tidy`.

## 3. `-vendor` Flag

```bash
go mod why -vendor -m all
```

Considers only imports reachable in the vendored build — matches what
`Vendoring` would actually include.

## 4. Gotchas

- It reports the **shortest** path, not all paths. Removing that one import may
  not remove the dependency.
- Test-only imports count, so a dependency may exist purely because of a test in
  a dependency.
- For version questions rather than existence questions, use `go mod graph`.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
