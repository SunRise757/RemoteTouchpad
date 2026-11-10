package main

import (
	"encoding/json"
	"log"
	"math"
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
	cmd := exec.Command("wl-copy", str)
	err := cmd.Run()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	cmd = exec.Command("ydotool", "key", "135.1", "135.0")
	err = cmd.Run()
	if err != nil {
		log.Println("Error: %v\n")
	}

}

func mouseMove(x, y int) {
	cmd := exec.Command("xdotool", "mousemove_relative", "--",
		strconv.Itoa(x*min(intAbs(x), 7)),
		strconv.Itoa(y*min(intAbs(y), 7)))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: %v\n%s", err, out)
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
		log.Printf("Error: %v\n%s", err, out)
	}
}

func mouseClick(button int) {
	cmd := exec.Command("xdotool", "click", strconv.Itoa(button))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: %v\n%s", err, out)
	}
}

func keyPress(key string) {
	cmd := exec.Command("xdotool", "keydown", key)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: %v\n%s", err, out)
	}
}

func keyRelease(key string) {
	cmd := exec.Command("xdotool", "keyup", key)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: %v\n%s", err, out)
	}
}
