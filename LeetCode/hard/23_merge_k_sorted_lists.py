# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

    @staticmethod
    def create_list_from_2d_list(l2d):
        lists = list()
        for l in l2d:
            head = None
            last = None
            for v in l:
                node = ListNode(v)
                if last is not None:
                    last.next = node
                else:
                    head = node
                last = node
            lists.append(head)
        return lists


def mergeKLists(lists):
    """
    :type lists: List[ListNode]
    :rtype: ListNode
    """
    l = list()
    for node in lists:
        if node:
            while node:
                l.append(node)
                node = node.next
    head = None
    if l:
        l = sorted(l, key=lambda node: node.val)
        head = l[0]
        for i in range(0, len(l)-1):
            l[i].next = l[i+1]
    return head


if __name__ == '__main__':
    lists = ListNode.create_list_from_2d_list([[-10,-6,-4,-4,0,2,2,2,2],[-5,1,3,4,4,4],[-9],[-9],[-10,-8,-5,-4,-3,-3,-2,4],[-10,-8,-7,-4,-4,0,1,2],[-8,-1,4],[-6,-5,-2,-2,-1,1,3,4]])
    l = mergeKLists(lists)
    sorted_list = list()
    while l is not None:
        sorted_list.append(l.val)
        l = l.next
    print 'Sorted list:', sorted_list