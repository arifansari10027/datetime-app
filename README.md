# Datetime App - GoLang + Kubernetes

This project is a simple GoLang web application that returns the current date and time.
It is containerized using Docker and deployed using Kubernetes.

---

## How to Run this Project on Your Machine

### 1. Prerequisites

Make sure you have installed:

- [Docker](https://www.docker.com/products/docker-desktop/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/) **or** access to [Play With Kubernetes (PWK)](https://labs.play-with-k8s.com/)
- [Git](https://git-scm.com/)

---

### 2. Clone the Repository

```bash
git clone https://github.com/<your-username>/datetime-app.git
cd datetime-app
```

---

### 3. Build and Push Docker Image

**Build the Docker Image:**

```bash
docker build -t <your-dockerhub-username>/datetime-app:latest .
```

**Login to Docker Hub:**

```bash
docker login
```

**Push the Docker Image:**

```bash
docker push <your-dockerhub-username>/datetime-app:latest
```

---

### 4. Set Up Kubernetes

#### Option 1: Minikube (Local Machine)

- Start Minikube:

  ```bash
  minikube start
  ```

- (If using LoadBalancer) Start Tunnel:

  ```bash
  minikube tunnel
  ```

#### Option 2: Play With Kubernetes (PWK)

- Visit: [https://labs.play-with-k8s.com/](https://labs.play-with-k8s.com/)
- Start a session and create a Kubernetes cluster (master node).
- Confirm Kubernetes is running:

  ```bash
  kubectl cluster-info
  ```

---

### 5. Deploy the Application

**Apply Deployment:**

```bash
kubectl apply -f deployment.yaml
```

**Apply Service:**

```bash
kubectl apply -f service.yaml
```

**Check Status:**

```bash
kubectl get pods
kubectl get services
```

---

### 6. Access the Application

#### If using NodePort:

- Get your **Node IP**:
  - **Minikube:**  
    ```bash
    minikube ip
    ```
  - **PWK:**  
    Public IP is visible on the dashboard.

- Access the app at:

  ```
  http://<NODE-IP>:<NODE-PORT>
  ```

(Example: `http://192.168.49.2:30007`)

---


## 📚 Author

- [Arif Ansari](https://github.com/arifansari10027>)