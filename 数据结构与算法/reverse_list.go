package main

import (
	"fmt"
)

type Item struct {
	Key, Val string
}

// byKeys is a comparison function that compares item keys and returns true
// when a is less than b.
func byKeys(a, b interface{}) bool {
	i1, i2 := a.(*Item), b.(*Item)
	return i1.Key < i2.Key
}

// byVals is a comparison function that compares item values and returns true
// when a is less than b.
func byVals(a, b interface{}) bool {
	i1, i2 := a.(*Item), b.(*Item)
	if i1.Val < i2.Val {
		return true
	}
	if i1.Val > i2.Val {
		return false
	}
	// Both vals are equal so we should fall though
	// and let the key comparison take over.
	return byKeys(a, b)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fmt.Printf("head: %d\n", head.Val)
	fmt.Printf("head.Next: %d\n", head.Next.Val)
	newHead := reverseList(head.Next)
	fmt.Printf("newHead: %d\n", newHead.Val)
	fmt.Printf("head.Val: %d\n", head.Val)
	head.Next.Next = head // 5 -> 4    // 2 -> 1
	head.Next = nil       // 4 -> nil  // 1 -> nil
	return newHead
}

// 1 -> 2 -> 3 -> 4 -> 5
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next) // head = 5, head.Next = nil, 此时递归返回
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 双指针法
// pre -> 1 -> 2 -> 3 -> 4 -> 5
// pre   cur

func reverseList1(head *ListNode) *ListNode {
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

func printList(head *ListNode) {
	for head != nil {
		nextNode := head.Next
		if nextNode != nil {
			fmt.Printf("%d -> ", head.Val)
		} else {
			fmt.Printf("%d", head.Val)
		}
		head = nextNode
	}
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	printList(head)
	fmt.Println()
	newHead := reverseList(head)
	printList(newHead)
	newHead = reverseList1(head)
	printList(newHead)
}
