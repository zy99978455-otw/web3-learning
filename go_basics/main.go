package main

import "fmt"

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，
// 其余每个元素均出现两次。找出那个只出现了一次的元素。
// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}

func singleNumber(nums []int) int {
	//1. 使用map 计算数量
	// var m = make(map[int]int)
	// for _, v := range nums {
	// 	value := m[v]
	// 	if value == 1 {
	// 		delete(m, v)
	// 	} else {
	// 		m[v] = 1
	// 	}
	// }
	// fmt.Println(m)
	// for k, _ := range m {
	// 	return k
	// }
	// return 0

	//2. 使用异或
	//相同数异或=0
	//0与任何数异或=任意数
	//异或满足交换律，所以异或所有数，最后剩下的数就是只出现一次的数
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
