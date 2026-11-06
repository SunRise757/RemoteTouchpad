package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/gorilla/websocket"
)

func mouseMove(x, y int) {
	cmd := exec.Command("xdotool", "mousemove_relative", "--", strconv.Itoa(x), strconv.Itoa(y))
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n%s", err, out)
	} else {
		fmt.Printf("output: %s\n", out)
	}
}

func mouseScroll(x int) {
	button := "5"
	if x > 0 {
		button = "4"
	}
	math.Abs(float64(x))

	cmd := exec.Command("xdotool", "click", "--repeat", strconv.Itoa(int((math.Abs(float64(x))))), "--delay", "0", button)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n%s", err, out)
	} else {
		fmt.Printf("output: %s\n", out)
	}
}

func mouseClick(button int) {
	cmd := exec.Command("xdotool", "click", strconv.Itoa(button))
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
		buttonDecoder(string(message))
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println(err)
			return
		}
	}
}

func buttonDecoder(message string) {
	var data map[string]any
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		log.Println(err)
		return
	}

	msgType, ok := data["type"].(string)
	if ok {
		switch msgType {
		case "mouse":
			fmt.Println("mouse")
			x, y := data["dx"].(float64), data["dy"].(float64)
			mouseMove(int(x), int(y))
		case "scroll":
			fmt.Println("scroll")
			y := data["dy"].(float64)
			mouseScroll(int(y))
		case "button":
			fmt.Println("button")
			switch data["name"] {
			case "buttonLeftClick":
				mouseClick(1)
			case "buttonRightClick":
				mouseClick(2)
			case "buttonMiddleClick":
				mouseClick(3)
			}
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
