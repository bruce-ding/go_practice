# 206.反转链表
# 力扣题目链接(https://leetcode-cn.com/problems/reverse-linked-list/)

# 题意：反转一个单链表。

# 示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL

require_relative './list_node.rb'

# 双指针法
def reverse_list(head)
    pre = nil
    cur = head
    while cur != nil
        tmp = cur.next
        cur.next = pre

        pre = cur
        cur = tmp
    end

    pre
end

# 递归法
def reverse_list_recursively(head)
    do_reverse(nil, head)
end

def do_reverse(pre, cur)
  return pre if cur.nil?

  tmp = cur.next
  cur.next = pre
  reverse(cur, tmp)	# 通过递归实现双指针法中的更新操作
end

# 示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL

arr = [1,2,3,4,5]
head = arr_to_list(arr)
print_list(head)
new_head = reverse_list(head)
print_list(new_head)
new_head = reverse_list(new_head)
print_list(new_head)