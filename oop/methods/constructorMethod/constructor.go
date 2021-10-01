package main

import (
	"fmt"
	"os"
)

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

func NewCar(model string, color string, fuelCapacity int, price float64, milage float64) (*Car, error) {
	if price == 0 {
		return nil, fmt.Errorf("p	lease give valid price")
	}

	car := &Car{
		Model:        model,
		Color:        color,
		FuelCapacity: fuelCapacity,
		Price:        price,
		Milage:       milage,
	}

	return car, nil
}

func main() {
	car, err := NewCar("Toyota", "White", 130, 1200.3, 1.4)

	if err != nil {
		fmt.Printf("Error: Cannot create car - %s\n", err)
		os.Exit(1)
	}

	fmt.Println(car.TotalMileUsingFullTank())
}
