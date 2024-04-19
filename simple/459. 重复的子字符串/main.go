package main

import (
	"fmt"
)

// 给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i*2 <= n; i++ {
		if n%i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("abcabcabc"))    // true
	fmt.Println(repeatedSubstringPattern("aba"))          // false
	fmt.Println(repeatedSubstringPattern("abcabcabcabc")) // true
}
