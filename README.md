# Self-Healing and Scalable Application Deployment using Kubernetes with CI/CD Automation

## Overview

This project demonstrates deployment of a scalable and fault-tolerant Go web application using Kubernetes and modern DevOps technologies.

The infrastructure supports:

- High availability using multiple replicas
- Kubernetes self-healing
- Horizontal Pod Autoscaling (HPA)
- Docker containerization
- CI/CD automation with GitHub Actions
- Docker Hub integration

---

## Technologies Used

- Go
- Docker
- Kubernetes
- GitHub Actions
- Docker Hub
- Git

---

## Project Structure

```text
.
├── k8s/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── hpa.yaml
├── .github/workflows/
│   └── ci-cd.yml
├── Dockerfile
├── main.go
└── README.md

```
## Installation and Running

### Requirements

Install the following tools:

- Docker
- Minikube
- kubectl
- Git

---

### Clone Repository

```bash
git clone https://github.com/ustkost/sna_devops_project
cd sna_devops_project
```

### Start multi-node Minikube cluster:
```bash
minikube start --nodes 3
```

### Deploy Application
```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/hpa.yaml
```

### Verify Deployment
```bash
kubectl get nodes
kubectl get pods -o wide
kubectl get svc
kubectl get hpa
```

### Open Application
```bash
minikube service sna-project
```

### Test Self-Healing
Delete one of the pods:
```bash
kubectl delete pod <pod-name>
```
Kubernetes automatically recreates the pod.

### Test Autoscaling
Generate load:
```bash
curl http://<service-ip>/load
```
---
## Docker
### Build image
```bash
docker build -t myapp .
```
### Run container
```bash
docker run -p 8080:8080 myapp
```
---
## CI/CD Pipeline
GitHub Actions pipeline automatically:

- Builds the Go application
- Builds Docker image
- Pushes image to Docker Hub

---
## Team Members
- Timur — Go application, Docker, CI/CD, documentation
- Konstantin — Kubernetes configuration, autoscaling, demo
---
## Demo
https://drive.google.com/drive/folders/1JFk4toS85OfSfH-M1GqjfMgPbrrxfpOh?usp=sharing
