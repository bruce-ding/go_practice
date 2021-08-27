package main

import (
	"fmt"
	"sort"
)

/**
排序 + 双指针
本题的难点在于如何去除重复解。

算法流程：
特判，对于数组长度 n，如果数组长度小于 4，返回 []。
对数组进行排序。
遍历排序后数组：
若 nums[i]>target：因为已经排序好，所以后面不可能有四个数加和等于 target，直接返回结果。
对于重复元素：跳过，避免出现重复解
令左指针 left=i+1，右指针 right=n-1，当 left<right 时，执行循环：
当 nums[i]+nums[left]+nums[right] == target，执行循环，判断左界和右界是否和下一位置重复，去除重复解。
并同时将 left,right 移到下一位置，寻找新的解
若和大于 target，说明 nums[right] right 左移
若和小于 target，说明 nums[left] 太小，left 右移
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
		// 若 nums[i]>target：因为已经排序好，所以后面不可能有四个数加和等于 target，直接返回结果
		if nums[i] > target {
			return res
		}

		// 对于重复元素：跳过，避免出现重复解
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n; j++ {
			if nums[j] > target {
				return res
			}
			// 对于重复元素：跳过，避免出现重复解
			if i > i+1 && nums[j] == nums[j-1] {
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

func main() {
	// nums := []int{1, 0, -1, 0, -2, 2}
	nums := []int{1, 2, 3, 4, 5, 6}
	target := 18
	res := fourSum(nums, target)
	fmt.Println(res)
}
