package main

/*
`合并区间`
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按区间的start排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			//有重叠，合并区间
			last[1] = max(last[1], current[1])
			res[len(res)-1] = last
		} else {
			// 没有重叠，直接加进去
			res = append(res, current)
		}
	}

}
