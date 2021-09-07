package main

// 链接：https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 迭代法
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	var pnode *TreeNode
	for node != nil {
		if val > node.Val {
			pnode = node
			node = node.Right
		} else {
			pnode = node
			node = node.Left
		}
	}
	if val > pnode.Val {
		pnode.Right = &TreeNode{Val: val}
	} else {
		pnode.Left = &TreeNode{Val: val}
	}
	return root
}

func main() {
	root := NewTreeNode(4)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(1)
	root.Left.Right = NewTreeNode(3)
	root.Right = NewTreeNode(7)

	res := insertIntoBST1(root, 5)
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
