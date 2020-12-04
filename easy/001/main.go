// 两数之和
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum

// 涉及到的算法和数据结构
//
package main

import "fmt"

func main() {
	nums := []int{3, 3, 4, 15, 3, 9}
	target := 7

	ret := twoSum(nums, target)
	fmt.Println(ret)
}

// 暴力枚举
// 时间复杂度 O(n²) 其中n是数组中的元素数量，最坏情况下数组中任意两个数都要被匹配一次
// 空间复杂度 O(1)
// 枚举数组中的每一个数 x ，寻找数组中是否存在 target-x
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 哈希表
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func twoSum2(nums []int, target int) []int {
	// key: value
	// value: index
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]] = i
	}
	// bug：
	// 重新遍历的话，会导致 x 和自己匹配了
	for i := 0; i < len(nums); i++ {
		if index, ok := hashMap[target-nums[i]]; ok {
			if index == i {
				continue
			}
			return []int{index, i}
		}
	}
	return nil
}

// 哈希表(时间换空间方式)
// 时间复杂度 O(n) 其中n是数组中元素数量、对于每一个元素 x，我们可以O(1)地寻找 target-x
// 空间复杂度 O(n) 其中n是数组中元素数量。主要为哈希表的开销
// 注意到暴力枚举中的时间复杂度较高的原因是寻找 target-x 的时间复杂度过高，因此我们需要一种更优秀的方法，能够快速寻找数组中是否存在目标元素
// 如果存在，我们需要找出它的索引。
// 使用哈希表，可以将寻找 target-x 的时间复杂度降低到从O(n)到O(1)
// 这样我们创建一个哈希表，对于每一个 x ，我们首先查询哈希表中是否存在 target-x，然后将 x 插入到哈希表中，即可保证不会让 x 和自己匹配

// 对twoSum2做优化
func twoSum(nums []int, target int) []int {
	// key: value
	// value: index
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, ok := hashMap[target-nums[i]]; ok {
			return []int{index, i}
		}
		hashMap[nums[i]] = i
	}
	return nil
}
