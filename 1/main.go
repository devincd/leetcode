package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if index, ok := numMap[temp]; ok {
			return []int{index, i}
		}
		numMap[nums[i]] = i

	}
	return nil
}
