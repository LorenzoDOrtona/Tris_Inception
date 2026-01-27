import { useState } from 'react'

function App() {
  const [token, setToken] = useState("")

  const prendiToken = async () => {
    try {
      const response = await fetch('http://localhost:8080/playerToken', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ 
            player_name: "ReactPlayer",
            user_type: "human" 
        }),
      });
      
      const data = await response.json();
      console.log(data); // Guarda la console del browser (F12)
      setToken(data.response.token); // Salva il token nello stato
      
    } catch (error) {
      console.error("Errore:", error);
    }
  }

  return (
    <div>
      <h1>Tris Inception</h1>
      <button onClick={prendiToken}>Login e Prendi Token</button>
      {token && <p>Il tuo Token è: {token}</p>}
    </div>
  )
}

export default App