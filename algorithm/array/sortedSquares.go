package main

// 977.有序数组的平方
// 力扣题目链接(https://leetcode-cn.com/problems/squares-of-a-sorted-array/)

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

// 示例 1： 输入：nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

// 示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]

import "fmt"

func sortedSquares(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1
	ans := make([]int, n)
	for i <= j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > rm {
			ans[k] = lm
			i++
		} else {
			ans[k] = rm
			j--
		}
		k--
	}
	return ans
}

func main() {
	// nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

	nums := []int{-4, -1, 0, 3, 10}
	res := sortedSquares(nums)
	fmt.Println(res)
}
