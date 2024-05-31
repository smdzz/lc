//  两个数组的交集

package main

func intersection(nums1 []int, nums2 []int) []int {
	nums1Map := map[int]struct{}{}
	intersectionMap := map[int]struct{}{}
	for _, v := range nums1 {
		nums1Map[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := nums1Map[v]; ok {
			intersectionMap[v] = struct{}{}
		}
	}
	res := make([]int, 0, len(intersectionMap))
	for k, _ := range intersectionMap {
		res = append(res, k)
	}
	return res
}

func main() {

}
