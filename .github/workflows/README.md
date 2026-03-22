
# Microservices Apps Project

This repository contains the **application code for all microservices** in the platform.

It is responsible for:
- Service source code
- Docker build context
- Triggering CI/CD pipelines
- Calling reusable workflow tasks from the platform repository

---

# ⚠️ Important: CI Workflows Are Shared

The CI pipeline is **triggered here**, but the actual workflow logic is **defined in the platform repository**:


Muhammad-Ibra3/platform-engineering-project


This repository acts as the **entrypoint**, while the platform repo provides the reusable CI/CD building blocks.

---

# 📁 Repository Structure


.github/
└── workflows/
├── CI-pipeline.yaml
└── preview-env-destroy.yaml

api_gateway_service/
reader_service/
writer_service/


---

# 🚀 CI Pipeline Overview

The CI pipeline runs on pull requests and follows this flow:


Detect changed services
↓
Build Docker images
↓
Scan images for vulnerabilities
↓
Push images to registry
↓
Sign images (Cosign)
↓
Update GitOps preview values
↓
Argo CD deploys preview environment


---

# 🔄 CI Pipeline Stages

## 1. Detect Services
Uses a reusable workflow to determine which services changed.

Only affected services proceed through the pipeline.

---

## 2. Build Images
Each changed service is built into a Docker image.

---

## 3. Security Scan
Images are scanned before being pushed.

Acts as a security gate.

---

## 4. Push & Sign Image
Images are:
- pushed to the registry
- signed using Cosign

⚠️ Signing happens **after push** because signatures are tied to the registry image.

---

## 5. Update GitOps
The pipeline updates preview environment values in the platform repository:


platform-helm/envs/preview/microservices/<pr-number>/


This triggers Argo CD to deploy/update the preview environment.

---

## 6. Deploy Preview Environment
Argo CD automatically:
- detects Git changes
- deploys preview workloads

---

## 7. Destroy Preview (on PR close)
When a PR is closed:
- preview values are removed
- Argo CD prunes the environment

---

# 🔗 Relationship with Platform Repo

## Platform Repo Responsibilities
- CI workflow task definitions
- GitOps manifests
- Helm charts
- Environment configuration
- Argo CD deployment logic

## This Repo Responsibilities
- Application code
- Dockerfiles
- CI trigger workflows
- Passing inputs to reusable workflows

---

# 🧠 Why This Design?

This separation provides:

- Centralized CI/CD logic
- Consistent deployment behavior
- Easier governance and security
- Less duplication across services

---

# 🔁 Conceptual Flow


Developer opens PR
↓
This repo triggers CI pipeline
↓
Calls reusable workflows from platform repo
↓
Images are built, scanned, pushed, signed
↓
GitOps values updated
↓
Argo CD deploys preview environment


---

# 🧪 Preview Environments

Each PR gets:

- Dedicated namespace
- Isolated deployment
- Unique ingress routing (e.g. preview URLs)

Preview environments are fully automated via GitOps.

---

# ⚠️ Important Notes

- Only changed services are built
- Image tags use commit SHA
- GitOps repo must be accessible via token
- Preview environments are ephemeral

---

# 📌 Summary

This repository:

- Contains application code
- Triggers CI/CD pipelines
- Delegates CI/CD logic to the platform repository

In short:


This repo → triggers pipeline
Platform repo → defines pipeline logic
Argo CD → deploys everything