package main

import ("fmt")

func main() {
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------")
	// break statement in loop
	for i := 0; i < 5; i++ {
		if i > 3 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("---------")
	// continue statement in loop
	for i := 0; i < 5; i++ {
		if i < 2 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("---------")
	// while loop structure
	a := 1

	for a < 5 {
		fmt.Println(a)
		a++
	}

	fmt.Println("---------")
	// do while loop structure
	b := 1
	for {
		if (b > 4) {
			break
		}
		fmt.Println(b)
		b++
	}

}