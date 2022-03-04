package main

import (
	"fmt"
	"strconv"
)

func solution(cell1 string, cell2 string) bool {
	c1 := cell1[0] - 64
	c2 := cell2[0] - 64

	c3, _ := strconv.Atoi(string(cell1[1]))
	c4, _ := strconv.Atoi(string(cell2[1]))

	if ((c1 - c2) % 2 != 0 ) && ((c3 - c4) % 2 == 0) {
		return false
	} else if ((c1 - c2) % 2 == 0 ) && ((c3 - c4) % 2 != 0) {
		return false
	}

	return true
}

func main() {
	cell1 := "B2"
	cell2 := "A1"
	fmt.Println(solution(cell1, cell2))
}