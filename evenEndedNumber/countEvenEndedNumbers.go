/*
 * An even-ended numbers are the numbers with same first and last digits e.g. 11, 121
 *
 * This program counts the total number of even-ended numbers by multiplying two four digit numbers
 *
 */

 package main

 import ("fmt")

 func main() {
	 // counter
	 count := 0

	 // for every pair of four digit numbers
	 for a := 1000; a <= 9999; a++ {
		 for b := a; b <= 9999; b++ { // don't count twice
			n := a * b // Multiply numbers

			s := fmt.Sprintf("%d", n) // convert to string

			// check if number is even ended
			if s[0] == s[len(s) - 1] {
				count++
			}

		 }
	 }

	 fmt.Println(count)
 }