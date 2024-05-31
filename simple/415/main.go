// 字符串相加

package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	res := []byte{}
	needAdd := false
	for n, m := len(num1)-1, len(num2)-1; n >= 0 || m >= 0; n, m = n-1, m-1 {
		x, y := '0', '0'
		if n >= 0 {
			x = int32(num1[n])
		}
		if m >= 0 {
			y = int32(num2[m])
		}
		if needAdd {
			if (x+y-'0')+1 > '9' {
				res = append([]byte{byte(x + y - '9' - 1 + 1)}, res...)
				needAdd = true
			} else {
				res = append([]byte{byte(x + y - '0' + 1)}, res...)
				needAdd = false
			}
		} else {
			if (x + y - '0') > '9' {
				// 大于9之后要减10,所以要x+y-'9'之后在-1
				res = append([]byte{byte(x+y-'9') - 1}, res...)
				needAdd = true
			} else {
				res = append([]byte{byte(x + y - '0')}, res...)
				needAdd = false
			}
		}
	}
	if needAdd {
		res = append([]byte{'1'}, res...)
	}

	return string(res)
}

func main() {
	// 533
	fmt.Println(addStrings("1", "9"))
}
