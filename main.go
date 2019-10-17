package main

import (
	"fmt"
	"net/http"

	"github.com/SwanHub/chat-app-backend/backend/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Distributed Chat System Over the Airwaves")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
