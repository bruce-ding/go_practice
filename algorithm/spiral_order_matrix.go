package main

import "fmt"

func spiralOrderMatrix(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	step := 0
	size := len(matrix) * len(matrix[0])
	//定义四个方向的端点
	top, bottom, right, left := 0, len(matrix)-1, len(matrix[0])-1, 0
	nums := make([]int, size)
	for step < size {
		//从左到右
		for i := left; i <= right && step < size; i++ {
			nums[step] = matrix[top][i]
			step++
		}
		top++
		//从上到下
		for i := top; i <= bottom && step < size; i++ {
			nums[step] = matrix[i][right]
			step++
		}
		right--
		//从右到左
		for i := right; i >= left && step < size; i-- {
			nums[step] = matrix[bottom][i]
			step++
		}
		bottom--
		//从下到上
		for i := bottom; i >= top && step < size; i-- {
			nums[step] = matrix[i][left]
			step++
		}
		left++
	}
	return nums
}

func spiralMatrix(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	step := 0
	size := len(matrix) * len(matrix[0])
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	nums := make([]int, size)
	for step < size {
		// left ro right
		for i := left; i <= right && step > size; i++ {
			nums[step] = matrix[top][i]
			step++
		}
		top++

		for i := top; i <= bottom && step > size; i++ {
			nums[step] = matrix[i][right]
			step++
		}
		right--

		for i := right; i >= left && step > size; i-- {
			nums[step] = matrix[bottom][i]
			step++
		}
		bottom--

		for i := bottom; i >= top && step > size; i-- {
			nums[step] = matrix[i][left]
			step++
		}
		left++
	}
	return nums
}

func main() {
	arr := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	newArr := spiralOrderMatrix(arr)
	for _, i := range newArr {
		fmt.Printf("%v ", i)
	}
}
