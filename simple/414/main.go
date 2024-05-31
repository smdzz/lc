//  第三大的数

package main

//func thirdMax(nums []int) int {
//	sort.Ints(nums)
//	for i, diff := len(nums)-1, 1; i > 0; i-- {
//		if nums[i] != nums[i-1] {
//			diff++
//			if diff == 3 { // 此时 nums[i] 就是第三大的数
//				return nums[i-1]
//			}
//		}
//	}
//	return nums[len(nums)-1]
//}

// 一次遍历, 没毛病,初始化为0是不可以的,但是初始为空指针,或者无穷大/无穷小是可以的
func thirdMax(nums []int) int {
	var a, b, c *int
	for _, num := range nums {
		if a == nil || num > *a {
			a, b, c = &num, a, b
		} else if *a > num && (b == nil || num > *b) {
			b, c = &num, b
		} else if b != nil && *b > num && (c == nil || num > *c) {
			c = &num
		}
	}
	if c == nil {
		return *a
	}
	return *c
}

func main() {
	thirdMax([]int{3, 2, 1})
}
