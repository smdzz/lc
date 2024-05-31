//  位1的个数

package main

import "fmt"

//func hammingWeight(num uint32) int {
//	var count int
//	for num > 0 {
//		if num&1 == 1 {
//			count += 1
//		}
//		num >>= 1
//	}
//	return count
//}

// 6 = (110)2   5 = (101)2   6 & 5 = (100)2
// 10 = (1010)2 9 = (1001)2  10 & 9 = (1000)2
// 在实际代码中，我们不断让当前的 n 与 n−1做与运算，直到 n 变为 0 即可。因为每次运算会使得 n 的最低位的 1 被翻转，因此运算次数就等于 n 的二进制位中 1 的个数。

func hammingWeight(num uint32) (ones int) {
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return
}

func main() {
	fmt.Println(hammingWeight(11))
}
