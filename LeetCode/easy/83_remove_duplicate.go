package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListFromSlice(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}

	l := new(ListNode)
	l.Val = s[0]
	for p, j := l, 1; j < len(s); j++ {
		p.Next = &ListNode{s[j], nil}
		p = p.Next
	}
	return l
}

func (l *ListNode) ToSlice() []int {
	p := l
	s := make([]int, 0)
	for p != nil {
		s = append(s, p.Val)
		p = p.Next
	}
	return s
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, q := head, head.Next
	for q != nil {
		for q != nil && p.Val == q.Val {
			q = q.Next
		}
		p.Next = q
		p = q
	}
	return head
}
