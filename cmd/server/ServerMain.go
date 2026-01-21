package server
import(
	"fmt"
	"log"       // <-- Devi importare il pacchetto log per usare log.Fatal/Println
	"net/http"  // <-- Devi importare il pacchetto net/http per usare http.HandleFunc/ListenAndServe
	"github.com/gorilla/websocket"
)
// Richiediamo la libreria esterna.
// (import "github.com/gorilla/websocket")
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    // Questo è fondamentale: permette connessioni da qualsiasi origine (CORS).
    // Dovrebbe essere più restrittivo in produzione!
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}
func wsHandler(w http.ResponseWriter, r *http.Request) {
    // 1. CHIAMATA ALL'UPGRADE
    // Trasforma la richiesta HTTP (w, r) in una connessione WebSocket (*conn).
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Errore nell'upgrade:", err)
        return
    }
    // Una volta qui, l'handshake ha avuto successo (risposta 101 Switching Protocols).
    defer conn.Close() // Assicuriamoci di chiudere la connessione quando la funzione termina.

    log.Println("Nuovo client connesso!")

    // 2. LOOP DI COMUNICAZIONE PERSISTENTE
    // In un progetto reale, questo sarebbe in una goroutine separata.
    for {
        // Leggi il prossimo messaggio (frame) dal client.
        // mt (message type): 1=Testo, 2=Binario.
        // p (payload): I dati del messaggio.
        mt, p, err := conn.ReadMessage()
        if err != nil {
            log.Println("Errore di lettura (client disconnesso?):", err)
            break // Esci dal loop se c'è un errore
        }

        // STAMPA E RISPONDI (Eco Server)
        fmt.Printf("Ricevuto: %s\n", p)

        // Scrivi un messaggio di risposta al client (funziona come un "Eco").
        // Si usa lo stesso tipo di messaggio (mt) ricevuto.
        if err := conn.WriteMessage(mt, []byte("Server ricevuto: "+string(p))); err != nil {
            log.Println("Errore di scrittura:", err)
            break
        }
    }
}
func main() {
    // Collega l'URL "/ws" al gestore WebSocket.
    http.HandleFunc("/ws", wsHandler)

    log.Println("Server in ascolto su :8080")
    // Avvia il server HTTP.
    log.Fatal(http.ListenAndServe(":8080", nil))
}
