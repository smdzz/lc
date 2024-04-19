package main

//func hammingDistance(x, y int) int {
//	return bits.OnesCount(uint(x ^ y))
//}

//func hammingDistance(x, y int) (ans int) {
//	// 每次和1进行&,如果最后一位是1,那么就+1,如果最后一位不为1,就跳过
//	for s := x ^ y; s > 0; s >>= 1 {
//		ans += s & 1
//	}
//	return
//}

// Brian Kernighan 算法
func hammingDistance(x, y int) (ans int) {
	for s := x ^ y; s > 0; s &= s - 1 {
		ans++
	}
	return
}
