package binary_tree

// 二叉树的遍历

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrderTraversal(root.Left)
	preOrderTraversal(root.Right)
}

func midOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	midOrderTraversal(root.Left)
	fmt.Println(root.Val)
	midOrderTraversal(root.Right)
}

func postOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postOrderTraversal(root.Left)
	postOrderTraversal(root.Right)
	fmt.Println(root.Val)
}

func layerOrderTraversal(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil { //防止为空
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) //出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val) //将值加入本层切片中
		}
		res = append(res, tmpArr) //放入结果集
		tmpArr = []int{}          //清空层的数据
	}
	return res
}

func TestTraversal(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 9}
	root.Left.Right.Right = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 8}
	preOrderTraversal(root)
	midOrderTraversal(root)
	postOrderTraversal(root)

	res := layerOrderTraversal(root)
	expected := [][]int{
		{5},
		{4, 6},
		{1, 2, 7, 8},
		{9, 10},
	}
	assert.Equal(t, expected, res)
}
