package main

import "fmt"

//uppercase: globally visible
var I int = 27

// lowercase: visibile to the package only
var (
	playerName    string  = "Andre Agassi"
	sponsor       string  = "Head"
	worldRanking  int     = 1
	stringTension float32 = 65.5
)

func main() {
	// block: only visible within the main function
	var i int = 42
	fmt.Println(i)
	fmt.Println(I)
	fmt.Println(playerName, sponsor, worldRanking, stringTension)
}
