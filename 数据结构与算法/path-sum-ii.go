package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result = [][]int{}
var path = []int{}

// 递归函数不需要返回值，因为我们要遍历整个树
func traversal(cur *TreeNode, count int) {
	if cur.Left == nil && cur.Right == nil && count == 0 { // 遇到了叶子节点且找到了和为sum的路径
		result = append(result, path)
		return
	}

	if cur.Left == nil && cur.Right == nil { // 遇到叶子节点而没有找到合适的边，直接返回
		return
	}

	if cur.Left != nil { // 左 （空节点不遍历）
		path = append(path, cur.Left.Val)
		count -= cur.Left.Val
		traversal(cur.Left, count) // 递归
		count += cur.Left.Val      // 回溯
		path = path[:len(path)-1]  // 回溯
	}
	if cur.Right != nil { // 右 （空节点不遍历）
		path = append(path, cur.Right.Val)
		count -= cur.Right.Val
		traversal(cur.Right, count) // 递归
		count += cur.Right.Val      // 回溯
		path = path[:len(path)-1]   // 回溯
	}
	return
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return result
	}
	path = append(path, root.Val) // 把根节点放进路径
	traversal(root, sum-root.Val)
	return result
}

func main() {
	// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}

	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 14}
	root.Right.Right.Right = &TreeNode{Val: 1}

	res := pathSum(root, 22)
	fmt.Println(res)
}
