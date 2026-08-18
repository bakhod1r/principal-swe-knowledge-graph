---
title: System Design
tags:
  - system-design
  - architecture
  - distributed-systems
  - principal-swe
parent: "[[Principal SWE]]"
---

# 🏗️ System Design & Distributed Architecture

Comprehensive, production-proven blueprints for designing, scaling, and operating mission-critical distributed systems. Covering trade-off frameworks (CAP/PACELC), global multi-region topologies, data streaming, storage engines, reliability patterns, zero-trust security, fintech transaction pipelines, and real-world Big Tech architectures across 43 specialized pillars and 376 subdomains.

```text
System Design
│
├── [[Introduction|01. Introduction]]
│   ├── [[What Is System Design|01. What Is System Design]]
│   ├── [[How to Approach|02. How to Approach]]
│   ├── [[Functional vs Nonfunctional|03. Functional vs Nonfunctional]]
│   ├── [[Key Characteristics|04. Key Characteristics]]
│   └── [[Numbers Every Engineer Should Know|05. Numbers Every Engineer Should Know]]
├── [[Tradeoffs Framework|02. Tradeoffs Framework]]
│   ├── [[CAP Theorem (Tradeoffs Framework)|01. CAP Theorem]]
│   ├── [[PACELC|02. PACELC]]
│   ├── [[Consistency vs Availability|03. Consistency vs Availability]]
│   └── [[Consistency Models (Tradeoffs Framework)|04. Consistency Models]]
├── [[Capacity Estimation|03. Capacity Estimation]]
│   ├── [[QPS|01. QPS]]
│   ├── [[Storage|02. Storage]]
│   ├── [[Bandwidth|03. Bandwidth]]
│   └── [[Latency Budgets|04. Latency Budgets]]
├── [[Back of Envelope|04. Back of Envelope]]
│   ├── [[Number Tables|01. Number Tables]]
│   └── [[Fermi Estimation|02. Fermi Estimation]]
├── [[Networking Protocols|05. Networking Protocols]]
│   ├── [[Osi and TCP Ip|01. Osi and TCP Ip]]
│   ├── [[TCP vs UDP|02. TCP vs UDP]]
│   ├── [[TLS and Https|03. TLS and Https]]
│   ├── [[HTTP Evolution 1 2 3 QUIC|04. HTTP Evolution 1 2 3 QUIC]]
│   ├── [[Websockets|05. Websockets]]
│   ├── [[Server Sent Events|06. Server Sent Events]]
│   ├── [[Long Polling and Streaming|07. Long Polling and Streaming]]
│   ├── [[Network Proxies and NAT|08. Network Proxies and NAT]]
│   ├── [[Congestion Control and TCP Tuning|09. Congestion Control and TCP Tuning]]
│   ├── [[Container and Overlay Networking|10. Container and Overlay Networking]]
│   └── [[BGP and Internet Routing|11. BGP and Internet Routing]]
├── [[Domain Name System|06. Domain Name System]]
│   ├── [[DNS Resolution Flow|01. DNS Resolution Flow]]
│   ├── [[Record Types|02. Record Types]]
│   ├── [[DNS Load Balancing|03. DNS Load Balancing]]
│   ├── [[DNS Caching and Ttl|04. DNS Caching and Ttl]]
│   └── [[Geodns and Anycast|05. Geodns and Anycast]]
├── [[Content Delivery Networks|07. Content Delivery Networks]]
│   ├── [[Pull CDN|01. Pull CDN]]
│   ├── [[Push CDN|02. Push CDN]]
│   ├── [[Cache Invalidation|03. Cache Invalidation]]
│   ├── [[Edge Locations|04. Edge Locations]]
│   └── [[CDN Security|05. CDN Security]]
├── [[Load Balancers|08. Load Balancers]]
│   ├── [[Lb vs Reverse Proxy|01. Lb vs Reverse Proxy]]
│   ├── [[Load Balancing Algorithms|02. Load Balancing Algorithms]]
│   ├── [[Layer 4 Load Balancing|03. Layer 4 Load Balancing]]
│   ├── [[Layer 7 Load Balancing|04. Layer 7 Load Balancing]]
│   ├── [[Health Checks and Failover|05. Health Checks and Failover]]
│   ├── [[Horizontal Scaling|06. Horizontal Scaling]]
│   └── [[Global Server Load Balancing|07. Global Server Load Balancing]]
├── [[Communication|09. Communication]]
│   ├── [[HTTP (Communication)|01. HTTP]]
│   ├── [[TCP (Communication)|02. TCP]]
│   ├── [[UDP (Communication)|03. UDP]]
│   ├── [[RPC|04. RPC]]
│   ├── [[GRPC|05. GRPC]]
│   ├── [[REST|06. REST]]
│   ├── [[Graphql|07. Graphql]]
│   └── [[Idempotent Operations|08. Idempotent Operations]]
├── [[Application Layer|10. Application Layer]]
│   ├── [[Microservices|01. Microservices]]
│   ├── [[Monolith vs Microservices|02. Monolith vs Microservices]]
│   ├── [[Service Discovery|03. Service Discovery]]
│   ├── [[API Composition|04. API Composition]]
│   ├── [[Stateless Design|05. Stateless Design]]
│   ├── [[Service Mesh Intro|06. Service Mesh Intro]]
│   ├── [[Serverless Faas|07. Serverless Faas]]
│   └── [[Peer to Peer Architecture|08. Peer to Peer Architecture]]
├── [[API Design at Scale|11. API Design at Scale]]
│   ├── [[API Gateway|01. API Gateway]]
│   ├── [[REST Design at Scale|02. REST Design at Scale]]
│   ├── [[Graphql Federation|03. Graphql Federation]]
│   ├── [[GRPC and Streaming|04. GRPC and Streaming]]
│   ├── [[Versioning and Deprecation|05. Versioning and Deprecation]]
│   ├── [[Pagination and Filtering|06. Pagination and Filtering]]
│   ├── [[Idempotency and Retries|07. Idempotency and Retries]]
│   ├── [[Webhooks|08. Webhooks]]
│   └── [[Backends for Frontend|09. Backends for Frontend]]
├── [[Databases (System Design)|12. Databases]]
│   ├── [[Relational Rdbms|01. Relational Rdbms]]
│   ├── [[Key Value|02. Key Value]]
│   ├── [[Document|03. Document]]
│   ├── [[Wide Column|04. Wide Column]]
│   ├── [[Column Oriented Olap|05. Column Oriented Olap]]
│   ├── [[Graph|06. Graph]]
│   ├── [[Time Series|07. Time Series]]
│   ├── [[Search Engine|08. Search Engine]]
│   ├── [[Vector|09. Vector]]
│   ├── [[Newsql Distributed SQL|10. Newsql Distributed SQL]]
│   ├── [[Replication (Databases)|11. Replication]]
│   ├── [[Sharding and Partitioning|12. Sharding and Partitioning]]
│   ├── [[Indexing|13. Indexing]]
│   ├── [[Transactions and Isolation|14. Transactions and Isolation]]
│   ├── [[Denormalization|15. Denormalization]]
│   ├── [[SQL Tuning|16. SQL Tuning]]
│   ├── [[SQL vs NoSQL|17. SQL vs NoSQL]]
│   ├── [[Oltp vs Olap|18. Oltp vs Olap]]
│   ├── [[Polyglot Persistence|19. Polyglot Persistence]]
│   ├── [[Choosing a Database|20. Choosing a Database]]
│   ├── [[Connection Pooling Pgbouncer Proxysql|21. Connection Pooling Pgbouncer Proxysql]]
│   ├── [[Database HA and Failover|22. Database HA and Failover]]
│   └── [[Backup Restore and Pitr|23. Backup Restore and Pitr]]
├── [[Storage Systems|13. Storage Systems]]
│   ├── [[Object vs Block vs File|01. Object vs Block vs File]]
│   ├── [[Distributed File Systems Gfs Hdfs|02. Distributed File Systems Gfs Hdfs]]
│   ├── [[Blob Storage S3|03. Blob Storage S3]]
│   ├── [[LSM Trees and Btrees|04. LSM Trees and Btrees]]
│   ├── [[Data Warehouse vs Lake|05. Data Warehouse vs Lake]]
│   ├── [[File Formats Parquet Orc Iceberg|06. File Formats Parquet Orc Iceberg]]
│   └── [[Erasure Coding and Durability|07. Erasure Coding and Durability]]
├── [[Caching|14. Caching]]
│   ├── [[Cache Aside|01. Cache Aside]]
│   ├── [[Write Through|02. Write Through]]
│   ├── [[Write Behind|03. Write Behind]]
│   ├── [[Refresh Ahead|04. Refresh Ahead]]
│   ├── [[Eviction Policies|05. Eviction Policies]]
│   ├── [[Types of Caching|06. Types of Caching]]
│   ├── [[Cache Invalidation (Caching)|07. Cache Invalidation]]
│   └── [[Cache Stampede and Hot Keys|08. Cache Stampede and Hot Keys]]
├── [[Data Streaming|15. Data Streaming]]
│   ├── [[Batch Processing Mapreduce|01. Batch Processing Mapreduce]]
│   ├── [[Apache Spark|02. Apache Spark]]
│   ├── [[Stream Processing|03. Stream Processing]]
│   ├── [[Apache Kafka|04. Apache Kafka]]
│   ├── [[Lambda vs Kappa Architecture|05. Lambda vs Kappa Architecture]]
│   ├── [[Data Lake and Warehouse|06. Data Lake and Warehouse]]
│   ├── [[Change Data Capture|07. Change Data Capture]]
│   ├── [[Etl vs Elt|08. Etl vs Elt]]
│   ├── [[Apache Flink|09. Apache Flink]]
│   ├── [[Apache Pulsar|10. Apache Pulsar]]
│   ├── [[Trino Presto Distributed Query|11. Trino Presto Distributed Query]]
│   ├── [[Workflow Orchestration Airflow Dagster|12. Workflow Orchestration Airflow Dagster]]
│   ├── [[Lakehouse and Table Formats|13. Lakehouse and Table Formats]]
│   ├── [[Dimensional Data Modeling|14. Dimensional Data Modeling]]
│   └── [[Reverse Etl and Operational Analytics|15. Reverse Etl and Operational Analytics]]
├── [[Asynchronism|16. Asynchronism]]
│   ├── [[Message Queues|01. Message Queues]]
│   ├── [[Task Queues|02. Task Queues]]
│   ├── [[Back Pressure|03. Back Pressure]]
│   ├── [[Dead Letter Queues|04. Dead Letter Queues]]
│   ├── [[Delivery Guarantees|05. Delivery Guarantees]]
│   ├── [[Rabbitmq|06. Rabbitmq]]
│   └── [[Nats|07. Nats]]
├── [[Background Jobs|17. Background Jobs]]
│   ├── [[Event Driven|01. Event Driven]]
│   ├── [[Schedule Driven|02. Schedule Driven]]
│   ├── [[Returning Results|03. Returning Results]]
│   ├── [[Retries and Idempotency|04. Retries and Idempotency]]
│   └── [[Durable Execution Temporal|05. Durable Execution Temporal]]
├── [[Concurrency Coordination|18. Concurrency Coordination]]
│   ├── [[Idempotency Keys|01. Idempotency Keys]]
│   ├── [[Leases and Fencing|02. Leases and Fencing]]
│   ├── [[Exactly Once Semantics|03. Exactly Once Semantics]]
│   ├── [[Optimistic vs Pessimistic Locking|04. Optimistic vs Pessimistic Locking]]
│   ├── [[Coordination Services|05. Coordination Services]]
│   ├── [[Gossip Protocol|06. Gossip Protocol]]
│   └── [[Atomic Commit 2PC 3PC TCC|07. Atomic Commit 2PC 3PC TCC]]
├── [[Building Blocks|19. Building Blocks]]
│   ├── [[Rate Limiter|01. Rate Limiter]]
│   ├── [[Consistent Hashing|02. Consistent Hashing]]
│   ├── [[Unique Id Generator|03. Unique Id Generator]]
│   ├── [[Distributed Key Value Store|04. Distributed Key Value Store]]
│   ├── [[Distributed Cache|05. Distributed Cache]]
│   ├── [[Distributed Message Queue|06. Distributed Message Queue]]
│   ├── [[Pub Sub System|07. Pub Sub System]]
│   ├── [[Blob Object Store|08. Blob Object Store]]
│   ├── [[Distributed Search Typeahead|09. Distributed Search Typeahead]]
│   ├── [[Distributed Task Scheduler|10. Distributed Task Scheduler]]
│   ├── [[Distributed Lock|11. Distributed Lock]]
│   ├── [[Distributed Logging|12. Distributed Logging]]
│   ├── [[Sharded Counters Leaderboard|13. Sharded Counters Leaderboard]]
│   └── [[Bloom Filter (Building Blocks)|14. Bloom Filter]]
├── [[Reliability Patterns|20. Reliability Patterns]]
│   ├── [[Circuit Breaker|01. Circuit Breaker]]
│   ├── [[Bulkhead|02. Bulkhead]]
│   ├── [[Retry|03. Retry]]
│   ├── [[Throttling|04. Throttling]]
│   ├── [[Health Endpoint Monitoring|05. Health Endpoint Monitoring]]
│   ├── [[Leader Election|06. Leader Election]]
│   ├── [[Compensating Transaction|07. Compensating Transaction]]
│   ├── [[Deployment Stamps and Geodes|08. Deployment Stamps and Geodes]]
│   ├── [[Queue Based Load Leveling|09. Queue Based Load Leveling]]
│   ├── [[Redundancy and Failure Domains|10. Redundancy and Failure Domains]]
│   └── [[Shuffle Sharding|11. Shuffle Sharding]]
├── [[Cloud Design Patterns|21. Cloud Design Patterns]]
│   ├── [[Strangler Fig|01. Strangler Fig]]
│   ├── [[Sidecar|02. Sidecar]]
│   ├── [[Ambassador|03. Ambassador]]
│   ├── [[Anti Corruption Layer|04. Anti Corruption Layer]]
│   ├── [[CQRS|05. CQRS]]
│   ├── [[Event Sourcing|06. Event Sourcing]]
│   ├── [[Materialized View|07. Materialized View]]
│   ├── [[Pipes and Filters|08. Pipes and Filters]]
│   ├── [[External Config Store|09. External Config Store]]
│   ├── [[Valet Key|10. Valet Key]]
│   ├── [[Claim Check|11. Claim Check]]
│   ├── [[Competing Consumers|12. Competing Consumers]]
│   ├── [[Publisher Subscriber|13. Publisher Subscriber]]
│   └── [[Transactional Outbox|14. Transactional Outbox]]
├── [[Performance Antipatterns|22. Performance Antipatterns]]
│   ├── [[Busy Database|01. Busy Database]]
│   ├── [[Busy Frontend|02. Busy Frontend]]
│   ├── [[Chatty IO|03. Chatty IO]]
│   ├── [[Extraneous Fetching|04. Extraneous Fetching]]
│   ├── [[Improper Instantiation|05. Improper Instantiation]]
│   ├── [[Monolithic Persistence|06. Monolithic Persistence]]
│   ├── [[Noisy Neighbor|07. Noisy Neighbor]]
│   ├── [[Synchronous IO|08. Synchronous IO]]
│   ├── [[Retry Storm|09. Retry Storm]]
│   └── [[No Caching|10. No Caching]]
├── [[Monitoring|23. Monitoring]]
│   ├── [[Health Monitoring|01. Health Monitoring]]
│   ├── [[Availability Monitoring|02. Availability Monitoring]]
│   ├── [[Performance Monitoring|03. Performance Monitoring]]
│   ├── [[Security Monitoring|04. Security Monitoring]]
│   ├── [[Usage Monitoring|05. Usage Monitoring]]
│   ├── [[Instrumentation|06. Instrumentation]]
│   ├── [[Visualization and Alerts|07. Visualization and Alerts]]
│   └── [[Synthetic Monitoring and Rum|08. Synthetic Monitoring and Rum]]
├── [[Observability|24. Observability]]
│   ├── [[Logs Metrics Traces|01. Logs Metrics Traces]]
│   ├── [[SLO SLI Error Budgets|02. SLO SLI Error Budgets]]
│   ├── [[Red and Use Methods|03. Red and Use Methods]]
│   ├── [[Distributed Tracing|04. Distributed Tracing]]
│   ├── [[Metrics Pipelines|05. Metrics Pipelines]]
│   ├── [[Log Aggregation|06. Log Aggregation]]
│   ├── [[Alerting and on Call|07. Alerting and on Call]]
│   ├── [[Opentelemetry|08. Opentelemetry]]
│   ├── [[Cardinality and Metrics Cost|09. Cardinality and Metrics Cost]]
│   └── [[Trace and Log Sampling Strategies|10. Trace and Log Sampling Strategies]]
├── [[Chaos Engineering|25. Chaos Engineering]]
│   ├── [[Failure Modes|01. Failure Modes]]
│   ├── [[Fault Injection|02. Fault Injection]]
│   ├── [[Game Days|03. Game Days]]
│   ├── [[Resilience Testing|04. Resilience Testing]]
│   └── [[Blast Radius and Recovery|05. Blast Radius and Recovery]]
├── [[Deployment Infrastructure|26. Deployment Infrastructure]]
│   ├── [[Containers and Docker|01. Containers and Docker]]
│   ├── [[Kubernetes Orchestration|02. Kubernetes Orchestration]]
│   ├── [[Deployment Strategies|03. Deployment Strategies]]
│   ├── [[Ci Cd Pipelines|04. Ci Cd Pipelines]]
│   ├── [[Infrastructure As Code|05. Infrastructure As Code]]
│   ├── [[Multi Region Deployment|06. Multi Region Deployment]]
│   ├── [[Disaster Recovery|07. Disaster Recovery]]
│   ├── [[Autoscaling|08. Autoscaling]]
│   ├── [[Cloud Network Architecture Vpc|09. Cloud Network Architecture Vpc]]
│   └── [[Gitops Argocd Flux|10. Gitops Argocd Flux]]
├── [[Security at Scale|27. Security at Scale]]
│   ├── [[Authentication|01. Authentication]]
│   ├── [[Authorization|02. Authorization]]
│   ├── [[Oauth2 and Oidc|03. Oauth2 and Oidc]]
│   ├── [[JWT and Tokens|04. JWT and Tokens]]
│   ├── [[Encryption at REST and Transit|05. Encryption at REST and Transit]]
│   ├── [[Secrets Management|06. Secrets Management]]
│   ├── [[DDoS Mitigation|07. DDoS Mitigation]]
│   ├── [[WAF and API Security|08. WAF and API Security]]
│   ├── [[Rate Limiting for Abuse|09. Rate Limiting for Abuse]]
│   ├── [[Devsecops and Supply Chain Security|10. Devsecops and Supply Chain Security]]
│   ├── [[Zero Trust Architecture|11. Zero Trust Architecture]]
│   ├── [[PKI and Certificate Management|12. PKI and Certificate Management]]
│   ├── [[Threat Modeling Stride|13. Threat Modeling Stride]]
│   └── [[Envelope Encryption and KMS|14. Envelope Encryption and KMS]]
├── [[Data Privacy Compliance|28. Data Privacy Compliance]]
│   ├── [[PII and Data Classification|01. PII and Data Classification]]
│   ├── [[GDPR and Right to Be Forgotten|02. GDPR and Right to Be Forgotten]]
│   ├── [[Data Residency|03. Data Residency]]
│   ├── [[Audit Logging (Data Privacy Compliance)|04. Audit Logging]]
│   └── [[Encryption Key Lifecycle|05. Encryption Key Lifecycle]]
├── [[Multi Tenancy SaaS|29. Multi Tenancy SaaS]]
│   ├── [[Tenant Isolation Models|01. Tenant Isolation Models]]
│   ├── [[Data Partitioning Per Tenant|02. Data Partitioning Per Tenant]]
│   ├── [[Noisy Neighbor Mitigation|03. Noisy Neighbor Mitigation]]
│   ├── [[Per Tenant Scaling and Limits|04. Per Tenant Scaling and Limits]]
│   └── [[Tenant Onboarding and Config|05. Tenant Onboarding and Config]]
├── [[Geospatial Systems|30. Geospatial Systems]]
│   ├── [[Geohashing|01. Geohashing]]
│   ├── [[Quadtrees|02. Quadtrees]]
│   ├── [[S2 and H3|03. S2 and H3]]
│   ├── [[Proximity Search|04. Proximity Search]]
│   └── [[Map Tiling and Routing|05. Map Tiling and Routing]]
├── [[ML Recommendation Systems|31. ML Recommendation Systems]]
│   ├── [[Recommendation Architecture|01. Recommendation Architecture]]
│   ├── [[Feature Store|02. Feature Store]]
│   ├── [[Candidate Generation|03. Candidate Generation]]
│   ├── [[Ranking and Scoring|04. Ranking and Scoring]]
│   ├── [[Online vs Offline Inference|05. Online vs Offline Inference]]
│   ├── [[Ab Testing and Feedback Loops|06. Ab Testing and Feedback Loops]]
│   ├── [[LLM Inference and Serving|07. LLM Inference and Serving]]
│   ├── [[Retrieval Augmented Generation Rag|08. Retrieval Augmented Generation Rag]]
│   ├── [[LLM Application Architecture|09. LLM Application Architecture]]
│   └── [[Ai Agents and Orchestration|10. Ai Agents and Orchestration]]
├── [[Classic Problems|32. Classic Problems]]
│   ├── [[Url Shortener|01. Url Shortener]]
│   ├── [[News Feed and Timeline|02. News Feed and Timeline]]
│   ├── [[Chat and Messaging System|03. Chat and Messaging System]]
│   ├── [[Youtube Video Platform|04. Youtube Video Platform]]
│   ├── [[Google Drive File Sync|05. Google Drive File Sync]]
│   ├── [[Payment System|06. Payment System]]
│   ├── [[Digital Wallet|07. Digital Wallet]]
│   ├── [[Hotel Reservation System|08. Hotel Reservation System]]
│   ├── [[Distributed Email Service|09. Distributed Email Service]]
│   ├── [[Instagram Photo Feed|10. Instagram Photo Feed]]
│   ├── [[Stack Overflow|11. Stack Overflow]]
│   ├── [[Google Docs Collab Editor|12. Google Docs Collab Editor]]
│   ├── [[Proximity Service Maps|13. Proximity Service Maps]]
│   ├── [[Ticketmaster Booking|14. Ticketmaster Booking]]
│   ├── [[Notification System|15. Notification System]]
│   ├── [[Live Streaming|16. Live Streaming]]
│   ├── [[Distributed Job Scheduler|17. Distributed Job Scheduler]]
│   ├── [[Stock Exchange|18. Stock Exchange]]
│   ├── [[S3 Object Storage|19. S3 Object Storage]]
│   ├── [[Online Judge|20. Online Judge]]
│   ├── [[Distributed Analytics Counter|21. Distributed Analytics Counter]]
│   └── [[Web Crawler and Search Engine|22. Web Crawler and Search Engine]]
├── [[Real Architectures|33. Real Architectures]]
│   ├── [[Google Spanner|01. Google Spanner]]
│   ├── [[Facebook Tao|02. Facebook Tao]]
│   ├── [[Amazon Dynamodb|03. Amazon Dynamodb]]
│   ├── [[Netflix Stack|04. Netflix Stack]]
│   ├── [[Apache Kafka (Real Architectures)|05. Apache Kafka]]
│   ├── [[Apache Cassandra|06. Apache Cassandra]]
│   ├── [[Redis Internals|07. Redis Internals]]
│   ├── [[Discord Realtime|08. Discord Realtime]]
│   ├── [[Slack Messaging|09. Slack Messaging]]
│   └── [[Uber Lyft Dispatch|10. Uber Lyft Dispatch]]
├── [[Interview Playbook|34. Interview Playbook]]
│   ├── [[Reshaded Framework|01. Reshaded Framework]]
│   ├── [[Requirements Clarification|02. Requirements Clarification]]
│   ├── [[Capacity Estimation in Interview|03. Capacity Estimation in Interview]]
│   ├── [[API Design Step|04. API Design Step]]
│   ├── [[High Level Design|05. High Level Design]]
│   ├── [[Data Model and Storage|06. Data Model and Storage]]
│   ├── [[Deep Dives and Bottlenecks|07. Deep Dives and Bottlenecks]]
│   ├── [[Tradeoffs and Wrap Up|08. Tradeoffs and Wrap Up]]
│   ├── [[Common Mistakes|09. Common Mistakes]]
│   └── [[Mock Interview Walkthroughs|10. Mock Interview Walkthroughs]]
├── [[Architecture Decision Making|35. Architecture Decision Making]]
│   ├── [[Architecture Decision Records|01. Architecture Decision Records]]
│   ├── [[Rfc Process|02. Rfc Process]]
│   ├── [[Evolutionary Architecture|03. Evolutionary Architecture]]
│   ├── [[Fitness Functions|04. Fitness Functions]]
│   ├── [[Tech Radar|05. Tech Radar]]
│   ├── [[Build vs Buy|06. Build vs Buy]]
│   └── [[Tradeoff Analysis Frameworks|07. Tradeoff Analysis Frameworks]]
├── [[Large Scale Migrations|36. Large Scale Migrations]]
│   ├── [[Monolith to Microservices|01. Monolith to Microservices]]
│   ├── [[Strangler Fig at Scale|02. Strangler Fig at Scale]]
│   ├── [[Zero Downtime Migration|03. Zero Downtime Migration]]
│   ├── [[Expand Contract Pattern|04. Expand Contract Pattern]]
│   ├── [[Dual Write and Backfill|05. Dual Write and Backfill]]
│   ├── [[Data Migration at Scale|06. Data Migration at Scale]]
│   └── [[Deprecation Strategy|07. Deprecation Strategy]]
├── [[Sociotechnical Org Design|37. Sociotechnical Org Design]]
│   ├── [[Conways Law|01. Conways Law]]
│   ├── [[Team Topologies|02. Team Topologies]]
│   ├── [[Platform Engineering IDP|03. Platform Engineering IDP]]
│   ├── [[Ownership and Boundaries|04. Ownership and Boundaries]]
│   └── [[Cognitive Load (Sociotechnical Org Design)|05. Cognitive Load]]
├── [[Cost Efficiency FinOps|38. Cost Efficiency FinOps]]
│   ├── [[Cost Modeling|01. Cost Modeling]]
│   ├── [[Capacity Planning|02. Capacity Planning]]
│   ├── [[Efficiency As a Feature|03. Efficiency As a Feature]]
│   ├── [[Hardware Aware Design|04. Hardware Aware Design]]
│   ├── [[Performance Economics|05. Performance Economics]]
│   ├── [[Cloud Cost Optimization|06. Cloud Cost Optimization]]
│   ├── [[Storage Tiering and Data Lifecycle|07. Storage Tiering and Data Lifecycle]]
│   └── [[Data Transfer and Egress Costs|08. Data Transfer and Egress Costs]]
├── [[Global Multi Region|39. Global Multi Region]]
│   ├── [[Active Active Architecture|01. Active Active Architecture]]
│   ├── [[Data Sovereignty and Residency|02. Data Sovereignty and Residency]]
│   ├── [[Geo Routing|03. Geo Routing]]
│   ├── [[Global Consistency|04. Global Consistency]]
│   ├── [[Conflict Resolution|05. Conflict Resolution]]
│   └── [[Follow the Sun|06. Follow the Sun]]
├── [[SRE Reliability Engineering|40. SRE Reliability Engineering]]
│   ├── [[Error Budgets|01. Error Budgets]]
│   ├── [[SLO Ownership|02. SLO Ownership]]
│   ├── [[Incident Management|03. Incident Management]]
│   ├── [[Postmortems|04. Postmortems]]
│   ├── [[Toil Reduction|05. Toil Reduction]]
│   ├── [[Load Shedding|06. Load Shedding]]
│   └── [[Graceful Degradation|07. Graceful Degradation]]
├── [[Performance Engineering|41. Performance Engineering]]
│   ├── [[Tail Latency P99 P99.9|01. Tail Latency P99 P99.9]]
│   ├── [[Coordinated Omission|02. Coordinated Omission]]
│   ├── [[Hedged Requests|03. Hedged Requests]]
│   ├── [[Backpressure Deep|04. Backpressure Deep]]
│   ├── [[Queueing Theory Littles Law|05. Queueing Theory Littles Law]]
│   ├── [[Universal Scalability Law|06. Universal Scalability Law]]
│   ├── [[Amdahls Law|07. Amdahls Law]]
│   ├── [[Mechanical Sympathy Cpu Cache Numa|08. Mechanical Sympathy Cpu Cache Numa]]
│   ├── [[Zero Copy IO Sendfile IO Uring|09. Zero Copy IO Sendfile IO Uring]]
│   ├── [[Async IO and Event Loops|10. Async IO and Event Loops]]
│   ├── [[Thread Per Core Shared Nothing|11. Thread Per Core Shared Nothing]]
│   ├── [[Kernel Bypass Networking Dpdk Rdma|12. Kernel Bypass Networking Dpdk Rdma]]
│   ├── [[Continuous Profiling and Flame Graphs|13. Continuous Profiling and Flame Graphs]]
│   ├── [[Zero Copy Serialization Flatbuffers Arrow|14. Zero Copy Serialization Flatbuffers Arrow]]
│   └── [[Batching and Vectorized Processing|15. Batching and Vectorized Processing]]
├── [[Data Governance Contracts|42. Data Governance Contracts]]
│   ├── [[Schema Registry|01. Schema Registry]]
│   ├── [[Data Contracts|02. Data Contracts]]
│   ├── [[Data Lineage|03. Data Lineage]]
│   ├── [[Data Quality|04. Data Quality]]
│   ├── [[Master Data Management|05. Master Data Management]]
│   ├── [[Privacy by Design|06. Privacy by Design]]
│   └── [[Data Mesh|07. Data Mesh]]
└── [[Payments and Fintech|43. Payments and Fintech]]
│   ├── [[Payments Ecosystem and Rails|01. Payments Ecosystem and Rails]]
│   ├── [[Card Payments and Networks|02. Card Payments and Networks]]
│   ├── [[Bank Transfers ACH and Wire|03. Bank Transfers ACH and Wire]]
│   ├── [[Digital Wallets and Mobile Pay|04. Digital Wallets and Mobile Pay]]
│   ├── [[Realtime Payments UPI and FPS|05. Realtime Payments UPI and FPS]]
│   ├── [[Ledgers and Double Entry Accounting|06. Ledgers and Double Entry Accounting]]
│   ├── [[Reconciliation|07. Reconciliation]]
│   ├── [[Idempotency and Exactly Once Payments|08. Idempotency and Exactly Once Payments]]
│   ├── [[Hotspot Accounts and Contention|09. Hotspot Accounts and Contention]]
│   ├── [[Cross Border and Fx|10. Cross Border and Fx]]
│   └── [[Fraud Risk and PCI Compliance|11. Fraud Risk and PCI Compliance]]
```

---

## 🏛️ Core Knowledge Pillars

### 1. 📂 [[Introduction|01. Introduction]]
- 📂 [[What Is System Design]] — Architectural blueprints and trade-offs for What Is System Design.
- 📂 [[How to Approach]] — Architectural blueprints and trade-offs for How to Approach.
- 📂 [[Functional vs Nonfunctional]] — Architectural blueprints and trade-offs for Functional vs Nonfunctional.
- 📂 [[Key Characteristics]] — Architectural blueprints and trade-offs for Key Characteristics.
- 📂 [[Numbers Every Engineer Should Know]] — Architectural blueprints and trade-offs for Numbers Every Engineer Should Know.
### 2. 📂 [[Tradeoffs Framework|02. Tradeoffs Framework]]
- 📂 [[CAP Theorem (Tradeoffs Framework)]] — Architectural blueprints and trade-offs for CAP Theorem.
- 📂 [[PACELC]] — Architectural blueprints and trade-offs for PACELC.
- 📂 [[Consistency vs Availability]] — Architectural blueprints and trade-offs for Consistency vs Availability.
- 📂 [[Consistency Models (Tradeoffs Framework)]] — Architectural blueprints and trade-offs for Consistency Models.
### 3. 📂 [[Capacity Estimation|03. Capacity Estimation]]
- 📂 [[QPS]] — Architectural blueprints and trade-offs for QPS.
- 📂 [[Storage]] — Architectural blueprints and trade-offs for Storage.
- 📂 [[Bandwidth]] — Architectural blueprints and trade-offs for Bandwidth.
- 📂 [[Latency Budgets]] — Architectural blueprints and trade-offs for Latency Budgets.
### 4. 📂 [[Back of Envelope|04. Back of Envelope]]
- 📂 [[Number Tables]] — Architectural blueprints and trade-offs for Number Tables.
- 📂 [[Fermi Estimation]] — Architectural blueprints and trade-offs for Fermi Estimation.
### 5. 📂 [[Networking Protocols|05. Networking Protocols]]
- 📂 [[Osi and TCP Ip]] — Architectural blueprints and trade-offs for Osi and TCP Ip.
- 📂 [[TCP vs UDP]] — Architectural blueprints and trade-offs for TCP vs UDP.
- 📂 [[TLS and Https]] — Architectural blueprints and trade-offs for TLS and Https.
- 📂 [[HTTP Evolution 1 2 3 QUIC]] — Architectural blueprints and trade-offs for HTTP Evolution 1 2 3 QUIC.
- 📂 [[Websockets]] — Architectural blueprints and trade-offs for Websockets.
- 📂 [[Server Sent Events]] — Architectural blueprints and trade-offs for Server Sent Events.
- 📂 [[Long Polling and Streaming]] — Architectural blueprints and trade-offs for Long Polling and Streaming.
- 📂 [[Network Proxies and NAT]] — Architectural blueprints and trade-offs for Network Proxies and NAT.
- 📂 [[Congestion Control and TCP Tuning]] — Architectural blueprints and trade-offs for Congestion Control and TCP Tuning.
- 📂 [[Container and Overlay Networking]] — Architectural blueprints and trade-offs for Container and Overlay Networking.
- 📂 [[BGP and Internet Routing]] — Architectural blueprints and trade-offs for BGP and Internet Routing.
### 6. 📂 [[Domain Name System|06. Domain Name System]]
- 📂 [[DNS Resolution Flow]] — Architectural blueprints and trade-offs for DNS Resolution Flow.
- 📂 [[Record Types]] — Architectural blueprints and trade-offs for Record Types.
- 📂 [[DNS Load Balancing]] — Architectural blueprints and trade-offs for DNS Load Balancing.
- 📂 [[DNS Caching and Ttl]] — Architectural blueprints and trade-offs for DNS Caching and Ttl.
- 📂 [[Geodns and Anycast]] — Architectural blueprints and trade-offs for Geodns and Anycast.
### 7. 📂 [[Content Delivery Networks|07. Content Delivery Networks]]
- 📂 [[Pull CDN]] — Architectural blueprints and trade-offs for Pull CDN.
- 📂 [[Push CDN]] — Architectural blueprints and trade-offs for Push CDN.
- 📂 [[Cache Invalidation]] — Architectural blueprints and trade-offs for Cache Invalidation.
- 📂 [[Edge Locations]] — Architectural blueprints and trade-offs for Edge Locations.
- 📂 [[CDN Security]] — Architectural blueprints and trade-offs for CDN Security.
### 8. 📂 [[Load Balancers|08. Load Balancers]]
- 📂 [[Lb vs Reverse Proxy]] — Architectural blueprints and trade-offs for Lb vs Reverse Proxy.
- 📂 [[Load Balancing Algorithms]] — Architectural blueprints and trade-offs for Load Balancing Algorithms.
- 📂 [[Layer 4 Load Balancing]] — Architectural blueprints and trade-offs for Layer 4 Load Balancing.
- 📂 [[Layer 7 Load Balancing]] — Architectural blueprints and trade-offs for Layer 7 Load Balancing.
- 📂 [[Health Checks and Failover]] — Architectural blueprints and trade-offs for Health Checks and Failover.
- 📂 [[Horizontal Scaling]] — Architectural blueprints and trade-offs for Horizontal Scaling.
- 📂 [[Global Server Load Balancing]] — Architectural blueprints and trade-offs for Global Server Load Balancing.
### 9. 📂 [[Communication|09. Communication]]
- 📂 [[HTTP (Communication)]] — Architectural blueprints and trade-offs for HTTP.
- 📂 [[TCP (Communication)]] — Architectural blueprints and trade-offs for TCP.
- 📂 [[UDP (Communication)]] — Architectural blueprints and trade-offs for UDP.
- 📂 [[RPC]] — Architectural blueprints and trade-offs for RPC.
- 📂 [[GRPC]] — Architectural blueprints and trade-offs for GRPC.
- 📂 [[REST]] — Architectural blueprints and trade-offs for REST.
- 📂 [[Graphql]] — Architectural blueprints and trade-offs for Graphql.
- 📂 [[Idempotent Operations]] — Architectural blueprints and trade-offs for Idempotent Operations.
### 10. 📂 [[Application Layer|10. Application Layer]]
- 📂 [[Microservices]] — Architectural blueprints and trade-offs for Microservices.
- 📂 [[Monolith vs Microservices]] — Architectural blueprints and trade-offs for Monolith vs Microservices.
- 📂 [[Service Discovery]] — Architectural blueprints and trade-offs for Service Discovery.
- 📂 [[API Composition]] — Architectural blueprints and trade-offs for API Composition.
- 📂 [[Stateless Design]] — Architectural blueprints and trade-offs for Stateless Design.
- 📂 [[Service Mesh Intro]] — Architectural blueprints and trade-offs for Service Mesh Intro.
- 📂 [[Serverless Faas]] — Architectural blueprints and trade-offs for Serverless Faas.
- 📂 [[Peer to Peer Architecture]] — Architectural blueprints and trade-offs for Peer to Peer Architecture.
### 11. 📂 [[API Design at Scale|11. API Design at Scale]]
- 📂 [[API Gateway]] — Architectural blueprints and trade-offs for API Gateway.
- 📂 [[REST Design at Scale]] — Architectural blueprints and trade-offs for REST Design at Scale.
- 📂 [[Graphql Federation]] — Architectural blueprints and trade-offs for Graphql Federation.
- 📂 [[GRPC and Streaming]] — Architectural blueprints and trade-offs for GRPC and Streaming.
- 📂 [[Versioning and Deprecation]] — Architectural blueprints and trade-offs for Versioning and Deprecation.
- 📂 [[Pagination and Filtering]] — Architectural blueprints and trade-offs for Pagination and Filtering.
- 📂 [[Idempotency and Retries]] — Architectural blueprints and trade-offs for Idempotency and Retries.
- 📂 [[Webhooks]] — Architectural blueprints and trade-offs for Webhooks.
- 📂 [[Backends for Frontend]] — Architectural blueprints and trade-offs for Backends for Frontend.
### 12. 📂 [[Databases (System Design)|12. Databases]]
- 📂 [[Relational Rdbms]] — Architectural blueprints and trade-offs for Relational Rdbms.
- 📂 [[Key Value]] — Architectural blueprints and trade-offs for Key Value.
- 📂 [[Document]] — Architectural blueprints and trade-offs for Document.
- 📂 [[Wide Column]] — Architectural blueprints and trade-offs for Wide Column.
- 📂 [[Column Oriented Olap]] — Architectural blueprints and trade-offs for Column Oriented Olap.
- 📂 [[Graph]] — Architectural blueprints and trade-offs for Graph.
- 📂 [[Time Series]] — Architectural blueprints and trade-offs for Time Series.
- 📂 [[Search Engine]] — Architectural blueprints and trade-offs for Search Engine.
- 📂 [[Vector]] — Architectural blueprints and trade-offs for Vector.
- 📂 [[Newsql Distributed SQL]] — Architectural blueprints and trade-offs for Newsql Distributed SQL.
- 📂 [[Replication (Databases)]] — Architectural blueprints and trade-offs for Replication.
- 📂 [[Sharding and Partitioning]] — Architectural blueprints and trade-offs for Sharding and Partitioning.
- 📂 [[Indexing]] — Architectural blueprints and trade-offs for Indexing.
- 📂 [[Transactions and Isolation]] — Architectural blueprints and trade-offs for Transactions and Isolation.
- 📂 [[Denormalization]] — Architectural blueprints and trade-offs for Denormalization.
- 📂 [[SQL Tuning]] — Architectural blueprints and trade-offs for SQL Tuning.
- 📂 [[SQL vs NoSQL]] — Architectural blueprints and trade-offs for SQL vs NoSQL.
- 📂 [[Oltp vs Olap]] — Architectural blueprints and trade-offs for Oltp vs Olap.
- 📂 [[Polyglot Persistence]] — Architectural blueprints and trade-offs for Polyglot Persistence.
- 📂 [[Choosing a Database]] — Architectural blueprints and trade-offs for Choosing a Database.
- 📂 [[Connection Pooling Pgbouncer Proxysql]] — Architectural blueprints and trade-offs for Connection Pooling Pgbouncer Proxysql.
- 📂 [[Database HA and Failover]] — Architectural blueprints and trade-offs for Database HA and Failover.
- 📂 [[Backup Restore and Pitr]] — Architectural blueprints and trade-offs for Backup Restore and Pitr.
### 13. 📂 [[Storage Systems|13. Storage Systems]]
- 📂 [[Object vs Block vs File]] — Architectural blueprints and trade-offs for Object vs Block vs File.
- 📂 [[Distributed File Systems Gfs Hdfs]] — Architectural blueprints and trade-offs for Distributed File Systems Gfs Hdfs.
- 📂 [[Blob Storage S3]] — Architectural blueprints and trade-offs for Blob Storage S3.
- 📂 [[LSM Trees and Btrees]] — Architectural blueprints and trade-offs for LSM Trees and Btrees.
- 📂 [[Data Warehouse vs Lake]] — Architectural blueprints and trade-offs for Data Warehouse vs Lake.
- 📂 [[File Formats Parquet Orc Iceberg]] — Architectural blueprints and trade-offs for File Formats Parquet Orc Iceberg.
- 📂 [[Erasure Coding and Durability]] — Architectural blueprints and trade-offs for Erasure Coding and Durability.
### 14. 📂 [[Caching|14. Caching]]
- 📂 [[Cache Aside]] — Architectural blueprints and trade-offs for Cache Aside.
- 📂 [[Write Through]] — Architectural blueprints and trade-offs for Write Through.
- 📂 [[Write Behind]] — Architectural blueprints and trade-offs for Write Behind.
- 📂 [[Refresh Ahead]] — Architectural blueprints and trade-offs for Refresh Ahead.
- 📂 [[Eviction Policies]] — Architectural blueprints and trade-offs for Eviction Policies.
- 📂 [[Types of Caching]] — Architectural blueprints and trade-offs for Types of Caching.
- 📂 [[Cache Invalidation (Caching)]] — Architectural blueprints and trade-offs for Cache Invalidation.
- 📂 [[Cache Stampede and Hot Keys]] — Architectural blueprints and trade-offs for Cache Stampede and Hot Keys.
### 15. 📂 [[Data Streaming|15. Data Streaming]]
- 📂 [[Batch Processing Mapreduce]] — Architectural blueprints and trade-offs for Batch Processing Mapreduce.
- 📂 [[Apache Spark]] — Architectural blueprints and trade-offs for Apache Spark.
- 📂 [[Stream Processing]] — Architectural blueprints and trade-offs for Stream Processing.
- 📂 [[Apache Kafka]] — Architectural blueprints and trade-offs for Apache Kafka.
- 📂 [[Lambda vs Kappa Architecture]] — Architectural blueprints and trade-offs for Lambda vs Kappa Architecture.
- 📂 [[Data Lake and Warehouse]] — Architectural blueprints and trade-offs for Data Lake and Warehouse.
- 📂 [[Change Data Capture]] — Architectural blueprints and trade-offs for Change Data Capture.
- 📂 [[Etl vs Elt]] — Architectural blueprints and trade-offs for Etl vs Elt.
- 📂 [[Apache Flink]] — Architectural blueprints and trade-offs for Apache Flink.
- 📂 [[Apache Pulsar]] — Architectural blueprints and trade-offs for Apache Pulsar.
- 📂 [[Trino Presto Distributed Query]] — Architectural blueprints and trade-offs for Trino Presto Distributed Query.
- 📂 [[Workflow Orchestration Airflow Dagster]] — Architectural blueprints and trade-offs for Workflow Orchestration Airflow Dagster.
- 📂 [[Lakehouse and Table Formats]] — Architectural blueprints and trade-offs for Lakehouse and Table Formats.
- 📂 [[Dimensional Data Modeling]] — Architectural blueprints and trade-offs for Dimensional Data Modeling.
- 📂 [[Reverse Etl and Operational Analytics]] — Architectural blueprints and trade-offs for Reverse Etl and Operational Analytics.
### 16. 📂 [[Asynchronism|16. Asynchronism]]
- 📂 [[Message Queues]] — Architectural blueprints and trade-offs for Message Queues.
- 📂 [[Task Queues]] — Architectural blueprints and trade-offs for Task Queues.
- 📂 [[Back Pressure]] — Architectural blueprints and trade-offs for Back Pressure.
- 📂 [[Dead Letter Queues]] — Architectural blueprints and trade-offs for Dead Letter Queues.
- 📂 [[Delivery Guarantees]] — Architectural blueprints and trade-offs for Delivery Guarantees.
- 📂 [[Rabbitmq]] — Architectural blueprints and trade-offs for Rabbitmq.
- 📂 [[Nats]] — Architectural blueprints and trade-offs for Nats.
### 17. 📂 [[Background Jobs|17. Background Jobs]]
- 📂 [[Event Driven]] — Architectural blueprints and trade-offs for Event Driven.
- 📂 [[Schedule Driven]] — Architectural blueprints and trade-offs for Schedule Driven.
- 📂 [[Returning Results]] — Architectural blueprints and trade-offs for Returning Results.
- 📂 [[Retries and Idempotency]] — Architectural blueprints and trade-offs for Retries and Idempotency.
- 📂 [[Durable Execution Temporal]] — Architectural blueprints and trade-offs for Durable Execution Temporal.
### 18. 📂 [[Concurrency Coordination|18. Concurrency Coordination]]
- 📂 [[Idempotency Keys]] — Architectural blueprints and trade-offs for Idempotency Keys.
- 📂 [[Leases and Fencing]] — Architectural blueprints and trade-offs for Leases and Fencing.
- 📂 [[Exactly Once Semantics]] — Architectural blueprints and trade-offs for Exactly Once Semantics.
- 📂 [[Optimistic vs Pessimistic Locking]] — Architectural blueprints and trade-offs for Optimistic vs Pessimistic Locking.
- 📂 [[Coordination Services]] — Architectural blueprints and trade-offs for Coordination Services.
- 📂 [[Gossip Protocol]] — Architectural blueprints and trade-offs for Gossip Protocol.
- 📂 [[Atomic Commit 2PC 3PC TCC]] — Architectural blueprints and trade-offs for Atomic Commit 2PC 3PC TCC.
### 19. 📂 [[Building Blocks|19. Building Blocks]]
- 📂 [[Rate Limiter]] — Architectural blueprints and trade-offs for Rate Limiter.
- 📂 [[Consistent Hashing]] — Architectural blueprints and trade-offs for Consistent Hashing.
- 📂 [[Unique Id Generator]] — Architectural blueprints and trade-offs for Unique Id Generator.
- 📂 [[Distributed Key Value Store]] — Architectural blueprints and trade-offs for Distributed Key Value Store.
- 📂 [[Distributed Cache]] — Architectural blueprints and trade-offs for Distributed Cache.
- 📂 [[Distributed Message Queue]] — Architectural blueprints and trade-offs for Distributed Message Queue.
- 📂 [[Pub Sub System]] — Architectural blueprints and trade-offs for Pub Sub System.
- 📂 [[Blob Object Store]] — Architectural blueprints and trade-offs for Blob Object Store.
- 📂 [[Distributed Search Typeahead]] — Architectural blueprints and trade-offs for Distributed Search Typeahead.
- 📂 [[Distributed Task Scheduler]] — Architectural blueprints and trade-offs for Distributed Task Scheduler.
- 📂 [[Distributed Lock]] — Architectural blueprints and trade-offs for Distributed Lock.
- 📂 [[Distributed Logging]] — Architectural blueprints and trade-offs for Distributed Logging.
- 📂 [[Sharded Counters Leaderboard]] — Architectural blueprints and trade-offs for Sharded Counters Leaderboard.
- 📂 [[Bloom Filter (Building Blocks)]] — Architectural blueprints and trade-offs for Bloom Filter.
### 20. 📂 [[Reliability Patterns|20. Reliability Patterns]]
- 📂 [[Circuit Breaker]] — Architectural blueprints and trade-offs for Circuit Breaker.
- 📂 [[Bulkhead]] — Architectural blueprints and trade-offs for Bulkhead.
- 📂 [[Retry]] — Architectural blueprints and trade-offs for Retry.
- 📂 [[Throttling]] — Architectural blueprints and trade-offs for Throttling.
- 📂 [[Health Endpoint Monitoring]] — Architectural blueprints and trade-offs for Health Endpoint Monitoring.
- 📂 [[Leader Election]] — Architectural blueprints and trade-offs for Leader Election.
- 📂 [[Compensating Transaction]] — Architectural blueprints and trade-offs for Compensating Transaction.
- 📂 [[Deployment Stamps and Geodes]] — Architectural blueprints and trade-offs for Deployment Stamps and Geodes.
- 📂 [[Queue Based Load Leveling]] — Architectural blueprints and trade-offs for Queue Based Load Leveling.
- 📂 [[Redundancy and Failure Domains]] — Architectural blueprints and trade-offs for Redundancy and Failure Domains.
- 📂 [[Shuffle Sharding]] — Architectural blueprints and trade-offs for Shuffle Sharding.
### 21. 📂 [[Cloud Design Patterns|21. Cloud Design Patterns]]
- 📂 [[Strangler Fig]] — Architectural blueprints and trade-offs for Strangler Fig.
- 📂 [[Sidecar]] — Architectural blueprints and trade-offs for Sidecar.
- 📂 [[Ambassador]] — Architectural blueprints and trade-offs for Ambassador.
- 📂 [[Anti Corruption Layer]] — Architectural blueprints and trade-offs for Anti Corruption Layer.
- 📂 [[CQRS]] — Architectural blueprints and trade-offs for CQRS.
- 📂 [[Event Sourcing]] — Architectural blueprints and trade-offs for Event Sourcing.
- 📂 [[Materialized View]] — Architectural blueprints and trade-offs for Materialized View.
- 📂 [[Pipes and Filters]] — Architectural blueprints and trade-offs for Pipes and Filters.
- 📂 [[External Config Store]] — Architectural blueprints and trade-offs for External Config Store.
- 📂 [[Valet Key]] — Architectural blueprints and trade-offs for Valet Key.
- 📂 [[Claim Check]] — Architectural blueprints and trade-offs for Claim Check.
- 📂 [[Competing Consumers]] — Architectural blueprints and trade-offs for Competing Consumers.
- 📂 [[Publisher Subscriber]] — Architectural blueprints and trade-offs for Publisher Subscriber.
- 📂 [[Transactional Outbox]] — Architectural blueprints and trade-offs for Transactional Outbox.
### 22. 📂 [[Performance Antipatterns|22. Performance Antipatterns]]
- 📂 [[Busy Database]] — Architectural blueprints and trade-offs for Busy Database.
- 📂 [[Busy Frontend]] — Architectural blueprints and trade-offs for Busy Frontend.
- 📂 [[Chatty IO]] — Architectural blueprints and trade-offs for Chatty IO.
- 📂 [[Extraneous Fetching]] — Architectural blueprints and trade-offs for Extraneous Fetching.
- 📂 [[Improper Instantiation]] — Architectural blueprints and trade-offs for Improper Instantiation.
- 📂 [[Monolithic Persistence]] — Architectural blueprints and trade-offs for Monolithic Persistence.
- 📂 [[Noisy Neighbor]] — Architectural blueprints and trade-offs for Noisy Neighbor.
- 📂 [[Synchronous IO]] — Architectural blueprints and trade-offs for Synchronous IO.
- 📂 [[Retry Storm]] — Architectural blueprints and trade-offs for Retry Storm.
- 📂 [[No Caching]] — Architectural blueprints and trade-offs for No Caching.
### 23. 📂 [[Monitoring|23. Monitoring]]
- 📂 [[Health Monitoring]] — Architectural blueprints and trade-offs for Health Monitoring.
- 📂 [[Availability Monitoring]] — Architectural blueprints and trade-offs for Availability Monitoring.
- 📂 [[Performance Monitoring]] — Architectural blueprints and trade-offs for Performance Monitoring.
- 📂 [[Security Monitoring]] — Architectural blueprints and trade-offs for Security Monitoring.
- 📂 [[Usage Monitoring]] — Architectural blueprints and trade-offs for Usage Monitoring.
- 📂 [[Instrumentation]] — Architectural blueprints and trade-offs for Instrumentation.
- 📂 [[Visualization and Alerts]] — Architectural blueprints and trade-offs for Visualization and Alerts.
- 📂 [[Synthetic Monitoring and Rum]] — Architectural blueprints and trade-offs for Synthetic Monitoring and Rum.
### 24. 📂 [[Observability|24. Observability]]
- 📂 [[Logs Metrics Traces]] — Architectural blueprints and trade-offs for Logs Metrics Traces.
- 📂 [[SLO SLI Error Budgets]] — Architectural blueprints and trade-offs for SLO SLI Error Budgets.
- 📂 [[Red and Use Methods]] — Architectural blueprints and trade-offs for Red and Use Methods.
- 📂 [[Distributed Tracing]] — Architectural blueprints and trade-offs for Distributed Tracing.
- 📂 [[Metrics Pipelines]] — Architectural blueprints and trade-offs for Metrics Pipelines.
- 📂 [[Log Aggregation]] — Architectural blueprints and trade-offs for Log Aggregation.
- 📂 [[Alerting and on Call]] — Architectural blueprints and trade-offs for Alerting and on Call.
- 📂 [[Opentelemetry]] — Architectural blueprints and trade-offs for Opentelemetry.
- 📂 [[Cardinality and Metrics Cost]] — Architectural blueprints and trade-offs for Cardinality and Metrics Cost.
- 📂 [[Trace and Log Sampling Strategies]] — Architectural blueprints and trade-offs for Trace and Log Sampling Strategies.
### 25. 📂 [[Chaos Engineering|25. Chaos Engineering]]
- 📂 [[Failure Modes]] — Architectural blueprints and trade-offs for Failure Modes.
- 📂 [[Fault Injection]] — Architectural blueprints and trade-offs for Fault Injection.
- 📂 [[Game Days]] — Architectural blueprints and trade-offs for Game Days.
- 📂 [[Resilience Testing]] — Architectural blueprints and trade-offs for Resilience Testing.
- 📂 [[Blast Radius and Recovery]] — Architectural blueprints and trade-offs for Blast Radius and Recovery.
### 26. 📂 [[Deployment Infrastructure|26. Deployment Infrastructure]]
- 📂 [[Containers and Docker]] — Architectural blueprints and trade-offs for Containers and Docker.
- 📂 [[Kubernetes Orchestration]] — Architectural blueprints and trade-offs for Kubernetes Orchestration.
- 📂 [[Deployment Strategies]] — Architectural blueprints and trade-offs for Deployment Strategies.
- 📂 [[Ci Cd Pipelines]] — Architectural blueprints and trade-offs for Ci Cd Pipelines.
- 📂 [[Infrastructure As Code]] — Architectural blueprints and trade-offs for Infrastructure As Code.
- 📂 [[Multi Region Deployment]] — Architectural blueprints and trade-offs for Multi Region Deployment.
- 📂 [[Disaster Recovery]] — Architectural blueprints and trade-offs for Disaster Recovery.
- 📂 [[Autoscaling]] — Architectural blueprints and trade-offs for Autoscaling.
- 📂 [[Cloud Network Architecture Vpc]] — Architectural blueprints and trade-offs for Cloud Network Architecture Vpc.
- 📂 [[Gitops Argocd Flux]] — Architectural blueprints and trade-offs for Gitops Argocd Flux.
### 27. 📂 [[Security at Scale|27. Security at Scale]]
- 📂 [[Authentication]] — Architectural blueprints and trade-offs for Authentication.
- 📂 [[Authorization]] — Architectural blueprints and trade-offs for Authorization.
- 📂 [[Oauth2 and Oidc]] — Architectural blueprints and trade-offs for Oauth2 and Oidc.
- 📂 [[JWT and Tokens]] — Architectural blueprints and trade-offs for JWT and Tokens.
- 📂 [[Encryption at REST and Transit]] — Architectural blueprints and trade-offs for Encryption at REST and Transit.
- 📂 [[Secrets Management]] — Architectural blueprints and trade-offs for Secrets Management.
- 📂 [[DDoS Mitigation]] — Architectural blueprints and trade-offs for DDoS Mitigation.
- 📂 [[WAF and API Security]] — Architectural blueprints and trade-offs for WAF and API Security.
- 📂 [[Rate Limiting for Abuse]] — Architectural blueprints and trade-offs for Rate Limiting for Abuse.
- 📂 [[Devsecops and Supply Chain Security]] — Architectural blueprints and trade-offs for Devsecops and Supply Chain Security.
- 📂 [[Zero Trust Architecture]] — Architectural blueprints and trade-offs for Zero Trust Architecture.
- 📂 [[PKI and Certificate Management]] — Architectural blueprints and trade-offs for PKI and Certificate Management.
- 📂 [[Threat Modeling Stride]] — Architectural blueprints and trade-offs for Threat Modeling Stride.
- 📂 [[Envelope Encryption and KMS]] — Architectural blueprints and trade-offs for Envelope Encryption and KMS.
### 28. 📂 [[Data Privacy Compliance|28. Data Privacy Compliance]]
- 📂 [[PII and Data Classification]] — Architectural blueprints and trade-offs for PII and Data Classification.
- 📂 [[GDPR and Right to Be Forgotten]] — Architectural blueprints and trade-offs for GDPR and Right to Be Forgotten.
- 📂 [[Data Residency]] — Architectural blueprints and trade-offs for Data Residency.
- 📂 [[Audit Logging (Data Privacy Compliance)]] — Architectural blueprints and trade-offs for Audit Logging.
- 📂 [[Encryption Key Lifecycle]] — Architectural blueprints and trade-offs for Encryption Key Lifecycle.
### 29. 📂 [[Multi Tenancy SaaS|29. Multi Tenancy SaaS]]
- 📂 [[Tenant Isolation Models]] — Architectural blueprints and trade-offs for Tenant Isolation Models.
- 📂 [[Data Partitioning Per Tenant]] — Architectural blueprints and trade-offs for Data Partitioning Per Tenant.
- 📂 [[Noisy Neighbor Mitigation]] — Architectural blueprints and trade-offs for Noisy Neighbor Mitigation.
- 📂 [[Per Tenant Scaling and Limits]] — Architectural blueprints and trade-offs for Per Tenant Scaling and Limits.
- 📂 [[Tenant Onboarding and Config]] — Architectural blueprints and trade-offs for Tenant Onboarding and Config.
### 30. 📂 [[Geospatial Systems|30. Geospatial Systems]]
- 📂 [[Geohashing]] — Architectural blueprints and trade-offs for Geohashing.
- 📂 [[Quadtrees]] — Architectural blueprints and trade-offs for Quadtrees.
- 📂 [[S2 and H3]] — Architectural blueprints and trade-offs for S2 and H3.
- 📂 [[Proximity Search]] — Architectural blueprints and trade-offs for Proximity Search.
- 📂 [[Map Tiling and Routing]] — Architectural blueprints and trade-offs for Map Tiling and Routing.
### 31. 📂 [[ML Recommendation Systems|31. ML Recommendation Systems]]
- 📂 [[Recommendation Architecture]] — Architectural blueprints and trade-offs for Recommendation Architecture.
- 📂 [[Feature Store]] — Architectural blueprints and trade-offs for Feature Store.
- 📂 [[Candidate Generation]] — Architectural blueprints and trade-offs for Candidate Generation.
- 📂 [[Ranking and Scoring]] — Architectural blueprints and trade-offs for Ranking and Scoring.
- 📂 [[Online vs Offline Inference]] — Architectural blueprints and trade-offs for Online vs Offline Inference.
- 📂 [[Ab Testing and Feedback Loops]] — Architectural blueprints and trade-offs for Ab Testing and Feedback Loops.
- 📂 [[LLM Inference and Serving]] — Architectural blueprints and trade-offs for LLM Inference and Serving.
- 📂 [[Retrieval Augmented Generation Rag]] — Architectural blueprints and trade-offs for Retrieval Augmented Generation Rag.
- 📂 [[LLM Application Architecture]] — Architectural blueprints and trade-offs for LLM Application Architecture.
- 📂 [[Ai Agents and Orchestration]] — Architectural blueprints and trade-offs for Ai Agents and Orchestration.
### 32. 📂 [[Classic Problems|32. Classic Problems]]
- 📂 [[Url Shortener]] — Architectural blueprints and trade-offs for Url Shortener.
- 📂 [[News Feed and Timeline]] — Architectural blueprints and trade-offs for News Feed and Timeline.
- 📂 [[Chat and Messaging System]] — Architectural blueprints and trade-offs for Chat and Messaging System.
- 📂 [[Youtube Video Platform]] — Architectural blueprints and trade-offs for Youtube Video Platform.
- 📂 [[Google Drive File Sync]] — Architectural blueprints and trade-offs for Google Drive File Sync.
- 📂 [[Payment System]] — Architectural blueprints and trade-offs for Payment System.
- 📂 [[Digital Wallet]] — Architectural blueprints and trade-offs for Digital Wallet.
- 📂 [[Hotel Reservation System]] — Architectural blueprints and trade-offs for Hotel Reservation System.
- 📂 [[Distributed Email Service]] — Architectural blueprints and trade-offs for Distributed Email Service.
- 📂 [[Instagram Photo Feed]] — Architectural blueprints and trade-offs for Instagram Photo Feed.
- 📂 [[Stack Overflow]] — Architectural blueprints and trade-offs for Stack Overflow.
- 📂 [[Google Docs Collab Editor]] — Architectural blueprints and trade-offs for Google Docs Collab Editor.
- 📂 [[Proximity Service Maps]] — Architectural blueprints and trade-offs for Proximity Service Maps.
- 📂 [[Ticketmaster Booking]] — Architectural blueprints and trade-offs for Ticketmaster Booking.
- 📂 [[Notification System]] — Architectural blueprints and trade-offs for Notification System.
- 📂 [[Live Streaming]] — Architectural blueprints and trade-offs for Live Streaming.
- 📂 [[Distributed Job Scheduler]] — Architectural blueprints and trade-offs for Distributed Job Scheduler.
- 📂 [[Stock Exchange]] — Architectural blueprints and trade-offs for Stock Exchange.
- 📂 [[S3 Object Storage]] — Architectural blueprints and trade-offs for S3 Object Storage.
- 📂 [[Online Judge]] — Architectural blueprints and trade-offs for Online Judge.
- 📂 [[Distributed Analytics Counter]] — Architectural blueprints and trade-offs for Distributed Analytics Counter.
- 📂 [[Web Crawler and Search Engine]] — Architectural blueprints and trade-offs for Web Crawler and Search Engine.
### 33. 📂 [[Real Architectures|33. Real Architectures]]
- 📂 [[Google Spanner]] — Architectural blueprints and trade-offs for Google Spanner.
- 📂 [[Facebook Tao]] — Architectural blueprints and trade-offs for Facebook Tao.
- 📂 [[Amazon Dynamodb]] — Architectural blueprints and trade-offs for Amazon Dynamodb.
- 📂 [[Netflix Stack]] — Architectural blueprints and trade-offs for Netflix Stack.
- 📂 [[Apache Kafka (Real Architectures)]] — Architectural blueprints and trade-offs for Apache Kafka.
- 📂 [[Apache Cassandra]] — Architectural blueprints and trade-offs for Apache Cassandra.
- 📂 [[Redis Internals]] — Architectural blueprints and trade-offs for Redis Internals.
- 📂 [[Discord Realtime]] — Architectural blueprints and trade-offs for Discord Realtime.
- 📂 [[Slack Messaging]] — Architectural blueprints and trade-offs for Slack Messaging.
- 📂 [[Uber Lyft Dispatch]] — Architectural blueprints and trade-offs for Uber Lyft Dispatch.
### 34. 📂 [[Interview Playbook|34. Interview Playbook]]
- 📂 [[Reshaded Framework]] — Architectural blueprints and trade-offs for Reshaded Framework.
- 📂 [[Requirements Clarification]] — Architectural blueprints and trade-offs for Requirements Clarification.
- 📂 [[Capacity Estimation in Interview]] — Architectural blueprints and trade-offs for Capacity Estimation in Interview.
- 📂 [[API Design Step]] — Architectural blueprints and trade-offs for API Design Step.
- 📂 [[High Level Design]] — Architectural blueprints and trade-offs for High Level Design.
- 📂 [[Data Model and Storage]] — Architectural blueprints and trade-offs for Data Model and Storage.
- 📂 [[Deep Dives and Bottlenecks]] — Architectural blueprints and trade-offs for Deep Dives and Bottlenecks.
- 📂 [[Tradeoffs and Wrap Up]] — Architectural blueprints and trade-offs for Tradeoffs and Wrap Up.
- 📂 [[Common Mistakes]] — Architectural blueprints and trade-offs for Common Mistakes.
- 📂 [[Mock Interview Walkthroughs]] — Architectural blueprints and trade-offs for Mock Interview Walkthroughs.
### 35. 📂 [[Architecture Decision Making|35. Architecture Decision Making]]
- 📂 [[Architecture Decision Records]] — Architectural blueprints and trade-offs for Architecture Decision Records.
- 📂 [[Rfc Process]] — Architectural blueprints and trade-offs for Rfc Process.
- 📂 [[Evolutionary Architecture]] — Architectural blueprints and trade-offs for Evolutionary Architecture.
- 📂 [[Fitness Functions]] — Architectural blueprints and trade-offs for Fitness Functions.
- 📂 [[Tech Radar]] — Architectural blueprints and trade-offs for Tech Radar.
- 📂 [[Build vs Buy]] — Architectural blueprints and trade-offs for Build vs Buy.
- 📂 [[Tradeoff Analysis Frameworks]] — Architectural blueprints and trade-offs for Tradeoff Analysis Frameworks.
### 36. 📂 [[Large Scale Migrations|36. Large Scale Migrations]]
- 📂 [[Monolith to Microservices]] — Architectural blueprints and trade-offs for Monolith to Microservices.
- 📂 [[Strangler Fig at Scale]] — Architectural blueprints and trade-offs for Strangler Fig at Scale.
- 📂 [[Zero Downtime Migration]] — Architectural blueprints and trade-offs for Zero Downtime Migration.
- 📂 [[Expand Contract Pattern]] — Architectural blueprints and trade-offs for Expand Contract Pattern.
- 📂 [[Dual Write and Backfill]] — Architectural blueprints and trade-offs for Dual Write and Backfill.
- 📂 [[Data Migration at Scale]] — Architectural blueprints and trade-offs for Data Migration at Scale.
- 📂 [[Deprecation Strategy]] — Architectural blueprints and trade-offs for Deprecation Strategy.
### 37. 📂 [[Sociotechnical Org Design|37. Sociotechnical Org Design]]
- 📂 [[Conways Law]] — Architectural blueprints and trade-offs for Conways Law.
- 📂 [[Team Topologies]] — Architectural blueprints and trade-offs for Team Topologies.
- 📂 [[Platform Engineering IDP]] — Architectural blueprints and trade-offs for Platform Engineering IDP.
- 📂 [[Ownership and Boundaries]] — Architectural blueprints and trade-offs for Ownership and Boundaries.
- 📂 [[Cognitive Load (Sociotechnical Org Design)]] — Architectural blueprints and trade-offs for Cognitive Load.
### 38. 📂 [[Cost Efficiency FinOps|38. Cost Efficiency FinOps]]
- 📂 [[Cost Modeling]] — Architectural blueprints and trade-offs for Cost Modeling.
- 📂 [[Capacity Planning]] — Architectural blueprints and trade-offs for Capacity Planning.
- 📂 [[Efficiency As a Feature]] — Architectural blueprints and trade-offs for Efficiency As a Feature.
- 📂 [[Hardware Aware Design]] — Architectural blueprints and trade-offs for Hardware Aware Design.
- 📂 [[Performance Economics]] — Architectural blueprints and trade-offs for Performance Economics.
- 📂 [[Cloud Cost Optimization]] — Architectural blueprints and trade-offs for Cloud Cost Optimization.
- 📂 [[Storage Tiering and Data Lifecycle]] — Architectural blueprints and trade-offs for Storage Tiering and Data Lifecycle.
- 📂 [[Data Transfer and Egress Costs]] — Architectural blueprints and trade-offs for Data Transfer and Egress Costs.
### 39. 📂 [[Global Multi Region|39. Global Multi Region]]
- 📂 [[Active Active Architecture]] — Architectural blueprints and trade-offs for Active Active Architecture.
- 📂 [[Data Sovereignty and Residency]] — Architectural blueprints and trade-offs for Data Sovereignty and Residency.
- 📂 [[Geo Routing]] — Architectural blueprints and trade-offs for Geo Routing.
- 📂 [[Global Consistency]] — Architectural blueprints and trade-offs for Global Consistency.
- 📂 [[Conflict Resolution]] — Architectural blueprints and trade-offs for Conflict Resolution.
- 📂 [[Follow the Sun]] — Architectural blueprints and trade-offs for Follow the Sun.
### 40. 📂 [[SRE Reliability Engineering|40. SRE Reliability Engineering]]
- 📂 [[Error Budgets]] — Architectural blueprints and trade-offs for Error Budgets.
- 📂 [[SLO Ownership]] — Architectural blueprints and trade-offs for SLO Ownership.
- 📂 [[Incident Management]] — Architectural blueprints and trade-offs for Incident Management.
- 📂 [[Postmortems]] — Architectural blueprints and trade-offs for Postmortems.
- 📂 [[Toil Reduction]] — Architectural blueprints and trade-offs for Toil Reduction.
- 📂 [[Load Shedding]] — Architectural blueprints and trade-offs for Load Shedding.
- 📂 [[Graceful Degradation]] — Architectural blueprints and trade-offs for Graceful Degradation.
### 41. 📂 [[Performance Engineering|41. Performance Engineering]]
- 📂 [[Tail Latency P99 P99.9]] — Architectural blueprints and trade-offs for Tail Latency P99 P99.9.
- 📂 [[Coordinated Omission]] — Architectural blueprints and trade-offs for Coordinated Omission.
- 📂 [[Hedged Requests]] — Architectural blueprints and trade-offs for Hedged Requests.
- 📂 [[Backpressure Deep]] — Architectural blueprints and trade-offs for Backpressure Deep.
- 📂 [[Queueing Theory Littles Law]] — Architectural blueprints and trade-offs for Queueing Theory Littles Law.
- 📂 [[Universal Scalability Law]] — Architectural blueprints and trade-offs for Universal Scalability Law.
- 📂 [[Amdahls Law]] — Architectural blueprints and trade-offs for Amdahls Law.
- 📂 [[Mechanical Sympathy Cpu Cache Numa]] — Architectural blueprints and trade-offs for Mechanical Sympathy Cpu Cache Numa.
- 📂 [[Zero Copy IO Sendfile IO Uring]] — Architectural blueprints and trade-offs for Zero Copy IO Sendfile IO Uring.
- 📂 [[Async IO and Event Loops]] — Architectural blueprints and trade-offs for Async IO and Event Loops.
- 📂 [[Thread Per Core Shared Nothing]] — Architectural blueprints and trade-offs for Thread Per Core Shared Nothing.
- 📂 [[Kernel Bypass Networking Dpdk Rdma]] — Architectural blueprints and trade-offs for Kernel Bypass Networking Dpdk Rdma.
- 📂 [[Continuous Profiling and Flame Graphs]] — Architectural blueprints and trade-offs for Continuous Profiling and Flame Graphs.
- 📂 [[Zero Copy Serialization Flatbuffers Arrow]] — Architectural blueprints and trade-offs for Zero Copy Serialization Flatbuffers Arrow.
- 📂 [[Batching and Vectorized Processing]] — Architectural blueprints and trade-offs for Batching and Vectorized Processing.
### 42. 📂 [[Data Governance Contracts|42. Data Governance Contracts]]
- 📂 [[Schema Registry]] — Architectural blueprints and trade-offs for Schema Registry.
- 📂 [[Data Contracts]] — Architectural blueprints and trade-offs for Data Contracts.
- 📂 [[Data Lineage]] — Architectural blueprints and trade-offs for Data Lineage.
- 📂 [[Data Quality]] — Architectural blueprints and trade-offs for Data Quality.
- 📂 [[Master Data Management]] — Architectural blueprints and trade-offs for Master Data Management.
- 📂 [[Privacy by Design]] — Architectural blueprints and trade-offs for Privacy by Design.
- 📂 [[Data Mesh]] — Architectural blueprints and trade-offs for Data Mesh.
### 43. 📂 [[Payments and Fintech|43. Payments and Fintech]]
- 📂 [[Payments Ecosystem and Rails]] — Architectural blueprints and trade-offs for Payments Ecosystem and Rails.
- 📂 [[Card Payments and Networks]] — Architectural blueprints and trade-offs for Card Payments and Networks.
- 📂 [[Bank Transfers ACH and Wire]] — Architectural blueprints and trade-offs for Bank Transfers ACH and Wire.
- 📂 [[Digital Wallets and Mobile Pay]] — Architectural blueprints and trade-offs for Digital Wallets and Mobile Pay.
- 📂 [[Realtime Payments UPI and FPS]] — Architectural blueprints and trade-offs for Realtime Payments UPI and FPS.
- 📂 [[Ledgers and Double Entry Accounting]] — Architectural blueprints and trade-offs for Ledgers and Double Entry Accounting.
- 📂 [[Reconciliation]] — Architectural blueprints and trade-offs for Reconciliation.
- 📂 [[Idempotency and Exactly Once Payments]] — Architectural blueprints and trade-offs for Idempotency and Exactly Once Payments.
- 📂 [[Hotspot Accounts and Contention]] — Architectural blueprints and trade-offs for Hotspot Accounts and Contention.
- 📂 [[Cross Border and Fx]] — Architectural blueprints and trade-offs for Cross Border and Fx.
- 📂 [[Fraud Risk and PCI Compliance]] — Architectural blueprints and trade-offs for Fraud Risk and PCI Compliance.

---

## 🔗 Navigation
- ⬆️ Parent: [[Principal SWE]]

