---
title: Code Generation & Metaprogramming Tooling
tags:
  - golang
  - toolchain
  - principal-swe
parent: "[[Go Toolchain & Developer Experience]]"
---

# Code Generation & Metaprogramming Tooling

go generate workflows, stringer enum generators, zero-reflection binary serialization (msgp), and OpenAPI/gRPC generators.

```text
Code Generation & Metaprogramming Tooling
│
├── [[go generate Workflow & Directives Architecture]]
├── [[Type-Safe Enums with stringer & jsonenums]]
├── [[Deep Copy & Serialization Generators (msgp, easyjson)]]
├── [[OpenAPI & Protobuf Code Generators (protoc-gen-go)]]
└── [[go fix & Automated Source Code Migration Tooling]]
```

---

## 🗂️ Topics

- [[go generate Workflow & Directives Architecture]] — Automated build-time generators: //go:generate stringer -type=Pill, //go:generate mockgen.
- [[Type-Safe Enums with stringer & jsonenums]] — Generating zero-allocation constant lookup strings and JSON serializers from AST definitions.
- [[Deep Copy & Serialization Generators (msgp, easyjson)]] — Build-time AST generation for zero-reflection binary encoding and decoding performance.
- [[OpenAPI & Protobuf Code Generators (protoc-gen-go)]] — Managing protoc, buf, and generating high-performance gRPC client/server stubs.
- [[go fix & Automated Source Code Migration Tooling]] — Leveraging AST rewrite tools to automate breaking syntax migrations across Go versions.

---

## 🔗 References
- ⬆️ Parent: [[Go Toolchain & Developer Experience]]

