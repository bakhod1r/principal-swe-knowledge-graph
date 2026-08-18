---
title: strconv Package
tags:
  - golang
  - standard-library
  - conversion
parent: "[[Standard Library]]"
---

# `strconv`

Conversions between strings and the basic numeric and boolean types.

## 1. The Core Four

```go
n, err := strconv.Atoi("42")                  // string → int
s := strconv.Itoa(42)                          // int → string
f, err := strconv.ParseFloat("3.14", 64)
b, err := strconv.ParseBool("true")

i64, err := strconv.ParseInt("ff", 16, 64)     // base 16, fits in int64
s = strconv.FormatInt(255, 16)                 // "ff"
```

## 2. Never Use `string(int)`

```go
string(65)              // "A" — a rune conversion, not a number
strconv.Itoa(65)        // "65" — what you meant
```

`go vet` flags this (`stringintconv`); it is a classic silent bug.

## 3. Error Handling

```go
n, err := strconv.Atoi(input)
if err != nil {
    var ne *strconv.NumError
    if errors.As(err, &ne) {
        log.Printf("bad %q in %s: %v", ne.Num, ne.Func, ne.Err)
    }
}
```

`NumError` carries the offending input — useful in validation messages. See
`errors`.

## 4. Quoting

```go
strconv.Quote(`he said "hi"`)      // "\"he said \\\"hi\\\"\""
strconv.Unquote(s)
strconv.AppendInt(buf, 42, 10)     // allocation-free formatting into a buffer
```

The `Append*` family writes into an existing `[]byte` — the allocation-free path
used inside `fmt` and logging libraries.

## 5. Gotchas

- `Atoi` is `ParseInt(s, 10, 0)`: platform-width `int`, so it overflows
  differently on 32-bit.
- `ParseFloat` accepts `"NaN"`, `"Inf"`, `"1e10"`, and underscores in Go 1.13+
  literals.
- For performance-critical formatting, `AppendInt` into a reused buffer beats
  `fmt.Sprintf` by an order of magnitude.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
