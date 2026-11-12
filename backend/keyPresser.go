package main

import (
	"encoding/json"
	"errors"
	"log"
	"os/exec"
	"strconv"
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
	cmd := exec.Command("cb", "copy", str)
	err := cmd.Run()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	cmd = exec.Command("ydotool", "key", "135:1", "135:0")
	err = cmd.Run()
	if err != nil {
		log.Println("Error: %v\n")
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
	err := exec.Command("ydotool").Run()
	if errors.Is(err, exec.ErrNotFound) {
		log.Fatal("ydotool not found. Get it here: https://github.com/ReimuNotMoe/ydotool")
	}
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("cb").Run()
	if errors.Is(err, exec.ErrNotFound) {
		log.Fatal("cb not found. Get it here: https://github.com/Slackadays/Clipboard")
	}
	if err != nil {
		log.Fatal(err)
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
