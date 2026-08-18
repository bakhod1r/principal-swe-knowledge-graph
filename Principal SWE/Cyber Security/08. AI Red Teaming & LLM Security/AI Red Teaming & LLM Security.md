---
title: AI Red Teaming & LLM Security
tags:
  - cyber-security
  - appsec
  - ai-red-teaming-and-llm-security
  - principal-swe
parent: "[[Cyber Security]]"
---

# 🏛️ AI Red Teaming & LLM Security

Artificial Intelligence & Large Language Model security: Direct and indirect Prompt Injection, Jailbreak techniques, Training Data Poisoning, Adversarial inputs, Model Inversion & Weight Theft, and LLM Guardrail pipelines.

```text
AI Red Teaming & LLM Security
│
├── [[Prompt Injection and LLM Jailbreak Techniques|01. Prompt Injection and Jailbreaks]]
├── [[Training Data Poisoning and Model Backdoors|02. Training Data Poisoning and Backdoors]]
├── [[Adversarial Attacks and Machine Learning Robustness|03. Adversarial Attacks and Robustness]]
├── [[Model Extraction, Inversion, and Weight Theft|04. Model Extraction and Weight Theft]]
├── [[LLM Guardrails and Input-output Sanitization|05. LLM Guardrails and Input Sanitization]]
└── [[AI Red Teaming Frameworks and Automated Evals|06. AI Red Teaming Frameworks and Evals]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Prompt Injection and LLM Jailbreak Techniques|01. Prompt Injection and Jailbreaks]] — Direct prompt injection, indirect context injection via RAG/web scraping, DAN jailbreaks, base64 payload obfuscation, and defenses.
- 📂 [[Training Data Poisoning and Model Backdoors|02. Training Data Poisoning and Backdoors]] — Poisoning pre-training datasets, trigger phrases inducing hidden backdoors, and data curation integrity validation.
- 📂 [[Adversarial Attacks and Machine Learning Robustness|03. Adversarial Attacks and Robustness]] — Fast Gradient Sign Method (FGSM), adversarial perturbation attacks, evasion attacks on classifiers, and robust model training.
- 📂 [[Model Extraction, Inversion, and Weight Theft|04. Model Extraction and Weight Theft]] — API query distillation attacks, model inversion extracting private training data, and securing model artifact storage (S3/KMS).
- 📂 [[LLM Guardrails and Input-output Sanitization|05. LLM Guardrails and Input Sanitization]] — NeMo Guardrails, Llama Guard, regex/embedding-based semantic filtering, output validation, and preventing toxic/PII leakage.
- 📂 [[AI Red Teaming Frameworks and Automated Evals|06. AI Red Teaming Frameworks and Evals]] — Automated adversarial red teaming pipelines (Garak, PyRIT), benchmark evaluation suites, and responsible disclosure.

---

## 🔗 References
- ⬆️ Parent: [[Cyber Security]]
- 🎓 Root: [[Principal SWE]]
