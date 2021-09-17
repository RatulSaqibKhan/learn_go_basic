package main

import ("fmt")

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is Big!")
	} else if x == 5 {
		fmt.Println("x is Equal!")
	} else {
		fmt.Println("x is Less!")
	}
}