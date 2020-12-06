// 罗马数字转整数
// 字符          数值
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/roman-to-integer
package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var ret int
	for i := 0; i < len(s); i++ {
		left := roman[string(s[i])]
		// 判断右边后一位是否比本位大
		if i+1 < len(s) {
			right := roman[string(s[i+1])]
			if left < right {
				left = -left
			}
		}
		ret += left
	}
	return ret
}
