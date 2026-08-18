---
title: path/filepath Package
tags:
  - golang
  - standard-library
  - filesystem
parent: "[[Standard Library]]"
---

# `path/filepath`

OS-aware path manipulation. Distinct from `path`, which is for slash-separated
paths like URLs.

## 1. `path` vs `path/filepath`

| | Separator | Use for |
|---|---|---|
| `path` | always `/` | URLs, import paths, `io/fs` paths |
| `path/filepath` | OS-specific | Real filesystem paths |

Using `path.Join` on Windows filesystem paths produces `a/b` where `a\b` was
needed — a bug that never surfaces on the developer's Mac.

## 2. Core Functions

```go
filepath.Join("a", "b", "..", "c")   // "a/c" — Join cleans the result
filepath.Abs(p)
filepath.Base(p)  filepath.Dir(p)  filepath.Ext(p)
filepath.Clean(p)
filepath.Rel(base, target)
filepath.ToSlash(p)  filepath.FromSlash(p)
```

## 3. Walking

```go
filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
    if err != nil { return err }
    if d.IsDir() && d.Name() == ".git" {
        return filepath.SkipDir
    }
    return nil
})
```

`WalkDir` (Go 1.16) beats `Walk`: it uses `fs.DirEntry` and avoids a `stat` per
file.

## 4. Path Traversal Safety

```go
// UNSAFE: user input can contain ../../etc/passwd
p := filepath.Join(baseDir, userInput)

// SAFE
p := filepath.Join(baseDir, filepath.Clean("/"+userInput))
// or Go 1.24:
root, _ := os.OpenRoot(baseDir)
f, err := root.Open(userInput)      // cannot escape baseDir
```

`os.OpenRoot` (Go 1.24) is the proper fix — it enforces containment at the
syscall level.

## 5. Gotchas

- `filepath.Join` cleans, so it silently resolves `..` — that is exactly why the
  naive form above is unsafe.
- `Ext(".bashrc")` returns `".bashrc"`, not `""`.
- Symlinks are not resolved; use `filepath.EvalSymlinks` when identity matters.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
