//  找到所有数组中消失的数字

package main

import "fmt"

//	func findDisappearedNumbers(nums []int) []int {
//		res := []int{}
//		m := map[int]struct{}{}
//		for _, v := range nums {
//			m[v] = struct{}{}
//		}
//		for n := 1; n <= len(nums); n++ {
//			if _, ok := m[n]; !ok {
//				res = append(res, n)
//			}
//		}
//		return res
//	}
func findDisappearedNumbers(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
