package main

// 206.反转链表
// 力扣题目链接(https://leetcode-cn.com/problems/reverse-linked-list/)

// 题意：反转一个单链表。

// 示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针法
// pre -> 1 -> 2 -> 3 -> 4 -> 5
// pre   cur
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// 递归法
func reverseList1(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	tmp := cur.Next
	cur.Next = pre
	return reverse(cur, tmp) // 通过递归实现双指针法中的更新操作
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
	arr := []int{1, 2, 3, 4, 5}
	head := arrToList(arr)
	printList(head)
	newHead := reverseList(head)
	printList(newHead)
	newHead = reverseList1(newHead)
	printList(newHead)
}
