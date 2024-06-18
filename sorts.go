package main

import (
	"errors"
	"fmt"
	"math"
)

type Matrix struct {
	matrix  [][]float64
	rows    int
	columns int
}

func (v Matrix) getRows() int {
	return v.rows
}

func (v Matrix) getColumns() int {
	return v.columns
}

func (v Matrix) print() {
	for i := 0; i < v.columns; i++ {
		fmt.Printf("|")
		for j := 0; j < v.rows; j++ {
			fmt.Printf("%d|", v.matrix[i][j])
		}
		fmt.Println("")
	}
}

func newMatrix(matrix [][]float64) Matrix {
	return Matrix{matrix, len(matrix), len(matrix[0])}
}

func newMatrixFromInt(matrix [][]int) Matrix {
	var newMatrixVar [][]float64
	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		var newRow []float64
		for j := 0; j < col; j++ {
			if j != col {
				newRow = append(newRow, float64(matrix[i][j]))
			}
		}
		newMatrixVar = append(newMatrixVar, newRow)
	}
	return newMatrix(newMatrixVar)
}

func (v Matrix) cut(row, col int) Matrix {
	var tempMatrix [][]float64
	n := v.rows
	for i := 0; i < n; i++ {
		if i == row {
			continue
		}
		var newRow []float64
		for j := 0; j < n; j++ {
			if j != col {
				newRow = append(newRow, v.matrix[i][j])
			}
		}
		tempMatrix = append(tempMatrix, newRow)
	}
	return newMatrix(tempMatrix)
}

func (v Matrix) determinant() (float64, error) {
	if v.columns != v.rows {
		return 0.0, errors.New("matrix dimensions do not match for determinant")
	}

	if v.rows == 2 {
		return v.matrix[0][0]*v.matrix[1][1] - v.matrix[0][1]*v.matrix[1][0], nil
	}

	var result = 0.0

	for i := 0; i < v.columns; i++ {
		matrixTemp := v.cut(0, i)
		x, _ := matrixTemp.determinant()
		result += math.Pow(-1, float64(i)) * v.matrix[0][i] * x
	}
	return result, nil
}

func (v Matrix) transpose() Matrix {
	transposed := make([][]float64, v.rows)
	for i := range transposed {
		transposed[i] = make([]float64, v.columns)
	}
	for i := 0; i < v.rows; i++ {
		for j := 0; j < v.columns; j++ {
			transposed[j][i] = v.matrix[i][j]
		}
	}
	return newMatrix(transposed)
}

func (v Matrix) cofactor() Matrix {
	cofactors := make([][]float64, v.rows)
	for i := range cofactors {
		cofactors[i] = make([]float64, v.columns)
	}
	for i := 0; i < v.rows; i++ {
		for j := 0; j < v.columns; j++ {
			minor := v.cut(i, j)
			det, _ := minor.determinant()
			cofactors[i][j] = math.Pow(-1, float64(i+j)) * det
		}
	}
	return newMatrix(cofactors)
}

func (v Matrix) inverse() (Matrix, error) {
	det, _ := v.determinant()
	if det == 0 {
		return newMatrix([][]float64{{1}}), errors.New("matrix is singular and cannot be inverted")
	}
	cofactor := v.cofactor()
	ad := cofactor.transpose()
	inverse := make([][]float64, v.rows)
	for i := range inverse {
		inverse[i] = make([]float64, v.columns)
	}
	for i := 0; i < v.rows; i++ {
		for j := 0; j < v.columns; j++ {
			inverse[i][j] = ad.matrix[i][j] / det
		}
	}
	return newMatrix(inverse), nil
}

func main() {
	m := newMatrix([][]float64{{1, 2}, {3, 4}})
	m.print()
	fmt.Println()
	m = m.transpose()
	m.print()
}
