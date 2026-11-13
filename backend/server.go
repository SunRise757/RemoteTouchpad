package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/websocket"
)

func launchServer() {
	path, _ := os.Executable()
	dir := filepath.Dir(path) + "/mobile"

	mux := http.NewServeMux()
	mux.HandleFunc("/register", registerHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/protected", protectedRouteHandler)
	mux.Handle("/", http.FileServer(http.Dir(dir)))

	log.Println("Serving app at " + getLanIp() + ":8080")

	log.Fatal(http.ListenAndServe(":8080", mux))

}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "invalid request method", err)
		return
	}

	if err := Authorize(r); err != nil {
		err := http.StatusUnauthorized
		http.Error(w, "Unauthorized", err)
		return
	}

	log.Printf("WebSocket request: origin=%q host=%q", r.Header.Get("Origin"), r.Host)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}
		buttonDecoder(string(message))

	}
}
