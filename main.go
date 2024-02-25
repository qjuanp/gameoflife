package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Project setup done")
	game := NewGameOfSize(100, 100)

	clear()
	for {
		fmt.Print(game.toString())
		time.Sleep(time.Second)
		game = game.next()
	}
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
