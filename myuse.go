package main

import (
	"fmt"

	"github.com/zhanggy1984/mathutil"
)

func main() {
	// 使用基本运算
	sum := mathutil.Add(2, 11)
	fmt.Println("Sum:", sum) // 输出: Sum: 5

	// 使用统计功能
	numbers := []int{1, 2, 3, 4, 5, 6}
	avg := mathutil.Average(numbers)
	fmt.Println("Average:", avg) // 输出: Average: 3

	// 使用几何运算
	area := mathutil.CircleArea(1.15)
	fmt.Println("Circle Area:", area) // 输出: Circle Area: 12.566370614359172
}
