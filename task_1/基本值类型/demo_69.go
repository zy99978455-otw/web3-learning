package main

/*
`x的平方根`
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
*/

import (
	"fmt"
)

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right //注意：返回right，因为right 是小于x的最大整数
}

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(16))
}
