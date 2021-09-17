package main

import (
	"fmt"
)

func main() {
	// Same type
	animals := []string{"Lion", "Wolf", "Hayena", "Zebra", "Deer", "Tiger"}

	// Length
	fmt.Println(len(animals))  // 6

	fmt.Println("-------------")

	// 0 indexing
	fmt.Println(animals[1]) //Wolf

	fmt.Println("-------------")

	// slices
	fmt.Println(animals[1:]) // [Wolf Hayena Zebra Deer Tiger]

	fmt.Println("-------------")
	// loop through slice
	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i])
	}

	fmt.Println("-------------")
	// Single value range
	for i := range animals {
		fmt.Println(i, "=> " + animals[i])
	}
	
	fmt.Println("-------------")
	// Double value range
	for i, name := range animals {
		fmt.Println(i, "=> " + name)
	}

	fmt.Println("-------------")
	// Double value range ignoring index by using _
	for _, name := range animals {
		fmt.Println(name)
	}

	fmt.Println("-------------")
	// Use of append
	animals = append(animals, "Gorilla")
	fmt.Println(animals)
}
