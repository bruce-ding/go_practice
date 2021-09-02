package main

// 题目链接： https://leetcode-cn.com/problems/delete-node-in-a-bst/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func trimBST(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func main() {
	root := NewTreeNode(2)
	root.Left = NewTreeNode(1)
	root.Right = NewTreeNode(7)
	root.Right.Left = NewTreeNode(5)
	root.Right.Left.Left = NewTreeNode(4)
	root.Right.Left.Right = NewTreeNode(6)
	root.Right.Right = NewTreeNode(9)
	root.Right.Right.Left = NewTreeNode(8)
	root.Right.Right.Right = NewTreeNode(10)

	res := trimBST(root, 4, 8)
	inorderTraversal(res)
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	inorderTraversal(root.Left)
	inorderTraversal(root.Right)
}
