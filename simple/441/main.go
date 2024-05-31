//  排列硬币

package main

import (
	"fmt"
)

//func arrangeCoins(n int) int {
//	lines := 0
//	lineCount := 1
//	for n > 0 {
//		if n >= lineCount {
//			lines += 1
//		}
//		n -= lineCount
//		lineCount += 1
//	}
//	return lines
//}

// 二分查找
func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	left, right := 1, n
	for left < right {
		middle := (left + right) / 2
		if middle*(middle+1) <= 2*n {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left - 1
}
func main() {
	fmt.Println(arrangeCoins(3))
}
