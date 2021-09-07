package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针法 易于理解
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//递归1
func reverseListRecursive(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return help(head, next)
}

//递归2
func reverseListRecursive1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseListRecursive1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	newHead := reverseListRecursive1(head)
	printList(newHead)
}
