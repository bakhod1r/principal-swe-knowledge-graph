---
title: Networking (Computer Science)
tags:
  - computer-science
  - networking-(computer-science)
  - principal-swe
parent: "[[Computer Science]]"
---

# 🏛️ Networking (Computer Science) (Foundations & Systems Architecture)

TCP transport internals, congestion control algorithms (BBR, CUBIC), head-of-line blocking, TLS 1.3 cryptographic handshakes, QUIC/HTTP-3 UDP multiplexing, L4/L7 load balancing, and network namespaces.

```text
Networking (Computer Science)
│
├── [[TCP Congestion Control and Flow Control|01. TCP Congestion and Flow Control]]
├── [[TCP State Machine and Head of Line Blocking|02. TCP Internals and Hol Blocking]]
├── [[TLS 1.3 Cryptographic Protocol and Handshake|03. TLS 1.3 Handshake Internals]]
├── [[QUIC Transport Protocol and HTTP 3|04. QUIC and HTTP 3]]
├── [[Layer 4 vs Layer 7 Load Balancing Architecture|05. L4 vs L7 Load Balancing]]
├── [[Network Latency Numbers Every Engineer Should Know|06. Network Latency Numbers]]
├── [[DNS Resolution Internals and BGP Anycast|07. DNS Internals and Anycast]]
└── [[Network Namespaces and Overlay Networking|08. Network Namespaces and Overlays]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[TCP Congestion Control and Flow Control|01. TCP Congestion and Flow Control]] — Sliding window flow control, slow start, AIMD, loss-based CUBIC, and model-based BBR congestion control.
- 📂 [[TCP State Machine and Head of Line Blocking|02. TCP Internals and Hol Blocking]] — 3-way handshake, TIME_WAIT socket states, packet reordering, TCP buffer autotuning, and HOL blocking.
- 📂 [[TLS 1.3 Cryptographic Protocol and Handshake|03. TLS 1.3 Handshake Internals]] — 1-RTT handshake, Diffie-Hellman ephemeral key exchange (ECDHE), 0-RTT early data, and perfect forward secrecy.
- 📂 [[QUIC Transport Protocol and HTTP 3|04. QUIC and HTTP 3]] — UDP-based multiplexed transport, connection migration (CID), zero HOL blocking, and encrypted transport headers.
- 📂 [[Layer 4 vs Layer 7 Load Balancing Architecture|05. L4 vs L7 Load Balancing]] — Direct Server Return (DSR), Consistent Hashing, IPVS, Envoy proxy multiplexing, and connection draining.
- 📂 [[Network Latency Numbers Every Engineer Should Know|06. Network Latency Numbers]] — Intra-datacenter RTT (<0.5ms), cross-region speed of light fiber latency (70-100ms), and packet serialization delay.
- 📂 [[DNS Resolution Internals and BGP Anycast|07. DNS Internals and Anycast]] — Recursive resolution, DNS caching TTL, authoritative nameservers, EDNS Client Subnet, and Anycast routing.
- 📂 [[Network Namespaces and Overlay Networking|08. Network Namespaces and Overlays]] — Linux netns, veth pairs, iptables NAT/MASQUERADE, VXLAN packet encapsulation, and Calico/Cilium CNI.

---

## 🔗 References
- ⬆️ Parent: [[Computer Science]]

