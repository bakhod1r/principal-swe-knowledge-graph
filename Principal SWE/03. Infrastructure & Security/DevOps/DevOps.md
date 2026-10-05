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

Comprehensive, production-grade master architecture covering the complete spectrum of cloud infrastructure, container runtimes, Kubernetes orchestration, Infrastructure as Code (Terraform), Linux systems engineering, edge computing, MLOps, enterprise networking (Network Engineer Roadmap), DevSecOps, and Git & GitHub CI/CD automation across 21 master pillars:

```text
DevOps
│
├── [[Linux Systems Administration|01. Linux Systems Administration]]
├── [[Network Engineering|02. Network Engineering]]
├── [[Git Version Control|03. Git Version Control]]
├── [[GitHub and CI-CD Automation|04. CI-CD Automation]]
├── [[Docker|05. Docker]]
├── [[Docker Swarm|06. Docker Swarm]]
├── [[Container Runtime Internals|07. Container Runtime Internals]]
├── [[Kubernetes|08. Kubernetes]]
├── [[Cloud-Native Orchestration|09. Cloud-Native Orchestration]]
├── [[Infrastructure as Code|10. Infrastructure as Code]]
├── [[Terraform|11. Terraform]]
├── [[DevOps Automation Tooling|12. DevOps Automation Tooling]]
├── [[AWS Cloud Platform|13. AWS Cloud Platform]]
├── [[AWS Enterprise Infrastructure|14. AWS Enterprise Infrastructure]]
├── [[Cloudflare Edge Computing|15. Cloudflare Edge Computing]]
├── [[CDN Infrastructure|16. CDN Infrastructure]]
├── [[Enterprise Protocols|17. Enterprise Protocols]]
├── [[DevSecOps|18. DevSecOps]]
├── [[Cloud-Native Security Automation|19. Cloud-Native Security Automation]]
├── [[MLOps & Machine Learning Operations|20. MLOps & Machine Learning Operations]]
└── [[Kernel Engineering|21. Kernel Engineering]]
```

---

## 🚀 Core Knowledge Pillars

### 📂 [[Linux Systems Administration|01. Linux Systems Administration]]
- 📂 [[Linux Filesystem Hierarchy Standard (fhs) and Special Mounts|01. Linux Directory Hierarchy and Filesystem Standards]]
- 📂 [[Linux Core Commands, Text Manipulation (grep, Sed, Awk, Cut, Tr)|02. Linux Core Commands and Text Manipulation]]
- 📂 [[Linux User and Group Management, Posix Permissions, and Suid Sgid|03. Linux User, Group, and Permission Models]]
- 📂 [[Posix Access Control Lists (acls) and Extended File Attributes|04. Posix Access Control Lists ACLs and File Attributes]]
- 📂 [[Linux Process Lifecycle, Priority (nice), and Signal Handling|05. Process Lifecycle, Signals, and Daemon Management]]
- 📂 [[Systemd Architecture, Unit Files, Targets, and Journald Logs|06. Systemd Architecture, Targets, and Journald]]
- 📂 [[Linux Storage Management: Lvm, Software RAID (mdadm), and Partitions|07. Storage Management Lvm, Raid, and Partitioning]]
- 📂 [[Linux Networking Cli Tools, Socket Inspection (ip, Ss, Netstat, Ethtool)|08. Linux Networking Tools and Socket Inspection]]
- 📂 [[Openssh Server Hardening, Key Authentication, and Bastions|09. SSH Daemon Hardening and Key Management]]
- 📂 [[Automated Task Scheduling: Cron, Anacron, and Systemd Timers|10. Cron, Anacron, and Systemd Timers]]
- 📂 [[Linux Logging Infrastructure: Syslog, Rsyslog, and Logrotate|11. Linux Logging Architecture Syslog and Rsyslog]]
- 📂 [[Linux Backup Utilities: Tar, Gzip, Rsync, and Rclone|12. Linux Backup and Archiving Utilities]]
- 📂 [[Linux Systems Performance Troubleshooting: the Use and Red Methods|13. Linux Systems Performance Troubleshooting Runbook]]
- 📂 [[Linux Package Managers, Repositories, and Artifact Storage|14. Package Managers and Repositories]]

### 📂 [[Network Engineering|02. Network Engineering]]
- 📂 [[Network Security Architecture - Next Gen Firewalls, Ids Ips, and VPNs|01. Network Security Firewalls, Ids Ips, and VPN Tunnels]]
- 📂 [[Network Load Balancing - Layer 4 Direct Server Return (dsr) vs Layer 7 Proxies|02. Layer 4 vs Layer 7 Load Balancing and Traffic Routing]]
- 📂 [[Network Automation Engineering - Netmiko, Napalm, and Ansible Playbooks|03. Network Automation with Python, Netmiko, and Ansible]]
- 📂 [[Network Observability - Wireshark Packet Analysis, Tcpdump, and Ebpf|04. Network Observability, Packet Analysis, and Troubleshooting]]
- 📂 [[Software Defined Networking (sdn) and Sd WAN Architecture|05. Software Defined Networking Sdn and Sd WAN Architecture]]
- 📂 [[Cloud Networking Architecture - Aws Vpc, Peering, Transit Gateway, and Direct Connect|06. Cloud Virtual Private Clouds VPC and Hybrid Interconnects]]
- 📂 [[Kernel Bypass Networking - Data Plane Development Kit (dpdk) and Rdma|07. High Performance Kernel Bypass Networking Dpdk and Rdma]]
- 📂 [[Web Servers, Reverse Proxies, and Edge Ingress (nginx, Caddy)|08. Web Servers and Reverse Proxies]]

### 📂 [[Git Version Control|03. Git Version Control]]
- 📂 [[Git Plumbing, Internals & Core Mechanics|01. Git Plumbing, Internals & Core Mechanics]]
- 📂 [[Branching Strategies & Merge Topologies|02. Branching Strategies & Merge Topologies]]
- 📂 [[Advanced Rebasing, Cherry-Picking & History Rewriting|03. Advanced Rebasing, Cherry-Picking & History Rewriting]]
- 📂 [[Conflict Resolution & Interactive Debugging|04. Conflict Resolution & Interactive Debugging]]
- 📂 [[Git and Version Control Standards for Infrastructure (GitOps)|05. Git and Version Control Best Practices]]

### 📂 [[GitHub and CI-CD Automation|04. CI-CD Automation]]
- 📂 [[GitHub Enterprise Workflows|01. GitHub Enterprise Workflows]]
- 📂 [[GitHub Actions CI-CD & Workflow Automation|02. GitHub Actions CI-CD & Workflow Automation]]
- 📂 [[Repository Security, Secrets & Supply Chain Hardening|03. Repository Security, Secrets & Supply Chain Hardening]]
- 📂 [[GitOps, Enterprise CLI & Automation Tooling|04. GitOps, Enterprise CLI & Automation Tooling]]
- 📂 [[PR Engineering|05. PR Engineering]]

### 📂 [[Docker|05. Docker]]
- 📂 [[Docker Engine Architecture, Containerd, Runc, and Oci Standards|01. Docker Engine Architecture and Oci Standards]]
- 📂 [[Dockerfile Best Practices, Layer Caching, and Multi Stage Builds|02. Dockerfile Optimization and Multi Stage Builds]]
- 📂 [[Docker Networking Models (bridge, Host, None, Macvlan, Overlay)|03. Container Networking Models Bridge, Host, Overlay]]
- 📂 [[Docker Storage Options (volumes, Bind Mounts, and Tmpfs Mounts)|04. Storage Volumes, Bind Mounts, and Tmpfs]]
- 📂 [[Docker Compose Specification and Local Microservice Topology|05. Docker Compose for Multi Container Development]]
- 📂 [[Rootless Docker, User Namespaces, and Capability Dropping|06. Rootless Docker and Container Security Hardening]]
- 📂 [[Container Image Registries, Harbor, and Vulnerability Scanning|07. Container Image Registries and Scanning]]
- 📂 [[Container Image Signing, Cryptographic Provenance, and Cosign|08. Container Image Signing with Cosign and Notary]]
- 📂 [[Docker Cli Debugging, Container Exec, and Resource Profiling|09. Docker Cli Power Tools and Container Debugging]]
- 📂 [[Container Lifecycle Automation, Pruning, and Garbage Collection|10. Container Lifecycle Management and Clean Up]]
- 📂 [[Production Container Troubleshooting and Failure Mode Analysis|11. Production Container Troubleshooting Runbook]]

### 📂 [[Docker Swarm|06. Docker Swarm]]
- 📂 [[Swarm Architecture, Managers, Workers, and Raft Consensus|01. Swarm Architecture, Managers, Workers, and Raft Consensus]]
- 📂 [[Services, Tasks, and Scheduling Constraints|02. Services, Tasks, and Scheduling Constraints]]
- 📂 [[Overlay Networking, Ingress Routing Mesh, and Service Discovery|03. Overlay Networking, Ingress Routing Mesh, and Service Discovery]]
- 📂 [[Secrets and Configs Management|04. Secrets and Configs Management]]
- 📂 [[Stack Deployments with Compose Files|05. Stack Deployments with Compose Files]]
- 📂 [[Rolling Updates, Rollbacks, and Health Checks|06. Rolling Updates, Rollbacks, and Health Checks]]
- 📂 [[Cluster Security, Mutual TLS, and Autolock|07. Cluster Security, Mutual TLS, and Autolock]]
- 📂 [[High Availability, Backup, and Disaster Recovery|08. High Availability, Backup, and Disaster Recovery]]
- 📂 [[Swarm vs Kubernetes and Migration Strategy|09. Swarm vs Kubernetes and Migration Strategy]]

### 📂 [[Container Runtime Internals|07. Container Runtime Internals]]
- 📂 [[Linux Namespaces (pid, Net, Mnt, Ipc, Uts, User) Deep Dive|01. Linux Namespaces and Process Isolation]]
- 📂 [[Control Groups (cgroups V1 and V2) Cpu and Memory Controls|02. Control Groups cgroups V1 and V2 Resource Limits]]
- 📂 [[Union Filesystems, Copy on Write (cow), and Overlayfs Internals|03. Union Filesystems and Overlayfs Storage Drivers]]
- 📂 [[Alternative Container Toolchains (podman, Buildah, Skopeo)|04. Alternative Container Runtimes Podman and Buildah]]
- 📂 [[Microvms and Sandboxed Container Runtimes (firecracker, Gvisor, Kata)|05. Microvms and Sandbox Runtimes Firecracker and Gvisor]]

### 📂 [[Kubernetes|08. Kubernetes]]
- 📂 [[Kubernetes Control Plane (api Server, Etcd, Scheduler, Controllers) and Kubelet|01. Kubernetes Control Plane and Worker Node Architecture]]
- 📂 [[Kubernetes Workloads: Pods, Replicasets, and Deployments|02. Pods, Replicasets, and Deployments]]
- 📂 [[Statefulsets, Headless Services, and Stable Network Identifiers|03. Statefulsets and Stateful Cluster Orchestration]]
- 📂 [[Daemonsets, Batch Jobs, and Scheduled Cronjobs|04. Daemonsets and Job Cronjob Workloads]]
- 📂 [[Kubernetes Networking Model, Pod to Pod Communication, and Cni Plugins|05. Kubernetes Networking and Cni Plugins]]
- 📂 [[Kubernetes Services (clusterip, Nodeport, Loadbalancer) and Coredns|06. Services, Kube Proxy, and Coredns]]
- 📂 [[Kubernetes Ingress Controllers and Modern Gateway Api|07. Ingress Controllers and Gateway Api]]
- 📂 [[Kubernetes Configmaps, Secrets Management, and External Secrets Operator|08. Configmaps, Secrets, and External Secrets Operator]]
- 📂 [[Persistent Volumes (pv), Pvcs, Storageclasses, and CSI Drivers|09. Persistent Volumes, Pvcs, and CSI Storage Drivers]]
- 📂 [[Kubernetes Autoscaling: Horizontal Pod Autoscaler, Vpa, and Karpenter|10. Autoscaling Hpa, Vpa, and Cluster Autoscaler]]
- 📂 [[Kubernetes Security: Rbac, Pod Security Standards, and OPA Kyverno|11. Kubernetes Security, Rbac, and Admission Controllers]]

### 📂 [[Cloud-Native Orchestration|09. Cloud-Native Orchestration]]
- 📂 [[Custom Resource Definitions (crds) and Kubernetes Operator Sdk|01. Custom Resource Definitions CRDs and Kubernetes Operators]]
- 📂 [[Kubernetes Package Management: Helm Charts and Kustomize Overlays|02. Package Management with Helm and Kustomize]]
- 📂 [[Service Mesh Architecture, Mtls, Traffic Shifting, and Istio|03. Service Mesh Architecture Istio and Linkerd]]
- 📂 [[Gitops Continuous Delivery with Argocd, Flux, and Declarative Sync|04. Gitops Continuous Delivery with Argocd and Flux]]

### 📂 [[Infrastructure as Code|10. Infrastructure as Code]]
- 📂 [[Opentofu Architecture, State Encryption, and Migration From Terraform|01. Opentofu Fork and Open Source Ecosystem]]
- 📂 [[Policy As Code for Infrastructure (sentinel and OPA Rego)|02. Policy As Code with Sentinel and OPA Rego]]
- 📂 [[Static Security Analysis and Linting for Terraform (tfsec, Checkov, Tflint)|03. Static Security Analysis with Checkov and Tfsec]]
- 📂 [[Terraform Drift Detection, Scheduled Plans, and Automated Remediation|04. Drift Detection and Continuous Reconciliation]]
- 📂 [[Terraform Multi Account AWS Landing Zones and Control Tower|05. Multi Cloud and Multi Account Landing Zones]]
- 📂 [[Infrastructure Cost Estimation with Infracost in Pull Requests|06. Infrastructure Cost Estimation with Infracost]]
- 📂 [[Zero Downtime Infrastructure Refactoring and Database Migrations|07. Zero Downtime Infrastructure Refactoring]]

### 📂 [[Terraform|11. Terraform]]
- 📂 [[Terraform Architecture, Core Engine, and Init Plan Apply Workflow|01. Terraform Architecture and Core Workflow]]
- 📂 [[Hashicorp Configuration Language (hcl2) Syntax and Structure|02. Hashicorp Configuration Language HCL Syntax]]
- 📂 [[Terraform State Management, S3 Backends, and Dynamodb Locking|03. State Management and Remote Backends]]
- 📂 [[Terraform State Manipulation, State Migration, and Import|04. State Manipulation and Refactoring]]
- 📂 [[Terraform Providers Architecture, Aliases, and Resource Schemas|05. Providers and Resource Definitions]]
- 📂 [[Terraform Input Variables, Outputs, and Custom Validation Rules|06. Input Variables, Outputs, and Validation]]
- 📂 [[Building Reusable, Modular Terraform Blueprints and Modules|07. Modular Terraform Blueprints and Best Practices]]
- 📂 [[Dynamic Blocks, for Expressions, and Splat Syntax in Terraform|08. Dynamic Blocks and Advanced Expressions]]
- 📂 [[Terraform Built in Functions, String Manipulation, and Templates|09. Built in Functions and Template Generation]]
- 📂 [[Terraform Workspaces vs Directory Based Multi Environment Architecture|10. Terraform Workspaces and Multi Environment Topologies]]
- 📂 [[Terraform Resource Lifecycles and Provisioners|11. Resource Lifecycles and Provisioners]]
- 📂 [[Terragrunt: Dry Terraform Code, Remote State Auto Init, and Dag Execution|12. Terragrunt for Dry Infrastructure Architectures]]
- 📂 [[Terraform Cloud, Enterprise, and Infrastructure Orchestration Platforms|13. Terraform Cloud, Enterprise, and Spacelift]]
- 📂 [[Automated Infrastructure Testing with Terratest in Go|14. Automated Testing with Terratest]]
- 📂 [[Production Terraform Troubleshooting, Lock Clearing, and State Recovery|15. Production Terraform Troubleshooting and State Recovery]]

### 📂 [[DevOps Automation Tooling|12. DevOps Automation Tooling]]
- 📂 [[Continuous Integration (ci) Principles and Pipeline Automation|01. Continuous Integration and CI Tools]]
- 📂 [[Continuous Delivery and Automated Deployment Strategies|02. Continuous Delivery and Deployment CD]]
- 📂 [[Infrastructure Provisioning, Cloud Automation, and Declarative IaC|03. Infrastructure Provisioning Tools]]
- 📂 [[Configuration Management, Ansible Playbooks, and Idempotency|04. Configuration Management and Ansible]]
- 📂 [[Secret Management in CI CD Pipelines and Infrastructure|05. Secret Management in Pipelines]]
- 📂 [[Infrastructure Monitoring, Metrics Collection, and Prometheus|06. Infrastructure Monitoring and Alerting]]
- 📂 [[Centralized Logging, Log Shippers, and Elasticsearch Loki|07. Centralized Logging and Log Aggregation]]
- 📂 [[Distributed Tracing, Span Context Propagation, and Opentelemetry|08. Distributed Tracing and Opentelemetry]]
- 📂 [[Disaster Recovery Planning, Backup Automation, and Rto Rpo|09. Disaster Recovery and Backup Automation]]
- 📂 [[Programming Languages for DevOps (Python, Go, Rust, Bash)|10. Learn a Programming Language for DevOps]]
- 📂 [[Site Reliability Engineering (SRE) Principles and Incident Command|11. SRE Culture and Incident Response]]

### 📂 [[AWS Cloud Platform|13. AWS Cloud Platform]]
- 📂 [[Amazon EC2 Architecture, Instance Types, and Auto Scaling Groups (asg)|01. Elastic Compute Cloud EC2 and Auto Scaling Groups]]
- 📂 [[Amazon S3 Storage Classes, Lifecycle Policies, and Bucket Hardening|02. Simple Storage Service S3 Architecture and Security]]
- 📂 [[Amazon Ebs Volume Types (gp3, Io2) and Amazon Efs Shared Storage|03. Elastic Block Store Ebs and Elastic File System Efs]]
- 📂 [[Amazon RDS Multi Az, Read Replicas, and Aurora Distributed Storage|04. Relational Database Service RDS and Aurora Architecture]]
- 📂 [[Amazon Dynamodb Architecture, Single Table Design, and Partitioning|05. Dynamodb Distributed Nosql Database]]
- 📂 [[AWS Lambda Internals, Event Source Mappings, and Step Functions|06. AWS Lambda and Serverless Compute Architectures]]
- 📂 [[Amazon ECS Architecture, Task Definitions, and AWS Fargate Serverless|07. Elastic Container Service ECS and Fargate]]
- 📂 [[Amazon EKS Production Cluster Management and Managed Node Groups|08. Elastic Kubernetes Service EKS Production Clusters]]
- 📂 [[AWS IAM Policies, Roles, Sts Assumerole, and Permission Boundaries|09. Identity and Access Management IAM Deep Dive]]
- 📂 [[AWS Key Management Service (kms) and AWS Secrets Manager|10. Key Management Service Kms and Secrets Manager]]
- 📂 [[Amazon Cloudwatch Metrics, Logs Insights, and Cloudtrail Audit Trails|11. Cloudwatch, Cloudtrail, and AWS Observability]]
- 📂 [[AWS Security Services: Guardduty, AWS Waf, and AWS Shield Advanced|12. AWS Security Services Guardduty, Waf, and Shield]]
- 📂 [[AWS Well Architected Framework (6 Pillars) and Finops Cost Governance|13. AWS Well Architected Framework and Finops Cost Management]]
- 📂 [[AWS Cloudformation and AWS Cloud Development Kit (cdk)|14. Cloudformation and AWS Cloud Development Kit Cdk]]
- 📂 [[Public Cloud Providers (AWS, GCP, Azure) and Hybrid Cloud|15. Cloud Providers and Hybrid Deployments]]

### 📂 [[AWS Enterprise Infrastructure|14. AWS Enterprise Infrastructure]]
- 📂 [[AWS Global Infrastructure, Regions, Availability Zones, and Edge Locations|01. AWS Global Infrastructure and Multi Region Design]]
- 📂 [[Amazon VPC Architecture, Subnets, Route Tables, and Internet Gateways|02. Virtual Private Cloud VPC and Network Topologies]]
- 📂 [[AWS Transit Gateway, Direct Connect, and Site to Site VPN|03. AWS Transit Gateway and Hybrid Cloud Connectivity]]
- 📂 [[Elastic Load Balancing (alb, Nlb, Glb) and Target Groups|04. Elastic Load Balancing Elb and Application Load Balancer]]
- 📂 [[Amazon Route 53 DNS Routing Policies and Health Checks|05. Route 53 DNS and Global Traffic Management]]

### 📂 [[Cloudflare Edge Computing|15. Cloudflare Edge Computing]]
- 📂 [[Cloudflare Global Architecture, Anycast BGP Routing, and Pops|01. Cloudflare Architecture and Anycast Global Network]]
- 📂 [[Cloudflare Workers Serverless Runtime and V8 Isolate Mechanics|02. Cloudflare Workers and V8 Isolate Architecture]]
- 📂 [[Cloudflare WAF Architecture, Managed Rulesets, and Custom Expressions|03. Cloudflare Web Application Firewall WAF and Rulesets]]
- 📂 [[Edge Ddos Mitigation, L3 L4 L7 Protection, and Rate Limiting|04. Ddos Mitigation and Rate Limiting at the Edge]]
- 📂 [[Cloudflare Bot Management, Super Bot Fight Mode, and Turnstile|05. Cloudflare Bot Management and Turnstile]]
- 📂 [[Cloudflare Zero Trust Architecture, Access, and Tunnel (cloudflared)|06. Cloudflare Zero Trust and Access Gateway]]
- 📂 [[Cloudflare Analytics, Real Time Web Traffic, and Logpush Pipelines|07. Cloudflare Observability, Analytics, and Logpush]]

### 📂 [[CDN Infrastructure|16. CDN Infrastructure]]
- 📂 [[Cloudflare Managed Dns, Proxied vs DNS Only Records, and Dnssec|01. Edge DNS and Managed DNS Records]]
- 📂 [[Cloudflare CDN Caching Rules, Cache Keys, and Instant Purging|02. CDN Caching Strategies and Edge Purging]]
- 📂 [[Cloudflare Edge Storage: Workers Kv, D1 Sql Database, and R2 Storage|03. Edge Storage Kv, D1, and R2 Object Storage]]
- 📂 [[Cloudflare Transform Rules, Url Rewrites, and Header Modification|04. Page Rules, Transform Rules, and Url Rewrites]]
- 📂 [[Cloudflare SSL TLS Modes (off, Flexible, Full, Full Strict)|05. SSL TLS Encryption Modes and Origin Certificates]]

### 📂 [[Enterprise Protocols|17. Enterprise Protocols]]
- 📂 [[Osi Model and Tcp Ip Protocol Suite - Encapsulation and Data Flow|01. Osi Model and Tcp Ip Protocol Suite Architecture]]
- 📂 [[Ip Addressing Architecture - Ipv4 Subnetting, Cidr, Vlsm, and Ipv6|02. Ip Addressing, Cidr Subnetting, and Ipv6 Migration]]
- 📂 [[Enterprise Routing Protocols - BGP (border Gateway Protocol) and OSPF|03. Enterprise Routing Protocols Bgp, Ospf, and Rip]]
- 📂 [[Layer 2 Switching Architecture - Vlans, 802.1q Trunking, and Stp-rstp|04. Layer 2 Switching, Vlans, and Spanning Tree Protocol]]
- 📂 [[Enterprise Network Services - DNS Anycast, Dhcp Relay, and Ntp Stratums|05. Enterprise Network Services DNS Anycast, Dhcp, and Ntp]]

### 📂 [[DevSecOps|18. DevSecOps]]
- 📂 [[Shift Left Security Culture and Devsecops Engineering Standards|01. Shift Left Security Culture and Devsecops Frameworks]]
- 📂 [[Automated Sast, Dast, and Secret Scanning in CI CD Pipelines|02. Automated SAST and DAST in CI CD Pipelines]]
- 📂 [[Infrastructure As Code (iac) Security Linting and Policy Enforcement|03. Infrastructure As Code Security Linting and Guardrails]]
- 📂 [[Software Supply Chain Security: SBOM Generation, Slsa, and Cosign Verification|04. Software Supply Chain Security Slsa, Cosign, and SBOM]]
- 📂 [[Continuous Security Compliance, Audit Trails, and Devsecops Runbooks|05. Continuous Compliance, Audit Trails, and Devsecops Runbooks]]

### 📂 [[Cloud-Native Security Automation|19. Cloud-Native Security Automation]]
- 📂 [[Container Image Security Hardening and Vulnerability Scanning (trivy)|01. Container Image Security and Vulnerability Scanning]]
- 📂 [[Dynamic Secrets Injection, Vault Agent, and Cloud IAM Workload Identity|02. Dynamic Secrets Injection and Ephemeral Credentials]]
- 📂 [[Cloud Native Runtime Security and Anomaly Detection with Falco and Ebpf|03. Cloud Native Runtime Security with Falco and Ebpf]]
- 📂 [[Kubernetes Policy As Code: Open Policy Agent (gatekeeper) and Kyverno|04. Policy As Code with Open Policy Agent OPA and Kyverno]]
- 📂 [[Cloud Security Posture Management (cspm) and Automated Remediation|05. Cloud Security Posture Management Cspm Automation]]

### 📂 [[MLOps & Machine Learning Operations|20. MLOps & Machine Learning Operations]]
- 📂 [[MLOps Architecture, Maturity Levels, and CI CD for Machine Learning|01. MLOps Architecture and ML Lifecycle Standards]]
- 📂 [[Data Versioning, Pipeline Lineage, and Data Version Control (dvc)|02. Data Versioning and Lineage with DVC]]
- 📂 [[Feature Stores Architecture (feast, Hopsworks) and Point in Time Joins|03. Feature Stores Architecture Feast and Hopsworks]]
- 📂 [[Experiment Tracking, Hyperparameter Logging, and Mlflow Model Registry|04. Experiment Tracking and Model Registry with Mlflow]]
- 📂 [[Distributed ML Training Pipelines, Kubeflow, and Ray Train|05. Distributed Training Pipelines and Kubeflow]]
- 📂 [[High Throughput Model Inference Serving (triton Inference Server, Vllm)|06. High Throughput Model Serving Triton and Vllm]]
- 📂 [[Model Packaging, Export Formats (onnx, Torchscript), and Bentoml|07. Model Packaging and Containerization Bentoml]]
- 📂 [[Model Monitoring in Production: Data Drift, Concept Drift, and Evidently Ai|08. Model Monitoring, Data Drift, and Concept Drift]]
- 📂 [[Model Deployment Strategies (shadow Deployments, a B Testing, Canary)|09. Model Deployment Strategies Canary and Shadow Deployments]]
- 📂 [[GPU Cluster Orchestration, Nvidia GPU Operator, and Mig in Kubernetes|10. GPU Cluster Orchestration and Scheduling in Kubernetes]]
- 📂 [[Llmops: Prompt Management, Vector Db Sync, and Continuous Evaluation|11. Production Llm Operations and Serving Llmops]]

### 📂 [[Kernel Engineering|21. Kernel Engineering]]
- 📂 [[Linux Memory Management, Virtual Memory, Buffers Cache, and Swap|01. Linux Memory Architecture and Swap Management]]
- 📂 [[Linux Filesystems Internals: Inodes, Superblocks, Ext4, and Xfs|02. Linux Filesystem Internals Ext4 and Xfs]]
- 📂 [[Linux Packet Filtering with Netfilter, Iptables, and Nftables|03. Linux Firewalling Netfilter, Iptables, and Nftables]]
- 📂 [[Linux Kernel Performance Tuning via Sysctl and Procfs|04. Linux Kernel Performance Tuning with Sysctl]]
- 📂 [[Linux Security Modules: Selinux Contexts and Apparmor Profiles|05. Linux Security Modules Apparmor and Selinux]]

## 🔗 Navigation
- ⬆️ Parent: [[Infrastructure & Security]]
- 🏛️ Software Architecture: `Architecture`
- 💻 Computer Science Foundations: `Computer Science`
- 🛡️ Cyber Security: `Cyber Security`

---

## 🗂️ Topics

- [[Linux Systems Administration]]
- [[Network Engineering]]
- [[Git Version Control]]
- [[GitHub and CI-CD Automation]]
- [[Docker]]
- [[Docker Swarm]]
- [[Container Runtime Internals]]
- [[Kubernetes]]
- [[Cloud-Native Orchestration]]
- [[Infrastructure as Code]]
- [[Terraform]]
- [[DevOps Automation Tooling]]
- [[AWS Cloud Platform]]
- [[AWS Enterprise Infrastructure]]
- [[Cloudflare Edge Computing]]
- [[CDN Infrastructure]]
- [[Enterprise Protocols]]
- [[DevSecOps]]
- [[Cloud-Native Security Automation]]
- [[MLOps & Machine Learning Operations]]
- [[Kernel Engineering]]
