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

require_relative './list_node.rb'

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

arr = [1, 2, 3, 4, 5]
head = arr_to_list(arr)
newHead = remove_nth_from_end(head, 2)
print_list(newHead)
