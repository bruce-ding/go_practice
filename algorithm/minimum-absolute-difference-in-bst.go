package main

import (
	"fmt"
	"math"
)

// 题目地址：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func dfs(node *TreeNode, ans, pre int) int {
	if node == nil {
		return ans
	}
	ans = dfs(node.Left, ans, pre)
	if pre != -1 && node.Val-pre < ans {
		ans = node.Val - pre
	}
	pre = node.Val
	ans = dfs(node.Right, ans, pre)
	return ans
}

func getMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	// var dfs func(*TreeNode)
	// dfs = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	dfs(node.Left)
	// 	if pre != -1 && node.Val-pre < ans {
	// 		ans = node.Val - pre
	// 	}
	// 	pre = node.Val
	// 	dfs(node.Right)
	// }
	// dfs(root)
	return dfs(root, ans, pre)
	// return ans
}

func main() {
	root := NewTreeNode(4)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(1)
	root.Left.Right = NewTreeNode(3)
	root.Right = NewTreeNode(7)

	res := getMinimumDifference(root)
	fmt.Println(res)

}
