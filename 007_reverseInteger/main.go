/**
官方解法:
弹出和推入数字 & 溢出前进行检查
 */
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(-123))
	fmt.Println(reverseWithString(-123))
}

func reverse(x int) int {
	num := 0
	for {
		//(1)弹出数字
		pop := x % 10
		x = x / 10

		//(2)推送数字
		num = num * 10 + pop
		if x == 0 {
			break
		}
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		num = 0
	}
	return num
}

// 另外的做法
func reverseWithString(x int) int {
	str := strconv.Itoa(x)
	var (
		newStr string
		prefix string
	)
	if string(str[0]) == "-" {
		prefix = "-"
		str = str[1:]
	}
	for i := len(str)-1; i >= 0; i-- {
		newStr += string(str[i])
	}
	newX, _ := strconv.Atoi(prefix + newStr)
	if newX < math.MinInt32 || newX > math.MaxInt32 {
		return 0
	}
	return newX
}