package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	leftHeight := maxHeight(root.Left) + 1
	rightHeight := maxHeight(root.Right) + 1
	if abs(leftHeight-rightHeight) > 1 {
		return false
	}
	return true
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxHeight(root.Left), maxHeight(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	res := isBalanced(root)
	fmt.Println(res)
}
