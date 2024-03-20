package main

import "fmt"

func longestPalindrome(s string) int {
	m := map[byte]int{}
	for _, v := range []byte(s) {
		m[v] += 1
	}
	res := 0
	hasSingle := false
	for _, v := range m {
		res += v / 2
		if v%2 == 1 {
			hasSingle = true
		}
	}
	res = res * 2
	if hasSingle {
		res += 1
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("a"))
}
