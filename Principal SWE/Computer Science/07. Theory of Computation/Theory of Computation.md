---
title: Theory of Computation
tags:
  - computer-science
  - theory-of-computation
  - principal-swe
parent: "[[Computer Science]]"
---

# 🏛️ Theory of Computation (Foundations & Systems Architecture)

Automata theory, formal grammars, Turing machines, decidability, P vs NP completeness, space complexity (PSPACE), Shannon information theory, discrete mathematics, and quantum complexity.

```text
Theory of Computation
│
├── [[Finite Automata and Regular Languages|01. Finite Automata and Regular Languages]]
├── [[Context Free Grammars and Pushdown Automata|02. Context Free Grammars]]
├── [[Turing Machines and Computability Theory|03. Turing Machines and Computability]]
├── [[The Halting Problem and Rice's Theorem|04. Halting Problem and Decidability]]
├── [[P vs NP and NP Completeness Foundations|05. P vs NP and NP Completeness]]
├── [[Polynomial Time Reductions and NP Hardness|06. NP Reductions and Hardness]]
├── [[Space Complexity Classes (l, Nl, Pspace, Expspace)|07. Space Complexity Classes]]
├── [[Randomized and Quantum Complexity Classes|08. Randomized and Quantum Complexity]]
├── [[Propositional Logic and Mathematical Proof Techniques|09. Logic and Mathematical Proofs]]
├── [[Set Theory, Relations, and Functions|10. Set Theory and Relations]]
├── [[Combinatorics and Generating Functions|11. Combinatorics and Generating Functions]]
├── [[Discrete Probability and Random Variables|12. Discrete Probability]]
├── [[Graph Theory Axioms and Structural Theorems|13. Graph Theory Fundamentals]]
├── [[Mathematical Foundations of Number Theory|14. Number Theory Foundations]]
├── [[Recurrence Relations and Master Theorem|15. Recurrences and Master Theorem]]
├── [[Boolean Algebra and Digital Circuit Minimization|16. Boolean Algebra and Circuit Minimization]]
├── [[Shannon Entropy and Information Theory Foundations|17. Shannon Entropy and Information]]
├── [[Lossless Data Compression (huffman and Lz77)|18. Data Compression Huffman and Lz77]]
├── [[Error Correcting Codes (hamming and Reed Solomon)|19. Error Correcting Codes]]
├── [[Channel Capacity and Noisy Channel Coding Theorem|20. Channel Capacity and Coding Theorem]]
├── [[Kolmogorov Complexity and Algorithmic Information|21. Kolmogorov Complexity]]
├── [[Arithmetic Coding and Range Coding|22. Arithmetic and Range Coding]]
└── [[Mutual Information and Rate Distortion Theory|23. Mutual Information and Rate Distortion]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Finite Automata and Regular Languages|01. Finite Automata and Regular Languages]] — DFA, NFA, epsilon-transitions, Subset Construction algorithm, Regular Expressions, and the Pumping Lemma.
- 📂 [[Context Free Grammars and Pushdown Automata|02. Context Free Grammars]] — CFG production rules, Chomsky Normal Form (CNF), CYK parsing algorithm, Pushdown Automata (PDA), and ambiguity.
- 📂 [[Turing Machines and Computability Theory|03. Turing Machines and Computability]] — Single-tape and multi-tape Turing Machines, Church-Turing Thesis, Universal Turing Machine (UTM), and decidability.
- 📂 [[The Halting Problem and Rice's Theorem|04. Halting Problem and Decidability]] — Diagonalization proof of undecidability, Post Correspondence Problem (PCP), and Rice's theorem on semantic properties.
- 📂 [[P vs NP and NP Completeness Foundations|05. P vs NP and NP Completeness]] — Deterministic vs Non-Deterministic Polynomial Time, polynomial verifiers, Cook-Levin theorem (SAT), and millennium problem.
- 📂 [[Polynomial Time Reductions and NP Hardness|06. NP Reductions and Hardness]] — Karp reductions (SAT -> 3SAT -> Clique -> Vertex Cover -> Subset Sum), and proving hardness of combinatorial problems.
- 📂 [[Space Complexity Classes (l, Nl, Pspace, Expspace)|07. Space Complexity Classes]] — Savitch's Theorem (PSPACE = NPSPACE), PSPACE-completeness (QBF, Generalized Geography), and L vs NL.
- 📂 [[Randomized and Quantum Complexity Classes|08. Randomized and Quantum Complexity]] — BPP (Bounded-error Probabilistic Polynomial), RP, ZPP, BQP (Bounded-error Quantum Polynomial), and Shor's algorithm.
- 📂 [[Propositional Logic and Mathematical Proof Techniques|09. Logic and Mathematical Proofs]] — Truth tables, natural deduction, first-order predicate logic, proof by induction, contradiction, and contrapositive.
- 📂 [[Set Theory, Relations, and Functions|10. Set Theory and Relations]] — Zermelo-Fraenkel set axioms, equivalence relations, partial orders, lattices, cardinality, and Cantor's diagonal argument.
- 📂 [[Combinatorics and Generating Functions|11. Combinatorics and Generating Functions]] — Permutations, combinations, Pigeonhole Principle, Inclusion-Exclusion, and Ordinary/Exponential Generating Functions.
- 📂 [[Discrete Probability and Random Variables|12. Discrete Probability]] — Sample spaces, Bayes Theorem, Markov and Chebyshev inequalities, Chernoff bounds, and linearity of expectation.
- 📂 [[Graph Theory Axioms and Structural Theorems|13. Graph Theory Fundamentals]] — Planar graphs and Euler's formula (V - E + F = 2), Hall's Marriage Theorem, graph coloring, and Hamiltonian paths.
- 📂 [[Mathematical Foundations of Number Theory|14. Number Theory Foundations]] — Divisibility, Euclidean algorithm, Fundamental Theorem of Arithmetic, modular congruences, and Euler's phi function.
- 📂 [[Recurrence Relations and Master Theorem|15. Recurrences and Master Theorem]] — Homogeneous/non-homogeneous recurrences, recursion tree method, Akra-Bazzi method, and asymptotic recurrence bounds.
- 📂 [[Boolean Algebra and Digital Circuit Minimization|16. Boolean Algebra and Circuit Minimization]] — De Morgan's laws, Karnaugh maps (K-Maps), Quine-McCluskey algorithm, and NAND/NOR universal gate completeness.
- 📂 [[Shannon Entropy and Information Theory Foundations|17. Shannon Entropy and Information]] — Self-information $I(X) = -\log_2 P(X)$, Shannon Entropy $H(X)$, joint entropy, and conditional entropy bounds.
- 📂 [[Lossless Data Compression (huffman and Lz77)|18. Data Compression Huffman and Lz77]] — Optimal prefix codes, Huffman coding proof, LZ77 sliding window dictionary parsing, and DEFLATE.
- 📂 [[Error Correcting Codes (hamming and Reed Solomon)|19. Error Correcting Codes]] — Hamming distance $d_{\min}$, parity-check matrices, Reed-Solomon polynomial evaluation codes, and erasure recovery.
- 📂 [[Channel Capacity and Noisy Channel Coding Theorem|20. Channel Capacity and Coding Theorem]] — Shannon-Hartley theorem $C = B \log_2(1 + S/N)$, binary symmetric channels, and maximum achievable data rates.
- 📂 [[Kolmogorov Complexity and Algorithmic Information|21. Kolmogorov Complexity]] — Incompressibility theorem, minimal program length $K(x)$, Berry paradox, and algorithmic randomness.
- 📂 [[Arithmetic Coding and Range Coding|22. Arithmetic and Range Coding]] — Fractional interval subdivision for high-efficiency entropy coding outperforming Huffman on skewed probabilities.
- 📂 [[Mutual Information and Rate Distortion Theory|23. Mutual Information and Rate Distortion]] — Kullback-Leibler divergence (relative entropy), information gain, and lossy compression rate-distortion limits.

---

## 🔗 References
- ⬆️ Parent: [[Computer Science]]

