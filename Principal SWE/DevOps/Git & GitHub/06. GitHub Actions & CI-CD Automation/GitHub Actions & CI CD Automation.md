---
title: GitHub Actions & CI CD Automation
tags:
  - git
  - github
  - version-control
  - github-actions-and-ci-cd-automation
  - principal-swe
parent: "[[Git & GitHub]]"
---

# 🏛️ GitHub Actions & CI CD Automation

Continuous integration and delivery automation on GitHub: Workflow YAML syntax, event triggers, matrix builds, self-hosted runners, action secrets & OIDC cloud auth, dependency caching, artifact management, and custom composite actions.

```text
GitHub Actions & CI CD Automation
│
├── [[GitHub Actions Workflow Syntax and Event Triggers|01. GitHub Actions Workflow Syntax and Triggers]]
├── [[GitHub Actions Matrix Strategy and Parallel Builds|02. Matrix Strategy and Parallel Builds]]
├── [[GitHub Actions Runners (cloud Hosted vs Self Hosted)|03. Runners Hosted vs Self Hosted]]
├── [[GitHub Actions Secrets and Cloud Openid Connect (oidc)|04. Secrets Management and Cloud Oidc]]
├── [[GitHub Actions Dependency Caching and Artifact Storage|05. Dependency Caching and Artifact Storage]]
└── [[Authoring Custom Composite and Docker Container Actions|06. Custom Composite and Docker Actions]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[GitHub Actions Workflow Syntax and Event Triggers|01. GitHub Actions Workflow Syntax and Triggers]] — Push, pull_request, workflow_dispatch, and scheduled cron triggers; job dependencies, outputs, and conditional expressions.
- 📂 [[GitHub Actions Matrix Strategy and Parallel Builds|02. Matrix Strategy and Parallel Builds]] — Running cross-platform test suites across multiple OS versions and runtime language versions simultaneously.
- 📂 [[GitHub Actions Runners (cloud Hosted vs Self Hosted)|03. Runners Hosted vs Self Hosted]] — Configuring enterprise self-hosted runner fleets on Kubernetes (Actions Runner Controller ARC) with ephemeral pods.
- 📂 [[GitHub Actions Secrets and Cloud Openid Connect (oidc)|04. Secrets Management and Cloud Oidc]] — Eliminating long-lived cloud credentials using short-lived OIDC tokens for passwordless AWS/GCP/Azure role assumption.
- 📂 [[GitHub Actions Dependency Caching and Artifact Storage|05. Dependency Caching and Artifact Storage]] — Accelerating CI build times via actions/cache for package managers and publishing build binaries via actions/upload-artifact.
- 📂 [[Authoring Custom Composite and Docker Container Actions|06. Custom Composite and Docker Actions]] — Building reusable organization-wide Actions (action.yml) with composite steps and containerized execution.

---

## 🔗 References
- ⬆️ Parent: [[Git & GitHub]]
- 🎓 Root: [[Principal SWE]]
