//  有效的字母异位词

package main

import "fmt"

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	m := map[byte]int{}
//	for _, v := range []byte(s) {
//		m[v]++
//	}
//	for _, v := range []byte(t) {
//		m[v]--
//		if m[v] < 0 {
//			return false
//		}
//	}
//	return true
//}

func isAnagram(s string, t string) bool {
	var sl, tl [26]int
	for _, v := range s {
		sl[v-'a']++
	}
	for _, v := range t {
		tl[v-'a']++
	}
	return sl == tl
}

func main() {
	fmt.Println(isAnagram("rat", "car"))
}
