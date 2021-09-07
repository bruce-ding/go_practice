package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNodesNum(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	leftNum := getNodesNum(cur.Left)   // 左
	rightNum := getNodesNum(cur.Right) // 右
	treeNum := leftNum + rightNum + 1  // 中
	return treeNum
}
func countNodes(root *TreeNode) int {
	return getNodesNum(root)
}

func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	res := countNodes(root)
	fmt.Println(res)
}
