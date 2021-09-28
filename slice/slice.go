package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes)      //print all primes
	fmt.Println(primes[1:4]) //fromt 1st index (2nd item) to 4th item
	fmt.Println(primes[:4])  //from zero index (1st item) to 4th item
	fmt.Println(primes[4:])  //from 4th index (5th iem) to the last item

	primes = primes[:0] //change length t0 zero (values still remain)
	printSliceInfo(primes)

	primes = primes[:4] //change length t0 4 (you'll see the values reappear)
	printSliceInfo(primes)

	primes = primes[2:] //change the starting index to 2 (will cut the first 2) of the 4 items
	printSliceInfo(primes)
}

func printSliceInfo(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
