/**
* Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

* You may assume that each input would have exactly one solution,
and you may not use the same element twice.

* You can return the answer in any order.
*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int 
	for i, num1 := range nums {
			for j, num2 := range nums[i+1:] {
					if num1 + num2 == target {
						result = []int {i, i+j+1}
						break
					}
			}
	}
	return result
}

func main() {
	res := twoSum([]int {3,2,1}, 5)

	fmt.Println(res)
}