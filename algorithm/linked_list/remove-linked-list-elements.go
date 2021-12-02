package main

//203.移除链表元素
//力扣题目链接(https://leetcode-cn.com/problems/remove-linked-list-elements/)
//
//题意：删除链表中等于给定值 val 的所有节点。
//
//示例 1：
//输入：head = [1,2,6,3,4,5,6], val = 6
//输出：[1,2,3,4,5]
//
//示例 2：
//输入：head = [], val = 1
//输出：[]
//
//示例 3：
//输入：head = [7,7,7,7], val = 7
//输出：[]

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, target int) *ListNode {
	dummyHead := &ListNode{Next: head}
	curNode := dummyHead

	for curNode != nil && curNode.Next != nil {
		if curNode.Next.Val == target {
			curNode.Next = curNode.Next.Next
		} else {
			curNode = curNode.Next
		}
	}
	return dummyHead.Next
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func arrToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	dummyHead := &ListNode{}
	curNode := dummyHead

	for i := 0; i < len(arr); i++ {
		newNode := &ListNode{Val: arr[i]}
		curNode.Next = newNode
		curNode = newNode
	}

	return dummyHead.Next
}

func main() {
	arr := []int{1, 2, 6, 3, 4, 5, 6}
	head := arrToList(arr)
	printList(head)
	newHead := removeElements(head, 6)
	printList(newHead)

}
