package main

import (
	"fmt"
)

func solution(inputString string) string {
	retString := ""
	stringToReverse := ""
	rev := false
	for i := 0; i < len(inputString); i++ {
		char := string(inputString[i])
		fmt.Println(char)
		if char == ")" {
			rev = false
			retString = retString + stringToReverse
			stringToReverse = ""
		}
		if rev {
			stringToReverse = char + stringToReverse
		} else if char != "(" && char != ")" {
			retString = retString + char
		}
		if char == "(" {
			rev = true
		}
	}

	return retString
}

func main() {
	fmt.Println(solution("foo(bar(baz))blim"))
}
