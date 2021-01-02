package main

import "fmt"

var (
	playerName    string  = "Andre Agassi"
	sponser       string  = "Head"
	worldRanking  int     = 1
	stringTension float32 = 65.5
)

var (
	matchesPlayed int = 0
)

func main() {
	fmt.Println(playerName, sponser, worldRanking, stringTension, matchesPlayed)
}
