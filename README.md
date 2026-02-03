# Tris Inception – Ultimate Tic-Tac-Toe Engine

**Tris Inception** is an **Ultimate Tic-Tac-Toe** engine written in **Go**, focused on building a scalable, maintainable, and testable backend. I’m particularly interested in backend development, and I study cybersecurity out of personal interest.

---

## Demo

Try it live: [https://tris-inception.onrender.com/](https://tris-inception.onrender.com/)

> Focus is on engine correctness and API behavior rather than UI polish.
> Note: On the Render free plan, there may be a short delay between submitting your name and accessing the game lobby, as the backend takes a few seconds to wake up after inactivity.

---

## Key Highlights

* Deterministic game engine for Ultimate Tic-Tac-Toe
* Terminal-based UI for local play
* HTTP API with server-side game state management
* Live backend sessions maintained on the server
* Clean architecture ready for distributed deployment

---

## Roadmap

* User accounts and authentication
* Persistent storage (game history, rankings)
* Multiple game modes (local, ranked, bot)
* Observability: metrics, logging, monitoring

---

## Design Principles

* **Separation of concerns**: Game logic fully decoupled from UI and API layers
* **Stateful backend**: Ensures consistent live gameplay
* **Explicit game states**: Improves clarity and correctness
* **Go language**: Chosen for simplicity, performance, and deployment efficiency

---

## Planned Persistence

The backend maintains game state and will handle:

* User accounts and authentication
* Game history storage
* Rankings and statistics

> This keeps game rules separate from storage, making the project maintainable and testable while supporting live sessions.
