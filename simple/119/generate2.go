// 杨辉三角2

// 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
//在「杨辉三角」中，每个数是它左上方和右上方的数的和。

// 示例 1:
// 输入: rowIndex = 3
// 输出: [1,3,3,1]
// 示例 2:
// 输入: rowIndex = 0
// 输出: [1]
// 示例 3:
// 输入: rowIndex = 1
// 输出: [1,1]
package main

import "fmt"

// 能否只用一个数组呢？
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}

func main() {
	fmt.Println(getRow(5))
}
