//  存在重复元素

package main

import (
	"fmt"
)

//// 排序看相邻的两个是否相等,相等则认为存在重复元素
//func containsDuplicate(nums []int) bool {
//	sort.Ints(nums)
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] == nums[i+1] {
//			return true
//		}
//	}
//	return false
//}

// map
func containsDuplicate(nums []int) bool {
	m := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = struct{}{}
		}
	}
	return false
}

func main() {
	//[1,2,3,1]
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1}))
}
