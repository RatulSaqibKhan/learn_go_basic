package main

import "fmt"

type Car struct {
	Model        string
	Color        string
	FuelCapacity int
	Price        float64
}

func main() {
	car1 := Car{"Toyota Alien", "Black", 200, 2000.2}
	fmt.Println(car1)
	
	fmt.Printf("%+v\n", car1)
	
	fmt.Println(car1.Model)

	car2 := Car{
		Model: "Hyundai Classic",
		Color: "Magenta",
		FuelCapacity: 120,
		Price: 2120.3,
	}
	fmt.Printf("%+v\n", car2)
	
	car3 := Car{}
	fmt.Printf("%+v\n", car3)

}
