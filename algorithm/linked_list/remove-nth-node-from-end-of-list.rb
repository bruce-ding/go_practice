# // 19.删除链表的倒数第N个节点
# // 力扣题目链接(https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)

# // 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

# // 进阶：你能尝试使用一趟扫描实现吗？

# // 示例1：
# // 输入：head = [1,2,3,4,5], n = 2 输出：[1,2,3,5]

# // 示例 2：
# // 输入：head = [1], n = 1 输出：[]

# // 示例 3：
# // 输入：head = [1,2], n = 1 输出：[1]

# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
end

def remove_nth_from_end(head, n)
    dummy = ListNode.new
    dummy.next = head
    slow = fast = dummy

    while n > 0 && fast != nil
        fast = fast.next
        n -= 1
    end

    fast = fast.next

    while fast != nil
        slow = slow.next
        fast = fast.next
    end
    slow.next = slow.next.next

    dummy.next
end

def print_list(head)
	cur = head
    str = ''
	while cur != nil
		str += "#{cur.val} -> "
		cur = cur.next
    end
    p str
end

def arr_to_list(arr)
    if arr.length == 0
        return nil
    end

    dummy_head = ListNode.new
    cur_node = dummy_head

    arr.each do |val|
        new_node = ListNode.new(val)
        cur_node.next = new_node
        cur_node = new_node
    end

    dummy_head.next
end

arr = [1, 2, 3, 4, 5]
head = arr_to_list(arr)
newHead = remove_nth_from_end(head, 2)
print_list(newHead)
