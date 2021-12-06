package array

//59.螺旋矩阵II
//力扣题目链接(https://leetcode-cn.com/problems/spiral-matrix-ii/)
//
//给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
//
//示例:
//
//输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]

import "fmt"

func spiralMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	size := n * n
	step := 0
	for step < size {
		for i := left; i <= right && step < size; i++ {
			matrix[top][i] = step + 1
			step++
		}
		top++

		for i := top; i <= bottom && step < size; i++ {
			matrix[i][right] = step + 1
			step++
		}
		right--

		for i := right; i >= left && step < size; i-- {
			matrix[bottom][i] = step + 1
			step++
		}
		bottom--

		for i := bottom; i >= top && step < size; i-- {
			matrix[i][left] = step + 1
			step++
		}
		left++
	}

	return matrix
}

// 1  2  3
// 8  9  4
// 7  6  5
func main() {
	matrix := spiralMatrix(3)
	fmt.Println(matrix)
}
