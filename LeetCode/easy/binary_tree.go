package easy

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeFromSlice(s []int, idx int) *TreeNode {
	if len(s) == 0 || idx >= len(s) || s[idx] == -1 {
		return nil
	}

	root := &TreeNode{s[idx], nil, nil}
	root.Left = CreateTreeFromSlice(s, 2*idx+1)
	root.Right = CreateTreeFromSlice(s, 2*idx+2)
	return root
}

func PrintTree(t *TreeNode, l int) {
	prefix := strings.Repeat(" ", 4*l) + strings.Repeat("-", 4)
	if t == nil {
		fmt.Println(prefix, "null")
	} else {
		fmt.Println(prefix, t.Val)
	}

	if t != nil {
		PrintTree(t.Left, l+1)
		PrintTree(t.Right, l+1)
	}
}
