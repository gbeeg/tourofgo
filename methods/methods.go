package main

import "fmt"

type Car struct {
	brand    string
	model    string
	year     int
	codename string
	totalkm  int
}

//Normal method declaration
func (c Car) DisplayName() string {
	return fmt.Sprintf("%d %s %s %s", c.year, c.brand, c.model, c.codename)
}

//Method pointer declaration (if you want a method that changes the value of the struct)
func (c *Car) DriveCar(noOfKm int) {
	c.totalkm += noOfKm
}

func main() {
	toyota_corolla := Car{"Toyota", "Corolla", 2018, "E210", 0}

	//you can call it like a "class" method. Note: there's no classes i ngo
	fmt.Println(toyota_corolla.DisplayName())

	toyota_corolla.DriveCar(10)
	toyota_corolla.DriveCar(5)
	fmt.Println("Total km driven", toyota_corolla.totalkm)
}
