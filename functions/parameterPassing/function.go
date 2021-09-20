package main

import ("fmt")

// double a value of a slice's specific index
func doubleAt(value []int, i int) {
	value[i] *= 2
}

func doubleInt(a *int) {
	*a *= 2
}

func main() {
	values := []int{1,2,3,4}

	fmt.Println(values)

	doubleAt(values, 2)

	fmt.Println(values)

	a := 6

	doubleInt(&a)

	println(a)

}