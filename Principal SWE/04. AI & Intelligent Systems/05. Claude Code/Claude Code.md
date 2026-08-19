---
title: Claude Code
tags:
  - ai-and-machine-learning
  - ai-engineering
  - claude-code,-subagents,-model-context-protocol-mcp-and-tooling
  - principal-swe
parent: "[[AI & Intelligent Systems]]"
---

# 🤖 Claude Code, Subagents, Model Context Protocol MCP & Tooling

Anthropic Claude Code and agentic tooling: Model Context Protocol (MCP) servers and clients, subagent delegation, multi-file code synthesis, tool execution safety, conversation history compaction, slash commands, skills/rules customization, and production codebase refactoring.

```text
Claude Code, Subagents, Model Context Protocol MCP & Tooling
│
├── [[Claude Code Architecture, Terminal Integration, and Tool Invocation Engine|01. Claude Code Architecture and Terminal Execution Engine]]
├── `02. Model Context Protocol MCP Standards and Architecture`
├── `03. Developing Custom MCP Servers in Typescript and Python`
├── [[Subagent Spawning, Delegation Protocols, and Clean Context Handoffs|04. Subagent Delegation, Swarms, and Context Handoffs]]
├── [[Multi File Code Synthesis, Ast Refactoring, and Atomic Patch Applications|05. Multi File Code Synthesis and Atomic Patching]]
├── `06. Terminal and Tool Execution Safety and Sandboxing`
├── [[Conversation History Compaction, Semantic Summarization, and Transcripts|07. Conversation History Compaction and Transcript Management]]
├── [[Custom Slash Commands, Task Scheduling, and Background Daemons|08. Custom Slash Commands and Background Execution]]
├── [[Agent Skill and Rule Customization Engine (yaml Frontmatter, Directory Scopes)|09. Skill and Rule Customization Engine and Discovery]]
├── [[Automated Pull Request Generation and Headless CI Runs with Claude Code|10. Automated PR Generation and CI Integration with Claude Code]]
├── [[Large Scale Production Codebase Refactoring and Migration with AI Agents|11. Production Codebase Refactoring with Autonomous Agents]]
├── [[Debugging Agent Failure Modes, Tool Call Loops, and Hallucinated APIs|12. Debugging Autonomous Agent Failure Modes and Hallucinations]]
└── [[Evaluating Coding Agents: Swe Bench, Humaneval, and Real World Bug Resolution|13. Benchmark Evaluations for Coding Agents Swe Bench]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Claude Code Architecture, Terminal Integration, and Tool Invocation Engine|01. Claude Code Architecture and Terminal Execution Engine]] — Autonomous terminal agent design, real-time command execution, file system read/write primitives, and streaming token response handling.
- 📂 `02. Model Context Protocol MCP Standards and Architecture` — The MCP client-server standard, JSON-RPC 2.0 transport (stdio, SSE), Resources, Tools, and Prompts primitives, and ecosystem integration.
- 📂 `03. Developing Custom MCP Servers in Typescript and Python` — Exposing enterprise internal APIs, databases, and internal tools as MCP endpoints; parameter schema validation, and error serialization.
- 📂 [[Subagent Spawning, Delegation Protocols, and Clean Context Handoffs|04. Subagent Delegation, Swarms, and Context Handoffs]] — Spawning specialized subagents for isolated tasks, passing constrained context packets, aggregating subagent results, and preventing context window bloat.
- 📂 [[Multi File Code Synthesis, Ast Refactoring, and Atomic Patch Applications|05. Multi File Code Synthesis and Atomic Patching]] — Generating consistent multi-file diffs, applying targeted replacements without corrupting whitespace, syntax tree validation, and rolling back failed edits.
- 📂 `06. Terminal and Tool Execution Safety and Sandboxing` — Detecting dangerous shell commands (`rm -rf`, `DROP TABLE`), requiring explicit confirmation, read-only modes, and sandbox containment.
- 📂 [[Conversation History Compaction, Semantic Summarization, and Transcripts|07. Conversation History Compaction and Transcript Management]] — Pruning large terminal tool outputs, sliding window transcript compression, maintaining critical state invariants across long multi-turn sessions.
- 📂 [[Custom Slash Commands, Task Scheduling, and Background Daemons|08. Custom Slash Commands and Background Execution]] — Extending agent CLI capabilities with custom slash commands (`/goal`, `/schedule`), background async tasks, and reactive wakeups.
- 📂 [[Agent Skill and Rule Customization Engine (yaml Frontmatter, Directory Scopes)|09. Skill and Rule Customization Engine and Discovery]] — Modular skill definitions (`SKILL.md`), hierarchical rule loading (global vs workspace), plugin architectures, and dynamic on-demand capability injection.
- 📂 [[Automated Pull Request Generation and Headless CI Runs with Claude Code|10. Automated PR Generation and CI Integration with Claude Code]] — Running Claude Code in non-interactive CI/CD pipelines, automated bug fixing from stack traces, generating clean PR diffs with testing evidence.
- 📂 [[Large Scale Production Codebase Refactoring and Migration with AI Agents|11. Production Codebase Refactoring with Autonomous Agents]] — Orchestrating multi-day codebase upgrades (e.g. library migrations, TypeScript strict mode conversion), test-driven verification, and tracking progress.
- 📂 [[Debugging Agent Failure Modes, Tool Call Loops, and Hallucinated APIs|12. Debugging Autonomous Agent Failure Modes and Hallucinations]] — Analyzing agent decision transcripts, diagnosing incorrect tool argument generation, debugging prompt injection leaks, and calibrating agent prompts.
- 📂 [[Evaluating Coding Agents: Swe Bench, Humaneval, and Real World Bug Resolution|13. Benchmark Evaluations for Coding Agents Swe Bench]] — Setting up SWE-bench evaluation harnesses, measuring pass@1 and pass@5 resolution rates on real GitHub issues, and regression testing agent capabilities.

---

## 🔗 References
- ⬆️ Parent: `AI & Machine Learning`

