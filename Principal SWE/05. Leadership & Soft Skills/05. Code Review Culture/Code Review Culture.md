---
title: Code Review Culture
tags:
  - soft-skills
  - leadership
  - engineering-management
  - team-lead
  - code-review-culture-and-engineering-standards
  - principal-swe
parent: "[[Leadership & Soft Skills]]"
---

# 🤝 Code Review Culture & Engineering Standards

Code review excellence: Egoless code reviews, review SLAs & turnaround times, small PR sizing & stacked diffs, automated CI linting gates, constructive feedback, reviewing architecture vs style, and CODEOWNERS governance.

```text
Code Review Culture & Engineering Standards
│
├── [[Egoless Code Reviews - Separating Self Worth From Source Code|01. Egoless Code Review Principles and Team Culture]]
├── [[Code Review Slas - Minimizing Review Latency and Blocking Queues|02. Code Review Slas and Review Latency Optimization]]
├── [[Small PR Discipline - Sub 200 Line PRs and Stacked Diffs|03. Small Pull Requests and Stacked Diffs Methodology]]
├── `04. Automated Ci Linting and Static Analysis Gates`
├── [[Constructive Feedback - Conventional Comments and Non Violent Phrasing|05. Constructive Feedback and Non Violent Communication]]
├── `06. Reviewing Architecture, Security, and Edge Cases`
├── `07. Security and Performance Review Checklists`
└── `08. Repository Codeowners and Mandatory Domain Approvals`
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Egoless Code Reviews - Separating Self Worth From Source Code|01. Egoless Code Review Principles and Team Culture]] — Treating code reviews as learning opportunities, fostering curiosity over criticism, assuming positive intent, and building high-trust code review culture.
- 📂 [[Code Review Slas - Minimizing Review Latency and Blocking Queues|02. Code Review Slas and Review Latency Optimization]] — Setting 4-hour review turnaround SLAs, prioritizing reviews over starting new work, preventing WIP accumulation, and maintaining team momentum.
- 📂 [[Small PR Discipline - Sub 200 Line PRs and Stacked Diffs|03. Small Pull Requests and Stacked Diffs Methodology]] — Why large PRs receive superficial reviews, breaking features into small atomic pull requests, and managing stacked branches with tools like Graphite/git-town.
- 📂 `04. Automated Ci Linting and Static Analysis Gates` — Moving nitpicky formatting and stylistic debates into automated pre-commit and CI linters (Prettier, ESLint, golangci-lint), freeing humans for architectural review.
- 📂 [[Constructive Feedback - Conventional Comments and Non Violent Phrasing|05. Constructive Feedback and Non Violent Communication]] — Using Conventional Comments prefixes (`nit:`, `suggestion:`, `question:`, `blocking:`), asking open questions instead of commands, and praising good patterns.
- 📂 `06. Reviewing Architecture, Security, and Edge Cases` — Focusing human reviews on domain boundaries, database query performance (N+1s), race conditions, error handling, backward compatibility, and security vulnerabilities.
- 📂 `07. Security and Performance Review Checklists` — Structured verification checklists: Input validation, authentication/authorization checks, SQL injection, memory leaks, unclosed resources, and timeouts.
- 📂 `08. Repository Codeowners and Mandatory Domain Approvals` — Configuring `.github/CODEOWNERS` to ensure specialized domain architects (Security, DB, Infrastructure) review and approve changes touching critical paths.

---

## 🔗 References
- ⬆️ Parent: `Soft Skills`

