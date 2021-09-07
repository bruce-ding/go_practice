package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA := 0
	lenB := 0
	for curA != nil { // 求链表A的长度
		lenA++
		curA = curA.Next
	}
	for curB != nil { // 求链表B的长度
		lenB++
		curB = curB.Next
	}

	curA = headA
	curB = headB
	// 让curA为最长链表的头，lenA为其长度
	if lenB > lenA {
		lenA, lenB = lenB, lenA
		curA, curB = curB, curA
	}
	// 求长度差
	gap := lenA - lenB
	// 让curA和curB在同一起点上（末尾位置对齐）
	for gap > 0 {
		curA = curA.Next
		gap--
	}
	// 遍历curA 和 curB，遇到相同则直接返回
	for curA != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}

/**
思路

A长度为 a, B长度为b， 假设存在交叉点，此时 A到交叉点距离为 c， 而B到交叉点距离为d
后续交叉后长度是一样的，那么就是 a-c = b-d 交换后得到 a+d = b+c
这里意味着只要分别让A和B额外多走一遍B和A，那么必然会走到交叉，注意这里边缘情况是，大家都走到null依然没交叉，那么正好返回null即可
**/
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}

func main() {
	// listA = [4,1,8,4,5], listB = [5,0,1,8,4,5]
	// [4,1,8,4,5,5,0,1,8,4,5]
	// [5,0,1,8,4,5,4,1,8,4,5]
	a := &ListNode{Val: 8}
	b := &ListNode{Val: 4}
	c := &ListNode{Val: 5}

	headA := &ListNode{Val: 4}
	headA.Next = &ListNode{Val: 1}
	headA.Next.Next = a
	headA.Next.Next.Next = b
	headA.Next.Next.Next.Next = c

	headB := &ListNode{Val: 5}
	headB.Next = &ListNode{Val: 0}
	headB.Next.Next = &ListNode{Val: 1}
	headB.Next.Next.Next = a
	headB.Next.Next.Next.Next = b
	headB.Next.Next.Next.Next.Next = c

	res := getIntersectionNode1(headA, headB)
	fmt.Println(res.Val)
}
