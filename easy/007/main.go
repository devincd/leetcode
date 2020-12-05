// 整数反转
// case-1
// 输入: 123
// 输出: 321

// case-2
// 输入: -123
// 输出: -321

// case-3
// 输入: 120
// 输出: 21

// 注意:
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为[−2^31, 2^31-1]。请根据这个假设，如果反转后整数溢出那么就返回 0

//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-integer

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("input: 123")
	fmt.Println("output:", reverse(123))

	fmt.Println("input: -123")
	fmt.Println("output:", reverse(-123))

	fmt.Println("input: 120")
	fmt.Println("output:", reverse(120))
}

// 解法-1
// 借助 栈 数据结构
// 1.先将数据从带符号个位开始入栈
// 2.出栈：从数据带符号最高位出栈
// 3.计算规则如下: ret = ret + 位数*10^x
func reverse1(x int) int {
	var ret = 0
	retSlice := make([]int, 0)
	// 数据入栈
	for x != 0 {
		index := x % 10
		retSlice = append(retSlice, index)
		x = x / 10
	}
	// 数据出栈
	for i := len(retSlice) - 1; i >= 0; i-- {
		temp := ret + retSlice[i]*int(math.Pow(float64(10), float64(len(retSlice)-1-i)))
		// 数据溢出判断
		if temp > math.MaxInt32 || temp < math.MinInt32 {
			return 0
		}
		ret = temp
	}
	return ret
}

// 解法-2
// 借助 队列 数据结构
// 1.先将数据从带符号个位开始入队列
// 2.出栈：从数据带符号个位出队列
// 3.计算规则如下: ret = ret*10 +x
func reverse2(x int) int {
	var ret = 0
	retSlice := make([]int, 0)
	// 数据入栈
	for x != 0 {
		index := x % 10
		retSlice = append(retSlice, index)
		x = x / 10
	}
	// 数据出栈
	for _, val := range retSlice {
		temp := ret*10 + val
		// 数据溢出判断
		if temp > math.MaxInt32 || temp < math.MinInt32 {
			return 0
		}
		ret = temp
	}
	return ret
}

// 最终解法
// 队列优化版，去掉队列存储逻辑
func reverse(x int) int {
	var ret int
	for x != 0 {
		// 弹出
		index := x % 10
		// 推入数据&溢出前进行检查
		temp := ret*10 + index
		if temp > math.MaxInt32 || temp < math.MinInt32 {
			return 0
		}
		ret = temp
		x = x / 10
	}
	return ret
}
