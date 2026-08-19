---
title: TypeScript
tags:
  - programming
  - languages
  - typescript
  - principal-swe
parent: "[[Language]]"
---

# 💻 TypeScript

Large-scale TypeScript engineering: Static type system, structural subtyping, union/intersection types, type narrowing, generics, conditional types, utility types, template literals, tsconfig tuning, AST transformations, and Node/Bun runtime ecosystems.

```text
TypeScript
│
├── [[TypeScript Architecture, Compiler (tsc), and AST Pipeline|01. Introduction to TypeScript and Compiler Architecture]]
├── [[TypeScript Basic and Primitive Types (any, Unknown, Never, Void)|02. TypeScript Basic and Primitive Type System]]
├── [[TypeScript Combining Types: Union Types, Intersections, and Discriminated Unions|03. Combining Types Unions, Intersections, and Discriminated Unions]]
├── [[TypeScript Type Narrowing, Custom Type Guards, and Assertion Functions|04. Type Guards, Narrowing, and Assertion Signatures]]
├── [[TypeScript Interfaces vs Type Aliases and Declaration Merging|05. Interfaces vs Type Aliases and Declaration Merging]]
├── [[TypeScript Functions: Parameter Types, Overloads, and This Parameters|06. Functions, Overloads, and This Typing]]
├── [[Object Oriented Typescript: Classes, Visibility, and Abstract Classes|07. Object Oriented Typescript, Classes, and Access Modifiers]]
├── [[TypeScript Generics, Generic Functions, and Type Parameter Constraints|08. Generics and Generic Constraints]]
├── [[TypeScript Built in Utility Types (partial, Required, Pick, Omit, Record)|09. Built in Utility Types and Type Transformations]]
├── [[Advanced Typescript: Conditional Types, Mapped Types, and Keyof|10. Advanced Types Conditional, Mapped, and Indexed Types]]
├── [[Template Literal Types, Recursive Types, and Type Level Computation|11. Template Literal Types and Recursive Type Construction]]
├── [[TypeScript Module Systems, Path Aliases, and Tsconfig Optimization|12. Modules, Namespaces, and Tsconfig Optimization]]
└── [[TypeScript Tooling Ecosystem: Eslint, Prettier, Vite, Tsx, and Bun|13. TypeScript Ecosystem, Linters, and Bundlers]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[TypeScript Architecture, Compiler (tsc), and AST Pipeline|01. Introduction to TypeScript and Compiler Architecture]] — TypeScript as a typed superset of JavaScript, structural type system vs nominal typing, compiler phases (Scanner, Parser, Binder, Checker, Emitter), and `@types` declarations.
- 📂 [[TypeScript Basic and Primitive Types (any, Unknown, Never, Void)|02. TypeScript Basic and Primitive Type System]] — Type annotations vs type inference, primitive types, why `unknown` is safer than `any`, exhaustiveness checking with `never`, and literal types.
- 📂 [[TypeScript Combining Types: Union Types, Intersections, and Discriminated Unions|03. Combining Types Unions, Intersections, and Discriminated Unions]] — Union types (`|`), Intersection types (`&`), Discriminated unions with literal tag fields, exhaustive switch statements, and domain modeling.
- 📂 [[TypeScript Type Narrowing, Custom Type Guards, and Assertion Functions|04. Type Guards, Narrowing, and Assertion Signatures]] — Control flow analysis, `typeof` and `instanceof` narrowing, `in` operator narrowing, user-defined type guards (`is`), and assertion functions (`asserts`).
- 📂 [[TypeScript Interfaces vs Type Aliases and Declaration Merging|05. Interfaces vs Type Aliases and Declaration Merging]] — Object shape definitions, extending interfaces, intersection with types, declaration merging in global namespace, and performance implications on compiler checker.
- 📂 [[TypeScript Functions: Parameter Types, Overloads, and This Parameters|06. Functions, Overloads, and This Typing]] — Optional and default parameters, rest parameters, function overloads and implementation signatures, context-aware `this` typing, and generic functions.
- 📂 [[Object Oriented Typescript: Classes, Visibility, and Abstract Classes|07. Object Oriented Typescript, Classes, and Access Modifiers]] — Class member visibility (`public`, `private`, `protected`), ECMAScript private fields (`#`), `readonly` modifiers, parameter properties, and abstract classes.
- 📂 [[TypeScript Generics, Generic Functions, and Type Parameter Constraints|08. Generics and Generic Constraints]] — Generic type parameters, generic constraints (`extends`), multiple type parameters, default type arguments, and generic classes/interfaces.
- 📂 [[TypeScript Built in Utility Types (partial, Required, Pick, Omit, Record)|09. Built in Utility Types and Type Transformations]] — Standard utility types: `Partial<T>`, `Required<T>`, `Readonly<T>`, `Record<K, T>`, `Pick<T, K>`, `Omit<T, K>`, `Exclude<T, U>`, `Extract<T, U>`, and `ReturnType<T>`.
- 📂 [[Advanced Typescript: Conditional Types, Mapped Types, and Keyof|10. Advanced Types Conditional, Mapped, and Indexed Types]] — Conditional types (`T extends U ? X : Y`), `infer` keyword type deduction, Mapped types (`[K in keyof T]`), Key remapping (`as`), and Index access types (`T[K]`).
- 📂 [[Template Literal Types, Recursive Types, and Type Level Computation|11. Template Literal Types and Recursive Type Construction]] — String manipulation with template literal types (`Uppercase`, `Lowercase`, `Capitalize`), recursive deeply-nested object paths (`a.b.c`), and type-level programming.
- 📂 [[TypeScript Module Systems, Path Aliases, and Tsconfig Optimization|12. Modules, Namespaces, and Tsconfig Optimization]] — ESM vs CommonJS resolution (`moduleResolution: NodeNext`), path aliases (`paths`), `strict: true` compiler flags, `noUncheckedIndexedAccess`, and project references.
- 📂 [[TypeScript Tooling Ecosystem: Eslint, Prettier, Vite, Tsx, and Bun|13. TypeScript Ecosystem, Linters, and Bundlers]] — Type-aware ESLint rules (`@typescript-eslint`), fast compilation with SWC and esbuild, running TypeScript natively with Bun/Deno, and runtime schema validation with Zod.

---

## 🔗 References
- ⬆️ Parent: [[Language]]

