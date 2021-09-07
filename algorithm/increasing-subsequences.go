package main

import "fmt"

// https://leetcode-cn.com/problems/increasing-subsequences/

var result [][]int
var path []int

func backtracking(nums []int, startIndex int) {
	if len(path) > 1 {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp) // 收集子集
	}
	var used = [201]bool{} // 这里使用数组来进行去重操作，题目说数值范围[-100, 100]
	for i := startIndex; i < len(nums); i++ {
		if len(path) > 0 && nums[i] < path[len(path)-1] || used[nums[i]+100] {
			continue
		}
		used[nums[i]+100] = false // 记录这个元素在本层用过了，本层后面不能再用了
		path = append(path, nums[i])
		backtracking(nums, i+1)
		path = path[:len(path)-1]
	}
}

func findSubsequences(nums []int) [][]int {
	backtracking(nums, 0)
	return result
}

func main() {
	//输入: [4, 6, 7, 7] 输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
	nums := []int{4, 6, 7, 7}
	res := findSubsequences(nums)
	fmt.Println(res)
}
