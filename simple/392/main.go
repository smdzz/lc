//  判断子序列

package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	si, ti := 0, 0
	for si < len(s) && ti < len(t) {
		if t[ti] != s[si] {
			ti += 1
		} else {
			ti += 1
			si += 1
		}
	}
	if si == len(s) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
