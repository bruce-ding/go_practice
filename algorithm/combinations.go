package main

import "fmt"

var result [][]int // 存放符合条件结果的集合
var path []int     // 用来存放符合条件的结果

func backtracking(n, k, start int) {
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		result = append(result, temp) // 处理节点
		return
	}

	fmt.Println("start:", start)
	fmt.Println("i <=:", n-(k-len(path))+1)

	for i := start; i <= n-(k-len(path))+1; i++ {
		path = append(path, i) // 处理节点
		fmt.Println("path:", path)

		backtracking(n, k, i+1)   // 递归
		path = path[:len(path)-1] // 回溯，撤销处理的节点
	}
}

func combine(n, k int) [][]int {
	backtracking(n, k, 1)
	return result
}

// var res [][]int

// func combine(n int, k int) [][]int {
// 	res = [][]int{}
// 	if n <= 0 || k <= 0 || k > n {
// 		return res
// 	}
// 	backtrack(n, k, 1, []int{})
// 	return res
// }

// func backtrack(n, k, start int, track []int) {
// 	if len(track) == k {
// 		temp := make([]int, k)
// 		copy(temp, track)
// 		res = append(res, temp)
// 	}
// 	if len(track)+n-start+1 < k {
// 		return
// 	}
// 	for i := start; i <= n; i++ {
// 		track = append(track, i)
// 		backtrack(n, k, i+1, track)
// 		track = track[:len(track)-1]
// 	}
// }

func combine1(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(temp)+(n-cur+1) < k {
			return
		}
		// 记录合法的答案
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}

func main() {
	// 输入: n = 4, k = 2
	res := combine(4, 2)
	fmt.Println(res)
}
