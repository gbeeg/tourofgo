package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomint() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}

func main() {
	if x := randomint(); x > 5 {
		fmt.Println("Random number generated is more than 5")
		fmt.Println("The random number is", x)
	} else {
		fmt.Println("Random number generated is less than 5")
		fmt.Println("The random number is", x)
	}

	//i cant use x here
}
