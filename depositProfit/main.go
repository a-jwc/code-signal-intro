package main

import (
	"fmt"
)

func solution(deposit int, rate int, threshold int) int {
	acc := float64(deposit)
	year := 0
	for {
		acc = acc + ((float64(rate) / 100) * acc)
		year++
		if acc >= float64(threshold) {
			return year
		}
	}
}

func main() {
	fmt.Println(solution(100, 20, 170))
}
