---
title: Data Engineering
tags:
  - ai-and-machine-learning
  - ai-engineering
  - data-engineering,-etl-pipelines-and-lakehouse-architecture
  - principal-swe
parent: "[[AI & Intelligent Systems]]"
---

# 🤖 Data Engineering, ETL Pipelines & Lakehouse Architecture

Modern enterprise data platforms: Batch vs streaming, Apache Spark, Apache Flink, Lakehouse formats (Delta Lake, Apache Iceberg), orchestration (Airflow, Dagster), data modeling (Kimball), dbt transformations, and data quality.

```text
Data Engineering, ETL Pipelines & Lakehouse Architecture
│
├── [[The Modern Data Stack (mds), ETL vs Elt, and Cloud Data Warehouses|01. Modern Data Stack Architecture and Paradigms]]
├── `02. Batch vs Streaming Data Processing Paradigms`
├── `03. Apache Spark Distributed Computation Internals`
├── `04. Stateful Stream Processing with Apache Flink`
├── [[Open Lakehouse Table Formats: Apache Iceberg, Delta Lake, and Apache Hudi|05. Lakehouse Table Formats Delta Lake, Iceberg, and Hudi]]
├── `06. Data Pipeline Orchestration Airflow, Dagster, and Prefect`
├── [[Dimensional Data Modeling: Kimball Star Schemas, Snowflake, and Data Vault|07. Dimensional Data Modeling and Star Schemas]]
├── `08. Data Transformation and Modeling with dbt`
├── [[Data Quality Verification, Great Expectations, and Anomaly Detection|09. Data Quality, Anomaly Detection, and Great Expectations]]
└── [[Data Governance, End to End Lineage (openlineage), and Catalogs (datahub)|10. Data Governance, Lineage, and Metadata Catalogs]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[The Modern Data Stack (mds), ETL vs Elt, and Cloud Data Warehouses|01. Modern Data Stack Architecture and Paradigms]] — Evolution from on-premise ETL to cloud-native ELT (Snowflake, BigQuery, Databricks), decoupling compute from storage, and modular data architectures.
- 📂 `02. Batch vs Streaming Data Processing Paradigms` — Windowing strategies (Tumbling, Sliding, Session), event time vs processing time, watermarks for late-arriving data, and Kappa streaming architecture.
- 📂 `03. Apache Spark Distributed Computation Internals` — Distributed in-memory computation, Catalyst query optimization, Tungsten binary memory management, shuffle partitioning, and memory spill debugging.
- 📂 `04. Stateful Stream Processing with Apache Flink` — Distributed asynchronous snapshotting (Chandy-Lamport), state backends (RocksDB), keyed streams, complex event processing (CEP), and exactly-once guarantees.
- 📂 [[Open Lakehouse Table Formats: Apache Iceberg, Delta Lake, and Apache Hudi|05. Lakehouse Table Formats Delta Lake, Iceberg, and Hudi]] — ACID transactions on cloud object storage (S3), metadata layers, time travel and snapshot isolation, partition evolution without rewrites, and file compaction.
- 📂 `06. Data Pipeline Orchestration Airflow, Dagster, and Prefect` — Directed Acyclic Graphs (DAGs), asset-based orchestration in Dagster, declarative dependencies, dynamic retries, backfilling historic data, and SLA alerts.
- 📂 [[Dimensional Data Modeling: Kimball Star Schemas, Snowflake, and Data Vault|07. Dimensional Data Modeling and Star Schemas]] — Fact tables (transactional, periodic snapshot, accumulating), dimension tables, Slowly Changing Dimensions (SCD Type 1, 2, 3), and Data Vault 2.0.
- 📂 `08. Data Transformation and Modeling with dbt` — Writing modular SQL transformations, Jinja macros, ephemeral/view/table materializations, automated schema testing, and documentation generation in dbt.
- 📂 [[Data Quality Verification, Great Expectations, and Anomaly Detection|09. Data Quality, Anomaly Detection, and Great Expectations]] — Automated data validation assertions (null checks, uniqueness, distribution bounds), CI data quality gates, anomaly detection in pipeline runs, and data SLAs.
- 📂 [[Data Governance, End to End Lineage (openlineage), and Catalogs (datahub)|10. Data Governance, Lineage, and Metadata Catalogs]] — Tracking automated data lineage across Spark/dbt/Airflow, data discovery catalogs (DataHub, Amundsen), column-level access control, and GDPR/CCPA compliance.

---

## 🔗 References
- ⬆️ Parent: `AI & Machine Learning`

