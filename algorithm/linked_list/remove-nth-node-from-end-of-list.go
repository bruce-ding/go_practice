package main

// 19.删除链表的倒数第N个节点
// 力扣题目链接(https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

// 进阶：你能尝试使用一趟扫描实现吗？

// 示例1：
// 输入：head = [1,2,3,4,5], n = 2 输出：[1,2,3,5]

// 示例 2：
// 输入：head = [1], n = 1 输出：[]

// 示例 3：
// 输入：head = [1,2], n = 1 输出：[1]

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

	// fast再提前走一步，因为需要让slow指向删除节点的上一个节点
	fast = fast.Next

	// fast和slow同时移动，直到fast指向末尾。这个时候slow.Next的位置正好是倒数第n个节点
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
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
	newHead := removeNthFromEnd(head, 2)
	printList(newHead)
}
