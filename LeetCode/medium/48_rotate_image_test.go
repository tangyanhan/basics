package medium

import (
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		expect [][]int
	}{
		{
			matrix: [][]int{},
			expect: [][]int{},
		},
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expect: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expect: [][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
	}

	for _, test := range testCases {
		rotate(test.matrix)
		for i, row := range test.expect {
			for j, value := range row {
				if value != test.matrix[i][j] {
					t.Fatalf("Expected matrix: %v Got: %v", test.expect, test.matrix)
				}
			}
		}
	}
}
