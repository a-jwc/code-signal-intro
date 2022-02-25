package main

import (
	"fmt"
)

func solution(s1 string, s2 string) int {
	s1Map := make(map[rune]int)
	s2Map := make(map[rune]int)
	count := 0

	for _, v := range s1 {
		s1Map[v] += 1
	}
	for _, v := range s2 {
		s2Map[v] += 1
	}
	for i := range s1Map {
		s2Ele, s2Ok := s2Map[i]
		s1Ele, s1Ok := s1Map[i]
		if s2Ok && s1Ok {
			if s1Ele >= s2Ele {
				count += s2Ele
			} else {
				count += s1Ele
			}
		}
	}

	return count
}

func main() {
	s1 := "aabcc"
	s2 := "adcaa"
	fmt.Println(solution(s1, s2))
}
