package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	num := findRepeatNumber(nums)
	fmt.Println("重复的一个数字为: ", num)
}
