# Tris Inception – Ultimate Tic-Tac-Toe Engine 

**Tris Inception** is an **Ultimate Tic-Tac-Toe** engine written in **Go**, focused on building a scalable, maintainable, and testable backend.
The backend and React frontend are currently deployed separately on Render’s free hosting plan and communicate via REST APIs.

---

## Demo

Try it live: [https://tris-inception.onrender.com/](https://tris-inception.onrender.com/)

> Note: For a demo, I suggest playing against a bot due to the lack of online players.
---

## Roadmap

* User accounts and authentication
* Persistent storage (game history, rankings)
* Multiple game modes (local, ranked, bot)
* Observability: metrics, logging, monitoring

---

## Design Choices

* **Go language**: Chosen for simplicity, performance, and deployment efficiency
* **Stateful backend**: Ensures consistent live gameplay
* **REST-APIs Polling**: Used for the MVP for simplicity, even though WebSockets would be better for real-time updates
* **MVC pattern**: Choosen for better separation on concerns
