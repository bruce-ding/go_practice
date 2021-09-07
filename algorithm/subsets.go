package main

import "fmt"

// https://leetcode-cn.com/problems/subsets/

var result [][]int
var path []int

func backtracking(nums []int, startIndex int) {
	temp := make([]int, len(path))
	copy(temp, path)
	result = append(result, temp) // 收集子集，要放在终止添加的上面，否则会漏掉自己

	if startIndex >= len(nums) { // 终止条件可以不加
		return
	}
	for i := startIndex; i < len(nums); i++ {
		path = append(path, nums[i])
		backtracking(nums, i+1)
		path = path[:len(path)-1]
	}
}

func subsets(nums []int) [][]int {
	backtracking(nums, 0)
	return result
}

func main() {
	// 示例: 输入: nums = [1,2,3] 输出: [ [3],   [1],   [2],   [1,2,3],   [1,3],   [2,3],   [1,2],   [] ]
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
}
