package main

import "fmt"

type Matrix struct {
	Rows    int
	Columns int
	Matr    [][]int
}

type MatrixType int

const (
	G MatrixType = iota
	H
)

func NewMatrix(matr [][]int) *Matrix {
	m := new(Matrix)
	m.Rows = len(matr)
	m.Columns = len(matr[0])
	m.Matr = matr
	return m
}

func (m1 *Matrix) multiplyTo(m2 *Matrix) *Matrix {
	ans := new(Matrix)
	ans.Rows = m1.Rows
	ans.Columns = m2.Columns
	ans.Matr = make([][]int, ans.Rows)
	for i := 0; i < ans.Rows; i++ {
		ans.Matr[i] = make([]int, ans.Columns)
	}

	for h := 0; h < ans.Rows; h++ {
		for i := 0; i < ans.Columns; i++ {
			ans.Matr[h][i] = 0
			for j := 0; j < m1.Columns; j++ {
				ans.Matr[h][i] += m1.Matr[h][j] * m2.Matr[j][i]
			}
		}
	}

	return ans
}

func (m *Matrix) toPower(n int) *Matrix {
	ans := new(Matrix)
	ans.Rows = m.Rows
	ans.Columns = m.Columns
	ans.Matr = m.Matr
	for i := 0; i < n-1; i++ {
		ans = ans.multiplyTo(m)
	}
	return ans
}

func (m *Matrix) T() *Matrix {
	ans := new(Matrix)
	ans.Rows = m.Columns
	ans.Columns = m.Rows
	ans.Matr = make([][]int, ans.Rows)
	for i := 0; i < ans.Rows; i++ {
		ans.Matr[i] = make([]int, ans.Columns)
	}
	for i := 0; i < ans.Rows; i++ {
		for j := 0; j < ans.Columns; j++ {
			ans.Matr[i][j] = m.Matr[j][i]
		}
	}
	return ans
}

func (m *Matrix) toBinary() *Matrix {
	ans := new(Matrix)
	ans.Rows = m.Rows
	ans.Columns = m.Columns
	ans.Matr = m.Matr
	for i := 0; i < ans.Rows; i++ {
		for j := 0; j < ans.Columns; j++ {
			ans.Matr[i][j] = m.Matr[i][j] % 2
		}
	}
	return ans
}

func (m *Matrix) getCodeInfo(mt MatrixType) {
	switch mt {
	case G:
		fmt.Printf("Resulting message length: %d\n", m.Columns)
		fmt.Printf("Number of verification symbols: %d\n", m.Columns-m.Rows)
		fmt.Printf("Number of informative symbols (length of original message): %d\n", m.Rows)
	case H:
		fmt.Printf("Resulting message length: %d\n", m.Columns)
		fmt.Printf("Number of verification symbols: %d\n", m.Rows)
		fmt.Printf("Number of informative symbols (length of original message): %d\n", m.Columns-m.Rows)
	}
}
