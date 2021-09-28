package main

import "fmt"

func main() {
	var arr []int
	printSlice(arr)

	arr = append(arr, 1)
	printSlice(arr)
	fmt.Println(arr[0])

	arr = append(arr, 2, 3, 4, 5)
	printSlice(arr)
	fmt.Println(arr[1], arr[2])
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
