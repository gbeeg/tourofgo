package main

import "fmt"

func countToTen(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	go countToTen(c)
	for val := range c {
		fmt.Println(val)
	}
}
