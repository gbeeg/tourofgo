package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v is %v years old", p.Name, p.Age)
}

func main() {
	a := Person{"Ben", 31}
	b := Person{"Steve", 0}

	fmt.Println(a, b)
}
