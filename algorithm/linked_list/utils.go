package linked_list

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func listToArr(head *ListNode) []int {
	s := make([]int, 0)
	cur := head
	for cur != nil {
		s = append(s, cur.Val)
		cur = cur.Next
	}
	return s
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
