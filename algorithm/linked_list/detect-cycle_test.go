package linked_list

import (
	"fmt"
	"testing"
)

//142.环形链表II
//力扣题目链接（https://leetcode-cn.com/problems/linked-list-cycle-ii/）
//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
//不允许修改链表。

//示例 1：
//输入：head = [3,2,0,-4], pos = 1
//输出：返回索引为 1 的链表节点
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//示例2：
//输入：head = [1,2], pos = 0
//输出：返回索引为 0 的链表节点
//解释：链表中有一个环，其尾部连接到第一个节点。

//示例 3：
//输入：head = [1], pos = -1
//输出：返回 null
//解释：链表中没有环。

//提示：
//链表中节点的数目范围在范围 [0, 104] 内
//-105 <= Node.val <= 105
//pos 的值为 -1 或者链表中的一个有效索引
//
//进阶：你是否可以使用 O(1) 空间解决此题?

// 使用map
func detectCycle(head *ListNode) *ListNode {
	m := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}

func TestDetectCycle(t *testing.T) {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next

	node := detectCycle(head)
	node1 := detectCycle1(head)
	fmt.Println(node.Val)
	fmt.Println(node1.Val)

	if node != head.Next {
		t.Errorf("expected %v, got %v", head.Next, node)
	}
	if node1 != head.Next {
		t.Errorf("expected %v, got %v", head.Next, node1)
	}
}
