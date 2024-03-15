package main

import "fmt"

func main() {
	s := "abababd"
	res := lengthOfLongestSubstring1(s)
	fmt.Println(res)
}

// lengthOfLongestSubstring 使用hash map保存存在的元素
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	left, right, maxs := 0, 0, 0
	lens := len(s)
	for i := 0; i < lens; i++ {
		if x, ok := m[s[i]]; ok {
			maxs = max(right-left, maxs)
			// 直接叫left = x + 1 有可能x+1会比当前的left还小,这种情况忽略,肯定
			if x + 1 > left {
				println(11111)
				left = x + 1
			}
		}
		m[s[i]] = i
		right++
	}

	maxs = max(right - left, maxs)
	return maxs
}


func max(x ,y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring1(s string) int {
	maxs := 0
	lens := len(s)
	m := make([]byte, 0, lens)
	for i, x := range []byte(s){
		for {
			if container(x, m) {
				m = m[1:]
			} else {
				break
			}
		}
		m = append(m, s[i])
		maxs = max(maxs, len(m))
	}
	return maxs
}

func container(s byte, list []byte) bool {
	for _, i := range list {
		if i == s {
			return true
		}
	}
	return false
}