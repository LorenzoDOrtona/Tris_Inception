# Tris Inception – Ultimate Tic-Tac-Toe Engine 

**Tris Inception** is an **Ultimate Tic-Tac-Toe** engine written in **Go**, focused on building a scalable, maintainable, and testable backend.
The backend and React frontend are currently deployed separately on Render’s free hosting plan and communicate via REST APIs.

---
<img width="926" height="766" alt="image" src="https://github.com/user-attachments/assets/3c76750e-101a-4471-b80b-f2687ee87566" />

## Demo

Try it live: [https://tris-inception.onrender.com/](https://tris-inception.onrender.com/)

> Note: For a demo, I suggest playing against a bot due to the lack of online players.
---
## Accomplished
* Working MVC pattern
* Implemented Rest APIs communication
* Working React frontend
* Multiplayer Mode, vs Bot mode
---
## Roadmap
* User accounts and authentication
* Persistent storage (game history, rankings)
* Multiple game searching parameters (Play against a friend, Specific game mode, etc.)
* GitHub Actions to run game tests before every production deploy
* Observability: metrics statistics, logging, monitoring of server resources

---

## Design Choices

* **Go language**: Chosen for simplicity, performance, and deployment efficiency
* **Stateful backend**: Ensures consistent live gameplay
* **REST APIs Polling**: Used for the MVP for simplicity, even though WebSockets would be better for real-time updates
* **MVC pattern**: Chosen for better separation of concerns
