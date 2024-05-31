//  二进制手表

package main

import (
	"fmt"
	"math/bits"
)

// 时分枚举
func readBinaryWatch(turnedOn int) (ans []string) {
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return
}

//// 二进制枚举
//func readBinaryWatch(turnedOn int) (ans []string) {
//	// 高四位最大为11,也就是1011,所以i最大为1011111111,只要i<767 + 1;不需要扫描到1024
//	for i := 0; i < 768; i++ {
//		h, m := i>>6, i&63 // 用位运算取出高 4 位和低 6 位
//		if h < 12 && m < 60 && bits.OnesCount(uint(i)) == turnedOn {
//			ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
//		}
//	}
//	return
//}

func main() {
	fmt.Println(readBinaryWatch(1))
}
