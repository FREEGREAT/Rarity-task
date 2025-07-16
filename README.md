
# 🧩 Rarible Go App — Kubernetes + Helm Deployment Guide

---

## 📦 Quick Start 


## 🚀 Running in Kubernetes with Helm

### 🔧 Requirements

- Kubernetes cluster (local or cloud)
- [Helm](https://helm.sh/) installed
- Docker (for local runs)
- (Optional) Minikube for local testing

---


### 1. Clone the repository
```sh
git clone https://github.com/FREEGREAT/Rarity-task.git
cd Rarible-task
```

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

3. Local Run with Minikube
### Access via:
```sh
minikube service rarible
# or
kubectl port-forward svc/rarible 8080:8080
```

---

# 🧪 TEST_TASK API Collection

This repository contains the `TEST_TASK` Postman collection for local testing of the NFT service.

## 📦 Collection Contents

The collection includes two main requests:

---

### 🔍 1. GET `/nft/ownerships/{id}`

Retrieve ownership information for a specific NFT.

#### 📥 Request:
GET http://localhost:8080/nft/ownerships/ETHEREUM:0xb66a603f4cfe17e3d27b87a8bfcad319856518b8:32292934596187112148346015918544186536963932779440027682601542850818403729410:0x4765273c477c2dc484da4f1984639e943adccfeb

#### 📌 Parameters:
- `id` — composite NFT identifier in the format:  
  BLOCKCHAIN:CONTRACT:TOKEN_ID:OWNER_ADDRESS

---

### 🧮 2. POST `/nft/trait-rarities`

Calculate rarity based on provided NFT properties in a collection.

#### 📥 Request:
POST http://localhost:8080/nft/trait-rarities

#### 🧾 Request body:
{
  "collectionId": "ETHEREUM:0x60e4d786628fea6478f785a6d7e704777c86a7c6",
  "properties": [
    { "key": "Hat", "value": "Halo" },
    { "key": "Color", "value": "Red" }
  ]
}

---

## 🚀 How to Use

1. Import the `.json` file of this collection into Postman
2. Use the provided requests to test the API






