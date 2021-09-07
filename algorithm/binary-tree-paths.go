package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths1(root *TreeNode) []string {
	res := make([]string, 0)
	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val)
			res = append(res, v)
			return
		}
		s = s + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			travel(node.Left, s)
		}
		if node.Right != nil {
			travel(node.Right, s)
		}
	}
	travel(root, "")
	return res
}

// var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths := []string{}
	// slice其实为runtime.slice结构体（array unsafe.Pointer, len int, cap int）
	// golang因为参数值传递的原因，需要传递指针类型
	constructPaths(root, "", &paths)
	return paths
}

func constructPaths(root *TreeNode, path string, paths *[]string) {
	if root == nil {
		return
	}

	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path)
	} else {
		path += "->"
		constructPaths(root.Left, path, paths)
		constructPaths(root.Right, path, paths)
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	res := binaryTreePaths(root)
	fmt.Println(res)
}
