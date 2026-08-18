---
title: I-O, OS & System
tags:
  - golang
  - stdlib
  - principal-swe
parent: "[[Standard Library Mastery]]"
---

# I-O, OS & System

Stream processing with io.Reader/Writer, buffered I/O, OS interactions, flags, and embedded assets.

```text
I-O, OS & System
│
├── [[io Package (Reader, Writer, Closer)]]
├── [[bufio Package]]
├── [[os Package]]
├── [[path-filepath Package]]
├── [[flag Package]]
├── [[embed Standard Package & Static Assets]]
└── [[io-fs Virtual Filesystem]]
```

---

## 🗂️ Topics

- [[io Package (Reader, Writer, Closer)]] — Streaming abstractions, io.Copy, io.Pipe, io.MultiReader, io.TeeReader, io.LimitReader.
- [[bufio Package]] — bufio.Reader, bufio.Writer, bufio.Scanner for high-throughput buffered stream processing.
- [[os Package]] — Environment variables, process management, POSIX signal handling, exit codes, os.File operations.
- [[path-filepath Package]] — Cross-platform file path manipulation, filepath.Walk, filepath.Clean, filepath.Join.
- [[flag Package]] — Command-line argument parsing and custom flag value types.
- [[embed Standard Package & Static Assets]] — Embedding static assets, templates, and directory trees directly into binary executables.
- [[io-fs Virtual Filesystem]] — fs.FS, fs.File, fs.WalkDir virtual filesystem interface abstraction.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]
- 🎓 Root: [[Principal SWE]]
