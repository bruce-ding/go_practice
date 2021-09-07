package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func compare(left, right *TreeNode) bool {
	// 首先排除空节点的情况
	if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right == nil {
		return true
	} else if left.Val != right.Val { // 排除了空节点，再排除数值不相同的情况
		return false
	}

	// 此时就是：左右节点都不为空，且数值相同的情况
	// 此时才做递归，做下一层的判断
	outside := compare(left.Left, right.Right) // 左子树：左、 右子树：右
	inside := compare(left.Right, right.Left)  // 左子树：右、 右子树：左
	isSame := outside && inside                // 左子树：中、 右子树：中 （逻辑处理）
	return isSame

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

// 迭代法 使用队列
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := list.New()
	queue.PushBack(root.Left)  // 将左子树头结点加入队列
	queue.PushBack(root.Right) // 将右子树头结点加入队列

	for queue.Len() > 0 { // 接下来就要判断这这两个树是否相互翻转
		leftNode := queue.Remove(queue.Front()).(*TreeNode)
		rightNode := queue.Remove(queue.Front()).(*TreeNode)

		if leftNode != nil && rightNode != nil { // 左节点为空、右节点为空，此时说明是对称的
			continue
		}

		// 左右一个节点不为空，或者都不为空但数值不相同，返回false
		if leftNode != nil || rightNode != nil || (leftNode.Val != rightNode.Val) {
			return false
		}

		queue.PushBack(leftNode.Left)   // 加入左节点左孩子
		queue.PushBack(rightNode.Right) // 加入右节点右孩子
		queue.PushBack(leftNode.Right)  // 加入左节点右孩子
		queue.PushBack(rightNode.Left)  // 加入右节点左孩子
	}
	return true
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}
	res := isSymmetric1(root)
	fmt.Println(res)
}
