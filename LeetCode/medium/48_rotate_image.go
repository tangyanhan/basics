package medium

func rotate(matrix [][]int) {
	n := len(matrix)
	// 1. left-top to right-bottom transition
	for rowIdx := 1; rowIdx < n; rowIdx++ {
		for colIdx := 0; colIdx < rowIdx; colIdx++ {
			matrix[rowIdx][colIdx], matrix[colIdx][rowIdx] = matrix[colIdx][rowIdx], matrix[rowIdx][colIdx]
		}
	}

	// 2. left to right transition
	var colMax int
	colMax = n / 2
	for rowIdx := 0; rowIdx < n; rowIdx++ {
		for colIdx := 0; colIdx < colMax; colIdx++ {
			chRowIdx := n - colIdx - 1
			matrix[rowIdx][colIdx], matrix[rowIdx][chRowIdx] = matrix[rowIdx][chRowIdx], matrix[rowIdx][colIdx]
		}
	}
}
