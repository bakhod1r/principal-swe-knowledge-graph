---
title: AI Engineering & RAG Systems
tags:
  - ai
  - machine-learning
  - ai-engineering-and-rag-systems
  - principal-swe
parent: "[[AI & Machine Learning]]"
---

# 🏛️ AI Engineering & RAG Systems

Production Retrieval-Augmented Generation (RAG) and embedding architectures: Foundation model APIs, local model serving (vLLM, Ollama), vector databases (HNSW indexing), chunking strategies, hybrid search, reranking, and multimodal LLMs.

```text
AI Engineering & RAG Systems
│
├── [[Foundation Model Ecosystem and Cloud APIs|01. Foundation Models and Cloud APIs]]
├── [[Open Source Models and Local High Throughput Serving|02. Open Source Models and Local Serving]]
├── [[Embedding Models and High Dimensional Vector Spaces|03. Embeddings and Vector Search Spaces]]
├── [[Vector Databases and HNSW Index Architectures|04. Vector Databases and HNSW Indexing]]
├── [[Document Parsing, Chunking, and Metadata Enrichment|05. Document Parsing and Chunking Strategies]]
├── [[Advanced Retrieval, Hybrid Search, and Re Ranking|06. Advanced Retrieval and Hybrid Search]]
├── [[Context Compression, Lost in the Middle, and Caching|07. Context Compression and Context Scaling]]
├── [[Multimodal AI Architectures (vision, Audio, Video)|08. Multimodal AI Architectures]]
├── [[AI Code Generation and Developer Assistant Tooling|09. AI Code Generation and Assistant Tooling]]
├── [[RAG Evaluation Frameworks (ragas, Trulens)|10. RAG Evaluation Frameworks]]
└── [[Production RAG Architecture and Semantic Caching|11. Production RAG Architecture and Caching]]
```

---

## 🗂️ Core Knowledge Domains

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

---

## 🔗 References
- ⬆️ Parent: [[AI & Machine Learning]]

