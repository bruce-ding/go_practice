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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil { // 第一种情况：没找到删除的节点，遍历到空节点直接返回了
		return root
	}
	if root.Val == key {
		// 第二种情况：左右孩子都为空（叶子节点），直接删除节点， 返回NULL为根节点
		// 第三种情况：其左孩子为空，右孩子不为空，删除节点，右孩子补位 ，返回右孩子为根节点
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil { // 第四种情况：其右孩子为空，左孩子不为空，删除节点，左孩子补位，返回左孩子为根节点
			return root.Left
		} else { // 第五种情况：左右孩子节点都不为空，则将删除节点的左子树放到删除节点的右子树的最左面节点的左孩子的位置
			// 并返回删除节点右孩子为新的根节点。
			cur := root.Right // 找右子树最左面的节点
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left // 把要删除的节点（root）左子树放在cur的左孩子的位置
			root = root.Right    // 返回旧root的右孩子作为新root
			return root
		}
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
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

	res := deleteNode(root, 7)
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
