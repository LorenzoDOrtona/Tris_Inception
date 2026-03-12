# Tris Inception – Ultimate Tic-Tac-Toe Engine 
 
**Tris Inception** is an **Ultimate Tic-Tac-Toe** engine written in **Go**, focused on building a scalable, maintainable, and testable backend.
The application is fully containerized and runs on a custom **Kubernetes (k3s)** cluster deployed on a remote VPS.

[![codecov](https://codecov.io/github/LorenzoDOrtona/Tris_Inception/graph/badge.svg?token=ZHU72H7R0G)](https://codecov.io/github/LorenzoDOrtona/Tris_Inception)

---

<img width="926" height="766" alt="image" src="https://github.com/user-attachments/assets/3c76750e-101a-4471-b80b-f2687ee87566" />

## Demo

Try it live: [https://tris.lorenzodortona.com](https://tris.lorenzodortona.com)

> Note: For a quick demo, I suggest playing against the bot due to the potential lack of online players in the lobby.

---

## ✨ Features & Infrastructure

**Game Engine**
* ~90% test coverage on core game engine logic
* Clean backend architecture (MVC + State pattern)
* Multiplayer mode and dedicated bot mode
* RESTful API for game and user interactions

**Tech Stack**
* Go backend for high concurrency and performance
* React frontend built with Vite
* PostgreSQL integration for user persistence
* User authentication system (bcrypt password hashing)

**DevOps & GitOps**
* Infrastructure as Code (IaC) provisioning via Ansible
* Fully automated GitOps CI/CD pipeline (GitHub Actions → GHCR → k3s)
* Deployed on a self-managed k3s cluster running on a remote VPS
* Encrypted secrets management via SOPS + Age
* Automated HTTPS provisioning via cert-manager and Let's Encrypt
---
## 📚 Documentation

For a deep dive into how Tris Inception is built and how to interact with it, check out the detailed documentation:

* [Architecture & Deployment](./docs/md/architecture.md) - Details on the k3s GitOps setup, Traefik routing, and the game engine's State Pattern.

---
## 🛠 Local Development

The local environment is fully containerized using Docker Compose.  
No SOPS or Kubernetes setup is required for development.

### 1️⃣ Environment Variables

Create a `.env` file inside the `backend/` directory:

```env
# backend/.env
PORT=8080
DATABASE_URL=postgres://user:password@db:5432/tris_db?sslmode=disable
JWT_SECRET=local-development-secret
FRONTEND_URL=http://localhost
ENV=development
````

> The `db` hostname matches the PostgreSQL service name defined in `docker-compose.yml`.

---

### 2️⃣ Start the Application

```bash
git clone https://github.com/LorenzoDOrtona/Tris_Inception.git
cd Tris_Inception

docker compose up --build
```

---

### 3️⃣ Access the Services

* Frontend → [http://localhost](http://localhost)
* Backend API → [http://localhost:8080/api](http://localhost:8080/api)
* PostgreSQL → running inside Docker (internal network only)

---

### ℹ️ Note on Secrets

In production, secrets are encrypted using SOPS and injected via Kubernetes Secrets.
For local development, environment variables are provided through the `.env` file.

---

## Roadmap
- [x] User accounts and authentication
- [ ] Persistent storage for matches and move history (PostgreSQL)
- [ ] JWT-based authentication (stateless tokens)
- [ ] Real-time gameplay via WebSockets (replace REST polling)
- [ ] Global ranking system (leaderboard)
- [ ] Advanced matchmaking options (friend games, custom parameters)
- [ ] Observability stack (Prometheus, Grafana, structured logging)
---

## Design Choices
- **Go for backend development**  
  Chosen for its simplicity, strong concurrency model, and ability to produce lightweight, deployable binaries suitable for containerized environments.
- **Stateful game engine**  
  Core game logic is fully server-side validated to ensure deterministic gameplay and prevent client-side manipulation.
- **REST-first approach (MVP)**  
  REST polling was selected to keep the initial architecture simple and predictable. WebSockets are planned to enable real-time bidirectional communication.
- **k3s on VPS instead of PaaS**  
  Migrated from a managed PaaS to a self-managed Kubernetes (k3s) cluster running on a VPS to gain full control over networking, deployment strategy, and scalability.
- **PostgreSQL for persistence**  
  Selected for strong consistency guarantees and relational modeling between users and future match data.
- **Encrypted secrets management (SOPS + Age)**  
  Ensures sensitive configuration never appears in plaintext within the repository or CI/CD pipelines.
- **Ingress-based routing (Traefik)**  
  Path-based routing (`/` for frontend, `/api` for backend) simplifies domain management and avoids cross-origin complexity.
