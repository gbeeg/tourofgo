package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("The sum of 2 and 3 is ", add(2, 3))
}
