package main

import (
	"fmt"
)

func solution(a []int, b []int) bool {
	swapped := false

	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] && i == len(a) - 1 {
			return false
		}
		if v != b[i] {
			if swapped {
				return false
			}
			b = append(b[:i], swap(b[i:], a[i:], v)...)
			swapped = true
		}
		if comp(a, b) {
			return true
		}
	}

	return true
}

func swap(a []int, b []int, value int) []int {
	firstIndex := 0
	for secondIndex, v := range a {
		if value == v {
			temp := a[firstIndex]
			a[firstIndex] = value
			a[secondIndex] = temp
			if comp(a, b) {
				return a
			} else {
				a[secondIndex] = value
				a[firstIndex] = temp
				continue
			}
		}
	}
	return a
}

func comp(a []int, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// b := []int{1, 2, 3, 4, 5, 2}
	// a := []int{2, 2, 3, 4, 5, 1}
	c := []int{832, 998, 148, 570, 533, 561, 894, 147, 455, 279}
	d := []int{832, 998, 148, 570, 533, 561, 455, 147, 894, 279}
	fmt.Println(solution(c, d))
}
