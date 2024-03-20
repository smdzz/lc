package main

import "fmt"

//func firstUniqChar(s string) int {
//	charMap := map[byte]int{}
//	for i, v := range []byte(s) {
//		if ei, ok := charMap[v]; ok {
//			if ei == -1 {
//				continue
//			} else {
//				charMap[v] = -1
//			}
//		} else {
//			charMap[v] = i
//		}
//	}
//	min := -1
//	for _, v := range charMap {
//		if v != -1 {
//			if min == -1 {
//				min = v
//			} else {
//				if v < min {
//					min = v
//				}
//			}
//		}
//	}
//	return min
//}

type pair struct {
	ch  byte
	pos int
}

func firstUniqChar(s string) int {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	q := []pair{}
	for i := range s {
		ch := s[i] - 'a'
		if pos[ch] == n {
			pos[ch] = i
			q = append(q, pair{ch, i})
		} else {
			pos[ch] = n + 1
			for len(q) > 0 && pos[q[0].ch] == n+1 {
				q = q[1:]
			}
		}
	}
	if len(q) > 0 {
		return q[0].pos
	}
	return -1
}

func main() {
	// 0
	fmt.Println(firstUniqChar("leetcode"))
	// 2
	fmt.Println(firstUniqChar("loveleetcode"))
	//-1
	fmt.Println(firstUniqChar("aabb"))
}
