package main

import (
	"fmt"
)

// 题目链接：https://leetcode-cn.com/problems/combination-sum/

var result [][]int
var path []int

func backtracking(candidates []int, target, sum, startIndex int) {
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp) // 处理节点
		return
	}

	// 如果 sum + candidates[i] > target 就终止遍历
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		sum += candidates[i]
		path = append(path, candidates[i])
		backtracking(candidates, target, sum, i)
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	// sort(candidates.begin(), candidates.end()); // 需要排序
	backtracking(candidates, target, 0, 0)
	return result
}

func main() {
	// 示例 1： 输入：candidates = [2,3,6,7], target = 7, 所求解集为： [ [7], [2,2,3] ]
	// 示例 2： 输入：candidates = [2,3,5],   target = 8, 所求解集为： [ [2,2,2,2], [2,3,3], [3,5] ]
	candidates := []int{2, 3, 5}
	target := 8
	res := combinationSum(candidates, target)
	fmt.Println(res)

}
