def parent_idx(node_idx):
    if node_idx % 2 == 0:
        return int((node_idx-2)/2)
    return int((node_idx-1)/2)

def left_child_idx(root_idx):
    return 2 * root_idx + 1

def right_child_idx(root_idx):
    return 2 * root_idx + 2

def max_heapify(l, root_idx, heap_len):
    max_idx = root_idx
    l_idx = left_child_idx(root_idx)
    if l_idx < heap_len and l[l_idx] > l[max_idx]:
        max_idx = l_idx
    r_idx = right_child_idx(root_idx)
    if r_idx < heap_len and l[r_idx] > l[max_idx]:
        max_idx = l_idx
    if max_idx != root_idx:
        l[max_idx], l[root_idx] = l[root_idx], l[max_idx]


def build_max_heap(l, heap_len):
    bottom_root_idx = parent_idx(heap_len-1)
    for i in range(bottom_root_idx, -1, -1):
        print 'root-', bottom_root_idx
        max_heapify(l, i, heap_len)

def heap_sort(l):
    build_max_heap(l, len(l))
    for i in range(len(l)-1, 0, -1):
        l[i], l[0] = l[0], l[i]
        print 'before', l
        build_max_heap(l, i)
        print 'after', l

if __name__ == '__main__':
    nums = [1, 3, 2, 0, 7, 9, 10]
    heap_sort(nums)
    print nums




