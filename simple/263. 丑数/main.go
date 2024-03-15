package main

import "fmt"

func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}
	for n > 1 || n*-1 > 1 {
		if n%3 != 0 && n%5 != 0 && n%2 != 0 {
			return false
		}
		if n%3 == 0 {
			n /= 3
			continue
		}
		if n%5 == 0 {
			n /= 5
			continue
		}
		if n%2 == 0 {
			n /= 2
			continue
		}
	}
	return true
}

func main() {
	fmt.Println(-10 / 2)
	fmt.Println(isUgly(-2147483648))
}
