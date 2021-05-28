package main

import "log"

func main() {
	game, err := NewGame()
	if err != nil {
		log.Fatalf("new game error: %v", err)
	}

	game.Run()
}
