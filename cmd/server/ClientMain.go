package main
import(
	"fmt"
	"log"       // <-- Devi importare il pacchetto log per usare log.Fatal/Println
	"github.com/gorilla/websocket" 
)
// Richiediamo la libreria esterna.
// (import "github.com/gorilla/websocket")
var dialer = websocket.DefaultDialer
func startClient(url string) {
    // 1. CHIAMATA AL DIAL
    // Tenta l'handshake con l'URL fornito (es. "ws://localhost:8080/ws")
    conn, _, err := dialer.Dial(url, nil)
    if err != nil {
        log.Fatal("Errore nel Dial (Connessione fallita):", err)
        return
    }
    log.Println("Connessione WebSocket stabilita con successo!")
    defer conn.Close() // Assicurati che la connessione sia chiusa quando la funzione termina

    // 2. INVIO MESSAGGIO INIZIALE (Client Push)
    // Inviamo subito il primo messaggio di testo (websocket.TextMessage)
    message := []byte("Ciao, sono il client Go!")
    err = conn.WriteMessage(websocket.TextMessage, message)
    if err != nil {
        log.Println("Errore nell'invio:", err)
        return
    }

    // 3. LOOP DI RICEZIONE (Ascolto del Server)
    // Il client deve essere costantemente in ascolto per i dati del server.
    for {
        mt, p, err := conn.ReadMessage()
        if err != nil {
            log.Println("Errore di lettura (Server disconnesso?):", err)
            break
        }

        // 4. ELABORAZIONE DEL MESSAGGIO
        if mt == websocket.TextMessage {
            fmt.Printf("Client Ricevuto (Testo): %s\n", p)
        } else {
            fmt.Printf("Client Ricevuto (Binario) di %d byte\n", len(p))
        }
    }

}
func startMenuInputRoutine(){

}
func main() {	
    serverURL := "ws://localhost:8080/ws"
		m:=selectMode()
		if m=="1"{
		}
		if m=="2"{
			startClient(serverURL)
			go startMenuInputRoutine()
		}
	
}
func selectMode() string{
	var m string
	fmt.Println("Seleziona la modalità:")
	fmt.Println("1- 1v1 local ")
	fmt.Println("2- Online")
	m= "1"
	return m	
}
