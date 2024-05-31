//  有效的完全平方数

package main

//func isPerfectSquare(num int) bool {
//	x := int(math.Sqrt(float64(num)))
//	return x*x == num
//}

//// 暴力
//func isPerfectSquare(num int) bool {
//	for i := 1; i <= num; i++ {
//		if i*i == num {
//			return true
//		} else if i*i > num {
//			return false
//		}
//	}
//	return false
//}

// 二分查找
func isPerfectSquare(num int) bool {
	left, right := 1, num
	for left <= right {
		middle := (left + right) >> 1
		if middle*middle < num {
			left = middle + 1
		} else if middle*middle > num {
			right = middle - 1
		} else {
			return true
		}
	}
	return false
}

// 牛顿迭代法
//func isPerfectSquare(num int) bool {
//	x0 := float64(num)
//	for {
//		x1 := (x0 + float64(num)/x0) / 2
//		if x0-x1 < 1e-6 {
//			x := int(x0)
//			return x*x == num
//		}
//		x0 = x1
//	}
//}

func main() {

}
