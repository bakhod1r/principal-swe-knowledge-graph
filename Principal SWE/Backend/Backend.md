---
title: Backend
tags:
  - backend
  - architecture
  - distributed-systems
  - principal-swe
parent: "[[Principal SWE]]"
---

# ⚙️ Backend Engineering & High-Scale Infrastructure

Comprehensive, production-grade knowledge architecture covering the full lifecycle of enterprise backend systems: API design standards (REST, gRPC, GraphQL), relational and distributed database internals, advanced PostgreSQL DBA operations, in-memory caching topologies (Redis), distributed full-text search (Elasticsearch), document persistence (MongoDB), and high-throughput microservice infrastructure across 7 master pillars and 107 specialized subdomains.

```text
Backend
│
├── [[Core Backend Engineering|01. Core Backend Engineering]]
│   ├── [[Internet|01. Internet]]
│   ├── [[Pick a Language|02. Pick a Language]]
│   ├── [[Version Control Systems|03. Version Control Systems]]
│   ├── [[Relational Databases|04. Relational Databases]]
│   ├── [[NoSQL Databases|05. NoSQL Databases]]
│   ├── [[More About Databases|06. More About Databases]]
│   ├── [[Learn About APIs|07. Learn About APIs]]
│   ├── [[Caching (Core Backend Engineering)|08. Caching]]
│   ├── [[Web Security|09. Web Security]]
│   ├── [[Testing (Core Backend Engineering)|10. Testing]]
│   ├── [[CI CD|11. CI CD]]
│   ├── [[Software Design Architecture|12. Software Design Architecture]]
│   ├── [[Containerization vs Virtualization|13. Containerization vs Virtualization]]
│   ├── [[Web Servers|14. Web Servers]]
│   ├── [[Search Engines|15. Search Engines]]
│   ├── [[Message Brokers|16. Message Brokers]]
│   ├── [[Real Time Data|17. Real Time Data]]
│   ├── [[Building for Scale|18. Building for Scale]]
│   ├── [[Observability (Core Backend Engineering)|19. Observability]]
│   └── [[Basic Infrastructure Knowledge|20. Basic Infrastructure Knowledge]]
├── [[API Design & Architecture|02. API Design & Architecture]]
│   ├── [[Learn the Basics|01. Learn the Basics]]
│   ├── [[Different API Styles|02. Different API Styles]]
│   ├── [[Building JSON RESTful APIs|03. Building JSON RESTful APIs]]
│   ├── [[API Authentication and Authorization|04. API Authentication and Authorization]]
│   ├── [[API Documentation Tools|05. API Documentation Tools]]
│   ├── [[API Security|06. API Security]]
│   ├── [[API Performance|07. API Performance]]
│   ├── [[API Integration Patterns|08. API Integration Patterns]]
│   ├── [[API Testing|09. API Testing]]
│   ├── [[Real Time APIs|10. Real Time APIs]]
│   ├── [[API Lifecycle Management|11. API Lifecycle Management]]
│   └── [[Standards and Compliance|12. Standards and Compliance]]
├── [[Relational & Distributed Databases|03. Relational & Distributed Databases]]
│   ├── [[Relational Model|01. Relational Model]]
│   ├── [[Normalization and Denormalization|02. Normalization and Denormalization]]
│   ├── [[ER Modeling|03. ER Modeling]]
│   ├── [[SQL DDL DML DQL DCL|04. SQL DDL DML DQL DCL]]
│   ├── [[Views|05. Views]]
│   ├── [[Stored Procedures and Triggers|06. Stored Procedures and Triggers]]
│   ├── [[Transactions and ACID|07. Transactions and ACID]]
│   ├── [[Isolation Levels|08. Isolation Levels]]
│   ├── [[Locking and Concurrency Control|09. Locking and Concurrency Control]]
│   ├── [[MVCC|10. MVCC]]
│   ├── [[BASE and Eventual Consistency|11. BASE and Eventual Consistency]]
│   ├── [[CAP Theorem|12. CAP Theorem]]
│   ├── [[PACELC (Relational & Distributed Databases)|13. PACELC]]
│   ├── [[Indexing (Relational & Distributed Databases)|14. Indexing]]
│   ├── [[Query Optimization|15. Query Optimization]]
│   ├── [[Replication|16. Replication]]
│   ├── [[Partitioning and Sharding|17. Partitioning and Sharding]]
│   ├── [[Database Federation|18. Database Federation]]
│   ├── [[SQL vs NoSQL (Relational & Distributed Databases)|19. SQL vs NoSQL]]
│   ├── [[NoSQL Data Models|20. NoSQL Data Models]]
│   ├── [[OLTP vs OLAP and Warehousing|21. OLTP vs OLAP and Warehousing]]
│   ├── [[Connection Pooling|22. Connection Pooling]]
│   ├── [[Caching at the DB Layer|23. Caching at the DB Layer]]
│   └── [[Backup and Recovery|24. Backup and Recovery]]
├── [[Postgresql Mastery & DBA|04. PostgreSQL Mastery & DBA]]
│   ├── [[Introduction (Postgresql Mastery & DBA)|01. Introduction]]
│   ├── [[Installation and Setup|02. Installation and Setup]]
│   ├── [[Learn SQL|03. Learn SQL]]
│   ├── [[Configuring|04. Configuring]]
│   ├── [[Security|05. Security]]
│   ├── [[Infrastructure Skills|06. Infrastructure Skills]]
│   ├── [[Application Skills|07. Application Skills]]
│   ├── [[Fine Grained Tuning|08. Fine Grained Tuning]]
│   ├── [[Advanced SQL|09. Advanced SQL]]
│   ├── [[Troubleshooting Techniques|10. Troubleshooting Techniques]]
│   ├── [[SQL Optimization Techniques|11. SQL Optimization Techniques]]
│   └── [[Get Involved in Development|12. Get Involved in Development]]
├── [[Redis & in Memory Architectures|05. Redis & In-Memory Architectures]]
│   ├── [[Overview of Redis What Is It|01. Overview of Redis What Is It]]
│   ├── [[Getting Started with Redis|02. Getting Started with Redis]]
│   ├── [[First Steps|03. First Steps]]
│   ├── [[Core Data Structures|04. Core Data Structures]]
│   ├── [[Working with Redis|05. Working with Redis]]
│   ├── [[Advanced Data Structures|06. Advanced Data Structures]]
│   ├── [[Pub Sub|07. Pub Sub]]
│   ├── [[Transactions|08. Transactions]]
│   ├── [[Lua Scripting|09. Lua Scripting]]
│   ├── [[Persistence Options|10. Persistence Options]]
│   ├── [[Replication HA|11. Replication HA]]
│   ├── [[Security (Redis & in Memory Architectures)|12. Security]]
│   ├── [[Monitoring Optimization|13. Monitoring Optimization]]
│   ├── [[Redis Modules|14. Redis Modules]]
│   ├── [[Managing Redis in Production|15. Managing Redis in Production]]
│   └── [[Redis Enterprise|16. Redis Enterprise]]
├── [[Elasticsearch & Distributed Search|06. Elasticsearch & Distributed Search]]
│   ├── [[Introduction (Elasticsearch & Distributed Search)|01. Introduction]]
│   ├── [[Core Architecture|02. Core Architecture]]
│   ├── [[Data Modelling|03. Data Modelling]]
│   ├── [[Data Ingestion|04. Data Ingestion]]
│   ├── [[Search Fundamentals|05. Search Fundamentals]]
│   ├── [[How Search Works|06. How Search Works]]
│   ├── [[Text Analysis|07. Text Analysis]]
│   ├── [[Aggregations|08. Aggregations]]
│   ├── [[Transformations|09. Transformations]]
│   ├── [[Relevance Tuning|10. Relevance Tuning]]
│   ├── [[Production|11. Production]]
│   └── [[Advanced Features|12. Advanced Features]]
└── [[Mongodb & Document Stores|07. MongoDB & Document Stores]]
│   ├── [[Mongodb Basics|01. Mongodb Basics]]
│   ├── [[Data Model Data Types|02. Data Model Data Types]]
│   ├── [[Collections Methods|03. Collections Methods]]
│   ├── [[Useful Concepts|04. Useful Concepts]]
│   ├── [[Query Operators|05. Query Operators]]
│   ├── [[Performance Optimization|06. Performance Optimization]]
│   ├── [[Aggregation|07. Aggregation]]
│   ├── [[Transactions (Mongodb & Document Stores)|08. Transactions]]
│   ├── [[Developer Tools|09. Developer Tools]]
│   ├── [[Scaling Mongodb|10. Scaling Mongodb]]
│   └── [[Mongodb Security|11. Mongodb Security]]
```

---

## 🏛️ Core Knowledge Pillars

### 1. 📂 [[Core Backend Engineering|01. Core Backend Engineering]]
- 📂 [[Internet|01. Internet]] — Architectural blueprints and production operations for Internet.
- 📂 [[Pick a Language|02. Pick a Language]] — Architectural blueprints and production operations for Pick a Language.
- 📂 [[Version Control Systems|03. Version Control Systems]] — Architectural blueprints and production operations for Version Control Systems.
- 📂 [[Relational Databases|04. Relational Databases]] — Architectural blueprints and production operations for Relational Databases.
- 📂 [[NoSQL Databases|05. NoSQL Databases]] — Architectural blueprints and production operations for NoSQL Databases.
- 📂 [[More About Databases|06. More About Databases]] — Architectural blueprints and production operations for More About Databases.
- 📂 [[Learn About APIs|07. Learn About APIs]] — Architectural blueprints and production operations for Learn About APIs.
- 📂 [[Caching (Core Backend Engineering)|08. Caching]] — Architectural blueprints and production operations for Caching.
- 📂 [[Web Security|09. Web Security]] — Architectural blueprints and production operations for Web Security.
- 📂 [[Testing (Core Backend Engineering)|10. Testing]] — Architectural blueprints and production operations for Testing.
- 📂 [[CI CD|11. CI CD]] — Architectural blueprints and production operations for CI CD.
- 📂 [[Software Design Architecture|12. Software Design Architecture]] — Architectural blueprints and production operations for Software Design Architecture.
- 📂 [[Containerization vs Virtualization|13. Containerization vs Virtualization]] — Architectural blueprints and production operations for Containerization vs Virtualization.
- 📂 [[Web Servers|14. Web Servers]] — Architectural blueprints and production operations for Web Servers.
- 📂 [[Search Engines|15. Search Engines]] — Architectural blueprints and production operations for Search Engines.
- 📂 [[Message Brokers|16. Message Brokers]] — Architectural blueprints and production operations for Message Brokers.
- 📂 [[Real Time Data|17. Real Time Data]] — Architectural blueprints and production operations for Real Time Data.
- 📂 [[Building for Scale|18. Building for Scale]] — Architectural blueprints and production operations for Building for Scale.
- 📂 [[Observability (Core Backend Engineering)|19. Observability]] — Architectural blueprints and production operations for Observability.
- 📂 [[Basic Infrastructure Knowledge|20. Basic Infrastructure Knowledge]] — Architectural blueprints and production operations for Basic Infrastructure Knowledge.
### 2. 📂 [[API Design & Architecture|02. API Design & Architecture]]
- 📂 [[Learn the Basics|01. Learn the Basics]] — Architectural blueprints and production operations for Learn the Basics.
- 📂 [[Different API Styles|02. Different API Styles]] — Architectural blueprints and production operations for Different API Styles.
- 📂 [[Building JSON RESTful APIs|03. Building JSON RESTful APIs]] — Architectural blueprints and production operations for Building JSON RESTful APIs.
- 📂 [[API Authentication and Authorization|04. API Authentication and Authorization]] — Architectural blueprints and production operations for API Authentication and Authorization.
- 📂 [[API Documentation Tools|05. API Documentation Tools]] — Architectural blueprints and production operations for API Documentation Tools.
- 📂 [[API Security|06. API Security]] — Architectural blueprints and production operations for API Security.
- 📂 [[API Performance|07. API Performance]] — Architectural blueprints and production operations for API Performance.
- 📂 [[API Integration Patterns|08. API Integration Patterns]] — Architectural blueprints and production operations for API Integration Patterns.
- 📂 [[API Testing|09. API Testing]] — Architectural blueprints and production operations for API Testing.
- 📂 [[Real Time APIs|10. Real Time APIs]] — Architectural blueprints and production operations for Real Time APIs.
- 📂 [[API Lifecycle Management|11. API Lifecycle Management]] — Architectural blueprints and production operations for API Lifecycle Management.
- 📂 [[Standards and Compliance|12. Standards and Compliance]] — Architectural blueprints and production operations for Standards and Compliance.
### 3. 📂 [[Relational & Distributed Databases|03. Relational & Distributed Databases]]
- 📂 [[Relational Model|01. Relational Model]] — Architectural blueprints and production operations for Relational Model.
- 📂 [[Normalization and Denormalization|02. Normalization and Denormalization]] — Architectural blueprints and production operations for Normalization and Denormalization.
- 📂 [[ER Modeling|03. ER Modeling]] — Architectural blueprints and production operations for ER Modeling.
- 📂 [[SQL DDL DML DQL DCL|04. SQL DDL DML DQL DCL]] — Architectural blueprints and production operations for SQL DDL DML DQL DCL.
- 📂 [[Views|05. Views]] — Architectural blueprints and production operations for Views.
- 📂 [[Stored Procedures and Triggers|06. Stored Procedures and Triggers]] — Architectural blueprints and production operations for Stored Procedures and Triggers.
- 📂 [[Transactions and ACID|07. Transactions and ACID]] — Architectural blueprints and production operations for Transactions and ACID.
- 📂 [[Isolation Levels|08. Isolation Levels]] — Architectural blueprints and production operations for Isolation Levels.
- 📂 [[Locking and Concurrency Control|09. Locking and Concurrency Control]] — Architectural blueprints and production operations for Locking and Concurrency Control.
- 📂 [[MVCC|10. MVCC]] — Architectural blueprints and production operations for MVCC.
- 📂 [[BASE and Eventual Consistency|11. BASE and Eventual Consistency]] — Architectural blueprints and production operations for BASE and Eventual Consistency.
- 📂 [[CAP Theorem|12. CAP Theorem]] — Architectural blueprints and production operations for CAP Theorem.
- 📂 [[PACELC (Relational & Distributed Databases)|13. PACELC]] — Architectural blueprints and production operations for PACELC.
- 📂 [[Indexing (Relational & Distributed Databases)|14. Indexing]] — Architectural blueprints and production operations for Indexing.
- 📂 [[Query Optimization|15. Query Optimization]] — Architectural blueprints and production operations for Query Optimization.
- 📂 [[Replication|16. Replication]] — Architectural blueprints and production operations for Replication.
- 📂 [[Partitioning and Sharding|17. Partitioning and Sharding]] — Architectural blueprints and production operations for Partitioning and Sharding.
- 📂 [[Database Federation|18. Database Federation]] — Architectural blueprints and production operations for Database Federation.
- 📂 [[SQL vs NoSQL (Relational & Distributed Databases)|19. SQL vs NoSQL]] — Architectural blueprints and production operations for SQL vs NoSQL.
- 📂 [[NoSQL Data Models|20. NoSQL Data Models]] — Architectural blueprints and production operations for NoSQL Data Models.
- 📂 [[OLTP vs OLAP and Warehousing|21. OLTP vs OLAP and Warehousing]] — Architectural blueprints and production operations for OLTP vs OLAP and Warehousing.
- 📂 [[Connection Pooling|22. Connection Pooling]] — Architectural blueprints and production operations for Connection Pooling.
- 📂 [[Caching at the DB Layer|23. Caching at the DB Layer]] — Architectural blueprints and production operations for Caching at the DB Layer.
- 📂 [[Backup and Recovery|24. Backup and Recovery]] — Architectural blueprints and production operations for Backup and Recovery.
### 4. 📂 [[Postgresql Mastery & DBA|04. PostgreSQL Mastery & DBA]]
- 📂 [[Introduction (Postgresql Mastery & DBA)|01. Introduction]] — Architectural blueprints and production operations for Introduction.
- 📂 [[Installation and Setup|02. Installation and Setup]] — Architectural blueprints and production operations for Installation and Setup.
- 📂 [[Learn SQL|03. Learn SQL]] — Architectural blueprints and production operations for Learn SQL.
- 📂 [[Configuring|04. Configuring]] — Architectural blueprints and production operations for Configuring.
- 📂 [[Security|05. Security]] — Architectural blueprints and production operations for Security.
- 📂 [[Infrastructure Skills|06. Infrastructure Skills]] — Architectural blueprints and production operations for Infrastructure Skills.
- 📂 [[Application Skills|07. Application Skills]] — Architectural blueprints and production operations for Application Skills.
- 📂 [[Fine Grained Tuning|08. Fine Grained Tuning]] — Architectural blueprints and production operations for Fine Grained Tuning.
- 📂 [[Advanced SQL|09. Advanced SQL]] — Architectural blueprints and production operations for Advanced SQL.
- 📂 [[Troubleshooting Techniques|10. Troubleshooting Techniques]] — Architectural blueprints and production operations for Troubleshooting Techniques.
- 📂 [[SQL Optimization Techniques|11. SQL Optimization Techniques]] — Architectural blueprints and production operations for SQL Optimization Techniques.
- 📂 [[Get Involved in Development|12. Get Involved in Development]] — Architectural blueprints and production operations for Get Involved in Development.
### 5. 📂 [[Redis & in Memory Architectures|05. Redis & In-Memory Architectures]]
- 📂 [[Overview of Redis What Is It|01. Overview of Redis What Is It]] — Architectural blueprints and production operations for Overview of Redis What Is It.
- 📂 [[Getting Started with Redis|02. Getting Started with Redis]] — Architectural blueprints and production operations for Getting Started with Redis.
- 📂 [[First Steps|03. First Steps]] — Architectural blueprints and production operations for First Steps.
- 📂 [[Core Data Structures|04. Core Data Structures]] — Architectural blueprints and production operations for Core Data Structures.
- 📂 [[Working with Redis|05. Working with Redis]] — Architectural blueprints and production operations for Working with Redis.
- 📂 [[Advanced Data Structures|06. Advanced Data Structures]] — Architectural blueprints and production operations for Advanced Data Structures.
- 📂 [[Pub Sub|07. Pub Sub]] — Architectural blueprints and production operations for Pub Sub.
- 📂 [[Transactions|08. Transactions]] — Architectural blueprints and production operations for Transactions.
- 📂 [[Lua Scripting|09. Lua Scripting]] — Architectural blueprints and production operations for Lua Scripting.
- 📂 [[Persistence Options|10. Persistence Options]] — Architectural blueprints and production operations for Persistence Options.
- 📂 [[Replication HA|11. Replication HA]] — Architectural blueprints and production operations for Replication HA.
- 📂 [[Security (Redis & in Memory Architectures)|12. Security]] — Architectural blueprints and production operations for Security.
- 📂 [[Monitoring Optimization|13. Monitoring Optimization]] — Architectural blueprints and production operations for Monitoring Optimization.
- 📂 [[Redis Modules|14. Redis Modules]] — Architectural blueprints and production operations for Redis Modules.
- 📂 [[Managing Redis in Production|15. Managing Redis in Production]] — Architectural blueprints and production operations for Managing Redis in Production.
- 📂 [[Redis Enterprise|16. Redis Enterprise]] — Architectural blueprints and production operations for Redis Enterprise.
### 6. 📂 [[Elasticsearch & Distributed Search|06. Elasticsearch & Distributed Search]]
- 📂 [[Introduction (Elasticsearch & Distributed Search)|01. Introduction]] — Architectural blueprints and production operations for Introduction.
- 📂 [[Core Architecture|02. Core Architecture]] — Architectural blueprints and production operations for Core Architecture.
- 📂 [[Data Modelling|03. Data Modelling]] — Architectural blueprints and production operations for Data Modelling.
- 📂 [[Data Ingestion|04. Data Ingestion]] — Architectural blueprints and production operations for Data Ingestion.
- 📂 [[Search Fundamentals|05. Search Fundamentals]] — Architectural blueprints and production operations for Search Fundamentals.
- 📂 [[How Search Works|06. How Search Works]] — Architectural blueprints and production operations for How Search Works.
- 📂 [[Text Analysis|07. Text Analysis]] — Architectural blueprints and production operations for Text Analysis.
- 📂 [[Aggregations|08. Aggregations]] — Architectural blueprints and production operations for Aggregations.
- 📂 [[Transformations|09. Transformations]] — Architectural blueprints and production operations for Transformations.
- 📂 [[Relevance Tuning|10. Relevance Tuning]] — Architectural blueprints and production operations for Relevance Tuning.
- 📂 [[Production|11. Production]] — Architectural blueprints and production operations for Production.
- 📂 [[Advanced Features|12. Advanced Features]] — Architectural blueprints and production operations for Advanced Features.
### 7. 📂 [[Mongodb & Document Stores|07. MongoDB & Document Stores]]
- 📂 [[Mongodb Basics|01. Mongodb Basics]] — Architectural blueprints and production operations for Mongodb Basics.
- 📂 [[Data Model Data Types|02. Data Model Data Types]] — Architectural blueprints and production operations for Data Model Data Types.
- 📂 [[Collections Methods|03. Collections Methods]] — Architectural blueprints and production operations for Collections Methods.
- 📂 [[Useful Concepts|04. Useful Concepts]] — Architectural blueprints and production operations for Useful Concepts.
- 📂 [[Query Operators|05. Query Operators]] — Architectural blueprints and production operations for Query Operators.
- 📂 [[Performance Optimization|06. Performance Optimization]] — Architectural blueprints and production operations for Performance Optimization.
- 📂 [[Aggregation|07. Aggregation]] — Architectural blueprints and production operations for Aggregation.
- 📂 [[Transactions (Mongodb & Document Stores)|08. Transactions]] — Architectural blueprints and production operations for Transactions.
- 📂 [[Developer Tools|09. Developer Tools]] — Architectural blueprints and production operations for Developer Tools.
- 📂 [[Scaling Mongodb|10. Scaling Mongodb]] — Architectural blueprints and production operations for Scaling Mongodb.
- 📂 [[Mongodb Security|11. Mongodb Security]] — Architectural blueprints and production operations for Mongodb Security.

---

## 🔗 Navigation
- ⬆️ Parent: [[Principal SWE]]
- 🎓 Root: [[Principal SWE]]
