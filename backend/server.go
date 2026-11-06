package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func launchServer() {
	//dir := "../frontend/build/WebAssembly_Qt_6_7_3_single_threaded_Debug"
	dir := "../frontend/build/WebAssembly_Qt_6_7_3_single_threaded_Debug"

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsHandler)
	mux.Handle("/", http.FileServer(http.Dir(dir)))

	log.Println("Serving app at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("WebSocket request: origin=%q host=%q", r.Header.Get("Origin"), r.Host)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//log.Println("Received message:", string(message))
		buttonDecoder(string(message))
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println(err)
			return
		}
	}
}
