# // 24. 两两交换链表中的节点
# // 力扣题目链接(https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

# // 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

# // 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

require_relative './list_node.rb'

# origin: 1 -> 2 -> 3 -> 4
# step0: cur = dummy -> 1 -> 2 -> 3 -> 4
# step1: cur -> 2
# step2: cur -> 2 -> 1
# step3: cur -> 2 -> 1 -> 3
# step4: cur = 1

def swap_pairs(head)
	dummy = ListNode.new
	dummy.next = head
	cur = dummy
	
	while cur.next != nil && cur.next.next != nil
		tmp = cur.next            # 记录临时节点
		tmp1 = cur.next.next.next # 记录临时节点

		cur.next = cur.next.next  # 步骤一
		cur.next.next = tmp       # 步骤二
		cur.next.next.next = tmp1 # 步骤三

		cur = cur.next.next # cur移动两位，准备下一轮交换
	end

	dummy.next
end

arr = [1, 2, 6, 3, 4, 5, 6]
head = arr_to_list(arr)
print_list(head)

new_head = swap_pairs(head)
print_list(new_head)

