---
title: encoding/json Package
tags:
  - golang
  - standard-library
  - serialization
parent: "[[Standard Library]]"
---

# `encoding/json`

JSON marshalling via struct tags and reflection.

## 1. Struct Tags

```go
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name,omitempty"`
    Password  string    `json:"-"`                 // never serialized
    CreatedAt time.Time `json:"created_at"`
    Meta      any       `json:"meta,omitempty"`
}
```

| Tag | Effect |
|---|---|
| `json:"x"` | Rename the field |
| `,omitempty` | Skip zero values (0, "", nil, empty slice/map) |
| `json:"-"` | Always skip |
| `,string` | Encode a number as a JSON string |

## 2. Streaming

```go
json.NewDecoder(r.Body).Decode(&v)     // no full buffering
json.NewEncoder(w).Encode(v)

dec := json.NewDecoder(r.Body)
dec.DisallowUnknownFields()            // reject unexpected keys
```

Prefer the streaming forms in HTTP handlers over `Marshal`/`Unmarshal` on a
buffered body. See `net_http`.

## 3. Unmarshalling Into `any`

```text
JSON        →  Go
object      →  map[string]any
array       →  []any
number      →  float64        ← every number, including integers
string      →  string
null        →  nil
```

The `float64` row is the trap: an int64 ID above 2^53 loses precision. Use
`json.Number` with `dec.UseNumber()`.

## 4. `omitempty` Does Not Mean "omit zero struct"

```go
Address Address `json:"address,omitempty"`   // struct is NEVER omitted
Address *Address `json:"address,omitempty"`  // pointer IS omitted when nil
```

`omitempty` has no notion of an empty struct. Go 1.24 added `omitzero`, which
does.

## 5. Gotchas

- Only **exported** fields are marshalled — a lowercase field silently vanishes.
- Unmarshalling into a non-empty map or slice merges rather than replaces.
- Reflection-based; on hot paths, code generators are ~5× faster.
- HTML characters are escaped by default; disable with `enc.SetEscapeHTML(false)`.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library]]
