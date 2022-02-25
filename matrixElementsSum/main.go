package main

import (
	"fmt"
)

func solution(matrix [][]int) int {
	count := 0
	// fmt.Println(matrix[2][0]) // = 2 
	for i, v := range matrix {
		for j, r := range v {
			// fmt.Println(r)
			if r != 0 {
				if i == 0 {
					count += r
				} else if matrix[i-1][j] != 0 {
					count += r
				} else {
					v[j] = 0
				}
			}
		}
	}
	return count
}

func main() {
	matrix := [][]int{
		{0, 1, 1, 2},
		{0, 5, 0, 0},
		{2, 0, 3, 3},
	}
	matrix2 := [][]int{
		{1, 1, 1, 0},
		{0, 5, 0, 1},
		{2, 1, 3, 10},
	}
	fmt.Println(solution(matrix))
	fmt.Println(solution(matrix2))
}