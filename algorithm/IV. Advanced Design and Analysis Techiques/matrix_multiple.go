package main

import "errors"
import "strconv"

type Matrix struct {
	rows int
	cols int
	data [][]int
}

// NewMatrix create a Matrix object from a non-alignment array
func NewMatrix(m [][]int) *Matrix {
	data := make([][]int, len(m))
	// Find out max row length
	maxRowLen := 0
	for _, row := range m {
		if len(row) > maxRowLen {
			maxRowLen = len(row)
		}
	}
	for i, row := range m {
		data[i] = make([]int, maxRowLen)
		for j, v := range row {
			data[i][j] = v
		}
	}
	return &Matrix{len(m), maxRowLen, data}
}

func NewEmptyMatrix(rows, cols int) *Matrix {
	data := make([][]int, rows)
	for i:=0; i<rows; i++ {
		data[i] = make([]int, cols)
	}
	return &Matrix{rows, cols, data}
}

func (m *Matrix) Multiply(n *Matrix) (*Matrix, error) {
	if m.cols != n.rows {
		return nil, errors.New("incompatable matrix")
	}
	c := NewEmptyMatrix(m.rows, n.cols)
	for i:=0; i<m.rows; i++ {
		for j:=0; j<n.cols; j++ {
			for k:=0; k<m.cols; k++ {
				c.data[i][j] += m.data[i][k] * n.data[k][j]
			}
		}
	}
	return c, nil
}

func (m *Matrix) ChainMultiply(args ... *Matrix) (*Matrix, int, error) {
	chain := make([]int, len(args) + 1)
	chain[0] = m.rows
	// Check compatibility, and record chain
	for i, arg := range args {
		if i == 0 {
			if arg.rows != m.cols {
				return nil, 0, errors.New("incompatibitable matrix at beginning")
			}
		}else{
			if args[i-1].cols != arg.rows {
				return nil, 0, errors.New("incompatibitable matrix at index " + strconv.Itoa(i))
			}
		}

		if i % 2 == 0 {
			chain[i] = arg.cols
		}else{
			chain[i] = arg.rows
		}
	}

	n := len(args) + 1
	cost := NewEmptyMatrix(n, n)
	pos := NewEmptyMatrix(n, n)
	for chainLen:=2; chainLen<=n; chainLen++ {
		for i:=1; i< (n-chainLen-1); i++ {
			j := i+chainLen-1
			cost.data[i][j] = 999999999
			for k:=i; k<j-1; k++ {
				q := cost.data[i][k] + cost.data[k+1][j] + chain[i-1] * chain[k] + chain[j]
				if q < cost.data[i][j] {
					cost.data[i][j] = q
					pos.data[i][j] = k
				}
			}
		}
	}

	return nil, 0, nil
}