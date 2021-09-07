package main

import (
	"fmt"
	"math"
)

func minSubArrayLen2(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= s {
				ans = min(ans, j-i+1)
				break
			}
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubArrayLen(target int, nums []int) int {
	i := 0               // 滑动窗口起始位置
	length := len(nums)  // 滑动窗口的长度
	sum := 0             // 滑动窗口数值之和
	result := length + 1 // 初始化返回长度为length+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < length; j++ {
		sum += nums[j]
		// 注意这里使用while，每次更新 i（起始位置），并不断比较子序列是否符合条件
		for sum >= target {
			subLength := j - i + 1 // 取子序列的长度
			if subLength < result {
				result = subLength
			}
			sum -= nums[i] // 这里体现出滑动窗口的精髓之处，不断变更i（子序列的起始位置）,和现在需要减去nums[i]
			i++            // 这里体现出滑动窗口的精髓之处，不断变更i（子序列的起始位置）
		}
	}
	// 如果result没有被赋值的话，就返回0，说明没有符合条件的子序列
	if result == length+1 {
		return 0
	} else {
		return result
	}
}

func minimumSizeSubArraySum(target int, nums []int) int {
	length := len(nums)
	res := length + 1
	i := 0
	sum := 0
	for j := 0; j < length; j++ {
		sum += nums[j]
		for sum >= target {
			subLen := j - i + 1
			if subLen < res {
				res = subLen
			}
			sum -= nums[i]
			i++
		}
	}
	if res == length+1 {
		return 0
	} else {
		return res
	}
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	res := minimumSizeSubArraySum(target, nums)
	fmt.Println(res)
}
