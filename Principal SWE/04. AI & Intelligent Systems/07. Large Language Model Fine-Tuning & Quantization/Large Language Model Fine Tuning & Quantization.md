---
title: Large Language Model Fine Tuning & Quantization
tags:
  - review
  - ai-and-machine-learning
  - ai-engineering
  - large-language-model-fine-tuning-and-quantization
  - principal-swe
parent: "[[AI & Intelligent Systems]]"
---

# 🤖 Large Language Model Fine Tuning & Quantization

Custom LLM adaptation: Parameter-Efficient Fine-Tuning (PEFT, LoRA, QLoRA), SFT dataset curation, DPO preference tuning, model quantization (GGUF, AWQ, GPTQ), model merging, and benchmark evaluation.

```text
Large Language Model Fine Tuning & Quantization
│
├── [[LLM Fine Tuning Paradigms: Full Fine Tuning vs PEFT (lora, Qlora)|01. Fine Tuning Paradigms Full vs Parameter Efficient PEFT]]
├── [[Supervised Fine Tuning (sft) Dataset Preparation, Cleaning, and Curation|02. SFT Dataset Preparation, Curation, and Quality Filtering]]
├── [[Executing SFT Fine Tuning Pipelines with Axolotl, Unsloth, and Trl|03. Supervised Fine Tuning SFT Pipelines with Axolotl and Unsloth]]
├── [[Direct Preference Optimization (dpo) and Odds Ratio Preference Optimization (orpo)|04. Direct Preference Optimization DPO and Alignment Tuning]]
├── [[Post Training Quantization (ptq): GGUF (llama.cpp), Awq, and GPTQ|05. Post Training Quantization Gguf, Awq, and GPTQ]]
├── [[Quantization Aware Training (qat) and Mixed Precision Training (fp16, Bf16, Fp8)|06. Quantization Aware Training QAT and Mixed Precision]]
├── [[Model Merging Architectures (slerp, Dare, Ties) with Mergekit|07. Model Merging Techniques with Mergekit]]
├── [[Open Weight Model Ecosystem: Llama 3, Mistral, Deepseek V3, Qwen 2.5|08. Open Weight Foundation Models Llama, Mistral, Deepseek]]
├── [[Standardized Model Evaluation: Mmlu, Humaneval, Gsm8k, and Mt Bench|09. Model Evaluation Benchmarks Mmlu, Humaneval, Mt Bench]]
└── [[Enterprise Domain Adaptation, Catastrophic Forgetting, and Continuous Training|10. Continuous Fine Tuning and Domain Adaptation in Enterprise]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[LLM Fine Tuning Paradigms: Full Fine Tuning vs PEFT (lora, Qlora)|01. Fine Tuning Paradigms Full vs Parameter Efficient PEFT]] — Why full fine-tuning is computationally prohibitive, Low-Rank Adaptation (LoRA) low-rank matrices decomposition, and QLoRA 4-bit NormalFloat quantization.
- 📂 [[Supervised Fine Tuning (sft) Dataset Preparation, Cleaning, and Curation|02. SFT Dataset Preparation, Curation, and Quality Filtering]] — Structuring instruction-response pairs (Alpaca, ShareGPT formats), synthetic data generation with frontier models, deduplication (MinHash), and quality filtering.
- 📂 [[Executing SFT Fine Tuning Pipelines with Axolotl, Unsloth, and Trl|03. Supervised Fine Tuning SFT Pipelines with Axolotl and Unsloth]] — Accelerating training with Unsloth fast kernels, configuring Axolotl YAML training configs, gradient accumulation, FlashAttention-2 integration, and loss curves.
- 📂 [[Direct Preference Optimization (dpo) and Odds Ratio Preference Optimization (orpo)|04. Direct Preference Optimization DPO and Alignment Tuning]] — Preparing chosen vs rejected preference datasets, tuning the beta hyperparameter, eliminating separate reward models, and stabilizing preference training.
- 📂 [[Post Training Quantization (ptq): GGUF (llama.cpp), Awq, and GPTQ|05. Post Training Quantization Gguf, Awq, and GPTQ]] — Comparing quantization schemes: Weight-only vs Weight-and-Activation quantization, 4-bit AWQ activation-aware quantization, and GGUF for CPU/Apple Silicon.
- 📂 [[Quantization Aware Training (qat) and Mixed Precision Training (fp16, Bf16, Fp8)|06. Quantization Aware Training QAT and Mixed Precision]] — Simulating quantization noise during forward passes, mixed-precision BF16 training stability, and hardware acceleration on NVIDIA Hopper/Blackwell FP8.
- 📂 [[Model Merging Architectures (slerp, Dare, Ties) with Mergekit|07. Model Merging Techniques with Mergekit]] — Combining capabilities of multiple fine-tuned models without training: Spherical Linear Interpolation (SLERP), TIES-Merging, and DARE sparsification.
- 📂 [[Open Weight Model Ecosystem: Llama 3, Mistral, Deepseek V3, Qwen 2.5|08. Open Weight Foundation Models Llama, Mistral, Deepseek]] — Architectural innovations: SwiGLU activations, RoPE base scaling, Mixture-of-Experts (MoE) routing, Multi-Head Latent Attention (MLA), and license analysis.
- 📂 [[Standardized Model Evaluation: Mmlu, Humaneval, Gsm8k, and Mt Bench|09. Model Evaluation Benchmarks Mmlu, Humaneval, Mt Bench]] — Running standardized evaluation harnesses (lm-evaluation-harness), measuring reasoning, coding accuracy, mathematical capability, and multi-turn chat quality.
- 📂 [[Enterprise Domain Adaptation, Catastrophic Forgetting, and Continuous Training|10. Continuous Fine Tuning and Domain Adaptation in Enterprise]] — Preventing catastrophic forgetting by mixing replay data, domain-specific vocabulary extension, and continuous retraining pipelines for enterprise data.

---

## 🔗 References
- ⬆️ Parent: `AI & Machine Learning`

