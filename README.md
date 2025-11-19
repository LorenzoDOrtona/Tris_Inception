# Ultimate Tic-Tac-Toe Engine  
[Tris_Inception](https://github.com/LorenzoDOrtona/Tris_Inception)

## 🔍 Overview  
This is the backend engine for an **Ultimate Tic-Tac-Toe** variant, developed in Go. The project aims to build a distributed gaming service with containerized deployment, orchestration via Kubernetes, and a strong focus on performance, scalability and developer experience.

## 🚀 Key Features  
- Written in **Go**, leveraging its concurrency model for efficient game state management.  
- Fully containerized with **Docker**, enabling consistent environments and easy deployment.  
- Orchestrated using **Kubernetes**, designed to scale and handle multiple simultaneous game sessions.  
- Modular architecture: core game logic separated from infrastructure concerns.  
- Clean code practices: maintainable, testable, and built with refactoring in mind.

## 🧱 Architecture & Components  
- **server/** – The main game engine: handles routing, game state machine, logic for Ultimate Tic-Tac-Toe.  
- **internal/model/board**, **internal/model/positionable** etc – Domain packages implementing the board structure and game rules.  
- **cmd/** – Entrypoint for the service, sets up HTTP handlers, WebSocket connections, and orchestrates game sessions.  
- **docker/** – Dockerfiles and compose files for local development and testing environments.  
- **k8s/** – Kubernetes manifests for deployment, service, scaling, and monitoring.  
- **docs/** – UML diagrams, architecture sketches, and design decisions (e.g., ProjectZ.drawio).

## 🛠 Tech Stack  
- **Languages**: Go  
- **Containers**: Docker  
- **Orchestration**: Kubernetes  
- **Testing**: Go’s built-in testing framework (unit & integration)  
- **CI/CD**: (planned) GitHub Actions for build, test, and deployment pipelines  
- **Monitoring/Logging**: (future) Prometheus + Grafana integration planned

## ✅ Why This Project Matters  
- Demonstrates system design skills: distributed backend, concurrency, container orchestration.  
- Shows hands-on use of modern cloud/devops tools—critical for backend and cybersecurity roles.  
- Offers tangible results: the service can be stress-tested, scaled, and observed in production-like environments.

## 📂 Getting Started  
1. Clone the repo:  
   ```bash
   git clone https://github.com/LorenzoDOrtona/Tris_Inception.git
   cd Tris_Inception
