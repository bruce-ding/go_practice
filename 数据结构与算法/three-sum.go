package main

import (
	"fmt"
	"sort"
)

/**
排序 + 双指针
本题的难点在于如何去除重复解。

算法流程：
特判，对于数组长度 n，如果数组长度小于 3，返回 []。
对数组进行排序。
遍历排序后数组：
若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
对于重复元素：跳过，避免出现重复解
令左指针 L=i+1，右指针 R=n-1，当 L<R 时，执行循环：
当 nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,R 移到下一位置，寻找新的解
若和大于 0，说明 nums[R] 太大，R 左移
若和小于 0，说明 nums[L] 太小，L 右移
**/

func threeSum(nums []int, target int) [][]int {
	res := [][]int{}
	n := len(nums)
	// 对于数组长度 n，如果数组长度小于 3，直接返回
	if n < 3 {
		return res
	}
	sort.Ints(nums) // O(nlgn)

	for i := 0; i < n; i++ {
		// 若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果
		if nums[i] > target {
			return res
		}

		// 对于重复元素：跳过，避免出现重复解
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				// 判断左界和右界是否和下一位置重复，去除重复解
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// -1, 0, 1, 2, -1, -4   target = 0
				// -4, -1, -1, 0, 1, 2   target = 0
				// nums[0] + nums[1] + nums[4]
				left++
				right--
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	target := 0
	res := threeSum(nums, target)
	fmt.Println(res)
}
