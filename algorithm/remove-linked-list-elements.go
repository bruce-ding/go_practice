package main

import "fmt"

// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

// 示例 1：

// 输入：head = [1,2,6,3,4,5,6], val = 6
// 输出：[1,2,3,4,5]
// 示例 2：

// 输入：head = [], val = 1
// 输出：[]
// 示例 3：

// 输入：head = [7,7,7,7], val = 7
// 输出：[]

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
}

func main() {
	// head = [1,2,6,3,4,5,6], val = 6
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	newHead := removeElements(head, 6)
	printList(newHead)

}
