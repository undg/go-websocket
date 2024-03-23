package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {

		// Read messages
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Print for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page hit by clien")

	fmt.Fprintf(w, "Welcome on 'home' endpoint")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ws Page hit by client")

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade to WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Client connected")

	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)

	fmt.Println("Router ready!")
}

func main() {
	fmt.Println("Setup router...")

	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
