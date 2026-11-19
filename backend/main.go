package main

import (
	"log"
	"os/exec"
)

func main() {
	checkDependencies()
	initDb()
	defer db.Close(ctx)

	cmd := exec.Command("ydotoold")
	defer cmd.Cancel()
	go cmd.Run()

	launchServer()
	log.Println("Hello, World!")
}
