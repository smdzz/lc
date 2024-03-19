package main

import "fmt"

type NumArray struct {
	Arr []int
}

//func Constructor(nums []int) NumArray {
//	return NumArray{Arr: nums}
//}
//
//func (this *NumArray) SumRange(left int, right int) int {
//	sum := 0
//	for i := left; i <= right; i++ {
//		sum += this.Arr[i]
//	}
//	return sum
//}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Arr[right+1] - this.Arr[left]
}

func main() {
	a := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(a.SumRange(0, 2))
	fmt.Println(a.SumRange(2, 5))
	fmt.Println(a.SumRange(0, 5))

}
