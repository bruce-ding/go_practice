package main

// 题目链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func traversal(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	root := NewTreeNode(nums[mid])
	root.Left = traversal(nums, left, mid-1)
	root.Right = traversal(nums, mid+1, right)
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	root := traversal(nums, 0, len(nums)-1)
	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(nums)
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
