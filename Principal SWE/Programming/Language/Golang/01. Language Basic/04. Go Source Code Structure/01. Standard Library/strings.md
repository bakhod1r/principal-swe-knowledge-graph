---
title: strings Package
tags:
  - golang
  - standard-library
  - text
parent: "[[Standard Library]]"
---

# `strings`

UTF-8 string manipulation. The most-used package after `fmt`.

## 1. Core Functions

```go
strings.Contains(s, "go")          strings.HasPrefix(s, "http")
strings.Split(s, ",")              strings.Join(parts, ", ")
strings.TrimSpace(s)               strings.TrimPrefix(s, "v")
strings.ReplaceAll(s, "a", "b")    strings.EqualFold(a, b)  // ASCII-ish case-insensitive
strings.Cut(s, "=")                // Go 1.18 — before, after, found
```

`strings.Cut` replaced most two-element `SplitN` uses:

```go
key, value, ok := strings.Cut("GOOS=linux", "=")
```

## 2. `strings.Builder` — the Concatenation Rule

```go
var b strings.Builder
for _, p := range parts {
    b.WriteString(p)          // amortized O(1)
}
return b.String()             // no copy — unsafe-backed
```

`s += p` in a loop is O(n²): every `+=` allocates a new string and copies. This is
the single most common Go performance bug.

## 3. Bytes vs Runes

```go
s := "héllo"
len(s)                  // 6 — BYTES, not characters
utf8.RuneCountInString(s) // 5
s[1]                    // 0xc3 — half of é
for i, r := range s { } // r is a rune, i jumps by encoded width
```

Indexing a string yields a **byte**. Only `range` decodes UTF-8.

## 4. Gotchas

- Strings are immutable; every "modification" allocates. Use `[]byte` and
  `bytes` for in-place work.
- `strings.Title` is deprecated (broken for Unicode) — use
  `golang.org/x/text/cases`.
- `strings.NewReplacer` beats chained `ReplaceAll` for multiple replacements.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
