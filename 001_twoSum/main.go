/**
官方解法:
一遍哈希表(采用以空间换取速度的方式)
在进行迭代并将元素插入到表中的同时，我们还会回过头来检查表中是否已经存在当前元素所对应的目标元素。如果它存在
那我们已经找到了对应解，并立即返回
 */

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
	numsMap := make(map[int]int)
	for key, value := range nums {
		//(1)检查当前元素是否在哈希表中已经存在
		if i, ok := numsMap[value]; ok {
			return []int{i, key}
		}
		//(2)将目标值-当前元素值存入到哈希表中去
		numsMap[target - value] = key
	}
	return nil
}