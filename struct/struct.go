package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(v.Y)

	//with pointers
	p := &v
	p.X = 3
	fmt.Println(v.X, p.X)

	//initialize only 1 X
	v1 := Vertex{X: 5}
	fmt.Println(v1)
}
