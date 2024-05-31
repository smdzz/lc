//  各位相加

package main

import "fmt"

//func addDigits(num int) int {
//	if num < 10 {
//		return num
//	}
//	for num >= 10 {
//		s := 0
//
//		for num != 0 {
//			s += num % 10
//			num /= 10
//		}
//		num = s
//	}
//	return num
//}

//func addDigits(num int) int {
//	if num == 0 {
//		return 0
//	}
//	if num%9 == 0 {
//		return 9
//	}
//	return num % 9
//}

func addDigits(num int) int {
	return (num-1)%9 + 1
}

func main() {
	fmt.Println(addDigits(38))
}
