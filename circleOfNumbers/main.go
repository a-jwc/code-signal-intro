package main

import (
	"fmt"
)

func solution(n int, firstNumber int) int {
	return (firstNumber + n/2) % n
}

func main() {
	fmt.Println(solution(10, 7))
}
