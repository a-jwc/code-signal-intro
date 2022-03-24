package main

import (
	"fmt"
)

func solution(upSpeed int, downSpeed int, desiredHeight int) int {
	initHeight := 0
	day := 1
	for initHeight < desiredHeight {
		initHeight += upSpeed
		if initHeight >= desiredHeight {
			return day
		} else {
			initHeight -= downSpeed
		}
		day++
	}
	return day
}

func main() {
	fmt.Println(solution(100, 10, 910))
}
