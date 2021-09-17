// The following code will calculate min value of two numbers
package main

import "fmt"

func main() {
	// let go define data type automatically
	// define variables in a single line
	x , y := 1.0 , 2.0

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("x=%v, type of %T\n", y, y)

	mean := ( x + y ) / 2
	fmt.Printf("x=%v, type of %T\n", mean, mean)
}