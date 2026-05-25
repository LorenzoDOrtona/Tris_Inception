import { useState, useEffect, useRef } from 'react';
import './App.css';

const API_BASE_URL="https://tris.lorenzodortona.com/api"
function App() {
  const [view, setView] = useState('login'); 
  const [playerName, setPlayerName] = useState('');
  const [token, setToken] = useState(null);
  
  const [gameId, setGameId] = useState(null);
  const [gameState, setGameState] = useState(null); 
  const lastTimestampRef = useRef(0);
  const [errorMessage, setErrorMessage] = useState('');
  const pollingRef = useRef(null);
  // New states for authentication
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isLoginMode, setIsLoginMode] = useState(true); // Toggle between Login and Register
  useEffect(() => {
      if (errorMessage) {
        // Set a timer to clear the message after 3 seconds (3000 ms)
        const timer = setTimeout(() => {
          setErrorMessage('');
        }, 3000);

        // Clear the timer if the component unmounts or if a new error arrives before 3 seconds
        return () => clearTimeout(timer);
      }
    }, [errorMessage]);
  // --- 1. LOGIN ---
  const handleAuth = async (e) => {
  e.preventDefault();
  
  // Decide endpoint based on mode
  const endpoint = isLoginMode ? '/login' : '/register';
  const payload = isLoginMode 
    ? { username: playerName, password: password } // Login uses username + pass
    : { username: playerName, email: email, password: password }; // Register adds email
  if (password.length < 6) {
    setErrorMessage("Password should be at least of 6 characters");
    return;
}
  try {
    const res = await fetch(`${API_BASE_URL}${endpoint}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    });

    const data = await res.json();
    if (!res.ok) throw new Error(data.error || "Auth failed");

    if (isLoginMode) {
      setToken(data.token); 
      setView('lobby'); 
    } else {
      alert("Account successfully created!");
      setIsLoginMode(true); // Switch back to login after registration
    }
  } catch (err) {
    console.error(err);
    setErrorMessage(err.message);
  }
};

  // --- 2. MATCHMAKING ---
  const handleFindGame = async (mode, customToken = null) => { // 'mode' can be "normal" or "BOT"
    const activeToken = customToken || token;
    try {
      const res = await fetch(`${API_BASE_URL}/match`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
           token: activeToken,
           game_mode: mode, 
           side_measure: 9
        })
      });
      
      const data = await res.json(); 
      if (!res.ok) throw new Error(data.error || "Matchmaking failed");

      setGameId(data.id_game);
      lastTimestampRef.current = 0;
      
      if (data.found) {
        setView('game');
      } else {
        setView('waiting');
      }
    } catch (err) {
      console.error("Matchmaking Error:", err);
      setErrorMessage(err.message);
    }
  };

  const handleDemoMatch = async () => {
    try {
      const res = await fetch(`${API_BASE_URL}/guest`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' }
      });
      const data = await res.json();
      if (!res.ok) throw new Error(data.error || "Guest login failed");

      setToken(data.token);
      setPlayerName(data.username);
      // Directly start a BOT game with the new token
      handleFindGame('BOT', data.token);
    } catch (err) {
      console.error("Demo Match Error:", err);
      setErrorMessage(err.message);
    }
  };

  
  // --- 3. POLLING ---
  const pollGameState = async () => {
    if (!token || !gameId) return; 

    try {
      const res = await fetch(`${API_BASE_URL}/polling`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            token: token,
            id_game: gameId,
            timestamp: lastTimestampRef.current
        })
      });

      if (!res.ok) return;

      const data = await res.json();

      // Check if we received a valid GameState
      if (data.game_state && data.game_state.board) {
         console.log("Updated State:", data);
         
         setGameState(data.game_state);
         
         // FUNDAMENTAL FIX:
         // Use the SERVER timestamp, not the client one!
         // Note: Check if it arrives in JSON as "LastTimestamp" or "last_timestamp"
          lastTimestampRef.current = data.game_state.last_timestamp;
         
         // View switch logic based on the 'started' flag added in the backend
         setView(prev =>
          prev === 'waiting' && data.game_state.started ? 'game' : prev
        );

      }
    } catch (err) {
      console.error("Polling Network Error:", err);
    }
  };
  // --- 4. POLLING (Effect) ---
  // This useEffect manages starting and stopping polling AUTOMATICALLY
  useEffect(() => {
    // Clean up previous timer
    if (pollingRef.current) clearInterval(pollingRef.current);

    // Conditions to start polling:
    // 1. We are in waiting or in game
    // 2. We have a valid gameId (prevents initial 400)
    if ((view === 'waiting' || view === 'game') && gameId) {
        
        pollGameState(); // Run once immediately
        pollingRef.current = setInterval(pollGameState, 1000); // Then every second
    }

    // Cleanup when component unmounts or dependencies change
    return () => {
        if (pollingRef.current) clearInterval(pollingRef.current);
    };
  }, [view, gameId, token]);


  // --- 5. MOVE ---
  const handleCellClick = async (row, col) => {
    if (!token || !gameId) return;

    // Mapping coordinates: 9x9 -> (X,Y) Macro, (I,J) Micro
    const X = Math.floor(row / 3);
    const Y = Math.floor(col / 3);
    const I = row % 3;
    const J = col % 3;

    try {
      const res = await fetch(`${API_BASE_URL}/move`, {
         method: 'POST',
         headers: { 'Content-Type': 'application/json' },
         body: JSON.stringify({
            token: token,
            id_game: gameId,
            x: X, y: Y, i: I, j: J
         })
      });
      const data = await res.json();
      if (!res.ok) {
      if (data.error_message) {
          setErrorMessage(data.error_message); 
      }
        return;
      }
      if (data.game_state) {
          setGameState(data.game_state);
          lastTimestampRef.current = data.game_state.last_timestamp;
      }
    } catch (err) {
      console.error("Move Error:", err);
    }
  };
// --- 6. RESET / BACK TO LOBBY ---
  const handleBackToLobby = () => {
    setGameId(null);
    setGameState(null);
    lastTimestampRef.current = 0; // Important reset of the timestamp
    setErrorMessage('');
    setView('lobby');
  };
  // --- RENDER ---
// Function to check if a specific cell (r, c) is among the available moves
// Corrected logic to match backend [x, y, i, j] with frontend (r, c)
const isMoveAvailable = (r, c) => {
  if (!gameState?.available_moves) return false;

  return gameState.available_moves.some(move => {
    // Assuming move is [x, y, i, j] based on your handleCellClick logic
    const [x, y, i, j] = move;
    const mappedRow = x * 3 + i;
    const mappedCol = y * 3 + j;
    
    return mappedRow === r && mappedCol === c;
  });
};

const renderBoard = () => {
  if (
    !gameState?.board ||
    !Array.isArray(gameState.board) ||
    gameState.board.length !== 9
  ) {
    return <div className="spinner" />;
  }

  return (
    <div className="super-board">
      {gameState.board.map((row, rIndex) => (
        <div key={rIndex} className="board-row">
          {row.map((cellValue, cIndex) => {
            // Check if this specific cell is playable
            const isAvailable = isMoveAvailable(rIndex, cIndex);
            return (
              <div
                key={cIndex}
                className={`cell val-${cellValue} ${
                  cIndex % 3 === 2 ? 'border-right' : ''
                } ${
                  rIndex % 3 === 2 ? 'border-bottom' : ''
                } ${isAvailable ? 'highlight-available' : ''}`} // Add CSS class if available
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
      <h1>Tris Inception</h1>
      {errorMessage && <div className="error-toast">{errorMessage}</div>}

      {view === 'login' && (
    <div className="card">
      <h2>{isLoginMode ? 'Sign In' : 'Create Account'}</h2>
      <form onSubmit={handleAuth}>
        {/* Username field (Used for both) */}
        <input 
          type="text" 
          placeholder="Username" 
          value={playerName} 
          onChange={e => setPlayerName(e.target.value)} 
          required 
        />
        
        {/* Email field (Only shown during registration) */}
        {!isLoginMode && (
          <input 
            type="email" 
            placeholder="Email" 
            value={email} 
            onChange={e => setEmail(e.target.value)} 
            required 
          />
        )}

        {/* Password field */}
        <input 
          type="password" 
          placeholder="Password" 
          value={password} 
          onChange={e => setPassword(e.target.value)} 
          required 
        />

        <button type="submit">
          {isLoginMode ? 'Login' : 'Register'}
        </button>
      </form>

      {isLoginMode && (
        <button 
          onClick={handleDemoMatch}
          className="demo-button"
          style={{ marginTop: '10px', backgroundColor: '#4caf50' }}
        >
          Try Demo Match vs Bot 🤖
        </button>
      )}

      {/* Toggle Link */}
      <p style={{ marginTop: '15px', fontSize: '0.9rem' }}>
        {isLoginMode ? "Don't have an account? " : "Already have an account? "}
        <span 
          style={{ color: '#646cff', cursor: 'pointer', textDecoration: 'underline' }}
          onClick={() => setIsLoginMode(!isLoginMode)}
        >
          {isLoginMode ? 'Create account' : 'Sign In'}
        </span>
      </p>
    </div>
  )}

      {view === 'lobby' && (
        <div className="card">
          <h2>Lobby: {playerName}</h2>
          <div className="lobby-buttons">
            <button onClick={() => handleFindGame('normal')}>Play Online</button>
            {/* BOT button*/}
            <button 
                className="bot-button" 
                onClick={() => handleFindGame('BOT')}
                style={{ marginLeft: '10px', backgroundColor: '#673ab7' }}
            >
                Play vs Bot 🤖
            </button>
          </div>
        </div>
      )}

      {view === 'waiting' && (
        <div className="card">
          <h2>Waiting for an opponent...</h2>
          <p>Game ID: {gameId}</p>
          <div className="spinner"></div>
        </div>
      )}

      {view === 'game' && (
        <div className="game-container">
          {renderBoard()}

          {/* --- WINNER OVERLAY --- */}
          {gameState && gameState.winner && gameState.winner !== "" && (
            <div className="winner-overlay">
              <div className="winner-card">
                <h2>🏆 GAMEOVER! 🏆</h2>
                <p>The winner is:</p>
                <h3>{gameState.winner}</h3>
                <button onClick={handleBackToLobby}>Go Back To Lobby</button>
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  );
}

export default App;