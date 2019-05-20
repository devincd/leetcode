/**
回文数
官方解法:

 */
package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	// 特殊情况
	// (1)当x < 0时, x不是回文数
	// (2)同样地，如果数字的最后一位是0，为了使该数字为回文，则第一位也应该是0
	if x < 0 ||  (x % 10 == 0 && x != 0) {
		return false
	}
	
	var revertedNumber int
	for x > revertedNumber {
		revertedNumber = revertedNumber * 10 + x %10
		x = x /10
	}
	// 当数字长度为奇数时，我们可以通过revertedNumber/10去除处于中位的数字
	return x == revertedNumber || x == revertedNumber/10
}
