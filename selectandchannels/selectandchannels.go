package main

import "fmt"

func checkValue(c chan int, quit chan int) {
	x := 0
	for { //this is an infinite loop
		select {
		case c <- x: //this wil assign x to the channel c and will be printed by the go func() inside the main function
			x = x + 1 //this will increment x by 1 for every loop
		case <-quit: //this will trigger when the go func() under main completes its main loop and exit the function
			fmt.Println("Quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //this will print out whenever there's a change of value in c
		}
		quit <- 0 //assign zero after this to make checkValue quit
	}()

	checkValue(c, quit)
}
