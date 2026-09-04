---
title: DevOps
tags:
  - devops
  - platform-engineering
  - infrastructure
  - cloud
  - kubernetes
  - terraform
  - network-engineering
  - git-and-github
  - principal-swe
parent: "[[Infrastructure & Security]]"
---

# 🚀 DevOps, Cloud Infrastructure & Platform Engineering

Comprehensive, production-grade master architecture covering the complete spectrum of cloud infrastructure, container runtimes, Kubernetes orchestration, Infrastructure as Code (Terraform), Linux systems engineering, edge computing, MLOps, enterprise networking (Network Engineer Roadmap), DevSecOps, and Git & GitHub CI/CD automation across 11 master pillars:

```text
DevOps
│
├── `01. Core DevOps Principles & Automation Tooling`
├── `02. Docker & Container Runtime Internals`
├── `03. Kubernetes & Cloud-Native Orchestration`
├── `04. Terraform & Infrastructure as Code (IaC)`
├── `05. AWS Cloud Platform & Enterprise Infrastructure`
├── `06. Linux Systems Administration & Kernel Engineering`
├── `07. Cloudflare, Edge Computing & CDN Infrastructure`
├── [[MLOps & Machine Learning Operations|08. MLOps & Machine Learning Operations]]
├── `09. Network Engineering & Enterprise Protocols`
├── `10. DevSecOps & Cloud-Native Security Automation`
└── `11. Git & GitHub Version Control & CI-CD Automation`
```

---

## 🚀 Core Knowledge Pillars

### 📂 [[Core DevOps Principles|01. Core DevOps Principles]]
- 📂 `01. Learn a Programming Language for Devops`
- 📂 `02. Operating System and Linux Basics`
- 📂 `03. Terminal Knowledge and Text Processing`
- 📂 `04. Networking Protocols and Socket Basics`
- 📂 `05. Package Managers and Repositories`
- 📂 `06. Linux Process and Service Management`
- 📂 `07. Web Servers and Reverse Proxies`
- 📂 `08. Building for Scale and Load Balancing`
- 📂 `09. Containerization and Container Basics`
- 📂 `10. Container Orchestration Paradigms`
- 📂 `11. Cloud Providers and Hybrid Deployments`
- 📂 `12. Git and Version Control Best Practices`
- 📂 `13. Sre Culture and Incident Response`

### 📂 [[DevOps Automation Tooling|15. DevOps Automation Tooling]]
- 📂 `01. Continuous Integration and CI Tools`
- 📂 `02. Continuous Delivery and Deployment CD`
- 📂 `03. Infrastructure Provisioning Tools`
- 📂 `04. Configuration Management and Ansible`
- 📂 `05. Secret Management in Pipelines`
- 📂 `06. Infrastructure Monitoring and Alerting`
- 📂 `07. Centralized Logging and Log Aggregation`
- 📂 `08. Distributed Tracing and Opentelemetry`
- 📂 `09. Disaster Recovery and Backup Automation`

### 📂 [[Docker|02. Docker]]
- 📂 `01. Docker Engine Architecture and Oci Standards`
- 📂 `02. Dockerfile Optimization and Multi Stage Builds`
- 📂 `03. Container Networking Models Bridge, Host, Overlay`
- 📂 `04. Storage Volumes, Bind Mounts, and Tmpfs`
- 📂 `05. Docker Compose for Multi Container Development`
- 📂 `06. Rootless Docker and Container Security Hardening`
- 📂 `07. Container Image Registries and Scanning`
- 📂 `08. Container Image Signing with Cosign and Notary`
- 📂 `09. Docker Cli Power Tools and Container Debugging`
- 📂 `10. Container Lifecycle Management and Clean Up`
- 📂 `11. Production Container Troubleshooting Runbook`

### 📂 [[Container Runtime Internals|18. Container Runtime Internals]]
- 📂 `01. Linux Namespaces and Process Isolation`
- 📂 `02. Control Groups cgroups V1 and V2 Resource Limits`
- 📂 `03. Union Filesystems and Overlayfs Storage Drivers`
- 📂 `04. Alternative Container Runtimes Podman and Buildah`
- 📂 `05. Microvms and Sandbox Runtimes Firecracker and Gvisor`

### 📂 [[Kubernetes|03. Kubernetes]]
- 📂 `01. Kubernetes Control Plane and Worker Node Architecture`
- 📂 `02. Pods, Replicasets, and Deployments`
- 📂 `03. Statefulsets and Stateful Cluster Orchestration`
- 📂 `04. Daemonsets and Job Cronjob Workloads`
- 📂 `05. Kubernetes Networking and Cni Plugins`
- 📂 `06. Services, Kube Proxy, and Coredns`
- 📂 `07. Ingress Controllers and Gateway Api`
- 📂 `08. Configmaps, Secrets, and External Secrets Operator`
- 📂 `09. Persistent Volumes, Pvcs, and CSI Storage Drivers`
- 📂 `10. Autoscaling Hpa, Vpa, and Cluster Autoscaler`
- 📂 `11. Kubernetes Security, Rbac, and Admission Controllers`

### 📂 [[Cloud-Native Orchestration|12. Cloud-Native Orchestration]]
- 📂 `01. Custom Resource Definitions CRDs and Kubernetes Operators`
- 📂 `02. Package Management with Helm and Kustomize`
- 📂 `03. Service Mesh Architecture Istio and Linkerd`
- 📂 `04. Gitops Continuous Delivery with Argocd and Flux`

### 📂 [[Terraform|04. Terraform]]
- 📂 `01. Terraform Architecture and Core Workflow`
- 📂 `02. Hashicorp Configuration Language HCL Syntax`
- 📂 `03. State Management and Remote Backends`
- 📂 `04. State Manipulation and Refactoring`
- 📂 `05. Providers and Resource Definitions`
- 📂 `06. Input Variables, Outputs, and Validation`
- 📂 `07. Modular Terraform Blueprints and Best Practices`
- 📂 `08. Dynamic Blocks and Advanced Expressions`
- 📂 `09. Built in Functions and Template Generation`
- 📂 `10. Terraform Workspaces and Multi Environment Topologies`
- 📂 `11. Resource Lifecycles and Provisioners`
- 📂 `12. Terragrunt for Dry Infrastructure Architectures`
- 📂 `13. Terraform Cloud, Enterprise, and Spacelift`
- 📂 `14. Automated Testing with Terratest`
- 📂 `15. Production Terraform Troubleshooting and State Recovery`

### 📂 [[Infrastructure as Code|14. Infrastructure as Code]]
- 📂 `01. Opentofu Fork and Open Source Ecosystem`
- 📂 `02. Policy As Code with Sentinel and OPA Rego`
- 📂 `03. Static Security Analysis with Checkov and Tfsec`
- 📂 `04. Drift Detection and Continuous Reconciliation`
- 📂 `05. Multi Cloud and Multi Account Landing Zones`
- 📂 `06. Infrastructure Cost Estimation with Infracost`
- 📂 `07. Zero Downtime Infrastructure Refactoring`

### 📂 [[AWS Cloud Platform|05. AWS Cloud Platform]]
- 📂 `01. Elastic Compute Cloud EC2 and Auto Scaling Groups`
- 📂 `02. Simple Storage Service S3 Architecture and Security`
- 📂 `03. Elastic Block Store Ebs and Elastic File System Efs`
- 📂 `04. Relational Database Service RDS and Aurora Architecture`
- 📂 `05. Dynamodb Distributed Nosql Database`
- 📂 `06. AWS Lambda and Serverless Compute Architectures`
- 📂 `07. Elastic Container Service ECS and Fargate`
- 📂 `08. Elastic Kubernetes Service EKS Production Clusters`
- 📂 `09. Identity and Access Management IAM Deep Dive`
- 📂 `10. Key Management Service Kms and Secrets Manager`
- 📂 `11. Cloudwatch, Cloudtrail, and AWS Observability`
- 📂 `12. AWS Security Services Guardduty, Waf, and Shield`
- 📂 `13. AWS Well Architected Framework and Finops Cost Management`
- 📂 `14. Cloudformation and AWS Cloud Development Kit Cdk`

### 📂 [[AWS Enterprise Infrastructure|20. AWS Enterprise Infrastructure]]
- 📂 `01. AWS Global Infrastructure and Multi Region Design`
- 📂 `02. Virtual Private Cloud VPC and Network Topologies`
- 📂 `03. AWS Transit Gateway and Hybrid Cloud Connectivity`
- 📂 `04. Elastic Load Balancing Elb and Application Load Balancer`
- 📂 `05. Route 53 DNS and Global Traffic Management`

### 📂 [[Linux Systems Administration|06. Linux Systems Administration]]
- 📂 `01. Linux Directory Hierarchy and Filesystem Standards`
- 📂 `02. Linux Core Commands and Text Manipulation`
- 📂 `03. Linux User, Group, and Permission Models`
- 📂 `04. Posix Access Control Lists ACLs and File Attributes`
- 📂 `05. Process Lifecycle, Signals, and Daemon Management`
- 📂 `06. Systemd Architecture, Targets, and Journald`
- 📂 `07. Storage Management Lvm, Raid, and Partitioning`
- 📂 `08. Linux Networking Tools and Socket Inspection`
- 📂 `09. SSH Daemon Hardening and Key Management`
- 📂 `10. Cron, Anacron, and Systemd Timers`
- 📂 `11. Linux Logging Architecture Syslog and Rsyslog`
- 📂 `12. Linux Backup and Archiving Utilities`
- 📂 `13. Linux Systems Performance Troubleshooting Runbook`

### 📂 [[Kernel Engineering|19. Kernel Engineering]]
- 📂 `01. Linux Memory Architecture and Swap Management`
- 📂 `02. Linux Filesystem Internals Ext4 and Xfs`
- 📂 `03. Linux Firewalling Netfilter, Iptables, and Nftables`
- 📂 `04. Linux Kernel Performance Tuning with Sysctl`
- 📂 `05. Linux Security Modules Apparmor and Selinux`

### 📂 [[Cloudflare Edge Computing|07. Cloudflare Edge Computing]]
- 📂 `01. Cloudflare Architecture and Anycast Global Network`
- 📂 `02. Cloudflare Workers and V8 Isolate Architecture`
- 📂 `03. Cloudflare Web Application Firewall WAF and Rulesets`
- 📂 `04. Ddos Mitigation and Rate Limiting at the Edge`
- 📂 `05. Cloudflare Bot Management and Turnstile`
- 📂 `06. Cloudflare Zero Trust and Access Gateway`
- 📂 `07. Cloudflare Observability, Analytics, and Logpush`

### 📂 [[CDN Infrastructure|17. CDN Infrastructure]]
- 📂 `01. Edge DNS and Managed DNS Records`
- 📂 `02. CDN Caching Strategies and Edge Purging`
- 📂 `03. Edge Storage Kv, D1, and R2 Object Storage`
- 📂 `04. Page Rules, Transform Rules, and Url Rewrites`
- 📂 `05. SSL TLS Encryption Modes and Origin Certificates`

### 📂 [[MLOps & Machine Learning Operations|08. MLOps & Machine Learning Operations]]
- 📂 `01. MLOps Architecture and ML Lifecycle Standards`
- 📂 `02. Data Versioning and Lineage with DVC`
- 📂 `03. Feature Stores Architecture Feast and Hopsworks`
- 📂 `04. Experiment Tracking and Model Registry with Mlflow`
- 📂 `05. Distributed Training Pipelines and Kubeflow`
- 📂 `06. High Throughput Model Serving Triton and Vllm`
- 📂 `07. Model Packaging and Containerization Bentoml`
- 📂 `08. Model Monitoring, Data Drift, and Concept Drift`
- 📂 `09. Model Deployment Strategies Canary and Shadow Deployments`
- 📂 `10. GPU Cluster Orchestration and Scheduling in Kubernetes`
- 📂 `11. Production Llm Operations and Serving Llmops`
### 📂 [[Network Engineering|09. Network Engineering]]
- 📂 `01. Network Security Firewalls, Ids Ips, and VPN Tunnels`
- 📂 `02. Layer 4 vs Layer 7 Load Balancing and Traffic Routing`
- 📂 `03. Network Automation with Python, Netmiko, and Ansible`
- 📂 `04. Network Observability, Packet Analysis, and Troubleshooting`
- 📂 `05. Software Defined Networking Sdn and Sd WAN Architecture`
- 📂 `06. Cloud Virtual Private Clouds VPC and Hybrid Interconnects`
- 📂 `07. High Performance Kernel Bypass Networking Dpdk and Rdma`

### 📂 [[Enterprise Protocols|16. Enterprise Protocols]]
- 📂 `01. Osi Model and Tcp Ip Protocol Suite Architecture`
- 📂 `02. Ip Addressing, Cidr Subnetting, and Ipv6 Migration`
- 📂 `03. Enterprise Routing Protocols Bgp, Ospf, and Rip`
- 📂 `04. Layer 2 Switching, Vlans, and Spanning Tree Protocol`
- 📂 `05. Enterprise Network Services DNS Anycast, Dhcp, and Ntp`

### 📂 [[DevSecOps|10. DevSecOps]]
- 📂 `01. Shift Left Security Culture and Devsecops Frameworks`
- 📂 `02. Automated SAST and DAST in CI CD Pipelines`
- 📂 `03. Infrastructure As Code Security Linting and Guardrails`
- 📂 `04. Software Supply Chain Security Slsa, Cosign, and SBOM`
- 📂 `05. Continuous Compliance, Audit Trails, and Devsecops Runbooks`

### 📂 [[Cloud-Native Security Automation|13. Cloud-Native Security Automation]]
- 📂 `01. Container Image Security and Vulnerability Scanning`
- 📂 `02. Dynamic Secrets Injection and Ephemeral Credentials`
- 📂 `03. Cloud Native Runtime Security with Falco and Ebpf`
- 📂 `04. Policy As Code with Open Policy Agent OPA and Kyverno`
- 📂 `05. Cloud Security Posture Management Cspm Automation`

### 📂 [[Git Version Control|11. Git Version Control]]
- 📂 [[Git Plumbing, Internals & Core Mechanics|01. 01. Git Plumbing, Internals & Core Mechanics]]
- 📂 [[Branching Strategies & Merge Topologies|02. 02. Branching Strategies & Merge Topologies]]
- 📂 [[Advanced Rebasing, Cherry-Picking & History Rewriting|03. 03. Advanced Rebasing, Cherry-Picking & History Rewriting]]
- 📂 [[Conflict Resolution & Interactive Debugging|04. 04. Conflict Resolution & Interactive Debugging]]

## 🔗 Navigation
- ⬆️ Parent: [[Infrastructure & Security]]
- 🏛️ Software Architecture: `Architecture`
- 💻 Computer Science Foundations: `Computer Science`
- 🛡️ Cyber Security: `Cyber Security`

---

## 🗂️ Topics

- [[AWS Cloud Platform]]
- [[AWS Enterprise Infrastructure]]
- [[CDN Infrastructure]]
- [[Cloud-Native Orchestration]]
- [[Cloud-Native Security Automation]]
- [[Cloudflare Edge Computing]]
- [[Container Runtime Internals]]
- [[Core DevOps Principles]]
- [[DevOps Automation Tooling]]
- [[DevSecOps]]
- [[Docker]]
- [[Enterprise Protocols]]
- [[Git Version Control]]
- [[GitHub and CI-CD Automation]]
- [[Infrastructure as Code]]
- [[Kernel Engineering]]
- [[Kubernetes]]
- [[Linux Systems Administration]]
- [[MLOps & Machine Learning Operations]]
- [[Network Engineering]]
- [[Terraform]]
