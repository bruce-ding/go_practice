package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针法
// dummy -> 1 -> 2 -> 3 -> 4 -> 5
// 假如要移除倒数第2个节点
// 一、fast先走n+1步
// 二、fast和slow同时移动，直到fast指向末尾
// 三、删除slow指向的下一个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	dummy.Next = head
	slow := dummy
	fast := dummy
	// fast先走n步
	for n > 0 && fast != nil {
		fast = fast.Next
		n--
	}
	fmt.Println(fast.Val)
	// fast再提前走一步，因为需要让slow指向删除节点的上一个节点
	fast = fast.Next
	fmt.Println(fast.Val)
	fmt.Println("-----------")
	// fast和slow同时移动，直到fast指向末尾。这个时候slow.Next的位置正好是倒数第n个节点
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		if fast != nil {
			fmt.Println("fast:", fast.Val)
		} else {
			fmt.Println("fast:", nil)
		}
		if slow != nil {
			fmt.Println("slow:", slow.Val)
		} else {
			fmt.Println("slow:", nil)
		}
	}

	slow.Next = slow.Next.Next
	return dummy.Next
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
	newHead := removeNthFromEnd(head, 2)
	printList(newHead)

}
