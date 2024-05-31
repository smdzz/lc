//  找不同

package main

import "fmt"

//func findTheDifference(s string, t string) byte {
//	m := map[byte]int{}
//	for _, v := range []byte(s) {
//		m[v] += 1
//	}
//	for _, v := range []byte(t) {
//		m[v] -= 1
//		if m[v] < 0 {
//			return v
//		}
//	}
//	return 0
//}

//// 异或, 同字符^同字符等于0
//func findTheDifference(s, t string) (diff byte) {
//	for i := range s {
//		diff ^= s[i] ^ t[i]
//	}
//	return diff ^ t[len(t)-1]
//}

func findTheDifference(s, t string) (diff byte) {
	sum := 0
	for _, v := range []byte(s) {
		sum -= int(v)
	}
	for _, v := range []byte(t) {
		sum += int(v)
	}
	return byte(sum)
}

func main() {
	fmt.Println(findTheDifference("meng", "maneg"))
}
