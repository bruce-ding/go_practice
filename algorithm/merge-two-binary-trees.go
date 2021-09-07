package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

//前序遍历（递归遍历，跟105 106差不多的思路）
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var value int
	var nullNode *TreeNode //空node，便于遍历
	nullNode = &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil}
	switch {
	case t1 == nil && t2 == nil:
		return nil //终止条件
	default: //如果其中一个节点为空，则将该节点置为nullNode，方便下次遍历
		if t1 == nil {
			value = t2.Val
			t1 = nullNode
		} else if t2 == nil {
			value = t1.Val
			t2 = nullNode
		} else {
			value = t1.Val + t2.Val
		}
	}
	root := &TreeNode{ //构造新的二叉树
		Val:   value,
		Left:  mergeTrees(t1.Left, t2.Left),
		Right: mergeTrees(t1.Right, t2.Right)}
	return root
}

// 前序遍历简洁版
func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees1(root1.Left, root2.Left)
	root1.Right = mergeTrees1(root1.Right, root2.Right)
	return root1
}

// 迭代版本
func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	queue := make([]*TreeNode, 0)
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue = append(queue, root1)
	queue = append(queue, root2)

	for size := len(queue); size > 0; size = len(queue) {
		node1 := queue[0]
		queue = queue[1:]
		node2 := queue[0]
		queue = queue[1:]
		node1.Val += node2.Val
		// 左子树都不为空
		if node1.Left != nil && node2.Left != nil {
			queue = append(queue, node1.Left)
			queue = append(queue, node2.Left)
		}
		// 右子树都不为空
		if node1.Right != nil && node2.Right != nil {
			queue = append(queue, node1.Right)
			queue = append(queue, node2.Right)
		}
		// 树 1 的左子树为 nil，直接接上树 2 的左子树
		if node1.Left == nil {
			node1.Left = node2.Left
		}
		// 树 1 的右子树为 nil，直接接上树 2 的右子树
		if node1.Right == nil {
			node1.Right = node2.Right
		}
	}
	return root1
}

func main() {
	t1 := NewTreeNode(1)
	t1.Left = NewTreeNode(3)
	t1.Left.Left = NewTreeNode(5)
	t1.Right = NewTreeNode(2)

	t2 := NewTreeNode(2)
	t2.Left = NewTreeNode(1)
	t1.Left.Right = NewTreeNode(4)
	t2.Right = NewTreeNode(3)
	t2.Right.Right = NewTreeNode(7)

	root := mergeTrees1(t1, t2)
	inorderTraversal(root)
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	inorderTraversal(root.Left)
	inorderTraversal(root.Right)
}
