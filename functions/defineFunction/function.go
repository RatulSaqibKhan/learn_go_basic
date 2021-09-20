package main

import (
	"fmt"
)
// Display message
func displayMsg(msg string) string {
	return msg
}

// adding two integers a and b
func add(a int, b int) int {
	return a + b
}

// divmod returns quotient and remainder
func divmod(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	message := displayMsg("Hello world!")
	
	fmt.Println(message)

	a := 12
	b := 4

	fmt.Println(add(a,b))

	fmt.Println(divmod(a,b))
}