package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// pre -> 1 -> 2 -> 3 -> 4 -> 5
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next   // pre.next = 2
		next := head.Next.Next // next = 3
		head.Next.Next = head  // 2.next = 1
		head.Next = next       // 1.next = 3
		pre = head             // pre = 1
		head = next            // head = 3
	}
	return dummy.Next
}

// 双指针法 1
// 1: pre 指针指向cur.Next,
// 2: cur.Next 指向cur,这一步 会把指向反转
// 3: cur 指向 cur.Next.Next
// 4: 移动cur和pre
func swapNodesInPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head

	for cur != nil && cur.Next != nil {
		pre.Next = cur.Next
		next := cur.Next.Next
		cur.Next.Next = cur
		cur.Next = next
		pre = cur
		cur = next
	}
	return dummy.Next
}

// 双指针法 2
func swapNodesInPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next            // 记录临时节点 1
		tmp1 := cur.Next.Next.Next // 记录临时节点 3

		cur.Next = cur.Next.Next // 步骤一 把dummy指向 2
		cur.Next.Next = tmp      // 步骤二 把2指向1
		tmp.Next = tmp1          // 步骤三 把1指向3

		cur = cur.Next.Next // cur移动两位，准备下一轮交换
	}

	return dummy.Next // 第一次for循环的时候dummy.Next 设为了2
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
	newHead := swapNodesInPairs(head)
	printList(newHead)

}
