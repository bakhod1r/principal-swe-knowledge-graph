---
title: bufio Package
tags:
  - golang
  - standard-library
  - io
parent: "[[Standard Library]]"
---

# `bufio`

Buffered wrappers around `io` readers and writers, plus the line/token scanner.

## 1. Why Buffering

```go
f, _ := os.Open("big.log")
r := bufio.NewReader(f)      // one syscall per 4096 bytes instead of per Read
```

An unbuffered `os.File.Read` of one byte is one syscall. Buffering turns
thousands of syscalls into a handful.

## 2. Scanner — the Line-Reading Idiom

```go
sc := bufio.NewScanner(f)
for sc.Scan() {
    line := sc.Text()
}
if err := sc.Err(); err != nil {   // ALWAYS check — Scan() returning false hides it
    return err
}
```

Split modes: `bufio.ScanLines` (default), `ScanWords`, `ScanRunes`, `ScanBytes`,
or a custom `SplitFunc`.

## 3. The 64 KB Line Limit

```text
bufio.Scanner: token too long
```

Default max token is 64 KB. For long lines (JSON logs):

```go
sc.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
```

Or use `bufio.Reader.ReadString('\n')`, which has no limit.

## 4. Writer Must Be Flushed

```go
w := bufio.NewWriter(f)
defer w.Flush()          // without this, buffered data is silently lost
fmt.Fprintln(w, "line")
```

Missing `Flush` is the most common bufio bug — the program exits successfully and
the file is short.

## 5. Gotchas

- `sc.Bytes()` returns a slice into the scanner's buffer, **invalidated by the
  next `Scan`**. Copy it if you keep it. `sc.Text()` allocates and is safe.
- `defer w.Flush()` swallows the flush error; check it explicitly when the write
  must succeed.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
