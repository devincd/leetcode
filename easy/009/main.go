// 回文数
// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

// case-1:
// 输入: 121
// 输出: true

// case-2:
// 输入: -121
// 输出: false
// 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

// case-3:
// 输入: 10
// 输出: false
// 解释: 从右向左读, 为 01 。因此它不是一个回文数。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/palindrome-number
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%t", isPalindrome(10))
}

// 通过数字反转，如果反转前后一致的话就是回文数
// 注: 负数显然都不可能是回文数，所以回文数只包括0和正整数
func isPalindrome1(x int) bool {
	// 负数都不是回文数
	if x < 0 {
		return false
	}
	// 正整数反转
	var reverseX int
	originalX := x
	for x != 0 {
		index := x % 10
		temp := reverseX*10 + index
		// 判断数字是否越界
		if temp > math.MaxInt32 {
			return false
		}
		reverseX = reverseX*10 + index
		x = x / 10
	}

	return reverseX == originalX
}

// 借助slice将正整数每位数字放到slice中去，判断收尾数字是否一致
// 注: 负数显然都不可能是回文数，所以回文数只包括0和正整数
func isPalindrome2(x int) bool {
	// 负数都不是回文数
	if x < 0 {
		return false
	}

	sliceX := make([]int, 0)
	for x != 0 {
		// 从个位开始将数字存入到sliceX
		sliceX = append(sliceX, x%10)
		x = x / 10
	}
	// 对sliceX做判断, 数字走到一半即可
	for i := 0; i < len(sliceX)/2; i++ {
		if sliceX[i] != sliceX[len(sliceX)-1-i] {
			return false
		}
	}
	return true
}

// 反转一半数字
func isPalindrome(x int) bool {
	// 负数不可能是回文数, 个位数为0的情况，下面的算法判断不出来
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	// 如何判断一个数字已经反转了一半的数字
	var reverseX int
	for reverseX < x {
		index := x % 10
		reverseX = reverseX*10 + index
		x = x / 10
	}
	// 需要同时处理正整数位数为 奇数和偶数的情况
	return reverseX == x || x == reverseX/10
}
