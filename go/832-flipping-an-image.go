package main

import (
	"fmt"
)

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		for m, n := 0, len(A[i])-1; m < n; m, n = m+1, n-1 {
			A[i][m], A[i][n] = A[i][n], A[i][m]
		}
		for j := 0; j < len(A[i]); j++ {
			A[i][j] ^= 1
		}
	}
	return A
}

func main() {
	arr := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(flipAndInvertImage(arr))
	arr = [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	fmt.Println(flipAndInvertImage(arr))
}
