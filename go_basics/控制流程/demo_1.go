package main

/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
*/

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	nums := []int{4, 1, 2, 1, 2}

	//调用函数
	unique := singleNumber(nums)

	//输出结果
	fmt.Println("只出现一次的数字是：", unique)
}
