package main

/*
`打家劫舍`
你是一个专业的小偷，计划偷窃沿街的房屋。
每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，
系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你不触动警报装置的情况下，一夜之内能够偷窃到的最高金额。
这道题可以使用动态规划的思想，
通过 for 循环遍历数组，利用 if 条件判断来决定是否选择当前房屋进行抢劫，
状态转移方程为 dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])。
*/

import "fmt"

// 定义打家劫舍函数， 接收一个整数切片，返回一个整数，表示偷窃到的最高金额
func rob(nums []int) int {

	n := len(nums) //获取房屋数量

	// 如果没有房屋，返回0
	if n == 0 {
		return 0
	}

	// 如果只有一个房屋，直接偷窃该房屋
	if n == 1 {
		return nums[0]
	}

	// prev表示dp[i-2]，即前前一个状态的最高金额
	// curr表示dp[i-1], 即前一个状态的最高金额
	prev, curr := nums[0], max(nums[0], nums[1])

	// 从第二个房屋开始遍历
	for i := 2; i < n; i++ {
		// next表示当前房屋的最高金额，选择偷窃当前房屋加上prev或者不偷窃当前房屋取最大值
		next := max(curr, prev+nums[i])
		// 更新prev和curr，为下一个循环做准备
		prev, curr = curr, next
	}
	// 返回最终的最高金额
	return curr
}

// 定义一个求两个整数最大值的函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))

}
