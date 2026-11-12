package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func intAbs(num int) int {
	if num > 0 {
		return num
	} else {
		return -num
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
			x, y := data["dx"].(float64), data["dy"].(float64)
			mouseMove(int(x), int(y))
		case "scroll":
			y := data["dy"].(float64)
			mouseScroll(int(y))
		case "button":
			switch data["name"] {
			case "buttonLeftClick":
				mouseClick(1)
			case "buttonMiddleClick":
				mouseClick(2)
			case "buttonRightClick":
				mouseClick(3)
			}
		case "string":
			str := data["string"].(string)
			typeString(str)
		case "key":
			log.Println(data)
		}
	}
}

func typeString(str string) {
	var err error
	switch strings.ToLower(os.Getenv("XDG_SESSION_TYPE")) {
	case "wayland":
		err = exec.Command("wl-copy", str).Run()
	case "x11":
		cmd := exec.Command("xclip", "-selection", "clipboard")
		cmd.Stdin = strings.NewReader(str)
		err = cmd.Run()
	default:
		return
	}
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	err = exec.Command("ydotool", "key", "29:1", "42:1", "47:1", "47:0", "42:0", "29:0").Run()
	if err != nil {
		log.Printf("ydotool error: %v", err)
	}

}

func mouseMove(x, y int) {
	cmd := exec.Command("ydotool", "mousemove", "--",
		strconv.Itoa(x*min(intAbs(x), 7)),
		strconv.Itoa(y*min(intAbs(y), 7)))
	err := cmd.Run()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func mouseScroll(x int) {
	cmd := exec.Command("ydotool", "mousemove", "--wheel", "--",
		"0",
		strconv.Itoa(x*min(intAbs(x), 7)))
	err := cmd.Run()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func mouseClick(button int) {
	m := map[int]string{
		1: "0x110",
		2: "0x112",
		3: "0x111",
	}
	cmd := exec.Command("ydotool", "key", m[button]+":1", m[button]+":0")
	err := cmd.Run()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func checkDependencies() {
	cmdYdotool := exec.Command("ydotool")

	err := cmdYdotool.Run()
	if errors.Is(err, exec.ErrNotFound) {
		log.Fatal("ydotool not found. Get it here: https://github.com/ReimuNotMoe/ydotool")
	}
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(os.Getenv("XDG_SESSION_TYPE")) {
	case "wayland":
		err = exec.Command("wl-copy").Run()
		if errors.Is(err, exec.ErrNotFound) {
			log.Fatal("wl-copy not found")
		}
	case "x11":
		err = exec.Command("xclip").Run()
		if errors.Is(err, exec.ErrNotFound) {
			log.Fatal("xclip not found")
		}
	default:
		return
	}
}

/*
func keyPress(key string) {
	cmd := exec.Command("ydotool", "key", key + ":1")
	err := cmd.Run()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}}
}

func keyRelease(key string) {
	cmd := exec.Command("ydotool", "key", key + ":0")
	err := cmd.Run()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}
