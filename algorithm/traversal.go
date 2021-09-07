package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrderTraversal(root.Left)
	preOrderTraversal(root.Right)
}

func midOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	midOrderTraversal(root.Left)
	fmt.Println(root.Val)
	midOrderTraversal(root.Right)
}

func postorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversal(root.Left)
	postorderTraversal(root.Right)
	fmt.Println(root.Val)
}

func layerOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 9}
	root.Left.Right.Right = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 8}
	preOrderTraversal(root)
}
