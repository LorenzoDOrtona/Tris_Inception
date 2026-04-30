package main
import(
	"fmt"
	"log"       // <-- You must import the log package to use log.Fatal/Println
	"github.com/gorilla/websocket" 
)
// External library requirement.
// (import "github.com/gorilla/websocket")
var dialer = websocket.DefaultDialer
func startClient(url string) {
    // 1. DIAL CALL
    // Attempts handshake with the provided URL (e.g., "ws://localhost:8080/ws")
    conn, _, err := dialer.Dial(url, nil)
    if err != nil {
        log.Fatal("Dial error (Connection failed):", err)
        return
    }
    log.Println("WebSocket connection established successfully!")
    defer conn.Close() // Ensure the connection is closed when the function terminates

    // 2. INITIAL MESSAGE SENDING (Client Push)
    // Send the first text message immediately (websocket.TextMessage)
    message := []byte("Hello, I am the Go client!")
    err = conn.WriteMessage(websocket.TextMessage, message)
    if err != nil {
        log.Println("Error in sending:", err)
        return
    }

    // 3. RECEPTION LOOP (Server Listening)
    // The client must be constantly listening for server data.
    for {
        mt, p, err := conn.ReadMessage()
        if err != nil {
            log.Println("Reading error (Server disconnected?):", err)
            break
        }

        // 4. MESSAGE PROCESSING
        if mt == websocket.TextMessage {
            fmt.Printf("Client Received (Text): %s\n", p)
        } else {
            fmt.Printf("Client Received (Binary) of %d bytes\n", len(p))
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
	fmt.Println("Select mode:")
	fmt.Println("1- 1v1 local ")
	fmt.Println("2- Online")
	m= "1"
	return m	
}
