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
	leftDepth := getDepth(node.Left)        // 左
	rightDepth := getDepth(node.Right)      // 右
	depth := 1 + max(leftDepth, rightDepth) // 中
	fmt.Println(leftDepth, rightDepth, depth)

	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	return getDepth(root)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 2}
	//root.Right.Left = &TreeNode{Val: 4}
	//root.Right.Right = &TreeNode{Val: 3}
	res := maxDepth(root)
	fmt.Println(res)
}
