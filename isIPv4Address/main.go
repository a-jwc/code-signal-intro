package main

import (
	"fmt"
	"regexp"
	"strings"
)

func solution(inputString string) bool {
	res := strings.Split(inputString, ".")
	if len(res) != 4 {
		return false
	}
	// regex, _ := regexp.Compile(`\b([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\b`)
	for _, v := range res {
		match, _ := regexp.MatchString(`\b([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\b`, v)
		if !match {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(solution("1.0.0.0"))
}