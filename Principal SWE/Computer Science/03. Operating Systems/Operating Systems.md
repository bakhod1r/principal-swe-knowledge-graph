---
title: Operating Systems
tags:
  - computer-science
  - operating-systems
  - principal-swe
parent: "[[Computer Science]]"
---

# 🏛️ Operating Systems (Foundations & Systems Architecture)

Kernel architectures, thread scheduling internals (CFS), virtual memory page replacement, IPC channels, epoll/io_uring I/O models, syscall traps, real-time deadlines, and hypervisor virtualization.

```text
Operating Systems
│
├── [[Kernel Cpu Scheduler Internals|01. Scheduler Internals]]
├── [[Virtual Memory and Page Table Management|02. Virtual Memory Internals]]
├── [[Page Replacement and Eviction Algorithms|03. Page Replacement Algorithms]]
├── [[Interprocess Communication (ipc)|04. Interprocess Communication]]
├── [[High Performance I-o Models (epoll and IO Uring)|05. IO Models epoll and IO Uring]]
├── [[Context Switching and System Call Execution|06. Context Switching and Syscalls]]
├── [[Hardware Interrupts, Exceptions, and Traps|07. Interrupts and Traps]]
├── [[Real Time Operating System Scheduling|08. Real Time Scheduling]]
└── [[Hardware Virtualization and Hypervisors|09. Virtualization and Hypervisors]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Kernel Cpu Scheduler Internals|01. Scheduler Internals]] — Linux Completely Fair Scheduler (CFS), vruntime tracking, red-black tree task queues, and multi-core load balancing.
- 📂 [[Virtual Memory and Page Table Management|02. Virtual Memory Internals]] — Address space layout, demand paging, anonymous vs file-backed pages, swap, and OOM killer heuristics.
- 📂 [[Page Replacement and Eviction Algorithms|03. Page Replacement Algorithms]] — LRU, Clock algorithm, CLOCK-Pro, 2Q, ARC (Adaptive Replacement Cache), and page fault handling.
- 📂 [[Interprocess Communication (ipc)|04. Interprocess Communication]] — Pipes, UNIX domain sockets, shared memory (shm_open), message queues, and lockless ring buffers.
- 📂 [[High Performance I-o Models (epoll and IO Uring)|05. IO Models epoll and IO Uring]] — Edge-triggered vs level-triggered epoll, zero-copy socket I/O, submission/completion queues in io_uring.
- 📂 [[Context Switching and System Call Execution|06. Context Switching and Syscalls]] — User-space to kernel-space transition, syscall trap tables, TSS register swap, and VDSO fast syscalls.
- 📂 [[Hardware Interrupts, Exceptions, and Traps|07. Interrupts and Traps]] — Interrupt Descriptor Table (IDT), Top-half vs Bottom-half ISR processing, tasklets, and softirqs.
- 📂 [[Real Time Operating System Scheduling|08. Real Time Scheduling]] — Rate Monotonic Scheduling (RMS), Earliest Deadline First (EDF), and Priority Inversion mitigation (Priority Ceiling).
- 📂 [[Hardware Virtualization and Hypervisors|09. Virtualization and Hypervisors]] — Type-1 vs Type-2 hypervisors, Intel VT-x/AMD-V hardware virtualization, EPT page tables, and KVM/QEMU.

---

## 🔗 References
- ⬆️ Parent: [[Computer Science]]
- 🎓 Root: [[Principal SWE]]
