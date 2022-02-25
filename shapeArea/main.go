package main

import (
	"fmt"
)

func solution(n int) int {
	if n == 1 {
		return 1
	}
	result := ((n - 1) * 4) + solution(n - 1)

	return result
}

func main() {
	// n = 1, 1
	// n = 2, 5
	// n = 3, 13
	// n = 4, 25
	// n = 5, 41
	fmt.Println(solution(2))
	fmt.Println(solution(3))
	fmt.Println(solution(4))
	fmt.Println(solution(5))
	fmt.Println(solution(6))
	fmt.Println(solution(7))
}