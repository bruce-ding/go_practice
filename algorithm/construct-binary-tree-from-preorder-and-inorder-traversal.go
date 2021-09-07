package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	mid := findRootIndex(preorder[0], inorder)
	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:mid+1], inorder[:mid]),
		Right: buildTree(preorder[mid+1:], inorder[mid+1:])}
	return root
}
func findRootIndex(target int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}

// 前序遍历 preorder = [3,9,20,15,7] 中序遍历 inorder = [9,3,15,20,7] 返回如下的二叉树：
func main() {
	inorder := []int{3, 9, 20, 15, 7}
	postorder := []int{9, 3, 15, 20, 7}
	root := buildTree(inorder, postorder)
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
