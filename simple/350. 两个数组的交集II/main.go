package main

func intersect(nums1 []int, nums2 []int) []int {
	nums1Map := map[int]int{}
	var res []int
	for _, v := range nums1 {
		if _, ok := nums1Map[v]; ok {
			nums1Map[v] += 1
		} else {
			nums1Map[v] = 1
		}
	}
	nums2Map := map[int]int{}
	for _, v := range nums2 {
		if _, ok := nums2Map[v]; ok {
			nums2Map[v] += 1
		} else {
			nums2Map[v] = 1
		}
	}
	for k, v1 := range nums1Map {
		if v2, ok := nums2Map[k]; ok {
			min := 0
			if v1 > v2 {
				min = v2
			} else {
				min = v1
			}
			for i := 0; i < min; i++ {
				res = append(res, k)
			}
		}
	}
	return res
}

func main() {

}
