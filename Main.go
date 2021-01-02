package main

import (
	"fmt"
)

func main() {
	// multi-line approach: useful for for loops
	var i int
	i = 14
	// single-line: useful for when you need control over the data type
	var j float32 = 15
	// colon-equals: least verbose, use when Go's compiler is able to determine data type
	k := 16
	fmt.Println(i)
	fmt.Printf("%v, %T", j, j)
	fmt.Printf("%v, %T", k, k)
}
