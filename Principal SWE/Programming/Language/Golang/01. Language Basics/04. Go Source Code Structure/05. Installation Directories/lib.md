---
title: lib
tags:
  - golang
  - goroot
  - assets
  - wasm
  - fips
parent: "[[Installation Directories]]"
---

# `$GOROOT/lib`

Non-Go support files the toolchain needs at build or run time.

## 1. Contents

```text
$GOROOT/lib/
├── fips140/     ← FIPS 140-3 validated crypto module snapshots (.zip)
├── time/        ← the embedded IANA tzdata (zoneinfo.zip)
├── wasm/        ← wasm_exec.js and the WASI/browser support shims
└── hg/          ← Mercurial helper config
```

## 2. `lib/time`

The tzdata database `time.LoadLocation` falls back to when the host has no
system zoneinfo — typically inside a `FROM scratch` container.

```go
import _ "time/tzdata"   // embed ~450 KB of tzdata into the binary
```

Without that import a `scratch` image returns
`unknown time zone America/New_York` at runtime, in production, at 2am.

## 3. `lib/wasm`

```bash
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./web/
GOOS=js GOARCH=wasm go build -o web/app.wasm ./cmd/web
```

`wasm_exec.js` is the JavaScript host glue; version it with the toolchain that
built the `.wasm`, mismatches fail obscurely.

## 4. `lib/fips140`

Frozen, validated snapshots of the crypto module, selected by `GOFIPS140`.
Regulatory builds pin one of these zips rather than using the live source.

## 5. Gotchas

- Moved here in recent releases (previously scattered under `misc/`); older
  tutorials point at `misc/wasm/wasm_exec.js`. See `misc`.
- These are toolchain assets, not importable packages.

---

## 🔗 References
- ⬆️ Parent: [[Installation Directories]]
