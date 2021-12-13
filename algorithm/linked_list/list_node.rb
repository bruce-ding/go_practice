# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
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