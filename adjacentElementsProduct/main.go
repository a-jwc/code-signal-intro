package main

import (
	"fmt"
	"math"
)

func solution(inputArray []int) int {
	i := 0
	j := i + 1
	max := int(math.Inf(-1))
	for j < len(inputArray) {
		curr := inputArray[i] * inputArray[j]
		if curr > max {
			max = curr
		} 
		i++
		j++
	}

	return max
}

func main() {
	// smolArr := []int{1, 2, 3}
	arr := []int{3, 6, -2, -5, 7, 3};
	fmt.Println(solution(arr))
}