//  多数元素

package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	_max, _maxN := 0, 0
	tmp := map[int]int{}
	for _, n := range nums {
		if _, ok := tmp[n]; ok {
			tmp[n] += 1
		} else {
			tmp[n] = 1
		}
		if _max < tmp[n] {
			_max = tmp[n]
			_maxN = n
		}
	}
	return _maxN
}

func a(nums []int) int {
	count := 0
	candidate := -1

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func b(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(a([]int{2, 2, 1, 1, 1, 2, 2}))
}
