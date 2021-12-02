# // https://leetcode-cn.com/problems/design-linked-list/
# // 707. 设计链表
# // 设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。
#
# // 在链表类中实现这些功能：
#
# // get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
# // addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
# // addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
# // addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
# // deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
#
# // 示例：
#
# // MyLinkedList linkedList = new MyLinkedList();
# // linkedList.addAtHead(1);
# // linkedList.addAtTail(3);
# // linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
# // linkedList.get(1);            //返回2
# // linkedList.deleteAtIndex(1);  //现在链表是1-> 3
# // linkedList.get(1);            //返回3


class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
end

class MyLinkedList
    attr_accessor :size, :dummy_head
    def initialize()
        @size = 0
        @dummy_head = ListNode.new
    end

=begin
    :type index: Integer
    :rtype: Integer
=end
    def get(index)
        if index < 0 || index > self.size - 1
            return -1
        end

        cur_node = self.dummy_head.next
        while index > 0
            cur_node = cur_node.next
            index -= 1
        end

        cur_node.val
    end

=begin
    :type val: Integer
    :rtype: Void
=end
    def add_at_head(val)
        new_node = ListNode.new(val)
        new_node.next = self.dummy_head.next
        self.dummy_head.next = new_node
        self.size += 1
    end

=begin
    :type val: Integer
    :rtype: Void
=end
    def add_at_tail(val)
        cur_node = self.dummy_head
        while cur_node.next != nil
            cur_node = cur_node.next
        end
        cur_node.next = ListNode.new(val)
        self.size += 1
    end

=begin
    :type index: Integer
    :type val: Integer
    :rtype: Void
=end
    def add_at_index(index, val)
        if index < 0 || index > self.size
            return
        end

        cur_node = self.dummy_head
        while index > 0
            cur_node = cur_node.next
            index -= 1
        end
        new_node = ListNode.new(val)
        new_node.next = cur_node.next
        cur_node.next = new_node
        self.size += 1
    end

=begin
    :type index: Integer
    :rtype: Void
=end
    def delete_at_index(index)
        if index < 0 || index > self.size - 1
            return
        end

        cur_node = self.dummy_head
        while index > 0
            cur_node = cur_node.next
            index -= 1
        end
        cur_node.next = cur_node.next.next
        self.size -= 1
    end

    def print_list
        cur_node = self.dummy_head
        str = ""
        while cur_node.next != nil
            str += "#{cur_node.next.val} -> "
            cur_node = cur_node.next
        end
        p str
    end

end

linked_list = MyLinkedList.new

linked_list.add_at_head(1)      #链表变为1->
linked_list.print_list()

linked_list.add_at_tail(3)      #链表变为1-> 3
linked_list.print_list()

linked_list.add_at_index(1,2)   #链表变为1-> 2-> 3
linked_list.print_list()

val = linked_list.get(1)        #返回2
p val

linked_list.delete_at_index(1)  #现在链表是1-> 3
linked_list.print_list()

val = linked_list.get(1)        #返回3
p val