package main

import (
	"fmt"
)

const (
	_ = iota
	dogSpecialist
	catSpecialist
	snakeSpecialist
)

func main() {
	var specialist int = 1
	fmt.Printf("%v\n", specialist == dogSpecialist)
}
