package main

import ("fmt")

func main() {
	book := "The colour of magic"

	fmt.Println(book)

	//length of a string in bytes
	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// uint8 is a byte

	// String in go is immutable
	// book[0] = 88

	// Slice (start, end), 0 based, ½ empty range
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])
	
	// Slice (no start)
	fmt.Println(book[:11])

	// Concatenate string using + symbol
	fmt.Println("t" + book[1:])

	// Unicode support
	fmt.Println("This is ½ price!")

	// Multiline using back tick
	poem :=`
	They jumped from the burning floors—
	one, two, a few more,
	higher, lower.`
	fmt.Println(poem)

}