package main

import (
	"fmt"
)

func solution(a []int) []int {
	team1 := 0
	team2 := 0

	for i, v := range a {
		if i%2 == 0 {
			team1 += v
		} else {
			team2 += v
		}
	}

	return []int{team1, team2}
}

func main() {
	arr := []int{50, 60, 60, 45, 70}
	fmt.Println(solution(arr))
}
