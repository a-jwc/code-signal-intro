package main

import (
	"fmt"
)

// y, x
func solution(image [][]int) [][]int {
	blurBox := [3][3]int{}
	imageHeight, imageWidth := len(image), len(image[0])
	if imageHeight < 3 || imageWidth < 3 {
		return [][]int{}
	}
	blurHeight, blurWidth := imageHeight-3, imageWidth-3
	retArr := make([][]int, blurHeight + 1)
	for i := range retArr {
		retArr[i] = make([]int, blurWidth + 1)
	}

	sum := 0
	for i := 0; i <= blurHeight; i++ {
		for j := 0; j <= blurWidth; j++ {

			for x := range blurBox {
				for y := range blurBox[x] {
					sum += image[x+i][y+j]
				}
			}
			retArr[i][j] = sum / 9
			sum = 0
		}
	}

	return retArr
}

func main() {
	image := [][]int{
		{36,0,18,9},
		{27,54,9,0},
		{81,63,72,45},
	}
	fmt.Println(solution(image))
}
