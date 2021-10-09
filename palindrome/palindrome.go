package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	var x_temp int = x
	var x_rev int = 0
	// make the number reverse to the half
	for x_temp > 0 {
		x_rev = (x_rev * 10) + (x_temp % 10)
		x_temp /= 10
	}
	return x == x_rev
}

func main() {
	var n int = 1123211

	if isPalindrome(n) {
		fmt.Printf("%d is a palindrome\n", n)
	} else {
		fmt.Printf("%d is not a palindrome\n", n)
	}
}
