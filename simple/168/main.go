// Excel表列名称

package main

import "fmt"

func convertToTitle(columnNumber int) string {
	s := []byte{}
	for columnNumber > 0 {
		columnNumber--
		s = append(s, 'A'+byte(columnNumber%26))
		columnNumber /= 26
	}
	for i, n := 0, len(s); i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
	return string(s)
}

func main() {
	fmt.Println(convertToTitle(701))
	// "FXSHRXW"
	fmt.Println(('Z'-'A'+1)*26 + ('Y' - 'A' + 1))
}
