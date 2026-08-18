---
title: Frontend Performance Best Practices
tags:
  - best-practices
  - engineering-excellence
  - frontend-performance-best-practices
  - principal-swe
parent: "[[Best Practices]]"
---

# 🏛️ Frontend Performance Best Practices

Ultra-fast web client engineering: Core Web Vitals optimization (LCP, INP, CLS), asset payload compression, JavaScript code splitting, DOM rendering acceleration, and client state caching.

```text
Frontend Performance Best Practices
│
├── [[Core Web Vitals Optimization (lcp, Inp, Cls)|01. Core Web Vitals Optimization]]
├── [[Frontend Network and Asset Delivery Optimization|02. Network and Asset Delivery]]
├── [[Javascript Bundle Optimization and Code Splitting|03. Javascript Bundle and Code Splitting]]
├── [[DOM and CSS Rendering Performance|04. DOM and CSS Rendering Performance]]
└── [[Client Side Caching and Offline Resilience|05. Client Caching and Offline Resilience]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Core Web Vitals Optimization (lcp, Inp, Cls)|01. Core Web Vitals Optimization]] — Optimizing Largest Contentful Paint (<2.5s), Interaction to Next Paint (<200ms), and Cumulative Layout Shift (<0.1).
- 📂 [[Frontend Network and Asset Delivery Optimization|02. Network and Asset Delivery]] — Next-gen image formats (AVIF, WebP), font preloading, CDN edge routing, HTTP/3, and minimizing TTFB (<1.3s).
- 📂 [[Javascript Bundle Optimization and Code Splitting|03. Javascript Bundle and Code Splitting]] — Route-based code splitting, dynamic imports, tree shaking unused exports, and eliminating render-blocking scripts.
- 📂 [[DOM and CSS Rendering Performance|04. DOM and CSS Rendering Performance]] — GPU layer acceleration, avoiding layout thrashing/forced reflows, virtual scrolling for giant lists, and CSS containment.
- 📂 [[Client Side Caching and Offline Resilience|05. Client Caching and Offline Resilience]] — Service Worker caching strategies, Cache-Control headers, IndexedDB client storage, and SWR/TanStack query deduplication.

---

## 🔗 References
- ⬆️ Parent: [[Best Practices]]
- 🎓 Root: [[Principal SWE]]
