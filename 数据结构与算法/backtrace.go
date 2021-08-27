package main

import "fmt"

//result = []
//func backtrack(路径, 选择列表) {
//  if 满足结束条件 {
//    result.add(路径)
//    return
//  }
//  for 选择 in 选择列表 {
//    做选择
//    backtrack(路径, 选择列表)
//    撤销选择
//  }
//}

// leetcode 用例测试中，直接在 Go语言 中使用全局变量会出错
// 解决方法：在函数中初始化该变量
var res [][]int
var visited = map[int]bool{}

func permute(nums []int) [][]int {
	// 初识化全局变量
	//res = [][]int{}
	var track []int
	backtrack(nums, track)
	return res
}

func backtrack(nums []int, track []int) {
	if len(nums) == len(track) {
		//temp := make([]int, len(track))
		//copy(temp, track)
		res = append(res, track)
		return
	}

	for _, num := range nums {
		// 排除不合法的选择
		if visited[num] {
			continue
		}
		track = append(track, num)
		visited[num] = true

		//fmt.Println(num)
		backtrack(nums, track)
		fmt.Println("before:", track)
		track = track[:len(track)-1] // 去掉最后一个元素
		fmt.Println("after:", track)
		visited[num] = false
		fmt.Println("visited:", visited)
	}
}

func main() {
	nums := []int{1,2,3}
	res := permute(nums)
	fmt.Println(res)
}
