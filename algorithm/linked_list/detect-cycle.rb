# //142.环形链表II
# //力扣题目链接（https://leetcode-cn.com/problems/linked-list-cycle-ii/）
# //给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
# //
# //如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
# //
# //不允许修改链表。
#
# //示例 1：
# //输入：head = [3,2,0,-4], pos = 1
# //输出：返回索引为 1 的链表节点
# //解释：链表中有一个环，其尾部连接到第二个节点。
# //
# //示例2：
# //输入：head = [1,2], pos = 0
# //输出：返回索引为 0 的链表节点
# //解释：链表中有一个环，其尾部连接到第一个节点。
#
# //示例 3：
# //输入：head = [1], pos = -1
# //输出：返回 null
# //解释：链表中没有环。
#
# //提示：
# //链表中节点的数目范围在范围 [0, 104] 内
# //-105 <= Node.val <= 105
# //pos 的值为 -1 或者链表中的一个有效索引
# //
# //进阶：你是否可以使用 O(1) 空间解决此题?

require_relative './list_node.rb'

# 使用hash
def detect_cycle(head)
	h = {}
	while head != nil
		return head if h[head]
		h[head] = true
		head = head.next
	end

	nil
end

def detect_cycle1(head)
	slow, fast = head, head
	while fast != nil && fast.next != nil
		slow = slow.next
		fast = fast.next.next
		if fast == slow
			while slow != head
				slow = slow.next
				head = head.next
			end
			return head
		end
	end
	
	nil
end

head = ListNode.new(3)
head.next = ListNode.new(2)
head.next.next = ListNode.new(0)
head.next.next.next = ListNode.new(-4)
head.next.next.next.next = head.next

node = detect_cycle(head)
node1 = detect_cycle1(head)
p(node.val)
p(node1.val)
