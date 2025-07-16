# Rarible Go App — Kubernetes & Helm Guide

## Generating Secrets

1. Create a secret from your .env file:
   ```sh
   kubectl create secret generic rarible-secret --from-env-file=.env
   ```

2. Install or upgrade the Helm release:
   ```sh
   helm install rarible-app ./helm/rarible-app --namespace default
   # or, if already installed:
   helm upgrade rarible-app ./helm/rarible-app --namespace default
   ```

---

## Running in Kubernetes with Helm

### Requirements
- Kubernetes cluster (local or cloud)
- [Helm](https://helm.sh/) installed
- Docker (for local runs)

---

### 2. Local run with Minikube

#### 2.1. Start Minikube
```sh
minikube start
```

#### 2.2. Deploy the app with Helm
```sh
helm upgrade --install rarible-app ./helm/rarible-app
```

#### 2.3. Access the service

##### Recommended: via minikube service
```sh
minikube service rarible
# This will open the correct URL in your browser
```

##### Using port-forward
```sh
kubectl port-forward svc/rarible 8080:8080
# Now available at http://localhost:8080
```

---

### 3. Running in cloud Kubernetes (GKE, EKS, DigitalOcean)

1. Create a cluster in your cloud provider (GKE, EKS, DigitalOcean, etc.)
2. Make sure kubectl is connected to your cluster
3. Deploy the app with Helm:
   ```sh
   helm upgrade --install rarible-app ./helm/rarible-app
   ```
4. Check EXTERNAL-IP:
   ```sh
   kubectl get svc
   # EXTERNAL-IP will appear automatically
   ```
5. Access your app:
   ```sh
   curl http://<EXTERNAL-IP>:8080
   ```

---

## Makefile commands

For convenience, main actions are automated via Makefile. Here is a short description of available commands:

- **secret-generic** — creates a Kubernetes Secret from your .env file:
  ```sh
  make secret-generic
  ```
- **helm-install** — installs the Helm release (first deploy):
  ```sh
  make helm-install
  ```
- **helm-upgrade** — upgrades the existing Helm release (after changes):
  ```sh
  make helm-upgrade
  ```

> ⚡️ Just run the needed command with `make <command-name>`

---

## Troubleshooting
- If EXTERNAL-IP = <pending> in a local cluster — you need MetalLB or minikube tunnel.
- If NodePort does not work — check if the port is open in your firewall and if the IP is correct.
- For questions — open an Issue!

---

## Useful commands
- Check pod status:
  ```sh
  kubectl get pods
  ```
- View logs:
  ```sh
  kubectl logs <pod-name>
  ```
- Uninstall Helm release:
  ```sh
  helm uninstall rarible-app
  ```


## Швидкий старт

### 1. Клонування репозиторію
```sh
git clone https://github.com/FREEGREAT/Rarity-task.git
cd Rarible-task
```

---

## Запуск у Kubernetes через Helm

### Вимоги
- Kubernetes кластер (локальний або хмарний)
- [Helm](https://helm.sh/) встановлений
- Docker (для локального запуску)

---

## Генерація секретів

1. Створи секрет з .env:
   ```sh
   kubectl create secret generic rarible-secret --from-env-file=.env 
   ```

2. Встанови або онови Helm-реліз:
   ```sh
   helm install rarible-app ./helm/rarible-app --namespace default
   # або, якщо вже встановлено:
   helm upgrade rarible-app ./helm/rarible-app --namespace default
   ```

---

### 2. Локальний запуск у Minikube

#### 2.1. Запусти Minikube
```sh
minikube start
```

#### 2.2. Деплой додатку через Helm
```sh
helm upgrade --install rarible-app ./helm/rarible-app
```

#### 2.3. Отримай доступ до сервісу

##### Через minikube service (рекомендується)
```sh
minikube service rarible
# Відкриє браузер з правильним URL
```

##### З використанням Port-forward
```sh
kubectl port-forward svc/rarible 8080:8080
# Тепер доступно на http://localhost:8080
```


## Makefile

Для зручності основні дії автоматизовані через Makefile. Ось короткий опис доступних команд:

- **secret-generic** — створює Kubernetes Secret з .env файлу:
  ```sh
  make secret-generic
  ```
- **helm-install** — встановлює Helm-реліз (перший деплой):
  ```sh
  make helm-install
  ```
- **helm-upgrade** — оновлює існуючий Helm-реліз (після змін):
  ```sh
  make helm-upgrade
  ```

>



## Troubleshooting
- Якщо EXTERNAL-IP = <pending> у локальному кластері — потрібен MetalLB або minikube tunnel.
- Якщо не працює NodePort — перевір, чи порт відкритий у фаєрволі, і чи правильний IP.
- Для питань — пиши у Issues!

---

## Корисні команди
- Перевірити статус подів:
  ```sh
  kubectl get pods
  ```
- Переглянути логи:
  ```sh
  kubectl logs <pod-name>
  ```
- Видалити реліз Helm:
  ```sh
  helm uninstall rarible-app
  ``` 

---

# English Guide

## Quick Start

### 1. Clone the repository
```sh
git clone https://github.com/FREEGREAT/Rarity-task.git
cd Rarible-task
```

---

