//  单词规律

package main

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	patternToWord := map[byte]string{}
	wordToPattern := map[string]byte{}
	sl := strings.Split(s, " ")
	if len(pattern) != len(sl) {
		return false
	}
	for i := range pattern {
		if word, ok := patternToWord[pattern[i]]; ok {
			if word != sl[i] {
				return false
			}
		} else {
			if p, pok := wordToPattern[sl[i]]; pok {
				if p != pattern[i] {
					return false
				}
			}
			patternToWord[pattern[i]] = sl[i]
			wordToPattern[sl[i]] = pattern[i]
		}
	}
	return true
}

func main() {

}
