package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	// use strconv package's str converter function to change from an integer to a string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)

}
