package main

import "fmt"

func main() {
	x := 1
	for x <= 10 {
		if x%2 == 0 {
			fmt.Println(x, " is even")
		} else {
			fmt.Println(x, " is odd")
		}
		x += 1
	}
}
