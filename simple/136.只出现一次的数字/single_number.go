package main

import "fmt"

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
//
// 示例 1 ：
// 输入：nums = [2,2,1]
// 输出：1
// 示例 2 ：
// 输入：nums = [4,1,2,1,2]
// 输出：4
// 示例 3 ：
// 输入：nums = [1]
// 输出：1
//
// 使用位运算。对于这道题，可使用异或运算 ⊕\oplus⊕。异或运算有以下三个性质。
//
// 任何数和 000 做异或运算，结果仍然是原来的数。
// 任何数和其自身做异或运算，结果是 0
// 数组中的全部元素的异或运算结果即为数组中只出现一次的数字。
func singleNumber(nums []int) int {
	single := 0
	for _, x := range nums {
		single ^= x
	}
	return single
}

func main() {
	fmt.Println(singleNumber([]int{8}))
}
