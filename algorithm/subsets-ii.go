package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/subsets-ii/

var result [][]int
var path []int

func backtracking(nums []int, startIndex int) {
	temp := make([]int, len(path))
	copy(temp, path)
	result = append(result, temp) // 收集子集，要放在终止添加的上面，否则会漏掉自己

	for i := startIndex; i < len(nums); i++ {
		// 而我们要对同一树层使用过的元素进行跳过

		if i > startIndex && nums[i] == nums[i-1] { // 注意这里使用i > startIndex
			continue
		}
		path = append(path, nums[i])
		backtracking(nums, i+1)
		path = path[:len(path)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums) // 去重需要排序
	backtracking(nums, 0)
	return result
}

func main() {
	// 示例: 输入: [1,2,2] 输出: [ [2], [1], [1,2,2], [2,2], [1,2], [] ]
	nums := []int{1, 2, 2}
	res := subsetsWithDup(nums)
	fmt.Println(res)
}
