package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}

	// use range function to get array index and value
	for i, v := range arr {
		fmt.Printf("value of index %v is %v\n", i, v)
	}

	// use range function to get the indexes only
	for i := range arr {
		fmt.Printf("value of index %v is %v\n", i, arr[i])
	}

	//use range function to get the values only
	for _, v := range arr {
		fmt.Println(v)
	}
}
