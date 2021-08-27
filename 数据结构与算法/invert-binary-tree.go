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

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left // 中
	invertTree(root.Left)                         // 左
	invertTree(root.Right)                        // 右
	return root
}

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode) // 中

		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			stack.PushBack(node.Right) // 右
		}
		if node.Left != nil {
			stack.PushBack(node.Left) // 左
		}

		printList(stack)

	}
	return root
}

func main() {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)
	root.Right = NewTreeNode(3)
	root.Right.Left = NewTreeNode(6)
	root.Right.Right = NewTreeNode(7)

	fmt.Println("before invert")
	fmt.Println(levelOrderTraversal(root))

	res := invertTree1(root)
	fmt.Println("after invert")
	fmt.Println(levelOrderTraversal(res))
}

func printList(l *list.List) {
	cur := l.Front()
	for cur != nil {
		fmt.Printf("%d -> ", cur.Value.(*TreeNode).Val)
		cur = cur.Next()
	}
	fmt.Println()
}

func levelOrderTraversal(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil { //防止为空
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) //出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val) //将值加入本层切片中
		}
		res = append(res, tmpArr) //放入结果集
		tmpArr = []int{}          //清空层的数据
	}
	return res
}
