package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) == 1 {
		return true
	}
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if (s2t[x] > 0 && s2t[x] != y) || (t2s[y] > 0 && t2s[y] != x) {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("badc", "bada"))
}
