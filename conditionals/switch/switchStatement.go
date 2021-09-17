package main

import (
	"fmt"
)

func main() {
	x := 2

	// In go break statement is not needed for each case of switch statement
	switch x % 2 {
	case 1:
		fmt.Println("x is odd!")
	default:
		fmt.Println("x is even!")
	}

	// We can use expression in case instead of switch
	switch {
	case x > 100:
		fmt.Println("x is very big")
	case x > 10:
		fmt.Println("x is big")
	default:
		fmt.Println("x is small")
	}
}
