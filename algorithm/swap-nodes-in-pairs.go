package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 -> 2 -> 3 -> 4
// dummy -> 1 -> 2 -> 3 -> 4

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next // dummy.next = 2

		next := head.Next.Next // next = 3
		head.Next.Next = head  // 2.next = head 1
		head.Next = next       //  1.next = 3
		pre = head             // pre = 1
		head = next            // head = 3
	}
	return dummy.Next
}

// 递归版本
func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

func swapPairsByStack(head *ListNode) *ListNode {
	//解法三：利用栈每次取出两个进行交换,时间复杂度O(n),空间复杂度O(1)
	if head == nil || head.Next == nil {
		return head
	}
	p := &ListNode{}
	cur := head
	//让head指向p，返回的时候返回head.next就好了。
	head = p
	var stack []*ListNode
	for cur != nil && cur.Next != nil {
		stack = append(stack, cur)
		stack = append(stack, cur.Next)
		cur = cur.Next.Next

		e := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		p.Next = e
		p = p.Next

		e = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		p.Next = e
		p = p.Next
	}
	//不管cur是否为空都可以让p.next指向cur。
	//把这句放到循环里面反而内存消耗变大了。所以还是拿出来了。
	p.Next = cur
	return head.Next
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
	newHead := swapPairsByStack(head)
	printList(newHead)
}
