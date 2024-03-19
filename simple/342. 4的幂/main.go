package main

import (
	"fmt"
	"math/bits"
)

//
//func isPowerOfFour(n int) bool {
//	if n&(n-1) != 0 || n < 1 {
//		return false
//	}
//
//	for n > 1 {
//		if n%4 != 0 {
//			return false
//		}
//		n /= 4
//	}
//	return true
//}

func isPowerOfFour(n int) bool {
	return n > 0 && bits.OnesCount(uint(n)) == 1 && bits.TrailingZeros(uint(n))%2 == 0
}

// n % 3 = 1这个绝了
//func isPowerOfFour(n int) bool {
//	return n & (n-1) == 0 && n > 0 && n % 3 == 1
//}

func main() {
	fmt.Println(isPowerOfFour(0))
}
