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

var maxDepth int // 全局变量 深度
var value int    // 全局变量 最终值

func findBottomLeftValue1(root *TreeNode) int {
	// 需要提前判断一下（不要这个if的话提交结果会出错，但执行代码不会。防止这种情况出现，故先判断是否只有一个节点）
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	findLeftValue(root, maxDepth)
	return value
}

func findLeftValue(root *TreeNode, depth int) {
	// 最左边的值在左边
	if root.Left == nil && root.Right == nil {
		if depth > maxDepth {
			value = root.Val
			maxDepth = depth
		}
	}
	// 递归
	if root.Left != nil {
		depth++
		findLeftValue(root.Left, depth)
		depth-- // 回溯
	}
	if root.Right != nil {
		depth++
		findLeftValue(root.Right, depth)
		depth-- // 回溯
	}
}

func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	var res int
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
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

	res := findBottomLeftValue(root)
	fmt.Println(res)
}
