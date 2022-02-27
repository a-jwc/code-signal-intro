package main

import (
	"fmt"
)

func solution(inputString string) bool {
	palinMap := make(map[rune]int)
	for _, v := range inputString {
		palinMap[v] += 1
	}
	fmt.Println(palinMap)

	singleCount := 0
	oddCount := 0
	for _, v := range palinMap {
		if v == 1 {
			singleCount += 1 
		}
		if v % 2 == 1 {
			oddCount += 1
		}
		if singleCount > 1 || oddCount > 1 {
			return false
		}
	}

	return true
}

func main() {
	inputString := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbccccaaaaaaaaaaaaa"
	fmt.Println(solution(inputString))
}
