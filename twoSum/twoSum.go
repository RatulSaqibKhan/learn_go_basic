/**
* Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

* You may assume that each input would have exactly one solution,
and you may not use the same element twice.

* You can return the answer in any order.
*/

package main

import "fmt"

// O(n2)
func twoSum(nums []int, target int) []int {
	var result []int

	for i, num1 := range nums {
		for j, num2 := range nums[i+1:] {
			if num1+num2 == target {
				result = []int{i, i + j + 1}
				break
			}
		}
	}
	return result
}

// O(n)
func twoSumUsingMap(nums []int, target int) []int {
	resultMap := map[int]int{
		nums[0]: 0,
	}
	var result []int

	for i, num1 := range nums {
		required_num := target - num1

		searchVal, ok := resultMap[required_num]
		if ok && i != searchVal {
			result = []int{searchVal, i}
			break
		} else {
			resultMap[num1] = i
		}
	}

	return result
}

func main() {
	res := twoSumUsingMap([]int{1, 6, 2, 3}, 5)

	fmt.Println(res)
}
