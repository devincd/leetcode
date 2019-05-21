/**
罗马数字转整数
解法：
建立一个哈希表，然后一一判断，若前面的罗马数比后面的罗马数值小，则减去该罗马值
若前面的罗马数值比后面的罗马数字大(或者等于)，则直接加上该罗马对应的数值
 */
package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	result := 0
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	//(1)特殊情况
	if len(s) == 1{
		return romanMap[s]
	}
	//(2)
	for i:=0; i< len(s)-1; i++ {
		if romanMap[string(s[i])] >= romanMap[string(s[i+1])] {
			result+= romanMap[string(s[i])]
		} else {
			result -= romanMap[string(s[i])]
		}
		//(3)如果处理到了倒数第二个字符的话,需要直接加上最后一个字符的值
		if i+1 == len(s) - 1 {
			result+= romanMap[string(s[i+1])]
		}
	}
	return result
}
