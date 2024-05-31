//  数字转换为十六进制数

package main

import (
	"fmt"
	"strings"
)

//func toHex(num int) string {
//	if num == 0 {
//		return "0"
//	}
//	l := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
//	res := []string{}
//	if num < 0 {
//		num += 2 << 31
//	}
//	for num > 0 {
//		res = append([]string{l[num%16]}, res...)
//		num /= 16
//	}
//	return strings.Join(res, "")
//}

// 位运算
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	sb := &strings.Builder{}
	for i := 7; i >= 0; i-- {
		val := num >> (4 * i) & 0xf
		if val > 0 || sb.Len() > 0 {
			var digit byte
			if val < 10 {
				digit = '0' + byte(val)
			} else {
				digit = 'a' + byte(val-10)
			}
			sb.WriteByte(digit)
		}
	}
	return sb.String()
}

// 26 1a   -1 ffffffff
func main() {
	fmt.Println(toHex(26))
}
