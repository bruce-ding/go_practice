package main

// 链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

//利用BSL的性质（前序遍历有序）
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val { //当前节点的值大于给定的值，则说明满足条件的在左边
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val { //当前节点的值小于各点的值，则说明满足条件的在右边
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	} //当前节点的值在给定值的中间（或者等于），即为最深的祖先
}

//递归会将值层层返回
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	//终止条件
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	} //最后为空或者找到一个值时，就返回这个值
	//后序遍历
	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)
	//处理单层逻辑
	if left != nil && right != nil {
		return root
	} //说明在root节点的两边
	if left == nil { //左边没找到，就说明在右边找到了
		return right
	} else {
		return left
	}
}

func main() {
	// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8 输出: 6 解释: 节点 2 和节点 8 的最近公共祖先是 6。
	root := NewTreeNode(6)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(0)
	root.Left.Right = NewTreeNode(4)
	root.Left.Right.Left = NewTreeNode(3)
	root.Left.Right.Right = NewTreeNode(5)

	root.Right = NewTreeNode(8)
	root.Right.Left = NewTreeNode(7)
	root.Right.Right = NewTreeNode(9)

	res := lowestCommonAncestor(root, root.Left, root.Right.Left)
	fmt.Println(res.Val)
}
