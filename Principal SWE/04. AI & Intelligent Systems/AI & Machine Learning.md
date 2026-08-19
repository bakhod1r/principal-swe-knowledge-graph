---
title: AI & Machine Learning
tags:
  - ai-and-machine-learning
  - ai-engineering
  - llms
  - rag
  - agents
  - data-engineering
  - principal-swe
parent: "[[Principal SWE]]"
---

# 🤖 AI, Machine Learning & Intelligent Systems Engineering

Comprehensive, production-grade master architecture covering the complete spectrum of transformer mechanics, prompt optimization, RAG architectures, autonomous AI agents, Claude Code & Model Context Protocol (MCP), AI-assisted software engineering, LLM fine-tuning, high-performance inference (vLLM), modern data engineering lakehouses, and statistical modeling across 10 master pillars and 112 specialized subdomains:

- **Machine Learning & Transformer Foundations:** Math foundations, supervised/unsupervised learning, deep learning, Self-Attention & Multi-Head Attention, Chinchilla scaling laws, subword BPE tokenization, AdamW optimization, and evaluation metrics.
- **Prompt Engineering & LLM Alignment:** Few-shot in-context learning, Chain-of-Thought (CoT), Tree of Thoughts (ToT), KV-cache optimization, JSON Schema enforcement, RLHF & DPO alignment, and prompt injection defense.
- **AI Engineering & RAG Architectures:** Multi-model APIs, dense embeddings, vector indexing (HNSW/IVF-PQ), vector databases (Qdrant, Milvus, pgvector), chunking strategies, hybrid search (BM25 + Dense), cross-encoder reranking, Advanced RAG (Self-RAG, GraphRAG), and Ragas evaluation.
- **AI Agents & Multi-Agent Systems:** ReAct loops, tool calling, memory systems, planning algorithms (Reflexion), LangGraph state machines, LlamaIndex workflows, CrewAI & AutoGen multi-agent swarms, Human-in-the-Loop gates, and secure code execution sandboxes.
- **Claude Code & Model Context Protocol MCP:** Autonomous terminal agents, MCP client-server specification (JSON-RPC), subagent delegation, multi-file code synthesis, conversation compaction, slash commands, and production codebase refactoring.
- **AI-Assisted Engineering & Vibe Coding:** Cursor Composer, GitHub Copilot, Plan-Before-You-Code methodologies, context engineering (`.cursorrules`), automated test generation, AI PR review, and rapid full-stack vibe coding ergonomics.
- **LLM Fine-Tuning & Quantization:** PEFT, LoRA & QLoRA, SFT dataset preparation, DPO preference tuning, quantization (GGUF, AWQ, GPTQ), model merging (MergeKit), and open-weight models (Llama 3, Mistral, DeepSeek).
- **High-Performance Inference & LLMOps:** vLLM, TensorRT-LLM, PagedAttention, FlashAttention-3, speculative decoding, distributed multi-GPU inference, Langfuse observability, and Llama Guard filtering.
- **Data Engineering & Lakehouses:** Modern Data Stack, batch vs streaming, Apache Spark internals, Apache Flink, Lakehouse formats (Delta Lake, Apache Iceberg), Airflow/Dagster orchestration, dimensional modeling (Kimball), and dbt transformations.
- **AI Data Science & Econometrics:** Bayesian statistics, hypothesis testing, econometrics & causal inference, EDA & feature engineering, time series forecasting (ARIMA, Prophet), dimensionality reduction (UMAP), and explainable AI (SHAP/LIME).

```text
AI & Machine Learning
│
├── `01. Machine Learning, Deep Learning & Transformer Foundations`
│   ├── `01. Mathematical Foundations for Machine Learning`
│   ├── `02. Supervised Learning Algorithms and Regression`
│   ├── `03. Unsupervised Learning, Clustering, and PCA`
│   ├── `04. Neural Network Architectures and Backpropagation`
│   ├── `05. Convolutional and Recurrent Architectures`
│   ├── `06. Attention Mechanism and Transformer Architecture`
│   ├── `07. Large Language Model Pre Training and Scaling Laws`
│   ├── `08. Subword Tokenization Algorithms BPE and Wordpiece`
│   ├── `09. Deep Learning Optimization Algorithms and Schedulers`
│   ├── `10. Loss Functions and Regularization Techniques`
│   └── `11. Model Evaluation Metrics and Validation Strategies`
├── [[Prompt Engineering, Context Optimization & LLM Alignment|02. Prompt Engineering, Context Optimization & LLM Alignment]]
│   ├── `01. Prompt Engineering Foundations and Best Practices`
│   ├── `02. In Context Learning and Few Shot Prompting`
│   ├── `03. Chain of Thought CoT and Reasoning Elicitation`
│   ├── `04. Advanced Reasoning Tree of Thoughts and Graph of Thoughts`
│   ├── `05. Context Window Management and Kv Cache Optimization`
│   ├── `06. Structured Outputs and Json Schema Enforcement`
│   ├── `07. Rlhf, Dpo, and Modern LLM Alignment`
│   ├── `08. System Prompts, Guardrail Directives, and Persona Design`
│   └── `09. Prompt Security, Jailbreak Defenses, and Leakage Prevention`
├── `03. AI Engineering, Vector Databases & RAG Architectures`
│   ├── `01. Foundation Model APIs and Multi Provider Gateways`
│   ├── `02. Dense Embeddings, Similarity Metrics, and Quantization`
│   ├── `03. Vector Database Internals and Indexing Algorithms`
│   ├── `04. Vector Databases in Production Qdrant, Milvus, Pgvector`
│   ├── `05. Document Ingestion and Chunking Strategies`
│   ├── `06. Dense vs Sparse Hybrid Search and Bm25`
│   ├── `07. Cross Encoder Reranking and Contextual Compression`
│   ├── `08. Advanced RAG Architectures Self RAG and Corrective RAG`
│   ├── `09. Graphrag and Knowledge Graph Augmented Retrieval`
│   ├── `10. Multi Modal Retrieval and Document Parsing`
│   ├── `11. RAG Evaluation Frameworks Ragas and Trulens`
│   ├── `12. Semantic Caching and LLM Response Optimization`
│   └── `13. Token Economics, Budgeting, and Cost Governance`
├── `04. AI Agents, Autonomous Systems & Multi-Agent Orchestration`
│   ├── `01. AI Agent Architecture and the ReAct Paradigm`
│   ├── `02. Tool Calling and Function Calling Specifications`
│   ├── `03. Agent Memory Systems Short Term, Long Term, Episodic`
│   ├── `04. Planning Algorithms Plan and Solve and Reflexion`
│   ├── `05. Langgraph State Machine Architecture and Workflows`
│   ├── `06. Llamaindex Agent Workflows and Query Engines`
│   ├── `07. Crewai Multi Agent Role Playing Collaboration`
│   ├── `08. Autogen Conversational Patterns and Group Chats`
│   ├── `09. Human in the Loop HITL and Approval Gates`
│   ├── `10. Agent Code Sandboxes and Secure Execution Runtimes`
│   ├── `11. Error Recovery, Loop Detection, and Self Healing Agents`
│   ├── `12. Guardrails, Safety Filters, and Output Governance`
│   ├── `13. Autonomous Web Browsing and Computer Use Agents`
│   └── `14. Agent Observability, Tracing, and Evaluation`
├── `05. Claude Code, Subagents, Model Context Protocol MCP & Tooling`
│   ├── `01. Claude Code Architecture and Terminal Execution Engine`
│   ├── `02. Model Context Protocol MCP Standards and Architecture`
│   ├── `03. Developing Custom MCP Servers in Typescript and Python`
│   ├── `04. Subagent Delegation, Swarms, and Context Handoffs`
│   ├── `05. Multi File Code Synthesis and Atomic Patching`
│   ├── `06. Terminal and Tool Execution Safety and Sandboxing`
│   ├── `07. Conversation History Compaction and Transcript Management`
│   ├── `08. Custom Slash Commands and Background Execution`
│   ├── `09. Skill and Rule Customization Engine and Discovery`
│   ├── `10. Automated PR Generation and CI Integration with Claude Code`
│   ├── `11. Production Codebase Refactoring with Autonomous Agents`
│   ├── `12. Debugging Autonomous Agent Failure Modes and Hallucinations`
│   └── `13. Benchmark Evaluations for Coding Agents Swe Bench`
├── `06. AI-Assisted Engineering, Cursor, Copilot & Vibe Coding`
│   ├── `01. The AI Assisted Software Engineering Paradigm`
│   ├── `02. Cursor IDE Architecture, Composer, and Multi File Edits`
│   ├── `03. Github Copilot, Inline Completions, and Workspace Indexing`
│   ├── `04. Plan Before You Code and Specification Driven Development`
│   ├── `05. Context Engineering .cursorrules and Repository Maps`
│   ├── `06. AI Generated Unit, Integration, and Property Tests`
│   ├── `07. Automated Pull Request Reviewers and Agentic CI`
│   ├── `08. Rapid Prototyping and Greenfield Project Scaffolding`
│   ├── `09. Detecting and Mitigating AI Hallucinations in Code`
│   ├── `10. Full Stack Vibe Coding Ergonomics and Velocity`
│   ├── `11. Refactoring Legacy Codebases with AI Pair Programming`
│   └── `12. Developer Ergonomics, Cognitive Flow, and AI Tooling Fatigue`
├── [[Large Language Model Fine Tuning & Quantization|07. Large Language Model Fine-Tuning & Quantization]]
│   ├── `01. Fine Tuning Paradigms Full vs Parameter Efficient PEFT`
│   ├── `02. SFT Dataset Preparation, Curation, and Quality Filtering`
│   ├── `03. Supervised Fine Tuning SFT Pipelines with Axolotl and Unsloth`
│   ├── `04. Direct Preference Optimization DPO and Alignment Tuning`
│   ├── `05. Post Training Quantization Gguf, Awq, and GPTQ`
│   ├── `06. Quantization Aware Training QAT and Mixed Precision`
│   ├── `07. Model Merging Techniques with Mergekit`
│   ├── `08. Open Weight Foundation Models Llama, Mistral, Deepseek`
│   ├── `09. Model Evaluation Benchmarks Mmlu, Humaneval, Mt Bench`
│   └── `10. Continuous Fine Tuning and Domain Adaptation in Enterprise`
├── `08. High-Performance Inference & LLMOps Infrastructure`
│   ├── `01. High Throughput Inference Engines Vllm, Tgi, Tensorrt LLM`
│   ├── `02. Pagedattention and Continuous Iteration Level Batching`
│   ├── `03. Flashattention Kernels and Memory Optimization`
│   ├── `04. Speculative Decoding and Multi Token Prediction`
│   ├── `05. Distributed Multi GPU Inference Tensor and Pipeline Parallelism`
│   ├── `06. Llmops Telemetry, Tracing, and Observability Langfuse`
│   ├── `07. Production LLM Guardrails and Content Moderation`
│   ├── `08. Semantic Caching and Intelligent Model Routing`
│   ├── `09. GPU Resource Scheduling and Autoscaling in Kubernetes`
│   └── `10. Disaster Recovery, Fallbacks, and High Availability for LLMs`
├── `09. Data Engineering, ETL Pipelines & Lakehouse Architecture`
│   ├── `01. Modern Data Stack Architecture and Paradigms`
│   ├── `02. Batch vs Streaming Data Processing Paradigms`
│   ├── `03. Apache Spark Distributed Computation Internals`
│   ├── `04. Stateful Stream Processing with Apache Flink`
│   ├── `05. Lakehouse Table Formats Delta Lake, Iceberg, and Hudi`
│   ├── `06. Data Pipeline Orchestration Airflow, Dagster, and Prefect`
│   ├── `07. Dimensional Data Modeling and Star Schemas`
│   ├── `08. Data Transformation and Modeling with dbt`
│   ├── `09. Data Quality, Anomaly Detection, and Great Expectations`
│   └── `10. Data Governance, Lineage, and Metadata Catalogs`
└── `10. AI Data Science, Statistical Modeling & Econometrics`
│   ├── `01. Advanced Probability and Bayesian Statistical Modeling`
│   ├── `02. Inferential Statistics and Rigorous Hypothesis Testing`
│   ├── `03. Econometrics and Causal Inference Methodologies`
│   ├── `04. Exploratory Data Analysis EDA and Advanced Feature Engineering`
│   ├── `05. Time Series Analysis and Forecasting Models`
│   ├── `06. Advanced Dimensionality Reduction and Manifold Learning`
│   ├── `07. Advanced Clustering and Density Estimation`
│   ├── `08. Machine Learning Model Interpretability and Explainable AI`
│   ├── `09. Experimentation Platforms and Statistical a B Testing`
│   └── `10. Automated Machine Learning AutoML and Hyperparameter Search`
```

---

## 🤖 Core Knowledge Pillars

### 1. 📂 `01. Machine Learning, Deep Learning & Transformer Foundations`
- 📂 `01. Mathematical Foundations for Machine Learning` — Matrix rank, SVD, eigenvalues, partial derivatives, chain rule, Bayes' theorem, expectation, variance, and probability distributions.
- 📂 `02. Supervised Learning Algorithms and Regression` — Cost functions (MSE, Cross-Entropy), gradient descent variants, decision trees, Random Forests, Gradient Boosted Trees (XGBoost, LightGBM), and bias-variance tradeoff.
- 📂 `03. Unsupervised Learning, Clustering, and PCA` — K-Means++, Hierarchical clustering, DBSCAN density clustering, PCA dimensionality reduction, and reconstruction loss.
- 📂 `04. Neural Network Architectures and Backpropagation` — Perceptrons, MLPs, activation functions (ReLU, GELU, Swish), computational graphs, reverse-mode automatic differentiation, and vanishing gradients.
- 📂 `05. Convolutional and Recurrent Architectures` — Convolution kernels, pooling, receptive fields, ResNet skip connections, LSTM cell gates (forget, input, output), and vanishing gradients in sequential data.
- 📂 `06. Attention Mechanism and Transformer Architecture` — Query-Key-Value formulation, Scaled Dot-Product Attention, Multi-Head Attention, Positional Encodings (Sinusoidal, RoPE, ALiBi), and Encoder-Decoder stacks.
- 📂 `07. Large Language Model Pre Training and Scaling Laws` — Causal language modeling (Next-Token Prediction), masked language modeling, compute-optimal training tokens (Chinchilla), compute scaling laws, and emergent abilities.
- 📂 `08. Subword Tokenization Algorithms BPE and Wordpiece` — Vocabulary construction, byte-level BPE, merge tables, out-of-vocabulary handling, tokenization overhead across languages, and token count optimization.
- 📂 `09. Deep Learning Optimization Algorithms and Schedulers` — Stochastic Gradient Descent with Momentum, Adam, AdamW weight decay, cosine annealing schedules, warmup steps, and gradient clipping.
- 📂 `10. Loss Functions and Regularization Techniques` — Categorical cross-entropy, focal loss for imbalanced classes, L1/L2 regularization, dropout, LayerNorm vs RMSNorm, and batch normalization.
- 📂 `11. Model Evaluation Metrics and Validation Strategies` — Confusion matrices, precision-recall curves, ROC-AUC score, perplexity in language models, stratified k-fold cross-validation, and data leakage prevention.
### 2. 📂 [[Prompt Engineering, Context Optimization & LLM Alignment|02. Prompt Engineering, Context Optimization & LLM Alignment]]
- 📂 `01. Prompt Engineering Foundations and Best Practices` — System prompts, user instructions, few-shot exemplars, formatting delimiters (XML, Markdown), context priming, and temperature/top-p parameter tuning.
- 📂 `02. In Context Learning and Few Shot Prompting` — Demonstration formatting, diversity in exemplar selection, ordering effects, negative examples, and dynamic in-context example retrieval.
- 📂 `03. Chain of Thought CoT and Reasoning Elicitation` — Zero-shot CoT ('Let's think step by step'), Manual Few-Shot CoT, Least-to-Most prompting, and mitigating cascading reasoning hallucinations.
- 📂 `04. Advanced Reasoning Tree of Thoughts and Graph of Thoughts` — Decomposing complex problems into thought units, BFS/DFS search across thought trees, heuristic evaluation, back-tracking, and graph-based synthesis.
- 📂 `05. Context Window Management and Kv Cache Optimization` — Effective context utilization, positioning critical facts at the beginning/end, needle-in-a-haystack retrieval, and prompt caching (Anthropic, OpenAI).
- 📂 `06. Structured Outputs and Json Schema Enforcement` — Enforcing valid JSON schemas (OpenAI Strict JSON mode, Outlines, Instructor), BNF grammar-constrained decoding, and zero-validation-failure workflows.
- 📂 `07. Rlhf, Dpo, and Modern LLM Alignment` — Reward modeling, PPO policy optimization, Direct Preference Optimization (DPO) without separate reward models, Constitutional AI, and alignment tax.
- 📂 `08. System Prompts, Guardrail Directives, and Persona Design` — Designing production system prompts, setting defensive boundaries, tone and style calibration, tool access permissions, and preventing persona drift.
- 📂 `09. Prompt Security, Jailbreak Defenses, and Leakage Prevention` — Defending against delimiter escaping, indirect prompt injection from untrusted web text, system prompt extraction attacks, and input sanitization layers.
### 3. 📂 `03. AI Engineering, Vector Databases & RAG Architectures`
- 📂 `01. Foundation Model APIs and Multi Provider Gateways` — REST/SDK integrations, streaming server-sent events (SSE), fallback routing (LiteLLM), token usage tracking, and multi-provider load balancing.
- 📂 `02. Dense Embeddings, Similarity Metrics, and Quantization` — Bi-encoder architectures (text-embedding-3, BGE), Cosine similarity vs Euclidean distance, Matryoshka representation learning, and vector scalar/binary quantization.
- 📂 `03. Vector Database Internals and Indexing Algorithms` — Hierarchical Navigable Small World (HNSW) graphs, Inverted File with Product Quantization (IVF-PQ), recall vs latency trade-offs, and memory footprints.
- 📂 `04. Vector Databases in Production Qdrant, Milvus, Pgvector` — Architecture comparison, hybrid filtering, distributed sharding, persistence, and choosing between dedicated vector DBs vs pgvector in PostgreSQL.
- 📂 `05. Document Ingestion and Chunking Strategies` — Parsing PDF/HTML/Markdown, fixed-size with overlap, sentence window chunking, recursive character chunking, and semantic chunking based on embedding distance.
- 📂 `06. Dense vs Sparse Hybrid Search and Bm25` — Combining lexical exact-match BM25 with semantic vector search, Reciprocal Rank Fusion (RRF), alpha parameter tuning, and avoiding semantic search blindspots.
- 📂 `07. Cross Encoder Reranking and Contextual Compression` — Two-stage retrieval pipeline: High-recall vector search -> High-precision Cross-Encoder reranking, and removing irrelevant tokens with contextual compression.
- 📂 `08. Advanced RAG Architectures Self RAG and Corrective RAG` — Self-reflection tokens, dynamic retrieval decisions, web search fallbacks when retrieved docs are insufficient, and query reformulation.
- 📂 `09. Graphrag and Knowledge Graph Augmented Retrieval` — Extracting entity-relationship graphs from documents (Neo4j, NetworkX), community detection summaries, combining vector search with graph traversal.
- 📂 `10. Multi Modal Retrieval and Document Parsing` — Retrieving text and images simultaneously, ColPali vision-language embeddings for PDF pages, and extracting complex structured tables.
- 📂 `11. RAG Evaluation Frameworks Ragas and Trulens` — Measuring Faithfulness (groundedness), Answer Relevance, Context Precision, and Context Recall; automated synthetic test set generation.
- 📂 `12. Semantic Caching and LLM Response Optimization` — Caching LLM responses by semantic embedding similarity threshold, exact-match prompt prefix caching, and slashing API latency and dollar costs.
- 📂 `13. Token Economics, Budgeting, and Cost Governance` — Tracking input/output token costs across teams, model routing based on query complexity (small fast model -> heavy frontier model), and budgeting guardrails.
### 4. 📂 `04. AI Agents, Autonomous Systems & Multi-Agent Orchestration`
- 📂 `01. AI Agent Architecture and the ReAct Paradigm` — The autonomous loop: Observe -> Think -> Act -> Observe; state management, tool execution cycle, and breaking complex objectives into subgoals.
- 📂 `02. Tool Calling and Function Calling Specifications` — Declaring JSON tool schemas, deterministic tool execution, handling multiple parallel tool calls, and error handling when tool invocations fail.
- 📂 `03. Agent Memory Systems Short Term, Long Term, Episodic` — Sliding window conversation buffers, summary buffers, long-term semantic memory in vector stores, and episodic memory retrieval for past successes.
- 📂 `04. Planning Algorithms Plan and Solve and Reflexion` — Static planning vs dynamic re-planning, verbal reinforcement learning via Reflexion, analyzing execution trace failures, and iterative self-repair.
- 📂 `05. Langgraph State Machine Architecture and Workflows` — Building cyclical agent graphs in LangGraph, state schemas, conditional edges, checkpointing for state persistence, and time-travel debugging.
- 📂 `06. Llamaindex Agent Workflows and Query Engines` — Router query engines, sub-question query engines, multi-document agent workflows, and combining retrieval tools with execution actions.
- 📂 `07. Crewai Multi Agent Role Playing Collaboration` — Defining agent personas (Role, Goal, Backstory), task dependencies, sequential vs hierarchical process execution, and inter-agent communication.
- 📂 `08. Autogen Conversational Patterns and Group Chats` — Multi-agent conversational patterns, UserProxy agents, group chat round-robin vs LLM-directed speaker selection, and code execution environments.
- 📂 `09. Human in the Loop HITL and Approval Gates` — Pausing agent execution before high-risk actions (e.g. database mutations, sending emails), human feedback injection, and resuming graph state.
- 📂 `10. Agent Code Sandboxes and Secure Execution Runtimes` — Isolating untrusted AI-generated code execution, ephemeral microVM sandboxes (E2B), network egress restrictions, CPU/memory limits, and timeout controls.
- 📂 `11. Error Recovery, Loop Detection, and Self Healing Agents` — Detecting repetitive cycling in tool calls, max-iteration caps, fallback strategies, and prompt re-anchoring when agents get stuck.
- 📂 `12. Guardrails, Safety Filters, and Output Governance` — Restricting allowed actions, parameter range validation, content moderation filters, and preventing unauthorized system access by autonomous agents.
- 📂 `13. Autonomous Web Browsing and Computer Use Agents` — DOM tree parsing for LLMs, accessibility tree navigation, coordinate-based clicking and typing (Anthropic Computer Use), and CAPTCHA handling.
- 📂 `14. Agent Observability, Tracing, and Evaluation` — Tracing nested tool calls and reasoning chains, measuring task success rate, tracking token costs per agent run, and benchmark evaluations (GAIA, SWE-bench).
### 5. 📂 `05. Claude Code, Subagents, Model Context Protocol MCP & Tooling`
- 📂 `01. Claude Code Architecture and Terminal Execution Engine` — Autonomous terminal agent design, real-time command execution, file system read/write primitives, and streaming token response handling.
- 📂 `02. Model Context Protocol MCP Standards and Architecture` — The MCP client-server standard, JSON-RPC 2.0 transport (stdio, SSE), Resources, Tools, and Prompts primitives, and ecosystem integration.
- 📂 `03. Developing Custom MCP Servers in Typescript and Python` — Exposing enterprise internal APIs, databases, and internal tools as MCP endpoints; parameter schema validation, and error serialization.
- 📂 `04. Subagent Delegation, Swarms, and Context Handoffs` — Spawning specialized subagents for isolated tasks, passing constrained context packets, aggregating subagent results, and preventing context window bloat.
- 📂 `05. Multi File Code Synthesis and Atomic Patching` — Generating consistent multi-file diffs, applying targeted replacements without corrupting whitespace, syntax tree validation, and rolling back failed edits.
- 📂 `06. Terminal and Tool Execution Safety and Sandboxing` — Detecting dangerous shell commands (`rm -rf`, `DROP TABLE`), requiring explicit confirmation, read-only modes, and sandbox containment.
- 📂 `07. Conversation History Compaction and Transcript Management` — Pruning large terminal tool outputs, sliding window transcript compression, maintaining critical state invariants across long multi-turn sessions.
- 📂 `08. Custom Slash Commands and Background Execution` — Extending agent CLI capabilities with custom slash commands (`/goal`, `/schedule`), background async tasks, and reactive wakeups.
- 📂 `09. Skill and Rule Customization Engine and Discovery` — Modular skill definitions (`SKILL.md`), hierarchical rule loading (global vs workspace), plugin architectures, and dynamic on-demand capability injection.
- 📂 `10. Automated PR Generation and CI Integration with Claude Code` — Running Claude Code in non-interactive CI/CD pipelines, automated bug fixing from stack traces, generating clean PR diffs with testing evidence.
- 📂 `11. Production Codebase Refactoring with Autonomous Agents` — Orchestrating multi-day codebase upgrades (e.g. library migrations, TypeScript strict mode conversion), test-driven verification, and tracking progress.
- 📂 `12. Debugging Autonomous Agent Failure Modes and Hallucinations` — Analyzing agent decision transcripts, diagnosing incorrect tool argument generation, debugging prompt injection leaks, and calibrating agent prompts.
- 📂 `13. Benchmark Evaluations for Coding Agents Swe Bench` — Setting up SWE-bench evaluation harnesses, measuring pass@1 and pass@5 resolution rates on real GitHub issues, and regression testing agent capabilities.
### 6. 📂 `06. AI-Assisted Engineering, Cursor, Copilot & Vibe Coding`
- 📂 `01. The AI Assisted Software Engineering Paradigm` — Shifting developer cognitive bandwidth from syntax typing to architectural specification, critical review, test verification, and high-level system design.
- 📂 `02. Cursor IDE Architecture, Composer, and Multi File Edits` — Using Cursor Composer for cross-file feature scaffolding, `@Files`, `@Codebase`, and `@Docs` context indexing, and reviewing side-by-side diffs.
- 📂 `03. Github Copilot, Inline Completions, and Workspace Indexing` — Optimizing inline completion acceptance rates, steering completions with descriptive comments and function signatures, and using Copilot in CLI.
- 📂 `04. Plan Before You Code and Specification Driven Development` — Writing exhaustive implementation plans (`implementation_plan.md`) before generating code, validating architecture upfront, and preventing AI drift.
- 📂 `05. Context Engineering .cursorrules and Repository Maps` — Configuring project-specific rules (`.cursorrules`), defining style invariants, technology constraints, and generating repo maps for LLM context.
- 📂 `06. AI Generated Unit, Integration, and Property Tests` — Using AI to generate edge-case test matrices, mocking external dependencies, generating property-based test fuzzers, and ensuring 100% test pass rates.
- 📂 `07. Automated Pull Request Reviewers and Agentic CI` — Embedding AI PR review bots in GitHub Actions, detecting security flaws, missing test cases, and performance regressions before human review.
- 📂 `08. Rapid Prototyping and Greenfield Project Scaffolding` — Going from PRD to running prototype in hours, generating database schemas, API routes, UI components, and authentication scaffolding.
- 📂 `09. Detecting and Mitigating AI Hallucinations in Code` — Spotting hallucinated library methods, verifying imports against installed packages, running local compiler checks, and static analysis verification.
- 📂 `10. Full Stack Vibe Coding Ergonomics and Velocity` — The 'Vibe Coding' methodology: Natural language iteration, rapid UI tweaking with hot-reload, fast-paced feature exploration, and refactoring to standards.
- 📂 `11. Refactoring Legacy Codebases with AI Pair Programming` — Converting legacy monolithic code into clean modular architectures, writing missing test harnesses first, and incrementally migrating frameworks.
- 📂 `12. Developer Ergonomics, Cognitive Flow, and AI Tooling Fatigue` — Balancing AI code generation with deep human understanding, avoiding rubber-stamp code reviews, and maintaining long-term code mastery.
### 7. 📂 [[Large Language Model Fine Tuning & Quantization|07. Large Language Model Fine-Tuning & Quantization]]
- 📂 `01. Fine Tuning Paradigms Full vs Parameter Efficient PEFT` — Why full fine-tuning is computationally prohibitive, Low-Rank Adaptation (LoRA) low-rank matrices decomposition, and QLoRA 4-bit NormalFloat quantization.
- 📂 `02. SFT Dataset Preparation, Curation, and Quality Filtering` — Structuring instruction-response pairs (Alpaca, ShareGPT formats), synthetic data generation with frontier models, deduplication (MinHash), and quality filtering.
- 📂 `03. Supervised Fine Tuning SFT Pipelines with Axolotl and Unsloth` — Accelerating training with Unsloth fast kernels, configuring Axolotl YAML training configs, gradient accumulation, FlashAttention-2 integration, and loss curves.
- 📂 `04. Direct Preference Optimization DPO and Alignment Tuning` — Preparing chosen vs rejected preference datasets, tuning the beta hyperparameter, eliminating separate reward models, and stabilizing preference training.
- 📂 `05. Post Training Quantization Gguf, Awq, and GPTQ` — Comparing quantization schemes: Weight-only vs Weight-and-Activation quantization, 4-bit AWQ activation-aware quantization, and GGUF for CPU/Apple Silicon.
- 📂 `06. Quantization Aware Training QAT and Mixed Precision` — Simulating quantization noise during forward passes, mixed-precision BF16 training stability, and hardware acceleration on NVIDIA Hopper/Blackwell FP8.
- 📂 `07. Model Merging Techniques with Mergekit` — Combining capabilities of multiple fine-tuned models without training: Spherical Linear Interpolation (SLERP), TIES-Merging, and DARE sparsification.
- 📂 `08. Open Weight Foundation Models Llama, Mistral, Deepseek` — Architectural innovations: SwiGLU activations, RoPE base scaling, Mixture-of-Experts (MoE) routing, Multi-Head Latent Attention (MLA), and license analysis.
- 📂 `09. Model Evaluation Benchmarks Mmlu, Humaneval, Mt Bench` — Running standardized evaluation harnesses (lm-evaluation-harness), measuring reasoning, coding accuracy, mathematical capability, and multi-turn chat quality.
- 📂 `10. Continuous Fine Tuning and Domain Adaptation in Enterprise` — Preventing catastrophic forgetting by mixing replay data, domain-specific vocabulary extension, and continuous retraining pipelines for enterprise data.
### 8. 📂 `08. High-Performance Inference & LLMOps Infrastructure`
- 📂 `01. High Throughput Inference Engines Vllm, Tgi, Tensorrt LLM` — Inference engine architecture comparison, server throughput benchmarks, streaming API protocols, and optimizing GPU utilization.
- 📂 `02. Pagedattention and Continuous Iteration Level Batching` — Solving memory fragmentation in KV-cache with PagedAttention (OS-style paging), eliminating static batching wait times via continuous iteration-level batching.
- 📂 `03. Flashattention Kernels and Memory Optimization` — IO-aware exact attention algorithms, avoiding slow HBM read/writes by tiling QKV computations in GPU SRAM, and FP8 FlashAttention-3 acceleration.
- 📂 `04. Speculative Decoding and Multi Token Prediction` — Accelerating autoregressive decoding by having a lightweight draft model speculate tokens and a frontier model verify them in a single forward pass.
- 📂 `05. Distributed Multi GPU Inference Tensor and Pipeline Parallelism` — Splitting matrix multiplications across GPUs with Megatron-LM Tensor Parallelism (TP), Pipeline Parallelism (PP) layer partitioning, and NCCL communication.
- 📂 `06. Llmops Telemetry, Tracing, and Observability Langfuse` — End-to-end distributed tracing of LLM inference requests, token latency breakdown (Time-To-First-Token TTFT vs Inter-Token-Latency ITL), and prompt versioning.
- 📂 `07. Production LLM Guardrails and Content Moderation` — Deploying real-time classifier guardrails (Llama Guard, NeMo), PII anonymization before model inference, hallucination checking, and prompt leak defense.
- 📂 `08. Semantic Caching and Intelligent Model Routing` — Embedding-based similarity response caching, tiered model routing (simple queries -> 8B model, complex reasoning -> 70B/Frontier model), and cost optimization.
- 📂 `09. GPU Resource Scheduling and Autoscaling in Kubernetes` — Sharing GPU resources with Time-Slicing vs Multi-Instance GPU (MIG), autoscaling inference pods based on vLLM queue depth metrics using KEDA.
- 📂 `10. Disaster Recovery, Fallbacks, and High Availability for LLMs` — Designing multi-region inference clusters, automatic failovers to cloud APIs during local GPU outages, graceful degradation to smaller models, and SLAs.
### 9. 📂 `09. Data Engineering, ETL Pipelines & Lakehouse Architecture`
- 📂 `01. Modern Data Stack Architecture and Paradigms` — Evolution from on-premise ETL to cloud-native ELT (Snowflake, BigQuery, Databricks), decoupling compute from storage, and modular data architectures.
- 📂 `02. Batch vs Streaming Data Processing Paradigms` — Windowing strategies (Tumbling, Sliding, Session), event time vs processing time, watermarks for late-arriving data, and Kappa streaming architecture.
- 📂 `03. Apache Spark Distributed Computation Internals` — Distributed in-memory computation, Catalyst query optimization, Tungsten binary memory management, shuffle partitioning, and memory spill debugging.
- 📂 `04. Stateful Stream Processing with Apache Flink` — Distributed asynchronous snapshotting (Chandy-Lamport), state backends (RocksDB), keyed streams, complex event processing (CEP), and exactly-once guarantees.
- 📂 `05. Lakehouse Table Formats Delta Lake, Iceberg, and Hudi` — ACID transactions on cloud object storage (S3), metadata layers, time travel and snapshot isolation, partition evolution without rewrites, and file compaction.
- 📂 `06. Data Pipeline Orchestration Airflow, Dagster, and Prefect` — Directed Acyclic Graphs (DAGs), asset-based orchestration in Dagster, declarative dependencies, dynamic retries, backfilling historic data, and SLA alerts.
- 📂 `07. Dimensional Data Modeling and Star Schemas` — Fact tables (transactional, periodic snapshot, accumulating), dimension tables, Slowly Changing Dimensions (SCD Type 1, 2, 3), and Data Vault 2.0.
- 📂 `08. Data Transformation and Modeling with dbt` — Writing modular SQL transformations, Jinja macros, ephemeral/view/table materializations, automated schema testing, and documentation generation in dbt.
- 📂 `09. Data Quality, Anomaly Detection, and Great Expectations` — Automated data validation assertions (null checks, uniqueness, distribution bounds), CI data quality gates, anomaly detection in pipeline runs, and data SLAs.
- 📂 `10. Data Governance, Lineage, and Metadata Catalogs` — Tracking automated data lineage across Spark/dbt/Airflow, data discovery catalogs (DataHub, Amundsen), column-level access control, and GDPR/CCPA compliance.
### 10. 📂 `10. AI Data Science, Statistical Modeling & Econometrics`
- 📂 `01. Advanced Probability and Bayesian Statistical Modeling` — Prior, likelihood, and posterior distributions; conjugate priors, Maximum A Posteriori (MAP), MCMC sampling algorithms (Gibbs, Metropolis-Hastings).
- 📂 `02. Inferential Statistics and Rigorous Hypothesis Testing` — Null hypothesis significance testing (NHST), Student's t-test, ANOVA, Chi-Square, Mann-Whitney U test, p-values, Type I/II errors, and power analysis.
- 📂 `03. Econometrics and Causal Inference Methodologies` — Correlation vs causation, potential outcomes framework, Difference-in-Differences (DiD), Propensity Score Matching, regression discontinuity, and DAGs.
- 📂 `04. Exploratory Data Analysis EDA and Advanced Feature Engineering` — Detecting outliers (IQR, Isolation Forests), handling missing data (MICE), feature encoding (Target, One-Hot), polynomial features, and feature scaling.
- 📂 `05. Time Series Analysis and Forecasting Models` — Stationarity testing (Augmented Dickey-Fuller), auto-correlation (ACF/PACF), seasonal decomposition, Facebook Prophet, and deep learning forecasting (DeepAR).
- 📂 `06. Advanced Dimensionality Reduction and Manifold Learning` — Linear vs non-linear dimensionality reduction, preserving local vs global data topology with UMAP, and visualizing high-dimensional embeddings.
- 📂 `07. Advanced Clustering and Density Estimation` — Soft clustering with Expectation-Maximization in GMMs, hierarchical density-based clustering (HDBSCAN), and Kernel Density Estimation (KDE).
- 📂 `08. Machine Learning Model Interpretability and Explainable AI` — Game-theoretic Shapley values with TreeSHAP/KernelSHAP, Local Interpretable Model-agnostic Explanations (LIME), and feature importance plots.
- 📂 `09. Experimentation Platforms and Statistical a B Testing` — Sample size determination, avoiding multiple testing pitfalls (Bonferroni, FDR), variance reduction via CUPED, and sequential testing algorithms.
- 📂 `10. Automated Machine Learning AutoML and Hyperparameter Search` — Hyperparameter optimization (Optuna, Hyperopt), Bayesian Optimization with Gaussian Processes, tree-structured Parzen estimators (TPE), and model search.

---

## 🔗 Navigation
- ⬆️ Parent: [[Principal SWE]]
- 💻 Computer Science Foundations: `Computer Science`
- 🏛️ Software Architecture: `Architecture`
- 🚀 Infrastructure & DevOps: `DevOps`
- 🛡️ Cyber Security: `Cyber Security`

