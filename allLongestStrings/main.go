package main

import (
	"fmt"
)

func solution(inputArray []string) []string {
	maxLen := 0
	outputArray := []string{}
	for _, v := range inputArray {
		if len(v) == maxLen {
			outputArray = append(outputArray, v)
		}
		if len(v) > maxLen {
			maxLen = len(v)
			outputArray = nil
			outputArray = append(outputArray, v)
		}
	}
	return outputArray
}

func main() {
	inputArray := []string{"aba", "aa", "ad", "vcd", "aba"}
	fmt.Println(solution(inputArray))
}