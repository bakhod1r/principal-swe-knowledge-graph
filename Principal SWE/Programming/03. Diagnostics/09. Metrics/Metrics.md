---
title: Metrics
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Metrics

Prometheus/OTel data models, cardinality explosion mitigation, percentile algorithms (HDR, DDSketch), and push vs pull scraping.

```text
Metrics
│
├── [[Metrics Data Models (Counters, Gauges, Histograms, Summaries)]]
├── [[High-Cardinality Metric Explosion Mitigation]]
├── [[Percentile Calculation Algorithms (HDR Histogram, DDSketch, t-digest)]]
└── [[Push vs Pull Metric Scraping Architecture]]
```

---

## 🗂️ Topics

- [[Metrics Data Models (Counters, Gauges, Histograms, Summaries)]] — Prometheus and OpenTelemetry metric types, aggregations, and mathematical properties.
- [[High-Cardinality Metric Explosion Mitigation]] — Managing multi-dimensional label cardinality and avoiding catastrophic memory leaks in time-series DBs.
- [[Percentile Calculation Algorithms (HDR Histogram, DDSketch, t-digest)]] — Accurate P95, P99, and P99.9 latency estimations across distributed clusters without raw data storage.
- [[Push vs Pull Metric Scraping Architecture]] — Evaluating Prometheus pull scraping vs OpenTelemetry push protocols under massive cloud scale.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]

