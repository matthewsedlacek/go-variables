package main

import (
	"fmt"
)

const (
	_ = iota + 2
	doctor
	lawyer
	developer
	dogSitter
	accountant
)

func main() {
	var occupationType int = lawyer
	fmt.Printf("%v\n", occupationType == lawyer)
	fmt.Printf("%v\n", lawyer)
}
