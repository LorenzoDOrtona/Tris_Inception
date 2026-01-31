import { useState, useEffect, useRef } from 'react';
import './App.css';

// TODO: CHANGE THIS TO YOUR ACTUAL GO BACKEND URL
const API_BASE_URL = "http://localhost:8080"; 

function App() {
  // --- STATE MANAGEMENT ---
  const [view, setView] = useState('login'); // 'login', 'lobby', 'waiting', 'game'
  
  // User Data
  const [playerName, setPlayerName] = useState('');
  const [token, setToken] = useState(null);
  
  // Game Data
  const [gameId, setGameId] = useState(null);
  const [gameState, setGameState] = useState(null); // Corresponds to GameStateDTO
  const [lastTimestamp, setLastTimestamp] = useState(0); // For ReqPooling
  const [errorMessage, setErrorMessage] = useState('');

  // Polling Interval Reference (to clear it if needed)
  const pollingRef = useRef(null);

  // --- 1. LOGIN (Maps to ReqToken) ---
  const handleLogin = async (e) => {
    e.preventDefault();
    if (!playerName) return;

    try {
      const payload = { 
        player_name: playerName, 
        user_type: "real" 
      };

      console.log("Sending request to Go server:", payload);

      // 1. SEND REQUEST & WAIT
      // The code pauses here until the server replies.
      const res = await fetch(`${API_BASE_URL}/playerToken`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
      });

      // 2. CHECK HTTP STATUS
      // If the server returns 400 or 500, we throw an error.
      if (!res.ok) {
        throw new Error(`Server Error: ${res.status}`);
      }

      // 3. PARSE JSON
      // We extract the JSON body from the response.
      // 'data' will now contain exactly what your Go struct 'RespToken' sends.
      const data = await res.json();
      
      console.log("Received from server:", data);

      // 4. USE THE DATA
      setToken(data.token); 
      
      // 5. UPDATE UI
      setView('lobby'); 

    } catch (err) {
      console.error("Login failed:", err);
      setErrorMessage("Connection failed. Is the Go server running?");
    }
  };
  // --- 2. CREATE/JOIN GAME (Maps to ReqCreateOrJoinGame) ---
  const handleFindGame = async (mode) => {
    try {
      const payload = {
        token: token,
        game_mode: "normal", // e.g., "normal"
        side_measure: 9  // Standard 9x9 Tris Inception
      };

      console.log("Searching game with payload:", payload);

      // TODO: Uncomment when backend is ready
      
      const res = await fetch(`${API_BASE_URL}/match`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
      });
      const data = await res.json(); // Maps to RespGameCreatedOrFound
      setGameId(data.id_game);
      
      if (data.found) {
        // Joined an existing game -> Go straight to game
        setView('game');
      } else {
        // Created a new game -> Wait for opponent
        setView('waiting');
        startPolling(); // Poll to see when someone joins
      }

    } catch (err) {
      console.error("Matchmaking Error:", err);
    }
  };

  // --- 3. POLLING (Maps to ReqPooling) ---
  // This runs periodically to check for moves or new players
  // App.js

const pollGameState = async () => {
  // Se non ho token o gameId (e sono in waiting o game), non ha senso fare polling
  if (!token) return; 

  try {
    const payload = {
      token: token,
      id_game: gameId,
      timestamp: lastTimestamp
    };

    // CORREZIONE 1: Usa i backtick ` per inserire la variabile nell'URL
    // CORREZIONE 2: Usa l'endpoint corretto "/polling" (non "pooling")
    // CORREZIONE 3: Aggiungi method POST e il body
    const res = await fetch(`${API_BASE_URL}/polling`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    });

    if (!res.ok) {
        // Gestisci errori silenziosamente o logga
        return;
    }

    const data = await res.json(); // GameStateDTO

    // CORREZIONE 4: Verifica la logica di aggiornamento
    // Se il backend risponde con una board popolata (timestamp aggiornato)
    if (data.game_state && data.game_state.board && data.game_state.board.length > 0) {
       console.log("Nuovo stato ricevuto:", data);
       setGameState(data.game_state); // Assicurati di prendere il campo giusto dalla risposta
      setLastTimestamp(Math.floor(Date.now() / 1000));       
       // Se eravamo in attesa, passiamo alla schermata di gioco
       if (view === 'waiting') setView('game');
    }

  } catch (err) {
    console.error("Polling Error:", err);
  }
};

  const startPolling = async() => {
    if (pollingRef.current) clearInterval(pollingRef.current);
    pollingRef.current = setInterval(pollGameState, 500); // Poll every 0.5 seconds
  };

  // Clean up polling when component unmounts
  useEffect(() => {
    return () => clearInterval(pollingRef.current);
  }, []);


  // --- 4. MAKE MOVE (Maps to ReqMove) ---
  const handleCellClick = async (row, col) => {
    if (!token || !gameId) return;

    // Convert flat 9x9 coordinates to the Backend's 4-coordinate system
    // The backend expects: X (Big Row), Y (Big Col), I (Small Row), J (Small Col)
    const X = Math.floor(row / 3);
    const Y = Math.floor(col / 3);
    const I = row % 3;
    const J = col % 3;

    try {
      const payload = {
        token: token,
        id_game: gameId,
        x: X,
        y: Y,
        i: I,
        j: J
      };

      console.log("Sending Move:", payload);

      // TODO: Actual implementation
      
      const res = await fetch(`${API_BASE_URL}/move`, {
         method: 'POST',
         headers: { 'Content-Type': 'application/json' },

         body: JSON.stringify(payload)
      });
      const data = await res.json(); // Maps to RespMove (contains GameStateDTO)
      setGameState(data.game_state);
      
    } catch (err) {
      console.error("Move Error:", err);
    }
  };


  // --- RENDER HELPERS ---
  
  // Renders the 9x9 grid
  const renderBoard = () => {
    // If no state yet, render empty grid
    const board = gameState?.board || Array(9).fill(Array(9).fill(0));

    return (
      <div className="super-board">
        {board.map((row, rIndex) => (
          <div key={rIndex} className="board-row">
            {row.map((cellValue, cIndex) => {
              // Determine if this cell is part of an "Active" macro-board
              // Logic: You would check gameState.available_boards here
              
              return (
                <div 
                  key={cIndex} 
                  className={`cell val-${cellValue} ${cIndex % 3 === 2 ? 'border-right' : ''} ${rIndex % 3 === 2 ? 'border-bottom' : ''}`}
                  onClick={() => handleCellClick(rIndex, cIndex)}
                >
                  {cellValue === 1 ? 'X' : cellValue === 2 ? 'O' : ''}
                </div>
              );
            })}
          </div>
        ))}
      </div>
    );
  };

  return (
    <div className="app-container">
      <h1>Tris Inception (Super Tic-Tac-Toe)</h1>
      
      {errorMessage && <div className="error">{errorMessage}</div>}

      {/* VIEW: LOGIN */}
      {view === 'login' && (
        <div className="card">
          <h2>Login</h2>
          <form onSubmit={handleLogin}>
            <input 
              type="text" 
              placeholder="Enter Player Name" 
              value={playerName}
              onChange={(e) => setPlayerName(e.target.value)}
            />
            <button type="submit">Get Token</button>
          </form>
        </div>
      )}

      {/* VIEW: LOBBY */}
      {view === 'lobby' && (
        <div className="card">
          <h2>Lobby</h2>
          <p>Welcome, <strong>{playerName}</strong></p>
          <div className="options">
             <button onClick={() => handleFindGame('normal')}>Play Normal Mode</button>
             <button onClick={() => handleFindGame('no_cards')}>Play No Cards</button>
          </div>
        </div>
      )}

      {/* VIEW: WAITING */}
      {view === 'waiting' && (
        <div className="card">
          <h2>Looking for opponent...</h2>
          <div className="spinner">⌛</div>
          <p>Game ID: {gameId}</p>
          <p><small>Polling for updates...</small></p>
        </div>
      )}

      {/* VIEW: GAME */}
      {view === 'game' && (
        <div className="game-container">
          <div className="info-panel">
            <p><strong>Game ID:</strong> {gameId}</p>
            <p>Status: {gameState?.is_complete ? "Game Over" : "Playing"}</p>
          </div>
          
          {renderBoard()}
          
          <div className="debug-panel">
            {/* Helpful for debugging DTOs */}
            <small>Debug: Last Update {lastTimestamp}</small>
          </div>
        </div>
      )}
    </div>
  );
}

export default App;