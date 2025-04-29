package main

/*
`我的日程安排表`
实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，则可以存储这个新的日程安排。
当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。
日程可以用一对整数 startTime 和 endTime 表示，这里的时间是半开区间，即 [startTime, endTime), 实数 x 的范围为，  startTime <= x < endTime 。
实现 MyCalendar 类：
MyCalendar() 初始化日历对象。
boolean book(int startTime, int endTime) 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true 。否则，返回 false 并且不要将该日程安排添加到日历中。
*/

import "fmt"

type MyCalendar struct {
	bookings [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		bookings: make([][2]int, 0),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, b := range this.bookings {
		//判断是否有交集(半开区间判断)
		if start < b[1] && end > b[0] {
			return false
		}
	}
	this.bookings = append(this.bookings, [2]int{start, end})
	return true
}

func main() {
	calendar := Constructor()
	fmt.Println(calendar.Book(10, 20)) //true
	fmt.Println(calendar.Book(15, 25)) //false, 和[10, 20)重叠
	fmt.Println(calendar.Book(20, 30)) //true, 不重叠
}
