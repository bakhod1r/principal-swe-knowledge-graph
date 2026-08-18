---
title: os/exec Package
tags:
  - golang
  - standard-library
  - process
  - io
parent: "[[Standard Library]]"
---

# `os/exec`

Running external commands.

## 1. The Correct Form

```go
cmd := exec.CommandContext(ctx, "git", "rev-parse", "HEAD")
cmd.Dir = repoPath
out, err := cmd.Output()
```

`CommandContext` kills the process when the context is cancelled — without it a
hung subprocess outlives your request. See `context`.

## 2. There Is No Shell

```go
exec.Command("ls -la")                      // ✗ looks for a binary named "ls -la"
exec.Command("ls", "-la")                   // ✓
exec.Command("sh", "-c", "ls -la | wc -l")  // ✓ explicit shell when you need pipes
```

Arguments are passed directly to `execve`. **This is a security feature** — no
shell means no shell injection. Only reach for `sh -c` deliberately.

## 3. Capturing Output

```go
out, err := cmd.Output()          // stdout; stderr lands in ExitError.Stderr
combined, err := cmd.CombinedOutput()

var stdout, stderr bytes.Buffer
cmd.Stdout, cmd.Stderr = &stdout, &stderr
err := cmd.Run()
```

## 4. Exit Codes

```go
if err := cmd.Run(); err != nil {
    var ee *exec.ExitError
    if errors.As(err, &ee) {
        log.Printf("exit %d: %s", ee.ExitCode(), ee.Stderr)
    }
}
```

See `errors`.

## 5. Gotchas

- `cmd.Wait` must not be called concurrently with reading `StdoutPipe`; read
  fully, **then** `Wait`, or it deadlocks.
- `exec.LookPath` respects `PATH` — in a container with a minimal image the
  binary may simply not be there.
- Since Go 1.19 a relative path found in the current directory is rejected
  (`ErrDot`), closing a Windows-style hijack.
- A killed context leaves the process's children orphaned unless you set a
  process group.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
