package main

import (
	"fmt"
	"time"
)

func add(x, y int, c chan int) {
	time.Sleep(time.Duration(x) * 100 * time.Millisecond)
	c <- x + y //send to c
}

func main() {
	c := make(chan int)

	go add(3, 4, c)
	go add(1, 2, c)

	x := <-c //receive  from c
	y := <-c //receive  from c

	fmt.Println("The value of x is ", x)
	fmt.Println("The value of y is ", y)
}
