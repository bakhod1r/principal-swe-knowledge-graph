---
title: log/slog Package
tags:
  - golang
  - standard-library
  - logging
  - observability
parent: "[[Standard Library]]"
---

# `log/slog`

**Go 1.21.** Structured, levelled logging in the standard library.

## 1. Basic Use

```go
slog.Info("request handled", "method", "GET", "path", p, "ms", d.Milliseconds())
slog.Error("query failed", "err", err, "table", "users")
```

Alternating key/value pairs. The typed form avoids allocation:

```go
slog.LogAttrs(ctx, slog.LevelInfo, "request handled",
    slog.String("method", "GET"),
    slog.Int64("ms", d.Milliseconds()),
)
```

## 2. Handlers

```go
h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level:     slog.LevelDebug,
    AddSource: true,
})
slog.SetDefault(slog.New(h))
```

`TextHandler` for humans, `JSONHandler` for log aggregators. A custom `Handler`
is the extension point — that is the whole design.

## 3. Context and Grouping

```go
log := slog.With("request_id", id, "user", uid)   // child logger, fields inherited
log.Info("started")

slog.InfoContext(ctx, "done")                      // handler can read ctx
```

`With` is how you avoid repeating correlation fields on every call.

## 4. Levels

`LevelDebug` (-4), `LevelInfo` (0), `LevelWarn` (4), `LevelError` (8). Integers,
so custom levels fit between them:

```go
const LevelTrace = slog.Level(-8)
```

## 5. Gotchas

- An odd number of key/value args produces a `!BADKEY` entry rather than a
  compile error — `LogAttrs` avoids this class of bug entirely.
- The old `log` package still writes to stderr unstructured; `slog` does not
  replace it automatically unless you call `slog.SetDefault`.
- Handlers must be safe for concurrent use; the built-in ones are.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
