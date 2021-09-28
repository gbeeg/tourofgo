package main

import "fmt"

func main() {
	s := make([]int, 10) //creates a slice with 10 length
	printSliceInfo(s)

	s[0] = 1
	printSliceInfo(s)
}

func printSliceInfo(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
