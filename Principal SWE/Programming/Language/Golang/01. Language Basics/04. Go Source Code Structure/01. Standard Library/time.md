---
title: time Package
tags:
  - golang
  - standard-library
  - time
parent: "[[Standard Library]]"
---

# `time`

Instants, durations, monotonic clocks, timers, and formatting.

## 1. Durations Are Typed

```go
time.Sleep(2 * time.Second)
d := 150 * time.Millisecond
timeout := time.Duration(n) * time.Second     // n is a variable → conversion needed
```

`time.Duration` is `int64` nanoseconds. `time.Sleep(2)` sleeps 2 **nanoseconds** —
a silent bug the type system permits.

## 2. The Reference-Layout Format

```go
t.Format("2006-01-02 15:04:05")
t.Format(time.RFC3339)
time.Parse(time.RFC3339, s)
```

The layout is the fixed instant `Mon Jan 2 15:04:05 MST 2006` — i.e. `1 2 3 4 5 6
-7`. Not `YYYY-MM-DD`.

## 3. Monotonic vs Wall Clock

```go
start := time.Now()
elapsed := time.Since(start)     // uses the monotonic reading — NTP-safe
```

`time.Now()` carries both readings. Subtraction uses the monotonic one, so
elapsed time is correct even if the wall clock jumps. But:

```go
t.Round(0)          // strips the monotonic reading
t.UTC()             // also strips it
```

Once stripped, durations become NTP-sensitive.

## 4. Timers and Tickers

```go
select {
case <-ctx.Done():
    return ctx.Err()
case <-time.After(5 * time.Second):     // leaks the timer until it fires
}

tk := time.NewTicker(time.Second)
defer tk.Stop()                          // ALWAYS — tickers never self-collect
```

Go 1.23 made unreferenced `Timer`s garbage-collectable, but `Ticker.Stop` is still
required.

## 5. Gotchas

- `==` on `time.Time` compares monotonic readings and location too — use
  `t.Equal(u)`.
- A `scratch` container has no tzdata; `time.LoadLocation` fails. Import
  `_ "time/tzdata"` — see `lib`.
- Storing `time.Time` in a struct copies 24 bytes; for large records store Unix
  nanos.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
