package main

import (
	"fmt"
)

func solution(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
	if maxW >= weight1 + weight2 {
		return value1 + value2
	}
	if weight1 <= maxW && weight2 <= maxW {
		if value1 > value2 {
			return value1
		} else {
			return value2
		}
	} else if weight1 <= maxW {
		return value1
	} else if weight2 <= maxW {
		return value2
	} else {
		return 0
	}
}

func main() {
	fmt.Println(solution(10, 5, 6, 4, 8))
}
