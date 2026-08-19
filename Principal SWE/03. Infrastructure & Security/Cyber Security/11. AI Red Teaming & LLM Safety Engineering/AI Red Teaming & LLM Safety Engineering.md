---
title: AI Red Teaming & LLM Safety Engineering
tags:
  - cyber-security
  - security-engineering
  - ai-red-teaming-and-llm-safety-engineering
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🛡️ AI Red Teaming & LLM Safety Engineering

Adversarial AI and LLM security: Direct/Indirect Prompt Injection, Jailbreaking, Training data poisoning, Adversarial perturbations, Model extraction & weight theft, and LLM Guardrails.

```text
AI Red Teaming & LLM Safety Engineering
│
├── [[Prompt Injection Attacks (direct System Overrides and Indirect Data Ingress)|01. Prompt Injection and Indirect Injection Vectors]]
├── [[LLM Jailbreaking, Roleplay Exploits, and Multi Turn Adversarial Prompts|02. Jailbreaking and Multi Turn Adversarial Prompts]]
├── [[Training Data Poisoning, Dataset Contamination, and Model Backdoors|03. Training Data Poisoning and Model Backdoors]]
├── [[Adversarial Perturbations, Gradient Based Attacks, and Model Robustness|04. Adversarial Robustness and Gradient Attacks]]
├── [[Model Extraction, Membership Inference, and Training Data Extraction|05. Model Extraction, Inversion, and Weight Theft]]
└── [[LLM Guardrails, Input Output Policy Filters (nemo, Llama Guard)|06. LLM Guardrails, Input Scrubbing, and Content Moderation]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Prompt Injection Attacks (direct System Overrides and Indirect Data Ingress)|01. Prompt Injection and Indirect Injection Vectors]] — Bypassing system instructions, malicious untrusted RAG document payloads, secondary prompt injection via web search, and delimiter defense strategies.
- 📂 [[LLM Jailbreaking, Roleplay Exploits, and Multi Turn Adversarial Prompts|02. Jailbreaking and Multi Turn Adversarial Prompts]] — Universal adversarial suffixes (GCG attacks), DAN exploits, multi-turn psychological priming, base64/rot13 obfuscation, and automated red team probing.
- 📂 [[Training Data Poisoning, Dataset Contamination, and Model Backdoors|03. Training Data Poisoning and Model Backdoors]] — Injecting stealthy triggers into pre-training/fine-tuning datasets, backdoor trigger activation, split-view poisoning, and dataset cryptographic hashing/deduplication.
- 📂 [[Adversarial Perturbations, Gradient Based Attacks, and Model Robustness|04. Adversarial Robustness and Gradient Attacks]] — Fast Gradient Sign Method (FGSM), Projected Gradient Descent (PGD), adversarial token substitutions, and adversarial training defensive fine-tuning.
- 📂 [[Model Extraction, Membership Inference, and Training Data Extraction|05. Model Extraction, Inversion, and Weight Theft]] — Stealing proprietary model weights via black-box API querying, reconstructing private training samples from LLM outputs, and differential privacy training.
- 📂 [[LLM Guardrails, Input Output Policy Filters (nemo, Llama Guard)|06. LLM Guardrails, Input Scrubbing, and Content Moderation]] — Semantic guardrail classifiers, PII masking before inference, hallucination detection, prompt-leakage prevention, and automated tool-call sandboxing.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]

