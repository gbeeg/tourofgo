package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("You entered an int: ", v)
	case string:
		fmt.Println("You enetered a string", v)
	default:
		fmt.Printf("Not sure what you keyed in but its provably a %T: %v\n", v, v)
	}
}

func main() {
	do(1)
	do("Hello world")
	do(1.23)
}
