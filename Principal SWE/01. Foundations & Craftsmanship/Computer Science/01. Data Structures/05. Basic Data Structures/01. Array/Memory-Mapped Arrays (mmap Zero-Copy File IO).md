---
title: "Memory-Mapped Arrays (mmap Zero-Copy File IO)"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Memory-Mapped Arrays (mmap Zero-Copy File IO)

## 1. Definition
A **Memory-Mapped Array** uses the POSIX `mmap` system call to map a disk file directly into the application's virtual address space.
The OS kernel handles page caching and demand paging transparently, enabling **zero-copy random access** across multi-terabyte arrays without loading the file entirely into RAM.

---

## 2. Mental Model
```text
Virtual Memory Address Space:
[ 0x7FFF0000 ───────────── mmap Address Space ───────────── 0x7FFFFFFF ]
      │                                                           │
      ▼ (Page Fault on demand)                                    ▼
[ OS Page Cache (RAM) ] <════ Zero-Copy DMA ════> [ NVMe SSD / Disk File ]
```

---

## 3. Usage
```go
// Reading a multi-gigabyte file as an array via mmap in Go
package main

import (
    "os"
    "golang.org/x/sys/unix"
)

func OpenMmapArray(filename string) ([]byte, error) {
    f, err := os.Open(filename)
    if err != nil { return nil, err }
    defer f.Close()
    
    info, _ := f.Stat()
    data, err := unix.Mmap(int(f.Fd()), 0, int(info.Size()), unix.PROT_READ, unix.MAP_SHARED)
    return data, err // Direct array access across multi-gigabyte files!
}
```

---

## 4. Gotchas
- **SIGBUS on Truncation:** If another process truncates the backing file while your application is reading memory-mapped addresses, the CPU triggers an immediate uncatchable `SIGBUS` crash.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

