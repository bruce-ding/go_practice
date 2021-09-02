package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftValue := sumOfLeftLeaves(root.Left)   // 左
	rightValue := sumOfLeftLeaves(root.Right) // 右

	res := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil { // 中
		res = root.Left.Val
	}

	fmt.Println(res, leftValue, rightValue)
	sum := res + leftValue + rightValue
	return sum
}

// 以上代码精简之后如下：
func sumOfLeftLeaves1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res = root.Left.Val
	}
	return res + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	res := sumOfLeftLeaves(root)
	fmt.Println(res)

}
