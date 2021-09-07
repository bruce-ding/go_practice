package main

// 题目链接：https://leetcode-cn.com/problems/convert-bst-to-greater-tree/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

var pre int // 记录前一个节点的数值
func traversal(cur *TreeNode) { // 右中左遍历
	if cur == nil {
		return
	}
	traversal(cur.Right)
	cur.Val += pre
	pre = cur.Val
	traversal(cur.Left)
}

func convertBST(root *TreeNode) *TreeNode {
	pre = 0
	traversal(root)
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

	res := convertBST(root)
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
