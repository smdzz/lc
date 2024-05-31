// 快乐数

package main

import "fmt"

// 快慢指针法
//func isHappy(n int) bool {
//	slow, fast := n, step(n)
//	for fast != 1 && fast != slow {
//		slow = step(slow)
//		fast = step(step(fast))
//	}
//	return fast == 1
//}

// 记录每个过程中的数
func isHappy(n int) bool {
	tmp := map[int]struct{}{n: {}}
	for {
		n = step(n)
		if n == 1 {
			return true
		}
		if _, ok := tmp[n]; ok {
			return n == 1
		}
		tmp[n] = struct{}{}
	}
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
func main() {
	fmt.Println(isHappy(19))
}
