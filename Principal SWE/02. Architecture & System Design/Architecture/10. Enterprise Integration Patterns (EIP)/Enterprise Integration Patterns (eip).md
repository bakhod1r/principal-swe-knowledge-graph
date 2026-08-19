---
title: Enterprise Integration Patterns (eip)
tags:
  - review
  - architecture
  - systems-architecture
  - enterprise-integration-patterns-(eip)
  - principal-swe
parent: "[[Architecture]]"
---

# 🏛️ Enterprise Integration Patterns (eip)

Gregor Hohpe Enterprise Integration Patterns: Message Channels, Pipes & Filters, Message Routers, Message Translators, Content-Based Routers, Scatter-Gather, Recipient Lists, Resequencers, and Claim Check patterns.

```text
Enterprise Integration Patterns (eip)
│
├── [[Message Channels: Point to Point vs Publish Subscribe Channels|01. Message Channels and Point to Point Topologies]]
├── [[Pipes and Filters Architecture: Composable Data Processing Pipelines|02. Pipes and Filters Architecture Pattern]]
├── [[Message Routing Patterns: Content Based Router and Message Filter|03. Message Routing Content Based and Filter Routers]]
├── [[Message Transformation: Message Translator, Envelope Wrapper, and Normalizer|04. Message Transformation, Translator, and Normalizer]]
├── [[Scatter Gather Pattern: Parallel Broadcast and Aggregated Response|05. Scatter Gather Pattern and Parallel Aggregation]]
├── [[Recipient List and Routing Slip: Dynamic Multi Destination Messaging|06. Recipient List and Dynamic Routing Slips]]
├── [[Message Sequencing: Resequencer, Aggregator, and Correlation Identifiers|07. Resequencer, Aggregator, and Message Correlator]]
├── [[Message Broker Topologies, Messaging Bridges, and Bus Interconnects|08. Message Broker Topologies and Enterprise Bridges]]
├── [[The Claim Check Pattern: Storing Large Payloads in Blob Storage|09. Claim Check Pattern for Large Payload Messaging]]
└── [[Idempotent Message Receiver and Enterprise Process Manager|10. Idempotent Message Receiver and Process Manager]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Message Channels: Point to Point vs Publish Subscribe Channels|01. Message Channels and Point to Point Topologies]] — Establishing decoupled communication channels between heterogeneous enterprise applications, channel adapters, and message endpoints.
- 📂 [[Pipes and Filters Architecture: Composable Data Processing Pipelines|02. Pipes and Filters Architecture Pattern]] — Building modular data processing pipelines where independent filter components process and pass messages through unidirectional pipes.
- 📂 [[Message Routing Patterns: Content Based Router and Message Filter|03. Message Routing Content Based and Filter Routers]] — Inspecting message headers and payload contents to dynamically route messages to specific destinations without sender awareness.
- 📂 [[Message Transformation: Message Translator, Envelope Wrapper, and Normalizer|04. Message Transformation, Translator, and Normalizer]] — Translating disparate enterprise data formats (XML, JSON, Protobuf, CSV), wrapping payloads in canonical message envelopes, and normalizers.
- 📂 [[Scatter Gather Pattern: Parallel Broadcast and Aggregated Response|05. Scatter Gather Pattern and Parallel Aggregation]] — Broadcasting a request message to multiple vendor/service endpoints in parallel, aggregating responses, and selecting the best offer with timeouts.
- 📂 [[Recipient List and Routing Slip: Dynamic Multi Destination Messaging|06. Recipient List and Dynamic Routing Slips]] — Calculating a dynamic list of recipients based on runtime business rules, and attaching a sequential itinerary (Routing Slip) to the message.
- 📂 [[Message Sequencing: Resequencer, Aggregator, and Correlation Identifiers|07. Resequencer, Aggregator, and Message Correlator]] — Reordering out-of-order event streams based on sequence numbers, and correlating related messages with unique `Correlation-ID` headers.
- 📂 [[Message Broker Topologies, Messaging Bridges, and Bus Interconnects|08. Message Broker Topologies and Enterprise Bridges]] — Bridging incompatible message brokers (Kafka to RabbitMQ, SQS to GCP PubSub), message routing hubs, and store-and-forward persistence.
- 📂 [[The Claim Check Pattern: Storing Large Payloads in Blob Storage|09. Claim Check Pattern for Large Payload Messaging]] — Splitting massive message payloads into external cloud blob storage (S3) and passing only a lightweight reference token (claim check) over the message bus.
- 📂 [[Idempotent Message Receiver and Enterprise Process Manager|10. Idempotent Message Receiver and Process Manager]] — Ensuring exactly-once business processing semantics on at-least-once message brokers, and orchestrating complex stateful business processes.

---

## 🔗 References
- ⬆️ Parent: [[Architecture]]

