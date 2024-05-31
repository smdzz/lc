//  丢失的数字

package main

import (
	"fmt"
)

func missingNumber(nums []int) (xor int) {
	for i, num := range nums {
		xor ^= i ^ num
	}
	return xor ^ len(nums)
}

//func missingNumber(nums []int) int {
//	lens := len(nums) + 1
//	sum := 0
//	if lens%2 == 0 {
//		sum = (lens - 1) * (lens / 2)
//	} else {
//		sum = lens * (lens / 2)
//	}
//	for _, num := range nums {
//		sum -= num
//	}
//	return sum
//}

//func missingNumber(nums []int) int {
//	has := map[int]bool{}
//	for _, v := range nums {
//		has[v] = true
//	}
//	for i := 0; ; i++ {
//		if !has[i] {
//			return i
//		}
//	}
//}

//func missingNumber(nums []int) int {
//	sort.Ints(nums)
//	if nums[0] != 0 {
//		return 0
//	}
//	for i := 1; i < len(nums); i++ {
//		if nums[i]-nums[i-1] != 1 {
//			return nums[i] - 1
//		}
//	}
//	return len(nums)
//}

func main() {
	fmt.Println(missingNumber([]int{0, 2, 3, 4}))
}
