package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	//先找到根节点（后续遍历的最后一个就是根节点）
	nodeValue := postorder[len(postorder)-1]
	//从中序遍历中找到一分为二的点，左边为左子树，右边为右子树
	mid := findRootIndex(inorder, nodeValue)
	//构造root
	root := &TreeNode{
		Val:   nodeValue,
		Left:  buildTree(inorder[:mid], postorder[:mid]), //将后续遍历一分为二，左边为左子树，右边为右子树
		Right: buildTree(inorder[mid+1:], postorder[mid:len(postorder)-1])}
	return root
}

func findRootIndex(inorder []int, target int) (index int) {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}

// 中序遍历 inorder = [9,3,15,20,7] 后序遍历 postorder = [9,15,7,20,3] 返回如下的二叉树：
func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
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
