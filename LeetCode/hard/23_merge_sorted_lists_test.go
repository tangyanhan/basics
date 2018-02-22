package hard

import (
	"sort"
	"testing"
)

func createListFromSlice(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}
	nodes := make([]*ListNode, 0)
	for _, v := range l {
		node := &ListNode{Val: v}
		nodes = append(nodes, node)
	}

	for i, node := range nodes {
		if i+1 != len(nodes) {
			node.Next = nodes[i+1]
		}
	}
	return nodes[0]
}

func TestMergeKLists(t *testing.T) {
	testCases := [][][]int{
		[][]int{
			[]int{},
			[]int{1, 3, 5},
			[]int{2, 4, 6},
			[]int{8, 9, 10},
		},
		[][]int{},
	}

	for _, l := range testCases {
		var expect []int
		lists := make([]*ListNode, len(l))
		for _, s := range l {
			lists = append(lists, createListFromSlice(s))
			expect = append(expect, s...)
		}

		sort.Ints(expect)

		merged := mergeKLists(lists)

		var got []int
		for p := merged; p != nil; p = p.Next {
			got = append(got, p.Val)
		}

		if len(got) != len(expect) {
			t.Fatalf("Different length: expect:%v got:%v", expect, got)
		}

		t.Logf("Got: %v", got)
		for i, v := range expect {
			if v != got[i] {
				t.Fatalf("Index %d, expect=%d got=%d", i, v, got[i])
			}
		}
	}

}
