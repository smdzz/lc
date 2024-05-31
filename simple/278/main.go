//  第一个错误的版本

package main

func isBadVersion(n int) bool {
	// xxxxxxx
	return false
}

func firstBadVersion(n int) int {
	begin, end := 1, n
	for begin < end {
		middle := (begin + end) >> 1
		if isBadVersion(middle) {
			end = middle
		} else {
			begin = middle + 1
		}
	}
	return begin
}

func main() {

}
