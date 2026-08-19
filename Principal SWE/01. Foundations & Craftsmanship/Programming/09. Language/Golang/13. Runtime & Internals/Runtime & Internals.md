---
title: Runtime & Internals
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Golang]]"
---

# ⚙️ Runtime & Internals

Exhaustive, line-by-line architectural source code dive into the Go runtime (`src/runtime/`): boot sequence, GMP scheduler (`proc.go`), Garbage Collector (`mgc.go`), TCMalloc allocator (`malloc.go`), panic/defer (`panic.go`), channels (`chan.go`), timers (`time.go`), stack growth (`stack.go`), primitives (`slice.go`, `map.go`, `iface.go`), signals (`signal_unix.go`), tracing (`trace.go`), and coroutines (`coro.go`).

```text
Runtime & Internals
│
├── `01. Runtime Bootstrapping & Core Architecture`
│   ├── `Runtime Boot Sequence (rt0_go & asm_amd64.s)`
│   ├── `schedinit Implementation & Runtime Initialization`
│   ├── `The main Goroutine Lifecycle (runtime.main)`
│   ├── `Goroutine Layout & g Struct Internals`
│   ├── `OS Thread Layout & m Struct Internals`
│   ├── `Logical Processor Layout & p Struct Internals`
│   ├── `The g0 System Stack & Stack Switching`
│   └── `moduledata & Global Symbol Table Layout (symtab.go)`
├── [[GMP Scheduler Implementation (proc.go)|02. GMP Scheduler Implementation (proc.go)]]
│   ├── `Scheduler Loop (schedule() function)`
│   ├── `findrunnable() Work Search Hierarchy`
│   ├── `Work Stealing Algorithm Implementation (runqsteal)`
│   ├── `Context Switch Mechanics (gogo & mcall)`
│   ├── `Goroutine Creation Mechanics (newproc1 & newproc)`
│   ├── `M Parking, Waking & Thread Pool (startm & stopm)`
│   ├── `Syscall Handling (entersyscall & exitsyscall)`
│   ├── `Sysmon Daemon Implementation`
│   └── `Non-Cooperative Async Preemption (preemptone & asyncPreempt)`
├── [[Garbage Collector Implementation (mgc.go)|03. Garbage Collector Implementation (mgc.go)]]
│   ├── `Concurrent Tricolor Mark-Sweep Engine`
│   ├── `Hybrid Write Barrier Implementation`
│   ├── `GC Pacer Algorithm & Target Trigger Math`
│   ├── `GC Mark Termination & Sweep Termination STW Phases`
│   ├── `GC Dedicated Worker Goroutines (gcBgMarkWorker)`
│   ├── `Mutator Assist Allocations`
│   ├── `Mark Work Buffers (gcWork & wbuf)`
│   ├── `GC Phase Transitions (gcStart, gcMarkDone, gcSweep)`
│   └── `Memory Purging & Scavenging (scavenger.go)`
├── [[Memory Allocator Implementation (malloc.go)|04. Memory Allocator Implementation (malloc.go)]]
│   ├── `TCMalloc-Based Allocation Hierarchy`
│   ├── `Memory Size Classes & Tiny Allocator (<16B)`
│   ├── `Small Object Allocation Path (<32KB)`
│   ├── `Large Object Allocation Path (>32KB)`
│   ├── `mspan & Page Allocator Architecture`
│   ├── `mcentral Span Management & Cache Refilling`
│   ├── `Arena Management & Virtual Memory Mapping`
│   └── `FixAlloc Fixed-Size Allocator for Runtime Metadata`
├── [[Panic, Defer, and Exception Flow (panic.go)|05. Panic, Defer, and Exception Flow (panic.go)]]
│   ├── `_defer Struct Architecture & Linked Lists`
│   ├── `Open-Coded Defer Implementation (inline defer bits)`
│   ├── `_panic Struct Architecture & Nested Panics`
│   ├── `gopanic() Implementation Walkthrough`
│   ├── `gorecover() Implementation Walkthrough`
│   └── `Fatal Runtime Errors & Throw Mechanics (runtime.throw)`
├── [[Synchronization & Network Internals (chan.go, netpoll.go)|06. Synchronization & Network Internals (chan.go, netpoll.go)]]
│   ├── `makechan, chansend, and chanrecv Implementation (chan.go)`
│   ├── `closechan Implementation & Broadcast Signaling`
│   ├── `selectgo Multi-Channel Select Algorithm (select.go)`
│   ├── `sudog Synchronization Node Layout & Caching`
│   ├── `Netpoller Implementation (netpoll_epoll.go, netpoll_kqueue.go)`
│   ├── `Futex & OS Mutex Implementation (lock_futex.go)`
│   └── `runtime Package Diagnostic & Inspection APIs`
├── [[Runtime Timers & Sleep Implementation (time.go)|07. Runtime Timers & Sleep Implementation (time.go)]]
│   ├── `4-Heap Timer Wheel Architecture & Per-P Timers`
│   ├── `timerproc & Sysmon Timer Stealing Mechanics`
│   ├── `time.Sleep & goparkunlock Mechanics`
│   ├── `Timer Bucket Resets (time.Reset) & State Machine`
│   └── `Ticker Garbage Collection Pitfalls & Stop Cleanup`
├── [[Runtime Stack Management & Unwinding (stack.go, traceback.go)|08. Runtime Stack Management & Unwinding (stack.go, traceback.go)]]
│   ├── `Stack Growth Engine (runtime.morestack & copystack)`
│   ├── `Stack Shrinking Mechanics & GC Stack Scavenging`
│   ├── `Stack Traceback Unwinding Algorithm (gentraceback)`
│   ├── `go:nosplit Pragma & Stack Overflow Prevention`
│   └── `Stack Segment Bounds & Guard Page Protection`
├── [[Primitive Runtime Helpers (slice.go, map.go, iface.go)|09. Primitive Runtime Helpers (slice.go, map.go, iface.go)]]
│   ├── `Slice Allocation & Dynamic Growth Math (growslice in slice.go)`
│   ├── `Legacy Map (hmap & bmap) Runtime Implementation (map.go)`
│   ├── `Swiss Table Map Runtime Engine (map_swiss.go Go 1.24+)`
│   ├── `String Concatenation & Memory Conversion (string.go)`
│   ├── `Interface Construction & Assertion Helpers (convT & assertI2I)`
│   └── `Finalizer Queue & Execution Lifecycle (runtime.SetFinalizer)`
├── [[OS Signals, Exceptions & Cgo Transitions (signal_unix.go, cgocall.go)|10. OS Signals, Exceptions & Cgo Transitions (signal_unix.go, cgocall.go)]]
│   ├── `Signal Handling Architecture & gsignal Stack (signal_unix.go)`
│   ├── `SIGSEGV & SIGBUS Recovery to Runtime Panics`
│   ├── `Profiling Signal Generation (SIGPROF Sampling Engine)`
│   ├── `Cross-ABI Cgo Transition Mechanics (cgocall & cgocallback)`
│   └── `Extra OS Threads Management for Cgo (extraM)`
├── [[Runtime Tracing & Execution Profiling (trace.go, mprof.go)|11. Runtime Tracing & Execution Profiling (trace.go, mprof.go)]]
│   ├── `Trace Event Buffer Generation & Encoding (trace.go)`
│   ├── `Per-P Trace Buffers & Event Flushing Mechanics`
│   ├── `Memory Profiling Sampling Engine (mprof.go)`
│   ├── `Block & Mutex Profile Recording Internals`
│   └── `CPU Profiler Sampling Clock & Call-Stack Collection`
└── [[Coroutine Runtime & Modern Concurrency (coro.go)|12. Coroutine Runtime & Modern Concurrency (coro.go)]]
│   ├── `Coroutine Stack-Switching Engine (coro.go Go 1.23+)`
│   ├── `corostart & coroswitch Implementation Details`
│   ├── `Integrating Coroutines with GMP Scheduler`
│   └── `Zero-Allocation Push-Pull Iterator Runtime Bridges`
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 `01. Runtime Bootstrapping & Core Architecture`
- `Runtime Boot Sequence (rt0_go & asm_amd64.s)` — Hardware entry point, CPU feature detection, argc/argv extraction, and initial OS thread creation.
- `schedinit Implementation & Runtime Initialization` — Stack initialization, memory allocator setup (mallocinit), mcommoninit, gcinit, and procresize.
- `The main Goroutine Lifecycle (runtime.main)` — Spawning runtime.main, executing package init() dependency graphs, initializing sysmon, calling main.main.
- `Goroutine Layout & g Struct Internals` — Dissecting the 80+ fields of g struct: stack bounds, sched context, m pointer, atomic status, and panic list.
- `OS Thread Layout & m Struct Internals` — Dissecting m struct: g0 system stack, gsignal, curg running goroutine, p pointer, and fastrand state.
- `Logical Processor Layout & p Struct Internals` — Dissecting p struct: lock-free local runqueues (runq), mcache span caches, sudogcache, and timers.
- `The g0 System Stack & Stack Switching` — Dedicated OS-sized system stack (8MB) used for scheduler execution, runtime memory allocation, and GC.
- `moduledata & Global Symbol Table Layout (symtab.go)` — First-class binary metadata struct: pclntab line mapping, function descriptors (funcInfo), and type descriptors.
### 2. 📂 [[GMP Scheduler Implementation (proc.go)|02. GMP Scheduler Implementation (proc.go)]]
- `Scheduler Loop (schedule() function)` — Step-by-step walkthrough of schedule() in proc.go, finding runnable work, and context switching.
- `findrunnable() Work Search Hierarchy` — Checking local runqueue, global runqueue (1/61 check), netpoller, work stealing from peers, and idle parking.
- `Work Stealing Algorithm Implementation (runqsteal)` — Lock-free stealing half of peer logical processor local runqueue using atomic CAS on head/tail indices.
- `Context Switch Mechanics (gogo & mcall)` — Assembly context switching: saving callee-saved registers into g.sched and restoring target goroutine state.
- `Goroutine Creation Mechanics (newproc1 & newproc)` — Allocating new g struct from per-P gFree cache or heap, setting up entry PC, and enqueueing to runq.
- `M Parking, Waking & Thread Pool (startm & stopm)` — Spinning threads, notesleep parking, waking idle threads, and managing OS thread lifecycle.
- `Syscall Handling (entersyscall & exitsyscall)` — Detaching logical processor P from blocking OS thread M, retaking P, and thread parking on notesleep.
- `Sysmon Daemon Implementation` — Background monitoring thread: retaking stuck Ps from long syscalls, forcing periodic GC, and preemption.
- `Non-Cooperative Async Preemption (preemptone & asyncPreempt)` — Injecting SIGURG POSIX signal handler into running thread to force safe-point preemption.
### 3. 📂 [[Garbage Collector Implementation (mgc.go)|03. Garbage Collector Implementation (mgc.go)]]
- `Concurrent Tricolor Mark-Sweep Engine` — White, Grey, Black object sets, concurrent marking phases, and concurrent sweeping in mgc.go.
- `Hybrid Write Barrier Implementation` — Dijkstra insertion barrier combined with Yuasa deletion barrier eliminating Stop-The-World stack rescans.
- `GC Pacer Algorithm & Target Trigger Math` — Dynamic feedback loop calculating next_gc trigger threshold based on allocation rate, GOGC, and GOMEMLIMIT.
- `GC Mark Termination & Sweep Termination STW Phases` — Sub-millisecond STW pause windows: disabling write barriers, flushing local caches, and starting sweep.
- `GC Dedicated Worker Goroutines (gcBgMarkWorker)` — Dedicated mark workers: gcMarkWorkerDedicated, gcMarkWorkerFractional, and gcMarkWorkerIdle.
- `Mutator Assist Allocations` — Forcing fast-allocating user goroutines to assist GC marking when allocation outpaces GC throughput.
- `Mark Work Buffers (gcWork & wbuf)` — Thread-local grey object work buffers (wbuf1, wbuf2) preventing lock contention during concurrent mark.
- `GC Phase Transitions (gcStart, gcMarkDone, gcSweep)` — State machine executing STW mark preparation, concurrent mark, STW mark termination, and concurrent sweep.
- `Memory Purging & Scavenging (scavenger.go)` — Background page scavenger returning unused virtual memory spans back to the OS via madvise.
### 4. 📂 [[Memory Allocator Implementation (malloc.go)|04. Memory Allocator Implementation (malloc.go)]]
- `TCMalloc-Based Allocation Hierarchy` — Three-tier architecture: thread-local mcache, central mcentral, and global heap mheap.
- `Memory Size Classes & Tiny Allocator (<16B)` — 67 size classes, allocating objects <16B grouped into single memory blocks without individual metadata.
- `Small Object Allocation Path (<32KB)` — Fast-path lockless allocation from mcache.alloc[spanClass] without global heap locks.
- `Large Object Allocation Path (>32KB)` — Direct allocation of contiguous memory pages from global mheap page allocator.
- `mspan & Page Allocator Architecture` — Radix tree page allocator (pageAlloc) managing 8KB page chunks and span descriptors.
- `mcentral Span Management & Cache Refilling` — Central span list (nonempty / empty lists) providing spans to local mcaches under central lock.
- `Arena Management & Virtual Memory Mapping` — 64MB memory arena chunks (heapArena), virtual memory reservations, and mmap kernel mapping.
- `FixAlloc Fixed-Size Allocator for Runtime Metadata` — Low-level free-list memory allocator used exclusively for runtime internal structs (mspan, mlink).
### 5. 📂 [[Panic, Defer, and Exception Flow (panic.go)|05. Panic, Defer, and Exception Flow (panic.go)]]
- `_defer Struct Architecture & Linked Lists` — _defer struct layout, function pointers, arguments, and Goroutine _defer chain.
- `Open-Coded Defer Implementation (inline defer bits)` — Compiler optimization storing defer execution bits in integer bitmask for zero runtime allocation.
- `_panic Struct Architecture & Nested Panics` — _panic struct layout, recovered flag, aborted flag, and active panic stack unwinding.
- `gopanic() Implementation Walkthrough` — Complete line-by-line execution flow of gopanic() in panic.go traversing defer lists.
- `gorecover() Implementation Walkthrough` — Intercepting active panic, setting recovered = true, and resuming execution at defer return site.
- `Fatal Runtime Errors & Throw Mechanics (runtime.throw)` — Unrecoverable runtime aborts: printing crash dumps and calling exit(2) directly.
### 6. 📂 [[Synchronization & Network Internals (chan.go, netpoll.go)|06. Synchronization & Network Internals (chan.go, netpoll.go)]]
- `makechan, chansend, and chanrecv Implementation (chan.go)` — Complete lock acquisition, sudog enqueueing, ring-buffer copy, and direct stack transfer algorithms.
- `closechan Implementation & Broadcast Signaling` — Locking channel, releasing all waiting receivers with zero values, and panicking senders.
- `selectgo Multi-Channel Select Algorithm (select.go)` — Case shuffling (fastrand), lock ordering by channel address to prevent deadlocks, and polling.
- `sudog Synchronization Node Layout & Caching` — Synchronization waiting node representing parked goroutine on channel or sync primitive.
- `Netpoller Implementation (netpoll_epoll.go, netpoll_kqueue.go)` — Non-blocking I/O event polling, descriptor registration, and waking parked goroutines.
- `Futex & OS Mutex Implementation (lock_futex.go)` — Low-level user-space mutexes, CAS spin-wait, and falling back to Linux SYS_futex kernel sleep.
- `runtime Package Diagnostic & Inspection APIs` — runtime.ReadMemStats, runtime.Gosched, runtime.LockOSThread, runtime.GC, runtime.KeepAlive.
### 7. 📂 [[Runtime Timers & Sleep Implementation (time.go)|07. Runtime Timers & Sleep Implementation (time.go)]]
- `4-Heap Timer Wheel Architecture & Per-P Timers` — How Go 1.14+ moved timers from a single central lock to lockless per-P heaps (p.timers).
- `timerproc & Sysmon Timer Stealing Mechanics` — How idle processors and sysmon steal expired timers from busy peers and fire callbacks.
- `time.Sleep & goparkunlock Mechanics` — Parking goroutines with waitReasonSleep and awakening via timer channel ready callbacks.
- `Timer Bucket Resets (time.Reset) & State Machine` — timerModifiedEarliest, timerDeleted, timerRunning state machine transitions in runtime/time.go.
- `Ticker Garbage Collection Pitfalls & Stop Cleanup` — Why unstopped tickers leak timers in per-P heaps until manual Stop() invocation.
### 8. 📂 [[Runtime Stack Management & Unwinding (stack.go, traceback.go)|08. Runtime Stack Management & Unwinding (stack.go, traceback.go)]]
- `Stack Growth Engine (runtime.morestack & copystack)` — Allocating new contiguous stack, copying frames, and updating pointer fixup tables.
- `Stack Shrinking Mechanics & GC Stack Scavenging` — Shrinking stack to half size during GC mark termination when utilization drops below 25%.
- `Stack Traceback Unwinding Algorithm (gentraceback)` — Traversing call frames using PC/SP tables for panic dumps, stack traces, and pprof.
- `go:nosplit Pragma & Stack Overflow Prevention` — Preventing infinite morestack recursion in scheduler and allocator leaf functions via //go:nosplit.
- `Stack Segment Bounds & Guard Page Protection` — Detecting hardware stack overflows via OS memory guard pages and stack bottom boundary checks.
### 9. 📂 [[Primitive Runtime Helpers (slice.go, map.go, iface.go)|09. Primitive Runtime Helpers (slice.go, map.go, iface.go)]]
- `Slice Allocation & Dynamic Growth Math (growslice in slice.go)` — Slice growth threshold math (2x below 256 elements, 1.25x + 192 above 256) and memory class rounding.
- `Legacy Map (hmap & bmap) Runtime Implementation (map.go)` — Bucket array allocation, tophash comparison, 6.5 load factor evacuation, and overflow buckets.
- `Swiss Table Map Runtime Engine (map_swiss.go Go 1.24+)` — Go 1.24+ Swiss Table layout: control bytes, 16-element groups, SIMD vector matching, and slot arrays.
- `String Concatenation & Memory Conversion (string.go)` — concatstrings algorithm, small string optimizations, and runtime byte-to-string casting.
- `Interface Construction & Assertion Helpers (convT & assertI2I)` — convT scalar boxing, type assertion checks (assertI2I), and dynamic itab caching (itabAdd).
- `Finalizer Queue & Execution Lifecycle (runtime.SetFinalizer)` — Special finalizer block allocation (fintab), special records, and the dedicated finq execution goroutine.
### 10. 📂 [[OS Signals, Exceptions & Cgo Transitions (signal_unix.go, cgocall.go)|10. OS Signals, Exceptions & Cgo Transitions (signal_unix.go, cgocall.go)]]
- `Signal Handling Architecture & gsignal Stack (signal_unix.go)` — Registering POSIX signals (initsig), dedicated signal stack (gsignal), and sighandler execution.
- `SIGSEGV & SIGBUS Recovery to Runtime Panics` — Catching hardware null-pointer dereferences and memory faults and converting them into catchable panics.
- `Profiling Signal Generation (SIGPROF Sampling Engine)` — OS interval timer delivering SIGPROF (100Hz) to capture thread program counters into trace buffers.
- `Cross-ABI Cgo Transition Mechanics (cgocall & cgocallback)` — Switching from Go stack to OS C thread stack, saving registers, calling C, and transitioning back.
- `Extra OS Threads Management for Cgo (extraM)` — Managing background OS threads (extraM) to handle incoming callbacks from C into Go.
### 11. 📂 [[Runtime Tracing & Execution Profiling (trace.go, mprof.go)|11. Runtime Tracing & Execution Profiling (trace.go, mprof.go)]]
- `Trace Event Buffer Generation & Encoding (trace.go)` — Binary execution trace format, event types (proc create, goroutine switch, syscall, GC phases).
- `Per-P Trace Buffers & Event Flushing Mechanics` — Lock-free event recording into local per-P byte buffers and flushing to global trace consumer.
- `Memory Profiling Sampling Engine (mprof.go)` — Poisson sampling distribution (1-in-512KB default) recording allocation call-stacks into mprof bucket tables.
- `Block & Mutex Profile Recording Internals` — Capturing stack traces and blocking durations for channel stalls and mutex contention events.
- `CPU Profiler Sampling Clock & Call-Stack Collection` — Extracting active goroutine call-stack on every SIGPROF tick and writing to pprof profile buffers.
### 12. 📂 [[Coroutine Runtime & Modern Concurrency (coro.go)|12. Coroutine Runtime & Modern Concurrency (coro.go)]]
- `Coroutine Stack-Switching Engine (coro.go Go 1.23+)` — Go 1.23+ lightweight asymmetric coroutine runtime implementation powering iter.Pull.
- `corostart & coroswitch Implementation Details` — Creating paired coroutine execution contexts and switching execution control between caller and callee.
- `Integrating Coroutines with GMP Scheduler` — How coroutines cooperate with goroutines, stack copying, and preemptible scheduler safepoints.
- `Zero-Allocation Push-Pull Iterator Runtime Bridges` — Transforming push-based yield functions into stateful pull iterators without heap allocations.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

