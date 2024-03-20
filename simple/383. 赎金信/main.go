package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	m := map[byte]int{}
	for _, v := range []byte(magazine) {
		m[v] += 1
	}
	for _, v := range []byte(ransomNote) {
		m[v] -= 1
		if m[v] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}
