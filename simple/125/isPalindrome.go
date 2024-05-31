// 验证回文串

package main

import (
	"fmt"
	"strings"
)

// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
// 字母和数字都属于字母数字字符。
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
//
// 示例 1：
// 输入: s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。
// 示例 2：
// 输入：s = "race a car"
// 输出：false
// 解释："raceacar" 不是回文串。
// 示例 3：
// 输入：s = " "
// 输出：true
// 解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
// 由于空字符串正着反着读都一样，所以是回文串。
//func isPalindrome(s string) bool {
//	s = strings.ToLower(s)
//	bs := []byte(s)
//	z := []byte{}
//	f := []byte{}
//	lenBs := len(bs)
//	for i := 0; i < lenBs; i++ {
//		if (bs[i] >= 'a' && bs[i] <= 'z') || (bs[i] >= '0' && bs[i] <= '9') {
//			z = append(z, bs[i])
//		}
//		if (bs[lenBs-i-1] >= 'a' && bs[lenBs-i-1] <= 'z') || (bs[lenBs-i-1] >= '0' && bs[lenBs-i-1] <= '9') {
//			f = append(f, bs[lenBs-i-1])
//		}
//	}
//	fmt.Println(string(z))
//	fmt.Println(string(f))
//	if string(z) == string(f) {
//		return true
//	}
//	return false
//}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := len(s) - 1
	for start <= end {
		if !isAlNum(s[start]) {
			start += 1
			continue
		}
		if !isAlNum(s[end]) {
			end -= 1
			continue
		}
		if s[start] != s[end] {
			return false
		} else {
			start += 1
			end -= 1
		}
	}

	return true
}

func isAlNum(s byte) bool {
	if (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9') {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome("abaa"))
}
