package hash

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
18. 四数之和
力扣题目链接(https://leetcode-cn.com/problems/4sum/)
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。

示例 1：
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

示例 2：
输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]

提示：
1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109

解题思路：
三数之和的双指针解法是一层for循环num[i]为确定值，然后循环内有left和right下表作为双指针，找到nums[i] + nums[left] + nums[right] == 0。

四数之和的双指针解法是两层for循环nums[k] + nums[i]为确定值，依然是循环内有left和right下表作为双指针，找出nums[k] + nums[i] + nums[left] + nums[right] == target的情况，
三数之和的时间复杂度是O(n^2)，四数之和的时间复杂度是O(n^3) 。

那么一样的道理，五数之和、六数之和等等都采用这种解法。

对于15.三数之和双指针法就是将原本暴力O(n^3)的解法，降为O(n^2)的解法，四数之和的双指针解法就是将原本暴力O(n^4)的解法，降为O(n^3)的解法。
**/

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	n := len(nums)
	// 对于数组长度 n，如果数组长度小于 3，直接返回
	if n < 4 {
		return res
	}
	sort.Ints(nums) // O(nlgn)

	for i := 0; i < n; i++ {
		// 不能这样写,因为target可能是负数
		// if nums[i] > target {
		// 	return res
		// }

		// 对于重复元素：跳过，避免出现重复解
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n; j++ {
			// 不能这样写,因为target可能是负数
			// if nums[j] > target {
			// 	return res
			// }
			// 对于重复元素：跳过，避免出现重复解
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := n - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					// 判断左界和右界是否和下一位置重复，去除重复解
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum > target {
					right--
				} else {
					left++
				}
			}

		}

	}
	return res
}

func TestFourSum(t *testing.T) {
	// 输入：nums = [1,0,-1,0,-2,2], target = 0
	// 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	res := fourSum(nums, target)
	expected := [][]int{
		{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1},
	}
	assert.Equal(t, expected, res)
}
