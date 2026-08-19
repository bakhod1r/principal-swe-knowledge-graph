---
title: AI Engineering
tags:
  - ai-and-machine-learning
  - ai-engineering
  - ai-engineering,-vector-databases-and-rag-architectures
  - principal-swe
parent: "[[AI & Intelligent Systems]]"
---

# 🤖 AI Engineering, Vector Databases & RAG Architectures

Production AI engineering: Multi-model API routing, dense embeddings, vector indexing (HNSW/IVF), vector databases (Qdrant, Milvus, pgvector), chunking strategies, hybrid search, Reranking, advanced RAG (Self-RAG, GraphRAG), and evaluation (Ragas).

```text
AI Engineering, Vector Databases & RAG Architectures
│
├── [[Foundation Model APIs (openai, Anthropic, Gemini, Deepseek) and Gateways|01. Foundation Model APIs and Multi Provider Gateways]]
├── `02. Dense Embeddings, Similarity Metrics, and Quantization`
├── `03. Vector Database Internals and Indexing Algorithms`
├── `04. Vector Databases in Production Qdrant, Milvus, Pgvector`
├── `05. Document Ingestion and Chunking Strategies`
├── `06. Dense vs Sparse Hybrid Search and Bm25`
├── `07. Cross Encoder Reranking and Contextual Compression`
├── `08. Advanced RAG Architectures Self RAG and Corrective RAG`
├── `09. Graphrag and Knowledge Graph Augmented Retrieval`
├── `10. Multi Modal Retrieval and Document Parsing`
├── `11. RAG Evaluation Frameworks Ragas and Trulens`
├── [[Semantic Caching (gptcache, Redis), Prompt Caching, and Latency Reduction|12. Semantic Caching and LLM Response Optimization]]
└── [[Token Economics, LLM Cost Attribution, and Budget Governance|13. Token Economics, Budgeting, and Cost Governance]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Foundation Model APIs (openai, Anthropic, Gemini, Deepseek) and Gateways|01. Foundation Model APIs and Multi Provider Gateways]] — REST/SDK integrations, streaming server-sent events (SSE), fallback routing (LiteLLM), token usage tracking, and multi-provider load balancing.
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
- 📂 [[Semantic Caching (gptcache, Redis), Prompt Caching, and Latency Reduction|12. Semantic Caching and LLM Response Optimization]] — Caching LLM responses by semantic embedding similarity threshold, exact-match prompt prefix caching, and slashing API latency and dollar costs.
- 📂 [[Token Economics, LLM Cost Attribution, and Budget Governance|13. Token Economics, Budgeting, and Cost Governance]] — Tracking input/output token costs across teams, model routing based on query complexity (small fast model -> heavy frontier model), and budgeting guardrails.

---

## 🔗 References
- ⬆️ Parent: `AI & Machine Learning`

