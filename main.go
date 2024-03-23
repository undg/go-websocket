package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page hit by clien")

	fmt.Fprintf(w, "Welcome on 'home' endpoint")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ws Page hit by client")

	fmt.Fprintf(w, "Welcome on 'ws' endpoint")
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
