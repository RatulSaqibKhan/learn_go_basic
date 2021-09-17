package main

import (
	"fmt"
)

func main() {
	x := 10

	// If Else Block
	if x > 100 {
		fmt.Println("x is very Big!")
	} else {
		fmt.Println("x is not that Big!")
	}

	// If Else If Block
	if x > 5 {
		fmt.Println("x is Big!")
	} else if x == 5 {
		fmt.Println("x is Equal!")
	} else {
		fmt.Println("x is Less!")
	}

	// Logical AND operator
	if x > 5 && x < 19 {
		fmt.Println("x is in the given range!")
	}

	// Logical OR operator
	if x < 24 || x > 30 {
		fmt.Println("x is out of range!")
	}

	// Optional initialization operator
	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b!")
	}
}
