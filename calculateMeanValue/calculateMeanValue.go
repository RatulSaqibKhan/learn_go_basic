// The following code will calculate min value of two numbers
package main

import "fmt"

func main() {
	// Explicitly define data type
	var x float64
	var y float64

	x = 1.0
	y = 2.0

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("x=%v, type of %T\n", y, y)

	var mean float64

	mean = ( x + y ) / 2
	fmt.Printf("x=%v, type of %T\n", mean, mean)
}