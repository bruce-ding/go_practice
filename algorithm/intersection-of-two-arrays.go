package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] = 1
	}
	var res []int
	// 利用count>0，实现重复值只拿一次放入返回结果中
	for _, v := range nums2 {
		if count, ok := m[v]; ok && count > 0 {
			res = append(res, v)
			m[v]--
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	res := intersection(nums1, nums2)
	fmt.Println(res)
}
