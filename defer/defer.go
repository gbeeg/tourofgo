package main

import "fmt"

func main() {
	for x := 1; x < 10; x++ {
		//this deferred function gets stacked first and executed after the "done" println is completed
		defer fmt.Println(x)
	}

	fmt.Println("done")
}
