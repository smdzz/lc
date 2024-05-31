//  汇总区间

package main

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}
	if len(nums) == 0 {
		return res
	}
	if len(res) == 1 {
		res = append(res, fmt.Sprintf("%d", nums[0]))
		return res
	}
	pre, next := 0, 1
	for next < len(nums) {
		if nums[next]-nums[next-1] > 1 {
			if next-pre > 1 {
				res = append(res, fmt.Sprintf("%d->%d", nums[pre], nums[next-1]))
			} else {
				res = append(res, fmt.Sprintf("%d", nums[pre]))
			}
			pre = next
			next = next + 1
		} else {
			next += 1
		}
	}
	if pre == len(nums)-1 {
		res = append(res, fmt.Sprintf("%d", nums[pre]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[pre], nums[len(nums)-1]))
	}

	return res
}

func main() {
	// ["0->2","4->5","7"]
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}
