package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/gorilla/websocket"
)

func mouseMove(x, y int) {
	cmd := exec.Command("xdotool", "mousemove_relative", strconv.Itoa(x), strconv.Itoa(y))
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n%s", err, out)
	} else {
		fmt.Printf("output: %s\n", out)
	}
}

func keyPress(key string) {
	cmd := exec.Command("xdotool", "keydown", key)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n%s", err, out)
	} else {
		fmt.Printf("output: %s\n", out)
	}
}

func keyRelease(key string) {
	cmd := exec.Command("xdotool", "keyup", key)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n%s", err, out)
	} else {
		fmt.Printf("output: %s\n", out)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Development example. In production, allow only your app's
		// expected origin(s).
		return true
		return r.Header.Get("Origin") == "http://localhost:8080"
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
		fmt.Println("Received message:", string(message))
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	fmt.Println("Hello, World!")

	dir := "../frontend/build/WebAssembly_Qt_6_7_3_single_threaded_Debug"

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsHandler)
	mux.Handle("/", http.FileServer(http.Dir(dir)))

	log.Println("Serving app at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

	fmt.Println("Hello, World!")
	/*
	   mouseMove(100, 100)

	   keyPress("Super_L")
	   keyPress("1")
	   keyRelease("Super_L")
	   keyRelease("1")
	*/
}

//dependency: xdtool (linux)
