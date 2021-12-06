# //面试题 02.07. 链表相交
# //力扣题目链接
# //给你两个单链表的头节点headA 和 headB，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
# //
# //图示两个链表在节点 c1 开始相交：
# //
# //题目数据 保证 整个链式结构中不存在环。
# //
# //注意，函数返回结果后，链表必须 保持其原始结构 。
# //
# //示例 1：
# //输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
# //输出：Intersected at '8'
# //解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
# //从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
# //在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
# //示例2：
# //输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
# //输出：Intersected at '2'
# //解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
# //从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。
# //在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
# //示例3：
# //输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
# //输出：null
# //解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
# //由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
# //这两个链表不相交，因此返回 null。
# //
# //提示：
# //
# //listA 中节点数目为 m
# //listB 中节点数目为 n
# //0 <= m, n <= 3 * 104
# //1 <= Node.val <= 105
# //0 <= skipA <= m
# //0 <= skipB <= n
# //如果 listA 和 listB 没有交点，intersectVal 为 0
# //如果 listA 和 listB 有交点，intersectVal == listA[skipA + 1] == listB[skipB + 1

# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
end

def get_intersection_node(headA, headB)
	curA = headA
	curB = headB
	lenA, lenB = 0, 0
	# 求A，B的长度
	while curA != nil
		curA = curA.next
		lenA += 1
	end
	while curB != nil
		curB = curB.next
		lenB += 1
	end

	step = 0
	fast, slow = nil, nil
	# 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB
		step = lenA - lenB
		fast, slow = headA, headB
	else
		step = lenB - lenA
		fast, slow = headB, headA
	end
	while step > 0
		fast = fast.next
		step -= 1
	end
	# 遍历两个链表遇到相同则跳出遍历
	while fast != slow
		fast = fast.next
		slow = slow.next
	end

	fast
end

# // listA = [4,1,8,4,5], listB = [5,0,1,8,4,5]
# // [4,1,8,4,5,5,0,1,8,4,5]
# // [5,0,1,8,4,5,4,1,8,4,5]
a = ListNode.new(8)
b = ListNode.new(4)
c = ListNode.new(5)

headA = ListNode.new(4)
headA.next = ListNode.new(1)
headA.next.next = a
headA.next.next.next = b
headA.next.next.next.next = c

headB = ListNode.new(5)
headB.next = ListNode.new(0)
headB.next.next = ListNode.new(1)
headB.next.next.next = a
headB.next.next.next.next = b
headB.next.next.next.next.next = c

res = get_intersection_node(headA, headB)
p res.val  # should be 8
