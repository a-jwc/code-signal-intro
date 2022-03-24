package main

import (
	"fmt"
	"math"
)

func solution(inputArray []int, k int) int {
	max := math.Inf(-1)
	sum := 0.0
	for i := 0; i < len(inputArray) - k + 1; i++ {
		for j := i; j < i + k; j++ {
			sum += float64(inputArray[j])
		}
		if sum > max {
			max = sum
		}
		sum = 0.0
	}
	return int(max)
}

func main() {
	inputArray := []int{2, 3, 5, 1, 6}
	fmt.Println(solution(inputArray, 2))
}
