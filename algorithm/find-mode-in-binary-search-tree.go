package main

import (
	"fmt"

	"github.com/sqs/goreturns/returns"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func findMode(root *TreeNode) []int {
	var history map[int]int
	var maxValue int
	var maxIndex int
	var result []int
	history = make(map[int]int)
	traversal(root, history)
	for k, value := range history {
		if value > maxValue {
			maxValue = value
			maxIndex = k
		}
	}
	for k, value := range history {
		if value == history[maxIndex] {
			result = append(result, k)
		}
	}
	return result
}

func traversal(root *TreeNode, history map[int]int) {
	if root.Left != nil {
		traversal(root.Left, history)
	}
	if value, ok := history[root.Val]; ok {
		history[root.Val] = value + 1
	} else {
		history[root.Val] = 1
	}
	if root.Right != nil {
		traversal(root.Right, history)
	}
}

func findMode1(root *TreeNode) []int {
	res := make([]int, 0) // 最红的结果数组
	count := 1 // 记录当前数字出现的次数
	max := 1   // 记录数字出现最多的次数
	var prev *TreeNode // 前一个节点
	var travel func(node *TreeNode) // 递归函数
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}
		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			max = count
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return res
}

func findMode2(root *TreeNode) []int {
	res := make([]int, 0)
	count := 1
	max := 1
	var prev *TreeNode
	var traversal func(node *TreeNode)

	traversal = func (node *TreeNode)  {
		if node == nil {
			return
		}
		traversal(node.Left)
		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}

		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			max = count
		}
		prev = node
		traversal(node.Right)
	}
	traversal(root)
	return res
}

func main() {
	root := NewTreeNode(4)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(1)
	root.Left.Right = NewTreeNode(3)
	root.Right = NewTreeNode(7)
	root.Right.Left = NewTreeNode(7)

	res := findMode(root)
	fmt.Println(res)
}
