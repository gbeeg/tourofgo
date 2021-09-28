package main

import "fmt"

func main() {
	adder := func(x int, y int) int {
		return x + y
	}
	fmt.Println(adder(1, 3))
	fmt.Println(compute(adder))
}

func compute(fn func(int, int) int) int {
	return fn(2, 3)
}
