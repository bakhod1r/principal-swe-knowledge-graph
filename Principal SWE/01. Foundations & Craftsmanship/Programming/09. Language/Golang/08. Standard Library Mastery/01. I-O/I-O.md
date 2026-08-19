---
title: I-O, OS & Virtual Filesystems
tags:
  - golang
  - stdlib
  - principal-swe
parent: "[[Standard Library Mastery]]"
---

# I-O, OS & Virtual Filesystems

io.Reader/Writer abstractions, io/fs virtual filesystem, OS process management, bufio streaming, embed static assets, and path handling.

```text
I-O, OS & Virtual Filesystems
│
├── [[io Package Architecture (Reader, Writer, Closer)]]
├── `io-fs Virtual Filesystem & Directory Sandboxing`
├── `os Package & Process Management`
├── [[bufio High-Performance Stream Buffering]]
├── `embed Standard Package & Static Assets`
├── `flag & pflag Command-Line Parsing`
└── `path-filepath Cross-Platform Path Handling`
```

---

## 🗂️ Topics

- [[io Package Architecture (Reader, Writer, Closer)]] — Core streaming interfaces: io.Reader, io.Writer, io.Closer, io.TeeReader, io.Pipe, and io.MultiWriter.
- `io-fs Virtual Filesystem & Directory Sandboxing` — Go 1.16+ read-only virtual filesystem abstraction (fs.FS), fs.WalkDir, and directory security.
- `os Package & Process Management` — File operations, environment variables, process spawning (os.StartProcess, exec.Command), and signal forwarding.
- [[bufio High-Performance Stream Buffering]] — Minimizing OS read/write syscall overhead with bufio.Reader, bufio.Writer, and custom Scanner splits.
- `embed Standard Package & Static Assets` — Embedding static templates, HTML/CSS assets, and database migrations directly into compiled Go binaries.
- `flag & pflag Command-Line Parsing` — Command-line flag parsing architecture, custom flag.Value interfaces, and POSIX-compliant pflag integration.
- `path-filepath Cross-Platform Path Handling` — Cross-platform file path manipulation, filepath.Clean, filepath.Rel, and path traversal defense.

- [[archive-tar & archive-zip Streaming Extraction]] — Safe archive decompression avoiding Zip Bomb memory exhaustion and Zip Slip directory traversal attacks.
- [[compress-gzip & flate High-Throughput Stream Compression]] — Optimizing compression levels, buffer pooling with sync.Pool, and low-allocation stream compression.
- `syscall & golang.org-x-sys-unix Low-Level System Calls` — Direct kernel system calls, raw socket options, epoll/kqueue descriptors, and flock advisory locking.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]

