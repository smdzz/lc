//  2的幂

package main

import (
	"fmt"
)

// 判断n是否是2的整数次方,n不可能为负数
func isPowerOfTwo(n int) bool {
	for {
		if n == 1 {
			return true
		}
		if n <= 0 {
			return false
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			return false
		}
	}
}

// 位运算前提是n > 0
// 使用位运算
// >>> bin(16)
// '0b10000'
// >>> bin(32)
// '0b100000'
// >>> bin(15)
// '0b1111'
// >>> bin(31)
// '0b11111'
// 当n为幂等时
//func isPowerOfTwo(n int) bool {
//	return n > 0 && n&(n-1) == 0
//}

func main() {
	fmt.Println(-8 & -9)
	fmt.Println(isPowerOfTwo(-8))
}
