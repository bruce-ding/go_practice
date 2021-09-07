package main

import "fmt"

// 题目地址：https://leetcode-cn.com/problems/maximum-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	//首先找到最大值
	index := findMax(nums)
	//然后构造二叉树
	root := &TreeNode{
		Val:   nums[index],
		Left:  constructMaximumBinaryTree(nums[:index]),   //左半边
		Right: constructMaximumBinaryTree(nums[index+1:]), //右半边
	}
	return root
}

func findMax(nums []int) int {
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return index
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
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
