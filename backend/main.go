package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ydotoold")
	defer cmd.Cancel()
	go cmd.Run()
	launchServer()
	fmt.Println("Hello, World!")
}
