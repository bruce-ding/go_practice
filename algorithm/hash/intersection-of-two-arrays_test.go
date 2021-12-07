package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 349. 两个数组的交集
// 力扣题目链接(https://leetcode-cn.com/problems/happy-number/)
// 给定两个数组，编写一个函数来计算它们的交集。

// 示例 1：
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2]

// 示例 2：
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[9,4]

// 说明：

// 输出结果中的每个元素一定是唯一的。
// 我们可以不考虑输出结果的顺序。

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		//如果存在于上一个数组中，则加入结果集，并清空该set值
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}

func TestIntersection(t *testing.T) {
	// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
	// 输出：[2]
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	res := intersection(nums1, nums2)
	assert.Equal(t, []int{2}, res)

	// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
	// 输出：[9,4]
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	res = intersection(nums1, nums2)
	assert.Equal(t, []int{9, 4}, res)
}
