package linked_list

// 206.反转链表
// 力扣题目链接(https://leetcode-cn.com/problems/reverse-linked-list/)

// 题意：反转一个单链表。

// 示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL

import (
	"testing"
)

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

func TestReverseList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	head := arrToList(arr)
	printList(head)
	newHead := reverseList(head)
	printList(newHead)
	newHead = reverseList1(newHead)
	printList(newHead)
}
