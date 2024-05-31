//  存在重复元素

package main

import "fmt"

//func containsNearbyDuplicate(nums []int, k int) bool {
//	m := map[int]int{}
//	for i := 0; i < len(nums); i++ {
//		if v, ok := m[nums[i]]; ok {
//			if i-v <= k {
//				return true
//			} else {
//				m[nums[i]] = i
//			}
//		} else {
//			m[nums[i]] = i
//		}
//	}
//	return false
//}

// 滑动窗口
func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]struct{}{}
	for i, v := range nums {
		if i-k > 0 {
			delete(m, nums[i-k-1])
		}
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = struct{}{}
		}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
