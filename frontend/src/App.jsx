import { useState, useEffect, useRef } from 'react';
import './App.css';

const API_BASE_URL = "http://localhost:8080"; 

function App() {
  const [view, setView] = useState('login'); 
  const [playerName, setPlayerName] = useState('');
  const [token, setToken] = useState(null);
  
  const [gameId, setGameId] = useState(null);
  const [gameState, setGameState] = useState(null); 
  const [lastTimestamp, setLastTimestamp] = useState(0); 
  const [errorMessage, setErrorMessage] = useState('');

  const pollingRef = useRef(null);

  // --- 1. LOGIN ---
  const handleLogin = async (e) => {
    e.preventDefault();
    if (!playerName) return;
    try {
      const res = await fetch(`${API_BASE_URL}/playerToken`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ player_name: playerName, user_type: "real" })
      });
      if (!res.ok) throw new Error(`Status: ${res.status}`);
      const data = await res.json();
      setToken(data.token); 
      setView('lobby'); 
    } catch (err) {
      console.error(err);
      setErrorMessage("Errore Login");
    }
  };

  // --- 2. MATCHMAKING ---
  const handleFindGame = async (mode) => {
    try {
      const res = await fetch(`${API_BASE_URL}/match`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
           token: token,
           game_mode: "normal",
           side_measure: 9
        })
      });
      
      const data = await res.json(); 
      setGameId(data.id_game); // Questo è asincrono!
      setLastTimestamp(0);     // Resetta timestamp per la nuova partita
      
      if (data.found) {
        setView('game');
      } else {
        setView('waiting');
      }
      // NOTA: Non chiamiamo startPolling() qui manualmente per evitare bug di stato
    } catch (err) {
      console.error("Matchmaking Error:", err);
    }
  };

  // --- 3. POLLING (Funzione) ---
  // --- 3. POLLING (Funzione Corretta) ---
  const pollGameState = async () => {
    if (!token || !gameId) return; 

    try {
      const res = await fetch(`${API_BASE_URL}/polling`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            token: token,
            id_game: gameId, 
            timestamp: lastTimestamp
        })
      });

      if (!res.ok) return;

      const data = await res.json();

      // Verifica se abbiamo ricevuto un GameState valido
      if (data.game_state && data.game_state.board) {
         console.log("Stato Aggiornato:", data);
         
         setGameState(data.game_state);
         
         // CORREZIONE FONDAMENTALE:
         // Usa il timestamp del SERVER, non quello del client!
         // Nota: Controlla se nel JSON ti arriva come "LastTimestamp" o "last_timestamp"
         setLastTimestamp(data.game_state.last_timestamp); 
         
         // Logica di switch view basata sul flag 'started' che hai aggiunto nel backend
         if (view === 'waiting' && data.game_state.started) {
             setView('game');
         }
      }
    } catch (err) {
      console.error("Polling Network Error:", err);
    }
  };
  // --- 4. POLLING (Effect) ---
  // Questo useEffect gestisce l'avvio e l'arresto del polling AUTOMATICAMENTE
  useEffect(() => {
    // Pulisci timer precedente
    if (pollingRef.current) clearInterval(pollingRef.current);

    // Condizioni per avviare il polling:
    // 1. Siamo in waiting o in game
    // 2. Abbiamo un gameId valido (evita il 400 iniziale)
    if ((view === 'waiting' || view === 'game') && gameId) {
        
        pollGameState(); // Esegui subito una volta
        pollingRef.current = setInterval(pollGameState, 1000); // Poi ogni secondo
    }

    // Cleanup quando il componente smonta o le dipendenze cambiano
    return () => {
        if (pollingRef.current) clearInterval(pollingRef.current);
    };
  }, [view, gameId, lastTimestamp]); // Riavvia se cambia view o gameId


  // --- 5. MOVE ---
  const handleCellClick = async (row, col) => {
    if (!token || !gameId) return;

    // Mapping coordinate: 9x9 -> (X,Y) Macro, (I,J) Micro
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
      if (data.game_state) {
          setGameState(data.game_state);
          setLastTimestamp(Math.floor(Date.now() / 1000));
      } else if (data.error) {
          console.warn(data.error);
      }
    } catch (err) {
      console.error("Move Error:", err);
    }
  };

  // --- RENDER ---
  const renderBoard = () => {
    const board = gameState?.board || Array(9).fill(Array(9).fill(0));
    return (
      <div className="super-board">
        {board.map((row, rIndex) => (
          <div key={rIndex} className="board-row">
            {row.map((cellValue, cIndex) => (
                <div 
                  key={cIndex} 
                  className={`cell val-${cellValue} ${cIndex % 3 === 2 ? 'border-right' : ''} ${rIndex % 3 === 2 ? 'border-bottom' : ''}`}
                  onClick={() => handleCellClick(rIndex, cIndex)}
                >
                  {cellValue === 1 ? 'X' : cellValue === 2 ? 'O' : ''}
                </div>
            ))}
          </div>
        ))}
      </div>
    );
  };

  return (
    <div className="app-container">
      <h1>Tris Inception</h1>
      {errorMessage && <div className="error">{errorMessage}</div>}

      {view === 'login' && (
        <div className="card">
          <h2>Login</h2>
          <form onSubmit={handleLogin}>
            <input type="text" placeholder="Nome" value={playerName} onChange={e => setPlayerName(e.target.value)} />
            <button type="submit">Entra</button>
          </form>
        </div>
      )}

      {view === 'lobby' && (
        <div className="card">
          <h2>Lobby: {playerName}</h2>
          <button onClick={() => handleFindGame('normal')}>Gioca</button>
        </div>
      )}

      {view === 'waiting' && (
        <div className="card">
          <h2>In attesa...</h2>
          <p>Game ID: {gameId}</p>
          <div className="spinner">⌛</div>
        </div>
      )}

      {view === 'game' && (
        <div className="game-container">
          {renderBoard()}
        </div>
      )}
    </div>
  );
}

export default App;