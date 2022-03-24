package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func solution(n int) int {

	regex := regexp.MustCompile(`^[0-9]$`)
	return check(strconv.Itoa(n), regex)
}

func check(n string, regex *regexp.Regexp) int {
	sum := 0

	if regex.MatchString(n) {
		return 0
	}
	for _, v := range n {
		sum += int(v - '0')
	}
	return check(strconv.Itoa(sum), regex) + 1
}

func main() {
	fmt.Println(solution(91))
}
