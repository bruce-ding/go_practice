package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//迭代法前序遍历
/**
前序遍历：中左右
压栈顺序：右左中
 **/
func preorderTraversal(root *TreeNode) []int {
	var stack = list.New()
	res := []int{}
	if root == nil {
		return res
	}

	stack.PushBack(root)

	for stack.Len() > 0 {
		e := stack.Back() // 中
		stack.Remove(e)
		node := e.Value.(*TreeNode)

		res = append(res, node.Val)

		if node.Right != nil {
			stack.PushBack(node.Right) // 右（空节点不入栈）
		}
		if node.Left != nil {
			stack.PushBack(node.Left) // 左（空节点不入栈）
		}
	}
	return res
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	res := preorderTraversal(root)
	fmt.Println(res)
}
