package Easy

func visitLeft(t *TreeNode) []int{
	stack := make([]*TreeNode, 0)
	data := make([]int, 0)
	for p := t; p != nil || len(stack) != 0; {
		for p != nil {
			stack = append(stack, p)
			data = append(data, p.Val)
			p = p.Left
			if p == nil {
				data = append(data, -1)
			}
		}
		if len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	return data
}

func visitRight(t *TreeNode) []int{
	stack := make([]*TreeNode, 0)
	data := make([]int, 0)
	for p := t; p != nil || len(stack) != 0; {
		for p != nil {
			stack = append(stack, p)
			data = append(data, p.Val)
			p = p.Right
			if p == nil {
				data = append(data, -1)
			}
		}
		if len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Left
		}
	}
	return data
}

func isSymmetric(root *TreeNode) bool {
	if root == nil { // A null tree is also symmetric
		return true
	}

	if root.Left == nil || root.Right == nil{
		return root.Left == root.Right
	}

	leftData := visitLeft(root.Left)
	rightData := visitRight(root.Right)

	if len(leftData) != len(rightData) {
		return false
	}

	for i, v := range leftData {
		if rightData[i] != v {
			return false
		}
	}
	return true
}
