package main

func moveZeroes(nums []int) {
	fast, slow, n := 0, 0, len(nums)
	for fast < n {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow += 1
		}
		fast += 1
	}
}

func main() {
	nums := []int{0}
	moveZeroes(nums)
}
