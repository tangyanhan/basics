package easy

var depth int = 0

func visitTree(t *TreeNode, l int) {
	if t.Left == nil && t.Right == nil {
		if l > depth {
			depth = l
		}
		return
	}
	if t.Left != nil {
		visitTree(t.Left, l+1)
	}
	if t.Right != nil {
		visitTree(t.Right, l+1)
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth = 0
	visitTree(root, 1)
	return depth
}
