package main

import (
	"fmt"
)

// 第一个字符为不为空加1,或者当前字符不为空,当前的前一个字符为空加一
func countSegments(s string) int {
	res := 0
	for i, ch := range s {
		if (i == 0 || string(s[i-1]) == " ") && string(ch) != " " {
			res += 1
		}
	}
	return res
}

func main() {
	fmt.Println(countSegments("Of all the gin joints in all the towns in all the world,   ")) // 13
	fmt.Println(countSegments(", , , ,        a, eaefa"))                                     // 6
}
