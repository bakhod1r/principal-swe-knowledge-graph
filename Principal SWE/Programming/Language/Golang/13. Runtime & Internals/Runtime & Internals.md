---
title: Runtime & Internals
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Golang]]"
---

# ⚙️ Runtime & Internals

Deep architectural source code dive into Go runtime internals: boot sequence, GMP scheduler (proc.go), Garbage Collector (mgc.go), TCMalloc allocator (malloc.go), panic/defer (panic.go), and channels (chan.go).

```text
Runtime & Internals
│
├── [[Runtime Bootstrapping & Core Architecture|01. Runtime Bootstrapping & Core Architecture]]
│   ├── [[Runtime Boot Sequence (rt0_go & asm_amd64.s)]]
│   ├── [[schedinit Implementation & Runtime Initialization]]
│   ├── [[The main Goroutine Lifecycle (runtime.main)]]
│   ├── [[Goroutine Layout & g Struct Internals]]
│   ├── [[OS Thread Layout & m Struct Internals]]
│   ├── [[Logical Processor Layout & p Struct Internals]]
│   └── [[The g0 System Stack & Stack Switching]]
├── [[GMP Scheduler Implementation (proc.go)|02. GMP Scheduler Implementation (proc.go)]]
│   ├── [[Scheduler Loop (schedule() function)]]
│   ├── [[findrunnable() Work Search Hierarchy]]
│   ├── [[Work Stealing Algorithm Implementation (runqsteal)]]
│   ├── [[Context Switch Mechanics (gogo & mcall)]]
│   ├── [[Syscall Handling (entersyscall & exitsyscall)]]
│   ├── [[Sysmon Daemon Implementation]]
│   └── [[Non-Cooperative Async Preemption (preemptone & asyncPreempt)]]
├── [[Garbage Collector Implementation (mgc.go)|03. Garbage Collector Implementation (mgc.go)]]
│   ├── [[Concurrent Tricolor Mark-Sweep Engine]]
│   ├── [[Hybrid Write Barrier Implementation]]
│   ├── [[GC Pacer Algorithm & Target Trigger Math]]
│   ├── [[GC Mark Termination & Sweep Termination STW Phases]]
│   ├── [[GC Dedicated Worker Goroutines (gcBgMarkWorker)]]
│   ├── [[Mutator Assist Allocations]]
│   └── [[Memory Purging & Scavenging (scavenger.go)]]
├── [[Memory Allocator Implementation (malloc.go)|04. Memory Allocator Implementation (malloc.go)]]
│   ├── [[TCMalloc-Based Allocation Hierarchy]]
│   ├── [[Memory Size Classes & Tiny Allocator (<16B)]]
│   ├── [[Small Object Allocation Path (<32KB)]]
│   ├── [[Large Object Allocation Path (>32KB)]]
│   ├── [[mspan & Page Allocator Architecture]]
│   └── [[Arena Management & Virtual Memory Mapping]]
├── [[Panic, Defer, and Exception Flow (panic.go)|05. Panic, Defer, and Exception Flow (panic.go)]]
│   ├── [[_defer Struct Architecture & Linked Lists]]
│   ├── [[Open-Coded Defer Implementation (inline defer bits)]]
│   ├── [[_panic Struct Architecture & Nested Panics]]
│   ├── [[gopanic() Implementation Walkthrough]]
│   └── [[gorecover() Implementation Walkthrough]]
└── [[Synchronization & Network Internals (chan.go, netpoll.go)|06. Synchronization & Network Internals (chan.go, netpoll.go)]]
│   ├── [[makechan, chansend, and chanrecv Implementation (chan.go)]]
│   ├── [[closechan Implementation & Broadcast Signaling]]
│   ├── [[Netpoller Implementation (netpoll_epoll.go, netpoll_kqueue.go)]]
│   ├── [[Futex & OS Mutex Implementation (lock_futex.go)]]
│   └── [[runtime Package Diagnostic & Inspection APIs]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Runtime Bootstrapping & Core Architecture|01. Runtime Bootstrapping & Core Architecture]]
- [[Runtime Boot Sequence (rt0_go & asm_amd64.s)]] — Hardware entry point, CPU feature detection, argc/argv extraction, and initial OS thread creation.
- [[schedinit Implementation & Runtime Initialization]] — Stack initialization, memory allocator setup (mallocinit), mcommoninit, gcinit, and procresize.
- [[The main Goroutine Lifecycle (runtime.main)]] — Spawning runtime.main, executing package init() dependency graphs, initializing sysmon, calling main.main.
- [[Goroutine Layout & g Struct Internals]] — Dissecting the 80+ fields of g struct: stack bounds, sched context, m pointer, atomic status, and panic list.
- [[OS Thread Layout & m Struct Internals]] — Dissecting m struct: g0 system stack, gsignal, curg running goroutine, p pointer, and fastrand state.
- [[Logical Processor Layout & p Struct Internals]] — Dissecting p struct: lock-free local runqueues (runq), mcache span caches, sudogcache, and timers.
- [[The g0 System Stack & Stack Switching]] — Dedicated OS-sized system stack (8MB) used for scheduler execution, runtime memory allocation, and GC.
### 2. 📂 [[GMP Scheduler Implementation (proc.go)|02. GMP Scheduler Implementation (proc.go)]]
- [[Scheduler Loop (schedule() function)]] — Step-by-step walkthrough of schedule() in proc.go, finding runnable work, and context switching.
- [[findrunnable() Work Search Hierarchy]] — Checking local runqueue, global runqueue (1/61 check), netpoller, work stealing from peers, and idle parking.
- [[Work Stealing Algorithm Implementation (runqsteal)]] — Lock-free stealing half of peer logical processor local runqueue using atomic CAS on head/tail indices.
- [[Context Switch Mechanics (gogo & mcall)]] — Assembly context switching: saving callee-saved registers into g.sched and restoring target goroutine state.
- [[Syscall Handling (entersyscall & exitsyscall)]] — Detaching logical processor P from blocking OS thread M, retaking P, and thread parking on notesleep.
- [[Sysmon Daemon Implementation]] — Background monitoring thread: retaking stuck Ps from long syscalls, forcing periodic GC, and preemption.
- [[Non-Cooperative Async Preemption (preemptone & asyncPreempt)]] — Injecting SIGURG POSIX signal handler into running thread to force safe-point preemption.
### 3. 📂 [[Garbage Collector Implementation (mgc.go)|03. Garbage Collector Implementation (mgc.go)]]
- [[Concurrent Tricolor Mark-Sweep Engine]] — White, Grey, Black object sets, concurrent marking phases, and concurrent sweeping in mgc.go.
- [[Hybrid Write Barrier Implementation]] — Dijkstra insertion barrier combined with Yuasa deletion barrier eliminating Stop-The-World stack rescans.
- [[GC Pacer Algorithm & Target Trigger Math]] — Dynamic feedback loop calculating next_gc trigger threshold based on allocation rate, GOGC, and GOMEMLIMIT.
- [[GC Mark Termination & Sweep Termination STW Phases]] — Sub-millisecond STW pause windows: disabling write barriers, flushing local caches, and starting sweep.
- [[GC Dedicated Worker Goroutines (gcBgMarkWorker)]] — Dedicated mark workers: gcMarkWorkerDedicated, gcMarkWorkerFractional, and gcMarkWorkerIdle.
- [[Mutator Assist Allocations]] — Forcing fast-allocating user goroutines to assist GC marking when allocation outpaces GC throughput.
- [[Memory Purging & Scavenging (scavenger.go)]] — Background page scavenger returning unused virtual memory spans back to the OS via madvise.
### 4. 📂 [[Memory Allocator Implementation (malloc.go)|04. Memory Allocator Implementation (malloc.go)]]
- [[TCMalloc-Based Allocation Hierarchy]] — Three-tier architecture: thread-local mcache, central mcentral, and global heap mheap.
- [[Memory Size Classes & Tiny Allocator (<16B)]] — 67 size classes, allocating objects <16B grouped into single memory blocks without individual metadata.
- [[Small Object Allocation Path (<32KB)]] — Fast-path lockless allocation from mcache.alloc[spanClass] without global heap locks.
- [[Large Object Allocation Path (>32KB)]] — Direct allocation of contiguous memory pages from global mheap page allocator.
- [[mspan & Page Allocator Architecture]] — Radix tree page allocator (pageAlloc) managing 8KB page chunks and span descriptors.
- [[Arena Management & Virtual Memory Mapping]] — 64MB memory arena chunks (heapArena), virtual memory reservations, and mmap kernel mapping.
### 5. 📂 [[Panic, Defer, and Exception Flow (panic.go)|05. Panic, Defer, and Exception Flow (panic.go)]]
- [[_defer Struct Architecture & Linked Lists]] — _defer struct layout, function pointers, arguments, and Goroutine _defer chain.
- [[Open-Coded Defer Implementation (inline defer bits)]] — Compiler optimization storing defer execution bits in integer bitmask for zero runtime allocation.
- [[_panic Struct Architecture & Nested Panics]] — _panic struct layout, recovered flag, aborted flag, and active panic stack unwinding.
- [[gopanic() Implementation Walkthrough]] — Complete line-by-line execution flow of gopanic() in panic.go traversing defer lists.
- [[gorecover() Implementation Walkthrough]] — Intercepting active panic, setting recovered = true, and resuming execution at defer return site.
### 6. 📂 [[Synchronization & Network Internals (chan.go, netpoll.go)|06. Synchronization & Network Internals (chan.go, netpoll.go)]]
- [[makechan, chansend, and chanrecv Implementation (chan.go)]] — Complete lock acquisition, sudog enqueueing, ring-buffer copy, and direct stack transfer algorithms.
- [[closechan Implementation & Broadcast Signaling]] — Locking channel, releasing all waiting receivers with zero values, and panicking senders.
- [[Netpoller Implementation (netpoll_epoll.go, netpoll_kqueue.go)]] — Non-blocking I/O event polling, descriptor registration, and waking parked goroutines.
- [[Futex & OS Mutex Implementation (lock_futex.go)]] — Low-level user-space mutexes, CAS spin-wait, and falling back to Linux SYS_futex kernel sleep.
- [[runtime Package Diagnostic & Inspection APIs]] — runtime.ReadMemStats, runtime.Gosched, runtime.LockOSThread, runtime.GC, runtime.KeepAlive.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
