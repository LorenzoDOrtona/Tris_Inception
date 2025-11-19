# Tris Inception — Tris Inception (Ultimate Tic-Tac-Toe Engine)

Tris Inception is an **Ultimate Tic-Tac-Toe** engine written in **Go**.

The goal of the project is to build a clean, extensible core for the game — starting from a solid domain model and a terminal UI — and later evolve it into a distributed, containerized service (with HTTP/WebSocket API, Docker, Kubernetes, monitoring, etc.).

At the current stage, the focus is mainly on:

- **game modelling** (boards, moves, rules),
- a **terminal user interface (TUI)** for local play / debugging,
- structuring the codebase so it can grow into a real backend service.

---

## 🎯 Goals

- Model Ultimate Tic-Tac-Toe in a clear, testable way.
- Provide a TUI layer to interact with the engine from the terminal.
- Keep the architecture modular so a future server (HTTP/WebSocket) can be added without rewriting the core.
- Use this project as a learning playground for:
  - Go best practices,
  - clean architecture / domain modelling,
  - (later) containers, orchestration, monitoring.

---

## 🧱 Project Structure

> Names may evolve as the project grows, but the high-level idea is:

- `cmd/server/`  
  Entrypoint of the application (the `main` package). Wires everything together and runs the current “game loop” / interaction.

- `internal/model/`  
  Core **domain model** of the game:
  - board representation,
  - cells / positions,
  - players,
  - rules for valid moves,
  - win / draw logic for both small boards and the global board.

- `internal/tui/`  
  **Terminal UI** components:
  - rendering of the board(s) in the terminal,
  - user input handling,
  - feedback / messages.

- `uml/` & `*.drawio`  
  Diagrams and design sketches used to reason about:
  - commands flow,
  - connections between components,
  - overall architecture of Project Z.

- `go.mod`, `go.sum`  
  Go module definition and dependencies.

---

## 🚦 Status & Roadmap

### ✅ Already implemented / in progress

- [x] Go module setup
- [x] Core game model under `internal/model`
- [x] Basic terminal UI under `internal/tui`
- [x] Main entrypoint in `cmd/server`
- [x] First versions of UML / architecture diagrams

### 🔜 Planned / not implemented yet

- [ ] Full Ultimate Tic-Tac-Toe rules validation (edge cases, illegal moves, etc.)
- [ ] Clean error handling and logging
- [ ] HTTP and/or WebSocket API to expose the game engine
- [ ] Dockerfile and containerized dev environment
- [ ] Kubernetes manifests for deployment and scaling
- [ ] Simple AI/bot opponent
- [ ] Metrics & monitoring (Prometheus/Grafana or similar)
- [ ] Automated tests with coverage tracking
- [ ] CI pipeline (GitHub Actions) for build + test

The idea is to iterate step by step and use this as a long-term learning project rather than a one-off script.

---

## 🛠 Tech Stack

- **Language**: Go
- **Architecture**: layered (model + UI), with clear separation between domain logic and presentation
- **Tooling**: Go modules, UML/Draw.io for design

In future iterations:

- **Containers**: Docker (planned)
- **Orchestration**: Kubernetes (planned)
- **CI/CD**: GitHub Actions (planned)

---

## 🚀 Getting Started

> Commands may change as the project evolves; this is the current expected flow.

### 1. Clone the repository

```bash
git clone https://github.com/LorenzoDOrtona/Tris_Inception.git
cd Tris_Inception
