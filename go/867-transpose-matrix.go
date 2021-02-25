package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	result := make([][]int, 0)
	wSize, hSize := len(matrix[0]), len(matrix)
	for i := 0; i < wSize; i++ {
		temp := []int{}
		for j := 0; j < hSize; j++ {
			temp = append(temp, matrix[j][i])
		}
		result = append(result, temp)
	}
	return result
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(transpose(matrix))
	matrix = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(transpose(matrix))
}
