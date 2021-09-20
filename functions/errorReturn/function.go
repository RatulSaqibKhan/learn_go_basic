package main

import (
	"fmt"
	"math"
)

func sqrt(a float64) (float64, error) {
	if (a < 0) {
		return 0.0, fmt.Errorf("ERROR: sqrt of negative value (%.2f)", a)
	}

	return math.Sqrt(a), nil
}

func main() {
	n1 := 16.0
	n2 := -6.0

	s1, err1 := sqrt(n1)
	if err1 == nil {
		fmt.Printf("Sqrt of %.2f is %.2f\n",n1,s1)
	} else {
		fmt.Println(err1)
	}
	
	s2, err2 := sqrt(n2)
	if err2 == nil {
		fmt.Printf("Sqrt of %.2f is %.2f\n",n2,s2)
	} else {
		fmt.Println(err2)
	}
}