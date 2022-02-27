package main

import (
	"fmt"
)

func solution(inputArray []int) int {
	moves := 0
	// if curr element > prev element -> nop
	// else curr element = update()
	for i := 1; i < len(inputArray); i++ {
		if inputArray[i] <= inputArray[i - 1] {
			value, newMove := update(inputArray[i], inputArray[i - 1])
			inputArray[i] = value
			moves += newMove
		}
	}
	return moves
}

func update(currElement int, prevElement int) (int, int) {
	// misinterpreted algo: moves += int(math.Abs(math.Abs(float64(firstElement+i)) - math.Abs(float64(v))))
	return prevElement + 1, prevElement + 1 - currElement 
}

func main() {
	inputArray := []int{-2, -3, 1}
	// [-2, -3, 1] -> 
	// [-1000, 0, -2, 0] 
	// [2, 1, 10, 1] -> 
	fmt.Println(solution(inputArray))
}
