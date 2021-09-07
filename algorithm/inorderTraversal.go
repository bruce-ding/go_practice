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

func inorderTraversal(root *TreeNode) []int {
	var stack = list.New()
	res := []int{}
	// if root == nil {
	// 	return res
	// }

	cur := root
	for cur != nil || stack.Len() > 0 {
		if cur != nil { // 指针来访问节点，访问到最底层
			stack.PushBack(cur) // 将访问的节点放进栈
			cur = cur.Left      // 左
		} else {
			e := stack.Back() // 从栈里弹出的数据，就是要处理的数据（放进result数组里的数据）
			stack.Remove(e)
			cur := e.Value.(*TreeNode)

			res = append(res, cur.Val) // 中
			cur = cur.Right            // 右
		}
	}
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}
	stack := list.New()
	node := root
	//先将所有左节点找到，加入栈中
	for node != nil {
		stack.PushBack(node)
		node = node.Left
	}
	//其次对栈中的每个节点先弹出加入到结果集中，再找到该节点的右节点的所有左节点加入栈中
	for stack.Len() > 0 {
		e := stack.Back()
		node := e.Value.(*TreeNode)
		stack.Remove(e)
		//找到该节点的右节点，再搜索他的所有左节点加入栈中
		res = append(res, node.Val)
		node = node.Right
		for node != nil {
			stack.PushBack(node)
			node = node.Left
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

	res := inorderTraversal1(root)
	fmt.Println(res)
}
