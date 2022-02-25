package main

import (
	"fmt"
	"strconv"
)

func solution(n int) bool {
	nStr := strconv.Itoa(n)

	firstHalf := nStr[0 : len(nStr)/2]
	secondHalf := nStr[len(nStr)/2:]

	firstCount := 0
	secondCount := 0

	for _, v := range firstHalf {
		// * subtracting a rune from any rune '0' - '9' will give an integer 0 - 9
		n := int(v - '0')
		firstCount += n
	}
	for _, v := range secondHalf {
		n := int(v - '0')
		secondCount += n
	}
	return firstCount == secondCount
}

func main() {
	fmt.Println(solution(1121))
}
