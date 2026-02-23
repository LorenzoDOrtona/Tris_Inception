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

## Accomplished
* 90% Game Logic tested
* Working MVC pattern
* Fully automated GitOps CI/CD pipeline via GitHub Actions (building to GHCR and deploying to k3s)
* Implemented REST APIs communication
* Working React frontend using Vite
* Multiplayer Mode & vs Bot mode
* Automated HTTPS certificate provisioning via cert-manager and Let's Encrypt

---

## Local Development (Docker Compose)

To run the entire stack locally for testing or development, you need to set up your environment variables and use Docker Compose.

### 1. Environment Setup
The backend requires a `.env` file to run. Create a file named `.env` inside the `backend/` directory and populate it with your local settings. You can use the following template:

```env
# backend/.env
PORT=8080
DATABASE_URL=postgres://user:password@db:5432/tris_db
JWT_SECRET=local-development-secret
FRONTEND_URL=http://localhost
ENV=development

```

Once the environment variables are set, you can start the containers:
```bash
# Clone the repository
git clone https://github.com/LorenzoDOrtona/Tris_Inception.git
cd Tris_Inception

# Start both frontend and backend
docker-compose up --build

```

* The frontend will be available at `http://localhost`
* The backend API will be available at `http://localhost:8080`

---

## Roadmap

* User accounts and authentication
* Persistent storage (game history, rankings) using PostgreSQL
* Multiple game searching parameters (Play against a friend, Specific game mode, etc.)
* Observability: metrics statistics, logging, and monitoring of server resources

---

## Design Choices

* **Go language**: Chosen for simplicity, performance, and deployment efficiency.
* **Stateful backend**: Ensures consistent live gameplay and strict server-side validation.
* **REST APIs Polling**: Used for the MVP for simplicity, paving the way for future WebSockets implementation for real-time updates.
* **Infrastructure**: Migrated from PaaS (Render) to a self-hosted **Kubernetes (k3s)** environment to maximize control over deployment, scaling, and networking (via **Traefik** Ingress).
