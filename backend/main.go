package main

import (
	"fmt"
	"os/exec"
	"strconv"
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

func main() {
	fmt.Println("Hello, World!")

	mouseMove(100, 100)

	keyPress("Super_L")
	keyPress("1")
	keyRelease("Super_L")
	keyRelease("1")
}

//dependency: xdtool (linux)
