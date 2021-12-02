# //203.移除链表元素
# //力扣题目链接(https://leetcode-cn.com/problems/remove-linked-list-elements/)
# //
# //题意：删除链表中等于给定值 val 的所有节点。
# //
# //示例 1：
# //输入：head = [1,2,6,3,4,5,6], val = 6
# //输出：[1,2,3,4,5]
# //
# //示例 2：
# //输入：head = [], val = 1
# //输出：[]
# //
# //示例 3：
# //输入：head = [7,7,7,7], val = 7
# //输出：[]

# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
end

# @param {ListNode} head
# @param {Integer} val
# @return {ListNode}
# 输入：head = [1,2,6,3,4,5,6], val = 6
def remove_elements(head, val)
    if head == nil
        return nil
    end

    dummy_head = ListNode.new(0, head)
    cur_node = dummy_head

    while cur_node != nil && cur_node.next != nil
        if cur_node.next.val == val
            cur_node.next = cur_node.next.next
        else
            cur_node = cur_node.next
        end
    end
    dummy_head.next
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

arr = [1,2,6,3,4,5,6]
head = arr_to_list(arr)
print_list(head)

new_head = remove_elements(head, 6)
print_list(new_head)
