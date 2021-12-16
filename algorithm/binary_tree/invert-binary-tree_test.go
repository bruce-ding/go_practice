package binary_tree

/**
226. 翻转二叉树
力扣题目链接(https://leetcode-cn.com/problems/invert-binary-tree/)
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9

输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

**/

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 前序遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 先交换左右孩子
	root.Left, root.Right = root.Right, root.Left // 中
	// 再分别翻转左右孩子
	invertTree(root.Left)  // 左
	invertTree(root.Right) // 右
	return root
}

// 层序遍历
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			node.Left, node.Right = node.Right, node.Left //交换
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root
}

func TestInvertTree(t *testing.T) {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)
	root.Right = NewTreeNode(3)
	root.Right.Left = NewTreeNode(6)
	root.Right.Right = NewTreeNode(7)

	fmt.Println("before invert")
	printTree(root)

	res := invertTree(root)
	fmt.Println("after invert")
	slice := printTree(res)
	expected := [][]int{
		{1}, {3, 2}, {7, 6, 5, 4},
	}
	assert.Equal(t, expected, slice)

	res = invertTree1(root)
	fmt.Println("invert back")
	slice = printTree(res)
	expected = [][]int{
		{1}, {2, 3}, {4, 5, 6, 7},
	}
	assert.Equal(t, expected, slice)
}
