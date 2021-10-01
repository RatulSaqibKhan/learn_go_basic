package main

import "fmt"

type Car struct {
	Model        string
	Color        string
	FuelCapacity int
	Price        float64
	Milage 			 float64
}

func (car *Car) TotalMileUsingFullTank() float64 {
	milesToGo := float64(car.FuelCapacity) / car.Milage

	return milesToGo
}

func main() {
	car1 := Car{"Toyota Alien", "Black", 200, 2000.2, 2}
	fmt.Println(car1)
	fmt.Println(car1.TotalMileUsingFullTank())

}
