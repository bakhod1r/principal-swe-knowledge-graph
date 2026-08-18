---
title: DevOps
tags:
  - devops
  - platform-engineering
  - cloud-infrastructure
  - principal-swe
parent: "[[Principal SWE]]"
---

# 🚀 DevOps, Cloud Infrastructure & Platform Engineering

Comprehensive, production-grade master architecture covering the complete spectrum of enterprise cloud platforms, continuous delivery, and infrastructure reliability: Docker containerization, Kubernetes orchestration, Terraform Infrastructure-as-Code (IaC), AWS enterprise cloud services, Linux systems administration & eBPF, Cloudflare edge computing (Workers, D1, R2), and production MLOps operations across 8 master pillars and 119 specialized subdomains.

```text
DevOps
│
├── [[Core DevOps Engineering|01. Core DevOps Engineering]]
│   ├── [[Learn a Programming Language|01. Learn a Programming Language]]
│   ├── [[Operating System|02. Operating System]]
│   ├── [[Terminal Knowledge|03. Terminal Knowledge]]
│   ├── [[Version Control Systems (Core DevOps Engineering)|04. Version Control Systems]]
│   ├── [[Containers|05. Containers]]
│   ├── [[Networking Protocols (Core DevOps Engineering)|06. Networking Protocols]]
│   ├── [[Cloud Providers|07. Cloud Providers]]
│   ├── [[Serverless|08. Serverless]]
│   ├── [[Configuration Management|09. Configuration Management]]
│   ├── [[Provisioning|10. Provisioning]]
│   ├── [[CI CD Tools|11. CI CD Tools]]
│   ├── [[Secret Management|12. Secret Management]]
│   ├── [[Infrastructure Monitoring|13. Infrastructure Monitoring]]
│   ├── [[Logs Management|14. Logs Management]]
│   ├── [[Container Orchestration|15. Container Orchestration]]
│   ├── [[Application Monitoring|16. Application Monitoring]]
│   ├── [[Artifact Management|17. Artifact Management]]
│   ├── [[GitOps|18. GitOps]]
│   ├── [[Service Mesh|19. Service Mesh]]
│   └── [[Cloud Design Patterns (Core DevOps Engineering)|20. Cloud Design Patterns]]
├── [[Docker & Containerization|02. Docker & Containerization]]
│   ├── [[Introduction (Docker & Containerization)|01. Introduction]]
│   ├── [[Underlying Technologies|02. Underlying Technologies]]
│   ├── [[Installation Setup|03. Installation Setup]]
│   ├── [[Basics of Docker|04. Basics of Docker]]
│   ├── [[Data Persistence|05. Data Persistence]]
│   ├── [[Using 3rd Party Container Images|06. Using 3rd Party Container Images]]
│   ├── [[Runtime Configuration Options|07. Runtime Configuration Options]]
│   ├── [[Building Container Images|08. Building Container Images]]
│   ├── [[Container Registries|09. Container Registries]]
│   ├── [[Running Containers|10. Running Containers]]
│   ├── [[Docker CLI|11. Docker CLI]]
│   ├── [[Container Security|12. Container Security]]
│   ├── [[Developer Experience|13. Developer Experience]]
│   └── [[Deploying Containers|14. Deploying Containers]]
├── [[Kubernetes & Container Orchestration|03. Kubernetes & Container Orchestration]]
│   ├── [[Introduction (Kubernetes & Container Orchestration)|01. Introduction]]
│   ├── [[Setting Up Kubernetes|02. Setting Up Kubernetes]]
│   ├── [[Running Applications|03. Running Applications]]
│   ├── [[Configuration Management (Kubernetes & Container Orchestration)|04. Configuration Management]]
│   ├── [[Services and Networking|05. Services and Networking]]
│   ├── [[Security (Kubernetes & Container Orchestration)|06. Security]]
│   ├── [[Resource Management|07. Resource Management]]
│   ├── [[Monitoring and Logging|08. Monitoring and Logging]]
│   ├── [[Storage and Volumes|09. Storage and Volumes]]
│   ├── [[Scheduling|10. Scheduling]]
│   ├── [[Autoscaling (Kubernetes & Container Orchestration)|11. Autoscaling]]
│   ├── [[Deployment Patterns|12. Deployment Patterns]]
│   └── [[Advanced Topics (Kubernetes & Container Orchestration)|13. Advanced Topics]]
├── [[Terraform & Infrastructure As Code|04. Terraform & Infrastructure as Code]]
│   ├── [[Introduction (Terraform & Infrastructure As Code)|01. Introduction]]
│   ├── [[Getting Started|02. Getting Started]]
│   ├── [[Providers|03. Providers]]
│   ├── [[Resources|04. Resources]]
│   ├── [[Variables|05. Variables]]
│   ├── [[Outputs|06. Outputs]]
│   ├── [[Format Validate|07. Format Validate]]
│   ├── [[State Management|08. State Management]]
│   ├── [[Deployment|09. Deployment]]
│   ├── [[Clean Up|10. Clean Up]]
│   ├── [[Modules|11. Modules]]
│   ├── [[Provisioners|12. Provisioners]]
│   ├── [[Data Sources|13. Data Sources]]
│   ├── [[Template Files|14. Template Files]]
│   ├── [[Workspaces|15. Workspaces]]
│   ├── [[CI CD Integration|16. CI CD Integration]]
│   ├── [[Testing (Terraform & Infrastructure As Code)|17. Testing]]
│   ├── [[Scaling Terraform|18. Scaling Terraform]]
│   ├── [[Security (Terraform & Infrastructure As Code)|19. Security]]
│   └── [[HCP|20. HCP]]
├── [[AWS Cloud Infrastructure|05. AWS Cloud Infrastructure]]
│   ├── [[Introduction (AWS Cloud Infrastructure)|01. Introduction]]
│   ├── [[EC2|02. EC2]]
│   ├── [[VPC|03. VPC]]
│   ├── [[IAM|04. IAM]]
│   ├── [[SES|05. SES]]
│   ├── [[Route 53|06. Route 53]]
│   ├── [[CloudWatch|07. CloudWatch]]
│   ├── [[S3|08. S3]]
│   ├── [[Auto Scaling|09. Auto Scaling]]
│   ├── [[RDS|10. RDS]]
│   ├── [[DynamoDB|11. DynamoDB]]
│   ├── [[CloudFront|12. CloudFront]]
│   ├── [[ElastiCache|13. ElastiCache]]
│   ├── [[ECR|14. ECR]]
│   ├── [[Lambda|15. Lambda]]
│   ├── [[EKS|16. EKS]]
│   └── [[ECS|17. ECS]]
├── [[Linux Systems & Administration|06. Linux Systems & Administration]]
│   ├── [[Navigation Basics|01. Navigation Basics]]
│   ├── [[Editing Files|02. Editing Files]]
│   ├── [[Shell and Other Basics|03. Shell and Other Basics]]
│   ├── [[Text Processing|04. Text Processing]]
│   ├── [[Working with Files|05. Working with Files]]
│   ├── [[User Management|06. User Management]]
│   ├── [[Process Management|07. Process Management]]
│   ├── [[Server Review|08. Server Review]]
│   ├── [[Service Management systemd|09. Service Management systemd]]
│   ├── [[Package Management|10. Package Management]]
│   ├── [[Disks and Filesystems|11. Disks and Filesystems]]
│   ├── [[Booting Linux|12. Booting Linux]]
│   ├── [[Networking|13. Networking]]
│   ├── [[Shell Programming|14. Shell Programming]]
│   ├── [[Troubleshooting|15. Troubleshooting]]
│   └── [[Containerization|16. Containerization]]
├── [[Cloudflare & Edge Computing|07. Cloudflare & Edge Computing]]
│   ├── [[Prerequisites|01. Prerequisites]]
│   ├── [[Core Development Skills|02. Core Development Skills]]
│   ├── [[Workers Core Concepts|03. Workers Core Concepts]]
│   ├── [[Frameworks and Tools|04. Frameworks and Tools]]
│   ├── [[Storage Solutions|05. Storage Solutions]]
│   ├── [[Durable Execution|06. Durable Execution]]
│   ├── [[Advanced Features (Cloudflare & Edge Computing)|07. Advanced Features]]
│   ├── [[Security Performance|08. Security Performance]]
│   ├── [[Integration Workflows|09. Integration Workflows]]
│   └── [[Development Tools|10. Development Tools]]
└── [[MLOps & Machine Learning Operations|08. MLOps & Machine Learning Operations]]
│   ├── [[Programming Fundamentals (MLOps & Machine Learning Operations)|01. Programming Fundamentals]]
│   ├── [[Version Control Systems (MLOps & Machine Learning Operations)|02. Version Control Systems]]
│   ├── [[Cloud Computing|03. Cloud Computing]]
│   ├── [[Containerization (MLOps & Machine Learning Operations)|04. Containerization]]
│   ├── [[Machine Learning Fundamentals|05. Machine Learning Fundamentals]]
│   ├── [[Data Engineering Fundamentals|06. Data Engineering Fundamentals]]
│   ├── [[MLOps Principles|07. MLOps Principles]]
│   ├── [[MLOps Components|08. MLOps Components]]
│   └── [[Infrastructure As Code (MLOps & Machine Learning Operations)|09. Infrastructure As Code]]
```

---

## 🏛️ Core Knowledge Pillars

### 1. 📂 [[Core DevOps Engineering|01. Core DevOps Engineering]]
- 📂 [[Learn a Programming Language|01. Learn a Programming Language]] — Platform blueprints and automation patterns for Learn a Programming Language.
- 📂 [[Operating System|02. Operating System]] — Platform blueprints and automation patterns for Operating System.
- 📂 [[Terminal Knowledge|03. Terminal Knowledge]] — Platform blueprints and automation patterns for Terminal Knowledge.
- 📂 [[Version Control Systems (Core DevOps Engineering)|04. Version Control Systems]] — Platform blueprints and automation patterns for Version Control Systems.
- 📂 [[Containers|05. Containers]] — Platform blueprints and automation patterns for Containers.
- 📂 [[Networking Protocols (Core DevOps Engineering)|06. Networking Protocols]] — Platform blueprints and automation patterns for Networking Protocols.
- 📂 [[Cloud Providers|07. Cloud Providers]] — Platform blueprints and automation patterns for Cloud Providers.
- 📂 [[Serverless|08. Serverless]] — Platform blueprints and automation patterns for Serverless.
- 📂 [[Configuration Management|09. Configuration Management]] — Platform blueprints and automation patterns for Configuration Management.
- 📂 [[Provisioning|10. Provisioning]] — Platform blueprints and automation patterns for Provisioning.
- 📂 [[CI CD Tools|11. CI CD Tools]] — Platform blueprints and automation patterns for CI CD Tools.
- 📂 [[Secret Management|12. Secret Management]] — Platform blueprints and automation patterns for Secret Management.
- 📂 [[Infrastructure Monitoring|13. Infrastructure Monitoring]] — Platform blueprints and automation patterns for Infrastructure Monitoring.
- 📂 [[Logs Management|14. Logs Management]] — Platform blueprints and automation patterns for Logs Management.
- 📂 [[Container Orchestration|15. Container Orchestration]] — Platform blueprints and automation patterns for Container Orchestration.
- 📂 [[Application Monitoring|16. Application Monitoring]] — Platform blueprints and automation patterns for Application Monitoring.
- 📂 [[Artifact Management|17. Artifact Management]] — Platform blueprints and automation patterns for Artifact Management.
- 📂 [[GitOps|18. GitOps]] — Platform blueprints and automation patterns for GitOps.
- 📂 [[Service Mesh|19. Service Mesh]] — Platform blueprints and automation patterns for Service Mesh.
- 📂 [[Cloud Design Patterns (Core DevOps Engineering)|20. Cloud Design Patterns]] — Platform blueprints and automation patterns for Cloud Design Patterns.
### 2. 📂 [[Docker & Containerization|02. Docker & Containerization]]
- 📂 [[Introduction (Docker & Containerization)|01. Introduction]] — Platform blueprints and automation patterns for Introduction.
- 📂 [[Underlying Technologies|02. Underlying Technologies]] — Platform blueprints and automation patterns for Underlying Technologies.
- 📂 [[Installation Setup|03. Installation Setup]] — Platform blueprints and automation patterns for Installation Setup.
- 📂 [[Basics of Docker|04. Basics of Docker]] — Platform blueprints and automation patterns for Basics of Docker.
- 📂 [[Data Persistence|05. Data Persistence]] — Platform blueprints and automation patterns for Data Persistence.
- 📂 [[Using 3rd Party Container Images|06. Using 3rd Party Container Images]] — Platform blueprints and automation patterns for Using 3rd Party Container Images.
- 📂 [[Runtime Configuration Options|07. Runtime Configuration Options]] — Platform blueprints and automation patterns for Runtime Configuration Options.
- 📂 [[Building Container Images|08. Building Container Images]] — Platform blueprints and automation patterns for Building Container Images.
- 📂 [[Container Registries|09. Container Registries]] — Platform blueprints and automation patterns for Container Registries.
- 📂 [[Running Containers|10. Running Containers]] — Platform blueprints and automation patterns for Running Containers.
- 📂 [[Docker CLI|11. Docker CLI]] — Platform blueprints and automation patterns for Docker CLI.
- 📂 [[Container Security|12. Container Security]] — Platform blueprints and automation patterns for Container Security.
- 📂 [[Developer Experience|13. Developer Experience]] — Platform blueprints and automation patterns for Developer Experience.
- 📂 [[Deploying Containers|14. Deploying Containers]] — Platform blueprints and automation patterns for Deploying Containers.
### 3. 📂 [[Kubernetes & Container Orchestration|03. Kubernetes & Container Orchestration]]
- 📂 [[Introduction (Kubernetes & Container Orchestration)|01. Introduction]] — Platform blueprints and automation patterns for Introduction.
- 📂 [[Setting Up Kubernetes|02. Setting Up Kubernetes]] — Platform blueprints and automation patterns for Setting Up Kubernetes.
- 📂 [[Running Applications|03. Running Applications]] — Platform blueprints and automation patterns for Running Applications.
- 📂 [[Configuration Management (Kubernetes & Container Orchestration)|04. Configuration Management]] — Platform blueprints and automation patterns for Configuration Management.
- 📂 [[Services and Networking|05. Services and Networking]] — Platform blueprints and automation patterns for Services and Networking.
- 📂 [[Security (Kubernetes & Container Orchestration)|06. Security]] — Platform blueprints and automation patterns for Security.
- 📂 [[Resource Management|07. Resource Management]] — Platform blueprints and automation patterns for Resource Management.
- 📂 [[Monitoring and Logging|08. Monitoring and Logging]] — Platform blueprints and automation patterns for Monitoring and Logging.
- 📂 [[Storage and Volumes|09. Storage and Volumes]] — Platform blueprints and automation patterns for Storage and Volumes.
- 📂 [[Scheduling|10. Scheduling]] — Platform blueprints and automation patterns for Scheduling.
- 📂 [[Autoscaling (Kubernetes & Container Orchestration)|11. Autoscaling]] — Platform blueprints and automation patterns for Autoscaling.
- 📂 [[Deployment Patterns|12. Deployment Patterns]] — Platform blueprints and automation patterns for Deployment Patterns.
- 📂 [[Advanced Topics (Kubernetes & Container Orchestration)|13. Advanced Topics]] — Platform blueprints and automation patterns for Advanced Topics.
### 4. 📂 [[Terraform & Infrastructure As Code|04. Terraform & Infrastructure as Code]]
- 📂 [[Introduction (Terraform & Infrastructure As Code)|01. Introduction]] — Platform blueprints and automation patterns for Introduction.
- 📂 [[Getting Started|02. Getting Started]] — Platform blueprints and automation patterns for Getting Started.
- 📂 [[Providers|03. Providers]] — Platform blueprints and automation patterns for Providers.
- 📂 [[Resources|04. Resources]] — Platform blueprints and automation patterns for Resources.
- 📂 [[Variables|05. Variables]] — Platform blueprints and automation patterns for Variables.
- 📂 [[Outputs|06. Outputs]] — Platform blueprints and automation patterns for Outputs.
- 📂 [[Format Validate|07. Format Validate]] — Platform blueprints and automation patterns for Format Validate.
- 📂 [[State Management|08. State Management]] — Platform blueprints and automation patterns for State Management.
- 📂 [[Deployment|09. Deployment]] — Platform blueprints and automation patterns for Deployment.
- 📂 [[Clean Up|10. Clean Up]] — Platform blueprints and automation patterns for Clean Up.
- 📂 [[Modules|11. Modules]] — Platform blueprints and automation patterns for Modules.
- 📂 [[Provisioners|12. Provisioners]] — Platform blueprints and automation patterns for Provisioners.
- 📂 [[Data Sources|13. Data Sources]] — Platform blueprints and automation patterns for Data Sources.
- 📂 [[Template Files|14. Template Files]] — Platform blueprints and automation patterns for Template Files.
- 📂 [[Workspaces|15. Workspaces]] — Platform blueprints and automation patterns for Workspaces.
- 📂 [[CI CD Integration|16. CI CD Integration]] — Platform blueprints and automation patterns for CI CD Integration.
- 📂 [[Testing (Terraform & Infrastructure As Code)|17. Testing]] — Platform blueprints and automation patterns for Testing.
- 📂 [[Scaling Terraform|18. Scaling Terraform]] — Platform blueprints and automation patterns for Scaling Terraform.
- 📂 [[Security (Terraform & Infrastructure As Code)|19. Security]] — Platform blueprints and automation patterns for Security.
- 📂 [[HCP|20. HCP]] — Platform blueprints and automation patterns for HCP.
### 5. 📂 [[AWS Cloud Infrastructure|05. AWS Cloud Infrastructure]]
- 📂 [[Introduction (AWS Cloud Infrastructure)|01. Introduction]] — Platform blueprints and automation patterns for Introduction.
- 📂 [[EC2|02. EC2]] — Platform blueprints and automation patterns for EC2.
- 📂 [[VPC|03. VPC]] — Platform blueprints and automation patterns for VPC.
- 📂 [[IAM|04. IAM]] — Platform blueprints and automation patterns for IAM.
- 📂 [[SES|05. SES]] — Platform blueprints and automation patterns for SES.
- 📂 [[Route 53|06. Route 53]] — Platform blueprints and automation patterns for Route 53.
- 📂 [[CloudWatch|07. CloudWatch]] — Platform blueprints and automation patterns for CloudWatch.
- 📂 [[S3|08. S3]] — Platform blueprints and automation patterns for S3.
- 📂 [[Auto Scaling|09. Auto Scaling]] — Platform blueprints and automation patterns for Auto Scaling.
- 📂 [[RDS|10. RDS]] — Platform blueprints and automation patterns for RDS.
- 📂 [[DynamoDB|11. DynamoDB]] — Platform blueprints and automation patterns for DynamoDB.
- 📂 [[CloudFront|12. CloudFront]] — Platform blueprints and automation patterns for CloudFront.
- 📂 [[ElastiCache|13. ElastiCache]] — Platform blueprints and automation patterns for ElastiCache.
- 📂 [[ECR|14. ECR]] — Platform blueprints and automation patterns for ECR.
- 📂 [[Lambda|15. Lambda]] — Platform blueprints and automation patterns for Lambda.
- 📂 [[EKS|16. EKS]] — Platform blueprints and automation patterns for EKS.
- 📂 [[ECS|17. ECS]] — Platform blueprints and automation patterns for ECS.
### 6. 📂 [[Linux Systems & Administration|06. Linux Systems & Administration]]
- 📂 [[Navigation Basics|01. Navigation Basics]] — Platform blueprints and automation patterns for Navigation Basics.
- 📂 [[Editing Files|02. Editing Files]] — Platform blueprints and automation patterns for Editing Files.
- 📂 [[Shell and Other Basics|03. Shell and Other Basics]] — Platform blueprints and automation patterns for Shell and Other Basics.
- 📂 [[Text Processing|04. Text Processing]] — Platform blueprints and automation patterns for Text Processing.
- 📂 [[Working with Files|05. Working with Files]] — Platform blueprints and automation patterns for Working with Files.
- 📂 [[User Management|06. User Management]] — Platform blueprints and automation patterns for User Management.
- 📂 [[Process Management|07. Process Management]] — Platform blueprints and automation patterns for Process Management.
- 📂 [[Server Review|08. Server Review]] — Platform blueprints and automation patterns for Server Review.
- 📂 [[Service Management systemd|09. Service Management systemd]] — Platform blueprints and automation patterns for Service Management systemd.
- 📂 [[Package Management|10. Package Management]] — Platform blueprints and automation patterns for Package Management.
- 📂 [[Disks and Filesystems|11. Disks and Filesystems]] — Platform blueprints and automation patterns for Disks and Filesystems.
- 📂 [[Booting Linux|12. Booting Linux]] — Platform blueprints and automation patterns for Booting Linux.
- 📂 [[Networking|13. Networking]] — Platform blueprints and automation patterns for Networking.
- 📂 [[Shell Programming|14. Shell Programming]] — Platform blueprints and automation patterns for Shell Programming.
- 📂 [[Troubleshooting|15. Troubleshooting]] — Platform blueprints and automation patterns for Troubleshooting.
- 📂 [[Containerization|16. Containerization]] — Platform blueprints and automation patterns for Containerization.
### 7. 📂 [[Cloudflare & Edge Computing|07. Cloudflare & Edge Computing]]
- 📂 [[Prerequisites|01. Prerequisites]] — Platform blueprints and automation patterns for Prerequisites.
- 📂 [[Core Development Skills|02. Core Development Skills]] — Platform blueprints and automation patterns for Core Development Skills.
- 📂 [[Workers Core Concepts|03. Workers Core Concepts]] — Platform blueprints and automation patterns for Workers Core Concepts.
- 📂 [[Frameworks and Tools|04. Frameworks and Tools]] — Platform blueprints and automation patterns for Frameworks and Tools.
- 📂 [[Storage Solutions|05. Storage Solutions]] — Platform blueprints and automation patterns for Storage Solutions.
- 📂 [[Durable Execution|06. Durable Execution]] — Platform blueprints and automation patterns for Durable Execution.
- 📂 [[Advanced Features (Cloudflare & Edge Computing)|07. Advanced Features]] — Platform blueprints and automation patterns for Advanced Features.
- 📂 [[Security Performance|08. Security Performance]] — Platform blueprints and automation patterns for Security Performance.
- 📂 [[Integration Workflows|09. Integration Workflows]] — Platform blueprints and automation patterns for Integration Workflows.
- 📂 [[Development Tools|10. Development Tools]] — Platform blueprints and automation patterns for Development Tools.
### 8. 📂 [[MLOps & Machine Learning Operations|08. MLOps & Machine Learning Operations]]
- 📂 [[Programming Fundamentals (MLOps & Machine Learning Operations)|01. Programming Fundamentals]] — Platform blueprints and automation patterns for Programming Fundamentals.
- 📂 [[Version Control Systems (MLOps & Machine Learning Operations)|02. Version Control Systems]] — Platform blueprints and automation patterns for Version Control Systems.
- 📂 [[Cloud Computing|03. Cloud Computing]] — Platform blueprints and automation patterns for Cloud Computing.
- 📂 [[Containerization (MLOps & Machine Learning Operations)|04. Containerization]] — Platform blueprints and automation patterns for Containerization.
- 📂 [[Machine Learning Fundamentals|05. Machine Learning Fundamentals]] — Platform blueprints and automation patterns for Machine Learning Fundamentals.
- 📂 [[Data Engineering Fundamentals|06. Data Engineering Fundamentals]] — Platform blueprints and automation patterns for Data Engineering Fundamentals.
- 📂 [[MLOps Principles|07. MLOps Principles]] — Platform blueprints and automation patterns for MLOps Principles.
- 📂 [[MLOps Components|08. MLOps Components]] — Platform blueprints and automation patterns for MLOps Components.
- 📂 [[Infrastructure As Code (MLOps & Machine Learning Operations)|09. Infrastructure As Code]] — Platform blueprints and automation patterns for Infrastructure As Code.

---

## 🔗 Navigation
- ⬆️ Parent: [[Principal SWE]]
- 🎓 Root: [[Principal SWE]]
