---
title: AI & Machine Learning
tags:
  - ai
  - machine-learning
  - llm
  - agents
  - principal-swe
parent: "[[Principal SWE]]"
---

# 🤖 AI Engineering, Machine Learning & Autonomous Agents

Comprehensive, production-grade master architecture covering the complete spectrum of modern artificial intelligence and Large Language Model engineering: Machine Learning & Deep Learning foundations, systematic Prompt Engineering & LLM alignment, production Retrieval-Augmented Generation (RAG) and vector search spaces, autonomous AI Agent architectures & Model Context Protocol (MCP), Claude Code CLI & subagent automation, and AI-assisted spec-driven software engineering (Vibe Coding) across 6 master pillars and 60 specialized subdomains.

```text
AI & Machine Learning
│
├── [[Machine Learning & Deep Learning Foundations|01. Machine Learning & Deep Learning Foundations]]
│   ├── [[Mathematical Foundations for Machine Learning|01. Mathematical Foundations for ML]]
│   ├── [[Programming and Numerical Computing Foundations|02. Programming and Numerical Foundations]]
│   ├── [[Data Sourcing, Ingestion, and Quality Assurance|03. Data Sourcing and Ingestion]]
│   ├── [[Data Preprocessing, Cleaning, and Feature Engineering|04. Data Preprocessing and Feature Engineering]]
│   ├── [[Classical Machine Learning Algorithms|05. Classical Machine Learning Algorithms]]
│   ├── [[Model Evaluation Metrics and Validation Strategies|06. Model Evaluation and Validation]]
│   ├── [[Deep Learning and Neural Network Architectures|07. Deep Learning and Neural Architectures]]
│   ├── [[Transformer Architecture and Self Attention Mechanics|08. Transformer Architecture and Attention]]
│   └── [[Advanced ML Optimization and Regularization Techniques|09. Advanced ML Training Techniques]]
├── [[Prompt Engineering & LLM Alignment|02. Prompt Engineering & LLM Alignment]]
│   ├── [[LLM Foundations and Tokenization Mechanics|01. LLM Terminology and Tokenization]]
│   ├── [[LLM Inference Parameters and Sampling Calibration|02. LLM Inference Parameters and Sampling]]
│   ├── [[Advanced Prompting Techniques (cot, React, Few Shot)|03. Advanced Prompting Techniques]]
│   ├── [[Structured Outputs, JSON Mode, and Function Calling|04. Structured Outputs and Function Calling]]
│   ├── [[Defensive Prompting and System Prompt Hardening|05. AI Red Teaming and Defensive Prompting]]
│   ├── [[Prompt Optimization, Versioning, and Evaluation|06. Prompt Optimization and Versioning]]
│   └── [[Improving LLM Reliability and Hallucination Reduction|07. Improving Output Reliability and Consistency]]
├── [[AI Engineering & RAG Systems|03. AI Engineering & RAG Systems]]
│   ├── [[Foundation Model Ecosystem and Cloud APIs|01. Foundation Models and Cloud APIs]]
│   ├── [[Open Source Models and Local High Throughput Serving|02. Open Source Models and Local Serving]]
│   ├── [[Embedding Models and High Dimensional Vector Spaces|03. Embeddings and Vector Search Spaces]]
│   ├── [[Vector Databases and HNSW Index Architectures|04. Vector Databases and HNSW Indexing]]
│   ├── [[Document Parsing, Chunking, and Metadata Enrichment|05. Document Parsing and Chunking Strategies]]
│   ├── [[Advanced Retrieval, Hybrid Search, and Re Ranking|06. Advanced Retrieval and Hybrid Search]]
│   ├── [[Context Compression, Lost in the Middle, and Caching|07. Context Compression and Context Scaling]]
│   ├── [[Multimodal AI Architectures (vision, Audio, Video)|08. Multimodal AI Architectures]]
│   ├── [[AI Code Generation and Developer Assistant Tooling|09. AI Code Generation and Assistant Tooling]]
│   ├── [[RAG Evaluation Frameworks (ragas, Trulens)|10. RAG Evaluation Frameworks]]
│   └── [[Production RAG Architecture and Semantic Caching|11. Production RAG Architecture and Caching]]
├── [[AI Agents & Multi Agent Architectures|04. AI Agents & Multi-Agent Architectures]]
│   ├── [[AI Agent Fundamentals and Execution State Machines|01. Agent Fundamentals and State Machines]]
│   ├── [[Tool Use, API Orchestration, and Sandboxed Execution|02. Tool Use, Apis, and Function Execution]]
│   ├── [[Model Context Protocol (mcp) Standards and Architecture|03. Model Context Protocol MCP Standards]]
│   ├── [[Agent Short Term, Long Term, and Working Memory|04. Agent Short and Long Term Memory]]
│   ├── [[Multi Agent Collaboration, Hierarchical Delegation, and Swarms|05. Multi Agent Collaboration and Swarms]]
│   ├── [[Agent Orchestration Frameworks (langgraph, Crewai)|06. Agent Frameworks and Orchestrators]]
│   ├── [[Agent Planning, Reflection, and Self Correction Loops|07. Planning, Reflection, and Self Correction]]
│   ├── [[Agent Observability, Distributed Tracing, and Debugging|08. Agent Observability, Tracing, and Debugging]]
│   ├── [[Agent Evaluation Benchmarks and Tool Call Evals|09. Evaluation and Benchmarking for Agents]]
│   ├── [[Agent Safety, Privilege Escalation, and Sandboxing|10. Agent Safety, Sandboxing, and Guardrails]]
│   ├── [[Autonomous Software Engineering Agents|11. Autonomous Code Generation Agents]]
│   └── [[Voice and Real Time Streaming Interactive Agents|12. Voice and Real Time Interactive Agents]]
├── [[Claude Code, Subagents & Workflows|05. Claude Code, Subagents & Workflows]]
│   ├── [[Claude Code CLI Architecture and Core Mechanics|01. Claude Code CLI Architecture]]
│   ├── [[Claude.md Context Engineering and System Prompts|02. Claude.md Context Engineering]]
│   ├── [[Custom Skills Architecture (skill.md) and Tool Extensions|03. Custom Skills and Tool Extensions]]
│   ├── [[Subagent Parallelism and Task Delegation|04. Subagent Delegation and Parallel Execution]]
│   ├── [[Custom Lifecycle Hooks (hooks.json) and Shell Automations|05. Custom Lifecycle Hooks and Automations]]
│   ├── [[Context Window Management and Token Compaction|06. Context Management and Token Compaction]]
│   ├── [[Test Driven Development (tdd) with Claude Code|07. Test Driven Development with Claude Code]]
│   ├── [[Large Scale Repository Refactoring with Claude Code|08. Large Scale Refactoring Workflows]]
│   ├── [[Claude Code Permission Controls and Command Sandboxing|09. Security Sandboxing and Permission Controls]]
│   ├── [[Headless Claude Code Execution and Ci-cd Integrations|10. Ci/cd and Headless Automation with Claude]]
│   └── [[Advanced Claude Code Workflows, Slash Commands, and Custom Tools|11. Advanced Claude Code Workflows and Patterns]]
└── [[AI Assisted Development & Vibe Coding|06. AI-Assisted Development & Vibe Coding]]
│   ├── [[Vibe Coding Philosophy and Spec Driven Architecture|01. Vibe Coding Philosophy and Spec Driven Design]]
│   ├── [[Planning Before Coding and Technical Design Artifacts|02. Planning Before Coding and Implementation Plans]]
│   ├── [[Context Optimization in AI IDEs (cursor, Windsurf, Antigravity)|03. Context Optimization in AI Coding IDEs]]
│   ├── [[Effective Prompting Patterns for Software Engineers|04. Prompting Patterns for Software Engineering]]
│   ├── [[AI Driven Debugging, Log Parsing, and Root Cause Analysis|05. AI Driven Debugging and Root Cause Analysis]]
│   ├── [[Git Mastery for AI Pair Programming and Micro Commits|06. Git and Version Control for AI Pair Programming]]
│   ├── [[Automated Test Generation and Boundary Fuzzing with AI|07. Automated Test Generation with AI]]
│   ├── [[Security Verification and Code Quality for AI Generated Code|08. Security Verification for AI Generated Code]]
│   ├── [[Domain Modeling, Schema Design, and Code Refactoring with AI|09. Domain Modeling and Refactoring with AI]]
│   └── [[The Future of Principal Engineering in the AI Era|10. Future of Software Engineering with AI]]
```

---

## 🏛️ Core Knowledge Pillars

### 1. 📂 [[Machine Learning & Deep Learning Foundations|01. Machine Learning & Deep Learning Foundations]]
- 📂 [[Mathematical Foundations for Machine Learning|01. Mathematical Foundations for ML]] — Multivariate calculus, matrix decompositions (SVD, Eigendecomposition), probability theory, and optimization gradients (SGD, AdamW).
- 📂 [[Programming and Numerical Computing Foundations|02. Programming and Numerical Foundations]] — Vectorized tensor computations in NumPy and PyTorch, broadcast semantics, and GPU acceleration abstractions.
- 📂 [[Data Sourcing, Ingestion, and Quality Assurance|03. Data Sourcing and Ingestion]] — Feature collection, data labeling pipelines, dataset balancing, and data distribution validation.
- 📂 [[Data Preprocessing, Cleaning, and Feature Engineering|04. Data Preprocessing and Feature Engineering]] — Handling missing values, one-hot encoding, normalization/standardization, dimensionality reduction (PCA), and embeddings.
- 📂 [[Classical Machine Learning Algorithms|05. Classical Machine Learning Algorithms]] — Linear regression, Logistic regression, Decision Trees, Random Forests, Gradient Boosted Trees (XGBoost), and SVMs.
- 📂 [[Model Evaluation Metrics and Validation Strategies|06. Model Evaluation and Validation]] — Confusion matrix, Precision, Recall, F1-Score, ROC-AUC, cross-validation, and bias-variance trade-off.
- 📂 [[Deep Learning and Neural Network Architectures|07. Deep Learning and Neural Architectures]] — Multi-Layer Perceptrons (MLP), Convolutional Neural Networks (CNNs), Recurrent Networks (RNN/LSTM), and Residual Networks (ResNet).
- 📂 [[Transformer Architecture and Self Attention Mechanics|08. Transformer Architecture and Attention]] — Scaled Dot-Product Attention, Multi-Head Attention, positional encodings, Encoder-Decoder topologies, and GPT decoders.
- 📂 [[Advanced ML Optimization and Regularization Techniques|09. Advanced ML Training Techniques]] — Dropout, Batch Normalization, LayerNorm, learning rate scheduling, mixed precision training (FP16/BF16), and LoRA fine-tuning.
### 2. 📂 [[Prompt Engineering & LLM Alignment|02. Prompt Engineering & LLM Alignment]]
- 📂 [[LLM Foundations and Tokenization Mechanics|01. LLM Terminology and Tokenization]] — Byte-Pair Encoding (BPE), SentencePiece, token context windows, context length scaling, and prompt token budget calculation.
- 📂 [[LLM Inference Parameters and Sampling Calibration|02. LLM Inference Parameters and Sampling]] — Temperature, Top-p (nucleus sampling), Top-k, frequency/presence penalties, seed determinism, and stop sequences.
- 📂 [[Advanced Prompting Techniques (cot, React, Few Shot)|03. Advanced Prompting Techniques]] — Chain-of-Thought (CoT), Reason+Act (ReAct), Few-Shot demonstrations, Tree-of-Thoughts (ToT), and Self-Consistency prompting.
- 📂 [[Structured Outputs, JSON Mode, and Function Calling|04. Structured Outputs and Function Calling]] — Constrained decoding, JSON schema enforcement, tool use parameter extraction, and deterministic structured formatting.
- 📂 [[Defensive Prompting and System Prompt Hardening|05. AI Red Teaming and Defensive Prompting]] — System prompt boundary isolation, delimiters, prompt injection defense, leak prevention, and adversarial testing.
- 📂 [[Prompt Optimization, Versioning, and Evaluation|06. Prompt Optimization and Versioning]] — Prompt as code, prompt evaluation suites, few-shot demonstration curation, and automated prompt tuning (DSPy).
- 📂 [[Improving LLM Reliability and Hallucination Reduction|07. Improving Output Reliability and Consistency]] — Self-refinement loops, majority voting, grounding assertions, cross-verification pipelines, and confidence scoring.
### 3. 📂 [[AI Engineering & RAG Systems|03. AI Engineering & RAG Systems]]
- 📂 [[Foundation Model Ecosystem and Cloud APIs|01. Foundation Models and Cloud APIs]] — OpenAI, Anthropic Claude, Google Gemini APIs; rate limits, token cost modeling, and fallback routing.
- 📂 [[Open Source Models and Local High Throughput Serving|02. Open Source Models and Local Serving]] — Quantization (GGUF, AWQ, GPTQ), vLLM continuous batching, PagedAttention, TensorRT-LLM, and Ollama.
- 📂 [[Embedding Models and High Dimensional Vector Spaces|03. Embeddings and Vector Search Spaces]] — Dense vector embeddings, cosine similarity, Euclidean distance, Dot product, and MTEB benchmark rankings.
- 📂 [[Vector Databases and HNSW Index Architectures|04. Vector Databases and HNSW Indexing]] — Hierarchical Navigable Small World (HNSW) graphs, IVF-Flat, Pinecone, Qdrant, Milvus, pgvector, and scaling billion-vector indexes.
- 📂 [[Document Parsing, Chunking, and Metadata Enrichment|05. Document Parsing and Chunking Strategies]] — Recursive character chunking, semantic chunking, document layout parsing, metadata extraction, and hierarchical indexing.
- 📂 [[Advanced Retrieval, Hybrid Search, and Re Ranking|06. Advanced Retrieval and Hybrid Search]] — BM25 keyword + Dense vector hybrid search (Reciprocal Rank Fusion RRF), cross-encoder re-rankers (Cohere), and HyDE.
- 📂 [[Context Compression, Lost in the Middle, and Caching|07. Context Compression and Context Scaling]] — Mitigating the Lost-in-the-Middle phenomenon, prompt compression (LLMLingua), prompt caching, and long-context RAG.
- 📂 [[Multimodal AI Architectures (vision, Audio, Video)|08. Multimodal AI Architectures]] — Vision-Language Models (VLMs), image embeddings (CLIP), speech-to-text (Whisper), and visual reasoning pipelines.
- 📂 [[AI Code Generation and Developer Assistant Tooling|09. AI Code Generation and Assistant Tooling]] — Repository-level AST context collection, fill-in-the-middle (FIM) code completion, and automated test synthesis.
- 📂 [[RAG Evaluation Frameworks (ragas, Trulens)|10. RAG Evaluation Frameworks]] — Faithfulness, Answer Relevance, Context Precision, and Context Recall metrics for automated RAG quality evaluation.
- 📂 [[Production RAG Architecture and Semantic Caching|11. Production RAG Architecture and Caching]] — Semantic cache layers (GPTCache), async ingestion workers, dead-letter indexing queues, and end-to-end RAG latency optimization.
### 4. 📂 [[AI Agents & Multi Agent Architectures|04. AI Agents & Multi-Agent Architectures]]
- 📂 [[AI Agent Fundamentals and Execution State Machines|01. Agent Fundamentals and State Machines]] — Perceive-Plan-Act-Reflect loops, deterministic state transitions, goal decomposition, and termination invariants.
- 📂 [[Tool Use, API Orchestration, and Sandboxed Execution|02. Tool Use, Apis, and Function Execution]] — Dynamic tool discovery, argument schema validation, secure sandboxed execution environments, and error recovery.
- 📂 [[Model Context Protocol (mcp) Standards and Architecture|03. Model Context Protocol MCP Standards]] — Anthropic MCP client-server protocol, standardized resource exposure, tool registration, and universal tool interoperability.
- 📂 [[Agent Short Term, Long Term, and Working Memory|04. Agent Short and Long Term Memory]] — Conversation context buffer, episodic memory, semantic memory graphs (Mem0), and entity relationship tracking.
- 📂 [[Multi Agent Collaboration, Hierarchical Delegation, and Swarms|05. Multi Agent Collaboration and Swarms]] — Manager-Worker hierarchies, peer-to-peer agent consensus, debate frameworks, and decentralized agent swarms.
- 📂 [[Agent Orchestration Frameworks (langgraph, Crewai)|06. Agent Frameworks and Orchestrators]] — State graph-based agent orchestration (LangGraph), role-playing crew models (CrewAI), and conversational AutoGen swarms.
- 📂 [[Agent Planning, Reflection, and Self Correction Loops|07. Planning, Reflection, and Self Correction]] — Plan-and-Solve, Reflexion, tree search planning, backtracking on tool failure, and iterative self-critique.
- 📂 [[Agent Observability, Distributed Tracing, and Debugging|08. Agent Observability, Tracing, and Debugging]] — Step-by-step agent execution tracing (Langfuse, Arize Phoenix), token attribution, cost tracking, and loop detection.
- 📂 [[Agent Evaluation Benchmarks and Tool Call Evals|09. Evaluation and Benchmarking for Agents]] — Tool call accuracy, goal completion rate, trajectory evaluation, and sandboxed coding benchmarks (SWE-bench).
- 📂 [[Agent Safety, Privilege Escalation, and Sandboxing|10. Agent Safety, Sandboxing, and Guardrails]] — Human-in-the-loop (HITL) authorization gates, blast radius minimization, container isolation, and safe tool execution.
- 📂 [[Autonomous Software Engineering Agents|11. Autonomous Code Generation Agents]] — Automated issue reproduction, repository search, patch generation, unit test verification, and automated PR authoring.
- 📂 [[Voice and Real Time Streaming Interactive Agents|12. Voice and Real Time Interactive Agents]] — WebRTC low-latency streaming, speech-to-speech models, interruption handling, and turn-taking state machines.
### 5. 📂 [[Claude Code, Subagents & Workflows|05. Claude Code, Subagents & Workflows]]
- 📂 [[Claude Code CLI Architecture and Core Mechanics|01. Claude Code CLI Architecture]] — Interactive terminal agent, filesystem tool execution, bash command spawning, and deterministic code manipulation.
- 📂 [[Claude.md Context Engineering and System Prompts|02. Claude.md Context Engineering]] — Structuring repository guidelines, project build commands, architectural boundaries, and coding standards in CLAUDE.md.
- 📂 [[Custom Skills Architecture (skill.md) and Tool Extensions|03. Custom Skills and Tool Extensions]] — Authoring modular skill packages, YAML frontmatter configuration, reference scripts, and on-demand skill activation.
- 📂 [[Subagent Parallelism and Task Delegation|04. Subagent Delegation and Parallel Execution]] — Spawning isolated subagents, context window isolation, background worker threads, and result synthesis.
- 📂 [[Custom Lifecycle Hooks (hooks.json) and Shell Automations|05. Custom Lifecycle Hooks and Automations]] — Pre-tool, post-tool, and user prompt interception hooks, automated linting on file edit, and security validation.
- 📂 [[Context Window Management and Token Compaction|06. Context Management and Token Compaction]] — Token budgeting, smart file pruning, transcript summarization, and avoiding context saturation in long coding sessions.
- 📂 [[Test Driven Development (tdd) with Claude Code|07. Test Driven Development with Claude Code]] — Authoring failing regression tests, running automated test runners, incremental code fixes, and test validation.
- 📂 [[Large Scale Repository Refactoring with Claude Code|08. Large Scale Refactoring Workflows]] — AST code searching, multi-file atomic replacements, API deprecation migrations, and whole-codebase consistency audits.
- 📂 [[Claude Code Permission Controls and Command Sandboxing|09. Security Sandboxing and Permission Controls]] — Auto-approved commands, bash execution permission boundaries, sensitive file ignore patterns, and API token security.
- 📂 [[Headless Claude Code Execution and Ci-cd Integrations|10. Ci/cd and Headless Automation with Claude]] — Running non-interactive Claude Code tasks in GitHub Actions, automated issue triaging, and automated PR review bots.
- 📂 [[Advanced Claude Code Workflows, Slash Commands, and Custom Tools|11. Advanced Claude Code Workflows and Patterns]] — Interactive interview modes, custom MCP server sidecar integration, memory management, and project onboarding.
### 6. 📂 [[AI Assisted Development & Vibe Coding|06. AI-Assisted Development & Vibe Coding]]
- 📂 [[Vibe Coding Philosophy and Spec Driven Architecture|01. Vibe Coding Philosophy and Spec Driven Design]] — Transitioning from line-by-line syntax writing to high-level system architecture, specification authoring, and verification.
- 📂 [[Planning Before Coding and Technical Design Artifacts|02. Planning Before Coding and Implementation Plans]] — Generating detailed implementation plans, architecture review artifacts, edge-case risk matrices, and user verification checklists.
- 📂 [[Context Optimization in AI IDEs (cursor, Windsurf, Antigravity)|03. Context Optimization in AI Coding IDEs]] — Rules files (.cursorrules, .agents), @-mentions, symbol indexing, semantic code search, and managing open editor context.
- 📂 [[Effective Prompting Patterns for Software Engineers|04. Prompting Patterns for Software Engineering]] — Providing explicit constraints, specifying acceptance criteria, asking for step-by-step diffs, and requesting failure modes.
- 📂 [[AI Driven Debugging, Log Parsing, and Root Cause Analysis|05. AI Driven Debugging and Root Cause Analysis]] — Feeding stack traces, core dumps, and distributed logs to AI models for rapid automated root-cause hypothesis generation.
- 📂 [[Git Mastery for AI Pair Programming and Micro Commits|06. Git and Version Control for AI Pair Programming]] — Atomic commit hygiene, branching per AI task, reviewing AI-generated diffs carefully, and automated rollback workflows.
- 📂 [[Automated Test Generation and Boundary Fuzzing with AI|07. Automated Test Generation with AI]] — Generating comprehensive unit test matrices, property-based tests, mock generation, and edge-case boundary fuzzing.
- 📂 [[Security Verification and Code Quality for AI Generated Code|08. Security Verification for AI Generated Code]] — Auditing AI code for OWASP Top 10 flaws, hardcoded credentials, hallucinated dependencies, and logic omissions.
- 📂 [[Domain Modeling, Schema Design, and Code Refactoring with AI|09. Domain Modeling and Refactoring with AI]] — Translating business domain requirements into clean Go/SQL schemas, domain entities, and refactoring legacy codebases.
- 📂 [[The Future of Principal Engineering in the AI Era|10. Future of Software Engineering with AI]] — The evolving role of Staff/Principal engineers: System-level reasoning, trade-off analysis, security leadership, and AI orchestration.

---

## 🔗 Navigation
- ⬆️ Parent: [[Principal SWE]]
- 🎓 Root: [[Principal SWE]]
