package main

import (
	"fmt"
)

func solution(inputArray []int, k int) []int {
	count := 1
	inputArrayLength := len(inputArray)
	if k == 1 {
		return []int{}
	}
	for i := k*count - 1; i < inputArrayLength; i = k*count - 1 {
		inputArray[i] = -21
		count++
	}
	inputArrayLength = len(inputArray)
	for i := 0; i < inputArrayLength; i++ {
		if inputArray[i] == -21 {
			inputArray = append(inputArray[:i], inputArray[i+1:]...)
			inputArrayLength -= 1
		}
	}
	return inputArray
}

func main() {
	inputArray := []int{1, 2, 1, 2, 1, 2, 1, 2}
	k := 1
	fmt.Println(solution(inputArray, k))
}
