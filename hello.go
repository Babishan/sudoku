package main

import (
	"fmt"
)

func printArray(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(a[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func getK(i, j int) int {
	return ((i / 3) * 3) + (j / 3)
}

func findBest(a [][]int) [][]int {
	var size = len(a)
	var colval = make([][]int, size)
	var rowval = make([][]int, size)
	var boxval = make([][]int, size)
	var count int = 0
	for i := 0; i < size; i++ {
		colval[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		rowval[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		boxval[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var k int = getK(i, j)
			if a[i][j] > 0 {
				rowval[i][a[i][j]-1] = 0
				colval[j][a[i][j]-1] = 0
				boxval[k][a[i][j]-1] = 0
				count = count + 1
			}

		}
	}
	if count == size*size {
		return a
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var k int = getK(i, j)
			if a[i][j] > 0 {
				fmt.Print(a[i][j])
				fmt.Print("\t")
				continue
			} else {
				var possibleValues []int
				for n := 0; n < size; n++ {
					if rowval[i][n] > 0 && colval[j][n] > 0 && boxval[k][n] > 0 {
						possibleValues = append(possibleValues, rowval[i][n])
					}
				}
				fmt.Print(possibleValues)
				if len(possibleValues) == 1 {
					a[i][j] = possibleValues[0]
				}
			}
			fmt.Print("\t")
		}
		fmt.Println()

	}
	return findBest(a)
}
func main() {
	var a = [][]int{
		{6, 0, 0, 0, 0, 0, 7, 0, 4},
		{0, 0, 5, 9, 0, 1, 0, 0, 0},
		{4, 0, 0, 7, 0, 6, 0, 9, 0},
		{0, 7, 6, 0, 9, 0, 5, 1, 0},
		{0, 0, 0, 8, 0, 5, 0, 0, 0},
		{0, 5, 3, 0, 6, 0, 4, 7, 0},
		{0, 8, 0, 3, 0, 9, 0, 0, 1},
		{0, 0, 0, 5, 0, 4, 3, 0, 0},
		{3, 0, 4, 0, 0, 0, 0, 0, 2},
	}
	printArray(findBest(a))
}
