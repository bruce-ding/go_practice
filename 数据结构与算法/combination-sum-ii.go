package main

import (
	"fmt"
	"sort"
)

// 题目链接：https://leetcode-cn.com/problems/combination-sum-ii/

var result [][]int
var path []int

func backtracking(candidates []int, target int, sum int, startIndex int, used []bool) {
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp) // 处理节点
		return
	}
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		// used[i - 1] == true，说明同一树支candidates[i - 1]使用过
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
		// 要对同一树层使用过的元素进行跳过
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		sum += candidates[i]
		path = append(path, candidates[i])
		used[i] = true
		backtracking(candidates, target, sum, i+1, used) // 和39.组合总和的区别1，这里是i+1，每个数字在每个组合中只能使用一次
		used[i] = false
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	used := make([]bool, len(candidates))
	// 首先把给candidates排序，让其相同的元素都挨在一起。
	sort.Ints(candidates)
	backtracking(candidates, target, 0, 0, used)
	return result
}

func main() {
	// 示例 1: 输入: candidates = [10,1,2,7,6,1,5], target = 8, 所求解集为: [ [1,7], [1,2,5], [2,6], [1,1,6] ]
	// 示例 2: 输入: candidates = [2,5,2,1,2], target = 5, 所求解集为: [ [1,2,2], [5] ]
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	res := combinationSum2(candidates, target)
	fmt.Println(res)
}
