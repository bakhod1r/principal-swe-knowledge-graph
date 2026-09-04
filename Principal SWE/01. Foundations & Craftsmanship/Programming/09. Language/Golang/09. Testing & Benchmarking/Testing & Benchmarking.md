---
title: Testing & Benchmarking
tags:
  - golang
  - testing
  - principal-swe
parent: "[[Golang]]"
---

# 🧪 Testing & Benchmarking

Unit testing primitives, test doubles, Testcontainers integration, benchmarking allocation profiles, native fuzzing, mutation testing, deterministic virtual time (synctest), and race detection.

```text
Testing & Benchmarking
│
├── [[Unit Testing & Test Framework Primitives|01. Unit Testing & Test Framework Primitives]]
│   ├── `testing.T Lifecycle & Fatal vs Error Reporting`
│   ├── `Table-Driven Tests Architecture (Idiomatic Pattern)`
│   ├── `Subtests (t.Run) Hierarchical Execution & Filtering`
│   ├── `Test Helpers (t.Helper) Stack Frame Stripping`
│   ├── `Parallel Tests (t.Parallel) Race Isolation`
│   ├── `t.Cleanup Reliable Resource Management`
│   ├── `TestMain Function Lifecycle (Global Setup & Teardown)`
│   └── `t.Setenv Environment Isolation`
├── [[Mocks, Stubs, Fakes & Contract Testing|02. Mocks, Stubs, Fakes & Contract Testing]]
│   ├── `Test Doubles Taxonomy (Dummies, Stubs, Fakes, Spies, Mocks)`
│   ├── `Hand-Written In-Memory Fakes (Zero-Dependency Testing)`
│   ├── `Mock Generators (mockery, gomock, moq)`
│   ├── `Mocking What You Do Not Own Anti-Pattern in Testing`
│   └── `Consumer-Driven Contract Testing (Pact in Go)`
├── [[HTTP, Integration & Database Testing|03. HTTP, Integration & Database Testing]]
│   ├── `httptest Package (ResponseRecorder & NewServer)`
│   ├── `Golden Files Snapshot Testing (-update flag)`
│   ├── `Testcontainers-go (Real Postgres, Redis, Kafka)`
│   ├── `Database Fixtures & Migration Lifecycle in Tests`
│   └── `Testing gRPC Services (bufconn In-Memory Listeners)`
├── [[Benchmarking, Allocation Profiling & benchstat|04. Benchmarking, Allocation Profiling & benchstat]]
│   ├── `testing.B Microbenchmarks (b.N & Iteration Scaling)`
│   ├── `Allocation Profiling in Benchmarks (-benchmem & b.ReportAllocs)`
│   ├── `Parallel Benchmarks (b.RunParallel & PB)`
│   ├── `Statistical Analysis with benchstat`
│   └── `Benchmarking Traps (Compiler Inlining & Dead Code Elimination)`
├── [[Fuzzing, Mutation & Property-Based Testing|05. Fuzzing, Mutation & Property-Based Testing]]
│   ├── `Native Go Fuzz Testing (testing.F Mutational Engine)`
│   ├── `Property-Based Testing (testing-quick & gopter)`
│   ├── `Mutation Testing in Go (go-mutesting)`
│   └── `Fault Injection & Chaos Testing in Go`
└── [[Concurrency Testing & Deterministic Virtual Time|06. Concurrency Testing & Deterministic Virtual Time]]
│   ├── `Race Detector (-race ThreadSanitizer Mechanics)`
│   ├── `Deterministic Virtual Time Testing (synctest Go 1.24+)`
│   ├── `Testing Goroutine Leaks (goleak)`
│   ├── `Deadlock & Starvation Testing in Concurrent Code`
│   └── `Code Coverage Profiling & Quality Gates (-cover, -coverprofile)`
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Unit Testing & Test Framework Primitives|01. Unit Testing & Test Framework Primitives]]
- [[testing.T Lifecycle & Fatal vs Error Reporting]] — t.Fail, t.FailNow, t.Fatal, t.Error, and t.Log execution semantics and failure propagation.
- [[Table-Driven Tests Architecture (Idiomatic Pattern)]] — Slices of anonymous structs, subtest encapsulation, and clear descriptive assertion failures.
- [[Subtests (t.Run) Hierarchical Execution & Filtering]] — Running specific subtests via regex (-run=TestRoot/SubCase) and test failure isolation.
- [[Test Helpers (t.Helper) Stack Frame Stripping]] — Marking helper functions to strip wrapper lines from assertion failure stack traces.
- [[Parallel Tests (t.Parallel) Race Isolation]] — Pausing and resuming tests concurrently to catch race conditions and shorten CI wall time.
- [[t.Cleanup Reliable Resource Management]] — Registering reliable LIFO cleanup callbacks replacing fragile defer statements in tests.
- [[TestMain Function Lifecycle (Global Setup & Teardown)]] — Global test suite setup, teardown hooks, flag parsing (flag.Parse), and m.Run() execution.
- [[t.Setenv Environment Isolation]] — Safely mutating environment variables per-test with automatic post-test restoration.
### 2. 📂 [[Mocks, Stubs, Fakes & Contract Testing|02. Mocks, Stubs, Fakes & Contract Testing]]
- [[Test Doubles Taxonomy (Dummies, Stubs, Fakes, Spies, Mocks)]] — Architectural definitions of test doubles and when to choose fakes vs generated mocks.
- [[Hand-Written In-Memory Fakes (Zero-Dependency Testing)]] — Writing high-performance in-memory repository fakes for deterministic unit testing.
- [[Mock Generators (mockery, gomock, moq)]] — Automated interface mock generation, call expectations, argument matchers, and assertions.
- [[Mocking What You Do Not Own Anti-Pattern in Testing]] — Why third-party libraries should never be mocked directly and must be wrapped in domain interfaces.
- [[Consumer-Driven Contract Testing (Pact in Go)]] — Validating API schemas, contracts, and backwards compatibility across distributed microservices.
### 3. 📂 [[HTTP, Integration & Database Testing|03. HTTP, Integration & Database Testing]]
- [[httptest Package (ResponseRecorder & NewServer)]] — Testing HTTP handlers, routing, and middlewares without opening real OS network ports.
- [[Golden Files Snapshot Testing (-update flag)]] — Validating complex JSON, HTML, and binary payloads against tracked golden snapshot files.
- [[Testcontainers-go (Real Postgres, Redis, Kafka)]] — Spinning up real Dockerized database and message broker containers in integration tests.
- [[Database Fixtures & Migration Lifecycle in Tests]] — Setting up clean transactional test databases, seeding fixtures, and rolling back migrations.
- [[Testing gRPC Services (bufconn In-Memory Listeners)]] — In-memory gRPC client-server integration testing without TCP network socket overhead.
### 4. 📂 [[Benchmarking, Allocation Profiling & benchstat|04. Benchmarking, Allocation Profiling & benchstat]]
- [[testing.B Microbenchmarks (b.N & Iteration Scaling)]] — Loop mechanics, b.N scaling, b.ResetTimer(), and avoiding compiler loop optimization optimizations.
- [[Allocation Profiling in Benchmarks (-benchmem & b.ReportAllocs)]] — Tracking heap allocations per operation (allocs/op and B/op) to prevent performance regressions.
- [[Parallel Benchmarks (b.RunParallel & PB)]] — Measuring throughput under multi-threaded concurrency load across all logical CPU cores.
- [[Statistical Analysis with benchstat]] — Comparing benchmark results before and after code changes with p-value statistical confidence.
- [[Benchmarking Traps (Compiler Inlining & Dead Code Elimination)]] — Forcing benchmark results into package-level sinks (var Sink T) to prevent dead-code removal.
### 5. 📂 [[Fuzzing, Mutation & Property-Based Testing|05. Fuzzing, Mutation & Property-Based Testing]]
- [[Native Go Fuzz Testing (testing.F Mutational Engine)]] — Corpus management, seed corpus (testdata/fuzz), and automated mutational fuzzing in CI.
- [[Property-Based Testing (testing-quick & gopter)]] — Generating randomized inputs to verify mathematical invariants, commutativity, and idempotency laws.
- [[Mutation Testing in Go (go-mutesting)]] — Injecting AST code mutations (swapping operators, altering branches) to evaluate test suite quality.
- [[Fault Injection & Chaos Testing in Go]] — Injecting artificial network latency, packet loss, and database errors into test pipelines.
### 6. 📂 [[Concurrency Testing & Deterministic Virtual Time|06. Concurrency Testing & Deterministic Virtual Time]]
- [[Race Detector (-race ThreadSanitizer Mechanics)]] — Runtime data race detection, 10x CPU / 20x memory overhead, and CI pipeline integration.
- [[Deterministic Virtual Time Testing (synctest Go 1.24+)]] — Testing concurrent goroutines and timers instantly without time.Sleep using virtual time bubbles.
- [[Testing Goroutine Leaks (goleak)]] — Uber goleak library verifying that no background goroutines leaked after test suite execution.
- [[Deadlock & Starvation Testing in Concurrent Code]] — Simulating high lock contention and verifying timeout-based deadlock recovery in tests.
- [[Code Coverage Profiling & Quality Gates (-cover, -coverprofile)]] — Generating HTML coverage reports, filtering generated code, and enforcing CI coverage thresholds.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

