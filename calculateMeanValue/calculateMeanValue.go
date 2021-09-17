// The following code will calculate min value of two numbers
package main

import "fmt"

func main() {
	// let go define data type automatically
	x := 1.0
	y := 2.8

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("x=%v, type of %T\n", y, y)

	mean := ( x + y ) / 2
	fmt.Printf("x=%v, type of %T\n", mean, mean)
}