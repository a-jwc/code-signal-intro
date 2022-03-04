package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func solution(name string) bool {
	if unicode.IsDigit(rune(name[0])) {
		return false
	}
	match, _ := regexp.MatchString(`^[\w\d_]+$`, name)
	return match
}

func main() {
	name := "sq_q"
	fmt.Println(solution(name))
}
