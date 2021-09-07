package main

import "fmt"

/**
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。

叶子节点 是指没有子节点的节点。

示例 1：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
**/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traversal(cur *TreeNode, count int) bool {
	if cur.Left == nil && cur.Right == nil && count == 0 { // 遇到叶子节点，并且计数为0
		return true
	}
	if cur.Left == nil && cur.Right == nil { // 遇到叶子节点直接返回
		return false
	}

	if cur.Left != nil { // 左
		count -= cur.Left.Val // 递归，处理节点;
		if traversal(cur.Left, count) {
			return true
		}
		count += cur.Left.Val // 回溯，撤销处理结果
	}
	if cur.Right != nil { // 右
		count -= cur.Right.Val // 递归，处理节点;
		if traversal(cur.Right, count) {
			return true
		}
		count += cur.Right.Val // 回溯，撤销处理结果
	}
	return false
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return traversal(root, sum-root.Val)
}

func main() {
	// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}

	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 14}
	root.Right.Right.Right = &TreeNode{Val: 1}

	res := hasPathSum(root, 22)
	fmt.Println(res)
}
