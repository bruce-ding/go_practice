package linked_list

import (
	"reflect"
	"testing"
)

// 24. 两两交换链表中的节点
// 力扣题目链接(https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

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

// 24. 两两交换链表中的节点
// 力扣题目链接(https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

func TestSwapPairs(t *testing.T) {
	arr := []int{1, 2, 6, 3, 4, 5, 6}
	head := arrToList(arr)
	printList(head)

	newHead := swapPairs(head)
	printList(newHead)

	newArr := listToArr(newHead)

	expected := []int{2,1,3,6,5,4,6}
	if !reflect.DeepEqual(expected, newArr) {
		t.Errorf("expected %v, but got %v", expected, newArr)
	}
}
