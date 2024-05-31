//  比特位计数

package main

import "fmt"

//func countBits(n int) []int {
//	res := make([]int, 0, n+1)
//	for i := 0; i <= n; i++ {
//		tmp := i
//		sum := 0
//		for tmp > 0 {
//			if tmp%2 == 1 {
//				sum += 1
//			}
//			tmp /= 2
//		}
//		res = append(res, sum)
//	}
//	return res
//}

//func onceCount(n int) int {
//	sum := 0
//	for ; n > 0; n &= n - 1 {
//		sum += 1
//	}
//	return sum
//}
//
//func countBits(n int) []int {
//	res := make([]int, n+1)
//	for x := range res {
//		res[x] = onceCount(x)
//	}
//	return res
//}

// 方法二：动态规划——最高有效位
func countBits(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}

func main() {
	fmt.Println(countBits(16))
}
