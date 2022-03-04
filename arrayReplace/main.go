package main

import (
	"fmt"
)

func solution(inputArray []int, elemToReplace int, substitutionElem int) []int {
	for i, v := range inputArray {
		if v == elemToReplace {
			inputArray[i] = substitutionElem
		}
	}
	return inputArray
}

func main() {
	inputArray := []int{1, 2, 1}
	elemToReplace := 1
	substitutionElem := 3
	fmt.Println(solution(inputArray, elemToReplace, substitutionElem))
}