package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numIndices := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := numIndices[complement]; found {
			return []int{index, i}
		}
		numIndices[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println("Indices of the two numbers:", result)
}
