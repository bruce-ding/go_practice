package main

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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// check
	if root == nil {
		return root
	}
	// 相等 直接返回root节点即可
	if root == p || root == q {
		return root
	}
	// Divide
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil {
		if left.Val == 0 {
			fmt.Println("root:", root.Val)
		}
		fmt.Println("left:", left.Val)
	} else {
		fmt.Println("left:", "nil")
	}
	if right != nil {
		fmt.Println("right:", right.Val)
	} else {
		fmt.Println("left:", "nil")
	}

	// Conquer
	// 左右两边都不为空，则根节点为祖先
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

/**
我们可以用哈希表存储所有节点的父节点，然后我们就可以利用节点的父节点信息从 p 结点开始不断往上跳，并记录已经访问过的节点，再从 q 节点开始不断往上跳，如果碰到已经访问过的节点，那么这个节点就是我们要找的最近公共祖先。

算法
从根节点开始遍历整棵二叉树，用哈希表记录每个节点的父节点指针。
从 p 节点开始不断往它的祖先移动，并用数据结构记录已经访问过的祖先节点。
同样，我们再从 q 节点开始不断往它的祖先移动，如果有祖先已经被访问过，即意味着这是 p 和 q 的深度最深的公共祖先，即 LCA 节点。
C++JavaGolang
**/
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}

	return nil
}

func main() {
	// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1 输出: 3
	root := NewTreeNode(3)
	root.Left = NewTreeNode(5)
	root.Left.Left = NewTreeNode(6)
	root.Left.Right = NewTreeNode(2)
	root.Left.Right.Left = NewTreeNode(7)
	root.Left.Right.Right = NewTreeNode(4)

	root.Right = NewTreeNode(1)
	root.Right.Left = NewTreeNode(0)
	root.Right.Right = NewTreeNode(8)

	res := lowestCommonAncestor(root, root.Left, root.Right.Left)
	fmt.Println(res.Val)
}
