package hard

// ListNode is for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func binSearch(l []*ListNode, node *ListNode) int {
	if node.Val <= l[0].Val {
		return 0
	}
	if node.Val >= l[len(l)-1].Val {
		return len(l)
	}

	low := 0
	high := len(l) - 1

	for low <= high {
		mid := (low + high) / 2
		if l[mid].Val == node.Val {
			return mid
		}
		if l[mid].Val > node.Val {
			high = mid - 1

		} else {
			low = mid + 1

		}
	}
	return low
}

func mergeKLists(lists []*ListNode) *ListNode {
	var sorted []*ListNode
	for _, l := range lists {
		if l == nil {
			continue
		}
		for l != nil {
			if sorted == nil {
				for ; l != nil; l = l.Next {
					sorted = append(sorted, l)
				}
				break
			}
			sorted = append(sorted, l)
			prevLen := len(sorted) - 1
			pos := binSearch(sorted[:prevLen], l)
			if pos < prevLen {
				copy(sorted[pos+1:], sorted[pos:prevLen])
				tmp := l
				l = l.Next
				sorted[pos] = tmp
				sorted[pos].Next = sorted[pos+1]
				if pos > 0 {
					sorted[pos-1].Next = sorted[pos]
				}

			} else {
				sorted[prevLen-1].Next = sorted[prevLen]
				for l = l.Next; l != nil; l = l.Next {
					sorted = append(sorted, l)
				}
				break
			}
		}
	}

	if sorted == nil {
		return nil
	}
	return sorted[0]
}
