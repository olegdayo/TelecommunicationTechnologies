package main

import "fmt"

const dashNumber int = 100

func printDelimiter() {
	for i := 0; i < dashNumber; i++ {
		fmt.Print("-")
	}
	fmt.Println("")
}

func output(m *Matrix) {
	fmt.Printf("Number of rows: %d\n", m.Rows)
	fmt.Printf("Number of columns: %d\n", m.Columns)
	fmt.Println("Matrix:")
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			fmt.Printf("%d ", m.Matr[i][j])
		}
		fmt.Println("")
	}
}

func main() {
	g := NewMatrix([][]int{
		{1, 0, 1, 1, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 1, 0, 1, 1, 0},
		{0, 0, 0, 1, 0, 1, 1},
	})
	h := NewMatrix([][]int{
		{0, 0, 1, 1, 1, 0, 1},
		{0, 1, 1, 1, 0, 1, 0},
		{1, 1, 1, 0, 1, 0, 0},
	})

	g.getCodeInfo(G)
	printDelimiter()
	h.getCodeInfo(H)
	printDelimiter()
	output(g.multiplyTo(h.T()))
	printDelimiter()
	output(g.multiplyTo(h.T()).toBinary())
}
