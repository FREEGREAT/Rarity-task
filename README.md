
# 🧩 Rarible Go App — Kubernetes + Helm Deployment Guide

---

## 📦 Quick Start 

### 1. Clone the repository
```sh
git clone https://github.com/FREEGREAT/Rarity-task.git
cd Rarible-task
```

---

## 🚀 Running in Kubernetes with Helm

### 🔧 Requirements

- Kubernetes cluster (local or cloud)
- [Helm](https://helm.sh/) installed
- Docker (for local runs)
- (Optional) Minikube for local testing

---

---

## 🛠️ Makefile Commands

| Command           | Description                                 |
|------------------|---------------------------------------------|
| `make secret-generic` | Create secret from .env file          |
| `make helm-install`   | First-time Helm installation           |
| `make helm-upgrade`   | Upgrade existing Helm release          |
| `make port-forward`   | Port forward service to localhost:8080 |

---


## 🛡️ Secret Management

1. Create Kubernetes Secret from `.env` file:
```sh
kubectl create secret generic rarible-secret --from-env-file=.env
# or
make secret-generic
```

2. Install or upgrade Helm chart:
```sh
#Install
helm install rarible-app ./helm/rarible-app --namespace default
# or
make helm-install

#Upgrade
helm upgrade rarible-app ./helm/rarible-app --namespace default
# or
make helm-upgrade
```

---

## 🧪 Local Run with Minikube

```sh
minikube start
helm upgrade --install rarible-app ./helm/rarible-app
```

### Access via:
```sh
minikube service rarible
# or
kubectl port-forward svc/rarible 8080:8080
```

---



