package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	res := make([]string, 0, n)
	for x := 1; x <= n; x++ {
		if x%3 == 0 && x%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if x%3 == 0 {
			res = append(res, "Fizz")
		} else if x%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(x))
		}
	}
	return res
}

func main() {
	fmt.Println(fizzBuzz(15))
}
