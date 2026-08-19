---
title: For Range
tags:
  - golang
  - loops
  - principal-swe
parent: "[[Loops]]"
---

# For Range

Iterating over slices, arrays, maps, strings, channels, integers, and user iterator functions.

```text
For Range
│
├── [[Range over Slices & Arrays]]
├── [[Range over Maps & Randomization]]
├── [[Range over Strings (Rune Decoding)]]
├── [[Range over Channels]]
├── [[Range over Integer (Go 1.22+)]]
├── [[Range over Func (Go 1.23+ Iterators)]]
└── [[Range Value Copying Pitfall]]
```

---

## 🗂️ Topics

- [[Range over Slices & Arrays]] — Iterating index and element value copies (for i, v := range slice).
- [[Range over Maps & Randomization]] — Randomized map iteration order and sorted key extraction patterns.
- [[Range over Strings (Rune Decoding)]] — Decoding UTF-8 Unicode runes vs raw byte offsets during range.
- [[Range over Channels]] — Receiving values from channel until closed (for msg := range ch).
- [[Range over Integer (Go 1.22+)]] — Syntactic sugar counting loops (for i := range 10).
- [[Range over Func (Go 1.23+ Iterators)]] — Custom user iterator functions yielding elements to for-range.
- [[Range Value Copying Pitfall]] — Mutating iteration value copy instead of slice element.


## 🗂️ Contents

- [[Coroutine Stack Switching in Pull Iterators]]
- [[Push Iterators (iter.Seq) vs Pull Iterators (iter.Pull)]]
- [[Range Value Copying Pitfall]]
- [[Range over Channels]]
- [[Range over Func (Go 1.23+ Iterators)]]
- [[Range over Integer (Go 1.22+)]]
- [[Range over Maps & Randomization]]
- [[Range over Slices & Arrays]]
- [[Range over Strings (Rune Decoding)]]
- [[SIMD Vectorization & BCE in Loops]]

---

## 🔗 References
- ⬆️ Parent: [[Loops]]

