package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftDepth := getDepth(node.Left)   // 左
	rightDepth := getDepth(node.Right) // 右
	// 中
	// 当一个左子树为空，右不为空，这时并不是最低点
	if node.Left == nil && node.Right != nil {
		return 1 + rightDepth
	}
	// 当一个右子树为空，左不为空，这时并不是最低点
	if node.Left != nil && node.Right == nil {
		return 1 + leftDepth
	}
	result := 1 + min(leftDepth, rightDepth)
	return result
}

func minDepth(root *TreeNode) int {
	return getDepth(root)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}
	if root.Left != nil && root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	return 1 + min(minDepth(root.Left), minDepth(root.Right))
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	// root.Right.Left = &TreeNode{Val: 6}
	// root.Right.Right = &TreeNode{Val: 7}

	res := minDepth1(root)
	fmt.Println(res)
}
