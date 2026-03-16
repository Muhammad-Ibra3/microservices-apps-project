# Go-CQRS-Kafka-gRPC Microservices (Clone)

This repository is a **clone of [Go-CQRS-Kafka-gRPC-Microservices](https://github.com/AleksK1NG/Go-CQRS-Kafka-gRPC-Microservices)**, optimized for **cloud-native deployments, GitOps, and enterprise-grade DevOps practices**.

> This clone exists because the microservices are integrated into my **Platform Engineering project**, where I created **custom Helm charts, GitOps pipelines, and developer self-service workflows via Backstage**.  
> The Helm charts for these services live in my **platform project**, allowing CI/CD, preview environments, and policy enforcement to be demonstrated.  

It contains three microservices:

1. **API Gateway Service** – handles external requests and routes them to internal services.
2. **Reader Service** – reads and processes events from Kafka.
3. **Writer Service** – publishes events to Kafka.

- All services are written in **Go**.  
- Communication between services is via **gRPC**.  
- Events are published and consumed via **Kafka**.  
- Services are **containerized** with **multi-stage Dockerfiles** using **distroless images** for security.  
- Images are **owned by root** but executed as a **non-root user**.

---

## 🧩 Dependencies

The services depend on:

- **Kafka** – for event streaming  
- **MongoDB** – for data persistence (Reader Service)  
- **Postgres** – for relational data storage  
- **Redis** – optional caching

Dependencies can be deployed either:

1. **Locally**, via `docker-compose`:

```bash
docker-compose -f docker-compose.yml up 
```

2. In Kubernetes, via the Platform Engineering project, where Helm charts manage all dependencies.

---

## 📦 Docker

Each service has its own Dockerfile:

api_gateway_service/Dockerfile

reader_service/Dockerfile

writer_service/Dockerfile


Build an image locally

# Example for API Gateway

```bash
docker build -t cqrs-api-gateway ./api_gateway_service
```

For local development

```bash
make migrate_up // run sql migrations
make mongo // run mongo init scripts
make swagger // generate swagger documentation
make local or docker_dev // for run docker compose files
```

---

## 🚀 Helm Charts

The Helm charts for these services live in my Platform Engineering project.
Minimal charts exist in this repo for local experimentation:

charts/
├─ api-gateway/
├─ reader/
└─ writer/

Install via Helm

```bash
helm install api-gateway charts/api-gateway
helm install reader charts/reader
helm install writer charts/writer
```

---

## ⚙️ CI/CD Pipelines

CI/CD pipelines live in this repository.

They are configured to build, scan, sign, and deploy the Docker images.

The pipelines use reusable templates defined in the Platform Engineering project, ensuring consistent GitOps workflows across multiple microservices.

PRs automatically trigger preview environments, followed by Dev → Prod deployments.

---

## 🔐 Security

- Distroless runtime images

- Non-root execution

- Binary owned by root

CI pipeline enforces supply-chain security:

- Trivy vulnerability scanning

- SBOM generation via Syft

- Cosign image signing


---

## ✅ Key Takeaways

This fork demonstrates:

- Containerized Go microservices

- Production-ready Dockerfiles

- Secure, minimal images

- GitOps-ready deployments with preview environments

- Secure CI pipelines using reusable templates with image scanning and signing
```

