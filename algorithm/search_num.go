package main

import (
	"fmt"
	"sort"
)

func findNumberIn2DArray(matrix [][]int, target int) bool {
	//以左下角为原点
	i := len(matrix) - 1 //获取右下角y坐标
	j := 0               //获取右下角x坐标
	for i >= 0 {
		if j < len(matrix[i]) {
			if target < matrix[i][j] {
				i-- //小于target,向上查找
			} else if target > matrix[i][j] {
				j++ //大于targat,向右查找
			} else {
				return true
			}
		} else {
			return false //超出数组返回false
		}
	}
	return false //超出matrix返回false
}

func findNumberIn2DArray2(matrix [][]int, target int) bool {
	for _, nums := range matrix {
		i := sort.SearchInts(nums, target)
		if i < len(nums) && target == nums[i] {
			return true
		}
	}
	return false
}

func main() {
	arr := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	res := findNumberIn2DArray2(arr, 3)
	fmt.Printf("res: %v\n", res)

}
