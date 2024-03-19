package main

import (
	"bytes"
	"fmt"
)

// a e i o u
func reverseVowels(s string) string {
	sb := []byte(s)
	lens := len(sb)
	vowels := []byte("aeiouAEIOU")
	left := 0
	right := lens - 1
	for left < right {
		if bytes.Contains(vowels, []byte{sb[left]}) && bytes.Contains(vowels, []byte{sb[right]}) {
			sb[left], sb[right] = sb[right], sb[left]
			right -= 1
			left += 1
		} else if bytes.Contains(vowels, []byte{sb[left]}) {
			right -= 1
		} else if bytes.Contains(vowels, []byte{sb[right]}) {
			left += 1
		} else {
			right -= 1
			left += 1
		}
	}
	return string(sb)
}

//// 官方题解
//func reverseVowels(s string) string {
//	t := []byte(s)
//	n := len(t)
//	i, j := 0, n-1
//	for i < j {
//		for i < n && !strings.Contains("aeiouAEIOU", string(t[i])) {
//			i++
//		}
//		for j > 0 && !strings.Contains("aeiouAEIOU", string(t[j])) {
//			j--
//		}
//		if i < j {
//			t[i], t[j] = t[j], t[i]
//			i++
//			j--
//		}
//	}
//	return string(t)
//}

func main() {
	//leotcede
	fmt.Println(reverseVowels("leetcode"))
	// holle
	fmt.Println(reverseVowels("hello"))
}
