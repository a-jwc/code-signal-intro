package main

import (
	"fmt"
)

func solution(matrix [][]bool) [][]int {
	matrixHeight, matrixWidth := len(matrix), len(matrix[0])
	retArr := make([][]int, matrixHeight)
	for i := range retArr {
		retArr[i] = make([]int, matrixWidth)
	}

	for i := 0; i < matrixHeight; i++ {
		for j := 0; j < matrixWidth; j++ {
			switch {
			case i == 0 && j == 0:
				{
					if matrix[i+1][j] {
						retArr[i][j] += 1
					}
					if matrix[i+1][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i][j+1] {
						retArr[i][j] += 1
					}
				}
			case i == 0 && j != 0 && j != matrixWidth-1:
				{
					if matrix[i+1][j] {
						retArr[i][j] += 1
					}
					if matrix[i+1][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i][j-1] {
						retArr[i][j] += 1
					}
					if matrix[i+1][j-1] {
						retArr[i][j] += 1
					}
				}
			case i == 0 && j == matrixWidth-1:
				{
					if matrix[i+1][j] {
						retArr[i][j] += 1
					}
					if matrix[i+1][j-1] {
						retArr[i][j] += 1
					}
					if matrix[i][j-1] {
						retArr[i][j] += 1
					}
				}
			case i != 0 && j == 0 && i != matrixHeight-1:
				{
					if matrix[i+1][j] {
						retArr[i][j] += 1
					}
					if matrix[i+1][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j] {
						retArr[i][j] += 1
					}
				}
			case i != 0 && j == matrixWidth-1 && i != matrixHeight-1:
				{
					if matrix[i+1][j] {
						retArr[i][j] += 1
					}
					if matrix[i+1][j-1] {
						retArr[i][j] += 1
					}
					if matrix[i][j-1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j-1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j] {
						retArr[i][j] += 1
					}
				}
			case i == matrixHeight-1 && j == 0:
				{
					if matrix[i][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j] {
						retArr[i][j] += 1
					}
				}
			case i == matrixHeight-1 && j != matrixWidth-1:
				{
					if matrix[i-1][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i][j-1] {
						retArr[i][j] += 1
					}
					if matrix[i][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j-1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j] {
						retArr[i][j] += 1
					}
				}
			case i == matrixHeight-1 && j == matrixWidth-1:
				{
					if matrix[i][j-1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j-1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j] {
						retArr[i][j] += 1
					}
				}
			default:
				{
					if matrix[i-1][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i][j-1] {
						retArr[i][j] += 1
					}
					if matrix[i][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j-1] {
						retArr[i][j] += 1
					}
					if matrix[i-1][j] {
						retArr[i][j] += 1
					}
					if matrix[i+1][j] {
						retArr[i][j] += 1
					}
					if matrix[i+1][j+1] {
						retArr[i][j] += 1
					}
					if matrix[i+1][j-1] {
						retArr[i][j] += 1
					}
				}
			}
		}
	}

	return retArr
}

func main() {
	matrix := [][]bool{
		{true,false,false,true},
		{false,false,true,false},
		{true,true,false,true},
	}
	fmt.Println(solution(matrix))
}
