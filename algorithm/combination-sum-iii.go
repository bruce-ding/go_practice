package main

// 链接：https://leetcode-cn.com/problems/combination-sum-iii/

import "fmt"

var result [][]int // 存放符合条件结果的集合
var path []int     // 用来存放符合条件的结果

func backtracking(targetSum, k, sum, start int) {
	if sum > targetSum { // 剪枝操作
		return // 如果path.size() == k 但sum != targetSum 直接返回
	}
	if len(path) == k {
		if sum == targetSum {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp) // 处理节点
		}
		return
	}
	for i := start; i <= 9-(k-len(path))+1; i++ { // 剪枝
		sum += i                             // 处理
		path = append(path, i)               // 处理
		backtracking(targetSum, k, sum, i+1) // 注意i+start
		sum -= i                             // 回溯
		path = path[:len(path)-1]            // 回溯
	}
}

func combinationSum3(k, n int) [][]int {
	backtracking(n, k, 0, 1)
	return result
}

func main() {
	// 示例 1: 输入: k = 3, n = 7 输出: [[1,2,4]]
	// 示例 2: 输入: k = 3, n = 9 输出: [[1,2,6], [1,3,5], [2,3,4]]
	res := combinationSum3(3, 9)
	fmt.Println(res)
}
