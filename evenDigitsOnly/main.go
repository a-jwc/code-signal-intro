package main

import (
	"fmt"
	"strconv"
)

func solution(n int) bool {
	intAsString := strconv.Itoa(n)
	for _, v := range intAsString {
		value := int(v - '0')
		if !(value % 2 == 0) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(solution(24582))
}