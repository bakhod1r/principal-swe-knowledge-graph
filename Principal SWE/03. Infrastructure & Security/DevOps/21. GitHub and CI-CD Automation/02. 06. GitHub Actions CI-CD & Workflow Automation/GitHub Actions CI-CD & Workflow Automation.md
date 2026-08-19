---
title: GitHub Actions CI-CD & Workflow Automation
tags:
  - review
  - devops
  - git-and-github
  - version-control
  - github-actions-ci-cd-and-workflow-automation
  - principal-swe
parent: "[[GitHub and CI-CD Automation]]"
---

# 🐙 GitHub Actions CI-CD & Workflow Automation

Continuous integration and delivery with GitHub Actions: Workflow syntax, matrix builds, self-hosted runners, caching, reusable workflows, composite actions, and OIDC cloud authentication.

```text
GitHub Actions CI-CD & Workflow Automation
│
├── [[GitHub Actions Workflow Syntax, Event Triggers, and Contexts|01. 01. GitHub Actions Workflow Syntax and Triggers]]
├── [[GitHub Actions Matrix Builds, Concurrency Groups, and Caching|02. 02. Matrix Builds, Concurrency, and Cache Optimization]]
├── [[Reusable Workflows (workflow_call) and Composite Actions|03. 03. Reusable Workflows and Composite Custom Actions]]
├── [[Self-Hosted Runners, Ephemeral Autoscaling, and ARC (Actions Runner Controller)|04. 04. Self-Hosted Runners and Ephemeral Infrastructure]]
├── [[GitHub Actions Cloud Authentication with OpenID Connect (OIDC)|05. 05. Cloud Authentication with OIDC and Workload Identity]]
└── [[GitHub Environments, Secrets Scoping, and Deployment Protection Rules|06. 06. Environments, Deployment Protection, and Approvals]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[GitHub Actions Workflow Syntax, Event Triggers, and Contexts|01. GitHub Actions Workflow Syntax and Triggers]] — Workflow YAML structure, triggers (`push`, `pull_request`, `schedule`, `workflow_dispatch`), path filtering, and context expressions (`github`, `secrets`, `env`).
- 📂 [[GitHub Actions Matrix Builds, Concurrency Groups, and Caching|02. Matrix Builds, Concurrency, and Cache Optimization]] — Cross-platform matrix builds (OS x Runtime), cancelling outdated PR runs (`concurrency: group`), caching dependencies with `actions/cache`, and build speedups.
- 📂 [[Reusable Workflows (workflow_call) and Composite Actions|03. Reusable Workflows and Composite Custom Actions]] — Creating centralized reusable CI/CD workflows across enterprise repos, inputs/secrets passing, and building custom Composite Actions with action.yml.
- 📂 [[Self-Hosted Runners, Ephemeral Autoscaling, and ARC (Actions Runner Controller)|04. Self-Hosted Runners and Ephemeral Infrastructure]] — Running CI on private Kubernetes clusters with Actions Runner Controller (ARC), runner security sandboxing, and ephemeral containerized runners.
- 📂 [[GitHub Actions Cloud Authentication with OpenID Connect (OIDC)|05. Cloud Authentication with OIDC and Workload Identity]] — Eliminating long-lived AWS/GCP access keys in GitHub Secrets, configuring IAM Workload Identity Federation, and short-lived STS tokens in CI.
- 📂 [[GitHub Environments, Secrets Scoping, and Deployment Protection Rules|06. Environments, Deployment Protection, and Approvals]] — Environment-scoped secrets (production vs staging), required manual approvers before deployment, wait timers, and deployment status webhooks.

---

## 🔗 References
- ⬆️ Parent: `Git & GitHub Version Control & CI-CD Automation`
- 📚 Module: `DevOps`

