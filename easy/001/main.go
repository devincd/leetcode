package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 0

	ret := twoSum(nums, target)
	fmt.Println(ret)
}

func twoSum(nums []int, target int) []int {
	numsHash := make(map[int]int)
	for i, j := range nums {
		if index, ok := numsHash[j]; ok {
			return []int{index, i}
		}
		numsHash[target-j] = i
	}

	return []int{}
}
