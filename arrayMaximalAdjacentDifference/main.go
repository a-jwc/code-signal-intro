package main

import (
	"fmt"
	"math"
)

func solution(inputArray []int) int {
	max := -31
	for i := 1; i < len(inputArray); i++ {
		v := inputArray[i]
		prevV := inputArray[i-1]
		dif := int(math.Abs(float64(v - prevV)))
		if dif > max {
			max = dif
		}
	}
	return max
}

func main() {
	inputArray := []int{2, 4, 1, 0}
	fmt.Println(solution(inputArray))
}
