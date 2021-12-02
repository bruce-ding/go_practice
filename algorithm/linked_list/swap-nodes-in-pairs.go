package main

// 24. 两两交换链表中的节点
// 力扣题目链接(https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// origin: 1 -> 2 -> 3 -> 4
// step0: cur = dummy
// step1: cur -> 2
// step2: cur -> 2 -> 1
// step3: cur -> 2 -> 1 -> 3
// step4: cur = 1
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head} // 设置一个虚拟头结点, 将虚拟头结点指向head，这样方面后面做删除操作
	cur := dummyHead

	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next            // 记录临时节点
		tmp1 := cur.Next.Next.Next // 记录临时节点

		cur.Next = cur.Next.Next  // 步骤一
		cur.Next.Next = tmp       // 步骤二
		cur.Next.Next.Next = tmp1 // 步骤三

		cur = cur.Next.Next // cur移动两位，准备下一轮交换
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

	newHead := swapPairs(head)
	printList(newHead)
}
