package main

import (
	"fmt"
)

func solution(s string) int {
	count := 0
	var arr []rune
	for _, v := range s {
		if !check(arr, v) {
			count++
			arr = append(arr, v)
		}
	}
	return count
}

func check(arr []rune, x rune) bool {
	for _, v := range arr {
		if v == x {
			return true
		}
	}
	return false
}

func main() {
	s := "cabca"
	fmt.Println(solution(s))
}
