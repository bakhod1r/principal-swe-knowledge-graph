---
title: AI Agents
tags:
  - ai-and-machine-learning
  - ai-engineering
  - ai-agents,-autonomous-systems-and-multi-agent-orchestration
  - principal-swe
parent: "[[AI & Intelligent Systems]]"
---

# 🤖 AI Agents, Autonomous Systems & Multi Agent Orchestration

Autonomous agent engineering: ReAct loop, tool calling, memory architectures, planning algorithms (Reflexion), LangGraph state machines, LlamaIndex workflows, CrewAI & AutoGen multi-agent systems, Human-in-the-Loop, and sandbox execution.

```text
AI Agents, Autonomous Systems & Multi Agent Orchestration
│
├── [[AI Agent Foundations: ReAct Loop (reasoning + Acting) and Agent Workflows|01. AI Agent Architecture and the ReAct Paradigm]]
├── [[LLM Tool Calling, Function Schemas, and Dynamic Parameter Extraction|02. Tool Calling and Function Calling Specifications]]
├── [[Agent Memory Architectures: Short Term, Long Term, and Episodic Memory|03. Agent Memory Systems Short Term, Long Term, Episodic]]
├── [[Agent Planning Algorithms: Plan and Solve, Reflexion, and Self Correction|04. Planning Algorithms Plan and Solve and Reflexion]]
├── `05. Langgraph State Machine Architecture and Workflows`
├── `06. Llamaindex Agent Workflows and Query Engines`
├── `07. Crewai Multi Agent Role Playing Collaboration`
├── `08. Autogen Conversational Patterns and Group Chats`
├── `09. Human in the Loop HITL and Approval Gates`
├── `10. Agent Code Sandboxes and Secure Execution Runtimes`
├── `11. Error Recovery, Loop Detection, and Self Healing Agents`
├── `12. Guardrails, Safety Filters, and Output Governance`
├── `13. Autonomous Web Browsing and Computer Use Agents`
└── [[Agent Observability, Execution Tracing (langfuse, Arize Phoenix), and Evals|14. Agent Observability, Tracing, and Evaluation]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[AI Agent Foundations: ReAct Loop (reasoning + Acting) and Agent Workflows|01. AI Agent Architecture and the ReAct Paradigm]] — The autonomous loop: Observe -> Think -> Act -> Observe; state management, tool execution cycle, and breaking complex objectives into subgoals.
- 📂 [[LLM Tool Calling, Function Schemas, and Dynamic Parameter Extraction|02. Tool Calling and Function Calling Specifications]] — Declaring JSON tool schemas, deterministic tool execution, handling multiple parallel tool calls, and error handling when tool invocations fail.
- 📂 [[Agent Memory Architectures: Short Term, Long Term, and Episodic Memory|03. Agent Memory Systems Short Term, Long Term, Episodic]] — Sliding window conversation buffers, summary buffers, long-term semantic memory in vector stores, and episodic memory retrieval for past successes.
- 📂 [[Agent Planning Algorithms: Plan and Solve, Reflexion, and Self Correction|04. Planning Algorithms Plan and Solve and Reflexion]] — Static planning vs dynamic re-planning, verbal reinforcement learning via Reflexion, analyzing execution trace failures, and iterative self-repair.
- 📂 `05. Langgraph State Machine Architecture and Workflows` — Building cyclical agent graphs in LangGraph, state schemas, conditional edges, checkpointing for state persistence, and time-travel debugging.
- 📂 `06. Llamaindex Agent Workflows and Query Engines` — Router query engines, sub-question query engines, multi-document agent workflows, and combining retrieval tools with execution actions.
- 📂 `07. Crewai Multi Agent Role Playing Collaboration` — Defining agent personas (Role, Goal, Backstory), task dependencies, sequential vs hierarchical process execution, and inter-agent communication.
- 📂 `08. Autogen Conversational Patterns and Group Chats` — Multi-agent conversational patterns, UserProxy agents, group chat round-robin vs LLM-directed speaker selection, and code execution environments.
- 📂 `09. Human in the Loop HITL and Approval Gates` — Pausing agent execution before high-risk actions (e.g. database mutations, sending emails), human feedback injection, and resuming graph state.
- 📂 `10. Agent Code Sandboxes and Secure Execution Runtimes` — Isolating untrusted AI-generated code execution, ephemeral microVM sandboxes (E2B), network egress restrictions, CPU/memory limits, and timeout controls.
- 📂 `11. Error Recovery, Loop Detection, and Self Healing Agents` — Detecting repetitive cycling in tool calls, max-iteration caps, fallback strategies, and prompt re-anchoring when agents get stuck.
- 📂 `12. Guardrails, Safety Filters, and Output Governance` — Restricting allowed actions, parameter range validation, content moderation filters, and preventing unauthorized system access by autonomous agents.
- 📂 `13. Autonomous Web Browsing and Computer Use Agents` — DOM tree parsing for LLMs, accessibility tree navigation, coordinate-based clicking and typing (Anthropic Computer Use), and CAPTCHA handling.
- 📂 [[Agent Observability, Execution Tracing (langfuse, Arize Phoenix), and Evals|14. Agent Observability, Tracing, and Evaluation]] — Tracing nested tool calls and reasoning chains, measuring task success rate, tracking token costs per agent run, and benchmark evaluations (GAIA, SWE-bench).

---

## 🔗 References
- ⬆️ Parent: `AI & Machine Learning`

